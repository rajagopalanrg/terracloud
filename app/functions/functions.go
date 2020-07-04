package functions

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	tfe "github.com/hashicorp/go-tfe"
)

type SecureParams struct {
	UserToken string
}

func GetWorkspaceID(ctx context.Context, client *tfe.Client, workspaceName string, orgName string) (string, error) {
	listOptions := tfe.ListOptions{
		PageSize: 20,
	}
	workspaceOptions := tfe.WorkspaceListOptions{
		ListOptions: listOptions,
		Search:      &workspaceName,
	}
	workspaces, err := client.Workspaces.List(ctx, orgName, workspaceOptions)
	workspaceID := workspaces.Items[0].ID
	return workspaceID, err
	//fmt.Println("workspace ID: ", workspaceID)
}
func getRunID(ctx context.Context, client *tfe.Client, workspaceID string, configVersionID string) (string, string) {
	var runID, runStatus string
	//fmt.Println("enter get run id")
	listOptions := tfe.ListOptions{
		PageSize: 20,
	}
	runlistOptions := tfe.RunListOptions{
		ListOptions: listOptions,
	}

	for {
		getRuns, err := client.Runs.List(ctx, workspaceID, runlistOptions)
		requiredRuns := getRuns.Items
		//fmt.Println(requiredRuns)
		if err != nil {
			fmt.Println("error in getrun")
			fmt.Println(err)
		}
		for _, run := range requiredRuns {
			if run.ConfigurationVersion.ID == configVersionID && run.Status == "planned" {
				runID = run.ID
				runStatus = "planned"
				break
			} else if run.ConfigurationVersion.ID == configVersionID && run.Status == "errored" {
				runID = run.ID
				runStatus = "errored"
				break
			} else if run.ConfigurationVersion.ID == configVersionID && run.Status == "discarded" {
				runID = run.ID
				runStatus = "discarded"
				break
			}
		}
		if runID != "" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	return runID, runStatus
}
func getRun(ctx context.Context, client *tfe.Client, runID string) *tfe.Run {
	fmt.Println("we are in get run")
	runStruct, err := client.Runs.Read(ctx, runID)
	if err != nil {
		fmt.Println("error is here in get run")
		log.Fatal(err)
	}
	return runStruct
}
func GetRuns(ctx context.Context, client *tfe.Client, workspaceID string) (*tfe.RunList, error) {
	listOptions := tfe.ListOptions{
		PageSize: 20,
	}
	runlistOptions := tfe.RunListOptions{
		ListOptions: listOptions,
	}
	getRuns, err := client.Runs.List(ctx, workspaceID, runlistOptions)

	return getRuns, err
}
func getApplyLog(ctx context.Context, client *tfe.Client, applyID string) {

	getApplyLogs, err := client.Applies.Logs(ctx, applyID)
	applyLog, err := ioutil.ReadAll(getApplyLogs)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Please find the log after applying")
	fmt.Println()
	fmt.Printf("%s", applyLog)
	fmt.Println()
}
func printPlan(ctx context.Context, client *tfe.Client, planID string) {
	getPlanLog, err := client.Plans.Logs(ctx, planID)
	planLog, err := ioutil.ReadAll(getPlanLog)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	fmt.Printf("%s", planLog)
	fmt.Println()
}
func apply(ctx context.Context, client *tfe.Client, applyComment *string, runID string) {
	fmt.Println("enter apply block")
	runApplyOptions := tfe.RunApplyOptions{
		Comment: applyComment,
	}
	err := client.Runs.Apply(ctx, runID, runApplyOptions)
	if err != nil {
		fmt.Println("error is here in apply")
		log.Fatal(err)
	} else {
		fmt.Println("Your template is applied successfully")
	}
}
func createConfigVersion(ctx context.Context, client *tfe.Client, autoqueueruns *bool, workspaceID string, uploadPATH string) string {
	configurationVersionOptions := tfe.ConfigurationVersionCreateOptions{
		AutoQueueRuns: autoqueueruns,
	}
	newConfigurationVersion, err := client.ConfigurationVersions.Create(ctx, workspaceID, configurationVersionOptions)
	if err != nil {
		log.Fatal(err)
	}
	UploadURL := newConfigurationVersion.UploadURL
	configVersionID := newConfigurationVersion.ID
	err = client.ConfigurationVersions.Upload(ctx, UploadURL, uploadPATH)
	if err != nil {
		fmt.Println("the upload block returns error")
		log.Fatal(err)
	}
	//fmt.Println("your run is created")
	return configVersionID
}
func CreateVariables(ctx context.Context, client *tfe.Client, workspaceID string, variableCreateOptions *tfe.VariableCreateOptions) (*tfe.Variable, error) {
	/* key := "clientID"
	value := "application_id"
	variableCreateOptions := tfe.VariableCreateOptions{
		Key:   &key,
		Value: &value,
	} */
	variable, err := client.Variables.Create(ctx, workspaceID, *variableCreateOptions)

	return variable, err
}
