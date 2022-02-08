# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['MountArgs', 'Mount']

@pulumi.input_type
class MountArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 type: pulumi.Input[str],
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_entropy_access: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 seal_wrap: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Mount resource.
        :param pulumi.Input[str] path: Where the secret backend will be mounted
        :param pulumi.Input[str] type: Type of the backend, such as "aws"
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[int] default_lease_ttl_seconds: Default lease duration for tokens and secrets in seconds
        :param pulumi.Input[str] description: Human-friendly description of the mount
        :param pulumi.Input[bool] external_entropy_access: Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        :param pulumi.Input[bool] local: Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        :param pulumi.Input[int] max_lease_ttl_seconds: Maximum possible lease duration for tokens and secrets in seconds
        :param pulumi.Input[Mapping[str, Any]] options: Specifies mount type specific options that are passed to the backend
        :param pulumi.Input[bool] seal_wrap: Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "type", type)
        if audit_non_hmac_request_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_request_keys", audit_non_hmac_request_keys)
        if audit_non_hmac_response_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_response_keys", audit_non_hmac_response_keys)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_entropy_access is not None:
            pulumi.set(__self__, "external_entropy_access", external_entropy_access)
        if local is not None:
            pulumi.set(__self__, "local", local)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if seal_wrap is not None:
            pulumi.set(__self__, "seal_wrap", seal_wrap)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        Where the secret backend will be mounted
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Type of the backend, such as "aws"
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @audit_non_hmac_request_keys.setter
    def audit_non_hmac_request_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_request_keys", value)

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @audit_non_hmac_response_keys.setter
    def audit_non_hmac_response_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_response_keys", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Default lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description of the mount
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalEntropyAccess")
    def external_entropy_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        """
        return pulumi.get(self, "external_entropy_access")

    @external_entropy_access.setter
    def external_entropy_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_entropy_access", value)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum possible lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Specifies mount type specific options that are passed to the backend
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="sealWrap")
    def seal_wrap(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        """
        return pulumi.get(self, "seal_wrap")

    @seal_wrap.setter
    def seal_wrap(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "seal_wrap", value)


@pulumi.input_type
class _MountState:
    def __init__(__self__, *,
                 accessor: Optional[pulumi.Input[str]] = None,
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_entropy_access: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 seal_wrap: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Mount resources.
        :param pulumi.Input[str] accessor: The accessor for this mount.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[int] default_lease_ttl_seconds: Default lease duration for tokens and secrets in seconds
        :param pulumi.Input[str] description: Human-friendly description of the mount
        :param pulumi.Input[bool] external_entropy_access: Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        :param pulumi.Input[bool] local: Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        :param pulumi.Input[int] max_lease_ttl_seconds: Maximum possible lease duration for tokens and secrets in seconds
        :param pulumi.Input[Mapping[str, Any]] options: Specifies mount type specific options that are passed to the backend
        :param pulumi.Input[str] path: Where the secret backend will be mounted
        :param pulumi.Input[bool] seal_wrap: Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        :param pulumi.Input[str] type: Type of the backend, such as "aws"
        """
        if accessor is not None:
            pulumi.set(__self__, "accessor", accessor)
        if audit_non_hmac_request_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_request_keys", audit_non_hmac_request_keys)
        if audit_non_hmac_response_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_response_keys", audit_non_hmac_response_keys)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_entropy_access is not None:
            pulumi.set(__self__, "external_entropy_access", external_entropy_access)
        if local is not None:
            pulumi.set(__self__, "local", local)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if seal_wrap is not None:
            pulumi.set(__self__, "seal_wrap", seal_wrap)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def accessor(self) -> Optional[pulumi.Input[str]]:
        """
        The accessor for this mount.
        """
        return pulumi.get(self, "accessor")

    @accessor.setter
    def accessor(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "accessor", value)

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @audit_non_hmac_request_keys.setter
    def audit_non_hmac_request_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_request_keys", value)

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @audit_non_hmac_response_keys.setter
    def audit_non_hmac_response_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_response_keys", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Default lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description of the mount
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalEntropyAccess")
    def external_entropy_access(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        """
        return pulumi.get(self, "external_entropy_access")

    @external_entropy_access.setter
    def external_entropy_access(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "external_entropy_access", value)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum possible lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Specifies mount type specific options that are passed to the backend
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Where the secret backend will be mounted
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="sealWrap")
    def seal_wrap(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        """
        return pulumi.get(self, "seal_wrap")

    @seal_wrap.setter
    def seal_wrap(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "seal_wrap", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of the backend, such as "aws"
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Mount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_entropy_access: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 seal_wrap: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.Mount("example",
            description="This is an example mount",
            path="dummy",
            type="generic")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        kvv2_example = vault.Mount("kvv2-example",
            description="This is an example KV Version 2 secret engine mount",
            path="version2-example",
            type="kv-v2")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        transit_example = vault.Mount("transit-example",
            description="This is an example transit secret engine mount",
            options={
                "convergent_encryption": False,
            },
            path="transit-example",
            type="transit")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        pki_example = vault.Mount("pki-example",
            default_lease_ttl_seconds=3600,
            description="This is an example PKI mount",
            max_lease_ttl_seconds=86400,
            path="pki-example",
            type="pki")
        ```

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:index/mount:Mount example dummy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[int] default_lease_ttl_seconds: Default lease duration for tokens and secrets in seconds
        :param pulumi.Input[str] description: Human-friendly description of the mount
        :param pulumi.Input[bool] external_entropy_access: Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        :param pulumi.Input[bool] local: Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        :param pulumi.Input[int] max_lease_ttl_seconds: Maximum possible lease duration for tokens and secrets in seconds
        :param pulumi.Input[Mapping[str, Any]] options: Specifies mount type specific options that are passed to the backend
        :param pulumi.Input[str] path: Where the secret backend will be mounted
        :param pulumi.Input[bool] seal_wrap: Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        :param pulumi.Input[str] type: Type of the backend, such as "aws"
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.Mount("example",
            description="This is an example mount",
            path="dummy",
            type="generic")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        kvv2_example = vault.Mount("kvv2-example",
            description="This is an example KV Version 2 secret engine mount",
            path="version2-example",
            type="kv-v2")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        transit_example = vault.Mount("transit-example",
            description="This is an example transit secret engine mount",
            options={
                "convergent_encryption": False,
            },
            path="transit-example",
            type="transit")
        ```

        ```python
        import pulumi
        import pulumi_vault as vault

        pki_example = vault.Mount("pki-example",
            default_lease_ttl_seconds=3600,
            description="This is an example PKI mount",
            max_lease_ttl_seconds=86400,
            path="pki-example",
            type="pki")
        ```

        ## Import

        Mounts can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:index/mount:Mount example dummy
        ```

        :param str resource_name: The name of the resource.
        :param MountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_entropy_access: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 seal_wrap: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MountArgs.__new__(MountArgs)

            __props__.__dict__["audit_non_hmac_request_keys"] = audit_non_hmac_request_keys
            __props__.__dict__["audit_non_hmac_response_keys"] = audit_non_hmac_response_keys
            __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
            __props__.__dict__["description"] = description
            __props__.__dict__["external_entropy_access"] = external_entropy_access
            __props__.__dict__["local"] = local
            __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
            __props__.__dict__["options"] = options
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["seal_wrap"] = seal_wrap
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["accessor"] = None
        super(Mount, __self__).__init__(
            'vault:index/mount:Mount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessor: Optional[pulumi.Input[str]] = None,
            audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_entropy_access: Optional[pulumi.Input[bool]] = None,
            local: Optional[pulumi.Input[bool]] = None,
            max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            path: Optional[pulumi.Input[str]] = None,
            seal_wrap: Optional[pulumi.Input[bool]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Mount':
        """
        Get an existing Mount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor for this mount.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[int] default_lease_ttl_seconds: Default lease duration for tokens and secrets in seconds
        :param pulumi.Input[str] description: Human-friendly description of the mount
        :param pulumi.Input[bool] external_entropy_access: Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        :param pulumi.Input[bool] local: Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        :param pulumi.Input[int] max_lease_ttl_seconds: Maximum possible lease duration for tokens and secrets in seconds
        :param pulumi.Input[Mapping[str, Any]] options: Specifies mount type specific options that are passed to the backend
        :param pulumi.Input[str] path: Where the secret backend will be mounted
        :param pulumi.Input[bool] seal_wrap: Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        :param pulumi.Input[str] type: Type of the backend, such as "aws"
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MountState.__new__(_MountState)

        __props__.__dict__["accessor"] = accessor
        __props__.__dict__["audit_non_hmac_request_keys"] = audit_non_hmac_request_keys
        __props__.__dict__["audit_non_hmac_response_keys"] = audit_non_hmac_response_keys
        __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
        __props__.__dict__["description"] = description
        __props__.__dict__["external_entropy_access"] = external_entropy_access
        __props__.__dict__["local"] = local
        __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
        __props__.__dict__["options"] = options
        __props__.__dict__["path"] = path
        __props__.__dict__["seal_wrap"] = seal_wrap
        __props__.__dict__["type"] = type
        return Mount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessor(self) -> pulumi.Output[str]:
        """
        The accessor for this mount.
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the list of keys that will not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        Default lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-friendly description of the mount
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalEntropyAccess")
    def external_entropy_access(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag that can be explicitly set to true to enable the secrets engine to access Vault's external entropy source
        """
        return pulumi.get(self, "external_entropy_access")

    @property
    @pulumi.getter
    def local(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean flag that can be explicitly set to true to enforce local mount in HA environment
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        Maximum possible lease duration for tokens and secrets in seconds
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Specifies mount type specific options that are passed to the backend
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        Where the secret backend will be mounted
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="sealWrap")
    def seal_wrap(self) -> pulumi.Output[bool]:
        """
        Boolean flag that can be explicitly set to true to enable seal wrapping for the mount, causing values stored by the mount to be wrapped by the seal's encryption capability
        """
        return pulumi.get(self, "seal_wrap")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of the backend, such as "aws"
        """
        return pulumi.get(self, "type")

