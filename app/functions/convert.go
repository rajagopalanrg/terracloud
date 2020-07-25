package functions

import (
	"terracloud/app/templates"
)

var disks *templates.DataDisks
var mvm *templates.MVMVARS

type vmdata map[string]map[string]string
type vmreturn interface{}

func CreateAzureVM(mvmvars *templates.MVMVARS) map[string]interface{} {
	vars := make(map[string]map[string]string)
	template := make(map[string]interface{})
	azurerm := make(map[string]map[string]interface{})
	module := make(map[string]map[string]interface{})
	//feature := make(map[string]string)
	if mvmvars.AdminPassword != "" {
		vars["adminPassword"] = make(map[string]string)
		vars["adminPassword"]["type"] = "string"
	}
	if mvmvars.AdminUsername != "" {
		vars["adminUsername"] = make(map[string]string)
		vars["adminUsername"]["type"] = "string"
	}
	if mvmvars.Location != "" {
		vars["location"] = make(map[string]string)
		vars["location"]["type"] = "string"
	}
	if mvmvars.VMName != "" {
		vars["vmname"] = make(map[string]string)
		vars["vmname"]["type"] = "string"
	}
	if mvmvars.ResourceGroupName != "" {
		vars["resourcegroupname"] = make(map[string]string)
		vars["resourcegroupname"]["type"] = "string"
	}
	if mvmvars.VMSku != "" {
		vars["vmsku"] = make(map[string]string)
		vars["vmsku"]["type"] = "string"
	}
	if mvmvars.VMSize != "" {
		vars["vmsize"] = make(map[string]string)
		vars["vmsize"]["type"] = "string"
	}
	if mvmvars.OSDataDiskSizeInGB != 0 {
		vars["osDiskSizeInGB"] = make(map[string]string)
		vars["osDiskSizeInGB"]["type"] = "number"
	}
	if mvmvars.StorageDataDisks != nil {
		vars["storagedatadisks"] = make(map[string]string)
		vars["storagedatadisks"]["type"] = "list(map(string))"
	}
	if mvmvars.VnetName != "" {
		vars["vnetName"] = make(map[string]string)
		vars["vnetName"]["type"] = "string"
	}
	vars["subnetName"] = make(map[string]string)
	if mvmvars.SubnetName != "" {
		vars["subnetName"]["type"] = "string"
	}
	if mvmvars.AvailabilitySet != "" {
		vars["availabilitySet"] = make(map[string]string)
		vars["availabilitySet"]["type"] = "string"
	}
	if mvmvars.IdentityIDs != nil {
		vars["identityIDs"] = make(map[string]string)
		vars["identityIDs"]["type"] = "string"
	}
	if mvmvars.BootDiagStorageAccount != "" {
		vars["bootDiagStorageAccount"] = make(map[string]string)
		vars["bootDiagStorageAccount"]["type"] = "string"
	}
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
	template["variable"] = vars

	//provider := make(map[string]map[string]map[string]interface{})
	azurerm["azurerm"] = make(map[string]interface{})
	azurerm["azurerm"]["version"] = "=2.4.0"
	azurerm["azurerm"]["subscription_id"] = "3dc3cd1a-d5cd-4e3e-a648-b2253048af83"
	azurerm["azurerm"]["client_id"] = "${var.client_id}"
	azurerm["azurerm"]["client_secret"] = "${var.client_secret}"
	azurerm["azurerm"]["tenant_id"] = "${var.tenant_id}"
	//feature["feature"] = ""
	azurerm["azurerm"]["feature"] = map[string]string{}
	//provider["provider"] = azurerm
	template["provider"] = azurerm
	module["azureVM"] = make(map[string]interface{})
	module["azureVM"]["source"] = "app.terraform.io/ClDevTeam/vm/azurerm"
	module["azureVM"]["version"] = "1.0.3"
	module["azureVM"]["adminPassword"] = "${var.adminPassword}"
	template["module"] = module
	return template
}
