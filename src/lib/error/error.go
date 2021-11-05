package e

import "golang.org/x/xerrors"

var (
	ErrMissingFields = xerrors.New("missing required fields")
	ErrQueryInvalid  = xerrors.New("query parameters invalid")
	ErrAmountInvalid = xerrors.New("amount invalid")
	ErrDenomInvalid  = xerrors.New("denom invalid")
)
