# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'AuthBackendGroupArgs',
    'AuthBackendUserArgs',
]

@pulumi.input_type
class AuthBackendGroupArgs:
    def __init__(__self__, *,
                 group_name: pulumi.Input[str],
                 policies: pulumi.Input[Sequence[pulumi.Input[str]]]):
        """
        :param pulumi.Input[str] group_name: Name of the Okta group
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Policies to associate with this group
        """
        pulumi.set(__self__, "group_name", group_name)
        pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Input[str]:
        """
        Name of the Okta group
        """
        return pulumi.get(self, "group_name")

    @group_name.setter
    def group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "group_name", value)

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Policies to associate with this group
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class AuthBackendUserArgs:
    def __init__(__self__, *,
                 username: pulumi.Input[str],
                 groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        :param pulumi.Input[str] username: Name of the user within Okta
        :param pulumi.Input[Sequence[pulumi.Input[str]]] groups: Groups within the Okta auth backend to associate with this user
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Policies to associate with this user
        """
        pulumi.set(__self__, "username", username)
        if groups is not None:
            pulumi.set(__self__, "groups", groups)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        Name of the user within Okta
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Groups within the Okta auth backend to associate with this user
        """
        return pulumi.get(self, "groups")

    @groups.setter
    def groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "groups", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Policies to associate with this user
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


