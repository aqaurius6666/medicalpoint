package api

import (
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/medicalpoint/gateway/src/db"
	"github.com/medicalpoint/gateway/src/db/interface/user"
	e "github.com/medicalpoint/gateway/src/lib/error"
	"github.com/medicalpoint/gateway/src/pb/api"
	"github.com/medicalpoint/gateway/src/pb/types"
	"github.com/medicalpoint/gateway/src/services/cosmos"
	"golang.org/x/xerrors"

	"github.com/sirupsen/logrus"
)

type BlockchainService struct {
	db     db.GateWayServiceRepo
	chain  *cosmos.CosmosServiceClient
	logger *logrus.Logger
	key    string
}

func (b *BlockchainService) GetAdmin(req *api.GetAdminRequest) (*api.GetAdminResponse, error) {

	return &api.GetAdminResponse{}, nil
}

func (b *BlockchainService) GetBalance(req *api.GetBalanceRequest) (*api.GetBalanceResponse, error) {
	if req.UserId == "" {
		return nil, xerrors.Errorf("%w", e.ErrQueryInvalid)
	}
	user, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	privKey, err := b.chain.Decrypt(*user.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	out, err := b.chain.QueryGetAllBalance(&bank.QueryAllBalancesRequest{
		Address: b.chain.GetAddress(privKey),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	balances := make([]*api.GetBalanceResponse_Point, 0)
	for i := 0; i < out.Balances.Len(); i++ {
		denom := out.Balances.GetDenomByIndex(i)
		balances = append(balances, &api.GetBalanceResponse_Point{
			Denom:  denom,
			Amount: out.Balances.AmountOf(denom).String(),
		})
	}
	return &api.GetBalanceResponse{
		Balances: balances,
	}, nil
}

func (b *BlockchainService) GetSystemBalance(req *api.GetSystemBalanceRequest) (*api.GetSystemBalanceResponse, error) {
	return &api.GetSystemBalanceResponse{}, nil
}

func (b BlockchainService) CreateUser(req *api.PostUserRequest) (*api.PostUserResponse, error) {
	if req.UserId == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	encryptedPrivateKey, err := b.chain.Encrypt(b.chain.GenPrivateKey(), b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	user, err := b.db.CreateUser(&user.User{
		UserID:              &req.UserId,
		EncryptedPrivateKey: &encryptedPrivateKey,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	return &api.PostUserResponse{
		UserId: *user.UserID,
	}, nil
}

func (b BlockchainService) InitSuperAdmin(req *api.PostSuperAdminRequest) (*api.PostSuperAdminResponse, error) {
	if req.UserId == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	encryptedPrivateKey, err := b.chain.Encrypt(b.chain.GenPrivateKey(), b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	user, err := b.db.CreateUser(&user.User{
		UserID:              &req.UserId,
		EncryptedPrivateKey: &encryptedPrivateKey,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostSuperAdminResponse{
		UserId: *user.UserID,
	}, nil
}

func (b BlockchainService) Mint(req *api.PostMintRequest) (*api.PostMintResponse, error) {
	if req.UserId == "" || req.Amount == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil || amount < 0 {
		return nil, xerrors.Errorf("%w", e.ErrAmountInvalid)
	}
	user, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	privKey, err := b.chain.Decrypt(*user.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	tx, err := b.chain.TxMint(privKey, &types.MsgMint{
		Creator: b.chain.GetAddress(privKey),
		Amount:  uint64(amount),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostMintResponse{
		UserId: *user.UserID,
		Amount: req.Amount,
		Txh:    tx.TxResponse.TxHash,
	}, nil
}

func (b BlockchainService) Burn(req *api.PostBurnRequest) (*api.PostBurnResponse, error) {
	if req.UserId == "" || req.Amount == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil || amount < 0 {
		return nil, xerrors.Errorf("%w", e.ErrAmountInvalid)
	}
	user, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	privKey, err := b.chain.Decrypt(*user.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	tx, err := b.chain.TxBurn(privKey, &types.MsgBurn{
		Creator: b.chain.GetAddress(privKey),
		Amount:  uint64(amount),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostBurnResponse{
		UserId: *user.UserID,
		Amount: req.Amount,
		Txh:    tx.TxResponse.TxHash,
	}, nil
}

func (b BlockchainService) Transfer(req *api.PostTransferRequest) (*api.PostTransferResponse, error) {
	if req.UserId == "" || req.Amount == "" || req.To == "" || req.Denom == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil || amount < 0 {
		return nil, xerrors.Errorf("%w", e.ErrAmountInvalid)
	}
	fromUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fromPrivateKey, err := b.chain.Decrypt(*fromUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	toUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.To},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	toPrivateKey, err := b.chain.Decrypt(*toUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	tx, err := b.chain.TxTransfer(fromPrivateKey, &bank.MsgSend{
		FromAddress: b.chain.GetAddress(fromPrivateKey),
		ToAddress:   b.chain.GetAddress(toPrivateKey),
		Amount:      sdk.Coins{sdk.NewInt64Coin(req.Denom, amount)},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostTransferResponse{
		UserId: *fromUser.UserID,
		To:     *toUser.UserID,
		Amount: req.Amount,
		Denom:  req.Denom,
		Txh:    tx.TxResponse.TxHash,
	}, nil
}

func (b BlockchainService) AdminTransfer(req *api.PostAdminTransferRequest) (*api.PostAdminTransferResponse, error) {
	if req.UserId == "" || req.Amount == "" || req.To == "" || req.Denom == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	if req.Denom != types.Denom {
		return nil, xerrors.Errorf("%w", e.ErrDenomInvalid)
	}
	amount, err := strconv.ParseInt(req.Amount, 10, 64)
	if err != nil || amount < 0 {
		return nil, xerrors.Errorf("%w", e.ErrAmountInvalid)
	}
	fromUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fromPrivateKey, err := b.chain.Decrypt(*fromUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	toUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.To},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	toPrivateKey, err := b.chain.Decrypt(*toUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	tx, err := b.chain.TxAdminTransfer(fromPrivateKey, &types.MsgAdminTransfer{
		Creator: b.chain.GetAddress(fromPrivateKey),
		Address: b.chain.GetAddress(toPrivateKey),
		Amount:  uint64(amount),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostAdminTransferResponse{
		UserId: *fromUser.UserID,
		To:     *toUser.UserID,
		Amount: req.Amount,
		Denom:  req.Denom,
		Txh:    tx.TxResponse.TxHash,
	}, nil
}

func (b BlockchainService) AddAdmin(req *api.PostAdminRequest) (*api.PostAdminResponse, error) {
	if req.UserId == "" || req.AdminId == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	fromUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fromPrivateKey, err := b.chain.Decrypt(*fromUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	privateKey := b.chain.GenPrivateKey()
	encryptedPrivateKey, err := b.chain.Encrypt(privateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	newAdmin, err := b.db.CreateUser(&user.User{
		UserID:              &req.AdminId,
		EncryptedPrivateKey: &encryptedPrivateKey,
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	tx, err := b.chain.TxCreateAdmin(fromPrivateKey, &types.MsgCreateAdmin{
		Creator: b.chain.GetAddress(fromPrivateKey),
		Address: b.chain.GetAddress(privateKey.Bytes()),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.PostAdminResponse{
		UserId:  *fromUser.UserID,
		AdminId: *newAdmin.UserID,
		Txh:     tx.TxResponse.TxHash,
	}, nil
}

func (b BlockchainService) DeleteAdmin(req *api.DeleteAdminRequest) (*api.DeleteAdminResponse, error) {
	if req.UserId == "" || req.AdminId == "" {
		return nil, xerrors.Errorf("%w", e.ErrMissingFields)
	}
	fromUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.UserId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	fromPrivateKey, err := b.chain.Decrypt(*fromUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	toUser, err := b.db.GetUser(&user.SearchUser{
		User: user.User{UserID: &req.AdminId},
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	toPrivateKey, err := b.chain.Decrypt(*toUser.EncryptedPrivateKey, b.key)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}

	tx, err := b.chain.TxDeleteAdmin(fromPrivateKey, &types.MsgDeleteAdmin{
		Creator: b.chain.GetAddress(fromPrivateKey),
		Address: b.chain.GetAddress(toPrivateKey),
	})
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return &api.DeleteAdminResponse{
		UserId:  *fromUser.UserID,
		AdminId: *toUser.UserID,
		Txh:     tx.TxResponse.TxHash,
	}, nil
}
