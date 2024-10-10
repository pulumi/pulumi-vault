~> **Important** It is important to ensure that the Vault token
has a long enough `time-to-live` to allow for all Vault resources to
be successfully provisioned. In the case where the `TTL` is insufficient,
you may encounter unexpected permission denied errors.
See [Vault Token TTLs](https://developer.hashicorp.com/vault/docs/concepts/tokens#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
for more details.

Most Terraform providers require credentials to interact with a third-party
service that they wrap. This provider allows such credentials to be obtained
from Vault, which means that operators or systems running Terraform need
only access to a suitably-privileged Vault token in order to temporarily
lease the credentials for other providers.

Currently, Terraform has no mechanism to redact or protect secrets that
are returned via data sources, so secrets read via this provider will be
persisted into the Terraform state, into any plan files, and in some cases
in the console output produced while planning and applying. These artifacts
must therefore all be protected accordingly.

To reduce the exposure of such secrets, the provider requests a Vault token
with a relatively-short TTL (20 minutes, by default) which in turn means
that where possible Vault will revoke any issued credentials after that
time, but in particular it is unable to retract any static secrets such as
those stored in Vault's "generic" secret backend.

The requested token TTL can be controlled by the `max_lease_ttl_seconds`
provider argument described below. It is important to consider that Terraform
reads from data sources during the `plan` phase and writes the result into
the plan. Thus, a subsequent `apply` will likely fail if it is run after the
intermediate token has expired, due to the revocation of the secrets that
are stored in the plan.

Except as otherwise noted, the resources that read secrets from Vault
are designed such that they require only the *read* capability on the relevant
resources.
