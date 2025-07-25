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

__all__ = ['SecretBackendArgs', 'SecretBackend']

@pulumi.input_type
class SecretBackendArgs:
    def __init__(__self__, *,
                 connection_uri: pulumi.Input[_builtins.str],
                 password: pulumi.Input[_builtins.str],
                 username: pulumi.Input[_builtins.str],
                 default_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 disable_remount: Optional[pulumi.Input[_builtins.bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 password_policy: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 username_template: Optional[pulumi.Input[_builtins.str]] = None,
                 verify_connection: Optional[pulumi.Input[_builtins.bool]] = None):
        """
        The set of arguments for constructing a SecretBackend resource.
        :param pulumi.Input[_builtins.str] connection_uri: Specifies the RabbitMQ connection URI.
        :param pulumi.Input[_builtins.str] password: Specifies the RabbitMQ management administrator password.
        :param pulumi.Input[_builtins.str] username: Specifies the RabbitMQ management administrator username.
        :param pulumi.Input[_builtins.int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[_builtins.str] description: A human-friendly description for this backend.
        :param pulumi.Input[_builtins.bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[_builtins.int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] password_policy: Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        :param pulumi.Input[_builtins.str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `rabbitmq`.
        :param pulumi.Input[_builtins.str] username_template: Template describing how dynamic usernames are generated.
        :param pulumi.Input[_builtins.bool] verify_connection: Specifies whether to verify connection URI, username, and password.
               Defaults to `true`.
        """
        pulumi.set(__self__, "connection_uri", connection_uri)
        pulumi.set(__self__, "password", password)
        pulumi.set(__self__, "username", username)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if password_policy is not None:
            pulumi.set(__self__, "password_policy", password_policy)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if username_template is not None:
            pulumi.set(__self__, "username_template", username_template)
        if verify_connection is not None:
            pulumi.set(__self__, "verify_connection", verify_connection)

    @_builtins.property
    @pulumi.getter(name="connectionUri")
    def connection_uri(self) -> pulumi.Input[_builtins.str]:
        """
        Specifies the RabbitMQ connection URI.
        """
        return pulumi.get(self, "connection_uri")

    @connection_uri.setter
    def connection_uri(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "connection_uri", value)

    @_builtins.property
    @pulumi.getter
    def password(self) -> pulumi.Input[_builtins.str]:
        """
        Specifies the RabbitMQ management administrator password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "password", value)

    @_builtins.property
    @pulumi.getter
    def username(self) -> pulumi.Input[_builtins.str]:
        """
        Specifies the RabbitMQ management administrator username.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "username", value)

    @_builtins.property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "disable_remount", value)

    @_builtins.property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        """
        return pulumi.get(self, "password_policy")

    @password_policy.setter
    def password_policy(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password_policy", value)

    @_builtins.property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `rabbitmq`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Template describing how dynamic usernames are generated.
        """
        return pulumi.get(self, "username_template")

    @username_template.setter
    def username_template(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_template", value)

    @_builtins.property
    @pulumi.getter(name="verifyConnection")
    def verify_connection(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies whether to verify connection URI, username, and password.
        Defaults to `true`.
        """
        return pulumi.get(self, "verify_connection")

    @verify_connection.setter
    def verify_connection(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "verify_connection", value)


@pulumi.input_type
class _SecretBackendState:
    def __init__(__self__, *,
                 connection_uri: Optional[pulumi.Input[_builtins.str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 disable_remount: Optional[pulumi.Input[_builtins.bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 password_policy: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None,
                 username_template: Optional[pulumi.Input[_builtins.str]] = None,
                 verify_connection: Optional[pulumi.Input[_builtins.bool]] = None):
        """
        Input properties used for looking up and filtering SecretBackend resources.
        :param pulumi.Input[_builtins.str] connection_uri: Specifies the RabbitMQ connection URI.
        :param pulumi.Input[_builtins.int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[_builtins.str] description: A human-friendly description for this backend.
        :param pulumi.Input[_builtins.bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[_builtins.int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] password: Specifies the RabbitMQ management administrator password.
        :param pulumi.Input[_builtins.str] password_policy: Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        :param pulumi.Input[_builtins.str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `rabbitmq`.
        :param pulumi.Input[_builtins.str] username: Specifies the RabbitMQ management administrator username.
        :param pulumi.Input[_builtins.str] username_template: Template describing how dynamic usernames are generated.
        :param pulumi.Input[_builtins.bool] verify_connection: Specifies whether to verify connection URI, username, and password.
               Defaults to `true`.
        """
        if connection_uri is not None:
            pulumi.set(__self__, "connection_uri", connection_uri)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if password is not None:
            pulumi.set(__self__, "password", password)
        if password_policy is not None:
            pulumi.set(__self__, "password_policy", password_policy)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if username is not None:
            pulumi.set(__self__, "username", username)
        if username_template is not None:
            pulumi.set(__self__, "username_template", username_template)
        if verify_connection is not None:
            pulumi.set(__self__, "verify_connection", verify_connection)

    @_builtins.property
    @pulumi.getter(name="connectionUri")
    def connection_uri(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the RabbitMQ connection URI.
        """
        return pulumi.get(self, "connection_uri")

    @connection_uri.setter
    def connection_uri(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "connection_uri", value)

    @_builtins.property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @_builtins.property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "description", value)

    @_builtins.property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "disable_remount", value)

    @_builtins.property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter
    def password(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the RabbitMQ management administrator password.
        """
        return pulumi.get(self, "password")

    @password.setter
    def password(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password", value)

    @_builtins.property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        """
        return pulumi.get(self, "password_policy")

    @password_policy.setter
    def password_policy(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "password_policy", value)

    @_builtins.property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `rabbitmq`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "path", value)

    @_builtins.property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the RabbitMQ management administrator username.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username", value)

    @_builtins.property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Template describing how dynamic usernames are generated.
        """
        return pulumi.get(self, "username_template")

    @username_template.setter
    def username_template(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_template", value)

    @_builtins.property
    @pulumi.getter(name="verifyConnection")
    def verify_connection(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies whether to verify connection URI, username, and password.
        Defaults to `true`.
        """
        return pulumi.get(self, "verify_connection")

    @verify_connection.setter
    def verify_connection(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "verify_connection", value)


@pulumi.type_token("vault:rabbitMq/secretBackend:SecretBackend")
class SecretBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_uri: Optional[pulumi.Input[_builtins.str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 disable_remount: Optional[pulumi.Input[_builtins.bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 password_policy: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None,
                 username_template: Optional[pulumi.Input[_builtins.str]] = None,
                 verify_connection: Optional[pulumi.Input[_builtins.bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        rabbitmq = vault.rabbitmq.SecretBackend("rabbitmq",
            connection_uri="https://.....",
            username="user",
            password="password")
        ```

        ## Import

        RabbitMQ secret backends can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:rabbitMq/secretBackend:SecretBackend rabbitmq rabbitmq
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] connection_uri: Specifies the RabbitMQ connection URI.
        :param pulumi.Input[_builtins.int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[_builtins.str] description: A human-friendly description for this backend.
        :param pulumi.Input[_builtins.bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[_builtins.int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] password: Specifies the RabbitMQ management administrator password.
        :param pulumi.Input[_builtins.str] password_policy: Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        :param pulumi.Input[_builtins.str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `rabbitmq`.
        :param pulumi.Input[_builtins.str] username: Specifies the RabbitMQ management administrator username.
        :param pulumi.Input[_builtins.str] username_template: Template describing how dynamic usernames are generated.
        :param pulumi.Input[_builtins.bool] verify_connection: Specifies whether to verify connection URI, username, and password.
               Defaults to `true`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        rabbitmq = vault.rabbitmq.SecretBackend("rabbitmq",
            connection_uri="https://.....",
            username="user",
            password="password")
        ```

        ## Import

        RabbitMQ secret backends can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:rabbitMq/secretBackend:SecretBackend rabbitmq rabbitmq
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 connection_uri: Optional[pulumi.Input[_builtins.str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 description: Optional[pulumi.Input[_builtins.str]] = None,
                 disable_remount: Optional[pulumi.Input[_builtins.bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 password: Optional[pulumi.Input[_builtins.str]] = None,
                 password_policy: Optional[pulumi.Input[_builtins.str]] = None,
                 path: Optional[pulumi.Input[_builtins.str]] = None,
                 username: Optional[pulumi.Input[_builtins.str]] = None,
                 username_template: Optional[pulumi.Input[_builtins.str]] = None,
                 verify_connection: Optional[pulumi.Input[_builtins.bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendArgs.__new__(SecretBackendArgs)

            if connection_uri is None and not opts.urn:
                raise TypeError("Missing required property 'connection_uri'")
            __props__.__dict__["connection_uri"] = connection_uri
            __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
            __props__.__dict__["description"] = description
            __props__.__dict__["disable_remount"] = disable_remount
            __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
            __props__.__dict__["namespace"] = namespace
            if password is None and not opts.urn:
                raise TypeError("Missing required property 'password'")
            __props__.__dict__["password"] = None if password is None else pulumi.Output.secret(password)
            __props__.__dict__["password_policy"] = password_policy
            __props__.__dict__["path"] = path
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = None if username is None else pulumi.Output.secret(username)
            __props__.__dict__["username_template"] = username_template
            __props__.__dict__["verify_connection"] = verify_connection
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["password", "username"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretBackend, __self__).__init__(
            'vault:rabbitMq/secretBackend:SecretBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            connection_uri: Optional[pulumi.Input[_builtins.str]] = None,
            default_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
            description: Optional[pulumi.Input[_builtins.str]] = None,
            disable_remount: Optional[pulumi.Input[_builtins.bool]] = None,
            max_lease_ttl_seconds: Optional[pulumi.Input[_builtins.int]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            password: Optional[pulumi.Input[_builtins.str]] = None,
            password_policy: Optional[pulumi.Input[_builtins.str]] = None,
            path: Optional[pulumi.Input[_builtins.str]] = None,
            username: Optional[pulumi.Input[_builtins.str]] = None,
            username_template: Optional[pulumi.Input[_builtins.str]] = None,
            verify_connection: Optional[pulumi.Input[_builtins.bool]] = None) -> 'SecretBackend':
        """
        Get an existing SecretBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] connection_uri: Specifies the RabbitMQ connection URI.
        :param pulumi.Input[_builtins.int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[_builtins.str] description: A human-friendly description for this backend.
        :param pulumi.Input[_builtins.bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[_builtins.int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] password: Specifies the RabbitMQ management administrator password.
        :param pulumi.Input[_builtins.str] password_policy: Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        :param pulumi.Input[_builtins.str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `rabbitmq`.
        :param pulumi.Input[_builtins.str] username: Specifies the RabbitMQ management administrator username.
        :param pulumi.Input[_builtins.str] username_template: Template describing how dynamic usernames are generated.
        :param pulumi.Input[_builtins.bool] verify_connection: Specifies whether to verify connection URI, username, and password.
               Defaults to `true`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendState.__new__(_SecretBackendState)

        __props__.__dict__["connection_uri"] = connection_uri
        __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
        __props__.__dict__["description"] = description
        __props__.__dict__["disable_remount"] = disable_remount
        __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["password"] = password
        __props__.__dict__["password_policy"] = password_policy
        __props__.__dict__["path"] = path
        __props__.__dict__["username"] = username
        __props__.__dict__["username_template"] = username_template
        __props__.__dict__["verify_connection"] = verify_connection
        return SecretBackend(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="connectionUri")
    def connection_uri(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies the RabbitMQ connection URI.
        """
        return pulumi.get(self, "connection_uri")

    @_builtins.property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> pulumi.Output[_builtins.int]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @_builtins.property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @_builtins.property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @_builtins.property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> pulumi.Output[_builtins.int]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter
    def password(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies the RabbitMQ management administrator password.
        """
        return pulumi.get(self, "password")

    @_builtins.property
    @pulumi.getter(name="passwordPolicy")
    def password_policy(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies a password policy to use when creating dynamic credentials. Defaults to generating an alphanumeric password if not set.
        """
        return pulumi.get(self, "password_policy")

    @_builtins.property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `rabbitmq`.
        """
        return pulumi.get(self, "path")

    @_builtins.property
    @pulumi.getter
    def username(self) -> pulumi.Output[_builtins.str]:
        """
        Specifies the RabbitMQ management administrator username.
        """
        return pulumi.get(self, "username")

    @_builtins.property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Template describing how dynamic usernames are generated.
        """
        return pulumi.get(self, "username_template")

    @_builtins.property
    @pulumi.getter(name="verifyConnection")
    def verify_connection(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Specifies whether to verify connection URI, username, and password.
        Defaults to `true`.
        """
        return pulumi.get(self, "verify_connection")

