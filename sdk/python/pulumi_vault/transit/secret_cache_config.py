# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SecretCacheConfigArgs', 'SecretCacheConfig']

@pulumi.input_type
class SecretCacheConfigArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 size: pulumi.Input[int]):
        """
        The set of arguments for constructing a SecretCacheConfig resource.
        :param pulumi.Input[str] backend: The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] size: The number of cache entries. 0 means unlimited.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "size", size)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def size(self) -> pulumi.Input[int]:
        """
        The number of cache entries. 0 means unlimited.
        """
        return pulumi.get(self, "size")

    @size.setter
    def size(self, value: pulumi.Input[int]):
        pulumi.set(self, "size", value)


class SecretCacheConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Configure the cache for the Transit Secret Backend in Vault.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        transit = vault.Mount("transit",
            default_lease_ttl_seconds=3600,
            description="Example description",
            max_lease_ttl_seconds=86400,
            path="transit",
            type="transit")
        cfg = vault.transit.SecretCacheConfig("cfg",
            backend=transit.path,
            size=500)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] size: The number of cache entries. 0 means unlimited.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretCacheConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the cache for the Transit Secret Backend in Vault.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        transit = vault.Mount("transit",
            default_lease_ttl_seconds=3600,
            description="Example description",
            max_lease_ttl_seconds=86400,
            path="transit",
            type="transit")
        cfg = vault.transit.SecretCacheConfig("cfg",
            backend=transit.path,
            size=500)
        ```

        :param str resource_name: The name of the resource.
        :param SecretCacheConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretCacheConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 size: Optional[pulumi.Input[int]] = None,
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

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if size is None and not opts.urn:
                raise TypeError("Missing required property 'size'")
            __props__['size'] = size
        super(SecretCacheConfig, __self__).__init__(
            'vault:transit/secretCacheConfig:SecretCacheConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            size: Optional[pulumi.Input[int]] = None) -> 'SecretCacheConfig':
        """
        Get an existing SecretCacheConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] size: The number of cache entries. 0 means unlimited.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["size"] = size
        return SecretCacheConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the transit secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def size(self) -> pulumi.Output[int]:
        """
        The number of cache entries. 0 means unlimited.
        """
        return pulumi.get(self, "size")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

