package databases

import (
	"encoding/json"
	"github.com/tidwall/buntdb"
	"pairinator/models"
)

type MemberBuntDB struct {
	db *buntdb.DB
}

func GetMemberDatabaseInstance(dbPath string) MemberBuntDB {
	md := MemberBuntDB{}
	md.db, _ = buntdb.Open(dbPath)
	return md
}

func (md MemberBuntDB) Add(member models.Member) {
	memberJson, _ := json.Marshal(member)

	md.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(member.Name, string(memberJson), nil)
		return err
	})
}

func (md MemberBuntDB) GetAll() []models.Member {
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

func (md MemberBuntDB) Update(member models.Member) {
	memberJson, _ := json.Marshal(member)

	md.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(member.Name, string(memberJson), nil)
		return err
	})
}