module github.com/pulumi/pulumi-vault/provider/v4

go 1.16

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20210915180734-eca3fa6fa12a
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/pkg/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)

replace github.com/hashicorp/go-getter v1.5.0 => github.com/hashicorp/go-getter v1.4.0
