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
                 algorithm_signer: Optional[pulumi.Input[str]] = None,
                 allow_bare_domains: Optional[pulumi.Input[bool]] = None,
                 allow_host_certificates: Optional[pulumi.Input[bool]] = None,
                 allow_subdomains: Optional[pulumi.Input[bool]] = None,
                 allow_user_certificates: Optional[pulumi.Input[bool]] = None,
                 allow_user_key_ids: Optional[pulumi.Input[bool]] = None,
                 allowed_critical_options: Optional[pulumi.Input[str]] = None,
                 allowed_domains: Optional[pulumi.Input[str]] = None,
                 allowed_extensions: Optional[pulumi.Input[str]] = None,
                 allowed_user_key_lengths: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 allowed_users: Optional[pulumi.Input[str]] = None,
                 allowed_users_template: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 cidr_list: Optional[pulumi.Input[str]] = None,
                 default_critical_options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 default_extensions: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 default_user: Optional[pulumi.Input[str]] = None,
                 key_id_format: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to manage roles in an SSH secret backend
        [SSH secret backend within Vault](https://www.vaultproject.io/docs/secrets/ssh/index.html).

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.Mount("example", type="ssh")
        foo = vault.ssh.SecretBackendRole("foo",
            allow_user_certificates=True,
            backend=example.path,
            key_type="ca")
        bar = vault.ssh.SecretBackendRole("bar",
            allowed_users="default,baz",
            backend=example.path,
            cidr_list="0.0.0.0/0",
            default_user="default",
            key_type="otp")
        ```

        ## Import

        SSH secret backend roles can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:ssh/secretBackendRole:SecretBackendRole foo ssh/roles/my-role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm_signer: When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
        :param pulumi.Input[bool] allow_bare_domains: Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_host_certificates: Specifies if certificates are allowed to be signed for use as a 'host'.
        :param pulumi.Input[bool] allow_subdomains: Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_user_certificates: Specifies if certificates are allowed to be signed for use as a 'user'.
        :param pulumi.Input[bool] allow_user_key_ids: Specifies if users can override the key ID for a signed certificate with the `key_id` field.
        :param pulumi.Input[str] allowed_critical_options: Specifies a comma-separated list of critical options that certificates can have when signed.
        :param pulumi.Input[str] allowed_domains: The list of domains for which a client can request a host certificate.
        :param pulumi.Input[str] allowed_extensions: Specifies a comma-separated list of extensions that certificates can have when signed.
        :param pulumi.Input[Mapping[str, Any]] allowed_user_key_lengths: Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
        :param pulumi.Input[str] allowed_users: Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
        :param pulumi.Input[bool] allowed_users_template: Specifies if `allowed_users` can be declared using identity template policies. Non-templated users are also permitted.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted.
        :param pulumi.Input[str] cidr_list: The comma-separated string of CIDR blocks for which this role is applicable.
        :param pulumi.Input[Mapping[str, Any]] default_critical_options: Specifies a map of critical options that certificates have when signed.
        :param pulumi.Input[Mapping[str, Any]] default_extensions: Specifies a map of extensions that certificates have when signed.
        :param pulumi.Input[str] default_user: Specifies the default username for which a credential will be generated.
        :param pulumi.Input[str] key_id_format: Specifies a custom format for the key id of a signed certificate.
        :param pulumi.Input[str] key_type: Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
        :param pulumi.Input[str] max_ttl: Specifies the maximum Time To Live value.
        :param pulumi.Input[str] name: Specifies the name of the role to create.
        :param pulumi.Input[str] ttl: Specifies the Time To Live value.
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

            __props__['algorithm_signer'] = algorithm_signer
            __props__['allow_bare_domains'] = allow_bare_domains
            __props__['allow_host_certificates'] = allow_host_certificates
            __props__['allow_subdomains'] = allow_subdomains
            __props__['allow_user_certificates'] = allow_user_certificates
            __props__['allow_user_key_ids'] = allow_user_key_ids
            __props__['allowed_critical_options'] = allowed_critical_options
            __props__['allowed_domains'] = allowed_domains
            __props__['allowed_extensions'] = allowed_extensions
            __props__['allowed_user_key_lengths'] = allowed_user_key_lengths
            __props__['allowed_users'] = allowed_users
            __props__['allowed_users_template'] = allowed_users_template
            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['cidr_list'] = cidr_list
            __props__['default_critical_options'] = default_critical_options
            __props__['default_extensions'] = default_extensions
            __props__['default_user'] = default_user
            __props__['key_id_format'] = key_id_format
            if key_type is None:
                raise TypeError("Missing required property 'key_type'")
            __props__['key_type'] = key_type
            __props__['max_ttl'] = max_ttl
            __props__['name'] = name
            __props__['ttl'] = ttl
        super(SecretBackendRole, __self__).__init__(
            'vault:ssh/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm_signer: Optional[pulumi.Input[str]] = None,
            allow_bare_domains: Optional[pulumi.Input[bool]] = None,
            allow_host_certificates: Optional[pulumi.Input[bool]] = None,
            allow_subdomains: Optional[pulumi.Input[bool]] = None,
            allow_user_certificates: Optional[pulumi.Input[bool]] = None,
            allow_user_key_ids: Optional[pulumi.Input[bool]] = None,
            allowed_critical_options: Optional[pulumi.Input[str]] = None,
            allowed_domains: Optional[pulumi.Input[str]] = None,
            allowed_extensions: Optional[pulumi.Input[str]] = None,
            allowed_user_key_lengths: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            allowed_users: Optional[pulumi.Input[str]] = None,
            allowed_users_template: Optional[pulumi.Input[bool]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            cidr_list: Optional[pulumi.Input[str]] = None,
            default_critical_options: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            default_extensions: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            default_user: Optional[pulumi.Input[str]] = None,
            key_id_format: Optional[pulumi.Input[str]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            max_ttl: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None) -> 'SecretBackendRole':
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] algorithm_signer: When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
        :param pulumi.Input[bool] allow_bare_domains: Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_host_certificates: Specifies if certificates are allowed to be signed for use as a 'host'.
        :param pulumi.Input[bool] allow_subdomains: Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
        :param pulumi.Input[bool] allow_user_certificates: Specifies if certificates are allowed to be signed for use as a 'user'.
        :param pulumi.Input[bool] allow_user_key_ids: Specifies if users can override the key ID for a signed certificate with the `key_id` field.
        :param pulumi.Input[str] allowed_critical_options: Specifies a comma-separated list of critical options that certificates can have when signed.
        :param pulumi.Input[str] allowed_domains: The list of domains for which a client can request a host certificate.
        :param pulumi.Input[str] allowed_extensions: Specifies a comma-separated list of extensions that certificates can have when signed.
        :param pulumi.Input[Mapping[str, Any]] allowed_user_key_lengths: Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
        :param pulumi.Input[str] allowed_users: Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
        :param pulumi.Input[bool] allowed_users_template: Specifies if `allowed_users` can be declared using identity template policies. Non-templated users are also permitted.
        :param pulumi.Input[str] backend: The path where the SSH secret backend is mounted.
        :param pulumi.Input[str] cidr_list: The comma-separated string of CIDR blocks for which this role is applicable.
        :param pulumi.Input[Mapping[str, Any]] default_critical_options: Specifies a map of critical options that certificates have when signed.
        :param pulumi.Input[Mapping[str, Any]] default_extensions: Specifies a map of extensions that certificates have when signed.
        :param pulumi.Input[str] default_user: Specifies the default username for which a credential will be generated.
        :param pulumi.Input[str] key_id_format: Specifies a custom format for the key id of a signed certificate.
        :param pulumi.Input[str] key_type: Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
        :param pulumi.Input[str] max_ttl: Specifies the maximum Time To Live value.
        :param pulumi.Input[str] name: Specifies the name of the role to create.
        :param pulumi.Input[str] ttl: Specifies the Time To Live value.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["algorithm_signer"] = algorithm_signer
        __props__["allow_bare_domains"] = allow_bare_domains
        __props__["allow_host_certificates"] = allow_host_certificates
        __props__["allow_subdomains"] = allow_subdomains
        __props__["allow_user_certificates"] = allow_user_certificates
        __props__["allow_user_key_ids"] = allow_user_key_ids
        __props__["allowed_critical_options"] = allowed_critical_options
        __props__["allowed_domains"] = allowed_domains
        __props__["allowed_extensions"] = allowed_extensions
        __props__["allowed_user_key_lengths"] = allowed_user_key_lengths
        __props__["allowed_users"] = allowed_users
        __props__["allowed_users_template"] = allowed_users_template
        __props__["backend"] = backend
        __props__["cidr_list"] = cidr_list
        __props__["default_critical_options"] = default_critical_options
        __props__["default_extensions"] = default_extensions
        __props__["default_user"] = default_user
        __props__["key_id_format"] = key_id_format
        __props__["key_type"] = key_type
        __props__["max_ttl"] = max_ttl
        __props__["name"] = name
        __props__["ttl"] = ttl
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="algorithmSigner")
    def algorithm_signer(self) -> pulumi.Output[str]:
        """
        When supplied, this value specifies a signing algorithm for the key. Possible values: ssh-rsa, rsa-sha2-256, rsa-sha2-512.
        """
        return pulumi.get(self, "algorithm_signer")

    @property
    @pulumi.getter(name="allowBareDomains")
    def allow_bare_domains(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if host certificates that are requested are allowed to use the base domains listed in `allowed_domains`.
        """
        return pulumi.get(self, "allow_bare_domains")

    @property
    @pulumi.getter(name="allowHostCertificates")
    def allow_host_certificates(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if certificates are allowed to be signed for use as a 'host'.
        """
        return pulumi.get(self, "allow_host_certificates")

    @property
    @pulumi.getter(name="allowSubdomains")
    def allow_subdomains(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if host certificates that are requested are allowed to be subdomains of those listed in `allowed_domains`.
        """
        return pulumi.get(self, "allow_subdomains")

    @property
    @pulumi.getter(name="allowUserCertificates")
    def allow_user_certificates(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if certificates are allowed to be signed for use as a 'user'.
        """
        return pulumi.get(self, "allow_user_certificates")

    @property
    @pulumi.getter(name="allowUserKeyIds")
    def allow_user_key_ids(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if users can override the key ID for a signed certificate with the `key_id` field.
        """
        return pulumi.get(self, "allow_user_key_ids")

    @property
    @pulumi.getter(name="allowedCriticalOptions")
    def allowed_critical_options(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a comma-separated list of critical options that certificates can have when signed.
        """
        return pulumi.get(self, "allowed_critical_options")

    @property
    @pulumi.getter(name="allowedDomains")
    def allowed_domains(self) -> pulumi.Output[Optional[str]]:
        """
        The list of domains for which a client can request a host certificate.
        """
        return pulumi.get(self, "allowed_domains")

    @property
    @pulumi.getter(name="allowedExtensions")
    def allowed_extensions(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a comma-separated list of extensions that certificates can have when signed.
        """
        return pulumi.get(self, "allowed_extensions")

    @property
    @pulumi.getter(name="allowedUserKeyLengths")
    def allowed_user_key_lengths(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Specifies a map of ssh key types and their expected sizes which are allowed to be signed by the CA type.
        """
        return pulumi.get(self, "allowed_user_key_lengths")

    @property
    @pulumi.getter(name="allowedUsers")
    def allowed_users(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a comma-separated list of usernames that are to be allowed, only if certain usernames are to be allowed.
        """
        return pulumi.get(self, "allowed_users")

    @property
    @pulumi.getter(name="allowedUsersTemplate")
    def allowed_users_template(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if `allowed_users` can be declared using identity template policies. Non-templated users are also permitted.
        """
        return pulumi.get(self, "allowed_users_template")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path where the SSH secret backend is mounted.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="cidrList")
    def cidr_list(self) -> pulumi.Output[Optional[str]]:
        """
        The comma-separated string of CIDR blocks for which this role is applicable.
        """
        return pulumi.get(self, "cidr_list")

    @property
    @pulumi.getter(name="defaultCriticalOptions")
    def default_critical_options(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Specifies a map of critical options that certificates have when signed.
        """
        return pulumi.get(self, "default_critical_options")

    @property
    @pulumi.getter(name="defaultExtensions")
    def default_extensions(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Specifies a map of extensions that certificates have when signed.
        """
        return pulumi.get(self, "default_extensions")

    @property
    @pulumi.getter(name="defaultUser")
    def default_user(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the default username for which a credential will be generated.
        """
        return pulumi.get(self, "default_user")

    @property
    @pulumi.getter(name="keyIdFormat")
    def key_id_format(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a custom format for the key id of a signed certificate.
        """
        return pulumi.get(self, "key_id_format")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Specifies the type of credentials generated by this role. This can be either `otp`, `dynamic` or `ca`.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[str]:
        """
        Specifies the maximum Time To Live value.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Specifies the name of the role to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[str]:
        """
        Specifies the Time To Live value.
        """
        return pulumi.get(self, "ttl")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

