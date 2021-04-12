# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities, _tables

__all__ = ['NamespaceArgs', 'Namespace']

@pulumi.input_type
class NamespaceArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str]):
        """
        The set of arguments for constructing a Namespace resource.
        :param pulumi.Input[str] path: The path of the namespace. Must not have a trailing `/`
        """
        pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        The path of the namespace. Must not have a trailing `/`
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)


class Namespace(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage [Namespaces](https://www.vaultproject.io/docs/enterprise/namespaces/index.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        ns1 = vault.Namespace("ns1", path="ns1")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] path: The path of the namespace. Must not have a trailing `/`
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NamespaceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage [Namespaces](https://www.vaultproject.io/docs/enterprise/namespaces/index.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        ns1 = vault.Namespace("ns1", path="ns1")
        ```

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
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['namespace_id'] = None
        super(Namespace, __self__).__init__(
            'vault:index/namespace:Namespace',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'Namespace':
        """
        Get an existing Namespace resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] namespace_id: ID of the namepsace.
        :param pulumi.Input[str] path: The path of the namespace. Must not have a trailing `/`
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["namespace_id"] = namespace_id
        __props__["path"] = path
        return Namespace(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        ID of the namepsace.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path of the namespace. Must not have a trailing `/`
        """
        return pulumi.get(self, "path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

