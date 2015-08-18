package main

import(
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"io"
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

func Get(id int) (*Thing, error) {
	t, found := things[id]
	if (!found) {
		return nil, errors.New("thing not found")
	}

	return t, nil
}

func Save(t *Thing) error {
	filename := fmt.Sprintf("store/%s.thing", t.Id)
	serialized, err := json.Marshal(t)
	if err != nil {
		return err
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.WriteString(file, string(serialized))
	if err != nil {
		return err
	}

	return nil
}













