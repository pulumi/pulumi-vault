module github.com/pulumi/pulumi-vault

go 1.12

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v13.0.1+incompatible
	github.com/Nvveen/Gotty => github.com/ijc25/Gotty v0.0.0-20170406111628-a8b993ba6abd
	github.com/golang/glog => github.com/pulumi/glog v0.0.0-20180820174630-7eaa6ffb71e4
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)

require (
	cloud.google.com/go/logging v1.0.0 // indirect
	github.com/Azure/go-autorest/autorest v0.9.1 // indirect
	github.com/Azure/go-autorest/autorest/azure/auth v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/to v0.3.0 // indirect
	github.com/Azure/go-autorest/autorest/validation v0.2.0 // indirect
	github.com/hashicorp/terraform v0.12.9
	github.com/kr/pty v1.1.3 // indirect
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.2.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20190905205929-ed0b5c29edd1
	github.com/stretchr/testify v1.4.0 // indirect
	github.com/terraform-providers/terraform-provider-vault v0.0.0-20190906165453-bf7562d42d48
)
