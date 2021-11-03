package cockroach

import (
	"github.com/medicalpoint/gateway/src/db/cockroach/user"
	"github.com/sonntuet1997/medical-chain-utils/cockroach"
)

type GatewayCDBRepo struct {
	cockroach.CDBRepo
	UserCDBRepo *user.UserCDBRepo
}
