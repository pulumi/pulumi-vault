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
from . import _utilities

__all__ = ['NamespaceArgs', 'Namespace']

@pulumi.input_type
class NamespaceArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[_builtins.str],
                 custom_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path_fq: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a Namespace resource.
        :param pulumi.Input[_builtins.str] path: The path of the namespace. Must not have a trailing `/`.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] custom_metadata: Custom metadata describing this namespace. Value type
               is `map[string]string`. Requires Vault version 1.12+.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path_fq: The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
               The path is relative to the provider's `namespace` argument.
        """
        pulumi.set(__self__, "path", path)
        if custom_metadata is not None:
            pulumi.set(__self__, "custom_metadata", custom_metadata)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path_fq is not None:
            pulumi.set(__self__, "path_fq", path_fq)

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Input[_builtins.str]:
        """
        The path of the namespace. Must not have a trailing `/`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter(name="customMetadata")
    def custom_metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Custom metadata describing this namespace. Value type
        is `map[string]string`. Requires Vault version 1.12+.
        """
        return pulumi.get(self, "custom_metadata")

    @custom_metadata.setter
    def custom_metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "custom_metadata", value)

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
    @pulumi.getter(name="pathFq")
    def path_fq(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        The path is relative to the provider's `namespace` argument.
        """
        return pulumi.get(self, "path_fq")

    @path_fq.setter
    def path_fq(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path_fq", value)


@pulumi.input_type
class _NamespaceState:
    def __init__(__self__, *,
                 custom_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 path_fq: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering Namespace resources.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] custom_metadata: Custom metadata describing this namespace. Value type
               is `map[string]string`. Requires Vault version 1.12+.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] namespace_id: Vault server's internal ID of the namespace.
        :param pulumi.Input[_builtins.str] path: The path of the namespace. Must not have a trailing `/`.
        :param pulumi.Input[_builtins.str] path_fq: The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
               The path is relative to the provider's `namespace` argument.
        """
        if custom_metadata is not None:
            pulumi.set(__self__, "custom_metadata", custom_metadata)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if path_fq is not None:
            pulumi.set(__self__, "path_fq", path_fq)

    @_builtins.property
    @pulumi.getter(name="customMetadata")
    def custom_metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]:
        """
        Custom metadata describing this namespace. Value type
        is `map[string]string`. Requires Vault version 1.12+.
        """
        return pulumi.get(self, "custom_metadata")

    @custom_metadata.setter
    def custom_metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "custom_metadata", value)

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
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Vault server's internal ID of the namespace.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_id", value)

    @_builtins.property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The path of the namespace. Must not have a trailing `/`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter(name="pathFq")
    def path_fq(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        The path is relative to the provider's `namespace` argument.
        """
        return pulumi.get(self, "path_fq")

    @path_fq.setter
    def path_fq(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path_fq", value)


@pulumi.type_token("vault:index/namespace:Namespace")
class Namespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 path_fq: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## Import

        Namespaces can be imported using its `name` as accessor id

        ```sh
        $ pulumi import vault:index/namespace:Namespace example <name>
        ```

        If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path.

        hcl

        provider "vault" {

        # Configuration options

          namespace = "example"

          alias     = "example"

        }

        resource "vault_namespace" "example2" {

          provider = vault.example

          path     = "example2"

        }

        ```sh
        $ pulumi import vault:index/namespace:Namespace example2 example2
        ```

        $ terraform state show vault_namespace.example2

        vault_namespace.example2:

        resource "vault_namespace" "example2" {

            id           = "example/example2/"
            
            namespace_id = <known after import>
            
            path         = "example2"
            
            path_fq      = "example2"

        }

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] custom_metadata: Custom metadata describing this namespace. Value type
               is `map[string]string`. Requires Vault version 1.12+.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] path: The path of the namespace. Must not have a trailing `/`.
        :param pulumi.Input[_builtins.str] path_fq: The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
               The path is relative to the provider's `namespace` argument.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Import

        Namespaces can be imported using its `name` as accessor id

        ```sh
        $ pulumi import vault:index/namespace:Namespace example <name>
        ```

        If the declared resource is imported and intends to support namespaces using a provider alias, then the name is relative to the namespace path.

        hcl

        provider "vault" {

        # Configuration options

          namespace = "example"

          alias     = "example"

        }

        resource "vault_namespace" "example2" {

          provider = vault.example

          path     = "example2"

        }

        ```sh
        $ pulumi import vault:index/namespace:Namespace example2 example2
        ```

        $ terraform state show vault_namespace.example2

        vault_namespace.example2:

        resource "vault_namespace" "example2" {

            id           = "example/example2/"
            
            namespace_id = <known after import>
            
            path         = "example2"
            
            path_fq      = "example2"

        }

        :param str resource_name: The name of the resource.
        :param NamespaceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NamespaceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 custom_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 path_fq: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NamespaceArgs.__new__(NamespaceArgs)

            __props__.__dict__["custom_metadata"] = custom_metadata
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["path_fq"] = path_fq
            __props__.__dict__["namespace_id"] = None
        super(Namespace, __self__).__init__(
            'vault:index/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            custom_metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
            path: Optional[pulumi.Input[_builtins.str]] = None,
            path_fq: Optional[pulumi.Input[_builtins.str]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, pulumi.Input[_builtins.str]]] custom_metadata: Custom metadata describing this namespace. Value type
               is `map[string]string`. Requires Vault version 1.12+.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] namespace_id: Vault server's internal ID of the namespace.
        :param pulumi.Input[_builtins.str] path: The path of the namespace. Must not have a trailing `/`.
        :param pulumi.Input[_builtins.str] path_fq: The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
               The path is relative to the provider's `namespace` argument.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NamespaceState.__new__(_NamespaceState)

        __props__.__dict__["custom_metadata"] = custom_metadata
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["path"] = path
        __props__.__dict__["path_fq"] = path_fq
        return Namespace(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="customMetadata")
    def custom_metadata(self) -> pulumi.Output[Mapping[str, _builtins.str]]:
        """
        Custom metadata describing this namespace. Value type
        is `map[string]string`. Requires Vault version 1.12+.
        """
        return pulumi.get(self, "custom_metadata")

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
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[_builtins.str]:
        """
        Vault server's internal ID of the namespace.
        """
        return pulumi.get(self, "namespace_id")

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Output[_builtins.str]:
        """
        The path of the namespace. Must not have a trailing `/`.
        """
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter(name="pathFq")
    def path_fq(self) -> pulumi.Output[_builtins.str]:
        """
        The fully qualified path to the namespace. Useful when provisioning resources in a child `namespace`.
        The path is relative to the provider's `namespace` argument.
        """
        return pulumi.get(self, "path_fq")

