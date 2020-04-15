module github.com/pulumi/pulumi-vault/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.0.0-20200414185723-c9aee63e6d4c
	github.com/pulumi/pulumi/sdk/v2 v2.0.0-beta.3
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20200403230059-d29ffdba020f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
