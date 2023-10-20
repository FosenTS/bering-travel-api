package storage

import (
	"fmt"
)

var ErrIDIsUndefined = fmt.Errorf("error id is undefined")

var ErrNoEffect = fmt.Errorf("error no effect")

var ErrNoRows = fmt.Errorf("error no rows")
