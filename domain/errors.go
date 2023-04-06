package domain

import (
	"fmt"
)

type Error error

var (
	ErrMessageNotFound Error = fmt.Errorf("not found message")
	ErrAuthorNotFound  Error = fmt.Errorf("not found author")
)