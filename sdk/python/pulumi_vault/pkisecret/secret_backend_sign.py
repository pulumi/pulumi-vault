# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendSignArgs', 'SecretBackendSign']

@pulumi.input_type
class SecretBackendSignArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 common_name: pulumi.Input[str],
                 csr: pulumi.Input[str],
                 alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 min_seconds_remaining: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecretBackendSign resource.
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[int] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alterative URIs
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "common_name", common_name)
        pulumi.set(__self__, "csr", csr)
        if alt_names is not None:
            pulumi.set(__self__, "alt_names", alt_names)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if exclude_cn_from_sans is not None:
            pulumi.set(__self__, "exclude_cn_from_sans", exclude_cn_from_sans)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if ip_sans is not None:
            pulumi.set(__self__, "ip_sans", ip_sans)
        if min_seconds_remaining is not None:
            pulumi.set(__self__, "min_seconds_remaining", min_seconds_remaining)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if other_sans is not None:
            pulumi.set(__self__, "other_sans", other_sans)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if uri_sans is not None:
            pulumi.set(__self__, "uri_sans", uri_sans)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The PKI secret backend the resource belongs to.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Input[str]:
        """
        CN of certificate to create
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Input[str]:
        """
        The CSR
        """
        return pulumi.get(self, "csr")

    @csr.setter
    def csr(self, value: pulumi.Input[str]):
        pulumi.set(self, "csr", value)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alternative names
        """
        return pulumi.get(self, "alt_names")

    @alt_names.setter
    def alt_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alt_names", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter(name="excludeCnFromSans")
    def exclude_cn_from_sans(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to exclude CN from SANs
        """
        return pulumi.get(self, "exclude_cn_from_sans")

    @exclude_cn_from_sans.setter
    def exclude_cn_from_sans(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_cn_from_sans", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of data
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="ipSans")
    def ip_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alternative IPs
        """
        return pulumi.get(self, "ip_sans")

    @ip_sans.setter
    def ip_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_sans", value)

    @property
    @pulumi.getter(name="minSecondsRemaining")
    def min_seconds_remaining(self) -> Optional[pulumi.Input[int]]:
        """
        Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        """
        return pulumi.get(self, "min_seconds_remaining")

    @min_seconds_remaining.setter
    def min_seconds_remaining(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_seconds_remaining", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role to create the certificate against
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="otherSans")
    def other_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @other_sans.setter
    def other_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "other_sans", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Time to live
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="uriSans")
    def uri_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alterative URIs
        """
        return pulumi.get(self, "uri_sans")

    @uri_sans.setter
    def uri_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uri_sans", value)


@pulumi.input_type
class _SecretBackendSignState:
    def __init__(__self__, *,
                 alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 ca_chains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 certificate: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 expiration: Optional[pulumi.Input[int]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 issuing_ca: Optional[pulumi.Input[str]] = None,
                 min_seconds_remaining: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 serial: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecretBackendSign resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ca_chains: The CA chain
        :param pulumi.Input[str] certificate: The certificate
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[int] expiration: The expiration date of the certificate in unix epoch format
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[str] issuing_ca: The issuing CA
        :param pulumi.Input[int] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] serial: The serial
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alterative URIs
        """
        if alt_names is not None:
            pulumi.set(__self__, "alt_names", alt_names)
        if auto_renew is not None:
            pulumi.set(__self__, "auto_renew", auto_renew)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if ca_chains is not None:
            pulumi.set(__self__, "ca_chains", ca_chains)
        if certificate is not None:
            pulumi.set(__self__, "certificate", certificate)
        if common_name is not None:
            pulumi.set(__self__, "common_name", common_name)
        if csr is not None:
            pulumi.set(__self__, "csr", csr)
        if exclude_cn_from_sans is not None:
            pulumi.set(__self__, "exclude_cn_from_sans", exclude_cn_from_sans)
        if expiration is not None:
            pulumi.set(__self__, "expiration", expiration)
        if format is not None:
            pulumi.set(__self__, "format", format)
        if ip_sans is not None:
            pulumi.set(__self__, "ip_sans", ip_sans)
        if issuing_ca is not None:
            pulumi.set(__self__, "issuing_ca", issuing_ca)
        if min_seconds_remaining is not None:
            pulumi.set(__self__, "min_seconds_remaining", min_seconds_remaining)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if other_sans is not None:
            pulumi.set(__self__, "other_sans", other_sans)
        if serial is not None:
            pulumi.set(__self__, "serial", serial)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)
        if uri_sans is not None:
            pulumi.set(__self__, "uri_sans", uri_sans)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alternative names
        """
        return pulumi.get(self, "alt_names")

    @alt_names.setter
    def alt_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "alt_names", value)

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> Optional[pulumi.Input[bool]]:
        """
        If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        """
        return pulumi.get(self, "auto_renew")

    @auto_renew.setter
    def auto_renew(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "auto_renew", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The PKI secret backend the resource belongs to.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="caChains")
    def ca_chains(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The CA chain
        """
        return pulumi.get(self, "ca_chains")

    @ca_chains.setter
    def ca_chains(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ca_chains", value)

    @property
    @pulumi.getter
    def certificate(self) -> Optional[pulumi.Input[str]]:
        """
        The certificate
        """
        return pulumi.get(self, "certificate")

    @certificate.setter
    def certificate(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "certificate", value)

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> Optional[pulumi.Input[str]]:
        """
        CN of certificate to create
        """
        return pulumi.get(self, "common_name")

    @common_name.setter
    def common_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "common_name", value)

    @property
    @pulumi.getter
    def csr(self) -> Optional[pulumi.Input[str]]:
        """
        The CSR
        """
        return pulumi.get(self, "csr")

    @csr.setter
    def csr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "csr", value)

    @property
    @pulumi.getter(name="excludeCnFromSans")
    def exclude_cn_from_sans(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to exclude CN from SANs
        """
        return pulumi.get(self, "exclude_cn_from_sans")

    @exclude_cn_from_sans.setter
    def exclude_cn_from_sans(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "exclude_cn_from_sans", value)

    @property
    @pulumi.getter
    def expiration(self) -> Optional[pulumi.Input[int]]:
        """
        The expiration date of the certificate in unix epoch format
        """
        return pulumi.get(self, "expiration")

    @expiration.setter
    def expiration(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "expiration", value)

    @property
    @pulumi.getter
    def format(self) -> Optional[pulumi.Input[str]]:
        """
        The format of data
        """
        return pulumi.get(self, "format")

    @format.setter
    def format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "format", value)

    @property
    @pulumi.getter(name="ipSans")
    def ip_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alternative IPs
        """
        return pulumi.get(self, "ip_sans")

    @ip_sans.setter
    def ip_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_sans", value)

    @property
    @pulumi.getter(name="issuingCa")
    def issuing_ca(self) -> Optional[pulumi.Input[str]]:
        """
        The issuing CA
        """
        return pulumi.get(self, "issuing_ca")

    @issuing_ca.setter
    def issuing_ca(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuing_ca", value)

    @property
    @pulumi.getter(name="minSecondsRemaining")
    def min_seconds_remaining(self) -> Optional[pulumi.Input[int]]:
        """
        Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        """
        return pulumi.get(self, "min_seconds_remaining")

    @min_seconds_remaining.setter
    def min_seconds_remaining(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "min_seconds_remaining", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role to create the certificate against
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="otherSans")
    def other_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @other_sans.setter
    def other_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "other_sans", value)

    @property
    @pulumi.getter
    def serial(self) -> Optional[pulumi.Input[str]]:
        """
        The serial
        """
        return pulumi.get(self, "serial")

    @serial.setter
    def serial(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "serial", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Time to live
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ttl", value)

    @property
    @pulumi.getter(name="uriSans")
    def uri_sans(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of alterative URIs
        """
        return pulumi.get(self, "uri_sans")

    @uri_sans.setter
    def uri_sans(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "uri_sans", value)


class SecretBackendSign(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 min_seconds_remaining: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Create a SecretBackendSign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[int] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alterative URIs
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendSignArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SecretBackendSign resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SecretBackendSignArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendSignArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 csr: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 min_seconds_remaining: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendSignArgs.__new__(SecretBackendSignArgs)

            __props__.__dict__["alt_names"] = alt_names
            __props__.__dict__["auto_renew"] = auto_renew
            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            if common_name is None and not opts.urn:
                raise TypeError("Missing required property 'common_name'")
            __props__.__dict__["common_name"] = common_name
            if csr is None and not opts.urn:
                raise TypeError("Missing required property 'csr'")
            __props__.__dict__["csr"] = csr
            __props__.__dict__["exclude_cn_from_sans"] = exclude_cn_from_sans
            __props__.__dict__["format"] = format
            __props__.__dict__["ip_sans"] = ip_sans
            __props__.__dict__["min_seconds_remaining"] = min_seconds_remaining
            __props__.__dict__["name"] = name
            __props__.__dict__["other_sans"] = other_sans
            __props__.__dict__["ttl"] = ttl
            __props__.__dict__["uri_sans"] = uri_sans
            __props__.__dict__["ca_chains"] = None
            __props__.__dict__["certificate"] = None
            __props__.__dict__["expiration"] = None
            __props__.__dict__["issuing_ca"] = None
            __props__.__dict__["serial"] = None
        super(SecretBackendSign, __self__).__init__(
            'vault:pkiSecret/secretBackendSign:SecretBackendSign',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alt_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            ca_chains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            csr: Optional[pulumi.Input[str]] = None,
            exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
            expiration: Optional[pulumi.Input[int]] = None,
            format: Optional[pulumi.Input[str]] = None,
            ip_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            issuing_ca: Optional[pulumi.Input[str]] = None,
            min_seconds_remaining: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            serial: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecretBackendSign':
        """
        Get an existing SecretBackendSign resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ca_chains: The CA chain
        :param pulumi.Input[str] certificate: The certificate
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[int] expiration: The expiration date of the certificate in unix epoch format
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[str] issuing_ca: The issuing CA
        :param pulumi.Input[int] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[Sequence[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] serial: The serial
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[Sequence[pulumi.Input[str]]] uri_sans: List of alterative URIs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendSignState.__new__(_SecretBackendSignState)

        __props__.__dict__["alt_names"] = alt_names
        __props__.__dict__["auto_renew"] = auto_renew
        __props__.__dict__["backend"] = backend
        __props__.__dict__["ca_chains"] = ca_chains
        __props__.__dict__["certificate"] = certificate
        __props__.__dict__["common_name"] = common_name
        __props__.__dict__["csr"] = csr
        __props__.__dict__["exclude_cn_from_sans"] = exclude_cn_from_sans
        __props__.__dict__["expiration"] = expiration
        __props__.__dict__["format"] = format
        __props__.__dict__["ip_sans"] = ip_sans
        __props__.__dict__["issuing_ca"] = issuing_ca
        __props__.__dict__["min_seconds_remaining"] = min_seconds_remaining
        __props__.__dict__["name"] = name
        __props__.__dict__["other_sans"] = other_sans
        __props__.__dict__["serial"] = serial
        __props__.__dict__["ttl"] = ttl
        __props__.__dict__["uri_sans"] = uri_sans
        return SecretBackendSign(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alternative names
        """
        return pulumi.get(self, "alt_names")

    @property
    @pulumi.getter(name="autoRenew")
    def auto_renew(self) -> pulumi.Output[Optional[bool]]:
        """
        If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        """
        return pulumi.get(self, "auto_renew")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The PKI secret backend the resource belongs to.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="caChains")
    def ca_chains(self) -> pulumi.Output[Sequence[str]]:
        """
        The CA chain
        """
        return pulumi.get(self, "ca_chains")

    @property
    @pulumi.getter
    def certificate(self) -> pulumi.Output[str]:
        """
        The certificate
        """
        return pulumi.get(self, "certificate")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[str]:
        """
        CN of certificate to create
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def csr(self) -> pulumi.Output[str]:
        """
        The CSR
        """
        return pulumi.get(self, "csr")

    @property
    @pulumi.getter(name="excludeCnFromSans")
    def exclude_cn_from_sans(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to exclude CN from SANs
        """
        return pulumi.get(self, "exclude_cn_from_sans")

    @property
    @pulumi.getter
    def expiration(self) -> pulumi.Output[int]:
        """
        The expiration date of the certificate in unix epoch format
        """
        return pulumi.get(self, "expiration")

    @property
    @pulumi.getter
    def format(self) -> pulumi.Output[Optional[str]]:
        """
        The format of data
        """
        return pulumi.get(self, "format")

    @property
    @pulumi.getter(name="ipSans")
    def ip_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alternative IPs
        """
        return pulumi.get(self, "ip_sans")

    @property
    @pulumi.getter(name="issuingCa")
    def issuing_ca(self) -> pulumi.Output[str]:
        """
        The issuing CA
        """
        return pulumi.get(self, "issuing_ca")

    @property
    @pulumi.getter(name="minSecondsRemaining")
    def min_seconds_remaining(self) -> pulumi.Output[Optional[int]]:
        """
        Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        """
        return pulumi.get(self, "min_seconds_remaining")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the role to create the certificate against
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="otherSans")
    def other_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @property
    @pulumi.getter
    def serial(self) -> pulumi.Output[str]:
        """
        The serial
        """
        return pulumi.get(self, "serial")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Time to live
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="uriSans")
    def uri_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of alterative URIs
        """
        return pulumi.get(self, "uri_sans")

