package cosmos

import (
	types3 "github.com/cosmos/cosmos-sdk/x/bank/types"
	types "github.com/medicalpoint/gateway/src/pb/types"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/status"
)

var (
	_ MedichainQuery = (*CosmosServiceClient)(nil)
)

type MedichainQuery interface {
	QueryAllAdmin(req *types.QueryAllAdminRequest) (*types.QueryAllAdminResponse, error)
	QueryAllSuperAdmin(req *types.QueryAllSuperAdminRequest) (*types.QueryAllSuperAdminResponse, error)
	QueryGetAdmin(req *types.QueryGetAdminRequest) (*types.QueryGetAdminResponse, error)
	QueryGetSuperAdmin(req *types.QueryGetSuperAdminRequest) (*types.QueryGetSuperAdminResponse, error)
	QueryGetAllBalance(req *types3.QueryAllBalancesRequest) (*types3.QueryAllBalancesResponse, error)
}

func (s *CosmosServiceClient) QueryGetAllBalance(req *types3.QueryAllBalancesRequest) (*types3.QueryAllBalancesResponse, error) {
	res, err := s.bankClient.AllBalances(s.ctx, req)

	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}
func (s *CosmosServiceClient) QueryAllAdmin(req *types.QueryAllAdminRequest) (*types.QueryAllAdminResponse, error) {
	res, err := s.queryClient.AdminAll(s.ctx, req)

	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) QueryAllSuperAdmin(req *types.QueryAllSuperAdminRequest) (*types.QueryAllSuperAdminResponse, error) {
	res, err := s.queryClient.SuperAdminAll(s.ctx, req)

	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) QueryGetAdmin(req *types.QueryGetAdminRequest) (*types.QueryGetAdminResponse, error) {
	res, err := s.queryClient.Admin(s.ctx, req)

	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}

func (s *CosmosServiceClient) QueryGetSuperAdmin(req *types.QueryGetSuperAdminRequest) (*types.QueryGetSuperAdminResponse, error) {
	res, err := s.queryClient.SuperAdmin(s.ctx, req)

	if err != nil {
		message := status.Convert(err).Message()
		return nil, xerrors.New(message)
	}
	return res, nil
}
