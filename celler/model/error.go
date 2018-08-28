package model

import (
	"fmt"
)

var (
	//ErrNoRow = errors.New("no rows in result set")
	ErrNoRow = fmt.Errorf("no rows in result set")
)
