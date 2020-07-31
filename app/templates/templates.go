package templates

type MVMVARS struct {
	Location               string
	VMName                 string
	ResourceGroupName      string
	AdminUsername          string
	AdminPassword          string
	VMSku                  string
	VMSize                 string
	OSDataDiskSizeInGB     int
	DataDisks              []int
	BootDiagStorageAccount string
	VnetName               string
	SubnetName             string
	AvailabilitySet        string
	IdentityID             []string
	Subscription_ID        string
}
