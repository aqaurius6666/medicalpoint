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
	if searchUser.UserID != nil {
		db = db.Where(`"users"."user_id" = ?`, searchUser.UserID)
	}
	return db
}

func (u *UserCDBRepo) CreateUser(value *user.User) (*user.User, error) {
	if err := u.Db.Create(value).Error; err != nil {
		return nil, user.ErrCreateFail
	}
	return value, nil
}

func (u *UserCDBRepo) GetUser(search *user.SearchUser) (*user.User, error) {
	var value user.User
	if err := applySearchUser(search, u.Db).First(&value).Error; err != nil {
		return nil, user.ErrNotFound
	}
	return &value, nil
}
