package controller

import (
	"github.com/Fedosin/manila-csi-operator/pkg/controller/manila"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, manila.Add)
}
