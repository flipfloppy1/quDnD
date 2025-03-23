package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func qudRest(endpoint string) string {
	resp, err := http.Get("https://wiki.cavesofqud.com/rest.php/v1" + endpoint)
	if err != nil {
		return "null"
	}
	buf := new(strings.Builder)
	_, copyErr := io.Copy(buf, resp.Body)
	if copyErr != nil {
		return "null"
	}
	return buf.String()
}

func qudAction(params string) string {
	resp, err := http.Get("https://wiki.cavesofqud.com/api.php?format=json&" + params)
	if err != nil {
		return "null"
	}
	buf := new(strings.Builder)
	_, copyErr := io.Copy(buf, resp.Body)
	if copyErr != nil {
		return "null"
	}
	return buf.String()
}

// App struct
type App struct {
	ctx context.Context
}

var categoryMap CategoryMembers

type categoryQueryMember struct {
	pageid int
	ns     int
	title  string
}

type categoryJsonQuery struct {
	categorymembers []categoryQueryMember
}

type categoryJson struct {
	query categoryJsonQuery
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

func (a *App) SearchForPage(query string) any {
	var results map[string]any
	err := json.Unmarshal([]byte(qudRest("/search/page?q="+query)), &results)
	if err != nil {
		fmt.Println(err.Error())
	}
	return results["pages"]
}

func getCategory(category string) []int {
	var resp categoryJson
	json.Unmarshal([]byte(qudAction("action=query&list=categorymembers&cmtitles="+category+"&cmlimit=max")), &resp)
	members := []int{}
	i := 0
	for {
		if len(resp.query.categorymembers) == i {
			break
		}
		member := resp.query.categorymembers[i]
		if member.ns == 14 {
			members = append(members, getCategory(member.title)...)
		} else if member.ns == 0 {
			members = append(members, member.pageid)
		}
		i++
	}
	return members
}

func (a *App) LoadPages() CategoryMembers {
	categoryMap.categoryMap = make(map[string][]int)
	// categoryMap.categoryMap["liquids"] = getCategory("Category:Liquids")
	// categoryMap.categoryMap["creatures"] = getCategory("Category:Creatures")
	// categoryMap.categoryMap["items"] = getCategory("Category:Items")
	// categoryMap.categoryMap["character"] = getCategory("Category:Character")
	// categoryMap.categoryMap["concepts"] = getCategory("Category:Concepts")
	// categoryMap.categoryMap["world"] = getCategory("Category:World")
	// categoryMap.categoryMap["mechanics"] = getCategory("Category:Mechanics")

	return categoryMap
}

func (a *App) GeneratePage(key string) PageInfo {
	resp := qudAction("action=query&prop=category&titles=" + key)
	if resp == "null" {
		return PageInfo{"Creatures"}
	}
	return PageInfo{"Creatures"}
}
