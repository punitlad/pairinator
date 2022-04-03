package views

import (
	"fmt"
	"pairinator/models"
	"strings"
)

type CurrentPairsView struct {
	stair models.PairStair
	output string
}

func NewCurrentPairsView(stair models.PairStair) CurrentPairsView {
	return CurrentPairsView{
		stair, "",
	}
}

func (v CurrentPairsView) ToString() string {
	for _, member := range v.stair.Members {
		containsPairOutput := strings.Contains(v.output, fmt.Sprintf("%s pairs with %s", member.CurrentPair, member.Name))
		if !containsPairOutput {
			v.outputBasedOnPair(member)

		}
	}

	return v.output
}

func (v *CurrentPairsView) outputBasedOnPair(pair models.Member) {
	if len(pair.CurrentPair) == 0 {
		v.output = fmt.Sprintf("%s%s is solo-ing\n", v.output, pair.Name)
	} else {
		v.output = fmt.Sprintf("%s%s pairs with %s\n", v.output, pair.Name, pair.CurrentPair)
	}
}
