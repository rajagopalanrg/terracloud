package controllers

import (
	"terracloud/app/functions"
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
	vars := make(map[string]interface{})
	vars = functions.CreateAzureVM(mvm)
	return c.RenderJSON(vars)
}
