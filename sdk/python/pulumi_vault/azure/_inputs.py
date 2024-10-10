# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = [
    'BackendRoleAzureGroupArgs',
    'BackendRoleAzureGroupArgsDict',
    'BackendRoleAzureRoleArgs',
    'BackendRoleAzureRoleArgsDict',
]

MYPY = False

if not MYPY:
    class BackendRoleAzureGroupArgsDict(TypedDict):
        group_name: pulumi.Input[str]
        object_id: NotRequired[pulumi.Input[str]]
elif False:
    BackendRoleAzureGroupArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BackendRoleAzureGroupArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 object_id: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "group_name", group_name)
        if object_id is not None:
            pulumi.set(__self__, "object_id", object_id)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter(name="objectId")
    def object_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "object_id")

    @object_id.setter
    def object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "object_id", value)


if not MYPY:
    class BackendRoleAzureRoleArgsDict(TypedDict):
        scope: pulumi.Input[str]
        role_id: NotRequired[pulumi.Input[str]]
        role_name: NotRequired[pulumi.Input[str]]
elif False:
    BackendRoleAzureRoleArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BackendRoleAzureRoleArgs:
    def __init__(__self__, *,
                 scope: pulumi.Input[str],
                 role_id: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "scope", scope)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Input[str]:
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)


