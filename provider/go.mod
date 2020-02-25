module github.com/pulumi/pulumi-vault/provider

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pulumi/pulumi v1.9.1
	github.com/pulumi/pulumi-terraform-bridge v1.7.0
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20200205231359-74bbb849786b
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
