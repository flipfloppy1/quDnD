package pageUtils

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

type Categories struct {
	ctx context.Context
}

var categoryMap map[string][]int

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

type CategoryMembersJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Continue      *struct {
		Cmcontinue string `json:"cmcontinue"`
		Continue   string `json:"continue"`
	} `json:"continue,omitempty"`
	Query struct {
		Categorymembers []struct {
			Pageid int    `json:"pageid"`
			Ns     int    `json:"ns"`
			Title  string `json:"title"`
		} `json:"categorymembers"`
	} `json:"query"`
}

type PageData struct {
	Pageid    int    `json:"pageid"`
	Namespace int    `json:"ns"`
	Title     string `json:"title"`
}

type RestPagesResultJson struct {
	Batchcomplete string `json:"batchcomplete"`
	Query         struct {
		PageMap map[string]PageData `json:"pages"`
	} `json:"query"`
}

func QudRest(endpoint string) string {
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

func QudAction(params string) string {
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

func GetCategory(category string) []int {
	var resp CategoryMembersJson
	json.Unmarshal([]byte(QudAction("action=query&list=categorymembers&cmtitle="+category+"&cmlimit=max")), &resp)
	members := []int{}
	i := 0
	for {
		if len(resp.Query.Categorymembers) == i {
			break
		}
		member := resp.Query.Categorymembers[i]
		if member.Ns == 14 {
			members = append(members, GetCategory(member.Title)...)
		} else if member.Ns == 0 {
			members = append(members, member.Pageid)
		}
		i++
	}
	return members
}

type CategoryMembers struct {
	Liquids   []int `json:"liquids"`
	Creatures []int `json:"creatures"`
	Items     []int `json:"items"`
	Character []int `json:"characters"`
	Concepts  []int `json:"concepts"`
	World     []int `json:"world"`
	Mechanics []int `json:"mechanics"`
	Mutations []int `json:"mutations"`
}

func (c *Categories) LoadCategories() CategoryMembers {
	fmt.Println("Loading categories...")
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("cacheDir error")
		fmt.Println(err.Error())
	}
	os.MkdirAll(filepath.Join(cacheDir, "quDnDFiles"), os.FileMode(0777))
	f, err := os.OpenFile(filepath.Join(cacheDir, "quDnDFiles", "pageCache.json"), os.O_RDWR, os.FileMode(0777))
	if err != nil {
		categoryMap = make(map[string][]int)
		categoryMap["liquids"] = GetCategory("Category:Liquids")
		categoryMap["creatures"] = GetCategory("Category:Creatures")
		categoryMap["items"] = GetCategory("Category:Items")
		categoryMap["character"] = GetCategory("Category:Character")
		categoryMap["concepts"] = GetCategory("Category:Concepts")
		categoryMap["world"] = GetCategory("Category:World")
		categoryMap["mechanics"] = GetCategory("Category:Mechanics")
		categoryMap["mutations"] = GetCategory("Category:Mutations")

		f, err = os.Create(filepath.Join(cacheDir, "quDnDFiles", "pageCache.json"))
		if err == nil {
			_, _ = f.Seek(0, 0)
			bytes, err := json.Marshal(categoryMap)
			if err != nil {
				fmt.Println("marshal result error")
				fmt.Println(err.Error())
			} else {
				f.Write(bytes)
			}

			f.Close()
		} else {
			fmt.Println("create error")
		}
		jsonStr, _ := json.Marshal(categoryMap)
		var retVal CategoryMembers
		json.Unmarshal(jsonStr, &retVal)
		fmt.Println("Loaded categories from cache")

		return retVal

	} else {
		stat, err := f.Stat()
		if err != nil {
			fmt.Println("stat error")
			fmt.Println(err)
		}
		fileContents := make([]byte, stat.Size())
		f.Read(fileContents)
		var categoryJson CategoryMembers
		json.Unmarshal(fileContents, &categoryJson)
		f.Close()
		categoryMap = make(map[string][]int)
		categoryMap["world"] = categoryJson.World
		categoryMap["creatures"] = categoryJson.Creatures
		categoryMap["character"] = categoryJson.Character
		categoryMap["items"] = categoryJson.Items
		categoryMap["concepts"] = categoryJson.Concepts
		categoryMap["mechanics"] = categoryJson.Mechanics
		categoryMap["mutations"] = categoryJson.Mutations
		categoryMap["liquids"] = categoryJson.Liquids
		return categoryJson
	}
}

func (c *Categories) GetScreen(pageid int) Screen {
	category := GetPageCategory(pageid)
	if category == "none" {
		return Other
	} else {
		return Screen(category)
	}
}
func GetPageCategory(pageid int) string {
	for cat, ids := range categoryMap {
		if slices.Contains(ids, pageid) {
			return cat
		}
	}

	return "none"
}
