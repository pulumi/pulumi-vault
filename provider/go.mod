module github.com/pulumi/pulumi-vault/provider/v3

go 1.14

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.13.2
	github.com/pulumi/pulumi/sdk/v2 v2.13.3-0.20201109230029-a6f8b9b205cd
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20201119160613-78798a9eb16c
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
