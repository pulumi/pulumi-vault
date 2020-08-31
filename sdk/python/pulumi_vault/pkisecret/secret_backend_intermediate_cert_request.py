# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendIntermediateCertRequest']


class SecretBackendIntermediateCertRequest(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alt_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 common_name: Optional[pulumi.Input[str]] = None,
                 country: Optional[pulumi.Input[str]] = None,
                 exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
                 format: Optional[pulumi.Input[str]] = None,
                 ip_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 key_bits: Optional[pulumi.Input[float]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 locality: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 other_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 ou: Optional[pulumi.Input[str]] = None,
                 postal_code: Optional[pulumi.Input[str]] = None,
                 private_key_format: Optional[pulumi.Input[str]] = None,
                 province: Optional[pulumi.Input[str]] = None,
                 street_address: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 uri_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a SecretBackendIntermediateCertRequest resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[List[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[float] key_bits: The number of bits to use
        :param pulumi.Input[str] key_type: The desired key type
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[List[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] type: Type of intermediate to create. Must be either \"exported\" or \"internal\"
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
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if common_name is None:
                raise TypeError("Missing required property 'common_name'")
            __props__['common_name'] = common_name
            __props__['country'] = country
            __props__['exclude_cn_from_sans'] = exclude_cn_from_sans
            __props__['format'] = format
            __props__['ip_sans'] = ip_sans
            __props__['key_bits'] = key_bits
            __props__['key_type'] = key_type
            __props__['locality'] = locality
            __props__['organization'] = organization
            __props__['other_sans'] = other_sans
            __props__['ou'] = ou
            __props__['postal_code'] = postal_code
            __props__['private_key_format'] = private_key_format
            __props__['province'] = province
            __props__['street_address'] = street_address
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['uri_sans'] = uri_sans
            __props__['csr'] = None
            __props__['private_key'] = None
            __props__['private_key_type'] = None
        super(SecretBackendIntermediateCertRequest, __self__).__init__(
            'vault:pkisecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alt_names: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            common_name: Optional[pulumi.Input[str]] = None,
            country: Optional[pulumi.Input[str]] = None,
            csr: Optional[pulumi.Input[str]] = None,
            exclude_cn_from_sans: Optional[pulumi.Input[bool]] = None,
            format: Optional[pulumi.Input[str]] = None,
            ip_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            key_bits: Optional[pulumi.Input[float]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            locality: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            other_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            ou: Optional[pulumi.Input[str]] = None,
            postal_code: Optional[pulumi.Input[str]] = None,
            private_key: Optional[pulumi.Input[str]] = None,
            private_key_format: Optional[pulumi.Input[str]] = None,
            private_key_type: Optional[pulumi.Input[str]] = None,
            province: Optional[pulumi.Input[str]] = None,
            street_address: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            uri_sans: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None) -> 'SecretBackendIntermediateCertRequest':
        """
        Get an existing SecretBackendIntermediateCertRequest resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[List[pulumi.Input[str]]] alt_names: List of alternative names
        :param pulumi.Input[str] backend: The PKI secret backend the resource belongs to.
        :param pulumi.Input[str] common_name: CN of intermediate to create
        :param pulumi.Input[str] country: The country
        :param pulumi.Input[str] csr: The CSR
        :param pulumi.Input[bool] exclude_cn_from_sans: Flag to exclude CN from SANs
        :param pulumi.Input[str] format: The format of data
        :param pulumi.Input[List[pulumi.Input[str]]] ip_sans: List of alternative IPs
        :param pulumi.Input[float] key_bits: The number of bits to use
        :param pulumi.Input[str] key_type: The desired key type
        :param pulumi.Input[str] locality: The locality
        :param pulumi.Input[str] organization: The organization
        :param pulumi.Input[List[pulumi.Input[str]]] other_sans: List of other SANs
        :param pulumi.Input[str] ou: The organization unit
        :param pulumi.Input[str] postal_code: The postal code
        :param pulumi.Input[str] private_key: The private key
        :param pulumi.Input[str] private_key_format: The private key format
        :param pulumi.Input[str] private_key_type: The private key type
        :param pulumi.Input[str] province: The province
        :param pulumi.Input[str] street_address: The street address
        :param pulumi.Input[str] type: Type of intermediate to create. Must be either \"exported\" or \"internal\"
        :param pulumi.Input[List[pulumi.Input[str]]] uri_sans: List of alternative URIs
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alt_names"] = alt_names
        __props__["backend"] = backend
        __props__["common_name"] = common_name
        __props__["country"] = country
        __props__["csr"] = csr
        __props__["exclude_cn_from_sans"] = exclude_cn_from_sans
        __props__["format"] = format
        __props__["ip_sans"] = ip_sans
        __props__["key_bits"] = key_bits
        __props__["key_type"] = key_type
        __props__["locality"] = locality
        __props__["organization"] = organization
        __props__["other_sans"] = other_sans
        __props__["ou"] = ou
        __props__["postal_code"] = postal_code
        __props__["private_key"] = private_key
        __props__["private_key_format"] = private_key_format
        __props__["private_key_type"] = private_key_type
        __props__["province"] = province
        __props__["street_address"] = street_address
        __props__["type"] = type
        __props__["uri_sans"] = uri_sans
        return SecretBackendIntermediateCertRequest(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="altNames")
    def alt_names(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of alternative names
        """
        return pulumi.get(self, "alt_names")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The PKI secret backend the resource belongs to.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> pulumi.Output[str]:
        """
        CN of intermediate to create
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter
    def country(self) -> pulumi.Output[Optional[str]]:
        """
        The country
        """
        return pulumi.get(self, "country")

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
    @pulumi.getter(name="keyBits")
    def key_bits(self) -> pulumi.Output[Optional[float]]:
        """
        The number of bits to use
        """
        return pulumi.get(self, "key_bits")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[Optional[str]]:
        """
        The desired key type
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def locality(self) -> pulumi.Output[Optional[str]]:
        """
        The locality
        """
        return pulumi.get(self, "locality")

    @property
    @pulumi.getter
    def organization(self) -> pulumi.Output[Optional[str]]:
        """
        The organization
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter(name="otherSans")
    def other_sans(self) -> pulumi.Output[Optional[List[str]]]:
        """
        List of other SANs
        """
        return pulumi.get(self, "other_sans")

    @property
    @pulumi.getter
    def ou(self) -> pulumi.Output[Optional[str]]:
        """
        The organization unit
        """
        return pulumi.get(self, "ou")

    @property
    @pulumi.getter(name="postalCode")
    def postal_code(self) -> pulumi.Output[Optional[str]]:
        """
        The postal code
        """
        return pulumi.get(self, "postal_code")

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
    @pulumi.getter
    def province(self) -> pulumi.Output[Optional[str]]:
        """
        The province
        """
        return pulumi.get(self, "province")

    @property
    @pulumi.getter(name="streetAddress")
    def street_address(self) -> pulumi.Output[Optional[str]]:
        """
        The street address
        """
        return pulumi.get(self, "street_address")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of intermediate to create. Must be either \"exported\" or \"internal\"
        """
        return pulumi.get(self, "type")

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

