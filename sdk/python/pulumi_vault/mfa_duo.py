# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MfaDuoArgs', 'MfaDuo']

@pulumi.input_type
class MfaDuoArgs:
    def __init__(__self__, *,
                 api_hostname: pulumi.Input[str],
                 integration_key: pulumi.Input[str],
                 mount_accessor: pulumi.Input[str],
                 secret_key: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 username_format: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a MfaDuo resource.
        :param pulumi.Input[str] api_hostname: `(string: <required>)` - API hostname for Duo.
        :param pulumi.Input[str] integration_key: `(string: <required>)` - Integration key for Duo.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] secret_key: `(string: <required>)` - Secret key for Duo.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] push_info: `(string)` - Push information for Duo.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        pulumi.set(__self__, "api_hostname", api_hostname)
        pulumi.set(__self__, "integration_key", integration_key)
        pulumi.set(__self__, "mount_accessor", mount_accessor)
        pulumi.set(__self__, "secret_key", secret_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if push_info is not None:
            pulumi.set(__self__, "push_info", push_info)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - API hostname for Duo.
        """
        return pulumi.get(self, "api_hostname")

    @api_hostname.setter
    def api_hostname(self, value: pulumi.Input[str]):
        pulumi.set(self, "api_hostname", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - Integration key for Duo.
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount_accessor", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Input[str]:
        """
        `(string: <required>)` - Secret key for Duo.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @push_info.setter
    def push_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_info", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
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
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 username_format: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering MfaDuo resources.
        :param pulumi.Input[str] api_hostname: `(string: <required>)` - API hostname for Duo.
        :param pulumi.Input[str] integration_key: `(string: <required>)` - Integration key for Duo.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] push_info: `(string)` - Push information for Duo.
        :param pulumi.Input[str] secret_key: `(string: <required>)` - Secret key for Duo.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        if api_hostname is not None:
            pulumi.set(__self__, "api_hostname", api_hostname)
        if integration_key is not None:
            pulumi.set(__self__, "integration_key", integration_key)
        if mount_accessor is not None:
            pulumi.set(__self__, "mount_accessor", mount_accessor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if push_info is not None:
            pulumi.set(__self__, "push_info", push_info)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - API hostname for Duo.
        """
        return pulumi.get(self, "api_hostname")

    @api_hostname.setter
    def api_hostname(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_hostname", value)

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - Integration key for Duo.
        """
        return pulumi.get(self, "integration_key")

    @integration_key.setter
    def integration_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "integration_key", value)

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount_accessor", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @push_info.setter
    def push_info(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "push_info", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        `(string: <required>)` - Secret key for Duo.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_format", value)


class MfaDuo(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_hostname: Optional[pulumi.Input[str]] = None,
                 integration_key: Optional[pulumi.Input[str]] = None,
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 username_format: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        userpass = vault.AuthBackend("userpass",
            type="userpass",
            path="userpass")
        my_duo = vault.MfaDuo("myDuo",
            mount_accessor=userpass.accessor,
            secret_key="8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz",
            integration_key="BIACEUEAXI20BNWTEYXT",
            api_hostname="api-2b5c39f5.duosecurity.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_hostname: `(string: <required>)` - API hostname for Duo.
        :param pulumi.Input[str] integration_key: `(string: <required>)` - Integration key for Duo.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] push_info: `(string)` - Push information for Duo.
        :param pulumi.Input[str] secret_key: `(string: <required>)` - Secret key for Duo.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaDuoArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage [Duo MFA](https://www.vaultproject.io/docs/enterprise/mfa/mfa-duo.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        userpass = vault.AuthBackend("userpass",
            type="userpass",
            path="userpass")
        my_duo = vault.MfaDuo("myDuo",
            mount_accessor=userpass.accessor,
            secret_key="8C7THtrIigh2rPZQMbguugt8IUftWhMRCOBzbuyz",
            integration_key="BIACEUEAXI20BNWTEYXT",
            api_hostname="api-2b5c39f5.duosecurity.com")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:index/mfaDuo:MfaDuo my_duo my_duo
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
                 mount_accessor: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 push_info: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
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
            if mount_accessor is None and not opts.urn:
                raise TypeError("Missing required property 'mount_accessor'")
            __props__.__dict__["mount_accessor"] = mount_accessor
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["push_info"] = push_info
            if secret_key is None and not opts.urn:
                raise TypeError("Missing required property 'secret_key'")
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["username_format"] = username_format
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["integrationKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MfaDuo, __self__).__init__(
            'vault:index/mfaDuo:MfaDuo',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_hostname: Optional[pulumi.Input[str]] = None,
            integration_key: Optional[pulumi.Input[str]] = None,
            mount_accessor: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            push_info: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            username_format: Optional[pulumi.Input[str]] = None) -> 'MfaDuo':
        """
        Get an existing MfaDuo resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_hostname: `(string: <required>)` - API hostname for Duo.
        :param pulumi.Input[str] integration_key: `(string: <required>)` - Integration key for Duo.
        :param pulumi.Input[str] mount_accessor: `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        :param pulumi.Input[str] name: `(string: <required>)` – Name of the MFA method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] push_info: `(string)` - Push information for Duo.
        :param pulumi.Input[str] secret_key: `(string: <required>)` - Secret key for Duo.
        :param pulumi.Input[str] username_format: `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
               - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
               - entity.name: The name configured for the Entity
               - alias.metadata.`<key>`: The value of the Alias's metadata parameter
               - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaDuoState.__new__(_MfaDuoState)

        __props__.__dict__["api_hostname"] = api_hostname
        __props__.__dict__["integration_key"] = integration_key
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["push_info"] = push_info
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["username_format"] = username_format
        return MfaDuo(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="apiHostname")
    def api_hostname(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - API hostname for Duo.
        """
        return pulumi.get(self, "api_hostname")

    @property
    @pulumi.getter(name="integrationKey")
    def integration_key(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - Integration key for Duo.
        """
        return pulumi.get(self, "integration_key")

    @property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - The mount to tie this method to for use in automatic mappings. The mapping will use the Name field of Aliases associated with this mount as the username in the mapping.
        """
        return pulumi.get(self, "mount_accessor")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` – Name of the MFA method.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="pushInfo")
    def push_info(self) -> pulumi.Output[Optional[str]]:
        """
        `(string)` - Push information for Duo.
        """
        return pulumi.get(self, "push_info")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[str]:
        """
        `(string: <required>)` - Secret key for Duo.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> pulumi.Output[Optional[str]]:
        """
        `(string)` - A format string for mapping Identity names to MFA method names. Values to substitute should be placed in `{{}}`. For example, `"{{alias.name}}@example.com"`. If blank, the Alias's Name field will be used as-is. Currently-supported mappings:
        - alias.name: The name returned by the mount configured via the `mount_accessor` parameter
        - entity.name: The name configured for the Entity
        - alias.metadata.`<key>`: The value of the Alias's metadata parameter
        - entity.metadata.`<key>`: The value of the Entity's metadata parameter
        """
        return pulumi.get(self, "username_format")

