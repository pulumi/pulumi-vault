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

import types

__config__ = pulumi.Config('vault')


class _ExportableConfig(types.ModuleType):
    @_builtins.property
    def add_address_to_env(self) -> Optional[str]:
        return __config__.get('addAddressToEnv')

    @_builtins.property
    def address(self) -> Optional[str]:
        """
        URL of the root of the target Vault server.
        """
        return __config__.get('address')

    @_builtins.property
    def auth_login(self) -> Optional[str]:
        """
        Login to vault with an existing auth method using auth/<mount>/login
        """
        return __config__.get('authLogin')

    @_builtins.property
    def auth_login_aws(self) -> Optional[str]:
        """
        Login to vault using the AWS method
        """
        return __config__.get('authLoginAws')

    @_builtins.property
    def auth_login_azure(self) -> Optional[str]:
        """
        Login to vault using the azure method
        """
        return __config__.get('authLoginAzure')

    @_builtins.property
    def auth_login_cert(self) -> Optional[str]:
        """
        Login to vault using the cert method
        """
        return __config__.get('authLoginCert')

    @_builtins.property
    def auth_login_gcp(self) -> Optional[str]:
        """
        Login to vault using the gcp method
        """
        return __config__.get('authLoginGcp')

    @_builtins.property
    def auth_login_jwt(self) -> Optional[str]:
        """
        Login to vault using the jwt method
        """
        return __config__.get('authLoginJwt')

    @_builtins.property
    def auth_login_kerberos(self) -> Optional[str]:
        """
        Login to vault using the kerberos method
        """
        return __config__.get('authLoginKerberos')

    @_builtins.property
    def auth_login_oci(self) -> Optional[str]:
        """
        Login to vault using the OCI method
        """
        return __config__.get('authLoginOci')

    @_builtins.property
    def auth_login_oidc(self) -> Optional[str]:
        """
        Login to vault using the oidc method
        """
        return __config__.get('authLoginOidc')

    @_builtins.property
    def auth_login_radius(self) -> Optional[str]:
        """
        Login to vault using the radius method
        """
        return __config__.get('authLoginRadius')

    @_builtins.property
    def auth_login_token_file(self) -> Optional[str]:
        """
        Login to vault using
        """
        return __config__.get('authLoginTokenFile')

    @_builtins.property
    def auth_login_userpass(self) -> Optional[str]:
        """
        Login to vault using the userpass method
        """
        return __config__.get('authLoginUserpass')

    @_builtins.property
    def ca_cert_dir(self) -> Optional[str]:
        """
        Path to directory containing CA certificate files to validate the server's certificate.
        """
        return __config__.get('caCertDir')

    @_builtins.property
    def ca_cert_file(self) -> Optional[str]:
        """
        Path to a CA certificate file to validate the server's certificate.
        """
        return __config__.get('caCertFile')

    @_builtins.property
    def client_auth(self) -> Optional[str]:
        """
        Client authentication credentials.
        """
        return __config__.get('clientAuth')

    @_builtins.property
    def headers(self) -> Optional[str]:
        """
        The headers to send with each Vault request.
        """
        return __config__.get('headers')

    @_builtins.property
    def max_lease_ttl_seconds(self) -> int:
        """
        Maximum TTL for secret leases requested by this provider.
        """
        return __config__.get_int('maxLeaseTtlSeconds') or (_utilities.get_env_int('TERRAFORM_VAULT_MAX_TTL') or 1200)

    @_builtins.property
    def max_retries(self) -> int:
        """
        Maximum number of retries when a 5xx error code is encountered.
        """
        return __config__.get_int('maxRetries') or (_utilities.get_env_int('VAULT_MAX_RETRIES') or 2)

    @_builtins.property
    def max_retries_ccc(self) -> Optional[int]:
        """
        Maximum number of retries for Client Controlled Consistency related operations
        """
        return __config__.get_int('maxRetriesCcc')

    @_builtins.property
    def namespace(self) -> Optional[str]:
        """
        The namespace to use. Available only for Vault Enterprise.
        """
        return __config__.get('namespace')

    @_builtins.property
    def set_namespace_from_token(self) -> Optional[bool]:
        """
        In the case where the Vault token is for a specific namespace and the provider namespace is not configured, use the
        token namespace as the root namespace for all resources.
        """
        return __config__.get_bool('setNamespaceFromToken')

    @_builtins.property
    def skip_child_token(self) -> Optional[bool]:
        """
        Set this to true to prevent the creation of ephemeral child token used by this provider.
        """
        return __config__.get_bool('skipChildToken')

    @_builtins.property
    def skip_get_vault_version(self) -> Optional[bool]:
        """
        Skip the dynamic fetching of the Vault server version.
        """
        return __config__.get_bool('skipGetVaultVersion')

    @_builtins.property
    def skip_tls_verify(self) -> Optional[bool]:
        """
        Set this to true only if the target Vault server is an insecure development instance.
        """
        return __config__.get_bool('skipTlsVerify') or _utilities.get_env_bool('VAULT_SKIP_VERIFY')

    @_builtins.property
    def tls_server_name(self) -> Optional[str]:
        """
        Name to use as the SNI host when connecting via TLS.
        """
        return __config__.get('tlsServerName')

    @_builtins.property
    def token(self) -> Optional[str]:
        """
        Token to use to authenticate to Vault.
        """
        return __config__.get('token')

    @_builtins.property
    def token_name(self) -> Optional[str]:
        """
        Token name to use for creating the Vault child token.
        """
        return __config__.get('tokenName')

    @_builtins.property
    def vault_version_override(self) -> Optional[str]:
        """
        Override the Vault server version, which is normally determined dynamically from the target Vault server
        """
        return __config__.get('vaultVersionOverride')

