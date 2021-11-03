package user

import (
	"github.com/medicalpoint/gateway/src/db/interface/user"

	"github.com/sonntuet1997/medical-chain-utils/cockroach"
	"gorm.io/gorm"
)

var (
	_ user.UserRepo = (*UserCDBRepo)(nil)
)

type UserCDBRepo struct {
	cockroach.CDBRepo
}

func applySearchUser(searchUser *user.SearchUser, db *gorm.DB) *gorm.DB {
	// if searchUser.ID != uuid.Nil {
	// 	db = db.Where(`"users"."id" = ?`, searchUser.ID)
	// }

	return db
}
