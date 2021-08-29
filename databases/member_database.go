package databases

import "pairinator/models"

type MemberDatabase interface {
	Add(member models.Member)
	GetAll() []models.Member
}
