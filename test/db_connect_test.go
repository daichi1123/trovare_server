package test

import (
	"go_api/db"
	"go_api/test/Helper"
	"testing"
)

func TestDbConnect(t *testing.T) {
	t.Run("Connect DB", func(t *testing.T) {
		got := db.DbConnect()
		want := true
		Helper.AssertCorrectBool(t, got, want)
	})
}
