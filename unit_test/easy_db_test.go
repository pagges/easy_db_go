package unittest

import (
	"easy_db_go/internal"
	"testing"
)

func TestPut(t *testing.T) {
	db, err := internal.Open()
	if err != nil {
		t.Error(err)
	}
	db.Put([]byte("test1"), []byte("test1"))
	db.Close()
}

func TestGet(t *testing.T) {
	db, err := internal.Open()
	if err != nil {
		t.Error(err)
	}
	e, err := db.Get([]byte("test"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s \n", e.Key)
	t.Logf("%s \n", e.Value)
	db.Close()
}
