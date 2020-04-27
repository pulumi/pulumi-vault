CHANGELOG
=========

## HEAD (Unreleased)
* Regenerate datasource examples to be async

---

## 2.0.0 (2020-04-17)
* Upgrade to Pulumi v2.0.0
* Upgrade to pulumi-terraform-bridge v2.0.0

## 1.11.0 (2020-04-14)
* Upgrade to v2.10.0 of the Vault Terraform Provider

## 1.10.0 (2020-03-23)
* Ensure JavaScript dependency for pulumi/pulumi isn't pinned to latest
* Upgrade to v2.9.0 of the Vault Terraform Provider
* Upgrade to Pulumi v1.12.1
* Upgrade to pulumi-terraform-bridge v1.8.2

## 1.9.0 (2020-02-07)
* Upgrade to v2.8.0 of the Vault Terraform Provider

## 1.8.0 (2020-01-29)
* Upgrade to pulumi-terraform-bridge v1.6.4

## 1.7.1 (2020-01-24)
* Upgrade default value for `max_lease_ttl_seconds` to 1200

## 1.7.0 (2020-01-06)
* Upgrade to pulumi-terraform-bridge v1.5.2
* Upgrade to v2.7.1 of the Vault Terraform Provider

## 1.6.0 (2019-12-18)
* Upgrade to v2.7.0 of the Vault Terraform Provider

## 1.5.0 (2019-12-11)

#### BREAKING CHANGES
* `vault.Token` module has been renamed to `vault.TokenAuth` this is to avoid a conflict between Token class and Token module
* `vault.lDAP` module has been renamed to `vault.ldap`

#### Improvements / Bug Fixes
* Remove the autonaming of `vault.identify.GroupAlias`. The `name` must match that of the Github Team Name.
* Namespace names in .NET SDK are adjusted to PascalCase
([#17](https://github.com/pulumi/pulumi-vault/pull/17)).

## 1.4.0 (2019-12-04)
* Upgrade to pulumi-terraform-bridge v1.4.3

## 1.3.0 (2019-11-27)
* Upgrade to support go 1.13.x
* Upgrade to pulumi-terraform-bridge v1.4.2
* Upgrade to v2.6.0 of the Vault Terraform Provider

## 1.2.0 (2019-11-15)
* Add support for DotNet SDK Generation

## 1.1.0 (2019-10-24)
* Upgrade to v2.5.0 of the Vault Terraform Provider
* Upgrade to v1.1.0 of the Pulumi-Terraform Bridge

## 1.0.0 (2019-10-05)
* Initial release of the provider
