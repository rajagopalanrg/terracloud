package functions

import (
	"fmt"
	"log"
	"os"
	"reflect"
	"terracloud/app/templates"

	"github.com/iancoleman/strcase"
)

const tagName = "required"

func CreateAzureVM(mvmvars *templates.MVMVARS, terraformfile string) error {

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
	inputs := reflect.ValueOf(*mvmvars)
	typeofinput := inputs.Type()
	for i := 0; i < inputs.NumField(); i++ {
		moduleKey := strcase.ToSnake(typeofinput.Field(i).Name)
		moduleValue := inputs.Field(i).Interface()
		//log.Println(typeofinput.Field(i).Tag.Get(tagName))
		if moduleKey == "subscription_id" {
			continue
		}
		//log.Print(inputs.Field(i).Kind())
		/* if typeofinput.Field(i).Tag.Get(tagName) == "true" {
			if (inputs.Field(i).Kind() == reflect.Slice && inputs.Field(i).IsNil()) || (inputs.Field(i).Kind() == reflect.String && inputs.Field(i).IsZero()) || (inputs.Field(i).Kind() == reflect.Int && inputs.Field(i).IsZero()) {
				return fmt.Errorf("The required field %v is missing", moduleKey)
			}
		} else if typeofinput.Field(i).Tag.Get(tagName) == "false" {
			if (inputs.Field(i).Kind() == reflect.Slice && inputs.Field(i).IsNil()) || (inputs.Field(i).Kind() == reflect.String && inputs.Field(i).IsZero()) || (inputs.Field(i).Kind() == reflect.Int && inputs.Field(i).IsZero()) {
				log.Print(typeofinput.Field(i).Name)
				continue
			}
		} */
		if (inputs.Field(i).Kind() == reflect.Slice && inputs.Field(i).IsNil()) || (inputs.Field(i).Kind() == reflect.String && inputs.Field(i).IsZero()) || (inputs.Field(i).Kind() == reflect.Int && inputs.Field(i).IsZero()) {
			log.Print(typeofinput.Field(i).Name)
			continue
		}
		//module["azureVM"][moduleKey] = moduleValue
		if moduleKey == "os_data_disk_size_in_gb" {
			fmt.Fprintf(file, "\t%s = %v\n", moduleKey, moduleValue)
		} else if moduleKey == "data_disks" {
			len := inputs.Field(i).Len()
			value := inputs.Field(i).Interface().([]int)
			for j := 0; j < len; j++ {
				if j == 0 {
					fmt.Fprintf(file, "\t%s = [ %v,", moduleKey, value[j])
				} else if j+1 == len {
					fmt.Fprintf(file, "%v ]\n", value[j])
				} else {
					fmt.Fprintf(file, "%v ,", value[j])
				}
			}
		} else {
			fmt.Fprintf(file, "\t%s = \"%s\"\n", moduleKey, moduleValue)
		}
	}
	fmt.Fprintf(file, "}\n")

	return nil
}
