# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
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

__all__ = [
    'GetOidcOpenidConfigResult',
    'AwaitableGetOidcOpenidConfigResult',
    'get_oidc_openid_config',
    'get_oidc_openid_config_output',
]

@pulumi.output_type
class GetOidcOpenidConfigResult:
    """
    A collection of values returned by getOidcOpenidConfig.
    """
    def __init__(__self__, authorization_endpoint=None, grant_types_supporteds=None, id=None, id_token_signing_alg_values_supporteds=None, issuer=None, jwks_uri=None, name=None, namespace=None, request_uri_parameter_supported=None, response_types_supporteds=None, scopes_supporteds=None, subject_types_supporteds=None, token_endpoint=None, token_endpoint_auth_methods_supporteds=None, userinfo_endpoint=None):
        if authorization_endpoint and not isinstance(authorization_endpoint, str):
            raise TypeError("Expected argument 'authorization_endpoint' to be a str")
        pulumi.set(__self__, "authorization_endpoint", authorization_endpoint)
        if grant_types_supporteds and not isinstance(grant_types_supporteds, list):
            raise TypeError("Expected argument 'grant_types_supporteds' to be a list")
        pulumi.set(__self__, "grant_types_supporteds", grant_types_supporteds)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if id_token_signing_alg_values_supporteds and not isinstance(id_token_signing_alg_values_supporteds, list):
            raise TypeError("Expected argument 'id_token_signing_alg_values_supporteds' to be a list")
        pulumi.set(__self__, "id_token_signing_alg_values_supporteds", id_token_signing_alg_values_supporteds)
        if issuer and not isinstance(issuer, str):
            raise TypeError("Expected argument 'issuer' to be a str")
        pulumi.set(__self__, "issuer", issuer)
        if jwks_uri and not isinstance(jwks_uri, str):
            raise TypeError("Expected argument 'jwks_uri' to be a str")
        pulumi.set(__self__, "jwks_uri", jwks_uri)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if request_uri_parameter_supported and not isinstance(request_uri_parameter_supported, bool):
            raise TypeError("Expected argument 'request_uri_parameter_supported' to be a bool")
        pulumi.set(__self__, "request_uri_parameter_supported", request_uri_parameter_supported)
        if response_types_supporteds and not isinstance(response_types_supporteds, list):
            raise TypeError("Expected argument 'response_types_supporteds' to be a list")
        pulumi.set(__self__, "response_types_supporteds", response_types_supporteds)
        if scopes_supporteds and not isinstance(scopes_supporteds, list):
            raise TypeError("Expected argument 'scopes_supporteds' to be a list")
        pulumi.set(__self__, "scopes_supporteds", scopes_supporteds)
        if subject_types_supporteds and not isinstance(subject_types_supporteds, list):
            raise TypeError("Expected argument 'subject_types_supporteds' to be a list")
        pulumi.set(__self__, "subject_types_supporteds", subject_types_supporteds)
        if token_endpoint and not isinstance(token_endpoint, str):
            raise TypeError("Expected argument 'token_endpoint' to be a str")
        pulumi.set(__self__, "token_endpoint", token_endpoint)
        if token_endpoint_auth_methods_supporteds and not isinstance(token_endpoint_auth_methods_supporteds, list):
            raise TypeError("Expected argument 'token_endpoint_auth_methods_supporteds' to be a list")
        pulumi.set(__self__, "token_endpoint_auth_methods_supporteds", token_endpoint_auth_methods_supporteds)
        if userinfo_endpoint and not isinstance(userinfo_endpoint, str):
            raise TypeError("Expected argument 'userinfo_endpoint' to be a str")
        pulumi.set(__self__, "userinfo_endpoint", userinfo_endpoint)

    @property
    @pulumi.getter(name="authorizationEndpoint")
    def authorization_endpoint(self) -> str:
        """
        The Authorization Endpoint for the provider.
        """
        return pulumi.get(self, "authorization_endpoint")

    @property
    @pulumi.getter(name="grantTypesSupporteds")
    def grant_types_supporteds(self) -> Sequence[str]:
        """
        The grant types supported by the provider.
        """
        return pulumi.get(self, "grant_types_supporteds")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="idTokenSigningAlgValuesSupporteds")
    def id_token_signing_alg_values_supporteds(self) -> Sequence[str]:
        """
        The signing algorithms supported by 
        the provider.
        """
        return pulumi.get(self, "id_token_signing_alg_values_supporteds")

    @property
    @pulumi.getter
    def issuer(self) -> str:
        """
        The URL of the issuer for the provider.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="jwksUri")
    def jwks_uri(self) -> str:
        """
        The well known keys URI for the provider.
        """
        return pulumi.get(self, "jwks_uri")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="requestUriParameterSupported")
    def request_uri_parameter_supported(self) -> bool:
        """
        Specifies whether Request URI Parameter is 
        supported by the provider.
        """
        return pulumi.get(self, "request_uri_parameter_supported")

    @property
    @pulumi.getter(name="responseTypesSupporteds")
    def response_types_supporteds(self) -> Sequence[str]:
        """
        The response types supported by the provider.
        """
        return pulumi.get(self, "response_types_supporteds")

    @property
    @pulumi.getter(name="scopesSupporteds")
    def scopes_supporteds(self) -> Sequence[str]:
        """
        The scopes supported by the provider.
        """
        return pulumi.get(self, "scopes_supporteds")

    @property
    @pulumi.getter(name="subjectTypesSupporteds")
    def subject_types_supporteds(self) -> Sequence[str]:
        """
        The subject types supported by the provider.
        """
        return pulumi.get(self, "subject_types_supporteds")

    @property
    @pulumi.getter(name="tokenEndpoint")
    def token_endpoint(self) -> str:
        """
        The Token Endpoint for the provider.
        """
        return pulumi.get(self, "token_endpoint")

    @property
    @pulumi.getter(name="tokenEndpointAuthMethodsSupporteds")
    def token_endpoint_auth_methods_supporteds(self) -> Sequence[str]:
        """
        The token endpoint auth methods supported by the provider.
        """
        return pulumi.get(self, "token_endpoint_auth_methods_supporteds")

    @property
    @pulumi.getter(name="userinfoEndpoint")
    def userinfo_endpoint(self) -> str:
        """
        The User Info Endpoint for the provider
        """
        return pulumi.get(self, "userinfo_endpoint")


class AwaitableGetOidcOpenidConfigResult(GetOidcOpenidConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOidcOpenidConfigResult(
            authorization_endpoint=self.authorization_endpoint,
            grant_types_supporteds=self.grant_types_supporteds,
            id=self.id,
            id_token_signing_alg_values_supporteds=self.id_token_signing_alg_values_supporteds,
            issuer=self.issuer,
            jwks_uri=self.jwks_uri,
            name=self.name,
            namespace=self.namespace,
            request_uri_parameter_supported=self.request_uri_parameter_supported,
            response_types_supporteds=self.response_types_supporteds,
            scopes_supporteds=self.scopes_supporteds,
            subject_types_supporteds=self.subject_types_supporteds,
            token_endpoint=self.token_endpoint,
            token_endpoint_auth_methods_supporteds=self.token_endpoint_auth_methods_supporteds,
            userinfo_endpoint=self.userinfo_endpoint)


def get_oidc_openid_config(name: Optional[str] = None,
                           namespace: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOidcOpenidConfigResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    key = vault.identity.OidcKey("key",
        name="key",
        allowed_client_ids=["*"],
        rotation_period=3600,
        verification_ttl=3600)
    app = vault.identity.OidcClient("app",
        name="application",
        key=key.name,
        redirect_uris=[
            "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
            "http://127.0.0.1:8251/callback",
            "http://127.0.0.1:8080/callback",
        ],
        id_token_ttl=2400,
        access_token_ttl=7200)
    provider = vault.identity.OidcProvider("provider",
        name="provider",
        allowed_client_ids=[test["clientId"]])
    config = vault.identity.get_oidc_openid_config_output(name=provider.name)
    ```


    :param str name: The name of the OIDC Provider in Vault.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:identity/getOidcOpenidConfig:getOidcOpenidConfig', __args__, opts=opts, typ=GetOidcOpenidConfigResult).value

    return AwaitableGetOidcOpenidConfigResult(
        authorization_endpoint=pulumi.get(__ret__, 'authorization_endpoint'),
        grant_types_supporteds=pulumi.get(__ret__, 'grant_types_supporteds'),
        id=pulumi.get(__ret__, 'id'),
        id_token_signing_alg_values_supporteds=pulumi.get(__ret__, 'id_token_signing_alg_values_supporteds'),
        issuer=pulumi.get(__ret__, 'issuer'),
        jwks_uri=pulumi.get(__ret__, 'jwks_uri'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        request_uri_parameter_supported=pulumi.get(__ret__, 'request_uri_parameter_supported'),
        response_types_supporteds=pulumi.get(__ret__, 'response_types_supporteds'),
        scopes_supporteds=pulumi.get(__ret__, 'scopes_supporteds'),
        subject_types_supporteds=pulumi.get(__ret__, 'subject_types_supporteds'),
        token_endpoint=pulumi.get(__ret__, 'token_endpoint'),
        token_endpoint_auth_methods_supporteds=pulumi.get(__ret__, 'token_endpoint_auth_methods_supporteds'),
        userinfo_endpoint=pulumi.get(__ret__, 'userinfo_endpoint'))
def get_oidc_openid_config_output(name: Optional[pulumi.Input[str]] = None,
                                  namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOidcOpenidConfigResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    key = vault.identity.OidcKey("key",
        name="key",
        allowed_client_ids=["*"],
        rotation_period=3600,
        verification_ttl=3600)
    app = vault.identity.OidcClient("app",
        name="application",
        key=key.name,
        redirect_uris=[
            "http://127.0.0.1:9200/v1/auth-methods/oidc:authenticate:callback",
            "http://127.0.0.1:8251/callback",
            "http://127.0.0.1:8080/callback",
        ],
        id_token_ttl=2400,
        access_token_ttl=7200)
    provider = vault.identity.OidcProvider("provider",
        name="provider",
        allowed_client_ids=[test["clientId"]])
    config = vault.identity.get_oidc_openid_config_output(name=provider.name)
    ```


    :param str name: The name of the OIDC Provider in Vault.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:identity/getOidcOpenidConfig:getOidcOpenidConfig', __args__, opts=opts, typ=GetOidcOpenidConfigResult)
    return __ret__.apply(lambda __response__: GetOidcOpenidConfigResult(
        authorization_endpoint=pulumi.get(__response__, 'authorization_endpoint'),
        grant_types_supporteds=pulumi.get(__response__, 'grant_types_supporteds'),
        id=pulumi.get(__response__, 'id'),
        id_token_signing_alg_values_supporteds=pulumi.get(__response__, 'id_token_signing_alg_values_supporteds'),
        issuer=pulumi.get(__response__, 'issuer'),
        jwks_uri=pulumi.get(__response__, 'jwks_uri'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        request_uri_parameter_supported=pulumi.get(__response__, 'request_uri_parameter_supported'),
        response_types_supporteds=pulumi.get(__response__, 'response_types_supporteds'),
        scopes_supporteds=pulumi.get(__response__, 'scopes_supporteds'),
        subject_types_supporteds=pulumi.get(__response__, 'subject_types_supporteds'),
        token_endpoint=pulumi.get(__response__, 'token_endpoint'),
        token_endpoint_auth_methods_supporteds=pulumi.get(__response__, 'token_endpoint_auth_methods_supporteds'),
        userinfo_endpoint=pulumi.get(__response__, 'userinfo_endpoint')))
