# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthBackendRoleSecretIdArgs', 'AuthBackendRoleSecretId']

@pulumi.input_type
class AuthBackendRoleSecretIdArgs:
    def __init__(__self__, *,
                 role_name: pulumi.Input[str],
                 backend: Optional[pulumi.Input[str]] = None,
                 cidr_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 with_wrapped_accessor: Optional[pulumi.Input[bool]] = None,
                 wrapping_ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthBackendRoleSecretId resource.
        :param pulumi.Input[str] role_name: The name of the role to create the SecretID for.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_lists: If set, specifies blocks of IP addresses which can
               perform the login operation using this SecretID.
        :param pulumi.Input[str] metadata: A JSON-encoded string containing metadata in
               key-value pairs to be set on tokens issued with this SecretID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] secret_id: The SecretID to be created. If set, uses "Push"
               mode.  Defaults to Vault auto-generating SecretIDs.
        :param pulumi.Input[bool] with_wrapped_accessor: Set to `true` to use the wrapped secret-id accessor as the resource ID.
               If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
               invalidated through unwrapping.
        :param pulumi.Input[str] wrapping_ttl: If set, the SecretID response will be
               [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
               and available for the duration specified. Only a single unwrapping of the
               token is allowed.
        """
        pulumi.set(__self__, "role_name", role_name)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if cidr_lists is not None:
            pulumi.set(__self__, "cidr_lists", cidr_lists)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if with_wrapped_accessor is not None:
            pulumi.set(__self__, "with_wrapped_accessor", with_wrapped_accessor)
        if wrapping_ttl is not None:
            pulumi.set(__self__, "wrapping_ttl", wrapping_ttl)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        The name of the role to create the SecretID for.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="cidrLists")
    def cidr_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If set, specifies blocks of IP addresses which can
        perform the login operation using this SecretID.
        """
        return pulumi.get(self, "cidr_lists")

    @cidr_lists.setter
    def cidr_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_lists", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON-encoded string containing metadata in
        key-value pairs to be set on tokens issued with this SecretID.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        The SecretID to be created. If set, uses "Push"
        mode.  Defaults to Vault auto-generating SecretIDs.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="withWrappedAccessor")
    def with_wrapped_accessor(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to use the wrapped secret-id accessor as the resource ID.
        If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
        invalidated through unwrapping.
        """
        return pulumi.get(self, "with_wrapped_accessor")

    @with_wrapped_accessor.setter
    def with_wrapped_accessor(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "with_wrapped_accessor", value)

    @property
    @pulumi.getter(name="wrappingTtl")
    def wrapping_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the SecretID response will be
        [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
        and available for the duration specified. Only a single unwrapping of the
        token is allowed.
        """
        return pulumi.get(self, "wrapping_ttl")

    @wrapping_ttl.setter
    def wrapping_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_ttl", value)


@pulumi.input_type
class _AuthBackendRoleSecretIdState:
    def __init__(__self__, *,
                 accessor: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cidr_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 with_wrapped_accessor: Optional[pulumi.Input[bool]] = None,
                 wrapping_accessor: Optional[pulumi.Input[str]] = None,
                 wrapping_token: Optional[pulumi.Input[str]] = None,
                 wrapping_ttl: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthBackendRoleSecretId resources.
        :param pulumi.Input[str] accessor: The unique ID for this SecretID that can be safely logged.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_lists: If set, specifies blocks of IP addresses which can
               perform the login operation using this SecretID.
        :param pulumi.Input[str] metadata: A JSON-encoded string containing metadata in
               key-value pairs to be set on tokens issued with this SecretID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: The name of the role to create the SecretID for.
        :param pulumi.Input[str] secret_id: The SecretID to be created. If set, uses "Push"
               mode.  Defaults to Vault auto-generating SecretIDs.
        :param pulumi.Input[bool] with_wrapped_accessor: Set to `true` to use the wrapped secret-id accessor as the resource ID.
               If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
               invalidated through unwrapping.
        :param pulumi.Input[str] wrapping_accessor: The unique ID for the response-wrapped SecretID that can
               be safely logged.
        :param pulumi.Input[str] wrapping_token: The token used to retrieve a response-wrapped SecretID.
        :param pulumi.Input[str] wrapping_ttl: If set, the SecretID response will be
               [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
               and available for the duration specified. Only a single unwrapping of the
               token is allowed.
        """
        if accessor is not None:
            pulumi.set(__self__, "accessor", accessor)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if cidr_lists is not None:
            pulumi.set(__self__, "cidr_lists", cidr_lists)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if secret_id is not None:
            pulumi.set(__self__, "secret_id", secret_id)
        if with_wrapped_accessor is not None:
            pulumi.set(__self__, "with_wrapped_accessor", with_wrapped_accessor)
        if wrapping_accessor is not None:
            pulumi.set(__self__, "wrapping_accessor", wrapping_accessor)
        if wrapping_token is not None:
            pulumi.set(__self__, "wrapping_token", wrapping_token)
        if wrapping_ttl is not None:
            pulumi.set(__self__, "wrapping_ttl", wrapping_ttl)

    @property
    @pulumi.getter
    def accessor(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for this SecretID that can be safely logged.
        """
        return pulumi.get(self, "accessor")

    @accessor.setter
    def accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accessor", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="cidrLists")
    def cidr_lists(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        If set, specifies blocks of IP addresses which can
        perform the login operation using this SecretID.
        """
        return pulumi.get(self, "cidr_lists")

    @cidr_lists.setter
    def cidr_lists(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_lists", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[str]]:
        """
        A JSON-encoded string containing metadata in
        key-value pairs to be set on tokens issued with this SecretID.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role to create the SecretID for.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> Optional[pulumi.Input[str]]:
        """
        The SecretID to be created. If set, uses "Push"
        mode.  Defaults to Vault auto-generating SecretIDs.
        """
        return pulumi.get(self, "secret_id")

    @secret_id.setter
    def secret_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_id", value)

    @property
    @pulumi.getter(name="withWrappedAccessor")
    def with_wrapped_accessor(self) -> Optional[pulumi.Input[bool]]:
        """
        Set to `true` to use the wrapped secret-id accessor as the resource ID.
        If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
        invalidated through unwrapping.
        """
        return pulumi.get(self, "with_wrapped_accessor")

    @with_wrapped_accessor.setter
    def with_wrapped_accessor(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "with_wrapped_accessor", value)

    @property
    @pulumi.getter(name="wrappingAccessor")
    def wrapping_accessor(self) -> Optional[pulumi.Input[str]]:
        """
        The unique ID for the response-wrapped SecretID that can
        be safely logged.
        """
        return pulumi.get(self, "wrapping_accessor")

    @wrapping_accessor.setter
    def wrapping_accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_accessor", value)

    @property
    @pulumi.getter(name="wrappingToken")
    def wrapping_token(self) -> Optional[pulumi.Input[str]]:
        """
        The token used to retrieve a response-wrapped SecretID.
        """
        return pulumi.get(self, "wrapping_token")

    @wrapping_token.setter
    def wrapping_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_token", value)

    @property
    @pulumi.getter(name="wrappingTtl")
    def wrapping_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        If set, the SecretID response will be
        [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
        and available for the duration specified. Only a single unwrapping of the
        token is allowed.
        """
        return pulumi.get(self, "wrapping_ttl")

    @wrapping_ttl.setter
    def wrapping_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "wrapping_ttl", value)


class AuthBackendRoleSecretId(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cidr_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 with_wrapped_accessor: Optional[pulumi.Input[bool]] = None,
                 wrapping_ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/approle) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
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
            role_name=example.role_name,
            metadata=json.dumps({
                "hello": "world",
            }))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_lists: If set, specifies blocks of IP addresses which can
               perform the login operation using this SecretID.
        :param pulumi.Input[str] metadata: A JSON-encoded string containing metadata in
               key-value pairs to be set on tokens issued with this SecretID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: The name of the role to create the SecretID for.
        :param pulumi.Input[str] secret_id: The SecretID to be created. If set, uses "Push"
               mode.  Defaults to Vault auto-generating SecretIDs.
        :param pulumi.Input[bool] with_wrapped_accessor: Set to `true` to use the wrapped secret-id accessor as the resource ID.
               If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
               invalidated through unwrapping.
        :param pulumi.Input[str] wrapping_ttl: If set, the SecretID response will be
               [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
               and available for the duration specified. Only a single unwrapping of the
               token is allowed.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendRoleSecretIdArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an AppRole auth backend SecretID in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/approle) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
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
            role_name=example.role_name,
            metadata=json.dumps({
                "hello": "world",
            }))
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param AuthBackendRoleSecretIdArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendRoleSecretIdArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cidr_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 metadata: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 secret_id: Optional[pulumi.Input[str]] = None,
                 with_wrapped_accessor: Optional[pulumi.Input[bool]] = None,
                 wrapping_ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendRoleSecretIdArgs.__new__(AuthBackendRoleSecretIdArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["cidr_lists"] = cidr_lists
            __props__.__dict__["metadata"] = metadata
            __props__.__dict__["namespace"] = namespace
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            __props__.__dict__["secret_id"] = None if secret_id is None else pulumi.Output.secret(secret_id)
            __props__.__dict__["with_wrapped_accessor"] = with_wrapped_accessor
            __props__.__dict__["wrapping_ttl"] = wrapping_ttl
            __props__.__dict__["accessor"] = None
            __props__.__dict__["wrapping_accessor"] = None
            __props__.__dict__["wrapping_token"] = None
        alias_opts = pulumi.ResourceOptions(aliases=[pulumi.Alias(type_="vault:appRole/authBackendRoleSecretID:AuthBackendRoleSecretID")])
        opts = pulumi.ResourceOptions.merge(opts, alias_opts)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["secretId", "wrappingToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AuthBackendRoleSecretId, __self__).__init__(
            'vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessor: Optional[pulumi.Input[str]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            cidr_lists: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            metadata: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            secret_id: Optional[pulumi.Input[str]] = None,
            with_wrapped_accessor: Optional[pulumi.Input[bool]] = None,
            wrapping_accessor: Optional[pulumi.Input[str]] = None,
            wrapping_token: Optional[pulumi.Input[str]] = None,
            wrapping_ttl: Optional[pulumi.Input[str]] = None) -> 'AuthBackendRoleSecretId':
        """
        Get an existing AuthBackendRoleSecretId resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The unique ID for this SecretID that can be safely logged.
        :param pulumi.Input[str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_lists: If set, specifies blocks of IP addresses which can
               perform the login operation using this SecretID.
        :param pulumi.Input[str] metadata: A JSON-encoded string containing metadata in
               key-value pairs to be set on tokens issued with this SecretID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: The name of the role to create the SecretID for.
        :param pulumi.Input[str] secret_id: The SecretID to be created. If set, uses "Push"
               mode.  Defaults to Vault auto-generating SecretIDs.
        :param pulumi.Input[bool] with_wrapped_accessor: Set to `true` to use the wrapped secret-id accessor as the resource ID.
               If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
               invalidated through unwrapping.
        :param pulumi.Input[str] wrapping_accessor: The unique ID for the response-wrapped SecretID that can
               be safely logged.
        :param pulumi.Input[str] wrapping_token: The token used to retrieve a response-wrapped SecretID.
        :param pulumi.Input[str] wrapping_ttl: If set, the SecretID response will be
               [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
               and available for the duration specified. Only a single unwrapping of the
               token is allowed.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendRoleSecretIdState.__new__(_AuthBackendRoleSecretIdState)

        __props__.__dict__["accessor"] = accessor
        __props__.__dict__["backend"] = backend
        __props__.__dict__["cidr_lists"] = cidr_lists
        __props__.__dict__["metadata"] = metadata
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["secret_id"] = secret_id
        __props__.__dict__["with_wrapped_accessor"] = with_wrapped_accessor
        __props__.__dict__["wrapping_accessor"] = wrapping_accessor
        __props__.__dict__["wrapping_token"] = wrapping_token
        __props__.__dict__["wrapping_ttl"] = wrapping_ttl
        return AuthBackendRoleSecretId(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessor(self) -> pulumi.Output[str]:
        """
        The unique ID for this SecretID that can be safely logged.
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="cidrLists")
    def cidr_lists(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        If set, specifies blocks of IP addresses which can
        perform the login operation using this SecretID.
        """
        return pulumi.get(self, "cidr_lists")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Optional[str]]:
        """
        A JSON-encoded string containing metadata in
        key-value pairs to be set on tokens issued with this SecretID.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Output[str]:
        """
        The name of the role to create the SecretID for.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="secretId")
    def secret_id(self) -> pulumi.Output[str]:
        """
        The SecretID to be created. If set, uses "Push"
        mode.  Defaults to Vault auto-generating SecretIDs.
        """
        return pulumi.get(self, "secret_id")

    @property
    @pulumi.getter(name="withWrappedAccessor")
    def with_wrapped_accessor(self) -> pulumi.Output[Optional[bool]]:
        """
        Set to `true` to use the wrapped secret-id accessor as the resource ID.
        If `false` (default value), a fresh secret ID will be regenerated whenever the wrapping token is expired or
        invalidated through unwrapping.
        """
        return pulumi.get(self, "with_wrapped_accessor")

    @property
    @pulumi.getter(name="wrappingAccessor")
    def wrapping_accessor(self) -> pulumi.Output[str]:
        """
        The unique ID for the response-wrapped SecretID that can
        be safely logged.
        """
        return pulumi.get(self, "wrapping_accessor")

    @property
    @pulumi.getter(name="wrappingToken")
    def wrapping_token(self) -> pulumi.Output[str]:
        """
        The token used to retrieve a response-wrapped SecretID.
        """
        return pulumi.get(self, "wrapping_token")

    @property
    @pulumi.getter(name="wrappingTtl")
    def wrapping_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        If set, the SecretID response will be
        [response-wrapped](https://www.vaultproject.io/docs/concepts/response-wrapping)
        and available for the duration specified. Only a single unwrapping of the
        token is allowed.
        """
        return pulumi.get(self, "wrapping_ttl")

