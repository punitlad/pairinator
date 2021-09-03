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

	mockDb.AssertCalled(t, "Update", models.Member{"Bender", []string{"Zoidberg"}, "Zoidberg"})
	mockDb.AssertCalled(t, "Update", models.Member{"Zoidberg", []string{"Bender"}, "Bender"})
	mockDb.AssertCalled(t, "Update", models.Member{"Fry", []string{"Leela"}, "Leela"})
	mockDb.AssertCalled(t, "Update", models.Member{"Leela", []string{"Fry"}, "Fry"})
}

func Test_PrintsOutPairsAfterRotation(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("GetAll").Return([]models.Member{models.NewMember("Bender"), models.NewMember("Zoidberg"), models.NewMember("Fry"), models.NewMember("Leela")})
	mockDb.On("Update", mock.Anything).Return()
	cmd := rotateCmd(mockDb)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Contains(t, stdout, "Bender pairs with Zoidberg")
	assert.Contains(t, stdout, "Zoidberg pairs with Bender")
	assert.Contains(t, stdout, "Fry pairs with Leela")
	assert.Contains(t, stdout, "Leela pairs with Fry")
}