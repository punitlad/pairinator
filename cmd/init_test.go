package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/buntdb"
	"testing"
)

func Test_ExecuteInitCommand(t *testing.T) {
	cmd := initCmd()
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Equal(t, "Created pairinator database\n", stdout)
	assert.FileExists(t, "pairinator.db")
	_, dbError := buntdb.Open("pairinator.db")
	assert.Nil(t, dbError)
}