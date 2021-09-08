package models

import (
	"errors"
	"fmt"
	"pairinator/util"
)

type PairStair struct {
	Members []Member
	maxRotations int
}

func NewPairStair(members []Member) (PairStair, error) {
	if len(members) < 4 {
		return PairStair{}, errors.New("not enough members to create pair stair! Make sure to add more than 3 members")
	}

	return PairStair{members, len(members) - 1}, nil
}

func (stair *PairStair) Rotate() {
	var assigned []string

	if shouldReset(stair) {
		fmt.Println("Pair stair complete! Resetting....")
		reset(stair)
	}

	for i, member := range stair.Members {
		if !util.Contains(assigned, member.Name) {
			for j, availablePair := range stair.Members {
				if !util.Contains(assigned, availablePair.Name) {
					if !member.HasPairedWith(availablePair.Name) {
						availablePair.SetCurrentPair(member.Name)
						stair.Members[j] = availablePair
						member.SetCurrentPair(availablePair.Name)
						stair.Members[i] = member
						assigned = append(assigned, member.Name, availablePair.Name)
						break
					}
				}
			}
		}
	}
}

func reset(stair *PairStair) {
	for i, member := range stair.Members {
		member.Reset()
		stair.Members[i] = member
	}
}

func shouldReset(stair *PairStair) bool {
	shouldReset := false
	for _, member := range stair.Members {
		if len(member.PairedWith) == stair.maxRotations {
			shouldReset = true
		}
	}

	return shouldReset
}