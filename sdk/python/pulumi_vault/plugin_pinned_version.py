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
from . import _utilities

__all__ = ['PluginPinnedVersionArgs', 'PluginPinnedVersion']

@pulumi.input_type
class PluginPinnedVersionArgs:
    def __init__(__self__, *,
                 type: pulumi.Input[str],
                 version: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a PluginPinnedVersion resource.
        :param pulumi.Input[str] type: Type of plugin; one of "auth", "secret", or "database".
        :param pulumi.Input[str] version: Semantic version of the plugin to pin.
        :param pulumi.Input[str] name: Name of the plugin.
        """
        pulumi.set(__self__, "type", type)
        pulumi.set(__self__, "version", version)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of plugin; one of "auth", "secret", or "database".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> pulumi.Input[str]:
        """
        Semantic version of the plugin to pin.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: pulumi.Input[str]):
        pulumi.set(self, "version", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the plugin.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _PluginPinnedVersionState:
    def __init__(__self__, *,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering PluginPinnedVersion resources.
        :param pulumi.Input[str] name: Name of the plugin.
        :param pulumi.Input[str] type: Type of plugin; one of "auth", "secret", or "database".
        :param pulumi.Input[str] version: Semantic version of the plugin to pin.
        """
        if name is not None:
            pulumi.set(__self__, "name", name)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if version is not None:
            pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the plugin.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of plugin; one of "auth", "secret", or "database".
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def version(self) -> Optional[pulumi.Input[str]]:
        """
        Semantic version of the plugin to pin.
        """
        return pulumi.get(self, "version")

    @version.setter
    def version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "version", value)


class PluginPinnedVersion(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        jwt = vault.Plugin("jwt",
            type="auth",
            name="jwt",
            command="vault-plugin-auth-jwt",
            version="v0.17.0",
            sha256="6bd0a803ed742aa3ce35e4fa23d2c8d550e6c1567bf63410cec489c28b68b0fc",
            envs=["HTTP_PROXY=http://proxy.example.com:8080"])
        jwt_pin = vault.PluginPinnedVersion("jwt_pin",
            type=jwt.type,
            name=jwt.name,
            version=jwt.version)
        jwt_auth = vault.AuthBackend("jwt_auth", type=jwt_pin.name)
        ```

        ## Import

        Pinned plugin versions can be imported using `type/name` as the ID, e.g.

        ```sh
        $ pulumi import vault:index/pluginPinnedVersion:PluginPinnedVersion jwt_pin auth/jwt
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the plugin.
        :param pulumi.Input[str] type: Type of plugin; one of "auth", "secret", or "database".
        :param pulumi.Input[str] version: Semantic version of the plugin to pin.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: PluginPinnedVersionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        jwt = vault.Plugin("jwt",
            type="auth",
            name="jwt",
            command="vault-plugin-auth-jwt",
            version="v0.17.0",
            sha256="6bd0a803ed742aa3ce35e4fa23d2c8d550e6c1567bf63410cec489c28b68b0fc",
            envs=["HTTP_PROXY=http://proxy.example.com:8080"])
        jwt_pin = vault.PluginPinnedVersion("jwt_pin",
            type=jwt.type,
            name=jwt.name,
            version=jwt.version)
        jwt_auth = vault.AuthBackend("jwt_auth", type=jwt_pin.name)
        ```

        ## Import

        Pinned plugin versions can be imported using `type/name` as the ID, e.g.

        ```sh
        $ pulumi import vault:index/pluginPinnedVersion:PluginPinnedVersion jwt_pin auth/jwt
        ```

        :param str resource_name: The name of the resource.
        :param PluginPinnedVersionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(PluginPinnedVersionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 version: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = PluginPinnedVersionArgs.__new__(PluginPinnedVersionArgs)

            __props__.__dict__["name"] = name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            if version is None and not opts.urn:
                raise TypeError("Missing required property 'version'")
            __props__.__dict__["version"] = version
        super(PluginPinnedVersion, __self__).__init__(
            'vault:index/pluginPinnedVersion:PluginPinnedVersion',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            name: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            version: Optional[pulumi.Input[str]] = None) -> 'PluginPinnedVersion':
        """
        Get an existing PluginPinnedVersion resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: Name of the plugin.
        :param pulumi.Input[str] type: Type of plugin; one of "auth", "secret", or "database".
        :param pulumi.Input[str] version: Semantic version of the plugin to pin.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _PluginPinnedVersionState.__new__(_PluginPinnedVersionState)

        __props__.__dict__["name"] = name
        __props__.__dict__["type"] = type
        __props__.__dict__["version"] = version
        return PluginPinnedVersion(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the plugin.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of plugin; one of "auth", "secret", or "database".
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def version(self) -> pulumi.Output[str]:
        """
        Semantic version of the plugin to pin.
        """
        return pulumi.get(self, "version")

