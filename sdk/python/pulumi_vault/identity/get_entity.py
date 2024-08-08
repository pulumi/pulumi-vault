# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs

__all__ = [
    'GetEntityResult',
    'AwaitableGetEntityResult',
    'get_entity',
    'get_entity_output',
]

@pulumi.output_type
class GetEntityResult:
    """
    A collection of values returned by getEntity.
    """
    def __init__(__self__, alias_id=None, alias_mount_accessor=None, alias_name=None, aliases=None, creation_time=None, data_json=None, direct_group_ids=None, disabled=None, entity_id=None, entity_name=None, group_ids=None, id=None, inherited_group_ids=None, last_update_time=None, merged_entity_ids=None, metadata=None, namespace=None, namespace_id=None, policies=None):
        if alias_id and not isinstance(alias_id, str):
            raise TypeError("Expected argument 'alias_id' to be a str")
        pulumi.set(__self__, "alias_id", alias_id)
        if alias_mount_accessor and not isinstance(alias_mount_accessor, str):
            raise TypeError("Expected argument 'alias_mount_accessor' to be a str")
        pulumi.set(__self__, "alias_mount_accessor", alias_mount_accessor)
        if alias_name and not isinstance(alias_name, str):
            raise TypeError("Expected argument 'alias_name' to be a str")
        pulumi.set(__self__, "alias_name", alias_name)
        if aliases and not isinstance(aliases, list):
            raise TypeError("Expected argument 'aliases' to be a list")
        pulumi.set(__self__, "aliases", aliases)
        if creation_time and not isinstance(creation_time, str):
            raise TypeError("Expected argument 'creation_time' to be a str")
        pulumi.set(__self__, "creation_time", creation_time)
        if data_json and not isinstance(data_json, str):
            raise TypeError("Expected argument 'data_json' to be a str")
        pulumi.set(__self__, "data_json", data_json)
        if direct_group_ids and not isinstance(direct_group_ids, list):
            raise TypeError("Expected argument 'direct_group_ids' to be a list")
        pulumi.set(__self__, "direct_group_ids", direct_group_ids)
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if entity_id and not isinstance(entity_id, str):
            raise TypeError("Expected argument 'entity_id' to be a str")
        pulumi.set(__self__, "entity_id", entity_id)
        if entity_name and not isinstance(entity_name, str):
            raise TypeError("Expected argument 'entity_name' to be a str")
        pulumi.set(__self__, "entity_name", entity_name)
        if group_ids and not isinstance(group_ids, list):
            raise TypeError("Expected argument 'group_ids' to be a list")
        pulumi.set(__self__, "group_ids", group_ids)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if inherited_group_ids and not isinstance(inherited_group_ids, list):
            raise TypeError("Expected argument 'inherited_group_ids' to be a list")
        pulumi.set(__self__, "inherited_group_ids", inherited_group_ids)
        if last_update_time and not isinstance(last_update_time, str):
            raise TypeError("Expected argument 'last_update_time' to be a str")
        pulumi.set(__self__, "last_update_time", last_update_time)
        if merged_entity_ids and not isinstance(merged_entity_ids, list):
            raise TypeError("Expected argument 'merged_entity_ids' to be a list")
        pulumi.set(__self__, "merged_entity_ids", merged_entity_ids)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if namespace_id and not isinstance(namespace_id, str):
            raise TypeError("Expected argument 'namespace_id' to be a str")
        pulumi.set(__self__, "namespace_id", namespace_id)
        if policies and not isinstance(policies, list):
            raise TypeError("Expected argument 'policies' to be a list")
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="aliasId")
    def alias_id(self) -> str:
        return pulumi.get(self, "alias_id")

    @property
    @pulumi.getter(name="aliasMountAccessor")
    def alias_mount_accessor(self) -> str:
        return pulumi.get(self, "alias_mount_accessor")

    @property
    @pulumi.getter(name="aliasName")
    def alias_name(self) -> str:
        return pulumi.get(self, "alias_name")

    @property
    @pulumi.getter
    def aliases(self) -> Sequence['outputs.GetEntityAliasResult']:
        """
        A list of entity alias. Structure is documented below.
        """
        return pulumi.get(self, "aliases")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> str:
        """
        Creation time of the Alias
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter(name="dataJson")
    def data_json(self) -> str:
        """
        A string containing the full data payload retrieved from
        Vault, serialized in JSON format.
        """
        return pulumi.get(self, "data_json")

    @property
    @pulumi.getter(name="directGroupIds")
    def direct_group_ids(self) -> Sequence[str]:
        """
        List of Group IDs of which the entity is directly a member of
        """
        return pulumi.get(self, "direct_group_ids")

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        Whether the entity is disabled
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> str:
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="entityName")
    def entity_name(self) -> str:
        return pulumi.get(self, "entity_name")

    @property
    @pulumi.getter(name="groupIds")
    def group_ids(self) -> Sequence[str]:
        """
        List of all Group IDs of which the entity is a member of
        """
        return pulumi.get(self, "group_ids")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="inheritedGroupIds")
    def inherited_group_ids(self) -> Sequence[str]:
        """
        List of all Group IDs of which the entity is a member of transitively
        """
        return pulumi.get(self, "inherited_group_ids")

    @property
    @pulumi.getter(name="lastUpdateTime")
    def last_update_time(self) -> str:
        """
        Last update time of the alias
        """
        return pulumi.get(self, "last_update_time")

    @property
    @pulumi.getter(name="mergedEntityIds")
    def merged_entity_ids(self) -> Sequence[str]:
        """
        Other entity IDs which is merged with this entity
        """
        return pulumi.get(self, "merged_entity_ids")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, Any]:
        """
        Arbitrary metadata
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> str:
        """
        Namespace of which the entity is part of
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter
    def policies(self) -> Sequence[str]:
        """
        List of policies attached to the entity
        """
        return pulumi.get(self, "policies")


class AwaitableGetEntityResult(GetEntityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntityResult(
            alias_id=self.alias_id,
            alias_mount_accessor=self.alias_mount_accessor,
            alias_name=self.alias_name,
            aliases=self.aliases,
            creation_time=self.creation_time,
            data_json=self.data_json,
            direct_group_ids=self.direct_group_ids,
            disabled=self.disabled,
            entity_id=self.entity_id,
            entity_name=self.entity_name,
            group_ids=self.group_ids,
            id=self.id,
            inherited_group_ids=self.inherited_group_ids,
            last_update_time=self.last_update_time,
            merged_entity_ids=self.merged_entity_ids,
            metadata=self.metadata,
            namespace=self.namespace,
            namespace_id=self.namespace_id,
            policies=self.policies)


def get_entity(alias_id: Optional[str] = None,
               alias_mount_accessor: Optional[str] = None,
               alias_name: Optional[str] = None,
               entity_id: Optional[str] = None,
               entity_name: Optional[str] = None,
               namespace: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntityResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    entity = vault.identity.get_entity(entity_name="entity_12345")
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `update` capability on `/identity/lookup/entity`.


    :param str alias_id: ID of the alias.
    :param str alias_mount_accessor: Accessor of the mount to which the alias belongs to.
           This should be supplied in conjunction with `alias_name`.
           
           The lookup criteria can be `entity_name`, `entity_id`, `alias_id`, or a combination of
           `alias_name` and `alias_mount_accessor`.
    :param str alias_name: Name of the alias. This should be supplied in conjunction with
           `alias_mount_accessor`.
    :param str entity_id: ID of the entity.
    :param str entity_name: Name of the entity.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['aliasId'] = alias_id
    __args__['aliasMountAccessor'] = alias_mount_accessor
    __args__['aliasName'] = alias_name
    __args__['entityId'] = entity_id
    __args__['entityName'] = entity_name
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:identity/getEntity:getEntity', __args__, opts=opts, typ=GetEntityResult).value

    return AwaitableGetEntityResult(
        alias_id=pulumi.get(__ret__, 'alias_id'),
        alias_mount_accessor=pulumi.get(__ret__, 'alias_mount_accessor'),
        alias_name=pulumi.get(__ret__, 'alias_name'),
        aliases=pulumi.get(__ret__, 'aliases'),
        creation_time=pulumi.get(__ret__, 'creation_time'),
        data_json=pulumi.get(__ret__, 'data_json'),
        direct_group_ids=pulumi.get(__ret__, 'direct_group_ids'),
        disabled=pulumi.get(__ret__, 'disabled'),
        entity_id=pulumi.get(__ret__, 'entity_id'),
        entity_name=pulumi.get(__ret__, 'entity_name'),
        group_ids=pulumi.get(__ret__, 'group_ids'),
        id=pulumi.get(__ret__, 'id'),
        inherited_group_ids=pulumi.get(__ret__, 'inherited_group_ids'),
        last_update_time=pulumi.get(__ret__, 'last_update_time'),
        merged_entity_ids=pulumi.get(__ret__, 'merged_entity_ids'),
        metadata=pulumi.get(__ret__, 'metadata'),
        namespace=pulumi.get(__ret__, 'namespace'),
        namespace_id=pulumi.get(__ret__, 'namespace_id'),
        policies=pulumi.get(__ret__, 'policies'))


@_utilities.lift_output_func(get_entity)
def get_entity_output(alias_id: Optional[pulumi.Input[Optional[str]]] = None,
                      alias_mount_accessor: Optional[pulumi.Input[Optional[str]]] = None,
                      alias_name: Optional[pulumi.Input[Optional[str]]] = None,
                      entity_id: Optional[pulumi.Input[Optional[str]]] = None,
                      entity_name: Optional[pulumi.Input[Optional[str]]] = None,
                      namespace: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEntityResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    entity = vault.identity.get_entity(entity_name="entity_12345")
    ```

    ## Required Vault Capabilities

    Use of this resource requires the `update` capability on `/identity/lookup/entity`.


    :param str alias_id: ID of the alias.
    :param str alias_mount_accessor: Accessor of the mount to which the alias belongs to.
           This should be supplied in conjunction with `alias_name`.
           
           The lookup criteria can be `entity_name`, `entity_id`, `alias_id`, or a combination of
           `alias_name` and `alias_mount_accessor`.
    :param str alias_name: Name of the alias. This should be supplied in conjunction with
           `alias_mount_accessor`.
    :param str entity_id: ID of the entity.
    :param str entity_name: Name of the entity.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    ...
