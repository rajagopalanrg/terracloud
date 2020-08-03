package templates

type Common struct {
	Resources interface{}
}

type MVMVARS struct {
	Location               string   `json:"location" validate:"required"`
	VMName                 string   `json:"vm_name" validate:"required"`
	ResourceGroupName      string   `json:"rg_name" validate:"required"`
	AdminUsername          string   `json:"admin_username" validate:"required"`
	AdminPassword          string   `json:"admin_password" validate:"required"`
	VMSku                  string   `json:"vm_sku" validate:"required"`
	VMSize                 string   `json:"vm_size" validate:"required"`
	OSDataDiskSizeInGB     int      `json:"osdatadisksizeingb" validate:"required"`
	DataDisks              []int    `json:"data_disks" validate:"required"`
	BootDiagStorageAccount string   `json:"boot_diag_storage"`
	VnetName               string   `json:"vnet_name" validate:"required"`
	SubnetName             string   `json:"subnet_name" validate:"required"`
	AvailabilitySet        string   `json:"availability_set"`
	IdentityID             []string `json:"identity_id"`
	SubscriptionID         string   `json:"subscription_Id" validate:"required"`
}
