# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['Endpoint']


class Endpoint(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data_json: Optional[pulumi.Input[str]] = None,
                 disable_delete: Optional[pulumi.Input[bool]] = None,
                 disable_read: Optional[pulumi.Input[bool]] = None,
                 ignore_absent_fields: Optional[pulumi.Input[bool]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 write_fields: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Endpoint resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written to the given path as the secret data.
        :param pulumi.Input[bool] disable_delete: Don't attempt to delete the path from Vault if true
        :param pulumi.Input[bool] disable_read: True/false. Set this to true if your vault
               authentication is not able to read the data or if the endpoint does
               not support the `GET` method. Setting this to `true` will break drift
               detection. You should set this to `true` for endpoints that are
               write-only. Defaults to false.
        :param pulumi.Input[bool] ignore_absent_fields: When reading, disregard fields not present in data_json
        :param pulumi.Input[str] path: The full logical path at which to write the given
               data. Consult each backend's documentation to see which endpoints
               support the `PUT` methods and to determine whether they also support
               `DELETE` and `GET`.
        :param pulumi.Input[List[pulumi.Input[str]]] write_fields: Top-level fields returned by write to persist in state
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
            __props__['disable_delete'] = disable_delete
            __props__['disable_read'] = disable_read
            __props__['ignore_absent_fields'] = ignore_absent_fields
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['write_fields'] = write_fields
            __props__['write_data'] = None
            __props__['write_data_json'] = None
        super(Endpoint, __self__).__init__(
            'vault:generic/endpoint:Endpoint',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            data_json: Optional[pulumi.Input[str]] = None,
            disable_delete: Optional[pulumi.Input[bool]] = None,
            disable_read: Optional[pulumi.Input[bool]] = None,
            ignore_absent_fields: Optional[pulumi.Input[bool]] = None,
            path: Optional[pulumi.Input[str]] = None,
            write_data: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            write_data_json: Optional[pulumi.Input[str]] = None,
            write_fields: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'Endpoint':
        """
        Get an existing Endpoint resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] data_json: String containing a JSON-encoded object that will be
               written to the given path as the secret data.
        :param pulumi.Input[bool] disable_delete: Don't attempt to delete the path from Vault if true
        :param pulumi.Input[bool] disable_read: True/false. Set this to true if your vault
               authentication is not able to read the data or if the endpoint does
               not support the `GET` method. Setting this to `true` will break drift
               detection. You should set this to `true` for endpoints that are
               write-only. Defaults to false.
        :param pulumi.Input[bool] ignore_absent_fields: When reading, disregard fields not present in data_json
        :param pulumi.Input[str] path: The full logical path at which to write the given
               data. Consult each backend's documentation to see which endpoints
               support the `PUT` methods and to determine whether they also support
               `DELETE` and `GET`.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] write_data: Map of strings returned by write operation
        :param pulumi.Input[str] write_data_json: JSON data returned by write operation
        :param pulumi.Input[List[pulumi.Input[str]]] write_fields: Top-level fields returned by write to persist in state
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["data_json"] = data_json
        __props__["disable_delete"] = disable_delete
        __props__["disable_read"] = disable_read
        __props__["ignore_absent_fields"] = ignore_absent_fields
        __props__["path"] = path
        __props__["write_data"] = write_data
        __props__["write_data_json"] = write_data_json
        __props__["write_fields"] = write_fields
        return Endpoint(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> pulumi.Output[str]:
        """
        String containing a JSON-encoded object that will be
        written to the given path as the secret data.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter(name="disableDelete")
    def disable_delete(self) -> pulumi.Output[Optional[bool]]:
        """
        Don't attempt to delete the path from Vault if true
        """
        return pulumi.get(self, "disable_delete")

    @property
    @pulumi.getter(name="disableRead")
    def disable_read(self) -> pulumi.Output[Optional[bool]]:
        """
        True/false. Set this to true if your vault
        authentication is not able to read the data or if the endpoint does
        not support the `GET` method. Setting this to `true` will break drift
        detection. You should set this to `true` for endpoints that are
        write-only. Defaults to false.
        """
        return pulumi.get(self, "disable_read")

    @property
    @pulumi.getter(name="ignoreAbsentFields")
    def ignore_absent_fields(self) -> pulumi.Output[Optional[bool]]:
        """
        When reading, disregard fields not present in data_json
        """
        return pulumi.get(self, "ignore_absent_fields")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The full logical path at which to write the given
        data. Consult each backend's documentation to see which endpoints
        support the `PUT` methods and to determine whether they also support
        `DELETE` and `GET`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="writeData")
    def write_data(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Map of strings returned by write operation
        """
        return pulumi.get(self, "write_data")

    @property
    @pulumi.getter(name="writeDataJson")
    def write_data_json(self) -> pulumi.Output[str]:
        """
        JSON data returned by write operation
        """
        return pulumi.get(self, "write_data_json")

    @property
    @pulumi.getter(name="writeFields")
    def write_fields(self) -> pulumi.Output[Optional[List[str]]]:
        """
        Top-level fields returned by write to persist in state
        """
        return pulumi.get(self, "write_fields")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

