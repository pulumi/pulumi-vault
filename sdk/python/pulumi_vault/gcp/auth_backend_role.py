# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['AuthBackendRole']


class AuthBackendRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 add_group_aliases: Optional[pulumi.Input[bool]] = None,
                 allow_gce_inference: Optional[pulumi.Input[bool]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 bound_instance_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bound_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bound_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bound_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bound_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 bound_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 max_jwt_exp: Optional[pulumi.Input[str]] = None,
                 max_ttl: Optional[pulumi.Input[str]] = None,
                 period: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 token_bound_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_explicit_max_ttl: Optional[pulumi.Input[int]] = None,
                 token_max_ttl: Optional[pulumi.Input[int]] = None,
                 token_no_default_policy: Optional[pulumi.Input[bool]] = None,
                 token_num_uses: Optional[pulumi.Input[int]] = None,
                 token_period: Optional[pulumi.Input[int]] = None,
                 token_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_ttl: Optional[pulumi.Input[int]] = None,
                 token_type: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Provides a resource to create a role in an [GCP auth backend within Vault](https://www.vaultproject.io/docs/auth/gcp.html).

        ## Import

        GCP authentication roles can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:gcp/authBackendRole:AuthBackendRole my_role auth/gcp/role/my_role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_gce_inference: A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        :param pulumi.Input[str] backend: Path to the mounted GCP auth backend
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_instance_groups: The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_labels: A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_projects: GCP Projects that the role exists within
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_regions: The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_service_accounts: GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_zones: The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        :param pulumi.Input[str] max_jwt_exp: The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of tokens
               issued using this role, provided as a number of seconds.
        :param pulumi.Input[str] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings
               specifying the policies to be set on tokens issued using this role.
        :param pulumi.Input[str] role: Name of the GCP role
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[int] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[int] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[int] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[int] token_period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[int] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[str] ttl: The TTL period of tokens issued
               using this role, provided as a number of seconds.
        :param pulumi.Input[str] type: Type of GCP authentication role (either `gce` or `iam`)
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

            __props__['add_group_aliases'] = add_group_aliases
            __props__['allow_gce_inference'] = allow_gce_inference
            __props__['backend'] = backend
            __props__['bound_instance_groups'] = bound_instance_groups
            __props__['bound_labels'] = bound_labels
            __props__['bound_projects'] = bound_projects
            __props__['bound_regions'] = bound_regions
            __props__['bound_service_accounts'] = bound_service_accounts
            __props__['bound_zones'] = bound_zones
            __props__['max_jwt_exp'] = max_jwt_exp
            if max_ttl is not None and not opts.urn:
                warnings.warn("""use `token_max_ttl` instead if you are running Vault >= 1.2""", DeprecationWarning)
                pulumi.log.warn("max_ttl is deprecated: use `token_max_ttl` instead if you are running Vault >= 1.2")
            __props__['max_ttl'] = max_ttl
            if period is not None and not opts.urn:
                warnings.warn("""use `token_period` instead if you are running Vault >= 1.2""", DeprecationWarning)
                pulumi.log.warn("period is deprecated: use `token_period` instead if you are running Vault >= 1.2")
            __props__['period'] = period
            if policies is not None and not opts.urn:
                warnings.warn("""use `token_policies` instead if you are running Vault >= 1.2""", DeprecationWarning)
                pulumi.log.warn("policies is deprecated: use `token_policies` instead if you are running Vault >= 1.2")
            __props__['policies'] = policies
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['token_bound_cidrs'] = token_bound_cidrs
            __props__['token_explicit_max_ttl'] = token_explicit_max_ttl
            __props__['token_max_ttl'] = token_max_ttl
            __props__['token_no_default_policy'] = token_no_default_policy
            __props__['token_num_uses'] = token_num_uses
            __props__['token_period'] = token_period
            __props__['token_policies'] = token_policies
            __props__['token_ttl'] = token_ttl
            __props__['token_type'] = token_type
            if ttl is not None and not opts.urn:
                warnings.warn("""use `token_ttl` instead if you are running Vault >= 1.2""", DeprecationWarning)
                pulumi.log.warn("ttl is deprecated: use `token_ttl` instead if you are running Vault >= 1.2")
            __props__['ttl'] = ttl
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
        super(AuthBackendRole, __self__).__init__(
            'vault:gcp/authBackendRole:AuthBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            add_group_aliases: Optional[pulumi.Input[bool]] = None,
            allow_gce_inference: Optional[pulumi.Input[bool]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            bound_instance_groups: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bound_labels: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bound_projects: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bound_regions: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bound_service_accounts: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            bound_zones: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            max_jwt_exp: Optional[pulumi.Input[str]] = None,
            max_ttl: Optional[pulumi.Input[str]] = None,
            period: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None,
            token_bound_cidrs: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token_explicit_max_ttl: Optional[pulumi.Input[int]] = None,
            token_max_ttl: Optional[pulumi.Input[int]] = None,
            token_no_default_policy: Optional[pulumi.Input[bool]] = None,
            token_num_uses: Optional[pulumi.Input[int]] = None,
            token_period: Optional[pulumi.Input[int]] = None,
            token_policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token_ttl: Optional[pulumi.Input[int]] = None,
            token_type: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'AuthBackendRole':
        """
        Get an existing AuthBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_gce_inference: A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        :param pulumi.Input[str] backend: Path to the mounted GCP auth backend
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_instance_groups: The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_labels: A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_projects: GCP Projects that the role exists within
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_regions: The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_service_accounts: GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] bound_zones: The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        :param pulumi.Input[str] max_jwt_exp: The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        :param pulumi.Input[str] max_ttl: The maximum allowed lifetime of tokens
               issued using this role, provided as a number of seconds.
        :param pulumi.Input[str] period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings
               specifying the policies to be set on tokens issued using this role.
        :param pulumi.Input[str] role: Name of the GCP role
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_bound_cidrs: List of CIDR blocks; if set, specifies blocks of IP
               addresses which can authenticate successfully, and ties the resulting token to these blocks
               as well.
        :param pulumi.Input[int] token_explicit_max_ttl: If set, will encode an
               [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
               onto the token in number of seconds. This is a hard cap even if `token_ttl` and
               `token_max_ttl` would otherwise allow a renewal.
        :param pulumi.Input[int] token_max_ttl: The maximum lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[bool] token_no_default_policy: If set, the default policy will not be set on
               generated tokens; otherwise it will be added to the policies set in token_policies.
        :param pulumi.Input[int] token_num_uses: The
               [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
               if any, in number of seconds to set on the token.
        :param pulumi.Input[int] token_period: If set, indicates that the
               token generated using this role should never expire. The token should be renewed within the
               duration specified by this value. At each renewal, the token's TTL will be set to the
               value of this field. Specified in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_policies: List of policies to encode onto generated tokens. Depending
               on the auth method, this list may be supplemented by user/group/other values.
        :param pulumi.Input[int] token_ttl: The incremental lifetime for generated tokens in number of seconds.
               Its current value will be referenced at renewal time.
        :param pulumi.Input[str] token_type: The type of token that should be generated. Can be `service`,
               `batch`, or `default` to use the mount's tuned default (which unless changed will be
               `service` tokens). For token store roles, there are two additional possibilities:
               `default-service` and `default-batch` which specify the type to return unless the client
               requests a different type at generation time.
        :param pulumi.Input[str] ttl: The TTL period of tokens issued
               using this role, provided as a number of seconds.
        :param pulumi.Input[str] type: Type of GCP authentication role (either `gce` or `iam`)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["add_group_aliases"] = add_group_aliases
        __props__["allow_gce_inference"] = allow_gce_inference
        __props__["backend"] = backend
        __props__["bound_instance_groups"] = bound_instance_groups
        __props__["bound_labels"] = bound_labels
        __props__["bound_projects"] = bound_projects
        __props__["bound_regions"] = bound_regions
        __props__["bound_service_accounts"] = bound_service_accounts
        __props__["bound_zones"] = bound_zones
        __props__["max_jwt_exp"] = max_jwt_exp
        __props__["max_ttl"] = max_ttl
        __props__["period"] = period
        __props__["policies"] = policies
        __props__["role"] = role
        __props__["token_bound_cidrs"] = token_bound_cidrs
        __props__["token_explicit_max_ttl"] = token_explicit_max_ttl
        __props__["token_max_ttl"] = token_max_ttl
        __props__["token_no_default_policy"] = token_no_default_policy
        __props__["token_num_uses"] = token_num_uses
        __props__["token_period"] = token_period
        __props__["token_policies"] = token_policies
        __props__["token_ttl"] = token_ttl
        __props__["token_type"] = token_type
        __props__["ttl"] = ttl
        __props__["type"] = type
        return AuthBackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="addGroupAliases")
    def add_group_aliases(self) -> pulumi.Output[bool]:
        return pulumi.get(self, "add_group_aliases")

    @property
    @pulumi.getter(name="allowGceInference")
    def allow_gce_inference(self) -> pulumi.Output[bool]:
        """
        A flag to determine if this role should allow GCE instances to authenticate by inferring service accounts from the GCE identity metadata token.
        """
        return pulumi.get(self, "allow_gce_inference")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Path to the mounted GCP auth backend
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="boundInstanceGroups")
    def bound_instance_groups(self) -> pulumi.Output[Sequence[str]]:
        """
        The instance groups that an authorized instance must belong to in order to be authenticated. If specified, either `bound_zones` or `bound_regions` must be set too.
        """
        return pulumi.get(self, "bound_instance_groups")

    @property
    @pulumi.getter(name="boundLabels")
    def bound_labels(self) -> pulumi.Output[Sequence[str]]:
        """
        A comma-separated list of GCP labels formatted as `"key:value"` strings that must be set on authorized GCE instances. Because GCP labels are not currently ACL'd, we recommend that this be used in conjunction with other restrictions.
        """
        return pulumi.get(self, "bound_labels")

    @property
    @pulumi.getter(name="boundProjects")
    def bound_projects(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        GCP Projects that the role exists within
        """
        return pulumi.get(self, "bound_projects")

    @property
    @pulumi.getter(name="boundRegions")
    def bound_regions(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of regions that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a regional group and the group must belong to this region. If bound_zones are provided, this attribute is ignored.
        """
        return pulumi.get(self, "bound_regions")

    @property
    @pulumi.getter(name="boundServiceAccounts")
    def bound_service_accounts(self) -> pulumi.Output[Sequence[str]]:
        """
        GCP Service Accounts allowed to issue tokens under this role. (Note: **Required** if role is `iam`)
        """
        return pulumi.get(self, "bound_service_accounts")

    @property
    @pulumi.getter(name="boundZones")
    def bound_zones(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of zones that a GCE instance must belong to in order to be authenticated. If bound_instance_groups is provided, it is assumed to be a zonal group and the group must belong to this zone.
        """
        return pulumi.get(self, "bound_zones")

    @property
    @pulumi.getter(name="maxJwtExp")
    def max_jwt_exp(self) -> pulumi.Output[str]:
        """
        The number of seconds past the time of authentication that the login param JWT must expire within. For example, if a user attempts to login with a token that expires within an hour and this is set to 15 minutes, Vault will return an error prompting the user to create a new signed JWT with a shorter `exp`. The GCE metadata tokens currently do not allow the `exp` claim to be customized.
        """
        return pulumi.get(self, "max_jwt_exp")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[str]:
        """
        The maximum allowed lifetime of tokens
        issued using this role, provided as a number of seconds.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def period(self) -> pulumi.Output[str]:
        """
        If set, indicates that the
        token generated using this role should never expire. The token should be renewed within the
        duration specified by this value. At each renewal, the token's TTL will be set to the
        value of this field. Specified in seconds.
        """
        return pulumi.get(self, "period")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        """
        An array of strings
        specifying the policies to be set on tokens issued using this role.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Name of the GCP role
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="tokenBoundCidrs")
    def token_bound_cidrs(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of CIDR blocks; if set, specifies blocks of IP
        addresses which can authenticate successfully, and ties the resulting token to these blocks
        as well.
        """
        return pulumi.get(self, "token_bound_cidrs")

    @property
    @pulumi.getter(name="tokenExplicitMaxTtl")
    def token_explicit_max_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        If set, will encode an
        [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
        onto the token in number of seconds. This is a hard cap even if `token_ttl` and
        `token_max_ttl` would otherwise allow a renewal.
        """
        return pulumi.get(self, "token_explicit_max_ttl")

    @property
    @pulumi.getter(name="tokenMaxTtl")
    def token_max_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The maximum lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        return pulumi.get(self, "token_max_ttl")

    @property
    @pulumi.getter(name="tokenNoDefaultPolicy")
    def token_no_default_policy(self) -> pulumi.Output[Optional[bool]]:
        """
        If set, the default policy will not be set on
        generated tokens; otherwise it will be added to the policies set in token_policies.
        """
        return pulumi.get(self, "token_no_default_policy")

    @property
    @pulumi.getter(name="tokenNumUses")
    def token_num_uses(self) -> pulumi.Output[Optional[int]]:
        """
        The
        [period](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls),
        if any, in number of seconds to set on the token.
        """
        return pulumi.get(self, "token_num_uses")

    @property
    @pulumi.getter(name="tokenPeriod")
    def token_period(self) -> pulumi.Output[Optional[int]]:
        """
        If set, indicates that the
        token generated using this role should never expire. The token should be renewed within the
        duration specified by this value. At each renewal, the token's TTL will be set to the
        value of this field. Specified in seconds.
        """
        return pulumi.get(self, "token_period")

    @property
    @pulumi.getter(name="tokenPolicies")
    def token_policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of policies to encode onto generated tokens. Depending
        on the auth method, this list may be supplemented by user/group/other values.
        """
        return pulumi.get(self, "token_policies")

    @property
    @pulumi.getter(name="tokenTtl")
    def token_ttl(self) -> pulumi.Output[Optional[int]]:
        """
        The incremental lifetime for generated tokens in number of seconds.
        Its current value will be referenced at renewal time.
        """
        return pulumi.get(self, "token_ttl")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of token that should be generated. Can be `service`,
        `batch`, or `default` to use the mount's tuned default (which unless changed will be
        `service` tokens). For token store roles, there are two additional possibilities:
        `default-service` and `default-batch` which specify the type to return unless the client
        requests a different type at generation time.
        """
        return pulumi.get(self, "token_type")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[str]:
        """
        The TTL period of tokens issued
        using this role, provided as a number of seconds.
        """
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Type of GCP authentication role (either `gce` or `iam`)
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

