# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendCert']


class SecretBackendCert(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alt_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 auto_renew: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 min_seconds_remaining: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 private_key_format: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SecretBackendCert resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[List[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[float] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[List[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[List[pulumi.Input[str]]] uri_sans: List of alternative URIs
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['alt_names'] = alt_names
            __props__['auto_renew'] = auto_renew
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if common_name is None:
                raise TypeError("Missing required property 'common_name'")
            __props__['common_name'] = common_name
            __props__['exclude_cn_from_sans'] = exclude_cn_from_sans
            __props__['format'] = format
            __props__['ip_sans'] = ip_sans
            __props__['min_seconds_remaining'] = min_seconds_remaining
            __props__['name'] = name
            __props__['other_sans'] = other_sans
            __props__['private_key_format'] = private_key_format
            __props__['ttl'] = ttl
            __props__['uri_sans'] = uri_sans
            __props__['ca_chain'] = None
            __props__['certificate'] = None
            __props__['expiration'] = None
            __props__['issuing_ca'] = None
            __props__['private_key'] = None
            __props__['private_key_type'] = None
            __props__['serial_number'] = None
        super(SecretBackendCert, __self__).__init__(
            'vault:pkisecret/secretBackendCert:SecretBackendCert',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alt_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            auto_renew: Optional[pulumi.Input[bool]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            ca_chain: Optional[pulumi.Input[str]] = None,
            certificate: Optional[pulumi.Input[str]] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
            expiration: Optional[pulumi.Input[float]] = None,
            format: Optional[pulumi.Input[str]] = None,
            ip_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            issuing_ca: Optional[pulumi.Input[str]] = None,
            min_seconds_remaining: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            other_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_format: Optional[pulumi.Input[str]] = None,
            private_key_type: Optional[pulumi.Input[str]] = None,
            serial_number: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            uri_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'SecretBackendCert':
        """
        Get an existing SecretBackendCert resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[bool] auto_renew: If set to `true`, certs will be renewed if the expiration is within `min_seconds_remaining`. Default `false`
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] ca_chain: The CA chain
        :param pulumi.Input[str] certificate: The certificate
        :param pulumi.Input[str] common_name: CN of certificate to create
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[float] expiration: The expiration date of the certificate in unix epoch format
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[List[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[str] issuing_ca: The issuing CA
        :param pulumi.Input[float] min_seconds_remaining: Generate a new certificate when the expiration is within this number of seconds, default is 604800 (7 days)
        :param pulumi.Input[str] name: Name of the role to create the certificate against
        :param pulumi.Input[List[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] private_key: The private key
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] private_key_type: The private key type
        :param pulumi.Input[str] serial_number: The serial number
        :param pulumi.Input[str] ttl: Time to live
        :param pulumi.Input[List[pulumi.Input[str]]] uri_sans: List of alternative URIs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alt_names"] = alt_names
        __props__["auto_renew"] = auto_renew
        __props__["backend"] = backend
        __props__["ca_chain"] = ca_chain
        __props__["certificate"] = certificate
        __props__["common_name"] = common_name
        __props__["exclude_cn_from_sans"] = exclude_cn_from_sans
        __props__["expiration"] = expiration
        __props__["format"] = format
        __props__["ip_sans"] = ip_sans
        __props__["issuing_ca"] = issuing_ca
        __props__["min_seconds_remaining"] = min_seconds_remaining
        __props__["name"] = name
        __props__["other_sans"] = other_sans
        __props__["private_key"] = private_key
        __props__["private_key_format"] = private_key_format
        __props__["private_key_type"] = private_key_type
        __props__["serial_number"] = serial_number
        __props__["ttl"] = ttl
        __props__["uri_sans"] = uri_sans
        return SecretBackendCert(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> pulumi.Output[Optional[List[str]]]:
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
    @pulumi.getter(name="caChain")
    def ca_chain(self) -> pulumi.Output[str]:
        """
        The CA chain
        """
        return pulumi.get(self, "ca_chain")

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
    @pulumi.getter(name="excludeCnFromSans")
    def exclude_cn_from_sans(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to exclude CN from SANs
        """
        return pulumi.get(self, "exclude_cn_from_sans")

    @property
    @pulumi.getter
    def expiration(self) -> pulumi.Output[float]:
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
    def ip_sans(self) -> pulumi.Output[Optional[List[str]]]:
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
    def min_seconds_remaining(self) -> pulumi.Output[Optional[float]]:
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
    def other_sans(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[str]:
        """
        The private key
        """
        return pulumi.get(self, "private_key")

    @property
    @pulumi.getter(name="privateKeyFormat")
    def private_key_format(self) -> pulumi.Output[Optional[str]]:
        """
        The private key format
        """
        return pulumi.get(self, "private_key_format")

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> pulumi.Output[str]:
        """
        The private key type
        """
        return pulumi.get(self, "private_key_type")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output[str]:
        """
        The serial number
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        Time to live
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="uriSans")
    def uri_sans(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of alternative URIs
        """
        return pulumi.get(self, "uri_sans")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

