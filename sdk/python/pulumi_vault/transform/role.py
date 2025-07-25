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

__all__ = ['RoleArgs', 'Role']

@pulumi.input_type
class RoleArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[_builtins.str],
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 transformations: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a Role resource.
        :param pulumi.Input[_builtins.str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[_builtins.str] name: The name of the role.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] transformations: A comma separated string or slice of transformations to use.
        """
        pulumi.set(__self__, "path", path)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if transformations is not None:
            pulumi.set(__self__, "transformations", transformations)

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Input[_builtins.str]:
        """
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

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
    def transformations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A comma separated string or slice of transformations to use.
        """
        return pulumi.get(self, "transformations")

    @transformations.setter
    def transformations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "transformations", value)


@pulumi.input_type
class _RoleState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 transformations: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering Role resources.
        :param pulumi.Input[_builtins.str] name: The name of the role.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] transformations: A comma separated string or slice of transformations to use.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if transformations is not None:
            pulumi.set(__self__, "transformations", transformations)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

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
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter
    def transformations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        A comma separated string or slice of transformations to use.
        """
        return pulumi.get(self, "transformations")

    @transformations.setter
    def transformations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "transformations", value)


@pulumi.type_token("vault:transform/role:Role")
class Role(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 transformations: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        This resource supports the "/transform/role/{name}" Vault endpoint.

        It creates or updates the role with the given name. If a role with the name does not exist, it will be created.
        If the role exists, it will be updated with the new attributes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mount_transform = vault.Mount("mount_transform",
            path="transform",
            type="transform")
        test = vault.transform.Role("test",
            path=mount_transform.path,
            name="payments",
            transformations=["ccn-fpe"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] name: The name of the role.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] transformations: A comma separated string or slice of transformations to use.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports the "/transform/role/{name}" Vault endpoint.

        It creates or updates the role with the given name. If a role with the name does not exist, it will be created.
        If the role exists, it will be updated with the new attributes.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mount_transform = vault.Mount("mount_transform",
            path="transform",
            type="transform")
        test = vault.transform.Role("test",
            path=mount_transform.path,
            name="payments",
            transformations=["ccn-fpe"])
        ```

        :param str resource_name: The name of the resource.
        :param RoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 transformations: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RoleArgs.__new__(RoleArgs)

            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["transformations"] = transformations
        super(Role, __self__).__init__(
            'vault:transform/role:Role',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            path: Optional[pulumi.Input[_builtins.str]] = None,
            transformations: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'Role':
        """
        Get an existing Role resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] name: The name of the role.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] transformations: A comma separated string or slice of transformations to use.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _RoleState.__new__(_RoleState)

        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["transformations"] = transformations
        return Role(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

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
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter
    def transformations(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        A comma separated string or slice of transformations to use.
        """
        return pulumi.get(self, "transformations")

