package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type Requests struct {
	ctx context.Context
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

func getCategory(category string) []int {
	var resp CategoryMembersJson
	json.Unmarshal([]byte(qudAction("action=query&list=categorymembers&cmtitle="+category+"&cmlimit=max")), &resp)
	members := []int{}
	i := 0
	for {
		if len(resp.Query.Categorymembers) == i {
			break
		}
		member := resp.Query.Categorymembers[i]
		if member.Ns == 14 {
			members = append(members, getCategory(member.Title)...)
		} else if member.Ns == 0 {
			members = append(members, member.Pageid)
		}
		i++
	}
	fmt.Println(members)
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
}

func (r *Requests) LoadPages() CategoryMembers {
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		fmt.Println("cacheDir error")
		fmt.Println(err.Error())
	}
	os.MkdirAll(filepath.Join(cacheDir, "quDnDFiles"), os.FileMode(0777))
	f, err := os.OpenFile(filepath.Join(cacheDir, "quDnDFiles", "pageCache.json"), os.O_RDWR, os.FileMode(0777))
	if err != nil {
		categoryMap := make(map[string][]int)
		categoryMap["liquids"] = getCategory("Category:Liquids")
		categoryMap["creatures"] = getCategory("Category:Creatures")
		categoryMap["items"] = getCategory("Category:Items")
		categoryMap["character"] = getCategory("Category:Character")
		categoryMap["concepts"] = getCategory("Category:Concepts")
		categoryMap["world"] = getCategory("Category:World")
		categoryMap["mechanics"] = getCategory("Category:Mechanics")

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
		return categoryJson
	}
}
