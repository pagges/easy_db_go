package unittest

import (
	"easy_db_go/internal"
	"fmt"
	"testing"
)

func TestIndexSave(t *testing.T) {

	indexMap := make(map[string]internal.Entry)
	indexMap["test"] = internal.Entry{0, []byte("test"), []byte("alice"), 0, nil}
	indexMap["test1"] = internal.Entry{0, []byte("test1"), []byte("alice"), 0, nil}
	indexMap["test2"] = internal.Entry{0, []byte("test2"), []byte("alice"), 0, nil}
	indexMap["test3"] = internal.Entry{0, []byte("test3"), []byte("alice"), 0, nil}
	indexMap["test4"] = internal.Entry{0, []byte("test4"), []byte("alice"), 0, nil}

	memoryIndex := internal.MemIndex{IndexEntryMap: &indexMap}
	err := memoryIndex.Save("../data/index.data")
	if err != nil {
		t.Error(err)
	}
}

func TestLoadIndex(t *testing.T) {
	indexMap := internal.MemIndex{}
	err := indexMap.Load("../data/index.data")
	if err != nil {
		t.Log(err)
	}
	entry := internal.Entry{0, []byte("test"), []byte("test"), 0, nil}
	e1, ok1 := indexMap.Search(entry.BuidIndexEntry())
	fmt.Println("result1", e1, ok1)
	indexMap.Put(entry.BuidIndexEntry())
	e2, ok2 := indexMap.Search(entry.BuidIndexEntry())
	fmt.Println("result2", e2, ok2)
	indexMap.Delete(entry.BuidIndexEntry())
	fmt.Println("resultMap", indexMap.IndexEntryMap)
}
