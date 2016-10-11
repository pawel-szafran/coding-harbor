package buildorder

import "errors"

var ErrDepCycle = errors.New("Dependency Cycle Error")
