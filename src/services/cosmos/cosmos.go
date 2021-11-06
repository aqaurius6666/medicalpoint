package cosmos

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"strconv"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	types4 "github.com/cosmos/cosmos-sdk/codec/types"
	codec2 "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	types2 "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	txtypes "github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
	signing2 "github.com/cosmos/cosmos-sdk/x/auth/signing"
	tx2 "github.com/cosmos/cosmos-sdk/x/auth/tx"
	types3 "github.com/cosmos/cosmos-sdk/x/auth/types"
	types5 "github.com/cosmos/cosmos-sdk/x/bank/types"
	types "github.com/medicalpoint/gateway/src/pb/types"
	"github.com/sirupsen/logrus"
	"github.com/sonntuet1997/medical-chain-utils/cryptography"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

var (
	GAS_LIMIT uint64 = 400000
)

type CosmosEndpoint string
type ChainID string
type Mnemonic string

type CosmosServiceClient struct {
	queryClient types.QueryClient
	authClient  types3.QueryClient
	bankClient  types5.QueryClient
	txClient    txtypes.ServiceClient
	ctx         context.Context
	chainId     string
	keyring     keyring.UnsafeKeyring
	cdc         *codec.ProtoCodec
	logger      *logrus.Logger
}

func InitCodec() (cdc *codec.ProtoCodec) {
	registry := types4.NewInterfaceRegistry()
	types3.RegisterInterfaces(registry)
	codec2.RegisterInterfaces(registry)
	types.RegisterInterfaces(registry)
	cdc = codec.NewProtoCodec(registry)
	return cdc
}

type CosmosOpts struct {
	ChainId  ChainID
	KeyRing  keyring.UnsafeKeyring
	CosmosEp CosmosEndpoint
	Mne      Mnemonic
}

func NewCosmosServiceClient(ctx context.Context, logger *logrus.Logger, opts CosmosOpts) *CosmosServiceClient {
	cc, err := grpc.DialContext(ctx, string(opts.CosmosEp), grpc.WithInsecure())
	if err != nil {
		return nil
	}
	authClient := types3.NewQueryClient(cc)
	qClient := types.NewQueryClient(cc)
	txClient := txtypes.NewServiceClient(cc)
	bankClient := types5.NewQueryClient(cc)

	ins := &CosmosServiceClient{
		queryClient: qClient,
		txClient:    txClient,
		authClient:  authClient,
		bankClient:  bankClient,
		ctx:         ctx,
		chainId:     string(opts.ChainId),
		keyring:     opts.KeyRing,
		cdc:         InitCodec(),
		logger:      logger,
	}
	_, err = ins.AddAccountFromMnemonic("admin", string(opts.Mne))
	if err != nil {
		return nil
	}
	return ins
}

func (s *CosmosServiceClient) GetAdminAddress() string {
	priv, err := s.getAdminKey()
	if err != nil {
		s.logger.Error(err)
		return ""
	}
	add, err := sdk.Bech32ifyAddressBytes("medipoint", priv.PubKey().Address().Bytes())
	if err != nil {
		s.logger.Error(err)
		return ""
	}
	return add
}
func (s *CosmosServiceClient) GetAddress(privKey []byte) string {
	priv := secp256k1.PrivKey{
		Key: privKey,
	}
	add, err := sdk.Bech32ifyAddressBytes("medipoint", priv.PubKey().Address().Bytes())
	if err != nil {
		s.logger.Error(err)
		return ""
	}
	return add
}

func (s *CosmosServiceClient) GenPrivateKey() *secp256k1.PrivKey {
	return secp256k1.GenPrivKey()
}

func (s *CosmosServiceClient) Encrypt(priv *secp256k1.PrivKey, passphrase string) (string, error) {
	// privByte, _ := hex.DecodeString("A7FF23E8D73FC8D6D1F462AC93AD92EF9A383A8771A6E9B9651D4294CA39BD6C")
	// cipher, err := cryptography.EncryptMessage(privByte, []byte(passphrase))
	cipher, err := cryptography.EncryptMessage(priv.Bytes(), []byte(passphrase))
	if err != nil {
		return "", xerrors.Errorf("%w", err)
	}
	return cryptography.ConvertBytesToBase64(cipher), nil
}

func (s *CosmosServiceClient) Decrypt(encrypted, passphrase string) ([]byte, error) {
	by, err := cryptography.ConvertBase64ToBytes(encrypted)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	data, err := cryptography.DecryptCipher(by, []byte(passphrase))
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return data, nil

}

func (s *CosmosServiceClient) GetAccount(address string) (*types3.BaseAccount, error) {
	req := types3.QueryAccountRequest{Address: address}
	res, err := s.authClient.Account(s.ctx, &req)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	var account = types3.BaseAccount{}
	bz, err := s.cdc.MarshalJSON(res.Account)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	bz, err = s.ConvertAccountByte(bz)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	err = json.Unmarshal(bz, &account)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	return &account, nil
}
func (s *CosmosServiceClient) ConvertAccountByte(bz []byte) ([]byte, error) {
	var temp map[string]interface{}
	var err error
	json.Unmarshal(bz, &temp)
	temp["account_number"], err = strconv.ParseInt(temp["account_number"].(string), 10, 64)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	temp["sequence"], err = strconv.ParseInt(temp["sequence"].(string), 10, 64)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	nbz, err := json.Marshal(temp)
	if err != nil {
		return nil, xerrors.Errorf("%w", err)
	}
	return nbz, nil
}
func (s *CosmosServiceClient) AddAccountFromMnemonic(uid string, mnemonic string) (keyring.Info, error) {
	return s.keyring.NewAccount(uid, mnemonic, "", "m/44'/118'/0'/0", hd.Secp256k1)
}

func (s *CosmosServiceClient) ShowAccount(uid string) (keyring.Info, error) {
	return s.keyring.Key(uid)
}

func (s *CosmosServiceClient) getAdminKey() (*secp256k1.PrivKey, error) {
	adminPrivStr, err := s.keyring.UnsafeExportPrivKeyHex("admin")
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	adminPrivBz, err := hex.DecodeString(adminPrivStr)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	adminPriv := secp256k1.PrivKey{Key: adminPrivBz}

	return &adminPriv, nil
}

func (s *CosmosServiceClient) sendTxs(msgs []sdk.Msg, privs []types2.PrivKey) (*txtypes.BroadcastTxResponse, error) {
	var accs []*types3.BaseAccount

	for _, priv := range privs {
		addr, err := sdk.Bech32ifyAddressBytes("medipoint", priv.PubKey().Address().Bytes())

		if err != nil {
			message := status.Convert(err).Message()
			return nil, xerrors.New(message)
		}
		acc, err := s.GetAccount(addr)
		if err != nil {
			message := status.Convert(err).Message()
			return nil, xerrors.New(message)
		}
		accs = append(accs, acc)
	}

	txConfig := tx2.NewTxConfig(s.cdc, tx2.DefaultSignModes)

	txBuilder := txConfig.NewTxBuilder()
	err := txBuilder.SetMsgs(msgs...)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txBuilder.SetGasLimit(GAS_LIMIT)

	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  txConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accs[i].Sequence,
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := signing2.SignerData{
			ChainID:       s.chainId,
			AccountNumber: accs[i].AccountNumber,
			Sequence:      accs[i].Sequence,
		}
		sigV2, err := tx.SignWithPrivKey(
			txConfig.SignModeHandler().DefaultMode(), signerData,
			txBuilder, priv, txConfig, accs[i].Sequence)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txRes, err := s.txClient.BroadcastTx(
		s.ctx,
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	if txRes.TxResponse.Logs == nil && txRes.TxResponse.Code != 0 {
		message := txRes.TxResponse.RawLog
		return nil, xerrors.New(message)
	}

	return txRes, nil
}

func (s *CosmosServiceClient) sendTx(msg sdk.Msg, privs []types2.PrivKey) (*txtypes.BroadcastTxResponse, error) {
	var accs []*types3.BaseAccount

	for _, priv := range privs {
		addr, err := sdk.Bech32ifyAddressBytes("medipoint", priv.PubKey().Address().Bytes())

		if err != nil {
			message := status.Convert(err).Message()
			return nil, xerrors.New(message)
		}
		acc, err := s.GetAccount(addr)
		if err != nil {
			message := status.Convert(err).Message()
			return nil, xerrors.New(message)
		}
		accs = append(accs, acc)
	}

	txConfig := tx2.NewTxConfig(s.cdc, tx2.DefaultSignModes)

	txBuilder := txConfig.NewTxBuilder()

	err := txBuilder.SetMsgs(msg)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txBuilder.SetGasLimit(GAS_LIMIT)

	var sigsV2 []signing.SignatureV2
	for i, priv := range privs {
		sigV2 := signing.SignatureV2{
			PubKey: priv.PubKey(),
			Data: &signing.SingleSignatureData{
				SignMode:  txConfig.SignModeHandler().DefaultMode(),
				Signature: nil,
			},
			Sequence: accs[i].Sequence,
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	sigsV2 = []signing.SignatureV2{}
	for i, priv := range privs {
		signerData := signing2.SignerData{
			ChainID:       s.chainId,
			AccountNumber: accs[i].AccountNumber,
			Sequence:      accs[i].Sequence,
		}
		sigV2, err := tx.SignWithPrivKey(
			txConfig.SignModeHandler().DefaultMode(), signerData,
			txBuilder, priv, txConfig, accs[i].Sequence)
		if err != nil {
			return nil, xerrors.Errorf("%w", err)
		}

		sigsV2 = append(sigsV2, sigV2)
	}

	err = txBuilder.SetSignatures(sigsV2...)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	txRes, err := s.txClient.BroadcastTx(
		s.ctx,
		&txtypes.BroadcastTxRequest{
			Mode:    txtypes.BroadcastMode_BROADCAST_MODE_BLOCK,
			TxBytes: txBytes, // Proto-binary of the signed transaction, see previous step.
		},
	)
	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}

	if txRes.TxResponse.Logs == nil && txRes.TxResponse.Code != 0 {
		message := txRes.TxResponse.RawLog
		return nil, xerrors.New(message)
	}

	return txRes, nil
}

// ParseTxResponse only valid if transaction response has only one response that is target data type
func (s *CosmosServiceClient) ParseTxResponse(tx *txtypes.BroadcastTxResponse, target codec.ProtoMarshaler) error {
	bz, err := hex.DecodeString(tx.GetTxResponse().Data)
	if err != nil {
		return err
	}
	var temp sdk.TxMsgData
	err = s.cdc.Unmarshal(bz, &temp)
	if err != nil {
		return err
	}
	return s.cdc.Unmarshal(temp.Data[0].Data, target)
}

// ParseTxResponse  valid if transaction response has more than one response
func (s *CosmosServiceClient) ParseTxResponses(tx *txtypes.BroadcastTxResponse, target []codec.ProtoMarshaler) error {
	bz, err := hex.DecodeString(tx.GetTxResponse().Data)
	if err != nil {
		return err
	}
	var temp sdk.TxMsgData
	err = s.cdc.Unmarshal(bz, &temp)
	if err != nil {
		return err
	}
	target = make([]codec.ProtoMarshaler, len(temp.Data))
	for i, body := range temp.Data {
		err := s.cdc.Unmarshal(body.Data, target[i])
		if err != nil {
			return err
		}
	}
	return nil
}
