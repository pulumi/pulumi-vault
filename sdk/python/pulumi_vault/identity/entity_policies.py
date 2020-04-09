# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class EntityPolicies(pulumi.CustomResource):
    entity_id: pulumi.Output[str]
    """
    Entity ID to assign policies to.
    """
    entity_name: pulumi.Output[str]
    """
    The name of the entity that are assigned the policies.
    """
    exclusive: pulumi.Output[bool]
    """
    Defaults to `true`.
    """
    policies: pulumi.Output[list]
    """
    List of policies to assign to the entity
    """
    def __init__(__self__, resource_name, opts=None, entity_id=None, exclusive=None, policies=None, __props__=None, __name__=None, __opts__=None):
        """
        Manages policies for an Identity Entity for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.



        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/identity_entity_policies.html.md.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: Entity ID to assign policies to.
        :param pulumi.Input[bool] exclusive: Defaults to `true`.
        :param pulumi.Input[list] policies: List of policies to assign to the entity
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if entity_id is None:
                raise TypeError("Missing required property 'entity_id'")
            __props__['entity_id'] = entity_id
            __props__['exclusive'] = exclusive
            if policies is None:
                raise TypeError("Missing required property 'policies'")
            __props__['policies'] = policies
            __props__['entity_name'] = None
        super(EntityPolicies, __self__).__init__(
            'vault:identity/entityPolicies:EntityPolicies',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, entity_id=None, entity_name=None, exclusive=None, policies=None):
        """
        Get an existing EntityPolicies resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] entity_id: Entity ID to assign policies to.
        :param pulumi.Input[str] entity_name: The name of the entity that are assigned the policies.
        :param pulumi.Input[bool] exclusive: Defaults to `true`.
        :param pulumi.Input[list] policies: List of policies to assign to the entity
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["entity_id"] = entity_id
        __props__["entity_name"] = entity_name
        __props__["exclusive"] = exclusive
        __props__["policies"] = policies
        return EntityPolicies(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

