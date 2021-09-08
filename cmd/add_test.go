package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pairinator/mocks"
	"testing"
)

func Test_AddsNewMember(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("Add", mock.Anything).Return()

	cmd := addCmd(mockDb)
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
