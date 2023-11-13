package types

import "errors"

var (
	InsertFailed  = errors.New("insert failed")
	AlreadyExists = errors.New("already exists")
)
