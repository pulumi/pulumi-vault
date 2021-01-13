# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendRole']


class SecretBackendRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_any_name: Optional[pulumi.Input[bool]] = None,
                 allow_bare_domains: Optional[pulumi.Input[bool]] = None,
                 allow_glob_domains: Optional[pulumi.Input[bool]] = None,
                 allow_ip_sans: Optional[pulumi.Input[bool]] = None,
                 allow_localhost: Optional[pulumi.Input[bool]] = None,
                 allow_subdomains: Optional[pulumi.Input[bool]] = None,
                 allowed_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 allowed_uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 basic_constraints_valid_for_non_ca: Optional[pulumi.Input[bool]] = None,
                 client_flag: Optional[pulumi.Input[bool]] = None,
                 code_signing_flag: Optional[pulumi.Input[bool]] = None,
                 countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 email_protection_flag: Optional[pulumi.Input[bool]] = None,
                 enforce_hostnames: Optional[pulumi.Input[bool]] = None,
                 ext_key_usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 generate_lease: Optional[pulumi.Input[bool]] = None,
                 key_bits: Optional[pulumi.Input[int]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 key_usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 localities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 no_store: Optional[pulumi.Input[bool]] = None,
                 not_before_duration: Optional[pulumi.Input[str]] = None,
                 organization_unit: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 organizations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy_identifiers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 postal_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 provinces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 require_cn: Optional[pulumi.Input[bool]] = None,
                 server_flag: Optional[pulumi.Input[bool]] = None,
                 street_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 use_csr_common_name: Optional[pulumi.Input[bool]] = None,
                 use_csr_sans: Optional[pulumi.Input[bool]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a role on an PKI Secret Backend for Vault.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        pki = vault.pki_secret.SecretBackend("pki",
            default_lease_ttl_seconds=3600,
            max_lease_ttl_seconds=86400,
            path="%s")
        role = vault.pki_secret.SecretBackendRole("role", backend=pki.path)
        ```

        ## Import

        PKI secret backend roles can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:pkiSecret/secretBackendRole:SecretBackendRole role pki/roles/my_role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_any_name: Flag to allow any name
        :param pulumi.Input[bool] allow_bare_domains: Flag to allow certificates matching the actual domain
        :param pulumi.Input[bool] allow_glob_domains: Flag to allow names containing glob patterns.
        :param pulumi.Input[bool] allow_ip_sans: Flag to allow IP SANs
        :param pulumi.Input[bool] allow_localhost: Flag to allow certificates for localhost
        :param pulumi.Input[bool] allow_subdomains: Flag to allow certificates matching subdomains
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_domains: List of allowed domains for certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_other_sans: Defines allowed custom SANs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uri_sans: Defines allowed URI SANs
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] basic_constraints_valid_for_non_ca: Flag to mark basic constraints valid when issuing non-CA certificates
        :param pulumi.Input[bool] client_flag: Flag to specify certificates for client use
        :param pulumi.Input[bool] code_signing_flag: Flag to specify certificates for code signing use
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: The country of generated certificates
        :param pulumi.Input[bool] email_protection_flag: Flag to specify certificates for email protection use
        :param pulumi.Input[bool] enforce_hostnames: Flag to allow only valid host names
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ext_key_usages: Specify the allowed extended key usage constraint on issued certificates
        :param pulumi.Input[bool] generate_lease: Flag to generate leases with certificates
        :param pulumi.Input[int] key_bits: The number of bits of generated keys
        :param pulumi.Input[str] key_type: The type of generated keys
        :param pulumi.Input[Sequence[pulumi.Input[str]]] key_usages: Specify the allowed key usage constraint on issued certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] localities: The locality of generated certificates
        :param pulumi.Input[str] max_ttl: The maximum TTL
        :param pulumi.Input[str] name: The name to identify this role within the backend. Must be unique within the backend.
        :param pulumi.Input[bool] no_store: Flag to not store certificates in the storage backend
        :param pulumi.Input[str] not_before_duration: Specifies the duration by which to backdate the NotBefore property.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organization_unit: The organization unit of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organizations: The organization of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_identifiers: Specify the list of allowed policies IODs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] postal_codes: The postal code of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provinces: The province of generated certificates
        :param pulumi.Input[bool] require_cn: Flag to force CN usage
        :param pulumi.Input[bool] server_flag: Flag to specify certificates for server use
        :param pulumi.Input[Sequence[pulumi.Input[str]]] street_addresses: The street address of generated certificates
        :param pulumi.Input[str] ttl: The TTL
        :param pulumi.Input[bool] use_csr_common_name: Flag to use the CN in the CSR
        :param pulumi.Input[bool] use_csr_sans: Flag to use the SANs in the CSR
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

            __props__['allow_any_name'] = allow_any_name
            __props__['allow_bare_domains'] = allow_bare_domains
            __props__['allow_glob_domains'] = allow_glob_domains
            __props__['allow_ip_sans'] = allow_ip_sans
            __props__['allow_localhost'] = allow_localhost
            __props__['allow_subdomains'] = allow_subdomains
            __props__['allowed_domains'] = allowed_domains
            __props__['allowed_other_sans'] = allowed_other_sans
            __props__['allowed_uri_sans'] = allowed_uri_sans
            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['basic_constraints_valid_for_non_ca'] = basic_constraints_valid_for_non_ca
            __props__['client_flag'] = client_flag
            __props__['code_signing_flag'] = code_signing_flag
            __props__['countries'] = countries
            __props__['email_protection_flag'] = email_protection_flag
            __props__['enforce_hostnames'] = enforce_hostnames
            __props__['ext_key_usages'] = ext_key_usages
            __props__['generate_lease'] = generate_lease
            __props__['key_bits'] = key_bits
            __props__['key_type'] = key_type
            __props__['key_usages'] = key_usages
            __props__['localities'] = localities
            __props__['max_ttl'] = max_ttl
            __props__['name'] = name
            __props__['no_store'] = no_store
            __props__['not_before_duration'] = not_before_duration
            __props__['organization_unit'] = organization_unit
            __props__['organizations'] = organizations
            __props__['policy_identifiers'] = policy_identifiers
            __props__['postal_codes'] = postal_codes
            __props__['provinces'] = provinces
            __props__['require_cn'] = require_cn
            __props__['server_flag'] = server_flag
            __props__['street_addresses'] = street_addresses
            __props__['ttl'] = ttl
            __props__['use_csr_common_name'] = use_csr_common_name
            __props__['use_csr_sans'] = use_csr_sans
        super(SecretBackendRole, __self__).__init__(
            'vault:pkiSecret/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_any_name: Optional[pulumi.Input[bool]] = None,
            allow_bare_domains: Optional[pulumi.Input[bool]] = None,
            allow_glob_domains: Optional[pulumi.Input[bool]] = None,
            allow_ip_sans: Optional[pulumi.Input[bool]] = None,
            allow_localhost: Optional[pulumi.Input[bool]] = None,
            allow_subdomains: Optional[pulumi.Input[bool]] = None,
            allowed_domains: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allowed_other_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            allowed_uri_sans: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            basic_constraints_valid_for_non_ca: Optional[pulumi.Input[bool]] = None,
            client_flag: Optional[pulumi.Input[bool]] = None,
            code_signing_flag: Optional[pulumi.Input[bool]] = None,
            countries: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            email_protection_flag: Optional[pulumi.Input[bool]] = None,
            enforce_hostnames: Optional[pulumi.Input[bool]] = None,
            ext_key_usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            generate_lease: Optional[pulumi.Input[bool]] = None,
            key_bits: Optional[pulumi.Input[int]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            key_usages: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            localities: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            max_ttl: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            no_store: Optional[pulumi.Input[bool]] = None,
            not_before_duration: Optional[pulumi.Input[str]] = None,
            organization_unit: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            organizations: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            policy_identifiers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            postal_codes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            provinces: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            require_cn: Optional[pulumi.Input[bool]] = None,
            server_flag: Optional[pulumi.Input[bool]] = None,
            street_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            use_csr_common_name: Optional[pulumi.Input[bool]] = None,
            use_csr_sans: Optional[pulumi.Input[bool]] = None) -> 'SecretBackendRole':
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_any_name: Flag to allow any name
        :param pulumi.Input[bool] allow_bare_domains: Flag to allow certificates matching the actual domain
        :param pulumi.Input[bool] allow_glob_domains: Flag to allow names containing glob patterns.
        :param pulumi.Input[bool] allow_ip_sans: Flag to allow IP SANs
        :param pulumi.Input[bool] allow_localhost: Flag to allow certificates for localhost
        :param pulumi.Input[bool] allow_subdomains: Flag to allow certificates matching subdomains
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_domains: List of allowed domains for certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_other_sans: Defines allowed custom SANs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_uri_sans: Defines allowed URI SANs
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[bool] basic_constraints_valid_for_non_ca: Flag to mark basic constraints valid when issuing non-CA certificates
        :param pulumi.Input[bool] client_flag: Flag to specify certificates for client use
        :param pulumi.Input[bool] code_signing_flag: Flag to specify certificates for code signing use
        :param pulumi.Input[Sequence[pulumi.Input[str]]] countries: The country of generated certificates
        :param pulumi.Input[bool] email_protection_flag: Flag to specify certificates for email protection use
        :param pulumi.Input[bool] enforce_hostnames: Flag to allow only valid host names
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ext_key_usages: Specify the allowed extended key usage constraint on issued certificates
        :param pulumi.Input[bool] generate_lease: Flag to generate leases with certificates
        :param pulumi.Input[int] key_bits: The number of bits of generated keys
        :param pulumi.Input[str] key_type: The type of generated keys
        :param pulumi.Input[Sequence[pulumi.Input[str]]] key_usages: Specify the allowed key usage constraint on issued certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] localities: The locality of generated certificates
        :param pulumi.Input[str] max_ttl: The maximum TTL
        :param pulumi.Input[str] name: The name to identify this role within the backend. Must be unique within the backend.
        :param pulumi.Input[bool] no_store: Flag to not store certificates in the storage backend
        :param pulumi.Input[str] not_before_duration: Specifies the duration by which to backdate the NotBefore property.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organization_unit: The organization unit of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] organizations: The organization of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policy_identifiers: Specify the list of allowed policies IODs
        :param pulumi.Input[Sequence[pulumi.Input[str]]] postal_codes: The postal code of generated certificates
        :param pulumi.Input[Sequence[pulumi.Input[str]]] provinces: The province of generated certificates
        :param pulumi.Input[bool] require_cn: Flag to force CN usage
        :param pulumi.Input[bool] server_flag: Flag to specify certificates for server use
        :param pulumi.Input[Sequence[pulumi.Input[str]]] street_addresses: The street address of generated certificates
        :param pulumi.Input[str] ttl: The TTL
        :param pulumi.Input[bool] use_csr_common_name: Flag to use the CN in the CSR
        :param pulumi.Input[bool] use_csr_sans: Flag to use the SANs in the CSR
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allow_any_name"] = allow_any_name
        __props__["allow_bare_domains"] = allow_bare_domains
        __props__["allow_glob_domains"] = allow_glob_domains
        __props__["allow_ip_sans"] = allow_ip_sans
        __props__["allow_localhost"] = allow_localhost
        __props__["allow_subdomains"] = allow_subdomains
        __props__["allowed_domains"] = allowed_domains
        __props__["allowed_other_sans"] = allowed_other_sans
        __props__["allowed_uri_sans"] = allowed_uri_sans
        __props__["backend"] = backend
        __props__["basic_constraints_valid_for_non_ca"] = basic_constraints_valid_for_non_ca
        __props__["client_flag"] = client_flag
        __props__["code_signing_flag"] = code_signing_flag
        __props__["countries"] = countries
        __props__["email_protection_flag"] = email_protection_flag
        __props__["enforce_hostnames"] = enforce_hostnames
        __props__["ext_key_usages"] = ext_key_usages
        __props__["generate_lease"] = generate_lease
        __props__["key_bits"] = key_bits
        __props__["key_type"] = key_type
        __props__["key_usages"] = key_usages
        __props__["localities"] = localities
        __props__["max_ttl"] = max_ttl
        __props__["name"] = name
        __props__["no_store"] = no_store
        __props__["not_before_duration"] = not_before_duration
        __props__["organization_unit"] = organization_unit
        __props__["organizations"] = organizations
        __props__["policy_identifiers"] = policy_identifiers
        __props__["postal_codes"] = postal_codes
        __props__["provinces"] = provinces
        __props__["require_cn"] = require_cn
        __props__["server_flag"] = server_flag
        __props__["street_addresses"] = street_addresses
        __props__["ttl"] = ttl
        __props__["use_csr_common_name"] = use_csr_common_name
        __props__["use_csr_sans"] = use_csr_sans
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowAnyName")
    def allow_any_name(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow any name
        """
        return pulumi.get(self, "allow_any_name")

    @property
    @pulumi.getter(name="allowBareDomains")
    def allow_bare_domains(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow certificates matching the actual domain
        """
        return pulumi.get(self, "allow_bare_domains")

    @property
    @pulumi.getter(name="allowGlobDomains")
    def allow_glob_domains(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow names containing glob patterns.
        """
        return pulumi.get(self, "allow_glob_domains")

    @property
    @pulumi.getter(name="allowIpSans")
    def allow_ip_sans(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow IP SANs
        """
        return pulumi.get(self, "allow_ip_sans")

    @property
    @pulumi.getter(name="allowLocalhost")
    def allow_localhost(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow certificates for localhost
        """
        return pulumi.get(self, "allow_localhost")

    @property
    @pulumi.getter(name="allowSubdomains")
    def allow_subdomains(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow certificates matching subdomains
        """
        return pulumi.get(self, "allow_subdomains")

    @property
    @pulumi.getter(name="allowedDomains")
    def allowed_domains(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of allowed domains for certificates
        """
        return pulumi.get(self, "allowed_domains")

    @property
    @pulumi.getter(name="allowedOtherSans")
    def allowed_other_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Defines allowed custom SANs
        """
        return pulumi.get(self, "allowed_other_sans")

    @property
    @pulumi.getter(name="allowedUriSans")
    def allowed_uri_sans(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Defines allowed URI SANs
        """
        return pulumi.get(self, "allowed_uri_sans")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="basicConstraintsValidForNonCa")
    def basic_constraints_valid_for_non_ca(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to mark basic constraints valid when issuing non-CA certificates
        """
        return pulumi.get(self, "basic_constraints_valid_for_non_ca")

    @property
    @pulumi.getter(name="clientFlag")
    def client_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to specify certificates for client use
        """
        return pulumi.get(self, "client_flag")

    @property
    @pulumi.getter(name="codeSigningFlag")
    def code_signing_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to specify certificates for code signing use
        """
        return pulumi.get(self, "code_signing_flag")

    @property
    @pulumi.getter
    def countries(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The country of generated certificates
        """
        return pulumi.get(self, "countries")

    @property
    @pulumi.getter(name="emailProtectionFlag")
    def email_protection_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to specify certificates for email protection use
        """
        return pulumi.get(self, "email_protection_flag")

    @property
    @pulumi.getter(name="enforceHostnames")
    def enforce_hostnames(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to allow only valid host names
        """
        return pulumi.get(self, "enforce_hostnames")

    @property
    @pulumi.getter(name="extKeyUsages")
    def ext_key_usages(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specify the allowed extended key usage constraint on issued certificates
        """
        return pulumi.get(self, "ext_key_usages")

    @property
    @pulumi.getter(name="generateLease")
    def generate_lease(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to generate leases with certificates
        """
        return pulumi.get(self, "generate_lease")

    @property
    @pulumi.getter(name="keyBits")
    def key_bits(self) -> pulumi.Output[Optional[int]]:
        """
        The number of bits of generated keys
        """
        return pulumi.get(self, "key_bits")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of generated keys
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="keyUsages")
    def key_usages(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specify the allowed key usage constraint on issued certificates
        """
        return pulumi.get(self, "key_usages")

    @property
    @pulumi.getter
    def localities(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The locality of generated certificates
        """
        return pulumi.get(self, "localities")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The maximum TTL
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to identify this role within the backend. Must be unique within the backend.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="noStore")
    def no_store(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to not store certificates in the storage backend
        """
        return pulumi.get(self, "no_store")

    @property
    @pulumi.getter(name="notBeforeDuration")
    def not_before_duration(self) -> pulumi.Output[str]:
        """
        Specifies the duration by which to backdate the NotBefore property.
        """
        return pulumi.get(self, "not_before_duration")

    @property
    @pulumi.getter(name="organizationUnit")
    def organization_unit(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The organization unit of generated certificates
        """
        return pulumi.get(self, "organization_unit")

    @property
    @pulumi.getter
    def organizations(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The organization of generated certificates
        """
        return pulumi.get(self, "organizations")

    @property
    @pulumi.getter(name="policyIdentifiers")
    def policy_identifiers(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Specify the list of allowed policies IODs
        """
        return pulumi.get(self, "policy_identifiers")

    @property
    @pulumi.getter(name="postalCodes")
    def postal_codes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The postal code of generated certificates
        """
        return pulumi.get(self, "postal_codes")

    @property
    @pulumi.getter
    def provinces(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The province of generated certificates
        """
        return pulumi.get(self, "provinces")

    @property
    @pulumi.getter(name="requireCn")
    def require_cn(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to force CN usage
        """
        return pulumi.get(self, "require_cn")

    @property
    @pulumi.getter(name="serverFlag")
    def server_flag(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to specify certificates for server use
        """
        return pulumi.get(self, "server_flag")

    @property
    @pulumi.getter(name="streetAddresses")
    def street_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The street address of generated certificates
        """
        return pulumi.get(self, "street_addresses")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[str]]:
        """
        The TTL
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter(name="useCsrCommonName")
    def use_csr_common_name(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to use the CN in the CSR
        """
        return pulumi.get(self, "use_csr_common_name")

    @property
    @pulumi.getter(name="useCsrSans")
    def use_csr_sans(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to use the SANs in the CSR
        """
        return pulumi.get(self, "use_csr_sans")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

