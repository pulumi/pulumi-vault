# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['GroupAlias']


class GroupAlias(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 canonical_id: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates an Identity Group Alias for Vault. The [Identity secrets engine](https://www.vaultproject.io/docs/secrets/identity/index.html) is the identity management solution for Vault.

        Group aliases allows entity membership in external groups to be managed semi-automatically. External group serves as a mapping to a group that is outside of the identity store. External groups can have one (and only one) alias. This alias should map to a notion of group that is outside of the identity store. For example, groups in LDAP, and teams in GitHub. A username in LDAP, belonging to a group in LDAP, can get its entity ID added as a member of a group in Vault automatically during logins and token renewals. This works only if the group in Vault is an external group and has an alias that maps to the group in LDAP. If the user is removed from the group in LDAP, that change gets reflected in Vault only upon the subsequent login or renewal operation.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        group = vault.identity.Group("group",
            policies=["test"],
            type="external")
        github = vault.AuthBackend("github",
            path="github",
            type="github")
        group_alias = vault.identity.GroupAlias("group-alias",
            canonical_id=group.id,
            mount_accessor=github.accessor,
            name="Github_Team_Slug")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] canonical_id: ID of the group to which this is an alias.
        :param pulumi.Input[str] mount_accessor: Mount accessor of the authentication backend to which this alias belongs to.
        :param pulumi.Input[str] name: Name of the group alias to create.
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

            if canonical_id is None:
                raise TypeError("Missing required property 'canonical_id'")
            __props__['canonical_id'] = canonical_id
            if mount_accessor is None:
                raise TypeError("Missing required property 'mount_accessor'")
            __props__['mount_accessor'] = mount_accessor
            if name is None:
                raise TypeError("Missing required property 'name'")
            __props__['name'] = name
        super(GroupAlias, __self__).__init__(
            'vault:identity/groupAlias:GroupAlias',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            canonical_id: Optional[pulumi.Input[str]] = None,
            mount_accessor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'GroupAlias':
        """
        Get an existing GroupAlias resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] canonical_id: ID of the group to which this is an alias.
        :param pulumi.Input[str] mount_accessor: Mount accessor of the authentication backend to which this alias belongs to.
        :param pulumi.Input[str] name: Name of the group alias to create.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["canonical_id"] = canonical_id
        __props__["mount_accessor"] = mount_accessor
        __props__["name"] = name
        return GroupAlias(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="canonicalId")
    def canonical_id(self) -> pulumi.Output[str]:
        """
        ID of the group to which this is an alias.
        """
        return pulumi.get(self, "canonical_id")

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[str]:
        """
        Mount accessor of the authentication backend to which this alias belongs to.
        """
        return pulumi.get(self, "mount_accessor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the group alias to create.
        """
        return pulumi.get(self, "name")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

