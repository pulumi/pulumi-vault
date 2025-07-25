# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities
from . import outputs

addAddressToEnv: Optional[str]

address: Optional[str]
"""
URL of the root of the target Vault server.
"""

authLogin: Optional[str]
"""
Login to vault with an existing auth method using auth/<mount>/login
"""

authLoginAws: Optional[str]
"""
Login to vault using the AWS method
"""

authLoginAzure: Optional[str]
"""
Login to vault using the azure method
"""

authLoginCert: Optional[str]
"""
Login to vault using the cert method
"""

authLoginGcp: Optional[str]
"""
Login to vault using the gcp method
"""

authLoginJwt: Optional[str]
"""
Login to vault using the jwt method
"""

authLoginKerberos: Optional[str]
"""
Login to vault using the kerberos method
"""

authLoginOci: Optional[str]
"""
Login to vault using the OCI method
"""

authLoginOidc: Optional[str]
"""
Login to vault using the oidc method
"""

authLoginRadius: Optional[str]
"""
Login to vault using the radius method
"""

authLoginTokenFile: Optional[str]
"""
Login to vault using
"""

authLoginUserpass: Optional[str]
"""
Login to vault using the userpass method
"""

caCertDir: Optional[str]
"""
Path to directory containing CA certificate files to validate the server's certificate.
"""

caCertFile: Optional[str]
"""
Path to a CA certificate file to validate the server's certificate.
"""

clientAuth: Optional[str]
"""
Client authentication credentials.
"""

headers: Optional[str]
"""
The headers to send with each Vault request.
"""

maxLeaseTtlSeconds: int
"""
Maximum TTL for secret leases requested by this provider.
"""

maxRetries: int
"""
Maximum number of retries when a 5xx error code is encountered.
"""

maxRetriesCcc: Optional[int]
"""
Maximum number of retries for Client Controlled Consistency related operations
"""

namespace: Optional[str]
"""
The namespace to use. Available only for Vault Enterprise.
"""

setNamespaceFromToken: Optional[bool]
"""
In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
token namespace as the root namespace for all resources.
"""

skipChildToken: Optional[bool]
"""
Set this to true to prevent the creation of ephemeral child token used by this provider.
"""

skipGetVaultVersion: Optional[bool]
"""
Skip the dynamic fetching of the Vault server version.
"""

skipTlsVerify: Optional[bool]
"""
Set this to true only if the target Vault server is an insecure development instance.
"""

tlsServerName: Optional[str]
"""
Name to use as the SNI host when connecting via TLS.
"""

token: Optional[str]
"""
Token to use to authenticate to Vault.
"""

tokenName: Optional[str]
"""
Token name to use for creating the Vault child token.
"""

vaultVersionOverride: Optional[str]
"""
Override the Vault server version, which is normally determined dynamically from the target Vault server
"""

