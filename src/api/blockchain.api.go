package api

import (
	"github.com/gin-gonic/gin"
	"github.com/medicalpoint/gateway/src/lib"
	"github.com/medicalpoint/gateway/src/pb/api"
	"github.com/sirupsen/logrus"
)

type BlockchainApi struct {
	service *BlockchainService
	logger  *logrus.Logger
}

func (b *BlockchainApi) HandleUserPost(g *gin.Context) {
	req := &api.PostUserRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.CreateUser(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleBalanceGet(g *gin.Context) {
	req := &api.GetBalanceRequest{
		Id: g.Query("id"),
	}
	res, err := b.service.GetBalance(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleMintPost(g *gin.Context) {
	req := &api.PostMintRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.Mint(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleBurnPost(g *gin.Context) {
	req := &api.PostBurnRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.Burn(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleSendSystemPost(g *gin.Context) {
	req := &api.PostSendSystemRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.SendSystem(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleTransferPost(g *gin.Context) {
	req := &api.PostTransferRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.Transfer(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
func (b *BlockchainApi) HandleSuperAdminPut(g *gin.Context) {
	req := &api.PutSuperAdminRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.UpdateSuperAdmin(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
func (b *BlockchainApi) HandleTotalSupplyGet(g *gin.Context) {
	req := &api.GetTotalSupplyRequest{}
	res, err := b.service.GetTotalSupply(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleAdminTransferPost(g *gin.Context) {
	req := &api.PostAdminTransferRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.AdminTransfer(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleAdminPost(g *gin.Context) {
	req := &api.PostAdminRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.AddAdmin(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleAdminDelete(g *gin.Context) {
	req := &api.DeleteAdminRequest{}
	err := g.BindJSON(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	res, err := b.service.DeleteAdmin(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}

func (b *BlockchainApi) HandleSystemBalanceGet(g *gin.Context) {
	req := &api.GetSystemBalanceRequest{}

	res, err := b.service.GetSystemBalance(req)
	if err != nil {
		lib.ErrBadRequest(g, err)
		return
	}
	lib.Success(g, res)
}
