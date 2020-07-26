package functions

import (
	"log"
	"reflect"
	"terracloud/app/templates"

	"github.com/iancoleman/strcase"
)

var disks *templates.DataDisks

type vmdata map[string]map[string]string
type vmreturn interface{}

func CreateAzureVM(mvmvars *templates.MVMVARS) map[string]interface{} {
	vars := make(map[string]map[string]string)
	template := make(map[string]interface{})
	azurerm := make(map[string]map[string]interface{})
	module := make(map[string]map[string]interface{})
	//feature := make(map[string]string)
	if mvmvars.Tenant_ID != "" {
		vars["tenant_id"] = make(map[string]string)
		vars["tenant_id"]["type"] = "string"
	}
	if mvmvars.Client_ID != "" {
		vars["client_id"] = make(map[string]string)
		vars["client_id"]["type"] = "string"
	}
	if mvmvars.Client_Secret != "" {
		vars["client_secret"] = make(map[string]string)
		vars["client_secret"]["type"] = "string"
	}
	if mvmvars.Subscription_ID != "" {
		vars["subscription_id"] = make(map[string]string)
		vars["subscription_id"]["type"] = "string"
	}
	template["variable"] = vars
	azurerm["azurerm"] = make(map[string]interface{})
	azurerm["azurerm"]["version"] = "=2.4.0"
	azurerm["azurerm"]["subscription_id"] = "${var.subscription_id}"
	azurerm["azurerm"]["client_id"] = "${var.client_id}"
	azurerm["azurerm"]["client_secret"] = "${var.client_secret}"
	azurerm["azurerm"]["tenant_id"] = "${var.tenant_id}"
	azurerm["azurerm"]["features"] = map[string]string{}
	template["provider"] = azurerm
	module["azureVM"] = make(map[string]interface{})
	module["azureVM"]["source"] = "app.terraform.io/ClDevTeam/vm/azurerm"
	module["azureVM"]["version"] = "1.0.4"
	inputs := reflect.ValueOf(*mvmvars)
	typeofinput := inputs.Type()
	for i := 0; i < inputs.NumField(); i++ {
		if typeofinput.Field(i).Name == "Client_ID" || typeofinput.Field(i).Name == "Client_Secret" || typeofinput.Field(i).Name == "Tenant_ID" || typeofinput.Field(i).Name == "Subscription_ID" {
			continue
		}

		moduleKey := strcase.ToSnake(typeofinput.Field(i).Name)
		moduleValue := inputs.Field(i).Interface()
		/* if moduleValue == "" {
			//log.Print(typeofinput.Field(i).Name)
			continue
		} */
		//log.Print(inputs.Field(i).Kind())
		if (inputs.Field(i).Kind() == reflect.Slice && inputs.Field(i).IsNil()) || (inputs.Field(i).Kind() == reflect.String && inputs.Field(i).IsZero()) {
			log.Print(typeofinput.Field(i).Name)
			continue
		}
		//fmt.Printf("Key: %s\tValue: %s", inputs.Field(i), inputs.Field(i).Kind())
		module["azureVM"][moduleKey] = moduleValue
	}
	template["module"] = module
	return template
}
