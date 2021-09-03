package models

import "pairinator/util"

type Member struct {
	Name string
	PairedWith []string
}

func NewMember(name string) Member {
	return Member{name, []string{}}
}

func (member *Member) AddMemberToPairList(name string) {
	member.PairedWith = append(member.PairedWith, name)
}

func (member *Member) HasPairedWith(name string) bool {
	return member.Name == name || util.Contains(member.PairedWith, name)
}