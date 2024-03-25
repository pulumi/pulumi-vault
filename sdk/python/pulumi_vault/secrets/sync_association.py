# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncAssociationArgs', 'SyncAssociation']

@pulumi.input_type
class SyncAssociationArgs:
    def __init__(__self__, *,
                 mount: pulumi.Input[str],
                 secret_name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SyncAssociation resource.
        :param pulumi.Input[str] mount: Specifies the mount where the secret is located.
        :param pulumi.Input[str] secret_name: Specifies the name of the secret to synchronize.
        :param pulumi.Input[str] type: Specifies the destination type.
        :param pulumi.Input[str] name: Specifies the name of the destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        pulumi.set(__self__, "mount", mount)
        pulumi.set(__self__, "secret_name", secret_name)
        pulumi.set(__self__, "type", type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Input[str]:
        """
        Specifies the mount where the secret is located.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Input[str]:
        """
        Specifies the name of the secret to synchronize.
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the destination type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the destination.
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
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _SyncAssociationState:
    def __init__(__self__, *,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 sync_status: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 updated_at: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyncAssociation resources.
        :param pulumi.Input[str] mount: Specifies the mount where the secret is located.
        :param pulumi.Input[str] name: Specifies the name of the destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name: Specifies the name of the secret to synchronize.
        :param pulumi.Input[str] sync_status: Specifies the status of the association (for eg. `SYNCED`).
        :param pulumi.Input[str] type: Specifies the destination type.
        :param pulumi.Input[str] updated_at: Duration string specifying when the secret was last updated.
        """
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_name is not None:
            pulumi.set(__self__, "secret_name", secret_name)
        if sync_status is not None:
            pulumi.set(__self__, "sync_status", sync_status)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if updated_at is not None:
            pulumi.set(__self__, "updated_at", updated_at)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the mount where the secret is located.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the destination.
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
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the name of the secret to synchronize.
        """
        return pulumi.get(self, "secret_name")

    @secret_name.setter
    def secret_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name", value)

    @property
    @pulumi.getter(name="syncStatus")
    def sync_status(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the status of the association (for eg. `SYNCED`).
        """
        return pulumi.get(self, "sync_status")

    @sync_status.setter
    def sync_status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sync_status", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the destination type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> Optional[pulumi.Input[str]]:
        """
        Duration string specifying when the secret was last updated.
        """
        return pulumi.get(self, "updated_at")

    @updated_at.setter
    def updated_at(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "updated_at", value)


class SyncAssociation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import pulumi_vault as vault

        kvv2 = vault.Mount("kvv2",
            path="kvv2",
            type="kv",
            options={
                "version": "2",
            },
            description="KV Version 2 secret engine mount")
        token = vault.kv.SecretV2("token",
            mount=kvv2.path,
            data_json=json.dumps({
                "dev": "B!gS3cr3t",
                "prod": "S3cureP4$$",
            }))
        gh = vault.secrets.SyncGhDestination("gh",
            access_token=var["access_token"],
            repository_owner=var["repo_owner"],
            repository_name="repo-name-example",
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}")
        gh_token = vault.secrets.SyncAssociation("ghToken",
            type=gh.type,
            mount=kvv2.path,
            secret_name=token.name)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mount: Specifies the mount where the secret is located.
        :param pulumi.Input[str] name: Specifies the name of the destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name: Specifies the name of the secret to synchronize.
        :param pulumi.Input[str] type: Specifies the destination type.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncAssociationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import json
        import pulumi_vault as vault

        kvv2 = vault.Mount("kvv2",
            path="kvv2",
            type="kv",
            options={
                "version": "2",
            },
            description="KV Version 2 secret engine mount")
        token = vault.kv.SecretV2("token",
            mount=kvv2.path,
            data_json=json.dumps({
                "dev": "B!gS3cr3t",
                "prod": "S3cureP4$$",
            }))
        gh = vault.secrets.SyncGhDestination("gh",
            access_token=var["access_token"],
            repository_owner=var["repo_owner"],
            repository_name="repo-name-example",
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}")
        gh_token = vault.secrets.SyncAssociation("ghToken",
            type=gh.type,
            mount=kvv2.path,
            secret_name=token.name)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param SyncAssociationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncAssociationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncAssociationArgs.__new__(SyncAssociationArgs)

            if mount is None and not opts.urn:
                raise TypeError("Missing required property 'mount'")
            __props__.__dict__["mount"] = mount
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if secret_name is None and not opts.urn:
                raise TypeError("Missing required property 'secret_name'")
            __props__.__dict__["secret_name"] = secret_name
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["sync_status"] = None
            __props__.__dict__["updated_at"] = None
        super(SyncAssociation, __self__).__init__(
            'vault:secrets/syncAssociation:SyncAssociation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            mount: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            secret_name: Optional[pulumi.Input[str]] = None,
            sync_status: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            updated_at: Optional[pulumi.Input[str]] = None) -> 'SyncAssociation':
        """
        Get an existing SyncAssociation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] mount: Specifies the mount where the secret is located.
        :param pulumi.Input[str] name: Specifies the name of the destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name: Specifies the name of the secret to synchronize.
        :param pulumi.Input[str] sync_status: Specifies the status of the association (for eg. `SYNCED`).
        :param pulumi.Input[str] type: Specifies the destination type.
        :param pulumi.Input[str] updated_at: Duration string specifying when the secret was last updated.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncAssociationState.__new__(_SyncAssociationState)

        __props__.__dict__["mount"] = mount
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["secret_name"] = secret_name
        __props__.__dict__["sync_status"] = sync_status
        __props__.__dict__["type"] = type
        __props__.__dict__["updated_at"] = updated_at
        return SyncAssociation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[str]:
        """
        Specifies the mount where the secret is located.
        """
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the destination.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="secretName")
    def secret_name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the secret to synchronize.
        """
        return pulumi.get(self, "secret_name")

    @property
    @pulumi.getter(name="syncStatus")
    def sync_status(self) -> pulumi.Output[str]:
        """
        Specifies the status of the association (for eg. `SYNCED`).
        """
        return pulumi.get(self, "sync_status")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the destination type.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updatedAt")
    def updated_at(self) -> pulumi.Output[str]:
        """
        Duration string specifying when the secret was last updated.
        """
        return pulumi.get(self, "updated_at")

