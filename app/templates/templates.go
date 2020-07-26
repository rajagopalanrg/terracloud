package templates

type DataDisks struct {
	CreateOption     string
	Lun              int
	Name             string
	DataDiskSizeInGB int
}

type MVMVARS struct {
	Location               string
	VMName                 string
	ResourceGroupName      string
	AdminUsername          string
	AdminPassword          string
	VMSku                  string
	VMSize                 string
	OSDataDiskSizeInGB     int
	StorageDataDisks       []DataDisks
	BootDiagStorageAccount string
	VnetName               string
	SubnetName             string
	AvailabilitySet        string
	IdentityID             []string
	Client_ID              string
	Client_Secret          string
	Tenant_ID              string
	Subscription_ID        string
}
