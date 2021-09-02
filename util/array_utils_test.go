package util

import (
	"github.com/stretchr/testify/assert"
	"pairinator/models"
	"testing"
)

func Test_ReturnTrueWhenValueExistsInArray(t *testing.T) {
	assert.True(t, Contains([]string{"Hello"}, "Hello"))
}

func Test_ReturnFalseWhenValueDoesNotExistsInArray(t *testing.T) {
	assert.False(t, Contains([]string{"Hello"}, "IAmNotThere"))
}

func Test_ReturnArrayWithSpecifiedElementRemoved(t *testing.T) {
	actual := RemoveOne([]models.Member{models.NewMember("Hello")}, 0)
	assert.Equal(t, []models.Member{}, actual)
}

func Test_ReturnArrayWithSpecifiedElementsRemoved(t *testing.T) {
	actual := RemoveTwo([]models.Member{models.NewMember("Hello"), models.NewMember("Person")}, 0, 1)
	assert.Equal(t, []models.Member{}, actual)
}
