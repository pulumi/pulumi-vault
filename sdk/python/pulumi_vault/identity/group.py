# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Group']


class Group(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_member_entity_ids: Optional[pulumi.Input[bool]] = None,
                 external_policies: Optional[pulumi.Input[bool]] = None,
                 member_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 member_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Identity Group for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.

        A group can contain multiple entities as its members. A group can also have subgroups. Policies set on the group is granted to all members of the group. During request time, when the token's entity ID is being evaluated for the policies that it has access to; along with the policies on the entity itself, policies that are inherited due to group memberships are also granted.

        ## Example Usage
        ### Internal Group

        ```python
        import pulumi
        import pulumi_vault as vault

        internal = vault.identity.Group("internal",
            metadata={
                "version": "2",
            },
            policies=[
                "dev",
                "test",
            ],
            type="internal")
        ```
        ### External Group

        ```python
        import pulumi
        import pulumi_vault as vault

        group = vault.identity.Group("group",
            metadata={
                "version": "1",
            },
            policies=["test"],
            type="external")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] external_member_entity_ids: Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
        :param pulumi.Input[bool] external_policies: Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] member_entity_ids: A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] member_group_ids: A list of Group IDs to be assigned as group members.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A Map of additional metadata to associate with the group.
        :param pulumi.Input[str] name: Name of the identity group to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of policies to apply to the group.
        :param pulumi.Input[str] type: Type of the group, internal or external. Defaults to `internal`.
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

            __props__['external_member_entity_ids'] = external_member_entity_ids
            __props__['external_policies'] = external_policies
            __props__['member_entity_ids'] = member_entity_ids
            __props__['member_group_ids'] = member_group_ids
            __props__['metadata'] = metadata
            __props__['name'] = name
            __props__['policies'] = policies
            __props__['type'] = type
        super(Group, __self__).__init__(
            'vault:identity/group:Group',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            external_member_entity_ids: Optional[pulumi.Input[bool]] = None,
            external_policies: Optional[pulumi.Input[bool]] = None,
            member_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            member_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Group':
        """
        Get an existing Group resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] external_member_entity_ids: Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
        :param pulumi.Input[bool] external_policies: Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] member_entity_ids: A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] member_group_ids: A list of Group IDs to be assigned as group members.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: A Map of additional metadata to associate with the group.
        :param pulumi.Input[str] name: Name of the identity group to create.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: A list of policies to apply to the group.
        :param pulumi.Input[str] type: Type of the group, internal or external. Defaults to `internal`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["external_member_entity_ids"] = external_member_entity_ids
        __props__["external_policies"] = external_policies
        __props__["member_entity_ids"] = member_entity_ids
        __props__["member_group_ids"] = member_group_ids
        __props__["metadata"] = metadata
        __props__["name"] = name
        __props__["policies"] = policies
        __props__["type"] = type
        return Group(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalMemberEntityIds")
    def external_member_entity_ids(self) -> pulumi.Output[Optional[bool]]:
        """
        Manage member entities externally through `vault_identity_group_policies_member_entity_ids`
        """
        return pulumi.get(self, "external_member_entity_ids")

    @property
    @pulumi.getter(name="externalPolicies")
    def external_policies(self) -> pulumi.Output[Optional[bool]]:
        """
        Manage policies externally through `vault_identity_group_policies`, allows using group ID in assigned policies.
        """
        return pulumi.get(self, "external_policies")

    @property
    @pulumi.getter(name="memberEntityIds")
    def member_entity_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of Entity IDs to be assigned as group members. Not allowed on `external` groups.
        """
        return pulumi.get(self, "member_entity_ids")

    @property
    @pulumi.getter(name="memberGroupIds")
    def member_group_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of Group IDs to be assigned as group members.
        """
        return pulumi.get(self, "member_group_ids")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A Map of additional metadata to associate with the group.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the identity group to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of policies to apply to the group.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        Type of the group, internal or external. Defaults to `internal`.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

