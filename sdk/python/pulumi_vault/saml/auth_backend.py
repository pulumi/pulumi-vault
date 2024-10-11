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

__all__ = ['AuthBackendArgs', 'AuthBackend']

@pulumi.input_type
class AuthBackendArgs:
    def __init__(__self__, *,
                 acs_urls: pulumi.Input[Sequence[pulumi.Input[str]]],
                 entity_id: pulumi.Input[str],
                 default_role: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 idp_cert: Optional[pulumi.Input[str]] = None,
                 idp_entity_id: Optional[pulumi.Input[str]] = None,
                 idp_metadata_url: Optional[pulumi.Input[str]] = None,
                 idp_sso_url: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 verbose_logging: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AuthBackend resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acs_urls: The well-formatted URLs of your Assertion Consumer Service (ACS)
               that should receive a response from the identity provider.
        :param pulumi.Input[str] entity_id: The entity ID of the SAML authentication service provider.
        :param pulumi.Input[str] default_role: The role to use if no role is provided during login.
        :param pulumi.Input[bool] disable_remount: If set to `true`, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[str] idp_cert: The PEM encoded certificate of the identity provider. Mutually exclusive
               with `idp_metadata_url`.
        :param pulumi.Input[str] idp_entity_id: The entity ID of the identity provider. Mutually exclusive with
               `idp_metadata_url`.
        :param pulumi.Input[str] idp_metadata_url: The metadata URL of the identity provider.
        :param pulumi.Input[str] idp_sso_url: The SSO URL of the identity provider. Mutually exclusive with 
               `idp_metadata_url`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where the auth backend will be mounted. Defaults to `auth/saml`
               if not specified.
        :param pulumi.Input[bool] verbose_logging: If set to `true`, logs additional, potentially sensitive
               information during the SAML exchange according to the current logging level. Not
               recommended for production.
        """
        pulumi.set(__self__, "acs_urls", acs_urls)
        pulumi.set(__self__, "entity_id", entity_id)
        if default_role is not None:
            pulumi.set(__self__, "default_role", default_role)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if idp_cert is not None:
            pulumi.set(__self__, "idp_cert", idp_cert)
        if idp_entity_id is not None:
            pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        if idp_metadata_url is not None:
            pulumi.set(__self__, "idp_metadata_url", idp_metadata_url)
        if idp_sso_url is not None:
            pulumi.set(__self__, "idp_sso_url", idp_sso_url)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if verbose_logging is not None:
            pulumi.set(__self__, "verbose_logging", verbose_logging)

    @property
    @pulumi.getter(name="acsUrls")
    def acs_urls(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        The well-formatted URLs of your Assertion Consumer Service (ACS)
        that should receive a response from the identity provider.
        """
        return pulumi.get(self, "acs_urls")

    @acs_urls.setter
    def acs_urls(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "acs_urls", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Input[str]:
        """
        The entity ID of the SAML authentication service provider.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> Optional[pulumi.Input[str]]:
        """
        The role to use if no role is provided during login.
        """
        return pulumi.get(self, "default_role")

    @default_role.setter
    def default_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_role", value)

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_remount", value)

    @property
    @pulumi.getter(name="idpCert")
    def idp_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The PEM encoded certificate of the identity provider. Mutually exclusive
        with `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_cert")

    @idp_cert.setter
    def idp_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_cert", value)

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The entity ID of the identity provider. Mutually exclusive with
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_entity_id")

    @idp_entity_id.setter
    def idp_entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_entity_id", value)

    @property
    @pulumi.getter(name="idpMetadataUrl")
    def idp_metadata_url(self) -> Optional[pulumi.Input[str]]:
        """
        The metadata URL of the identity provider.
        """
        return pulumi.get(self, "idp_metadata_url")

    @idp_metadata_url.setter
    def idp_metadata_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_metadata_url", value)

    @property
    @pulumi.getter(name="idpSsoUrl")
    def idp_sso_url(self) -> Optional[pulumi.Input[str]]:
        """
        The SSO URL of the identity provider. Mutually exclusive with 
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_sso_url")

    @idp_sso_url.setter
    def idp_sso_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_sso_url", value)

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
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the auth backend will be mounted. Defaults to `auth/saml`
        if not specified.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="verboseLogging")
    def verbose_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, logs additional, potentially sensitive
        information during the SAML exchange according to the current logging level. Not
        recommended for production.
        """
        return pulumi.get(self, "verbose_logging")

    @verbose_logging.setter
    def verbose_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verbose_logging", value)


@pulumi.input_type
class _AuthBackendState:
    def __init__(__self__, *,
                 acs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 idp_cert: Optional[pulumi.Input[str]] = None,
                 idp_entity_id: Optional[pulumi.Input[str]] = None,
                 idp_metadata_url: Optional[pulumi.Input[str]] = None,
                 idp_sso_url: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 verbose_logging: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AuthBackend resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acs_urls: The well-formatted URLs of your Assertion Consumer Service (ACS)
               that should receive a response from the identity provider.
        :param pulumi.Input[str] default_role: The role to use if no role is provided during login.
        :param pulumi.Input[bool] disable_remount: If set to `true`, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[str] entity_id: The entity ID of the SAML authentication service provider.
        :param pulumi.Input[str] idp_cert: The PEM encoded certificate of the identity provider. Mutually exclusive
               with `idp_metadata_url`.
        :param pulumi.Input[str] idp_entity_id: The entity ID of the identity provider. Mutually exclusive with
               `idp_metadata_url`.
        :param pulumi.Input[str] idp_metadata_url: The metadata URL of the identity provider.
        :param pulumi.Input[str] idp_sso_url: The SSO URL of the identity provider. Mutually exclusive with 
               `idp_metadata_url`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where the auth backend will be mounted. Defaults to `auth/saml`
               if not specified.
        :param pulumi.Input[bool] verbose_logging: If set to `true`, logs additional, potentially sensitive
               information during the SAML exchange according to the current logging level. Not
               recommended for production.
        """
        if acs_urls is not None:
            pulumi.set(__self__, "acs_urls", acs_urls)
        if default_role is not None:
            pulumi.set(__self__, "default_role", default_role)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if entity_id is not None:
            pulumi.set(__self__, "entity_id", entity_id)
        if idp_cert is not None:
            pulumi.set(__self__, "idp_cert", idp_cert)
        if idp_entity_id is not None:
            pulumi.set(__self__, "idp_entity_id", idp_entity_id)
        if idp_metadata_url is not None:
            pulumi.set(__self__, "idp_metadata_url", idp_metadata_url)
        if idp_sso_url is not None:
            pulumi.set(__self__, "idp_sso_url", idp_sso_url)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if verbose_logging is not None:
            pulumi.set(__self__, "verbose_logging", verbose_logging)

    @property
    @pulumi.getter(name="acsUrls")
    def acs_urls(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The well-formatted URLs of your Assertion Consumer Service (ACS)
        that should receive a response from the identity provider.
        """
        return pulumi.get(self, "acs_urls")

    @acs_urls.setter
    def acs_urls(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "acs_urls", value)

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> Optional[pulumi.Input[str]]:
        """
        The role to use if no role is provided during login.
        """
        return pulumi.get(self, "default_role")

    @default_role.setter
    def default_role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_role", value)

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_remount", value)

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The entity ID of the SAML authentication service provider.
        """
        return pulumi.get(self, "entity_id")

    @entity_id.setter
    def entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_id", value)

    @property
    @pulumi.getter(name="idpCert")
    def idp_cert(self) -> Optional[pulumi.Input[str]]:
        """
        The PEM encoded certificate of the identity provider. Mutually exclusive
        with `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_cert")

    @idp_cert.setter
    def idp_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_cert", value)

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> Optional[pulumi.Input[str]]:
        """
        The entity ID of the identity provider. Mutually exclusive with
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_entity_id")

    @idp_entity_id.setter
    def idp_entity_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_entity_id", value)

    @property
    @pulumi.getter(name="idpMetadataUrl")
    def idp_metadata_url(self) -> Optional[pulumi.Input[str]]:
        """
        The metadata URL of the identity provider.
        """
        return pulumi.get(self, "idp_metadata_url")

    @idp_metadata_url.setter
    def idp_metadata_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_metadata_url", value)

    @property
    @pulumi.getter(name="idpSsoUrl")
    def idp_sso_url(self) -> Optional[pulumi.Input[str]]:
        """
        The SSO URL of the identity provider. Mutually exclusive with 
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_sso_url")

    @idp_sso_url.setter
    def idp_sso_url(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "idp_sso_url", value)

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
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the auth backend will be mounted. Defaults to `auth/saml`
        if not specified.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="verboseLogging")
    def verbose_logging(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, logs additional, potentially sensitive
        information during the SAML exchange according to the current logging level. Not
        recommended for production.
        """
        return pulumi.get(self, "verbose_logging")

    @verbose_logging.setter
    def verbose_logging(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "verbose_logging", value)


class AuthBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 idp_cert: Optional[pulumi.Input[str]] = None,
                 idp_entity_id: Optional[pulumi.Input[str]] = None,
                 idp_metadata_url: Optional[pulumi.Input[str]] = None,
                 idp_sso_url: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 verbose_logging: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Manages a SAML Auth mount in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/saml/) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.saml.AuthBackend("test",
            path="saml",
            idp_metadata_url="https://company.okta.com/app/abc123eb9xnIfzlaf697/sso/saml/metadata",
            entity_id="https://my.vault/v1/auth/saml",
            acs_urls=["https://my.vault.primary/v1/auth/saml/callback"],
            default_role="admin")
        ```

        ## Import

        SAML authentication mounts can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:saml/authBackend:AuthBackend example saml
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acs_urls: The well-formatted URLs of your Assertion Consumer Service (ACS)
               that should receive a response from the identity provider.
        :param pulumi.Input[str] default_role: The role to use if no role is provided during login.
        :param pulumi.Input[bool] disable_remount: If set to `true`, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[str] entity_id: The entity ID of the SAML authentication service provider.
        :param pulumi.Input[str] idp_cert: The PEM encoded certificate of the identity provider. Mutually exclusive
               with `idp_metadata_url`.
        :param pulumi.Input[str] idp_entity_id: The entity ID of the identity provider. Mutually exclusive with
               `idp_metadata_url`.
        :param pulumi.Input[str] idp_metadata_url: The metadata URL of the identity provider.
        :param pulumi.Input[str] idp_sso_url: The SSO URL of the identity provider. Mutually exclusive with 
               `idp_metadata_url`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where the auth backend will be mounted. Defaults to `auth/saml`
               if not specified.
        :param pulumi.Input[bool] verbose_logging: If set to `true`, logs additional, potentially sensitive
               information during the SAML exchange according to the current logging level. Not
               recommended for production.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages a SAML Auth mount in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/saml/) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.saml.AuthBackend("test",
            path="saml",
            idp_metadata_url="https://company.okta.com/app/abc123eb9xnIfzlaf697/sso/saml/metadata",
            entity_id="https://my.vault/v1/auth/saml",
            acs_urls=["https://my.vault.primary/v1/auth/saml/callback"],
            default_role="admin")
        ```

        ## Import

        SAML authentication mounts can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:saml/authBackend:AuthBackend example saml
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 acs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_role: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 entity_id: Optional[pulumi.Input[str]] = None,
                 idp_cert: Optional[pulumi.Input[str]] = None,
                 idp_entity_id: Optional[pulumi.Input[str]] = None,
                 idp_metadata_url: Optional[pulumi.Input[str]] = None,
                 idp_sso_url: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 verbose_logging: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendArgs.__new__(AuthBackendArgs)

            if acs_urls is None and not opts.urn:
                raise TypeError("Missing required property 'acs_urls'")
            __props__.__dict__["acs_urls"] = acs_urls
            __props__.__dict__["default_role"] = default_role
            __props__.__dict__["disable_remount"] = disable_remount
            if entity_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_id'")
            __props__.__dict__["entity_id"] = entity_id
            __props__.__dict__["idp_cert"] = idp_cert
            __props__.__dict__["idp_entity_id"] = idp_entity_id
            __props__.__dict__["idp_metadata_url"] = idp_metadata_url
            __props__.__dict__["idp_sso_url"] = idp_sso_url
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["path"] = path
            __props__.__dict__["verbose_logging"] = verbose_logging
        super(AuthBackend, __self__).__init__(
            'vault:saml/authBackend:AuthBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            acs_urls: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            default_role: Optional[pulumi.Input[str]] = None,
            disable_remount: Optional[pulumi.Input[bool]] = None,
            entity_id: Optional[pulumi.Input[str]] = None,
            idp_cert: Optional[pulumi.Input[str]] = None,
            idp_entity_id: Optional[pulumi.Input[str]] = None,
            idp_metadata_url: Optional[pulumi.Input[str]] = None,
            idp_sso_url: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            verbose_logging: Optional[pulumi.Input[bool]] = None) -> 'AuthBackend':
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] acs_urls: The well-formatted URLs of your Assertion Consumer Service (ACS)
               that should receive a response from the identity provider.
        :param pulumi.Input[str] default_role: The role to use if no role is provided during login.
        :param pulumi.Input[bool] disable_remount: If set to `true`, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[str] entity_id: The entity ID of the SAML authentication service provider.
        :param pulumi.Input[str] idp_cert: The PEM encoded certificate of the identity provider. Mutually exclusive
               with `idp_metadata_url`.
        :param pulumi.Input[str] idp_entity_id: The entity ID of the identity provider. Mutually exclusive with
               `idp_metadata_url`.
        :param pulumi.Input[str] idp_metadata_url: The metadata URL of the identity provider.
        :param pulumi.Input[str] idp_sso_url: The SSO URL of the identity provider. Mutually exclusive with 
               `idp_metadata_url`.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path where the auth backend will be mounted. Defaults to `auth/saml`
               if not specified.
        :param pulumi.Input[bool] verbose_logging: If set to `true`, logs additional, potentially sensitive
               information during the SAML exchange according to the current logging level. Not
               recommended for production.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendState.__new__(_AuthBackendState)

        __props__.__dict__["acs_urls"] = acs_urls
        __props__.__dict__["default_role"] = default_role
        __props__.__dict__["disable_remount"] = disable_remount
        __props__.__dict__["entity_id"] = entity_id
        __props__.__dict__["idp_cert"] = idp_cert
        __props__.__dict__["idp_entity_id"] = idp_entity_id
        __props__.__dict__["idp_metadata_url"] = idp_metadata_url
        __props__.__dict__["idp_sso_url"] = idp_sso_url
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["verbose_logging"] = verbose_logging
        return AuthBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="acsUrls")
    def acs_urls(self) -> pulumi.Output[Sequence[str]]:
        """
        The well-formatted URLs of your Assertion Consumer Service (ACS)
        that should receive a response from the identity provider.
        """
        return pulumi.get(self, "acs_urls")

    @property
    @pulumi.getter(name="defaultRole")
    def default_role(self) -> pulumi.Output[Optional[str]]:
        """
        The role to use if no role is provided during login.
        """
        return pulumi.get(self, "default_role")

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @property
    @pulumi.getter(name="entityId")
    def entity_id(self) -> pulumi.Output[str]:
        """
        The entity ID of the SAML authentication service provider.
        """
        return pulumi.get(self, "entity_id")

    @property
    @pulumi.getter(name="idpCert")
    def idp_cert(self) -> pulumi.Output[Optional[str]]:
        """
        The PEM encoded certificate of the identity provider. Mutually exclusive
        with `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_cert")

    @property
    @pulumi.getter(name="idpEntityId")
    def idp_entity_id(self) -> pulumi.Output[Optional[str]]:
        """
        The entity ID of the identity provider. Mutually exclusive with
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_entity_id")

    @property
    @pulumi.getter(name="idpMetadataUrl")
    def idp_metadata_url(self) -> pulumi.Output[Optional[str]]:
        """
        The metadata URL of the identity provider.
        """
        return pulumi.get(self, "idp_metadata_url")

    @property
    @pulumi.getter(name="idpSsoUrl")
    def idp_sso_url(self) -> pulumi.Output[Optional[str]]:
        """
        The SSO URL of the identity provider. Mutually exclusive with 
        `idp_metadata_url`.
        """
        return pulumi.get(self, "idp_sso_url")

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
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        Path where the auth backend will be mounted. Defaults to `auth/saml`
        if not specified.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="verboseLogging")
    def verbose_logging(self) -> pulumi.Output[bool]:
        """
        If set to `true`, logs additional, potentially sensitive
        information during the SAML exchange according to the current logging level. Not
        recommended for production.
        """
        return pulumi.get(self, "verbose_logging")

