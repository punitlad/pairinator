package models

import (
	"github.com/kami-zh/go-capturer"
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

func Test_MembersAreRotatedInRandomOrder(t *testing.T) {
	stair, _ := NewPairStair(
		[]Member{NewMember("Bender"), NewMember("Zoidberg"), NewMember("Fry"), NewMember("Leela")},
	)

	stair.Rotate()

	assert.NotEqual(t,
		[]string{"Bender", "Zoidberg", "Fry", "Leela"},
		[]string{stair.Members[0].Name, stair.Members[1].Name, stair.Members[2].Name, stair.Members[3].Name},
	)
}

func Test_CannotCreateStairLessThanFourMembers(t *testing.T) {
	_, err := NewPairStair(
		[]Member{NewMember("Zoidberg"), NewMember("Fry"), NewMember("Leela")},
	)

	assert.Error(t, err)
	assert.EqualError(t, err, "not enough members to create pair stair! Make sure to add more than 3 members")
}

func Test_ResetStairsWhenRotationMaxIsReached(t *testing.T) {
	members := []Member{
		{"Bender", []string{"Zoidberg", "Leela", "Fry"}, "Zoidberg"},
		{"Zoidberg", []string{"Bender", "Leela", "Fry"}, "Bender"},
		{"Leela", []string{"Zoidberg", "Bender", "Fry"}, "Fry"},
		{"Fry", []string{"Zoidberg", "Leela", "Bender"}, "Leela"},
	}
	stair, _ := NewPairStair(members)

	out := capturer.CaptureStdout(func() {
		stair.Rotate()
	})

	assert.Equal(t, "Pair stair complete! Resetting....\n", out)
	for _, member := range stair.Members {
		assert.Len(t, member.PairedWith, 1)
	}
}

func Test_SetsCurrentPairToEmptyForOneMemberWhenThereAreOddNumberOfMembers(t *testing.T) {
	members := []Member{
		NewMember("Bender"),NewMember("Zoidberg"),NewMember("Leela"),NewMember("Fry"),NewMember("Professor"),
	}
	stair, _ := NewPairStair(members)
	stair.Rotate()
	assert.Equal(t, 1, numOfSolos(stair))
}

func numOfSolos(stair PairStair) int {
	numOfSolos := 0
	for _, member := range stair.Members {
		if member.CurrentPair == "" {
			numOfSolos += 1
		}
	}
	return numOfSolos
}
