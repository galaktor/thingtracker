package main

import(
	"strconv"
	"strings"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

var things map[int]*Thing

func init() {
	t, err := deserThings()
	guard(err)
	things = t
}

func thingFileNameToId(filename string) (int, error) {
	return strconv.Atoi(strings.TrimSuffix(filename, ".thing"))
}

func deserThings() (map[int]*Thing, error) {
	files, err := ioutil.ReadDir("store")
	guard(err)

	out := make(map[int]*Thing)
	for _, file := range files {
		if file.Mode().IsRegular() {
			if filepath.Ext(file.Name()) == ".thing" {
				id, err := thingFileNameToId(file.Name())
				guard(err)
				fullpath := "store/" + file.Name()
				content, err := ioutil.ReadFile(fullpath)
				guard(err)
				t := &Thing{}
				err = json.Unmarshal(content, t)
				guard(err)
				out[id] = t
			}
		}
	}
	
	return out, err	
}

func refreshThings() (err error) {
	things, err = deserThings()
	return
}

func getNextId() int {
	out := len(things)
	occupied := true
	for ;occupied; out++ {
		if _,occupied = things[out]; !occupied {
			break
		}
	}

	return out
}
