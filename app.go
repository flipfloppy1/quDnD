package main

import (
	"context"
	"encoding/json"
	"fmt"
)

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

func (a *App) GeneratePage(key string) PageInfo {
	resp := qudAction("action=query&prop=category&titles=" + key)
	if resp == "null" {
		return PageInfo{"Creatures"}
	}
	return PageInfo{"Creatures"}
}
