# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OidcKeyAllowedClientIDArgs', 'OidcKeyAllowedClientID']

@pulumi.input_type
class OidcKeyAllowedClientIDArgs:
    def __init__(__self__, *,
                 allowed_client_id: pulumi.Input[str],
                 key_name: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OidcKeyAllowedClientID resource.
        :param pulumi.Input[str] allowed_client_id: Client ID to allow usage with the OIDC named key
        :param pulumi.Input[str] key_name: Name of the OIDC Key allow the Client ID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "allowed_client_id", allowed_client_id)
        pulumi.set(__self__, "key_name", key_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="allowedClientId")
    def allowed_client_id(self) -> pulumi.Input[str]:
        """
        Client ID to allow usage with the OIDC named key
        """
        return pulumi.get(self, "allowed_client_id")

    @allowed_client_id.setter
    def allowed_client_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "allowed_client_id", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Input[str]:
        """
        Name of the OIDC Key allow the Client ID.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_name", value)

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


@pulumi.input_type
class _OidcKeyAllowedClientIDState:
    def __init__(__self__, *,
                 allowed_client_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OidcKeyAllowedClientID resources.
        :param pulumi.Input[str] allowed_client_id: Client ID to allow usage with the OIDC named key
        :param pulumi.Input[str] key_name: Name of the OIDC Key allow the Client ID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if allowed_client_id is not None:
            pulumi.set(__self__, "allowed_client_id", allowed_client_id)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="allowedClientId")
    def allowed_client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client ID to allow usage with the OIDC named key
        """
        return pulumi.get(self, "allowed_client_id")

    @allowed_client_id.setter
    def allowed_client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allowed_client_id", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the OIDC Key allow the Client ID.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

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


class OidcKeyAllowedClientID(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_client_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        key = vault.identity.OidcKey("key", algorithm="RS256")
        role_oidc_role = vault.identity.OidcRole("roleOidcRole", key=key.name)
        role_oidc_key_allowed_client_id = vault.identity.OidcKeyAllowedClientID("roleOidcKeyAllowedClientID",
            key_name=key.name,
            allowed_client_id=role_oidc_role.client_id)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_client_id: Client ID to allow usage with the OIDC named key
        :param pulumi.Input[str] key_name: Name of the OIDC Key allow the Client ID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OidcKeyAllowedClientIDArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        key = vault.identity.OidcKey("key", algorithm="RS256")
        role_oidc_role = vault.identity.OidcRole("roleOidcRole", key=key.name)
        role_oidc_key_allowed_client_id = vault.identity.OidcKeyAllowedClientID("roleOidcKeyAllowedClientID",
            key_name=key.name,
            allowed_client_id=role_oidc_role.client_id)
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param OidcKeyAllowedClientIDArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OidcKeyAllowedClientIDArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_client_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OidcKeyAllowedClientIDArgs.__new__(OidcKeyAllowedClientIDArgs)

            if allowed_client_id is None and not opts.urn:
                raise TypeError("Missing required property 'allowed_client_id'")
            __props__.__dict__["allowed_client_id"] = allowed_client_id
            if key_name is None and not opts.urn:
                raise TypeError("Missing required property 'key_name'")
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["namespace"] = namespace
        super(OidcKeyAllowedClientID, __self__).__init__(
            'vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_client_id: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None) -> 'OidcKeyAllowedClientID':
        """
        Get an existing OidcKeyAllowedClientID resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allowed_client_id: Client ID to allow usage with the OIDC named key
        :param pulumi.Input[str] key_name: Name of the OIDC Key allow the Client ID.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OidcKeyAllowedClientIDState.__new__(_OidcKeyAllowedClientIDState)

        __props__.__dict__["allowed_client_id"] = allowed_client_id
        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["namespace"] = namespace
        return OidcKeyAllowedClientID(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedClientId")
    def allowed_client_id(self) -> pulumi.Output[str]:
        """
        Client ID to allow usage with the OIDC named key
        """
        return pulumi.get(self, "allowed_client_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[str]:
        """
        Name of the OIDC Key allow the Client ID.
        """
        return pulumi.get(self, "key_name")

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

