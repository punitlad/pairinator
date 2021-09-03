package models

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RotateMembers(t *testing.T) {
	stair, _ := NewPairStair(
		[]Member{NewMember("Bender"), NewMember("Zoidberg"), NewMember("Fry"), NewMember("Leela")},
	)

	stair.Rotate()

	assert.Contains(t, stair.Members, Member{"Bender", []string{"Zoidberg"}})
	assert.Contains(t, stair.Members, Member{"Zoidberg", []string{"Bender"}})
	assert.Contains(t, stair.Members, Member{"Fry", []string{"Leela"}})
	assert.Contains(t, stair.Members, Member{"Leela", []string{"Fry"}})
}

func Test_CannotCreateStairLessThanFourMembers(t *testing.T) {
	_, err := NewPairStair(
		[]Member{NewMember("Zoidberg"), NewMember("Fry"), NewMember("Leela")},
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "Not enough members to create pair stair! Make sure to add more than 3 members.")
}
