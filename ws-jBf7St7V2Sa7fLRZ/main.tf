variable "client_id" {
	 type = string 
}
variable "client_secret" {
	 type = string 
}
variable "tenant_id" {
	 type = string 
}
variable "subscription_id" {
	 type = string 
}

provider "azurerm" {
	 version =  "=2.4.0"
	 client_id = var.client_id
	 client_secret = var.client_secret
	 subscription_id = var.subscription_id
	 tenant_id = var.tenant_id
	 features {}
}

module "vm" {
	version =  "1.0.4"
	source =  "app.terraform.io/ClDevTeam/vm/azurerm"
	location = "eastus"
	vm_name = "euwmvm01"
	resource_group_name = "Testing"
	admin_username = "winuser"
	admin_password = "Password@2020"
	vm_sku = "2016-Datacenter"
	vm_size = "Standard_DS1_v2"
	os_data_disk_size_in_gb = 0
	data_disks = [ 60,120 ]
	vnet_name = "vnet001"
	subnet_name = "subnet001"
}
