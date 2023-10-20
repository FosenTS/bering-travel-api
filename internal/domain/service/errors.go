package service

import (
	"errors"
	"fmt"
)

var ErrNotFound = fmt.Errorf("error not found")

var ErrAlreadyExists = fmt.Errorf("error already exists")

var ErrUnapprovedUser = errors.New("error unapproved user")
