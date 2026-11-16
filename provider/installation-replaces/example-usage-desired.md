## Example Usage

Configure the provider through the `VAULT_ADDR` and `VAULT_TOKEN` environment variables, so that each user can keep separate credentials in their own environment.
To log in with an auth method instead of a token, see the login examples below.

```bash
export VAULT_ADDR=https://vault.example.net:8200
export VAULT_TOKEN=<token>
```

```hcl
