module github.com/pulumi/pulumi-vault/provider

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.4.1
	github.com/pulumi/pulumi v1.12.2-0.20200313044354-8111d33438b9
	github.com/pulumi/pulumi-terraform-bridge v1.8.2
	github.com/terraform-providers/terraform-provider-vault v1.9.1-0.20200313194542-b2d81593a48e
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
