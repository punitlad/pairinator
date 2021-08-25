package models

type Member struct {
	Name string
	PairedWith []string
}

func NewMember(name string) Member {
	return Member{name, []string{}}
}