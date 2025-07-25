# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['SecretBackendConfigIssuersArgs', 'SecretBackendConfigIssuers']

@pulumi.input_type
class SecretBackendConfigIssuersArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[_builtins.str],
                 default: Optional[pulumi.Input[_builtins.str]] = None,
                 default_follows_latest_issuer: Optional[pulumi.Input[_builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a SecretBackendConfigIssuers resource.
        :param pulumi.Input[_builtins.str] backend: The path the PKI secret backend is mounted at, with no
               leading or trailing `/`s.
        :param pulumi.Input[_builtins.str] default: Specifies the default issuer by ID.
        :param pulumi.Input[_builtins.bool] default_follows_latest_issuer: Specifies whether a root creation
               or an issuer import operation updates the default issuer to the newly added issuer.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "backend", backend)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if default_follows_latest_issuer is not None:
            pulumi.set(__self__, "default_follows_latest_issuer", default_follows_latest_issuer)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> pulumi.Input[_builtins.str]:
        """
        The path the PKI secret backend is mounted at, with no
        leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "backend", value)

    @_builtins.property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the default issuer by ID.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "default", value)

    @_builtins.property
    @pulumi.getter(name="defaultFollowsLatestIssuer")
    def default_follows_latest_issuer(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies whether a root creation
        or an issuer import operation updates the default issuer to the newly added issuer.
        """
        return pulumi.get(self, "default_follows_latest_issuer")

    @default_follows_latest_issuer.setter
    def default_follows_latest_issuer(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "default_follows_latest_issuer", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _SecretBackendConfigIssuersState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 default: Optional[pulumi.Input[_builtins.str]] = None,
                 default_follows_latest_issuer: Optional[pulumi.Input[_builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering SecretBackendConfigIssuers resources.
        :param pulumi.Input[_builtins.str] backend: The path the PKI secret backend is mounted at, with no
               leading or trailing `/`s.
        :param pulumi.Input[_builtins.str] default: Specifies the default issuer by ID.
        :param pulumi.Input[_builtins.bool] default_follows_latest_issuer: Specifies whether a root creation
               or an issuer import operation updates the default issuer to the newly added issuer.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if default is not None:
            pulumi.set(__self__, "default", default)
        if default_follows_latest_issuer is not None:
            pulumi.set(__self__, "default_follows_latest_issuer", default_follows_latest_issuer)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The path the PKI secret backend is mounted at, with no
        leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "backend", value)

    @_builtins.property
    @pulumi.getter
    def default(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the default issuer by ID.
        """
        return pulumi.get(self, "default")

    @default.setter
    def default(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "default", value)

    @_builtins.property
    @pulumi.getter(name="defaultFollowsLatestIssuer")
    def default_follows_latest_issuer(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Specifies whether a root creation
        or an issuer import operation updates the default issuer to the newly added issuer.
        """
        return pulumi.get(self, "default_follows_latest_issuer")

    @default_follows_latest_issuer.setter
    def default_follows_latest_issuer(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "default_follows_latest_issuer", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)


@pulumi.type_token("vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers")
class SecretBackendConfigIssuers(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 default: Optional[pulumi.Input[_builtins.str]] = None,
                 default_follows_latest_issuer: Optional[pulumi.Input[_builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        pki = vault.Mount("pki",
            path="pki",
            type="pki",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=86400)
        root = vault.pkisecret.SecretBackendRootCert("root",
            backend=pki.path,
            type="internal",
            common_name="test",
            ttl="86400")
        example = vault.pkisecret.SecretBackendIssuer("example",
            backend=root.backend,
            issuer_ref=root.issuer_id,
            issuer_name="example-issuer")
        config = vault.pkisecret.SecretBackendConfigIssuers("config",
            backend=pki.path,
            default=example.issuer_id,
            default_follows_latest_issuer=True)
        ```

        ## Import

        PKI secret backend config issuers can be imported using the path, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers config pki/config/issuers
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] backend: The path the PKI secret backend is mounted at, with no
               leading or trailing `/`s.
        :param pulumi.Input[_builtins.str] default: Specifies the default issuer by ID.
        :param pulumi.Input[_builtins.bool] default_follows_latest_issuer: Specifies whether a root creation
               or an issuer import operation updates the default issuer to the newly added issuer.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendConfigIssuersArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        pki = vault.Mount("pki",
            path="pki",
            type="pki",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=86400)
        root = vault.pkisecret.SecretBackendRootCert("root",
            backend=pki.path,
            type="internal",
            common_name="test",
            ttl="86400")
        example = vault.pkisecret.SecretBackendIssuer("example",
            backend=root.backend,
            issuer_ref=root.issuer_id,
            issuer_name="example-issuer")
        config = vault.pkisecret.SecretBackendConfigIssuers("config",
            backend=pki.path,
            default=example.issuer_id,
            default_follows_latest_issuer=True)
        ```

        ## Import

        PKI secret backend config issuers can be imported using the path, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers config pki/config/issuers
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendConfigIssuersArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendConfigIssuersArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 default: Optional[pulumi.Input[_builtins.str]] = None,
                 default_follows_latest_issuer: Optional[pulumi.Input[_builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendConfigIssuersArgs.__new__(SecretBackendConfigIssuersArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["default"] = default
            __props__.__dict__["default_follows_latest_issuer"] = default_follows_latest_issuer
            __props__.__dict__["namespace"] = namespace
        super(SecretBackendConfigIssuers, __self__).__init__(
            'vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[_builtins.str]] = None,
            default: Optional[pulumi.Input[_builtins.str]] = None,
            default_follows_latest_issuer: Optional[pulumi.Input[_builtins.bool]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None) -> 'SecretBackendConfigIssuers':
        """
        Get an existing SecretBackendConfigIssuers resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] backend: The path the PKI secret backend is mounted at, with no
               leading or trailing `/`s.
        :param pulumi.Input[_builtins.str] default: Specifies the default issuer by ID.
        :param pulumi.Input[_builtins.bool] default_follows_latest_issuer: Specifies whether a root creation
               or an issuer import operation updates the default issuer to the newly added issuer.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendConfigIssuersState.__new__(_SecretBackendConfigIssuersState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["default"] = default
        __props__.__dict__["default_follows_latest_issuer"] = default_follows_latest_issuer
        __props__.__dict__["namespace"] = namespace
        return SecretBackendConfigIssuers(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> pulumi.Output[_builtins.str]:
        """
        The path the PKI secret backend is mounted at, with no
        leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @_builtins.property
    @pulumi.getter
    def default(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the default issuer by ID.
        """
        return pulumi.get(self, "default")

    @_builtins.property
    @pulumi.getter(name="defaultFollowsLatestIssuer")
    def default_follows_latest_issuer(self) -> pulumi.Output[_builtins.bool]:
        """
        Specifies whether a root creation
        or an issuer import operation updates the default issuer to the newly added issuer.
        """
        return pulumi.get(self, "default_follows_latest_issuer")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

