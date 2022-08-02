package unittest

import (
	"easy_db_go/internal"
	"fmt"
	"testing"
)

func TestPut(t *testing.T) {
	db, err := internal.Open()
	if err != nil {
		t.Error(err)
	}
	for i := 0; i < 10000000; i++ {
		key := fmt.Sprintf("key_%d", i)
		value := fmt.Sprintf("value_%d", i)
		db.Put([]byte(key), []byte(value))
	}
	db.Close()
}

func TestGet(t *testing.T) {
	db, err := internal.Open()
	if err != nil {
		t.Error(err)
	}
	e, err := db.Get([]byte("key_19"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s \n", e.Key)
	t.Logf("%s \n", e.Value)
	db.Close()
}

func TestDBSize(t *testing.T) {
	db, err := internal.Open()
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
	t.Log("DB SIZE: ", db.Size())
}
