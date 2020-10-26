module github.com/pulumi/pulumi-vault/provider/v3

go 1.14

require (
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.11.1-0.20201020163502-64cff1e50894
	github.com/pulumi/pulumi/sdk/v2 v2.12.0
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20201021161834-d3423e3099a3
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
