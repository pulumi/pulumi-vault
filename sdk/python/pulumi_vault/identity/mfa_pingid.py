# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['MfaPingidArgs', 'MfaPingid']

@pulumi.input_type
class MfaPingidArgs:
    def __init__(__self__, *,
                 settings_file_base64: pulumi.Input[_builtins.str],
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a MfaPingid resource.
        :param pulumi.Input[_builtins.str] settings_file_base64: A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        """
        pulumi.set(__self__, "settings_file_base64", settings_file_base64)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @_builtins.property
    @pulumi.getter(name="settingsFileBase64")
    def settings_file_base64(self) -> pulumi.Input[_builtins.str]:
        """
        A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        """
        return pulumi.get(self, "settings_file_base64")

    @settings_file_base64.setter
    def settings_file_base64(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "settings_file_base64", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_format", value)


@pulumi.input_type
class _MfaPingidState:
    def __init__(__self__, *,
                 admin_url: Optional[pulumi.Input[_builtins.str]] = None,
                 authenticator_url: Optional[pulumi.Input[_builtins.str]] = None,
                 idp_url: Optional[pulumi.Input[_builtins.str]] = None,
                 method_id: Optional[pulumi.Input[_builtins.str]] = None,
                 mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
                 org_alias: Optional[pulumi.Input[_builtins.str]] = None,
                 settings_file_base64: Optional[pulumi.Input[_builtins.str]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 use_signature: Optional[pulumi.Input[_builtins.bool]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 uuid: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering MfaPingid resources.
        :param pulumi.Input[_builtins.str] admin_url: The admin URL, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] authenticator_url: A unique identifier of the organization, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] idp_url: The IDP URL, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] org_alias: The name of the PingID client organization, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] settings_file_base64: A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.bool] use_signature: Use signature value, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        if admin_url is not None:
            pulumi.set(__self__, "admin_url", admin_url)
        if authenticator_url is not None:
            pulumi.set(__self__, "authenticator_url", authenticator_url)
        if idp_url is not None:
            pulumi.set(__self__, "idp_url", idp_url)
        if method_id is not None:
            pulumi.set(__self__, "method_id", method_id)
        if mount_accessor is not None:
            pulumi.set(__self__, "mount_accessor", mount_accessor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if org_alias is not None:
            pulumi.set(__self__, "org_alias", org_alias)
        if settings_file_base64 is not None:
            pulumi.set(__self__, "settings_file_base64", settings_file_base64)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if use_signature is not None:
            pulumi.set(__self__, "use_signature", use_signature)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @_builtins.property
    @pulumi.getter(name="adminUrl")
    def admin_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The admin URL, derived from "settings_file_base64"
        """
        return pulumi.get(self, "admin_url")

    @admin_url.setter
    def admin_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "admin_url", value)

    @_builtins.property
    @pulumi.getter(name="authenticatorUrl")
    def authenticator_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A unique identifier of the organization, derived from "settings_file_base64"
        """
        return pulumi.get(self, "authenticator_url")

    @authenticator_url.setter
    def authenticator_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "authenticator_url", value)

    @_builtins.property
    @pulumi.getter(name="idpUrl")
    def idp_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The IDP URL, derived from "settings_file_base64"
        """
        return pulumi.get(self, "idp_url")

    @idp_url.setter
    def idp_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "idp_url", value)

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @method_id.setter
    def method_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "method_id", value)

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mount_accessor", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_id", value)

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_path", value)

    @_builtins.property
    @pulumi.getter(name="orgAlias")
    def org_alias(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the PingID client organization, derived from "settings_file_base64"
        """
        return pulumi.get(self, "org_alias")

    @org_alias.setter
    def org_alias(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "org_alias", value)

    @_builtins.property
    @pulumi.getter(name="settingsFileBase64")
    def settings_file_base64(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        """
        return pulumi.get(self, "settings_file_base64")

    @settings_file_base64.setter
    def settings_file_base64(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "settings_file_base64", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter(name="useSignature")
    def use_signature(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Use signature value, derived from "settings_file_base64"
        """
        return pulumi.get(self, "use_signature")

    @use_signature.setter
    def use_signature(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "use_signature", value)

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_format", value)

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "uuid", value)


@pulumi.type_token("vault:identity/mfaPingid:MfaPingid")
class MfaPingid(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 settings_file_base64: Optional[pulumi.Input[_builtins.str]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Resource for configuring the pingid MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaPingid("example", settings_file_base64="CnVzZV9iYXNlNjR[...]HBtCg==")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaPingid:MfaPingid example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] settings_file_base64: A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaPingidArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for configuring the pingid MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaPingid("example", settings_file_base64="CnVzZV9iYXNlNjR[...]HBtCg==")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaPingid:MfaPingid example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param MfaPingidArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaPingidArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 settings_file_base64: Optional[pulumi.Input[_builtins.str]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaPingidArgs.__new__(MfaPingidArgs)

            __props__.__dict__["namespace"] = namespace
            if settings_file_base64 is None and not opts.urn:
                raise TypeError("Missing required property 'settings_file_base64'")
            __props__.__dict__["settings_file_base64"] = settings_file_base64
            __props__.__dict__["username_format"] = username_format
            __props__.__dict__["admin_url"] = None
            __props__.__dict__["authenticator_url"] = None
            __props__.__dict__["idp_url"] = None
            __props__.__dict__["method_id"] = None
            __props__.__dict__["mount_accessor"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_path"] = None
            __props__.__dict__["org_alias"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["use_signature"] = None
            __props__.__dict__["uuid"] = None
        super(MfaPingid, __self__).__init__(
            'vault:identity/mfaPingid:MfaPingid',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            admin_url: Optional[pulumi.Input[_builtins.str]] = None,
            authenticator_url: Optional[pulumi.Input[_builtins.str]] = None,
            idp_url: Optional[pulumi.Input[_builtins.str]] = None,
            method_id: Optional[pulumi.Input[_builtins.str]] = None,
            mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
            org_alias: Optional[pulumi.Input[_builtins.str]] = None,
            settings_file_base64: Optional[pulumi.Input[_builtins.str]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None,
            use_signature: Optional[pulumi.Input[_builtins.bool]] = None,
            username_format: Optional[pulumi.Input[_builtins.str]] = None,
            uuid: Optional[pulumi.Input[_builtins.str]] = None) -> 'MfaPingid':
        """
        Get an existing MfaPingid resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] admin_url: The admin URL, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] authenticator_url: A unique identifier of the organization, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] idp_url: The IDP URL, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] org_alias: The name of the PingID client organization, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] settings_file_base64: A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.bool] use_signature: Use signature value, derived from "settings_file_base64"
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaPingidState.__new__(_MfaPingidState)

        __props__.__dict__["admin_url"] = admin_url
        __props__.__dict__["authenticator_url"] = authenticator_url
        __props__.__dict__["idp_url"] = idp_url
        __props__.__dict__["method_id"] = method_id
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["org_alias"] = org_alias
        __props__.__dict__["settings_file_base64"] = settings_file_base64
        __props__.__dict__["type"] = type
        __props__.__dict__["use_signature"] = use_signature
        __props__.__dict__["username_format"] = username_format
        __props__.__dict__["uuid"] = uuid
        return MfaPingid(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="adminUrl")
    def admin_url(self) -> pulumi.Output[_builtins.str]:
        """
        The admin URL, derived from "settings_file_base64"
        """
        return pulumi.get(self, "admin_url")

    @_builtins.property
    @pulumi.getter(name="authenticatorUrl")
    def authenticator_url(self) -> pulumi.Output[_builtins.str]:
        """
        A unique identifier of the organization, derived from "settings_file_base64"
        """
        return pulumi.get(self, "authenticator_url")

    @_builtins.property
    @pulumi.getter(name="idpUrl")
    def idp_url(self) -> pulumi.Output[_builtins.str]:
        """
        The IDP URL, derived from "settings_file_base64"
        """
        return pulumi.get(self, "idp_url")

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[_builtins.str]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @_builtins.property
    @pulumi.getter(name="orgAlias")
    def org_alias(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the PingID client organization, derived from "settings_file_base64"
        """
        return pulumi.get(self, "org_alias")

    @_builtins.property
    @pulumi.getter(name="settingsFileBase64")
    def settings_file_base64(self) -> pulumi.Output[_builtins.str]:
        """
        A base64-encoded third-party settings contents as retrieved from PingID's configuration page.
        """
        return pulumi.get(self, "settings_file_base64")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter(name="useSignature")
    def use_signature(self) -> pulumi.Output[_builtins.bool]:
        """
        Use signature value, derived from "settings_file_base64"
        """
        return pulumi.get(self, "use_signature")

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[_builtins.str]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

