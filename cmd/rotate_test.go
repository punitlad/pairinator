package cmd

import (
	"github.com/kami-zh/go-capturer"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"pairinator/mocks"
	"pairinator/models"
	"testing"
)

func Test_CannotRotateLessThan3Members(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("GetAll").Return([]models.Member{})
	cmd := rotateCmd(mockDb)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Equal(t, "Not enough members to rotate! Make sure to add more than 3 members.\n", stdout)
}

func Test_RotateMembers(t *testing.T) {
	mockDb := &mocks.MemberDatabase{}
	mockDb.On("GetAll").Return([]models.Member{models.NewMember("Bender"), models.NewMember("Zoidberg"), models.NewMember("Fry"), models.NewMember("Leela")})
	mockDb.On("Update", mock.Anything).Return()
	cmd := rotateCmd(mockDb)
	stdout := capturer.CaptureStdout(func() {
		cmd.ExecuteC()
	})

	assert.Equal(t, "Pair Swapped! Bender-Zoidberg\nPair Swapped! Fry-Leela\n", stdout)
	mockDb.AssertCalled(t, "Update", models.Member{Name: "Bender", PairedWith: []string{"Zoidberg"}})
	mockDb.AssertCalled(t, "Update", models.Member{Name: "Zoidberg", PairedWith: []string{"Bender"}})
	mockDb.AssertCalled(t, "Update", models.Member{Name: "Fry", PairedWith: []string{"Leela"}})
	mockDb.AssertCalled(t, "Update", models.Member{Name: "Leela", PairedWith: []string{"Fry"}})
}