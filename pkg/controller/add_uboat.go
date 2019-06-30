package controller

import (
	"github.com/onlyascii/u-boat-operator/pkg/controller/uboat"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, uboat.Add)
}
