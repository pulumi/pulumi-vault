CHANGELOG
=========

## HEAD (Unreleased)
* Upgrade to v2.7.0 of the Vault Terraform Provider

---

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
