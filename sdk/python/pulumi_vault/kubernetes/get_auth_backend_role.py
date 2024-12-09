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
    'GetAuthBackendRoleResult',
    'AwaitableGetAuthBackendRoleResult',
    'get_auth_backend_role',
    'get_auth_backend_role_output',
]

@pulumi.output_type
class GetAuthBackendRoleResult:
    """
    A collection of values returned by getAuthBackendRole.
    """
    def __init__(__self__, alias_name_source=None, audience=None, backend=None, bound_service_account_names=None, bound_service_account_namespaces=None, id=None, namespace=None, role_name=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None):
        if alias_name_source and not isinstance(alias_name_source, str):
            raise TypeError("Expected argument 'alias_name_source' to be a str")
        pulumi.set(__self__, "alias_name_source", alias_name_source)
        if audience and not isinstance(audience, str):
            raise TypeError("Expected argument 'audience' to be a str")
        pulumi.set(__self__, "audience", audience)
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if bound_service_account_names and not isinstance(bound_service_account_names, list):
            raise TypeError("Expected argument 'bound_service_account_names' to be a list")
        pulumi.set(__self__, "bound_service_account_names", bound_service_account_names)
        if bound_service_account_namespaces and not isinstance(bound_service_account_namespaces, list):
            raise TypeError("Expected argument 'bound_service_account_namespaces' to be a list")
        pulumi.set(__self__, "bound_service_account_namespaces", bound_service_account_namespaces)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        pulumi.set(__self__, "role_name", role_name)
        if token_bound_cidrs and not isinstance(token_bound_cidrs, list):
            raise TypeError("Expected argument 'token_bound_cidrs' to be a list")
        pulumi.set(__self__, "token_bound_cidrs", token_bound_cidrs)
        if token_explicit_max_ttl and not isinstance(token_explicit_max_ttl, int):
            raise TypeError("Expected argument 'token_explicit_max_ttl' to be a int")
        pulumi.set(__self__, "token_explicit_max_ttl", token_explicit_max_ttl)
        if token_max_ttl and not isinstance(token_max_ttl, int):
            raise TypeError("Expected argument 'token_max_ttl' to be a int")
        pulumi.set(__self__, "token_max_ttl", token_max_ttl)
        if token_no_default_policy and not isinstance(token_no_default_policy, bool):
            raise TypeError("Expected argument 'token_no_default_policy' to be a bool")
        pulumi.set(__self__, "token_no_default_policy", token_no_default_policy)
        if token_num_uses and not isinstance(token_num_uses, int):
            raise TypeError("Expected argument 'token_num_uses' to be a int")
        pulumi.set(__self__, "token_num_uses", token_num_uses)
        if token_period and not isinstance(token_period, int):
            raise TypeError("Expected argument 'token_period' to be a int")
        pulumi.set(__self__, "token_period", token_period)
        if token_policies and not isinstance(token_policies, list):
            raise TypeError("Expected argument 'token_policies' to be a list")
        pulumi.set(__self__, "token_policies", token_policies)
        if token_ttl and not isinstance(token_ttl, int):
            raise TypeError("Expected argument 'token_ttl' to be a int")
        pulumi.set(__self__, "token_ttl", token_ttl)
        if token_type and not isinstance(token_type, str):
            raise TypeError("Expected argument 'token_type' to be a str")
        pulumi.set(__self__, "token_type", token_type)

    @property
    @pulumi.getter(name="aliasNameSource")
    def alias_name_source(self) -> str:
        """
        Method used for generating identity aliases. (vault-1.9+)
        """
        return pulumi.get(self, "alias_name_source")

    @property
    @pulumi.getter
    def audience(self) -> Optional[str]:
        """
        Audience claim to verify in the JWT.
        """
        return pulumi.get(self, "audience")

    @property
    @pulumi.getter
    def backend(self) -> Optional[str]:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="boundServiceAccountNames")
    def bound_service_account_names(self) -> Sequence[str]:
        """
        List of service account names able to access this role. If set to "*" all names are allowed, both this and bound_service_account_namespaces can not be "*".
        """
        return pulumi.get(self, "bound_service_account_names")

    @property
    @pulumi.getter(name="boundServiceAccountNamespaces")
    def bound_service_account_namespaces(self) -> Sequence[str]:
        """
        List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
        """
        return pulumi.get(self, "bound_service_account_namespaces")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="tokenBoundCidrs")
    def token_bound_cidrs(self) -> Optional[Sequence[str]]:
        """
        List of CIDR blocks; if set, specifies blocks of IP
        addresses which can authenticate successfully, and ties the resulting token to these blocks
        as well.
        """
        return pulumi.get(self, "token_bound_cidrs")

    @property
    @pulumi.getter(name="tokenExplicitMaxTtl")
    def token_explicit_max_ttl(self) -> Optional[int]:
        """
        If set, will encode an
        [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        `token_max_ttl` would otherwise allow a renewal.
        """
        return pulumi.get(self, "token_explicit_max_ttl")

    @property
    @pulumi.getter(name="tokenMaxTtl")
    def token_max_ttl(self) -> Optional[int]:
        """
        The maximum lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        return pulumi.get(self, "token_max_ttl")

    @property
    @pulumi.getter(name="tokenNoDefaultPolicy")
    def token_no_default_policy(self) -> Optional[bool]:
        """
        If set, the default policy will not be set on
        generated tokens; otherwise it will be added to the policies set in token_policies.
        """
        return pulumi.get(self, "token_no_default_policy")

    @property
    @pulumi.getter(name="tokenNumUses")
    def token_num_uses(self) -> Optional[int]:
        """
        The
        [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        if any, in number of seconds to set on the token.
        """
        return pulumi.get(self, "token_num_uses")

    @property
    @pulumi.getter(name="tokenPeriod")
    def token_period(self) -> Optional[int]:
        """
        (Optional) If set, indicates that the
        token generated using this role should never expire. The token should be renewed within the
        duration specified by this value. At each renewal, the token's TTL will be set to the
        value of this field. Specified in seconds.
        """
        return pulumi.get(self, "token_period")

    @property
    @pulumi.getter(name="tokenPolicies")
    def token_policies(self) -> Optional[Sequence[str]]:
        """
        List of policies to encode onto generated tokens. Depending
        on the auth method, this list may be supplemented by user/group/other values.
        """
        return pulumi.get(self, "token_policies")

    @property
    @pulumi.getter(name="tokenTtl")
    def token_ttl(self) -> Optional[int]:
        """
        The incremental lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        return pulumi.get(self, "token_ttl")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[str]:
        """
        The type of token that should be generated. Can be `service`,
        `batch`, or `default` to use the mount's tuned default (which unless changed will be
        `service` tokens). For token store roles, there are two additional possibilities:
        `default-service` and `default-batch` which specify the type to return unless the client
        requests a different type at generation time.
        """
        return pulumi.get(self, "token_type")


class AwaitableGetAuthBackendRoleResult(GetAuthBackendRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendRoleResult(
            alias_name_source=self.alias_name_source,
            audience=self.audience,
            backend=self.backend,
            bound_service_account_names=self.bound_service_account_names,
            bound_service_account_namespaces=self.bound_service_account_namespaces,
            id=self.id,
            namespace=self.namespace,
            role_name=self.role_name,
            token_bound_cidrs=self.token_bound_cidrs,
            token_explicit_max_ttl=self.token_explicit_max_ttl,
            token_max_ttl=self.token_max_ttl,
            token_no_default_policy=self.token_no_default_policy,
            token_num_uses=self.token_num_uses,
            token_period=self.token_period,
            token_policies=self.token_policies,
            token_ttl=self.token_ttl,
            token_type=self.token_type)


def get_auth_backend_role(audience: Optional[str] = None,
                          backend: Optional[str] = None,
                          namespace: Optional[str] = None,
                          role_name: Optional[str] = None,
                          token_bound_cidrs: Optional[Sequence[str]] = None,
                          token_explicit_max_ttl: Optional[int] = None,
                          token_max_ttl: Optional[int] = None,
                          token_no_default_policy: Optional[bool] = None,
                          token_num_uses: Optional[int] = None,
                          token_period: Optional[int] = None,
                          token_policies: Optional[Sequence[str]] = None,
                          token_ttl: Optional[int] = None,
                          token_type: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthBackendRoleResult:
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
    information.


    :param str audience: Audience claim to verify in the JWT.
    :param str backend: The unique name for the Kubernetes backend the role to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured namespace.
           *Available only for Vault Enterprise*.
    :param str role_name: The name of the role to retrieve the Role attributes for.
    :param Sequence[str] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
           addresses which can authenticate successfully, and ties the resulting token to these blocks
           as well.
    :param int token_explicit_max_ttl: If set, will encode an
           [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
           onto the token in number of seconds. This is a hard cap even if `token_ttl` and
           `token_max_ttl` would otherwise allow a renewal.
    :param int token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param bool token_no_default_policy: If set, the default policy will not be set on
           generated tokens; otherwise it will be added to the policies set in token_policies.
    :param int token_num_uses: The
           [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
           if any, in number of seconds to set on the token.
    :param int token_period: (Optional) If set, indicates that the
           token generated using this role should never expire. The token should be renewed within the
           duration specified by this value. At each renewal, the token's TTL will be set to the
           value of this field. Specified in seconds.
    :param Sequence[str] token_policies: List of policies to encode onto generated tokens. Depending
           on the auth method, this list may be supplemented by user/group/other values.
    :param int token_ttl: The incremental lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param str token_type: The type of token that should be generated. Can be `service`,
           `batch`, or `default` to use the mount's tuned default (which unless changed will be
           `service` tokens). For token store roles, there are two additional possibilities:
           `default-service` and `default-batch` which specify the type to return unless the client
           requests a different type at generation time.
    """
    __args__ = dict()
    __args__['audience'] = audience
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['roleName'] = role_name
    __args__['tokenBoundCidrs'] = token_bound_cidrs
    __args__['tokenExplicitMaxTtl'] = token_explicit_max_ttl
    __args__['tokenMaxTtl'] = token_max_ttl
    __args__['tokenNoDefaultPolicy'] = token_no_default_policy
    __args__['tokenNumUses'] = token_num_uses
    __args__['tokenPeriod'] = token_period
    __args__['tokenPolicies'] = token_policies
    __args__['tokenTtl'] = token_ttl
    __args__['tokenType'] = token_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getAuthBackendRole:getAuthBackendRole', __args__, opts=opts, typ=GetAuthBackendRoleResult).value

    return AwaitableGetAuthBackendRoleResult(
        alias_name_source=pulumi.get(__ret__, 'alias_name_source'),
        audience=pulumi.get(__ret__, 'audience'),
        backend=pulumi.get(__ret__, 'backend'),
        bound_service_account_names=pulumi.get(__ret__, 'bound_service_account_names'),
        bound_service_account_namespaces=pulumi.get(__ret__, 'bound_service_account_namespaces'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        role_name=pulumi.get(__ret__, 'role_name'),
        token_bound_cidrs=pulumi.get(__ret__, 'token_bound_cidrs'),
        token_explicit_max_ttl=pulumi.get(__ret__, 'token_explicit_max_ttl'),
        token_max_ttl=pulumi.get(__ret__, 'token_max_ttl'),
        token_no_default_policy=pulumi.get(__ret__, 'token_no_default_policy'),
        token_num_uses=pulumi.get(__ret__, 'token_num_uses'),
        token_period=pulumi.get(__ret__, 'token_period'),
        token_policies=pulumi.get(__ret__, 'token_policies'),
        token_ttl=pulumi.get(__ret__, 'token_ttl'),
        token_type=pulumi.get(__ret__, 'token_type'))
def get_auth_backend_role_output(audience: Optional[pulumi.Input[Optional[str]]] = None,
                                 backend: Optional[pulumi.Input[Optional[str]]] = None,
                                 namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                 role_name: Optional[pulumi.Input[str]] = None,
                                 token_bound_cidrs: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 token_explicit_max_ttl: Optional[pulumi.Input[Optional[int]]] = None,
                                 token_max_ttl: Optional[pulumi.Input[Optional[int]]] = None,
                                 token_no_default_policy: Optional[pulumi.Input[Optional[bool]]] = None,
                                 token_num_uses: Optional[pulumi.Input[Optional[int]]] = None,
                                 token_period: Optional[pulumi.Input[Optional[int]]] = None,
                                 token_policies: Optional[pulumi.Input[Optional[Sequence[str]]]] = None,
                                 token_ttl: Optional[pulumi.Input[Optional[int]]] = None,
                                 token_type: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAuthBackendRoleResult]:
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
    information.


    :param str audience: Audience claim to verify in the JWT.
    :param str backend: The unique name for the Kubernetes backend the role to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured namespace.
           *Available only for Vault Enterprise*.
    :param str role_name: The name of the role to retrieve the Role attributes for.
    :param Sequence[str] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
           addresses which can authenticate successfully, and ties the resulting token to these blocks
           as well.
    :param int token_explicit_max_ttl: If set, will encode an
           [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
           onto the token in number of seconds. This is a hard cap even if `token_ttl` and
           `token_max_ttl` would otherwise allow a renewal.
    :param int token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param bool token_no_default_policy: If set, the default policy will not be set on
           generated tokens; otherwise it will be added to the policies set in token_policies.
    :param int token_num_uses: The
           [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
           if any, in number of seconds to set on the token.
    :param int token_period: (Optional) If set, indicates that the
           token generated using this role should never expire. The token should be renewed within the
           duration specified by this value. At each renewal, the token's TTL will be set to the
           value of this field. Specified in seconds.
    :param Sequence[str] token_policies: List of policies to encode onto generated tokens. Depending
           on the auth method, this list may be supplemented by user/group/other values.
    :param int token_ttl: The incremental lifetime for generated tokens in number of seconds.
           Its current value will be referenced at renewal time.
    :param str token_type: The type of token that should be generated. Can be `service`,
           `batch`, or `default` to use the mount's tuned default (which unless changed will be
           `service` tokens). For token store roles, there are two additional possibilities:
           `default-service` and `default-batch` which specify the type to return unless the client
           requests a different type at generation time.
    """
    __args__ = dict()
    __args__['audience'] = audience
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['roleName'] = role_name
    __args__['tokenBoundCidrs'] = token_bound_cidrs
    __args__['tokenExplicitMaxTtl'] = token_explicit_max_ttl
    __args__['tokenMaxTtl'] = token_max_ttl
    __args__['tokenNoDefaultPolicy'] = token_no_default_policy
    __args__['tokenNumUses'] = token_num_uses
    __args__['tokenPeriod'] = token_period
    __args__['tokenPolicies'] = token_policies
    __args__['tokenTtl'] = token_ttl
    __args__['tokenType'] = token_type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:kubernetes/getAuthBackendRole:getAuthBackendRole', __args__, opts=opts, typ=GetAuthBackendRoleResult)
    return __ret__.apply(lambda __response__: GetAuthBackendRoleResult(
        alias_name_source=pulumi.get(__response__, 'alias_name_source'),
        audience=pulumi.get(__response__, 'audience'),
        backend=pulumi.get(__response__, 'backend'),
        bound_service_account_names=pulumi.get(__response__, 'bound_service_account_names'),
        bound_service_account_namespaces=pulumi.get(__response__, 'bound_service_account_namespaces'),
        id=pulumi.get(__response__, 'id'),
        namespace=pulumi.get(__response__, 'namespace'),
        role_name=pulumi.get(__response__, 'role_name'),
        token_bound_cidrs=pulumi.get(__response__, 'token_bound_cidrs'),
        token_explicit_max_ttl=pulumi.get(__response__, 'token_explicit_max_ttl'),
        token_max_ttl=pulumi.get(__response__, 'token_max_ttl'),
        token_no_default_policy=pulumi.get(__response__, 'token_no_default_policy'),
        token_num_uses=pulumi.get(__response__, 'token_num_uses'),
        token_period=pulumi.get(__response__, 'token_period'),
        token_policies=pulumi.get(__response__, 'token_policies'),
        token_ttl=pulumi.get(__response__, 'token_ttl'),
        token_type=pulumi.get(__response__, 'token_type')))
