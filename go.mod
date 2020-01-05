module github.com/pulumi/pulumi-vault

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.5.2
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20191206220152-8bebdb873141
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
