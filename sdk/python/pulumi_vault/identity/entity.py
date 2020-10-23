# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Entity']


class Entity(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disabled: Optional[pulumi.Input[bool]] = None,
                 external_policies: Optional[pulumi.Input[bool]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Entity resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: True/false Is this entity currently disabled. Defaults to `false`
        :param pulumi.Input[bool] external_policies: `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A Map of additional metadata to associate with the user.
        :param pulumi.Input[str] name: Name of the identity entity to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of policies to apply to the entity.
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

            __props__['disabled'] = disabled
            __props__['external_policies'] = external_policies
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['policies'] = policies
        super(Entity, __self__).__init__(
            'vault:identity/entity:Entity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disabled: Optional[pulumi.Input[bool]] = None,
            external_policies: Optional[pulumi.Input[bool]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'Entity':
        """
        Get an existing Entity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disabled: True/false Is this entity currently disabled. Defaults to `false`
        :param pulumi.Input[bool] external_policies: `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A Map of additional metadata to associate with the user.
        :param pulumi.Input[str] name: Name of the identity entity to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of policies to apply to the entity.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["disabled"] = disabled
        __props__["external_policies"] = external_policies
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["policies"] = policies
        return Entity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def disabled(self) -> pulumi.Output[Optional[bool]]:
        """
        True/false Is this entity currently disabled. Defaults to `false`
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="externalPolicies")
    def external_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        `false` by default. If set to `true`, this resource will ignore any policies return from Vault or specified in the resource. You can use `identity.EntityPolicies` to manage policies for this entity in a decoupled manner.
        """
        return pulumi.get(self, "external_policies")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A Map of additional metadata to associate with the user.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the identity entity to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of policies to apply to the entity.
        """
        return pulumi.get(self, "policies")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

