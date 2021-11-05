package cosmos

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	types2 "github.com/cosmos/cosmos-sdk/crypto/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	types3 "github.com/cosmos/cosmos-sdk/x/bank/types"
	types "github.com/medicalpoint/gateway/src/pb/types"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/status"
)

var (
	_ MedichainTx = (*CosmosServiceClient)(nil)
)

type MedichainTx interface {
	TxAdminTransfer(priv []byte, req *types.MsgAdminTransfer) (*txtypes.BroadcastTxResponse, error)
	TxBurn(priv []byte, req *types.MsgBurn) (*txtypes.BroadcastTxResponse, error)
	TxCreateAdmin(priv []byte, req *types.MsgCreateAdmin) (*txtypes.BroadcastTxResponse, error)
	TxCreateSuperAdmin(priv []byte, req *types.MsgCreateSuperAdmin) (*txtypes.BroadcastTxResponse, error)
	TxDeleteAdmin(priv []byte, req *types.MsgDeleteAdmin) (*txtypes.BroadcastTxResponse, error)
	TxMint(priv []byte, req *types.MsgMint) (*txtypes.BroadcastTxResponse, error)
	TxUpdateSuperAdmin(priv []byte, req *types.MsgUpdateSuperAdmin) (*txtypes.BroadcastTxResponse, error)
	TxTransfer(priv []byte, req *types3.MsgSend) (*txtypes.BroadcastTxResponse, error)
}

func (s *CosmosServiceClient) TxTransfer(priv []byte, req *types3.MsgSend) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxBurn(priv []byte, req *types.MsgBurn) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxCreateAdmin(priv []byte, req *types.MsgCreateAdmin) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxCreateSuperAdmin(priv []byte, req *types.MsgCreateSuperAdmin) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxDeleteAdmin(priv []byte, req *types.MsgDeleteAdmin) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxMint(priv []byte, req *types.MsgMint) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxUpdateSuperAdmin(priv []byte, req *types.MsgUpdateSuperAdmin) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) TxAdminTransfer(priv []byte, req *types.MsgAdminTransfer) (*txtypes.BroadcastTxResponse, error) {
	privs := []types2.PrivKey{&secp256k1.PrivKey{Key: priv}}

	res, err := s.sendTx(req, privs)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	
	return res, nil
}
