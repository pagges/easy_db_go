package unittest

import (
	"easy_db_go/internal"
	"encoding/json"
	"io/ioutil"
	"testing"
)

func TestIndexSave(t *testing.T) {
	indexEntry := internal.IndexEntry{}
	indexMap := make(map[string]internal.Entry)
	indexMap["test"] = internal.Entry{0, []byte("test"), []byte("alice"), 0, nil}
	indexMap["test1"] = internal.Entry{0, []byte("test1"), []byte("alice"), 0, nil}
	indexMap["test2"] = internal.Entry{0, []byte("test2"), []byte("alice"), 0, nil}
	indexMap["test3"] = internal.Entry{0, []byte("test3"), []byte("alice"), 0, nil}
	indexMap["test4"] = internal.Entry{0, []byte("test4"), []byte("alice"), 0, nil}

	err := indexEntry.Save("../data/index.data", &indexMap)
	if err != nil {
		t.Error(err)
	}
}

func TestLoadIndex(t *testing.T) {
	indexEntry := internal.IndexEntry{}
	indexMap, err := indexEntry.Load("../data/index.data")
	if err != nil {
		t.Log(err)
	}
	index := *indexMap
	t.Logf("type = %T", index)
	t.Log(index)
}

func SaveJson(path string, v interface{}) error {
	if err := internal.CreateFileIfNotExists(path); err != nil {
		return err
	}

	b, err := json.Marshal(v)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, b, 0755)
}
func TestPutIndex(t *testing.T) {
	indexMap := make(map[string]internal.Entry)
	entry := internal.Entry{0, []byte("test"), []byte("alice"), 0, nil}
	indexEntry := internal.BuidIndexEntry(entry)
	indexEntry.Put(&indexMap, indexEntry)
	e := indexEntry.Search(&indexMap, indexEntry)
	t.Log(e)
}
