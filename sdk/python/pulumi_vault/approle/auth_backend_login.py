# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['AuthBackendLoginArgs', 'AuthBackendLogin']

@pulumi.input_type
class AuthBackendLoginArgs:
    def __init__(__self__, *,
                 role_id: pulumi.Input[builtins.str],
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AuthBackendLogin resource.
        :param pulumi.Input[builtins.str] role_id: The ID of the role to log in with.
        :param pulumi.Input[builtins.str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[builtins.str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        pulumi.set(__self__, "role_id", role_id)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Input[builtins.str]:
        """
        The ID of the role to log in with.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique path of the Vault backend to log in with.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The secret ID of the role to log in with. Required
        unless `bind_secret_id` is set to false on the role.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_id", value)


@pulumi.input_type
class _AuthBackendLoginState:
    def __init__(__self__, *,
                 accessor: Optional[pulumi.Input[builtins.str]] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 client_token: Optional[pulumi.Input[builtins.str]] = None,
                 lease_duration: Optional[pulumi.Input[builtins.int]] = None,
                 lease_started: Optional[pulumi.Input[builtins.str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 renewable: Optional[pulumi.Input[builtins.bool]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering AuthBackendLogin resources.
        :param pulumi.Input[builtins.str] accessor: The accessor for the token.
        :param pulumi.Input[builtins.str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[builtins.str] client_token: The Vault token created.
        :param pulumi.Input[builtins.int] lease_duration: How long the token is valid for, in seconds.
        :param pulumi.Input[builtins.str] lease_started: The date and time the lease started, in RFC 3339 format.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] metadata: The metadata associated with the token.
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: A list of policies applied to the token.
        :param pulumi.Input[builtins.bool] renewable: Whether the token is renewable or not.
        :param pulumi.Input[builtins.str] role_id: The ID of the role to log in with.
        :param pulumi.Input[builtins.str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        if accessor is not None:
            pulumi.set(__self__, "accessor", accessor)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if client_token is not None:
            pulumi.set(__self__, "client_token", client_token)
        if lease_duration is not None:
            pulumi.set(__self__, "lease_duration", lease_duration)
        if lease_started is not None:
            pulumi.set(__self__, "lease_started", lease_started)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if renewable is not None:
            pulumi.set(__self__, "renewable", renewable)
        if role_id is not None:
            pulumi.set(__self__, "role_id", role_id)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)

    @property
    @pulumi.getter
    def accessor(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The accessor for the token.
        """
        return pulumi.get(self, "accessor")

    @accessor.setter
    def accessor(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "accessor", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The unique path of the Vault backend to log in with.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The Vault token created.
        """
        return pulumi.get(self, "client_token")

    @client_token.setter
    def client_token(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "client_token", value)

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> Optional[pulumi.Input[builtins.int]]:
        """
        How long the token is valid for, in seconds.
        """
        return pulumi.get(self, "lease_duration")

    @lease_duration.setter
    def lease_duration(self, value: Optional[pulumi.Input[builtins.int]]):
        pulumi.set(self, "lease_duration", value)

    @property
    @pulumi.getter(name="leaseStarted")
    def lease_started(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The date and time the lease started, in RFC 3339 format.
        """
        return pulumi.get(self, "lease_started")

    @lease_started.setter
    def lease_started(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "lease_started", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        The metadata associated with the token.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A list of policies applied to the token.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def renewable(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Whether the token is renewable or not.
        """
        return pulumi.get(self, "renewable")

    @renewable.setter
    def renewable(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "renewable", value)

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The ID of the role to log in with.
        """
        return pulumi.get(self, "role_id")

    @role_id.setter
    def role_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "role_id", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The secret ID of the role to log in with. Required
        unless `bind_secret_id` is set to false on the role.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "secret_id", value)


class AuthBackendLogin(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Logs into Vault using the AppRole auth backend. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/approle) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        approle = vault.AuthBackend("approle", type="approle")
        example = vault.app_role.AuthBackendRole("example",
            backend=approle.path,
            role_name="test-role",
            token_policies=[
                "default",
                "dev",
                "prod",
            ])
        id = vault.app_role.AuthBackendRoleSecretId("id",
            backend=approle.path,
            role_name=example.role_name)
        login = vault.app_role.AuthBackendLogin("login",
            backend=approle.path,
            role_id=example.role_id,
            secret_id=id.secret_id)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[builtins.str] role_id: The ID of the role to log in with.
        :param pulumi.Input[builtins.str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendLoginArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Logs into Vault using the AppRole auth backend. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/approle) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        approle = vault.AuthBackend("approle", type="approle")
        example = vault.app_role.AuthBackendRole("example",
            backend=approle.path,
            role_name="test-role",
            token_policies=[
                "default",
                "dev",
                "prod",
            ])
        id = vault.app_role.AuthBackendRoleSecretId("id",
            backend=approle.path,
            role_name=example.role_name)
        login = vault.app_role.AuthBackendLogin("login",
            backend=approle.path,
            role_id=example.role_id,
            secret_id=id.secret_id)
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendLoginArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendLoginArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 role_id: Optional[pulumi.Input[builtins.str]] = None,
                 secret_id: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendLoginArgs.__new__(AuthBackendLoginArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["namespace"] = namespace
            if role_id is None and not opts.urn:
                raise TypeError("Missing required property 'role_id'")
            __props__.__dict__["role_id"] = role_id
            __props__.__dict__["secret_id"] = None if secret_id is None else pulumi.Output.secret(secret_id)
            __props__.__dict__["accessor"] = None
            __props__.__dict__["client_token"] = None
            __props__.__dict__["lease_duration"] = None
            __props__.__dict__["lease_started"] = None
            __props__.__dict__["metadata"] = None
            __props__.__dict__["policies"] = None
            __props__.__dict__["renewable"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientToken", "secretId"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AuthBackendLogin, __self__).__init__(
            'vault:appRole/authBackendLogin:AuthBackendLogin',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessor: Optional[pulumi.Input[builtins.str]] = None,
            backend: Optional[pulumi.Input[builtins.str]] = None,
            client_token: Optional[pulumi.Input[builtins.str]] = None,
            lease_duration: Optional[pulumi.Input[builtins.int]] = None,
            lease_started: Optional[pulumi.Input[builtins.str]] = None,
            metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
            namespace: Optional[pulumi.Input[builtins.str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            renewable: Optional[pulumi.Input[builtins.bool]] = None,
            role_id: Optional[pulumi.Input[builtins.str]] = None,
            secret_id: Optional[pulumi.Input[builtins.str]] = None) -> 'AuthBackendLogin':
        """
        Get an existing AuthBackendLogin resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] accessor: The accessor for the token.
        :param pulumi.Input[builtins.str] backend: The unique path of the Vault backend to log in with.
        :param pulumi.Input[builtins.str] client_token: The Vault token created.
        :param pulumi.Input[builtins.int] lease_duration: How long the token is valid for, in seconds.
        :param pulumi.Input[builtins.str] lease_started: The date and time the lease started, in RFC 3339 format.
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] metadata: The metadata associated with the token.
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] policies: A list of policies applied to the token.
        :param pulumi.Input[builtins.bool] renewable: Whether the token is renewable or not.
        :param pulumi.Input[builtins.str] role_id: The ID of the role to log in with.
        :param pulumi.Input[builtins.str] secret_id: The secret ID of the role to log in with. Required
               unless `bind_secret_id` is set to false on the role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendLoginState.__new__(_AuthBackendLoginState)

        __props__.__dict__["accessor"] = accessor
        __props__.__dict__["backend"] = backend
        __props__.__dict__["client_token"] = client_token
        __props__.__dict__["lease_duration"] = lease_duration
        __props__.__dict__["lease_started"] = lease_started
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["policies"] = policies
        __props__.__dict__["renewable"] = renewable
        __props__.__dict__["role_id"] = role_id
        __props__.__dict__["secret_id"] = secret_id
        return AuthBackendLogin(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessor(self) -> pulumi.Output[builtins.str]:
        """
        The accessor for the token.
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The unique path of the Vault backend to log in with.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="clientToken")
    def client_token(self) -> pulumi.Output[builtins.str]:
        """
        The Vault token created.
        """
        return pulumi.get(self, "client_token")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> pulumi.Output[builtins.int]:
        """
        How long the token is valid for, in seconds.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseStarted")
    def lease_started(self) -> pulumi.Output[builtins.str]:
        """
        The date and time the lease started, in RFC 3339 format.
        """
        return pulumi.get(self, "lease_started")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, builtins.str]]:
        """
        The metadata associated with the token.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        A list of policies applied to the token.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def renewable(self) -> pulumi.Output[builtins.bool]:
        """
        Whether the token is renewable or not.
        """
        return pulumi.get(self, "renewable")

    @property
    @pulumi.getter(name="roleId")
    def role_id(self) -> pulumi.Output[builtins.str]:
        """
        The ID of the role to log in with.
        """
        return pulumi.get(self, "role_id")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The secret ID of the role to log in with. Required
        unless `bind_secret_id` is set to false on the role.
        """
        return pulumi.get(self, "secret_id")

