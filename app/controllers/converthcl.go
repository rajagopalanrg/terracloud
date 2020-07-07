package controllers

import (
	"terracloud/app/templates"

	"github.com/revel/revel"
)

type Convert struct {
	*revel.Controller
}

var disks *templates.DataDisks
var mvm *templates.MVMVARS

func (c Convert) AzureWindowsVM() revel.Result {
	c.Params.BindJSON(&mvm)
	return c.RenderJSON(mvm)
}
