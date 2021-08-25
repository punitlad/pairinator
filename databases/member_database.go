package databases

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/tidwall/buntdb"
	"pairinator/models"
)

type MemberDatabase struct {
	db *buntdb.DB
}

func GetMemberDatabaseInstance() MemberDatabase {
	md := MemberDatabase{}
	md.db, _ = buntdb.Open("pairinator.db")
	return md
}

func (md *MemberDatabase) Add(member models.Member) {
	memberJson, _ := json.Marshal(member)

	md.db.Update(func(tx *buntdb.Tx) error {
		u := uuid.New()
		_, _, err := tx.Set(u.String(), string(memberJson), nil)
		return err
	})
}

func (md *MemberDatabase) GetAll() []models.Member {
	var members []models.Member
	md.db.View(func(tx *buntdb.Tx) error {
		return tx.Ascend("", func(key, value string) bool {
			var member models.Member
			json.Unmarshal([]byte(value), &member)
			members = append(members, member)
			return true
		})
	})

	return members
}