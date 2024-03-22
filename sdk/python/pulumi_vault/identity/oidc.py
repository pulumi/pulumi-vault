# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OidcArgs', 'Oidc']

@pulumi.input_type
class OidcArgs:
    def __init__(__self__, *,
                 issuer: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Oidc resource.
        :param pulumi.Input[str] issuer: Issuer URL to be used in the iss claim of the token. If not set, Vault's
               `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
               scheme, host, and optionally, port number and path components, but no query or fragment
               components.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        Issuer URL to be used in the iss claim of the token. If not set, Vault's
        `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        scheme, host, and optionally, port number and path components, but no query or fragment
        components.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

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
class _OidcState:
    def __init__(__self__, *,
                 issuer: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Oidc resources.
        :param pulumi.Input[str] issuer: Issuer URL to be used in the iss claim of the token. If not set, Vault's
               `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
               scheme, host, and optionally, port number and path components, but no query or fragment
               components.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        Issuer URL to be used in the iss claim of the token. If not set, Vault's
        `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        scheme, host, and optionally, port number and path components, but no query or fragment
        components.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

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


class Oidc(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Configure the [Identity Tokens Backend](https://www.vaultproject.io/docs/secrets/identity/index.html#identity-tokens).

        The Identity secrets engine is the identity management solution for Vault. It internally maintains
        the clients who are recognized by Vault.

        > **NOTE:** Each Vault server may only have one Identity Tokens Backend configuration. Multiple configurations of the resource against the same Vault server will cause a perpetual difference.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        server = vault.identity.Oidc("server", issuer="https://www.acme.com")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] issuer: Issuer URL to be used in the iss claim of the token. If not set, Vault's
               `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
               scheme, host, and optionally, port number and path components, but no query or fragment
               components.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OidcArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Configure the [Identity Tokens Backend](https://www.vaultproject.io/docs/secrets/identity/index.html#identity-tokens).

        The Identity secrets engine is the identity management solution for Vault. It internally maintains
        the clients who are recognized by Vault.

        > **NOTE:** Each Vault server may only have one Identity Tokens Backend configuration. Multiple configurations of the resource against the same Vault server will cause a perpetual difference.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        server = vault.identity.Oidc("server", issuer="https://www.acme.com")
        ```
        <!--End PulumiCodeChooser -->

        :param str resource_name: The name of the resource.
        :param OidcArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OidcArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OidcArgs.__new__(OidcArgs)

            __props__.__dict__["issuer"] = issuer
            __props__.__dict__["namespace"] = namespace
        super(Oidc, __self__).__init__(
            'vault:identity/oidc:Oidc',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None) -> 'Oidc':
        """
        Get an existing Oidc resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] issuer: Issuer URL to be used in the iss claim of the token. If not set, Vault's
               `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
               scheme, host, and optionally, port number and path components, but no query or fragment
               components.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OidcState.__new__(_OidcState)

        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["namespace"] = namespace
        return Oidc(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[str]:
        """
        Issuer URL to be used in the iss claim of the token. If not set, Vault's
        `api_addr` will be used. The issuer is a case sensitive URL using the https scheme that contains
        scheme, host, and optionally, port number and path components, but no query or fragment
        components.
        """
        return pulumi.get(self, "issuer")

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

