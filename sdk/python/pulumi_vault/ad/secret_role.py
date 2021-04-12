# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SecretRoleArgs', 'SecretRole']

@pulumi.input_type
class SecretRoleArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 role: pulumi.Input[str],
                 service_account_name: pulumi.Input[str],
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SecretRole resource.
        :param pulumi.Input[str] backend: The mount path for the AD backend.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] service_account_name: The username/logon name for the service account with which this role will be associated.
        :param pulumi.Input[int] ttl: In seconds, the default password time-to-live.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "role", role)
        pulumi.set(__self__, "service_account_name", service_account_name)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The mount path for the AD backend.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="serviceAccountName")
    def service_account_name(self) -> pulumi.Input[str]:
        """
        The username/logon name for the service account with which this role will be associated.
        """
        return pulumi.get(self, "service_account_name")

    @service_account_name.setter
    def service_account_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_name", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        In seconds, the default password time-to-live.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


class SecretRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 service_account_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SecretRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The mount path for the AD backend.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] service_account_name: The username/logon name for the service account with which this role will be associated.
        :param pulumi.Input[int] ttl: In seconds, the default password time-to-live.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecretRole resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecretRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 service_account_name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
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

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            if service_account_name is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_name'")
            __props__['service_account_name'] = service_account_name
            __props__['ttl'] = ttl
            __props__['last_vault_rotation'] = None
            __props__['password_last_set'] = None
        super(SecretRole, __self__).__init__(
            'vault:ad/secretRole:SecretRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            last_vault_rotation: Optional[pulumi.Input[str]] = None,
            password_last_set: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            service_account_name: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None) -> 'SecretRole':
        """
        Get an existing SecretRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The mount path for the AD backend.
        :param pulumi.Input[str] last_vault_rotation: Last time Vault rotated this service account's password.
        :param pulumi.Input[str] password_last_set: Last time Vault set this service account's password.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] service_account_name: The username/logon name for the service account with which this role will be associated.
        :param pulumi.Input[int] ttl: In seconds, the default password time-to-live.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["last_vault_rotation"] = last_vault_rotation
        __props__["password_last_set"] = password_last_set
        __props__["role"] = role
        __props__["service_account_name"] = service_account_name
        __props__["ttl"] = ttl
        return SecretRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The mount path for the AD backend.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="lastVaultRotation")
    def last_vault_rotation(self) -> pulumi.Output[str]:
        """
        Last time Vault rotated this service account's password.
        """
        return pulumi.get(self, "last_vault_rotation")

    @property
    @pulumi.getter(name="passwordLastSet")
    def password_last_set(self) -> pulumi.Output[str]:
        """
        Last time Vault set this service account's password.
        """
        return pulumi.get(self, "password_last_set")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="serviceAccountName")
    def service_account_name(self) -> pulumi.Output[str]:
        """
        The username/logon name for the service account with which this role will be associated.
        """
        return pulumi.get(self, "service_account_name")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        In seconds, the default password time-to-live.
        """
        return pulumi.get(self, "ttl")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

