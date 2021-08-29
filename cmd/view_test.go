package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"pairinator/mocks"
	"pairinator/models"
	"testing"
)

func Test_ErrorNoMembersAdded(t *testing.T) {
	database := &mocks.MemberDatabase{}
	database.On("GetAll").Return([]models.Member{})

	cmd := viewCmd(database)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Equal(t, "No members added to pairinator!\n", stdout)
}

func Test_ShowAllMembers(t *testing.T) {
	md := &mocks.MemberDatabase{}
	zoidberg := models.NewMember("Zoidberg")
	bender := models.NewMember("Bender")
	md.On("GetAll").Return([]models.Member{zoidberg, bender})

	cmd := viewCmd(md)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Contains(t, stdout, "Bender")
	assert.Contains(t, stdout, "Zoidberg")
}