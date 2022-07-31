package unittest

import (
	"fmt"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func TestGetDataFiles(t *testing.T) {
	var ids []int
	fns, err := filepath.Glob(fmt.Sprintf("/Users/admin/Downloads/*.csv"))
	if err != nil {
		t.Error(err)
	}
	sort.Strings(fns)
	for _, fn := range fns {
		fn = filepath.Base(fn)
		ext := filepath.Ext(fn)
		if ext != ".csv" {
			continue
		}
		println(fn, ext)
		id, err := strconv.ParseInt(strings.TrimSuffix(fn, ext), 10, 32)
		println(id)
		if err != nil {
			t.Error(err)
		}
		ids = append(ids, int(id))
	}
	sort.Ints(ids)

}
