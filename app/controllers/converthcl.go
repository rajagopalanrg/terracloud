package controllers

import (
	"context"
	"log"
	"os"
	"terracloud/app/functions"
	"terracloud/app/templates"

	"github.com/hashicorp/go-tfe"
	"github.com/revel/revel"
	"gopkg.in/go-playground/validator.v9"
)

type Convert struct {
	*revel.Controller
}

var mvm *templates.MVMVARS

func (c Convert) AzureWindowsVM(workspaceName string, org string) revel.Result {
	userToken := c.Request.Header.Get("userToken")
	config := &tfe.Config{
		Token: userToken,
	}
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	workspaceID, err := functions.GetWorkspaceID(ctx, client, workspaceName, org)
	path, err := os.Getwd()
	terraformfile := path + "\\" + workspaceID + "\\main.tf"

	c.Params.BindJSON(&mvm)
	v := validator.New()
	err = v.Struct(mvm)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			log.Println(e)
			return c.RenderText(err.Error())
		}
	}
	//vars := make(map[string]interface{})

	err = functions.CreateAzureVM(mvm, terraformfile)
	if err != nil {
		return c.RenderText(err.Error())
	}
	gzipfile := functions.Gzip(terraformfile)
	//err = functions.WriteFileToDisk(filename, vars, filepath)
	return c.RenderText(gzipfile)
}
