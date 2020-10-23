# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Secret']


class Secret(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_json: Optional[pulumi.Input[str]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Secret resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[bool] disable_read: True/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        """
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

            if data_json is None:
                raise TypeError("Missing required property 'data_json'")
            __props__['data_json'] = data_json
            __props__['disable_read'] = disable_read
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['data'] = None
        super(Secret, __self__).__init__(
            'vault:generic/secret:Secret',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            data_json: Optional[pulumi.Input[str]] = None,
            disable_read: Optional[pulumi.Input[bool]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'Secret':
        """
        Get an existing Secret resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Mapping[str, Any]] data: A mapping whose keys are the top-level data keys returned from
               Vault and whose values are the corresponding values. This map can only
               represent string data, so any non-string values returned from Vault are
               serialized as JSON.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written as the secret data at the given path.
        :param pulumi.Input[bool] disable_read: True/false. Set this to true if your vault
               authentication is not able to read the data. Setting this to `true` will
               break drift detection. Defaults to false.
        :param pulumi.Input[str] path: The full logical path at which to write the given data.
               To write data into the "generic" secret backend mounted in Vault by default,
               this should be prefixed with `secret/`. Writing to other backends with this
               resource is possible; consult each backend's documentation to see which
               endpoints support the `PUT` and `DELETE` methods.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data"] = data
        __props__["data_json"] = data_json
        __props__["disable_read"] = disable_read
        __props__["path"] = path
        return Secret(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[Mapping[str, Any]]:
        """
        A mapping whose keys are the top-level data keys returned from
        Vault and whose values are the corresponding values. This map can only
        represent string data, so any non-string values returned from Vault are
        serialized as JSON.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> pulumi.Output[str]:
        """
        String containing a JSON-encoded object that will be
        written as the secret data at the given path.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter(name="disableRead")
    def disable_read(self) -> pulumi.Output[Optional[bool]]:
        """
        True/false. Set this to true if your vault
        authentication is not able to read the data. Setting this to `true` will
        break drift detection. Defaults to false.
        """
        return pulumi.get(self, "disable_read")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The full logical path at which to write the given data.
        To write data into the "generic" secret backend mounted in Vault by default,
        this should be prefixed with `secret/`. Writing to other backends with this
        resource is possible; consult each backend's documentation to see which
        endpoints support the `PUT` and `DELETE` methods.
        """
        return pulumi.get(self, "path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

