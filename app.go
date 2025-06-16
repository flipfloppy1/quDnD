package main

import (
	"context"

	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"quDnD/src/pageUtils"
	"quDnD/src/statblock"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/html"
)

type PageInfo struct {
	PageType    pageUtils.Screen     `json:"pageType"`
	PageTitle   string               `json:"pageTitle"`
	ImgLink     *string              `json:"imgSrc"`
	Description *string              `json:"description"`
	Statblock   *statblock.Statblock `json:"statblock"`
	PageId      int                  `json:"pageid"`
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
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
		fmt.Println(err.Error())
	}
	return results
}

func (a *App) GeneratePage(pageid int) PageInfo {
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

	return PageInfo{pageUtils.Screen(category), resp.Parse.Title, imgLink, description, pageSb, pageid}
}
