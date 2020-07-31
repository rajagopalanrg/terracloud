package functions

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"terracloud/app/templates"

	"github.com/iancoleman/strcase"
)

type vmdata map[string]map[string]string
type vmreturn interface{}

func CreateAzureVM(mvmvars *templates.MVMVARS, terraformfile string) error {
	/* vars := make(map[string]map[string]string)
	template := make(map[string]interface{})
	azurerm := make(map[string]map[string]interface{})
	module := make(map[string]map[string]interface{}) */

	file, err := os.Create(terraformfile)
	if err != nil {
		return err
	}

	fmt.Fprintf(file, "variable \"client_id\" {\n\t type = string \n}")
	fmt.Fprintf(file, "\nvariable \"client_secret\" {\n\t type = string \n}")
	fmt.Fprintf(file, "\nvariable \"tenant_id\" {\n\t type = string \n}")
	fmt.Fprintf(file, "\nvariable \"subscription_id\" {\n\t type = string \n}\n")
	fmt.Fprintf(file, "\nprovider \"azurerm\" {\n")
	fmt.Fprintf(file, "\t version =  \"=2.4.0\"\n")
	fmt.Fprintf(file, "\t client_id = var.client_id\n")
	fmt.Fprintf(file, "\t client_secret = var.client_secret\n")
	fmt.Fprintf(file, "\t subscription_id = var.subscription_id\n")
	fmt.Fprintf(file, "\t tenant_id = var.tenant_id\n")
	fmt.Fprintf(file, "\t features {}\n")
	fmt.Fprintf(file, "}\n")
	fmt.Fprintf(file, "\nmodule \"vm\" {\n")
	fmt.Fprintf(file, "\tversion =  \"1.0.4\"\n")
	fmt.Fprintf(file, "\tsource =  \"app.terraform.io/ClDevTeam/vm/azurerm\"\n")
	defer file.Close()
	//feature := make(map[string]string)

	/* vars["tenant_id"] = make(map[string]string)
	vars["tenant_id"]["type"] = "string"
	vars["client_id"] = make(map[string]string)
	vars["client_id"]["type"] = "string"
	vars["client_secret"] = make(map[string]string)
	vars["client_secret"]["type"] = "string"
	template["variable"] = vars

	azurerm["azurerm"] = make(map[string]interface{})
	azurerm["azurerm"]["version"] = "=2.4.0"
	azurerm["azurerm"]["subscription_id"] = mvmvars.Subscription_ID
	azurerm["azurerm"]["client_id"] = "${var.client_id}"
	azurerm["azurerm"]["client_secret"] = "${var.client_secret}"
	azurerm["azurerm"]["tenant_id"] = "${var.tenant_id}"
	azurerm["azurerm"]["features"] = map[string]string{}
	template["provider"] = azurerm

	module["azureVM"] = make(map[string]interface{})
	module["azureVM"]["source"] = "app.terraform.io/ClDevTeam/vm/azurerm"
	module["azureVM"]["version"] = "1.0.4" */
	inputs := reflect.ValueOf(*mvmvars)
	typeofinput := inputs.Type()
	for i := 0; i < inputs.NumField(); i++ {
		if typeofinput.Field(i).Name == "Client_ID" || typeofinput.Field(i).Name == "Client_Secret" || typeofinput.Field(i).Name == "Tenant_ID" || typeofinput.Field(i).Name == "Subscription_ID" {
			continue
		}
		moduleKey := strcase.ToSnake(typeofinput.Field(i).Name)
		moduleValue := inputs.Field(i).Interface()
		//log.Print(inputs.Field(i).Kind())
		if (inputs.Field(i).Kind() == reflect.Slice && inputs.Field(i).IsNil()) || (inputs.Field(i).Kind() == reflect.String && inputs.Field(i).IsZero()) {
			log.Print(typeofinput.Field(i).Name)
			continue
		}
		//fmt.Printf("Key: %s\tValue: %s", inputs.Field(i), inputs.Field(i).Kind())
		//module["azureVM"][moduleKey] = moduleValue
		if moduleKey == "os_data_disk_size_in_gb" {
			fmt.Fprintf(file, "\t%s = %b\n", moduleKey, moduleValue)
		} else if moduleKey == "data_disks" {
			fmt.Fprintf(file, "\t%s = %v\n", moduleKey, moduleValue)
		} else {
			fmt.Fprintf(file, "\t%s = \"%s\"\n", moduleKey, moduleValue)
		}
	}
	fmt.Fprintf(file, "}\n")
	//template["module"] = module

	return nil
}
