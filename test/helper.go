package test

import (
	"fmt"
	"path/filepath"
	"runtime"
	"testing"
)

func getFilepath(name string) string {
	_, b, _, _ := runtime.Caller(0)
	dir := filepath.Dir(b)
	return fmt.Sprintf("file://%s/testdata/%s", dir, name)
}

func check(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatal(err)
	}
}
