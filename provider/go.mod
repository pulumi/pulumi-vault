module github.com/pulumi/pulumi-vault/provider/v3

go 1.14

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20201215172718-c339b9ad9be6
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.17.0
	github.com/pulumi/pulumi/sdk/v2 v2.17.0
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
