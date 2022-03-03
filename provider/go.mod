module github.com/pulumi/pulumi-vault/provider/v5

go 1.16

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20220225202911-91f40a2aed19
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.18.0
	github.com/pulumi/pulumi/sdk/v3 v3.23.2
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20211230170131-3a7c83bfab87
	github.com/hashicorp/terraform-provider-vault => github.com/pulumi/terraform-provider-vault v1.9.1-0.20220303160649-7a58699a688c
)
