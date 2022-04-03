package views

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"pairinator/models"
	"strings"
	"testing"
)

func Test_ToStringsCurrentPairSets(t *testing.T) {
	stair, _ := models.NewPairStair([]models.Member{
		{"Bender", []string{}, "Zoidberg"},
		{"Zoidberg", []string{}, "Bender"},
		{"Leela", []string{}, "Fry"},
		{"Fry", []string{}, "Leela"},
	})

	pairsView := NewCurrentPairsView(stair)
	view := pairsView.ToString()

	assert.NotEqual(t,
		strings.Contains(view, "Bender pairs with Zoidberg"),
		strings.Contains(view, "Zoidberg pairs with Bender"),
		fmt.Sprintf("%s should only contain one statement of Zoidberg Bender pair", view),
	)

	assert.NotEqual(t,
		strings.Contains(view, "Leela pairs with Fry"),
		strings.Contains(view, "Fry pairs with Leela"),
		fmt.Sprintf("%s should only contains one statement of Fry Leela pair", view),
	)
}

func Test_ToStringsWithSoloPersonWhenOddNumberPresent(t *testing.T) {
	stair, _ := models.NewPairStair([]models.Member{
		{"Bender", []string{}, "Zoidberg"},
		{"Zoidberg", []string{}, "Bender"},
		{"Leela", []string{}, "Fry"},
		{"Fry", []string{}, "Leela"},
		{"Professor", []string{}, ""},
	})

	view := NewCurrentPairsView(stair).ToString()

	assert.Contains(t, view, "Professor is solo-ing")
}
