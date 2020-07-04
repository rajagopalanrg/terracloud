package controllers

import (
	"context"
	"log"
	"terracloud/app/functions"

	"github.com/hashicorp/go-tfe"
	"github.com/revel/revel"
)

type Deployment struct {
	*revel.Controller
}

var config *tfe.Config
var client *tfe.Client

/* func (c Deployment) GetHeader() revel.Result {
	userName := c.Request.Header.Get("userToken")
	return c.RenderText(userName)
} */
func (c Deployment) FileUpload() revel.Result {
	fileName := c.Params.Files["file"][0].Filename
	return c.RenderText(fileName)
}
func (c Deployment) CreateVariable(org string, workspaceName string) revel.Result {
	userToken := c.Request.Header.Get("userToken")
	config := &tfe.Config{
		Token: userToken,
	}
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	var variableOptions *tfe.VariableCreateOptions
	c.Params.BindJSON(&variableOptions)
	workspaceID, err := functions.GetWorkspaceID(ctx, client, workspaceName, org)
	variable, err := functions.CreateVariables(ctx, client, workspaceID, variableOptions)
	if err != nil {
		return c.RenderText(err.Error())
	}
	return c.RenderJSON(variable)
}
func (c Deployment) GetWorkspace(org string, workspaceName string) revel.Result {
	//var secureParams *functions.SecureParams
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
	if err != nil {
		return c.RenderText(err.Error())
	}
	return c.RenderText(workspaceID)
}
func (c Deployment) GetRuns(workspaceID string) revel.Result {
	//var secureParams *functions.SecureParams
	userToken := c.Request.Header.Get("userToken")
	config := &tfe.Config{
		Token: userToken,
	}
	client, err := tfe.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	ctx := context.Background()
	runs, err := functions.GetRuns(ctx, client, workspaceID)
	if err != nil {
		return c.RenderText(err.Error())
	}
	return c.RenderJSON(runs)
}
