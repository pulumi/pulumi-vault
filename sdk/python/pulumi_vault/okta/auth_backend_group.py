# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendGroup']


class AuthBackendGroup(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 group_name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create a group in an
        [Okta auth backend within Vault](https://www.vaultproject.io/docs/auth/okta.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.okta.AuthBackend("example",
            organization="dummy",
            path="group_okta")
        foo = vault.okta.AuthBackendGroup("foo",
            group_name="foo",
            path=example.path,
            policies=[
                "one",
                "two",
            ])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the group within the Okta
        :param pulumi.Input[str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Vault policies to associate with this group
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

            if group_name is None:
                raise TypeError("Missing required property 'group_name'")
            __props__['group_name'] = group_name
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['policies'] = policies
        super(AuthBackendGroup, __self__).__init__(
            'vault:okta/authBackendGroup:AuthBackendGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            group_name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'AuthBackendGroup':
        """
        Get an existing AuthBackendGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] group_name: Name of the group within the Okta
        :param pulumi.Input[str] path: The path where the Okta auth backend is mounted
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: Vault policies to associate with this group
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["group_name"] = group_name
        __props__["path"] = path
        __props__["policies"] = policies
        return AuthBackendGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="groupName")
    def group_name(self) -> pulumi.Output[str]:
        """
        Name of the group within the Okta
        """
        return pulumi.get(self, "group_name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path where the Okta auth backend is mounted
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Vault policies to associate with this group
        """
        return pulumi.get(self, "policies")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

