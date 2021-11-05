package user

import "github.com/sonntuet1997/medical-chain-utils/common"

type User struct {
	common.ModelBase
	UserID              *string `gorm:"index;unique"`
	EncryptedPrivateKey *string `gorm:"index;unique"`
}

type SearchUser struct {
	User
}
