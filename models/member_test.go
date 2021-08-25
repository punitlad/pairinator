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
	assert.Equal(t, "{\"Name\":\"Bender\",\"PairedWith\":[]}", string(memberJson))
}