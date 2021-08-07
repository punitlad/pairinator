package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_AddsNewMember(t *testing.T) {
	cmd := addCmd()
	stdout := capturer.CaptureStdout(func() {
		cmd.SetArgs([]string{"Bender"})
		cmd.ExecuteC()
	})

	assert.Equal(t, "Successfully added pair member Bender.\n", stdout)

	stdout = capturer.CaptureStdout(func() {
		cmd.SetArgs([]string{"Zoidberg"})
		cmd.ExecuteC()
	})

	assert.Equal(t, "Successfully added pair member Zoidberg.\n", stdout)
}
