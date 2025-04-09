# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['AuthBackendGroupArgs', 'AuthBackendGroup']

@pulumi.input_type
class AuthBackendGroupArgs:
    def __init__(__self__, *,
                 groupname: pulumi.Input[builtins.str],
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        The set of arguments for constructing a AuthBackendGroup resource.
        :param pulumi.Input[builtins.str] groupname: The LDAP groupname
        :param pulumi.Input[builtins.str] backend: Path to the authentication backend
               
               For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: Policies which should be granted to members of the group
        """
        pulumi.set(__self__, "groupname", groupname)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def groupname(self) -> pulumi.Input[builtins.str]:
        """
        The LDAP groupname
        """
        return pulumi.get(self, "groupname")

    @groupname.setter
    def groupname(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "groupname", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Path to the authentication backend

        For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Policies which should be granted to members of the group
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _AuthBackendGroupState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 groupname: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering AuthBackendGroup resources.
        :param pulumi.Input[builtins.str] backend: Path to the authentication backend
               
               For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        :param pulumi.Input[builtins.str] groupname: The LDAP groupname
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: Policies which should be granted to members of the group
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if groupname is not None:
            pulumi.set(__self__, "groupname", groupname)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Path to the authentication backend

        For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def groupname(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The LDAP groupname
        """
        return pulumi.get(self, "groupname")

    @groupname.setter
    def groupname(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "groupname", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Policies which should be granted to members of the group
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "policies", value)


class AuthBackendGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 groupname: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        ldap = vault.ldap.AuthBackend("ldap",
            path="ldap",
            url="ldaps://dc-01.example.org",
            userdn="OU=Users,OU=Accounts,DC=example,DC=org",
            userattr="sAMAccountName",
            upndomain="EXAMPLE.ORG",
            discoverdn=False,
            groupdn="OU=Groups,DC=example,DC=org",
            groupfilter="(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))")
        group = vault.ldap.AuthBackendGroup("group",
            groupname="dba",
            policies=["dba"],
            backend=ldap.path)
        ```

        ## Import

        LDAP authentication backend groups can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend: Path to the authentication backend
               
               For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        :param pulumi.Input[builtins.str] groupname: The LDAP groupname
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: Policies which should be granted to members of the group
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a group in an [LDAP auth backend within Vault](https://www.vaultproject.io/docs/auth/ldap.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        ldap = vault.ldap.AuthBackend("ldap",
            path="ldap",
            url="ldaps://dc-01.example.org",
            userdn="OU=Users,OU=Accounts,DC=example,DC=org",
            userattr="sAMAccountName",
            upndomain="EXAMPLE.ORG",
            discoverdn=False,
            groupdn="OU=Groups,DC=example,DC=org",
            groupfilter="(&(objectClass=group)(member:1.2.840.113556.1.4.1941:={{.UserDN}}))")
        group = vault.ldap.AuthBackendGroup("group",
            groupname="dba",
            policies=["dba"],
            backend=ldap.path)
        ```

        ## Import

        LDAP authentication backend groups can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:ldap/authBackendGroup:AuthBackendGroup foo auth/ldap/groups/foo
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 groupname: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendGroupArgs.__new__(AuthBackendGroupArgs)

            __props__.__dict__["backend"] = backend
            if groupname is None and not opts.urn:
                raise TypeError("Missing required property 'groupname'")
            __props__.__dict__["groupname"] = groupname
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["policies"] = policies
        super(AuthBackendGroup, __self__).__init__(
            'vault:ldap/authBackendGroup:AuthBackendGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[builtins.str]] = None,
            groupname: Optional[pulumi.Input[builtins.str]] = None,
            namespace: Optional[pulumi.Input[builtins.str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None) -> 'AuthBackendGroup':
        """
        Get an existing AuthBackendGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend: Path to the authentication backend
               
               For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        :param pulumi.Input[builtins.str] groupname: The LDAP groupname
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: Policies which should be granted to members of the group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendGroupState.__new__(_AuthBackendGroupState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["groupname"] = groupname
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["policies"] = policies
        return AuthBackendGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Path to the authentication backend

        For more details on the usage of each argument consult the [Vault LDAP API documentation](https://www.vaultproject.io/api-docs/auth/ldap).
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def groupname(self) -> pulumi.Output[builtins.str]:
        """
        The LDAP groupname
        """
        return pulumi.get(self, "groupname")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Policies which should be granted to members of the group
        """
        return pulumi.get(self, "policies")

