package user

import "github.com/sonntuet1997/medical-chain-utils/cockroach"

type UserRepo interface {
	cockroach.CommonDataService
}
