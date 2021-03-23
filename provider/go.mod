module github.com/pulumi/pulumi-vault/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20210317202421-93cac347e234
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.22.1
	github.com/pulumi/pulumi/sdk/v2 v2.22.1-0.20210310211618-1f16423ede4c
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
)
