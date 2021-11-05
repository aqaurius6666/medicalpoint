package cosmos

import "golang.org/x/xerrors"

var (
	ErrBlockchainCreateServiceFail     = xerrors.New("ERROR.BLOCKCHAIN.CREATE_SERVICE_FAILED")
	ErrBlockchainCreateUserFail        = xerrors.New("ERROR.BLOCKCHAIN.CREATE_USER_FAILED")
	ErrBlockchainCreateSharingFail     = xerrors.New("ERROR.BLOCKCHAIN.CREATE_SHARING_FAILED")
	ErrBlockchainDisconnectServiceFail = xerrors.New("ERROR.BLOCKCHAIN.DISCONNECT_SERVICE_FAILED")
	ErrBlockchainServiceUnavailable    = xerrors.New("ERROR.BLOCKCHAIN.SERVICE_UNAVAILABLE")
	ErrBlockchainUpdateSharingFail     = xerrors.New("ERROR.BLOCKCHAIN.UPDATE_SHARING_FAILED")
	ErrBlockchainBanUserFail           = xerrors.New("ERROR.BLOCKCHAIN.BAN_USER_FAILED")
	ErrBlockchainUnbanUserFail         = xerrors.New("ERROR.BLOCKCHAIN.UNBAN_USER_FAILED")
	ErrBlockchainUpdateUserFail        = xerrors.New("ERROR.BLOCKCHAIN.UPDATE_USER_FAILED")
)
