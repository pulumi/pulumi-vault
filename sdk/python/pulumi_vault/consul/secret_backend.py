# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendArgs', 'SecretBackend']

@pulumi.input_type
class SecretBackendArgs:
    def __init__(__self__, *,
                 address: pulumi.Input[str],
                 bootstrap: Optional[pulumi.Input[bool]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackend resource.
        :param pulumi.Input[str] address: Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        :param pulumi.Input[bool] bootstrap: Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        :param pulumi.Input[str] ca_cert: CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        :param pulumi.Input[str] client_cert: Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
               this is set you need to also set client_key.
        :param pulumi.Input[str] client_key: Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
               you need to also set client_cert.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[bool] local: Specifies if the secret backend is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
               to `consul`.
        :param pulumi.Input[str] scheme: Specifies the URL scheme to use. Defaults to `http`.
        :param pulumi.Input[str] token: Specifies the Consul token to use when managing or issuing new tokens.
        """
        pulumi.set(__self__, "address", address)
        if bootstrap is not None:
            pulumi.set(__self__, "bootstrap", bootstrap)
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if local is not None:
            pulumi.set(__self__, "local", local)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if scheme is not None:
            pulumi.set(__self__, "scheme", scheme)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Input[str]:
        """
        Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: pulumi.Input[str]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def bootstrap(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        """
        return pulumi.get(self, "bootstrap")

    @bootstrap.setter
    def bootstrap(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bootstrap", value)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        this is set you need to also set client_key.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        """
        Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        you need to also set client_cert.
        """
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The default TTL for credentials issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_remount", value)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the secret backend is local only.
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

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

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        to `consul`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the URL scheme to use. Defaults to `http`.
        """
        return pulumi.get(self, "scheme")

    @scheme.setter
    def scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scheme", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Consul token to use when managing or issuing new tokens.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


@pulumi.input_type
class _SecretBackendState:
    def __init__(__self__, *,
                 address: Optional[pulumi.Input[str]] = None,
                 bootstrap: Optional[pulumi.Input[bool]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackend resources.
        :param pulumi.Input[str] address: Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        :param pulumi.Input[bool] bootstrap: Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        :param pulumi.Input[str] ca_cert: CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        :param pulumi.Input[str] client_cert: Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
               this is set you need to also set client_key.
        :param pulumi.Input[str] client_key: Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
               you need to also set client_cert.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[bool] local: Specifies if the secret backend is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
               to `consul`.
        :param pulumi.Input[str] scheme: Specifies the URL scheme to use. Defaults to `http`.
        :param pulumi.Input[str] token: Specifies the Consul token to use when managing or issuing new tokens.
        """
        if address is not None:
            pulumi.set(__self__, "address", address)
        if bootstrap is not None:
            pulumi.set(__self__, "bootstrap", bootstrap)
        if ca_cert is not None:
            pulumi.set(__self__, "ca_cert", ca_cert)
        if client_cert is not None:
            pulumi.set(__self__, "client_cert", client_cert)
        if client_key is not None:
            pulumi.set(__self__, "client_key", client_key)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if disable_remount is not None:
            pulumi.set(__self__, "disable_remount", disable_remount)
        if local is not None:
            pulumi.set(__self__, "local", local)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if scheme is not None:
            pulumi.set(__self__, "scheme", scheme)
        if token is not None:
            pulumi.set(__self__, "token", token)

    @property
    @pulumi.getter
    def address(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        """
        return pulumi.get(self, "address")

    @address.setter
    def address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "address", value)

    @property
    @pulumi.getter
    def bootstrap(self) -> Optional[pulumi.Input[bool]]:
        """
        Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        """
        return pulumi.get(self, "bootstrap")

    @bootstrap.setter
    def bootstrap(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "bootstrap", value)

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        """
        return pulumi.get(self, "ca_cert")

    @ca_cert.setter
    def ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ca_cert", value)

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> Optional[pulumi.Input[str]]:
        """
        Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        this is set you need to also set client_key.
        """
        return pulumi.get(self, "client_cert")

    @client_cert.setter
    def client_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_cert", value)

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> Optional[pulumi.Input[str]]:
        """
        Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        you need to also set client_cert.
        """
        return pulumi.get(self, "client_key")

    @client_key.setter
    def client_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_key", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The default TTL for credentials issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> Optional[pulumi.Input[bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @disable_remount.setter
    def disable_remount(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_remount", value)

    @property
    @pulumi.getter
    def local(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies if the secret backend is local only.
        """
        return pulumi.get(self, "local")

    @local.setter
    def local(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "local", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

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

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        to `consul`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def scheme(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the URL scheme to use. Defaults to `http`.
        """
        return pulumi.get(self, "scheme")

    @scheme.setter
    def scheme(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "scheme", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the Consul token to use when managing or issuing new tokens.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)


class SecretBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 bootstrap: Optional[pulumi.Input[bool]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ### Creating a standard backend resource:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.consul.SecretBackend("test",
            address="127.0.0.1:8500",
            description="Manages the Consul backend",
            path="consul",
            token="4240861b-ce3d-8530-115a-521ff070dd29")
        ```
        <!--End PulumiCodeChooser -->

        ### Creating a backend resource to bootstrap a new Consul instance:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.consul.SecretBackend("test",
            address="127.0.0.1:8500",
            bootstrap=True,
            description="Bootstrap the Consul backend",
            path="consul")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Consul secret backends can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:consul/secretBackend:SecretBackend example consul
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        :param pulumi.Input[bool] bootstrap: Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        :param pulumi.Input[str] ca_cert: CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        :param pulumi.Input[str] client_cert: Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
               this is set you need to also set client_key.
        :param pulumi.Input[str] client_key: Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
               you need to also set client_cert.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[bool] local: Specifies if the secret backend is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
               to `consul`.
        :param pulumi.Input[str] scheme: Specifies the URL scheme to use. Defaults to `http`.
        :param pulumi.Input[str] token: Specifies the Consul token to use when managing or issuing new tokens.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ### Creating a standard backend resource:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.consul.SecretBackend("test",
            address="127.0.0.1:8500",
            description="Manages the Consul backend",
            path="consul",
            token="4240861b-ce3d-8530-115a-521ff070dd29")
        ```
        <!--End PulumiCodeChooser -->

        ### Creating a backend resource to bootstrap a new Consul instance:
        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.consul.SecretBackend("test",
            address="127.0.0.1:8500",
            bootstrap=True,
            description="Bootstrap the Consul backend",
            path="consul")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Consul secret backends can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:consul/secretBackend:SecretBackend example consul
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 address: Optional[pulumi.Input[str]] = None,
                 bootstrap: Optional[pulumi.Input[bool]] = None,
                 ca_cert: Optional[pulumi.Input[str]] = None,
                 client_cert: Optional[pulumi.Input[str]] = None,
                 client_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 disable_remount: Optional[pulumi.Input[bool]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 scheme: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendArgs.__new__(SecretBackendArgs)

            if address is None and not opts.urn:
                raise TypeError("Missing required property 'address'")
            __props__.__dict__["address"] = address
            __props__.__dict__["bootstrap"] = bootstrap
            __props__.__dict__["ca_cert"] = ca_cert
            __props__.__dict__["client_cert"] = None if client_cert is None else pulumi.Output.secret(client_cert)
            __props__.__dict__["client_key"] = None if client_key is None else pulumi.Output.secret(client_key)
            __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
            __props__.__dict__["description"] = description
            __props__.__dict__["disable_remount"] = disable_remount
            __props__.__dict__["local"] = local
            __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["path"] = path
            __props__.__dict__["scheme"] = scheme
            __props__.__dict__["token"] = None if token is None else pulumi.Output.secret(token)
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientCert", "clientKey", "token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretBackend, __self__).__init__(
            'vault:consul/secretBackend:SecretBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            address: Optional[pulumi.Input[str]] = None,
            bootstrap: Optional[pulumi.Input[bool]] = None,
            ca_cert: Optional[pulumi.Input[str]] = None,
            client_cert: Optional[pulumi.Input[str]] = None,
            client_key: Optional[pulumi.Input[str]] = None,
            default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            disable_remount: Optional[pulumi.Input[bool]] = None,
            local: Optional[pulumi.Input[bool]] = None,
            max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            scheme: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None) -> 'SecretBackend':
        """
        Get an existing SecretBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] address: Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        :param pulumi.Input[bool] bootstrap: Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        :param pulumi.Input[str] ca_cert: CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        :param pulumi.Input[str] client_cert: Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
               this is set you need to also set client_key.
        :param pulumi.Input[str] client_key: Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
               you need to also set client_cert.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[bool] disable_remount: If set, opts out of mount migration on path updates.
               See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        :param pulumi.Input[bool] local: Specifies if the secret backend is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
               to `consul`.
        :param pulumi.Input[str] scheme: Specifies the URL scheme to use. Defaults to `http`.
        :param pulumi.Input[str] token: Specifies the Consul token to use when managing or issuing new tokens.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendState.__new__(_SecretBackendState)

        __props__.__dict__["address"] = address
        __props__.__dict__["bootstrap"] = bootstrap
        __props__.__dict__["ca_cert"] = ca_cert
        __props__.__dict__["client_cert"] = client_cert
        __props__.__dict__["client_key"] = client_key
        __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
        __props__.__dict__["description"] = description
        __props__.__dict__["disable_remount"] = disable_remount
        __props__.__dict__["local"] = local
        __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["scheme"] = scheme
        __props__.__dict__["token"] = token
        return SecretBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def address(self) -> pulumi.Output[str]:
        """
        Specifies the address of the Consul instance, provided as "host:port" like "127.0.0.1:8500".
        """
        return pulumi.get(self, "address")

    @property
    @pulumi.getter
    def bootstrap(self) -> pulumi.Output[Optional[bool]]:
        """
        Denotes a backend resource that is used to bootstrap the Consul ACL system. Only one resource may be used to bootstrap.
        """
        return pulumi.get(self, "bootstrap")

    @property
    @pulumi.getter(name="caCert")
    def ca_cert(self) -> pulumi.Output[Optional[str]]:
        """
        CA certificate to use when verifying Consul server certificate, must be x509 PEM encoded.
        """
        return pulumi.get(self, "ca_cert")

    @property
    @pulumi.getter(name="clientCert")
    def client_cert(self) -> pulumi.Output[Optional[str]]:
        """
        Client certificate used for Consul's TLS communication, must be x509 PEM encoded and if
        this is set you need to also set client_key.
        """
        return pulumi.get(self, "client_cert")

    @property
    @pulumi.getter(name="clientKey")
    def client_key(self) -> pulumi.Output[Optional[str]]:
        """
        Client key used for Consul's TLS communication, must be x509 PEM encoded and if this is set
        you need to also set client_cert.
        """
        return pulumi.get(self, "client_key")

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The default TTL for credentials issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="disableRemount")
    def disable_remount(self) -> pulumi.Output[Optional[bool]]:
        """
        If set, opts out of mount migration on path updates.
        See here for more info on [Mount Migration](https://www.vaultproject.io/docs/concepts/mount-migration)
        """
        return pulumi.get(self, "disable_remount")

    @property
    @pulumi.getter
    def local(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the secret backend is local only.
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

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

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The unique location this backend should be mounted at. Must not begin or end with a `/`. Defaults
        to `consul`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def scheme(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the URL scheme to use. Defaults to `http`.
        """
        return pulumi.get(self, "scheme")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the Consul token to use when managing or issuing new tokens.
        """
        return pulumi.get(self, "token")

