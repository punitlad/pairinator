package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_ReturnTrueWhenValueExistsInArray(t *testing.T) {
	assert.True(t, Contains([]string{"Hello"}, "Hello"))
}

func Test_ReturnFalseWhenValueDoesNotExistsInArray(t *testing.T) {
	assert.False(t, Contains([]string{"Hello"}, "IAmNotThere"))
}
