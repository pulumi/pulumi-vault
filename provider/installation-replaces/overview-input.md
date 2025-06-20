The Vault provider allows Terraform to read from, write to, and configure
[HashiCorp Vault](https://developer.hashicorp.com/vault).

~> **Important** Interacting with Vault from Terraform causes any secrets
that you read and write to be persisted in both Terraform's state file
*and* in any generated plan files. For any Terraform module that reads or
writes Vault secrets, these files should be treated as sensitive and
protected accordingly.

This provider serves two pretty-distinct use-cases, which each have their
own security trade-offs and caveats that are covered in the sections that
follow. Consider these carefully before using this provider within your
Terraform configuration.

-> Visit the [Inject secrets into Terraform using the Vault provider](https://learn.hashicorp.com/tutorials/terraform/secrets-vault?utm_source=WEBSITE&utm_medium=WEB_IO&utm_offer=ARTICLE_PAGE&utm_content=DOCS) Learn tutorial to learn how to use
short-lived credentials from Vault's AWS Secrets Engine to authenticate the
AWS provider.