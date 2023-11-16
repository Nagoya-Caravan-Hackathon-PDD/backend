package types

import "errors"

var (
	InsertFailed  = errors.New("insert failed")
	AlreadyExists = errors.New("already exists")
	NotFound      = errors.New("not found")
)

var (
	ErrBadResponse = errors.New("bad response")
)
