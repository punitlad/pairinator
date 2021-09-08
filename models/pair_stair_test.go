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

	for _, member := range stair.Members {
		assert.Len(t, member.PairedWith, 1)
		assert.NotEmpty(t, member.CurrentPair)
		assert.NotEqual(t, member.Name, member.CurrentPair)
	}
}

func Test_CannotCreateStairLessThanFourMembers(t *testing.T) {
	_, err := NewPairStair(
		[]Member{NewMember("Zoidberg"), NewMember("Fry"), NewMember("Leela")},
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "Not enough members to create pair stair! Make sure to add more than 3 members.")
}

func Test_ResetStairsWhenRotationMaxIsReached(t *testing.T) {
	members := []Member{
		{"Bender", []string{"Zoidberg", "Leela", "Fry"}, "Zoidberg"},
		{"Zoidberg", []string{"Bender", "Leela", "Fry"}, "Bender"},
		{"Leela", []string{"Zoidberg", "Bender", "Fry"}, "Fry"},
		{"Fry", []string{"Zoidberg", "Leela", "Bender"}, "Leela"},
	}
	stair, _ := NewPairStair(members)

	stair.Rotate()

	for _, member := range stair.Members {
		assert.Len(t, member.PairedWith, 1)
	}
}
