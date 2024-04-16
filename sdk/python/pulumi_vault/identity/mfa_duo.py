# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['MfaDuoArgs', 'MfaDuo']

@pulumi.input_type
class MfaDuoArgs:
    def __init__(__self__, *,
                 api_hostname: pulumi.Input[str],
                 integration_key: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 use_passcode: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MfaDuo resource.
        :param pulumi.Input[str] api_hostname: API hostname for Duo
        :param pulumi.Input[str] integration_key: Integration key for Duo
        :param pulumi.Input[str] secret_key: Secret key for Duo
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] push_info: Push information for Duo.
        :param pulumi.Input[bool] use_passcode: Require passcode upon MFA validation.
        :param pulumi.Input[str] username_format: A template string for mapping Identity names to MFA methods.
        """
        pulumi.set(__self__, "api_hostname", api_hostname)
        pulumi.set(__self__, "integration_key", integration_key)
        pulumi.set(__self__, "secret_key", secret_key)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if push_info is not None:
            pulumi.set(__self__, "push_info", push_info)
        if use_passcode is not None:
            pulumi.set(__self__, "use_passcode", use_passcode)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> pulumi.Input[str]:
        """
        API hostname for Duo
        """
        return pulumi.get(self, "api_hostname")

    @api_hostname.setter
    def api_hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_hostname", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Input[str]:
        """
        Integration key for Duo
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        Secret key for Duo
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> Optional[pulumi.Input[str]]:
        """
        Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @push_info.setter
    def push_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_info", value)

    @property
    @pulumi.getter(name="usePasscode")
    def use_passcode(self) -> Optional[pulumi.Input[bool]]:
        """
        Require passcode upon MFA validation.
        """
        return pulumi.get(self, "use_passcode")

    @use_passcode.setter
    def use_passcode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_passcode", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_format", value)


@pulumi.input_type
class _MfaDuoState:
    def __init__(__self__, *,
                 api_hostname: Optional[pulumi.Input[str]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 method_id: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 namespace_id: Optional[pulumi.Input[str]] = None,
                 namespace_path: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 use_passcode: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 uuid: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MfaDuo resources.
        :param pulumi.Input[str] api_hostname: API hostname for Duo
        :param pulumi.Input[str] integration_key: Integration key for Duo
        :param pulumi.Input[str] method_id: Method ID.
        :param pulumi.Input[str] mount_accessor: Mount accessor.
        :param pulumi.Input[str] name: Method name.
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] namespace_id: Method's namespace ID.
        :param pulumi.Input[str] namespace_path: Method's namespace path.
        :param pulumi.Input[str] push_info: Push information for Duo.
        :param pulumi.Input[str] secret_key: Secret key for Duo
        :param pulumi.Input[str] type: MFA type.
        :param pulumi.Input[bool] use_passcode: Require passcode upon MFA validation.
        :param pulumi.Input[str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[str] uuid: Resource UUID.
        """
        if api_hostname is not None:
            pulumi.set(__self__, "api_hostname", api_hostname)
        if integration_key is not None:
            pulumi.set(__self__, "integration_key", integration_key)
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
        if push_info is not None:
            pulumi.set(__self__, "push_info", push_info)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if use_passcode is not None:
            pulumi.set(__self__, "use_passcode", use_passcode)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        API hostname for Duo
        """
        return pulumi.get(self, "api_hostname")

    @api_hostname.setter
    def api_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_hostname", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> Optional[pulumi.Input[str]]:
        """
        Integration key for Duo
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter(name="methodId")
    def method_id(self) -> Optional[pulumi.Input[str]]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @method_id.setter
    def method_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "method_id", value)

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[str]]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_accessor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[str]]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_id", value)

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[pulumi.Input[str]]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace_path", value)

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> Optional[pulumi.Input[str]]:
        """
        Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @push_info.setter
    def push_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_info", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        Secret key for Duo
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="usePasscode")
    def use_passcode(self) -> Optional[pulumi.Input[bool]]:
        """
        Require passcode upon MFA validation.
        """
        return pulumi.get(self, "use_passcode")

    @use_passcode.setter
    def use_passcode(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_passcode", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_format", value)

    @property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[str]]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "uuid", value)


class MfaDuo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_hostname: Optional[pulumi.Input[str]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 use_passcode: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Resource for configuring the duo MFA method.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaDuo("example",
            api_hostname="api-xxxxxxxx.duosecurity.com",
            secret_key="secret-key",
            integration_key="secret-int-key")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_hostname: API hostname for Duo
        :param pulumi.Input[str] integration_key: Integration key for Duo
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] push_info: Push information for Duo.
        :param pulumi.Input[str] secret_key: Secret key for Duo
        :param pulumi.Input[bool] use_passcode: Require passcode upon MFA validation.
        :param pulumi.Input[str] username_format: A template string for mapping Identity names to MFA methods.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaDuoArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for configuring the duo MFA method.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaDuo("example",
            api_hostname="api-xxxxxxxx.duosecurity.com",
            secret_key="secret-key",
            integration_key="secret-int-key")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaDuo:MfaDuo example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param MfaDuoArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaDuoArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_hostname: Optional[pulumi.Input[str]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 use_passcode: Optional[pulumi.Input[bool]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaDuoArgs.__new__(MfaDuoArgs)

            if api_hostname is None and not opts.urn:
                raise TypeError("Missing required property 'api_hostname'")
            __props__.__dict__["api_hostname"] = api_hostname
            if integration_key is None and not opts.urn:
                raise TypeError("Missing required property 'integration_key'")
            __props__.__dict__["integration_key"] = None if integration_key is None else pulumi.Output.secret(integration_key)
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["push_info"] = push_info
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["use_passcode"] = use_passcode
            __props__.__dict__["username_format"] = username_format
            __props__.__dict__["method_id"] = None
            __props__.__dict__["mount_accessor"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_path"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["uuid"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["integrationKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MfaDuo, __self__).__init__(
            'vault:identity/mfaDuo:MfaDuo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_hostname: Optional[pulumi.Input[str]] = None,
            integration_key: Optional[pulumi.Input[str]] = None,
            method_id: Optional[pulumi.Input[str]] = None,
            mount_accessor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            namespace_id: Optional[pulumi.Input[str]] = None,
            namespace_path: Optional[pulumi.Input[str]] = None,
            push_info: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            use_passcode: Optional[pulumi.Input[bool]] = None,
            username_format: Optional[pulumi.Input[str]] = None,
            uuid: Optional[pulumi.Input[str]] = None) -> 'MfaDuo':
        """
        Get an existing MfaDuo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_hostname: API hostname for Duo
        :param pulumi.Input[str] integration_key: Integration key for Duo
        :param pulumi.Input[str] method_id: Method ID.
        :param pulumi.Input[str] mount_accessor: Mount accessor.
        :param pulumi.Input[str] name: Method name.
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[str] namespace_id: Method's namespace ID.
        :param pulumi.Input[str] namespace_path: Method's namespace path.
        :param pulumi.Input[str] push_info: Push information for Duo.
        :param pulumi.Input[str] secret_key: Secret key for Duo
        :param pulumi.Input[str] type: MFA type.
        :param pulumi.Input[bool] use_passcode: Require passcode upon MFA validation.
        :param pulumi.Input[str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[str] uuid: Resource UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaDuoState.__new__(_MfaDuoState)

        __props__.__dict__["api_hostname"] = api_hostname
        __props__.__dict__["integration_key"] = integration_key
        __props__.__dict__["method_id"] = method_id
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["push_info"] = push_info
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["type"] = type
        __props__.__dict__["use_passcode"] = use_passcode
        __props__.__dict__["username_format"] = username_format
        __props__.__dict__["uuid"] = uuid
        return MfaDuo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> pulumi.Output[str]:
        """
        API hostname for Duo
        """
        return pulumi.get(self, "api_hostname")

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Output[str]:
        """
        Integration key for Duo
        """
        return pulumi.get(self, "integration_key")

    @property
    @pulumi.getter(name="methodId")
    def method_id(self) -> pulumi.Output[str]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[str]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[str]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Output[str]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> pulumi.Output[Optional[str]]:
        """
        Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        Secret key for Duo
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usePasscode")
    def use_passcode(self) -> pulumi.Output[Optional[bool]]:
        """
        Require passcode upon MFA validation.
        """
        return pulumi.get(self, "use_passcode")

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> pulumi.Output[Optional[str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[str]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

