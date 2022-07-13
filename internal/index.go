package internal

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

func (i *MemIndex) Delete(indexEntry IndexEntry) {
	delete(*i.IndexEntryMap, indexEntry.Key)
}

func (i *MemIndex) Put(indexEntry IndexEntry) {
	m := *i.IndexEntryMap
	m[indexEntry.Key] = *indexEntry.Entey
}

func (i *MemIndex) Search(indexEntry IndexEntry) Entry {
	m := *i.IndexEntryMap
	entry := m[indexEntry.Key]
	return entry
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
	var entrySlice []Entry
	if err := json.Unmarshal(b, &entrySlice); err != nil {
		return err
	}
	indexMap := make(map[string]Entry)
	for _, e := range entrySlice {
		indexMap[string(e.Key)] = e
	}
	i.IndexEntryMap = &indexMap
	return nil
}

// save IndexEntry to file
func (i *MemIndex) Save(path string) error {
	if err := CreateFileIfNotExists(path); err != nil {
		return err
	}
	entrySlice := []Entry{}
	for _, v := range *i.IndexEntryMap {
		entrySlice = append(entrySlice, v)
	}
	b, err := json.Marshal(entrySlice)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0755)
}
