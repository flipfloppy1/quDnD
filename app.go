package main

import (
	"context"
	"log"
	"os"
	"path"

	"encoding/json"
	"strconv"
	"strings"

	"quDnD/src/db"
	"quDnD/src/pageUtils"
	"quDnD/src/statblock"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

// App struct
type App struct {
	ctx    context.Context
	db     *db.DbHandler
	logger *log.Logger
}

// NewApp creates a new App application struct
func NewApp() *App {
	logger := log.New(os.Stdout, "[quDnD] ", log.Flags())
	dbHandler, err := db.NewSqliteHandler(logger)
	if err != nil {
		logger.Println("Error creating database handler:", err.Error())
	}
	conf, _ := os.UserConfigDir()
	os.MkdirAll(path.Join(conf, "quDnD", "db"), os.ModePerm)
	return &App{context.Background(), dbHandler, logger}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) SearchForPage(query string) pageUtils.RestPageSearch {
	var results pageUtils.RestPageSearch
	err := json.Unmarshal([]byte(pageUtils.QudRest("/search/page?q="+query)), &results)
	if err != nil {
		a.logger.Println(err.Error())
	}
	return results
}

func (a *App) GeneratePage(pageid int) statblock.PageInfo {
	category := pageUtils.GetPageCategory(pageid)
	var resp pageUtils.ParseHTMLPage
	json.Unmarshal([]byte(pageUtils.QudAction("action=parse&prop=text&pageid="+strconv.Itoa(pageid))), &resp)
	nodes, _ := html.ParseFragment(strings.NewReader(resp.Parse.Text.Root), nil)
	doc := goquery.NewDocumentFromNode(nodes[0])
	var pageSb *statblock.Statblock
	pageSb = statblock.ComposeStatblock(doc)
	var description *string
	description = pageUtils.GetDescription(doc)
	var imgLink *string
	imgLink = pageUtils.GetPageImg(doc)

	return statblock.PageInfo{pageUtils.Screen(category), resp.Parse.Title, imgLink, description, pageSb, pageid}
}

func (a *App) GetCachedPage(pageid int) statblock.DbPage {
	page, err := a.db.GetCachedPage(pageid)
	if err != nil {
		return statblock.DbPage{Page: statblock.PageInfo{}, Exists: false}
	}

	return statblock.DbPage{Page: page, Exists: true}
}

func (a *App) SetCachedPage(page statblock.PageInfo) error {
	return a.db.SetCachedPage(page)
}

func (a *App) GetCustomPage(pageid int) statblock.DbPage {
	page, err := a.db.GetCustomPage(pageid)
	if err != nil {
		return statblock.DbPage{Page: statblock.PageInfo{}, Exists: false}
	}

	return statblock.DbPage{Page: page, Exists: true}
}

func (a *App) SetCustomPage(page statblock.PageInfo) error {
	return a.db.SetCustomPage(page)
}

func (a *App) PageIdFromFriendly(friendlyId string) int {
	return pageUtils.GetPageIdFromFriendly(friendlyId)
}
