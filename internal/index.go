package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func (i *IndexEntry) Delete(indexMap *map[string]Entry, indexEntry IndexEntry) {
	m := *indexMap
	delete(m, indexEntry.Key)
}

func (i *IndexEntry) Put(indexMap *map[string]Entry, indexEntry IndexEntry) {
	m := *indexMap
	m[indexEntry.Key] = *indexEntry.Entey
}

func (i *IndexEntry) Search(indexMap *map[string]Entry, indexEntry IndexEntry) Entry {
	m := *indexMap
	entry := m[indexEntry.Key]
	return entry
}

// load IndexEntry from file
func (i *IndexEntry) Load(path string) (*map[string]Entry, error) {
	if !Exists(path) {
		return nil, errors.New("index file not exists!")
	}
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var entrySlice []Entry
	if err := json.Unmarshal(b, &entrySlice); err != nil {
		return nil, err
	}
	indexMap := make(map[string]Entry)
	for _, e := range entrySlice {
		indexMap[string(e.Key)] = e
	}
	return &indexMap, nil
}

// save IndexEntry to file
func (i *IndexEntry) Save(path string, indexMap *map[string]Entry) error {
	if err := CreateFileIfNotExists(path); err != nil {
		return err
	}
	iMap := *indexMap
	entrySlice := []Entry{}
	for _, v := range iMap {
		entrySlice = append(entrySlice, v)
	}
	b, err := json.Marshal(entrySlice)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0755)
}
