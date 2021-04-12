module github.com/pulumi/pulumi-vault/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20210317202421-93cac347e234
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.23.0
	github.com/pulumi/pulumi/pkg/v2 v2.24.1-0.20210411193841-b7d403204449
	github.com/pulumi/pulumi/sdk/v2 v2.24.1
)

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
