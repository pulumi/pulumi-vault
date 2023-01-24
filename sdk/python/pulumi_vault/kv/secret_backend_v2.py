# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendV2Args', 'SecretBackendV2']

@pulumi.input_type
class SecretBackendV2Args:
    def __init__(__self__, *,
                 mount: pulumi.Input[str],
                 cas_required: Optional[pulumi.Input[bool]] = None,
                 delete_version_after: Optional[pulumi.Input[int]] = None,
                 max_versions: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackendV2 resource.
        :param pulumi.Input[str] mount: Path where KV-V2 engine is mounted.
        :param pulumi.Input[bool] cas_required: If true, all keys will require the cas
               parameter to be set on all write requests.
        :param pulumi.Input[int] delete_version_after: If set, specifies the length of time before
               a version is deleted. Accepts duration in integer seconds.
        :param pulumi.Input[int] max_versions: The number of versions to keep per key.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "mount", mount)
        if cas_required is not None:
            pulumi.set(__self__, "cas_required", cas_required)
        if delete_version_after is not None:
            pulumi.set(__self__, "delete_version_after", delete_version_after)
        if max_versions is not None:
            pulumi.set(__self__, "max_versions", max_versions)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Input[str]:
        """
        Path where KV-V2 engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter(name="casRequired")
    def cas_required(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, all keys will require the cas
        parameter to be set on all write requests.
        """
        return pulumi.get(self, "cas_required")

    @cas_required.setter
    def cas_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cas_required", value)

    @property
    @pulumi.getter(name="deleteVersionAfter")
    def delete_version_after(self) -> Optional[pulumi.Input[int]]:
        """
        If set, specifies the length of time before
        a version is deleted. Accepts duration in integer seconds.
        """
        return pulumi.get(self, "delete_version_after")

    @delete_version_after.setter
    def delete_version_after(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_version_after", value)

    @property
    @pulumi.getter(name="maxVersions")
    def max_versions(self) -> Optional[pulumi.Input[int]]:
        """
        The number of versions to keep per key.
        """
        return pulumi.get(self, "max_versions")

    @max_versions.setter
    def max_versions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_versions", value)

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


@pulumi.input_type
class _SecretBackendV2State:
    def __init__(__self__, *,
                 cas_required: Optional[pulumi.Input[bool]] = None,
                 delete_version_after: Optional[pulumi.Input[int]] = None,
                 max_versions: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackendV2 resources.
        :param pulumi.Input[bool] cas_required: If true, all keys will require the cas
               parameter to be set on all write requests.
        :param pulumi.Input[int] delete_version_after: If set, specifies the length of time before
               a version is deleted. Accepts duration in integer seconds.
        :param pulumi.Input[int] max_versions: The number of versions to keep per key.
        :param pulumi.Input[str] mount: Path where KV-V2 engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        if cas_required is not None:
            pulumi.set(__self__, "cas_required", cas_required)
        if delete_version_after is not None:
            pulumi.set(__self__, "delete_version_after", delete_version_after)
        if max_versions is not None:
            pulumi.set(__self__, "max_versions", max_versions)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="casRequired")
    def cas_required(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, all keys will require the cas
        parameter to be set on all write requests.
        """
        return pulumi.get(self, "cas_required")

    @cas_required.setter
    def cas_required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "cas_required", value)

    @property
    @pulumi.getter(name="deleteVersionAfter")
    def delete_version_after(self) -> Optional[pulumi.Input[int]]:
        """
        If set, specifies the length of time before
        a version is deleted. Accepts duration in integer seconds.
        """
        return pulumi.get(self, "delete_version_after")

    @delete_version_after.setter
    def delete_version_after(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "delete_version_after", value)

    @property
    @pulumi.getter(name="maxVersions")
    def max_versions(self) -> Optional[pulumi.Input[int]]:
        """
        The number of versions to keep per key.
        """
        return pulumi.get(self, "max_versions")

    @max_versions.setter
    def max_versions(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_versions", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        Path where KV-V2 engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

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


class SecretBackendV2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cas_required: Optional[pulumi.Input[bool]] = None,
                 delete_version_after: Optional[pulumi.Input[int]] = None,
                 max_versions: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configures KV-V2 backend level settings that are applied to
        every key in the key-value store.

        For more information on Vault's KV-V2 secret backend
        [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        kvv2 = vault.Mount("kvv2",
            path="kvv2",
            type="kv",
            options={
                "version": "2",
            },
            description="KV Version 2 secret engine mount")
        example = vault.kv.SecretBackendV2("example",
            mount=kvv2.path,
            max_versions=5,
            delete_version_after=12600,
            cas_required=True)
        ```
        ## Required Vault Capabilities

        Use of this resource requires the `create` or `update` capability
        (depending on whether the resource already exists) on the given path,
        the `delete` capability if the resource is removed from configuration,
        and the `read` capability for drift detection (by default).

        ## Import

        The KV-V2 secret backend can be imported using its unique ID, the `${mount}/config`, e.g.

        ```sh
         $ pulumi import vault:kv/secretBackendV2:SecretBackendV2 example kvv2/config
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cas_required: If true, all keys will require the cas
               parameter to be set on all write requests.
        :param pulumi.Input[int] delete_version_after: If set, specifies the length of time before
               a version is deleted. Accepts duration in integer seconds.
        :param pulumi.Input[int] max_versions: The number of versions to keep per key.
        :param pulumi.Input[str] mount: Path where KV-V2 engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendV2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configures KV-V2 backend level settings that are applied to
        every key in the key-value store.

        For more information on Vault's KV-V2 secret backend
        [see here](https://www.vaultproject.io/docs/secrets/kv/kv-v2).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        kvv2 = vault.Mount("kvv2",
            path="kvv2",
            type="kv",
            options={
                "version": "2",
            },
            description="KV Version 2 secret engine mount")
        example = vault.kv.SecretBackendV2("example",
            mount=kvv2.path,
            max_versions=5,
            delete_version_after=12600,
            cas_required=True)
        ```
        ## Required Vault Capabilities

        Use of this resource requires the `create` or `update` capability
        (depending on whether the resource already exists) on the given path,
        the `delete` capability if the resource is removed from configuration,
        and the `read` capability for drift detection (by default).

        ## Import

        The KV-V2 secret backend can be imported using its unique ID, the `${mount}/config`, e.g.

        ```sh
         $ pulumi import vault:kv/secretBackendV2:SecretBackendV2 example kvv2/config
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendV2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendV2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cas_required: Optional[pulumi.Input[bool]] = None,
                 delete_version_after: Optional[pulumi.Input[int]] = None,
                 max_versions: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendV2Args.__new__(SecretBackendV2Args)

            __props__.__dict__["cas_required"] = cas_required
            __props__.__dict__["delete_version_after"] = delete_version_after
            __props__.__dict__["max_versions"] = max_versions
            if mount is None and not opts.urn:
                raise TypeError("Missing required property 'mount'")
            __props__.__dict__["mount"] = mount
            __props__.__dict__["namespace"] = namespace
        super(SecretBackendV2, __self__).__init__(
            'vault:kv/secretBackendV2:SecretBackendV2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cas_required: Optional[pulumi.Input[bool]] = None,
            delete_version_after: Optional[pulumi.Input[int]] = None,
            max_versions: Optional[pulumi.Input[int]] = None,
            mount: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None) -> 'SecretBackendV2':
        """
        Get an existing SecretBackendV2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] cas_required: If true, all keys will require the cas
               parameter to be set on all write requests.
        :param pulumi.Input[int] delete_version_after: If set, specifies the length of time before
               a version is deleted. Accepts duration in integer seconds.
        :param pulumi.Input[int] max_versions: The number of versions to keep per key.
        :param pulumi.Input[str] mount: Path where KV-V2 engine is mounted.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendV2State.__new__(_SecretBackendV2State)

        __props__.__dict__["cas_required"] = cas_required
        __props__.__dict__["delete_version_after"] = delete_version_after
        __props__.__dict__["max_versions"] = max_versions
        __props__.__dict__["mount"] = mount
        __props__.__dict__["namespace"] = namespace
        return SecretBackendV2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="casRequired")
    def cas_required(self) -> pulumi.Output[bool]:
        """
        If true, all keys will require the cas
        parameter to be set on all write requests.
        """
        return pulumi.get(self, "cas_required")

    @property
    @pulumi.getter(name="deleteVersionAfter")
    def delete_version_after(self) -> pulumi.Output[Optional[int]]:
        """
        If set, specifies the length of time before
        a version is deleted. Accepts duration in integer seconds.
        """
        return pulumi.get(self, "delete_version_after")

    @property
    @pulumi.getter(name="maxVersions")
    def max_versions(self) -> pulumi.Output[int]:
        """
        The number of versions to keep per key.
        """
        return pulumi.get(self, "max_versions")

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[str]:
        """
        Path where KV-V2 engine is mounted.
        """
        return pulumi.get(self, "mount")

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

