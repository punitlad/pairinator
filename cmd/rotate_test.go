package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pairinator/mocks"
	"pairinator/models"
	"testing"
)

func Test_RotateMembers(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("GetAll").Return([]models.Member{models.NewMember("Bender"), models.NewMember("Zoidberg"), models.NewMember("Fry"), models.NewMember("Leela")})
	mockDb.On("Update", mock.Anything).Return()
	cmd := rotateCmd(mockDb)
	capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	mockDb.AssertNumberOfCalls(t, "Update", 4)
}

func Test_PrintsOutPairsAfterRotation(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("GetAll").Return([]models.Member{models.NewMember("Bender"), models.NewMember("Zoidberg"), models.NewMember("Fry"), models.NewMember("Leela")})
	mockDb.On("Update", mock.Anything).Return()
	cmd := rotateCmd(mockDb)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.NotContains(t, stdout, "Bender pairs with Bender")
	assert.NotContains(t, stdout, "Zoidberg pairs with Zoidberg")
	assert.NotContains(t, stdout, "Fry pairs with Fry")
	assert.NotContains(t, stdout, "Leela pairs with Leela")
}