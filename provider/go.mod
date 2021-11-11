module github.com/pulumi/pulumi-vault/provider/v4

go 1.16

require (
	github.com/hashicorp/terraform-provider-vault v1.9.1-0.20211005125950-9dff41bf9149
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.11.0
	github.com/pulumi/pulumi/sdk/v3 v3.17.0
)

replace (
	github.com/hashicorp/go-getter => github.com/hashicorp/go-getter v1.4.0
	github.com/hashicorp/terraform-provider-vault => github.com/pulumi/terraform-provider-vault v1.9.1-0.20211005183512-18f85787fd84
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
