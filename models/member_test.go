package models

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateMember(t *testing.T) {
	member := NewMember("Bender")
	assert.Equal(t, member.Name, "Bender")
	assert.Empty(t, member.PairedWith)
}

func Test_MemberIsMarshallable(t *testing.T) {
	memberJson, _ := json.Marshal(NewMember("Bender"))
	assert.Equal(t, "{\"Name\":\"Bender\",\"PairedWith\":[],\"CurrentPair\":\"\"}", string(memberJson))
}

func Test_MemberIsAddedToPairList(t *testing.T) {
	member := NewMember("Bender")
	member.SetCurrentPair("Zoidberg")

	assert.Equal(t, []string{"Zoidberg"}, member.PairedWith)
	assert.Equal(t, "Zoidberg", member.CurrentPair)
}

func Test_ReturnsTrueWhenMemberHasPairedWithGivenName(t *testing.T) {
	member := NewMember("Bender")
	member.SetCurrentPair("Zoidberg")

	assert.True(t, member.HasPairedWith("Zoidberg"))
}

func Test_ReturnsTrueWhenMemberHasPairedWithSameName(t *testing.T) {
	member := NewMember("Bender")

	assert.True(t, member.HasPairedWith("Bender"))
}

func Test_ReturnsFalseWhenMemberHasNotPairedWithGivenName(t *testing.T) {
	member := NewMember("Bender")

	assert.False(t, member.HasPairedWith("Zoidberg"))
}