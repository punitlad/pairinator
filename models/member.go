package models

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