module github.com/pulumi/pulumi-vault

go 1.12

require (
	github.com/hashicorp/terraform-plugin-sdk v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform-bridge v1.4.2
	github.com/terraform-providers/terraform-provider-vault v0.0.0-20191017211552-55806c9f74a4
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
