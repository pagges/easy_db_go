package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func (i *MemIndex) Delete(item Item) {
	m := *i.IndexEntryMap
	delete(m, item.Key)
}

func (i *MemIndex) Put(item Item) {
	m := *i.IndexEntryMap
	m[item.Key] = item
}

func (i *MemIndex) Search(key string) (Item, bool) {
	m := *i.IndexEntryMap
	if _, ok := m[key]; !ok {
		return Item{}, false
	}
	item := m[key]
	return item, true
}

// load IndexEntry from file
func (i *MemIndex) Load(path string) error {
	if !Exists(path) {
		return errors.New("index file not exists!")
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	indexMap := make(map[string]Item)
	if len(b) == 0 {
		i.IndexEntryMap = &indexMap
		return nil
	}
	var itemSlice []Item
	if err := json.Unmarshal(b, &itemSlice); err != nil {
		return err
	}
	for _, item := range itemSlice {
		indexMap[string(item.Key)] = item
	}
	i.IndexEntryMap = &indexMap
	return nil
}

// save IndexEntry to file
func (i *MemIndex) Save(path string) error {
	if err := CreateFileIfNotExists(path); err != nil {
		return err
	}
	itemSlice := []Item{}
	for _, v := range *i.IndexEntryMap {
		itemSlice = append(itemSlice, v)
	}
	b, err := json.Marshal(itemSlice)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0755)
}
