package user

import "golang.org/x/xerrors"

var (
	ErrCreateFail = xerrors.New("create user fail")
	ErrNotFound   = xerrors.New("user not found")
)
