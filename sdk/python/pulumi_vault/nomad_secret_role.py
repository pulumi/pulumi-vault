# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['NomadSecretRoleArgs', 'NomadSecretRole']

@pulumi.input_type
class NomadSecretRoleArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 role: pulumi.Input[str],
                 global_: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NomadSecretRole resource.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "role", role)
        if global_ is not None:
            pulumi.set(__self__, "global_", global_)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The unique path this backend should be mounted at.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        The name to identify this role within the backend.
        Must be unique within the backend.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="global")
    def global_(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the generated token should be global. Defaults to 
        false.
        """
        return pulumi.get(self, "global_")

    @global_.setter
    def global_(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "global_", value)

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
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of policies attached to the generated token. This setting is only used 
        when `type` is 'client'.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of token to create when using this role. Valid 
        settings are 'client' and 'management'. Defaults to 'client'.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _NomadSecretRoleState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NomadSecretRole resources.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if global_ is not None:
            pulumi.set(__self__, "global_", global_)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="global")
    def global_(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the generated token should be global. Defaults to 
        false.
        """
        return pulumi.get(self, "global_")

    @global_.setter
    def global_(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "global_", value)

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
    @pulumi.getter
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of policies attached to the generated token. This setting is only used 
        when `type` is 'client'.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        The name to identify this role within the backend.
        Must be unique within the backend.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of token to create when using this role. Valid 
        settings are 'client' and 'management'. Defaults to 'client'.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class NomadSecretRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.NomadSecretBackend("config",
            backend="nomad",
            description="test description",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=7200,
            address="https://127.0.0.1:4646",
            token="ae20ceaa-...")
        test = vault.NomadSecretRole("test",
            backend=config.backend,
            role="test",
            type="client",
            policies=["readonly"])
        ```

        ## Import

        Nomad secret role can be imported using the `backend`, e.g.

        ```sh
        $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NomadSecretRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.NomadSecretBackend("config",
            backend="nomad",
            description="test description",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=7200,
            address="https://127.0.0.1:4646",
            token="ae20ceaa-...")
        test = vault.NomadSecretRole("test",
            backend=config.backend,
            role="test",
            type="client",
            policies=["readonly"])
        ```

        ## Import

        Nomad secret role can be imported using the `backend`, e.g.

        ```sh
        $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
        ```

        :param str resource_name: The name of the resource.
        :param NomadSecretRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NomadSecretRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NomadSecretRoleArgs.__new__(NomadSecretRoleArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["global_"] = global_
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["policies"] = policies
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["type"] = type
        super(NomadSecretRole, __self__).__init__(
            'vault:index/nomadSecretRole:NomadSecretRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            global_: Optional[pulumi.Input[bool]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'NomadSecretRole':
        """
        Get an existing NomadSecretRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NomadSecretRoleState.__new__(_NomadSecretRoleState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["global_"] = global_
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["policies"] = policies
        __props__.__dict__["role"] = role
        __props__.__dict__["type"] = type
        return NomadSecretRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The unique path this backend should be mounted at.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="global")
    def global_(self) -> pulumi.Output[bool]:
        """
        Specifies if the generated token should be global. Defaults to 
        false.
        """
        return pulumi.get(self, "global_")

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
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        """
        List of policies attached to the generated token. This setting is only used 
        when `type` is 'client'.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name to identify this role within the backend.
        Must be unique within the backend.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the type of token to create when using this role. Valid 
        settings are 'client' and 'management'. Defaults to 'client'.
        """
        return pulumi.get(self, "type")

