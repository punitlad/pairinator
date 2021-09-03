package models

import (
	"errors"
	"pairinator/util"
)

type PairStair struct {
	Members []Member
}

func NewPairStair(members []Member) (PairStair, error) {
	if len(members) < 4 {
		return PairStair{}, errors.New("Not enough members to create pair stair! Make sure to add more than 3 members.")
	}

	return PairStair{Members: members}, nil
}

func (stair *PairStair) Rotate() {
	var assigned []string

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