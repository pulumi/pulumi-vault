# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendConfigUrlsArgs', 'SecretBackendConfigUrls']

@pulumi.input_type
class SecretBackendConfigUrlsArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 crl_distribution_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_templating: Optional[pulumi.Input[bool]] = None,
                 issuing_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecretBackendConfigUrls resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[bool] enable_templating: Specifies that templating of AIA fields is allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ocsp_servers: Specifies the URL values for the OCSP Servers field.
        """
        pulumi.set(__self__, "backend", backend)
        if crl_distribution_points is not None:
            pulumi.set(__self__, "crl_distribution_points", crl_distribution_points)
        if enable_templating is not None:
            pulumi.set(__self__, "enable_templating", enable_templating)
        if issuing_certificates is not None:
            pulumi.set(__self__, "issuing_certificates", issuing_certificates)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if ocsp_servers is not None:
            pulumi.set(__self__, "ocsp_servers", ocsp_servers)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="crlDistributionPoints")
    def crl_distribution_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the CRL Distribution Points field.
        """
        return pulumi.get(self, "crl_distribution_points")

    @crl_distribution_points.setter
    def crl_distribution_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "crl_distribution_points", value)

    @property
    @pulumi.getter(name="enableTemplating")
    def enable_templating(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies that templating of AIA fields is allowed.
        """
        return pulumi.get(self, "enable_templating")

    @enable_templating.setter
    def enable_templating(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_templating", value)

    @property
    @pulumi.getter(name="issuingCertificates")
    def issuing_certificates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the Issuing Certificate field.
        """
        return pulumi.get(self, "issuing_certificates")

    @issuing_certificates.setter
    def issuing_certificates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "issuing_certificates", value)

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
    @pulumi.getter(name="ocspServers")
    def ocsp_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the OCSP Servers field.
        """
        return pulumi.get(self, "ocsp_servers")

    @ocsp_servers.setter
    def ocsp_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ocsp_servers", value)


@pulumi.input_type
class _SecretBackendConfigUrlsState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 crl_distribution_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_templating: Optional[pulumi.Input[bool]] = None,
                 issuing_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecretBackendConfigUrls resources.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[bool] enable_templating: Specifies that templating of AIA fields is allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ocsp_servers: Specifies the URL values for the OCSP Servers field.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if crl_distribution_points is not None:
            pulumi.set(__self__, "crl_distribution_points", crl_distribution_points)
        if enable_templating is not None:
            pulumi.set(__self__, "enable_templating", enable_templating)
        if issuing_certificates is not None:
            pulumi.set(__self__, "issuing_certificates", issuing_certificates)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if ocsp_servers is not None:
            pulumi.set(__self__, "ocsp_servers", ocsp_servers)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="crlDistributionPoints")
    def crl_distribution_points(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the CRL Distribution Points field.
        """
        return pulumi.get(self, "crl_distribution_points")

    @crl_distribution_points.setter
    def crl_distribution_points(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "crl_distribution_points", value)

    @property
    @pulumi.getter(name="enableTemplating")
    def enable_templating(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies that templating of AIA fields is allowed.
        """
        return pulumi.get(self, "enable_templating")

    @enable_templating.setter
    def enable_templating(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_templating", value)

    @property
    @pulumi.getter(name="issuingCertificates")
    def issuing_certificates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the Issuing Certificate field.
        """
        return pulumi.get(self, "issuing_certificates")

    @issuing_certificates.setter
    def issuing_certificates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "issuing_certificates", value)

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
    @pulumi.getter(name="ocspServers")
    def ocsp_servers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the URL values for the OCSP Servers field.
        """
        return pulumi.get(self, "ocsp_servers")

    @ocsp_servers.setter
    def ocsp_servers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ocsp_servers", value)


class SecretBackendConfigUrls(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 crl_distribution_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_templating: Optional[pulumi.Input[bool]] = None,
                 issuing_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        root = vault.Mount("root",
            path="pki-root",
            type="pki",
            description="root PKI",
            default_lease_ttl_seconds=8640000,
            max_lease_ttl_seconds=8640000)
        example = vault.pki_secret.SecretBackendConfigUrls("example",
            backend=root.path,
            issuing_certificates=["http://127.0.0.1:8200/v1/pki/ca"])
        ```

        ## Import

        The PKI config URLs can be imported using the resource's `id`.
        In the case of the example above the `id` would be `pki-root/config/urls`,
        where the `pki-root` component is the resource's `backend`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls example pki-root/config/urls
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[bool] enable_templating: Specifies that templating of AIA fields is allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ocsp_servers: Specifies the URL values for the OCSP Servers field.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendConfigUrlsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server endpoints that will be encoded into issued certificates.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        root = vault.Mount("root",
            path="pki-root",
            type="pki",
            description="root PKI",
            default_lease_ttl_seconds=8640000,
            max_lease_ttl_seconds=8640000)
        example = vault.pki_secret.SecretBackendConfigUrls("example",
            backend=root.path,
            issuing_certificates=["http://127.0.0.1:8200/v1/pki/ca"])
        ```

        ## Import

        The PKI config URLs can be imported using the resource's `id`.
        In the case of the example above the `id` would be `pki-root/config/urls`,
        where the `pki-root` component is the resource's `backend`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls example pki-root/config/urls
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendConfigUrlsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendConfigUrlsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 crl_distribution_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 enable_templating: Optional[pulumi.Input[bool]] = None,
                 issuing_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ocsp_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendConfigUrlsArgs.__new__(SecretBackendConfigUrlsArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["crl_distribution_points"] = crl_distribution_points
            __props__.__dict__["enable_templating"] = enable_templating
            __props__.__dict__["issuing_certificates"] = issuing_certificates
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["ocsp_servers"] = ocsp_servers
        super(SecretBackendConfigUrls, __self__).__init__(
            'vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            crl_distribution_points: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            enable_templating: Optional[pulumi.Input[bool]] = None,
            issuing_certificates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            ocsp_servers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecretBackendConfigUrls':
        """
        Get an existing SecretBackendConfigUrls resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] crl_distribution_points: Specifies the URL values for the CRL Distribution Points field.
        :param pulumi.Input[bool] enable_templating: Specifies that templating of AIA fields is allowed.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] issuing_certificates: Specifies the URL values for the Issuing Certificate field.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ocsp_servers: Specifies the URL values for the OCSP Servers field.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendConfigUrlsState.__new__(_SecretBackendConfigUrlsState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["crl_distribution_points"] = crl_distribution_points
        __props__.__dict__["enable_templating"] = enable_templating
        __props__.__dict__["issuing_certificates"] = issuing_certificates
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["ocsp_servers"] = ocsp_servers
        return SecretBackendConfigUrls(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="crlDistributionPoints")
    def crl_distribution_points(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the URL values for the CRL Distribution Points field.
        """
        return pulumi.get(self, "crl_distribution_points")

    @property
    @pulumi.getter(name="enableTemplating")
    def enable_templating(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies that templating of AIA fields is allowed.
        """
        return pulumi.get(self, "enable_templating")

    @property
    @pulumi.getter(name="issuingCertificates")
    def issuing_certificates(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the URL values for the Issuing Certificate field.
        """
        return pulumi.get(self, "issuing_certificates")

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
    @pulumi.getter(name="ocspServers")
    def ocsp_servers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specifies the URL values for the OCSP Servers field.
        """
        return pulumi.get(self, "ocsp_servers")

