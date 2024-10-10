# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['SecretBackendStaticRoleArgs', 'SecretBackendStaticRole']

@pulumi.input_type
class SecretBackendStaticRoleArgs:
    def __init__(__self__, *,
                 role_name: pulumi.Input[str],
                 rotation_period: pulumi.Input[int],
                 username: pulumi.Input[str],
                 dn: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 skip_import_rotation: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a SecretBackendStaticRole resource.
        :param pulumi.Input[str] role_name: Name of the role.
        :param pulumi.Input[int] rotation_period: How often Vault should rotate the password of the user entry.
        :param pulumi.Input[str] username: The username of the existing LDAP entry to manage password rotation for.
        :param pulumi.Input[str] dn: Distinguished name (DN) of the existing LDAP entry to manage
               password rotation for. If given, it will take precedence over `username` for the LDAP
               search performed during password rotation. Cannot be modified after creation.
        :param pulumi.Input[str] mount: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `ldap`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] skip_import_rotation: Causes vault to skip the initial secret rotation on import. Not applicable to updates.
               Requires Vault 1.16 or above.
        """
        pulumi.set(__self__, "role_name", role_name)
        pulumi.set(__self__, "rotation_period", rotation_period)
        pulumi.set(__self__, "username", username)
        if dn is not None:
            pulumi.set(__self__, "dn", dn)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if skip_import_rotation is not None:
            pulumi.set(__self__, "skip_import_rotation", skip_import_rotation)

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> pulumi.Input[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Input[int]:
        """
        How often Vault should rotate the password of the user entry.
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: pulumi.Input[int]):
        pulumi.set(self, "rotation_period", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The username of the existing LDAP entry to manage password rotation for.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def dn(self) -> Optional[pulumi.Input[str]]:
        """
        Distinguished name (DN) of the existing LDAP entry to manage
        password rotation for. If given, it will take precedence over `username` for the LDAP
        search performed during password rotation. Cannot be modified after creation.
        """
        return pulumi.get(self, "dn")

    @dn.setter
    def dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dn", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `ldap`.
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
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="skipImportRotation")
    def skip_import_rotation(self) -> Optional[pulumi.Input[bool]]:
        """
        Causes vault to skip the initial secret rotation on import. Not applicable to updates.
        Requires Vault 1.16 or above.
        """
        return pulumi.get(self, "skip_import_rotation")

    @skip_import_rotation.setter
    def skip_import_rotation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_import_rotation", value)


@pulumi.input_type
class _SecretBackendStaticRoleState:
    def __init__(__self__, *,
                 dn: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[int]] = None,
                 skip_import_rotation: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackendStaticRole resources.
        :param pulumi.Input[str] dn: Distinguished name (DN) of the existing LDAP entry to manage
               password rotation for. If given, it will take precedence over `username` for the LDAP
               search performed during password rotation. Cannot be modified after creation.
        :param pulumi.Input[str] mount: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `ldap`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: Name of the role.
        :param pulumi.Input[int] rotation_period: How often Vault should rotate the password of the user entry.
        :param pulumi.Input[bool] skip_import_rotation: Causes vault to skip the initial secret rotation on import. Not applicable to updates.
               Requires Vault 1.16 or above.
        :param pulumi.Input[str] username: The username of the existing LDAP entry to manage password rotation for.
        """
        if dn is not None:
            pulumi.set(__self__, "dn", dn)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if role_name is not None:
            pulumi.set(__self__, "role_name", role_name)
        if rotation_period is not None:
            pulumi.set(__self__, "rotation_period", rotation_period)
        if skip_import_rotation is not None:
            pulumi.set(__self__, "skip_import_rotation", skip_import_rotation)
        if username is not None:
            pulumi.set(__self__, "username", username)

    @property
    @pulumi.getter
    def dn(self) -> Optional[pulumi.Input[str]]:
        """
        Distinguished name (DN) of the existing LDAP entry to manage
        password rotation for. If given, it will take precedence over `username` for the LDAP
        search performed during password rotation. Cannot be modified after creation.
        """
        return pulumi.get(self, "dn")

    @dn.setter
    def dn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "dn", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `ldap`.
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
        Name of the role.
        """
        return pulumi.get(self, "role_name")

    @role_name.setter
    def role_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role_name", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> Optional[pulumi.Input[int]]:
        """
        How often Vault should rotate the password of the user entry.
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rotation_period", value)

    @property
    @pulumi.getter(name="skipImportRotation")
    def skip_import_rotation(self) -> Optional[pulumi.Input[bool]]:
        """
        Causes vault to skip the initial secret rotation on import. Not applicable to updates.
        Requires Vault 1.16 or above.
        """
        return pulumi.get(self, "skip_import_rotation")

    @skip_import_rotation.setter
    def skip_import_rotation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "skip_import_rotation", value)

    @property
    @pulumi.getter
    def username(self) -> Optional[pulumi.Input[str]]:
        """
        The username of the existing LDAP entry to manage password rotation for.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username", value)


class SecretBackendStaticRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dn: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[int]] = None,
                 skip_import_rotation: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.ldap.SecretBackend("config",
            path="my-custom-ldap",
            binddn="CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
            bindpass="SuperSecretPassw0rd",
            url="ldaps://localhost",
            insecure_tls=True,
            userdn="CN=Users,DC=corp,DC=example,DC=net")
        role = vault.ldap.SecretBackendStaticRole("role",
            mount=config.path,
            username="alice",
            dn="cn=alice,ou=Users,DC=corp,DC=example,DC=net",
            role_name="alice",
            rotation_period=60)
        ```

        ## Import

        LDAP secret backend static role can be imported using the full path to the role
        of the form: `<mount_path>/static-role/<role_name>` e.g.

        ```sh
        $ pulumi import vault:ldap/secretBackendStaticRole:SecretBackendStaticRole role ldap/static-role/example-role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dn: Distinguished name (DN) of the existing LDAP entry to manage
               password rotation for. If given, it will take precedence over `username` for the LDAP
               search performed during password rotation. Cannot be modified after creation.
        :param pulumi.Input[str] mount: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `ldap`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: Name of the role.
        :param pulumi.Input[int] rotation_period: How often Vault should rotate the password of the user entry.
        :param pulumi.Input[bool] skip_import_rotation: Causes vault to skip the initial secret rotation on import. Not applicable to updates.
               Requires Vault 1.16 or above.
        :param pulumi.Input[str] username: The username of the existing LDAP entry to manage password rotation for.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendStaticRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.ldap.SecretBackend("config",
            path="my-custom-ldap",
            binddn="CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
            bindpass="SuperSecretPassw0rd",
            url="ldaps://localhost",
            insecure_tls=True,
            userdn="CN=Users,DC=corp,DC=example,DC=net")
        role = vault.ldap.SecretBackendStaticRole("role",
            mount=config.path,
            username="alice",
            dn="cn=alice,ou=Users,DC=corp,DC=example,DC=net",
            role_name="alice",
            rotation_period=60)
        ```

        ## Import

        LDAP secret backend static role can be imported using the full path to the role
        of the form: `<mount_path>/static-role/<role_name>` e.g.

        ```sh
        $ pulumi import vault:ldap/secretBackendStaticRole:SecretBackendStaticRole role ldap/static-role/example-role
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendStaticRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendStaticRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 dn: Optional[pulumi.Input[str]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role_name: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[int]] = None,
                 skip_import_rotation: Optional[pulumi.Input[bool]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendStaticRoleArgs.__new__(SecretBackendStaticRoleArgs)

            __props__.__dict__["dn"] = dn
            __props__.__dict__["mount"] = mount
            __props__.__dict__["namespace"] = namespace
            if role_name is None and not opts.urn:
                raise TypeError("Missing required property 'role_name'")
            __props__.__dict__["role_name"] = role_name
            if rotation_period is None and not opts.urn:
                raise TypeError("Missing required property 'rotation_period'")
            __props__.__dict__["rotation_period"] = rotation_period
            __props__.__dict__["skip_import_rotation"] = skip_import_rotation
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__.__dict__["username"] = username
        super(SecretBackendStaticRole, __self__).__init__(
            'vault:ldap/secretBackendStaticRole:SecretBackendStaticRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            dn: Optional[pulumi.Input[str]] = None,
            mount: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            role_name: Optional[pulumi.Input[str]] = None,
            rotation_period: Optional[pulumi.Input[int]] = None,
            skip_import_rotation: Optional[pulumi.Input[bool]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'SecretBackendStaticRole':
        """
        Get an existing SecretBackendStaticRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] dn: Distinguished name (DN) of the existing LDAP entry to manage
               password rotation for. If given, it will take precedence over `username` for the LDAP
               search performed during password rotation. Cannot be modified after creation.
        :param pulumi.Input[str] mount: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `ldap`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role_name: Name of the role.
        :param pulumi.Input[int] rotation_period: How often Vault should rotate the password of the user entry.
        :param pulumi.Input[bool] skip_import_rotation: Causes vault to skip the initial secret rotation on import. Not applicable to updates.
               Requires Vault 1.16 or above.
        :param pulumi.Input[str] username: The username of the existing LDAP entry to manage password rotation for.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendStaticRoleState.__new__(_SecretBackendStaticRoleState)

        __props__.__dict__["dn"] = dn
        __props__.__dict__["mount"] = mount
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["role_name"] = role_name
        __props__.__dict__["rotation_period"] = rotation_period
        __props__.__dict__["skip_import_rotation"] = skip_import_rotation
        __props__.__dict__["username"] = username
        return SecretBackendStaticRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def dn(self) -> pulumi.Output[Optional[str]]:
        """
        Distinguished name (DN) of the existing LDAP entry to manage
        password rotation for. If given, it will take precedence over `username` for the LDAP
        search performed during password rotation. Cannot be modified after creation.
        """
        return pulumi.get(self, "dn")

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[Optional[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `ldap`.
        """
        return pulumi.get(self, "mount")

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
        Name of the role.
        """
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Output[int]:
        """
        How often Vault should rotate the password of the user entry.
        """
        return pulumi.get(self, "rotation_period")

    @property
    @pulumi.getter(name="skipImportRotation")
    def skip_import_rotation(self) -> pulumi.Output[Optional[bool]]:
        """
        Causes vault to skip the initial secret rotation on import. Not applicable to updates.
        Requires Vault 1.16 or above.
        """
        return pulumi.get(self, "skip_import_rotation")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The username of the existing LDAP entry to manage password rotation for.
        """
        return pulumi.get(self, "username")

