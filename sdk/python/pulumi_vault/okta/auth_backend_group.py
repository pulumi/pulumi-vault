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

__all__ = ['AuthBackendGroupInitArgs', 'AuthBackendGroup']

@pulumi.input_type
class AuthBackendGroupInitArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[_builtins.str],
                 path: pulumi.Input[_builtins.str],
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a AuthBackendGroup resource.
        :param pulumi.Input[_builtins.str] group_name: Name of the group within the Okta
        :param pulumi.Input[_builtins.str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] policies: Vault policies to associate with this group
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "path", path)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @_builtins.property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the group within the Okta
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "group_name", value)

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Input[_builtins.str]:
        """
        The path where the Okta auth backend is mounted
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Vault policies to associate with this group
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _AuthBackendGroupState:
    def __init__(__self__, *,
                 group_name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering AuthBackendGroup resources.
        :param pulumi.Input[_builtins.str] group_name: Name of the group within the Okta
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] policies: Vault policies to associate with this group
        """
        if group_name is not None:
            pulumi.set(__self__, "group_name", group_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @_builtins.property
    @pulumi.getter(name="groupName")
    def group_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the group within the Okta
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "group_name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The path where the Okta auth backend is mounted
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Vault policies to associate with this group
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.type_token("vault:okta/authBackendGroup:AuthBackendGroup")
class AuthBackendGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        Provides a resource to create a group in an
        [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.okta.AuthBackend("example",
            path="group_okta",
            organization="dummy")
        foo = vault.okta.AuthBackendGroup("foo",
            path=example.path,
            group_name="foo",
            policies=[
                "one",
                "two",
            ])
        ```

        ## Import

        Okta authentication backend groups can be imported using the format `backend/groupName` e.g.

        ```sh
        $ pulumi import vault:okta/authBackendGroup:AuthBackendGroup foo okta/foo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] group_name: Name of the group within the Okta
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] policies: Vault policies to associate with this group
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendGroupInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to create a group in an
        [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.okta.AuthBackend("example",
            path="group_okta",
            organization="dummy")
        foo = vault.okta.AuthBackendGroup("foo",
            path=example.path,
            group_name="foo",
            policies=[
                "one",
                "two",
            ])
        ```

        ## Import

        Okta authentication backend groups can be imported using the format `backend/groupName` e.g.

        ```sh
        $ pulumi import vault:okta/authBackendGroup:AuthBackendGroup foo okta/foo
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendGroupInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendGroupInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendGroupInitArgs.__new__(AuthBackendGroupInitArgs)

            if group_name is None and not opts.urn:
                raise TypeError("Missing required property 'group_name'")
            __props__.__dict__["group_name"] = group_name
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["policies"] = policies
        super(AuthBackendGroup, __self__).__init__(
            'vault:okta/authBackendGroup:AuthBackendGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            path: Optional[pulumi.Input[_builtins.str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'AuthBackendGroup':
        """
        Get an existing AuthBackendGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] group_name: Name of the group within the Okta
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] policies: Vault policies to associate with this group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendGroupState.__new__(_AuthBackendGroupState)

        __props__.__dict__["group_name"] = group_name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["policies"] = policies
        return AuthBackendGroup(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the group within the Okta
        """
        return pulumi.get(self, "group_name")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Output[_builtins.str]:
        """
        The path where the Okta auth backend is mounted
        """
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Vault policies to associate with this group
        """
        return pulumi.get(self, "policies")

