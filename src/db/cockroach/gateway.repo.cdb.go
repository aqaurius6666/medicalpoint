package cockroach

import (
	"github.com/medicalpoint/gateway/src/db/cockroach/user"
	iUser "github.com/medicalpoint/gateway/src/db/interface/user"
	"github.com/sonntuet1997/medical-chain-utils/cockroach"
)

type GatewayCDBRepo struct {
	cockroach.CDBRepo
	UserCDBRepo *user.UserCDBRepo
}

func (g *GatewayCDBRepo) GetUser(s *iUser.SearchUser) (*iUser.User, error) {
	return g.UserCDBRepo.GetUser(s)
}

func (g *GatewayCDBRepo) CreateUser(s *iUser.User) (*iUser.User, error) {
	return g.UserCDBRepo.CreateUser(s)
}
