# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['BackendRoleArgs', 'BackendRole']

@pulumi.input_type
class BackendRoleArgs:
    def __init__(__self__, *,
                 role: pulumi.Input[str],
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 azure_groups: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]] = None,
                 azure_roles: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 permanently_delete: Optional[pulumi.Input[bool]] = None,
                 sign_in_audience: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a BackendRole resource.
        :param pulumi.Input[str] role: Name of the Azure role
        :param pulumi.Input[str] application_object_id: Application Object ID for an existing service principal that will
               be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        :param pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]] azure_groups: List of Azure groups to be assigned to the generated service principal.
        :param pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]] azure_roles: List of Azure roles to be assigned to the generated service principal.
        :param pulumi.Input[str] backend: Path to the mounted Azure auth backend
        :param pulumi.Input[str] description: Human-friendly description of the mount for the backend.
        :param pulumi.Input[str] max_ttl: Specifies the maximum TTL for service principals generated using this role. Accepts time
               suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] permanently_delete: Indicates whether the applications and service principals created by Vault will be permanently
               deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        :param pulumi.Input[str] sign_in_audience: Specifies the security principal types that are allowed to sign in to the application.
               Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of Azure tags to attach to an application. Requires Vault 1.16+.
        :param pulumi.Input[str] ttl: Specifies the default TTL for service principals generated using this role.
               Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        pulumi.set(__self__, "role", role)
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if azure_groups is not None:
            pulumi.set(__self__, "azure_groups", azure_groups)
        if azure_roles is not None:
            pulumi.set(__self__, "azure_roles", azure_roles)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if permanently_delete is not None:
            pulumi.set(__self__, "permanently_delete", permanently_delete)
        if sign_in_audience is not None:
            pulumi.set(__self__, "sign_in_audience", sign_in_audience)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        Name of the Azure role
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application Object ID for an existing service principal that will
        be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="azureGroups")
    def azure_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]]:
        """
        List of Azure groups to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_groups")

    @azure_groups.setter
    def azure_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]]):
        pulumi.set(self, "azure_groups", value)

    @property
    @pulumi.getter(name="azureRoles")
    def azure_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]]:
        """
        List of Azure roles to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_roles")

    @azure_roles.setter
    def azure_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]]):
        pulumi.set(self, "azure_roles", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the mounted Azure auth backend
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description of the mount for the backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum TTL for service principals generated using this role. Accepts time
        suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_ttl", value)

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
    @pulumi.getter(name="permanentlyDelete")
    def permanently_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the applications and service principals created by Vault will be permanently
        deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        """
        return pulumi.get(self, "permanently_delete")

    @permanently_delete.setter
    def permanently_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanently_delete", value)

    @property
    @pulumi.getter(name="signInAudience")
    def sign_in_audience(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the security principal types that are allowed to sign in to the application.
        Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        """
        return pulumi.get(self, "sign_in_audience")

    @sign_in_audience.setter
    def sign_in_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sign_in_audience", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Azure tags to attach to an application. Requires Vault 1.16+.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the default TTL for service principals generated using this role.
        Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _BackendRoleState:
    def __init__(__self__, *,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 azure_groups: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]] = None,
                 azure_roles: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 permanently_delete: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sign_in_audience: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering BackendRole resources.
        :param pulumi.Input[str] application_object_id: Application Object ID for an existing service principal that will
               be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        :param pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]] azure_groups: List of Azure groups to be assigned to the generated service principal.
        :param pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]] azure_roles: List of Azure roles to be assigned to the generated service principal.
        :param pulumi.Input[str] backend: Path to the mounted Azure auth backend
        :param pulumi.Input[str] description: Human-friendly description of the mount for the backend.
        :param pulumi.Input[str] max_ttl: Specifies the maximum TTL for service principals generated using this role. Accepts time
               suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] permanently_delete: Indicates whether the applications and service principals created by Vault will be permanently
               deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        :param pulumi.Input[str] role: Name of the Azure role
        :param pulumi.Input[str] sign_in_audience: Specifies the security principal types that are allowed to sign in to the application.
               Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of Azure tags to attach to an application. Requires Vault 1.16+.
        :param pulumi.Input[str] ttl: Specifies the default TTL for service principals generated using this role.
               Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        if application_object_id is not None:
            pulumi.set(__self__, "application_object_id", application_object_id)
        if azure_groups is not None:
            pulumi.set(__self__, "azure_groups", azure_groups)
        if azure_roles is not None:
            pulumi.set(__self__, "azure_roles", azure_roles)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if permanently_delete is not None:
            pulumi.set(__self__, "permanently_delete", permanently_delete)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if sign_in_audience is not None:
            pulumi.set(__self__, "sign_in_audience", sign_in_audience)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> Optional[pulumi.Input[str]]:
        """
        Application Object ID for an existing service principal that will
        be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        """
        return pulumi.get(self, "application_object_id")

    @application_object_id.setter
    def application_object_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "application_object_id", value)

    @property
    @pulumi.getter(name="azureGroups")
    def azure_groups(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]]:
        """
        List of Azure groups to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_groups")

    @azure_groups.setter
    def azure_groups(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureGroupArgs']]]]):
        pulumi.set(self, "azure_groups", value)

    @property
    @pulumi.getter(name="azureRoles")
    def azure_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]]:
        """
        List of Azure roles to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_roles")

    @azure_roles.setter
    def azure_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['BackendRoleAzureRoleArgs']]]]):
        pulumi.set(self, "azure_roles", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path to the mounted Azure auth backend
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human-friendly description of the mount for the backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum TTL for service principals generated using this role. Accepts time
        suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_ttl", value)

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
    @pulumi.getter(name="permanentlyDelete")
    def permanently_delete(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether the applications and service principals created by Vault will be permanently
        deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        """
        return pulumi.get(self, "permanently_delete")

    @permanently_delete.setter
    def permanently_delete(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "permanently_delete", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Azure role
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="signInAudience")
    def sign_in_audience(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the security principal types that are allowed to sign in to the application.
        Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        """
        return pulumi.get(self, "sign_in_audience")

    @sign_in_audience.setter
    def sign_in_audience(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sign_in_audience", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A list of Azure tags to attach to an application. Requires Vault 1.16+.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the default TTL for service principals generated using this role.
        Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)


class BackendRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 azure_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureGroupArgs']]]]] = None,
                 azure_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureRoleArgs']]]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 permanently_delete: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sign_in_audience: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        azure = vault.azure.Backend("azure",
            subscription_id=var["subscription_id"],
            tenant_id=var["tenant_id"],
            client_secret=var["client_secret"],
            client_id=var["client_id"])
        generated_role = vault.azure.BackendRole("generatedRole",
            backend=azure.path,
            role="generated_role",
            sign_in_audience="AzureADMyOrg",
            tags=[
                "team:engineering",
                "environment:development",
            ],
            ttl="300",
            max_ttl="600",
            azure_roles=[vault.azure.BackendRoleAzureRoleArgs(
                role_name="Reader",
                scope=f"/subscriptions/{var['subscription_id']}/resourceGroups/azure-vault-group",
            )])
        existing_object_id = vault.azure.BackendRole("existingObjectId",
            backend=azure.path,
            role="existing_object_id",
            application_object_id="11111111-2222-3333-4444-44444444444",
            ttl="300",
            max_ttl="600")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: Application Object ID for an existing service principal that will
               be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureGroupArgs']]]] azure_groups: List of Azure groups to be assigned to the generated service principal.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureRoleArgs']]]] azure_roles: List of Azure roles to be assigned to the generated service principal.
        :param pulumi.Input[str] backend: Path to the mounted Azure auth backend
        :param pulumi.Input[str] description: Human-friendly description of the mount for the backend.
        :param pulumi.Input[str] max_ttl: Specifies the maximum TTL for service principals generated using this role. Accepts time
               suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] permanently_delete: Indicates whether the applications and service principals created by Vault will be permanently
               deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        :param pulumi.Input[str] role: Name of the Azure role
        :param pulumi.Input[str] sign_in_audience: Specifies the security principal types that are allowed to sign in to the application.
               Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of Azure tags to attach to an application. Requires Vault 1.16+.
        :param pulumi.Input[str] ttl: Specifies the default TTL for service principals generated using this role.
               Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackendRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        azure = vault.azure.Backend("azure",
            subscription_id=var["subscription_id"],
            tenant_id=var["tenant_id"],
            client_secret=var["client_secret"],
            client_id=var["client_id"])
        generated_role = vault.azure.BackendRole("generatedRole",
            backend=azure.path,
            role="generated_role",
            sign_in_audience="AzureADMyOrg",
            tags=[
                "team:engineering",
                "environment:development",
            ],
            ttl="300",
            max_ttl="600",
            azure_roles=[vault.azure.BackendRoleAzureRoleArgs(
                role_name="Reader",
                scope=f"/subscriptions/{var['subscription_id']}/resourceGroups/azure-vault-group",
            )])
        existing_object_id = vault.azure.BackendRole("existingObjectId",
            backend=azure.path,
            role="existing_object_id",
            application_object_id="11111111-2222-3333-4444-44444444444",
            ttl="300",
            max_ttl="600")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param BackendRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackendRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 application_object_id: Optional[pulumi.Input[str]] = None,
                 azure_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureGroupArgs']]]]] = None,
                 azure_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureRoleArgs']]]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 permanently_delete: Optional[pulumi.Input[bool]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 sign_in_audience: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackendRoleArgs.__new__(BackendRoleArgs)

            __props__.__dict__["application_object_id"] = application_object_id
            __props__.__dict__["azure_groups"] = azure_groups
            __props__.__dict__["azure_roles"] = azure_roles
            __props__.__dict__["backend"] = backend
            __props__.__dict__["description"] = description
            __props__.__dict__["max_ttl"] = max_ttl
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["permanently_delete"] = permanently_delete
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["sign_in_audience"] = sign_in_audience
            __props__.__dict__["tags"] = tags
            __props__.__dict__["ttl"] = ttl
        super(BackendRole, __self__).__init__(
            'vault:azure/backendRole:BackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            application_object_id: Optional[pulumi.Input[str]] = None,
            azure_groups: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureGroupArgs']]]]] = None,
            azure_roles: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureRoleArgs']]]]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            max_ttl: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            permanently_delete: Optional[pulumi.Input[bool]] = None,
            role: Optional[pulumi.Input[str]] = None,
            sign_in_audience: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[str]] = None) -> 'BackendRole':
        """
        Get an existing BackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] application_object_id: Application Object ID for an existing service principal that will
               be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureGroupArgs']]]] azure_groups: List of Azure groups to be assigned to the generated service principal.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['BackendRoleAzureRoleArgs']]]] azure_roles: List of Azure roles to be assigned to the generated service principal.
        :param pulumi.Input[str] backend: Path to the mounted Azure auth backend
        :param pulumi.Input[str] description: Human-friendly description of the mount for the backend.
        :param pulumi.Input[str] max_ttl: Specifies the maximum TTL for service principals generated using this role. Accepts time
               suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[bool] permanently_delete: Indicates whether the applications and service principals created by Vault will be permanently
               deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        :param pulumi.Input[str] role: Name of the Azure role
        :param pulumi.Input[str] sign_in_audience: Specifies the security principal types that are allowed to sign in to the application.
               Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] tags: A list of Azure tags to attach to an application. Requires Vault 1.16+.
        :param pulumi.Input[str] ttl: Specifies the default TTL for service principals generated using this role.
               Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackendRoleState.__new__(_BackendRoleState)

        __props__.__dict__["application_object_id"] = application_object_id
        __props__.__dict__["azure_groups"] = azure_groups
        __props__.__dict__["azure_roles"] = azure_roles
        __props__.__dict__["backend"] = backend
        __props__.__dict__["description"] = description
        __props__.__dict__["max_ttl"] = max_ttl
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["permanently_delete"] = permanently_delete
        __props__.__dict__["role"] = role
        __props__.__dict__["sign_in_audience"] = sign_in_audience
        __props__.__dict__["tags"] = tags
        __props__.__dict__["ttl"] = ttl
        return BackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="applicationObjectId")
    def application_object_id(self) -> pulumi.Output[Optional[str]]:
        """
        Application Object ID for an existing service principal that will
        be used instead of creating dynamic service principals. If present, `azure_roles` and `permanently_delete` will be ignored.
        """
        return pulumi.get(self, "application_object_id")

    @property
    @pulumi.getter(name="azureGroups")
    def azure_groups(self) -> pulumi.Output[Optional[Sequence['outputs.BackendRoleAzureGroup']]]:
        """
        List of Azure groups to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_groups")

    @property
    @pulumi.getter(name="azureRoles")
    def azure_roles(self) -> pulumi.Output[Optional[Sequence['outputs.BackendRoleAzureRole']]]:
        """
        List of Azure roles to be assigned to the generated service principal.
        """
        return pulumi.get(self, "azure_roles")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Path to the mounted Azure auth backend
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Human-friendly description of the mount for the backend.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the maximum TTL for service principals generated using this role. Accepts time
        suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine max TTL time.
        """
        return pulumi.get(self, "max_ttl")

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
    @pulumi.getter(name="permanentlyDelete")
    def permanently_delete(self) -> pulumi.Output[bool]:
        """
        Indicates whether the applications and service principals created by Vault will be permanently
        deleted when the corresponding leases expire. Defaults to `false`. For Vault v1.12+.
        """
        return pulumi.get(self, "permanently_delete")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Name of the Azure role
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="signInAudience")
    def sign_in_audience(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the security principal types that are allowed to sign in to the application.
        Valid values are: AzureADMyOrg, AzureADMultipleOrgs, AzureADandPersonalMicrosoftAccount, PersonalMicrosoftAccount. Requires Vault 1.16+.
        """
        return pulumi.get(self, "sign_in_audience")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A list of Azure tags to attach to an application. Requires Vault 1.16+.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the default TTL for service principals generated using this role.
        Accepts time suffixed strings ("1h") or an integer number of seconds. Defaults to the system/engine default TTL time.
        """
        return pulumi.get(self, "ttl")

