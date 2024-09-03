# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretRoleArgs', 'SecretRole']

@pulumi.input_type
class SecretRoleArgs:
    def __init__(__self__, *,
                 mount: pulumi.Input[str],
                 roles: pulumi.Input[Sequence[pulumi.Input[str]]],
                 cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretRole resource.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: Whitelist entry in CIDR notation to be added for the API key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: IP address to be added to the whitelist for the API key.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of credentials issued using this role.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization_id: Unique identifier for the organization to which the target API Key belongs.
               Required if `project_id` is not set.
        :param pulumi.Input[str] project_id: Unique identifier for the project to which the target API Key belongs.
               Required if `organization_id` is not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_roles: Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        :param pulumi.Input[str] ttl: Duration in seconds after which the issued credential should expire.
        """
        pulumi.set(__self__, "mount", mount)
        pulumi.set(__self__, "roles", roles)
        if cidr_blocks is not None:
            pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if project_roles is not None:
            pulumi.set(__self__, "project_roles", project_roles)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Input[str]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: pulumi.Input[str]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Whitelist entry in CIDR notation to be added for the API key.
        """
        return pulumi.get(self, "cidr_blocks")

    @cidr_blocks.setter
    def cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_blocks", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IP address to be added to the whitelist for the API key.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum allowed lifetime of credentials issued using this role.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_ttl", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role.
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
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the organization to which the target API Key belongs.
        Required if `project_id` is not set.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the project to which the target API Key belongs.
        Required if `organization_id` is not set.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        """
        return pulumi.get(self, "project_roles")

    @project_roles.setter
    def project_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_roles", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Duration in seconds after which the issued credential should expire.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _SecretRoleState:
    def __init__(__self__, *,
                 cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretRole resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: Whitelist entry in CIDR notation to be added for the API key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: IP address to be added to the whitelist for the API key.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of credentials issued using this role.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization_id: Unique identifier for the organization to which the target API Key belongs.
               Required if `project_id` is not set.
        :param pulumi.Input[str] project_id: Unique identifier for the project to which the target API Key belongs.
               Required if `organization_id` is not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_roles: Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        :param pulumi.Input[str] ttl: Duration in seconds after which the issued credential should expire.
        """
        if cidr_blocks is not None:
            pulumi.set(__self__, "cidr_blocks", cidr_blocks)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if organization_id is not None:
            pulumi.set(__self__, "organization_id", organization_id)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if project_roles is not None:
            pulumi.set(__self__, "project_roles", project_roles)
        if roles is not None:
            pulumi.set(__self__, "roles", roles)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Whitelist entry in CIDR notation to be added for the API key.
        """
        return pulumi.get(self, "cidr_blocks")

    @cidr_blocks.setter
    def cidr_blocks(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "cidr_blocks", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        IP address to be added to the whitelist for the API key.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        The maximum allowed lifetime of credentials issued using this role.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_ttl", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the role.
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
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the organization to which the target API Key belongs.
        Required if `project_id` is not set.
        """
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        Unique identifier for the project to which the target API Key belongs.
        Required if `organization_id` is not set.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        """
        return pulumi.get(self, "project_roles")

    @project_roles.setter
    def project_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "project_roles", value)

    @property
    @pulumi.getter
    def roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        """
        return pulumi.get(self, "roles")

    @roles.setter
    def roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "roles", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Duration in seconds after which the issued credential should expire.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


class SecretRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mongo = vault.Mount("mongo",
            path="%s",
            type="mongodbatlas",
            description="MongoDB Atlas secret engine mount")
        config = vault.mongodbatlas.SecretBackend("config",
            mount=mongo.path,
            private_key="privateKey",
            public_key="publicKey")
        role = vault.mongodbatlas.SecretRole("role",
            mount=mongo.path,
            name="tf-test-role",
            organization_id="7cf5a45a9ccf6400e60981b7",
            project_id="5cf5a45a9ccf6400e60981b6",
            roles=["ORG_READ_ONLY"],
            ip_addresses="192.168.1.5, 192.168.1.6",
            cidr_blocks="192.168.1.3/35",
            project_roles=["GROUP_READ_ONLY"],
            ttl="60",
            max_ttl="120")
        ```

        ## Import

        The MongoDB Atlas secret role can be imported using the full path to the role
        of the form: `<mount_path>/roles/<role_name>` e.g.

        ```sh
        $ pulumi import vault:mongodbatlas/secretRole:SecretRole example mongodbatlas/roles/example-role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: Whitelist entry in CIDR notation to be added for the API key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: IP address to be added to the whitelist for the API key.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of credentials issued using this role.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization_id: Unique identifier for the organization to which the target API Key belongs.
               Required if `project_id` is not set.
        :param pulumi.Input[str] project_id: Unique identifier for the project to which the target API Key belongs.
               Required if `organization_id` is not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_roles: Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        :param pulumi.Input[str] ttl: Duration in seconds after which the issued credential should expire.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mongo = vault.Mount("mongo",
            path="%s",
            type="mongodbatlas",
            description="MongoDB Atlas secret engine mount")
        config = vault.mongodbatlas.SecretBackend("config",
            mount=mongo.path,
            private_key="privateKey",
            public_key="publicKey")
        role = vault.mongodbatlas.SecretRole("role",
            mount=mongo.path,
            name="tf-test-role",
            organization_id="7cf5a45a9ccf6400e60981b7",
            project_id="5cf5a45a9ccf6400e60981b6",
            roles=["ORG_READ_ONLY"],
            ip_addresses="192.168.1.5, 192.168.1.6",
            cidr_blocks="192.168.1.3/35",
            project_roles=["GROUP_READ_ONLY"],
            ttl="60",
            max_ttl="120")
        ```

        ## Import

        The MongoDB Atlas secret role can be imported using the full path to the role
        of the form: `<mount_path>/roles/<role_name>` e.g.

        ```sh
        $ pulumi import vault:mongodbatlas/secretRole:SecretRole example mongodbatlas/roles/example-role
        ```

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
                 cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretRoleArgs.__new__(SecretRoleArgs)

            __props__.__dict__["cidr_blocks"] = cidr_blocks
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["max_ttl"] = max_ttl
            if mount is None and not opts.urn:
                raise TypeError("Missing required property 'mount'")
            __props__.__dict__["mount"] = mount
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["organization_id"] = organization_id
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["project_roles"] = project_roles
            if roles is None and not opts.urn:
                raise TypeError("Missing required property 'roles'")
            __props__.__dict__["roles"] = roles
            __props__.__dict__["ttl"] = ttl
        super(SecretRole, __self__).__init__(
            'vault:mongodbatlas/secretRole:SecretRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cidr_blocks: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            max_ttl: Optional[pulumi.Input[str]] = None,
            mount: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            organization_id: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            project_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[str]] = None) -> 'SecretRole':
        """
        Get an existing SecretRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] cidr_blocks: Whitelist entry in CIDR notation to be added for the API key.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: IP address to be added to the whitelist for the API key.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of credentials issued using this role.
        :param pulumi.Input[str] mount: Path where the MongoDB Atlas Secrets Engine is mounted.
        :param pulumi.Input[str] name: The name of the role.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization_id: Unique identifier for the organization to which the target API Key belongs.
               Required if `project_id` is not set.
        :param pulumi.Input[str] project_id: Unique identifier for the project to which the target API Key belongs.
               Required if `organization_id` is not set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] project_roles: Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] roles: List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        :param pulumi.Input[str] ttl: Duration in seconds after which the issued credential should expire.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretRoleState.__new__(_SecretRoleState)

        __props__.__dict__["cidr_blocks"] = cidr_blocks
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["max_ttl"] = max_ttl
        __props__.__dict__["mount"] = mount
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["organization_id"] = organization_id
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["project_roles"] = project_roles
        __props__.__dict__["roles"] = roles
        __props__.__dict__["ttl"] = ttl
        return SecretRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cidrBlocks")
    def cidr_blocks(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Whitelist entry in CIDR notation to be added for the API key.
        """
        return pulumi.get(self, "cidr_blocks")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        IP address to be added to the whitelist for the API key.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum allowed lifetime of credentials issued using this role.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[str]:
        """
        Path where the MongoDB Atlas Secrets Engine is mounted.
        """
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the role.
        """
        return pulumi.get(self, "name")

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
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[Optional[str]]:
        """
        Unique identifier for the organization to which the target API Key belongs.
        Required if `project_id` is not set.
        """
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        Unique identifier for the project to which the target API Key belongs.
        Required if `organization_id` is not set.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectRoles")
    def project_roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Roles assigned when an org API key is assigned to a project API key. Possible values are `GROUP_CLUSTER_MANAGER`, `GROUP_DATA_ACCESS_ADMIN`, `GROUP_DATA_ACCESS_READ_ONLY`, `GROUP_DATA_ACCESS_READ_WRITE`, `GROUP_OWNER` and `GROUP_READ_ONLY`.
        """
        return pulumi.get(self, "project_roles")

    @property
    @pulumi.getter
    def roles(self) -> pulumi.Output[Sequence[str]]:
        """
        List of roles that the API Key needs to have. Possible values are `ORG_OWNER`, `ORG_MEMBER`, `ORG_GROUP_CREATOR`, `ORG_BILLING_ADMIN` and `ORG_READ_ONLY`.
        """
        return pulumi.get(self, "roles")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Duration in seconds after which the issued credential should expire.
        """
        return pulumi.get(self, "ttl")

