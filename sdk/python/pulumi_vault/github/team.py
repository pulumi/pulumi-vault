# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class Team(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    Path where the github auth backend is mounted. Defaults to `github`
    if not specified.
    """
    policies: pulumi.Output[list]
    """
    An array of strings specifying the policies to be set on tokens
    issued using this role.
    """
    team: pulumi.Output[str]
    """
    GitHub team name in "slugified" format.
    """
    token_bound_cidrs: pulumi.Output[list]
    """
    Specifies the blocks of IP addresses which are allowed to use the generated token
    """
    token_explicit_max_ttl: pulumi.Output[float]
    """
    Generated Token's Explicit Maximum TTL in seconds
    """
    token_max_ttl: pulumi.Output[float]
    """
    The maximum lifetime of the generated token
    """
    token_no_default_policy: pulumi.Output[bool]
    """
    If true, the 'default' policy will not automatically be added to generated tokens
    """
    token_num_uses: pulumi.Output[float]
    """
    The maximum number of times a token may be used, a value of zero means unlimited
    """
    token_period: pulumi.Output[float]
    """
    Generated Token's Period
    """
    token_policies: pulumi.Output[list]
    """
    Generated Token's Policies
    """
    token_ttl: pulumi.Output[float]
    """
    The initial ttl of the token to generate in seconds
    """
    token_type: pulumi.Output[str]
    """
    The type of token to generate, service or batch
    """
    def __init__(__self__, resource_name, opts=None, backend=None, policies=None, team=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None, __props__=None, __name__=None, __opts__=None):
        """
        Create a Team resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[list] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        :param pulumi.Input[list] token_bound_cidrs: Specifies the blocks of IP addresses which are allowed to use the generated token
        :param pulumi.Input[float] token_explicit_max_ttl: Generated Token's Explicit Maximum TTL in seconds
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime of the generated token
        :param pulumi.Input[bool] token_no_default_policy: If true, the 'default' policy will not automatically be added to generated tokens
        :param pulumi.Input[float] token_num_uses: The maximum number of times a token may be used, a value of zero means unlimited
        :param pulumi.Input[float] token_period: Generated Token's Period
        :param pulumi.Input[list] token_policies: Generated Token's Policies
        :param pulumi.Input[float] token_ttl: The initial ttl of the token to generate in seconds
        :param pulumi.Input[str] token_type: The type of token to generate, service or batch
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['backend'] = backend
            __props__['policies'] = policies
            if team is None:
                raise TypeError("Missing required property 'team'")
            __props__['team'] = team
            if token_bound_cidrs is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_bound_cidrs is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_bound_cidrs'] = token_bound_cidrs
            if token_explicit_max_ttl is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_explicit_max_ttl is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_explicit_max_ttl'] = token_explicit_max_ttl
            if token_max_ttl is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_max_ttl is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_max_ttl'] = token_max_ttl
            if token_no_default_policy is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_no_default_policy is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_no_default_policy'] = token_no_default_policy
            if token_num_uses is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_num_uses is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_num_uses'] = token_num_uses
            if token_period is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_period is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_period'] = token_period
            if token_policies is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_policies is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_policies'] = token_policies
            if token_ttl is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_ttl is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_ttl'] = token_ttl
            if token_type is not None:
                warnings.warn("This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.", DeprecationWarning)
                pulumi.log.warn("token_type is deprecated: This parameter should be moved to the Github Auth backend config block. It does nothing in a user/team block.")
            __props__['token_type'] = token_type
        super(Team, __self__).__init__(
            'vault:github/team:Team',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, policies=None, team=None, token_bound_cidrs=None, token_explicit_max_ttl=None, token_max_ttl=None, token_no_default_policy=None, token_num_uses=None, token_period=None, token_policies=None, token_ttl=None, token_type=None):
        """
        Get an existing Team resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[list] policies: An array of strings specifying the policies to be set on tokens
               issued using this role.
        :param pulumi.Input[str] team: GitHub team name in "slugified" format.
        :param pulumi.Input[list] token_bound_cidrs: Specifies the blocks of IP addresses which are allowed to use the generated token
        :param pulumi.Input[float] token_explicit_max_ttl: Generated Token's Explicit Maximum TTL in seconds
        :param pulumi.Input[float] token_max_ttl: The maximum lifetime of the generated token
        :param pulumi.Input[bool] token_no_default_policy: If true, the 'default' policy will not automatically be added to generated tokens
        :param pulumi.Input[float] token_num_uses: The maximum number of times a token may be used, a value of zero means unlimited
        :param pulumi.Input[float] token_period: Generated Token's Period
        :param pulumi.Input[list] token_policies: Generated Token's Policies
        :param pulumi.Input[float] token_ttl: The initial ttl of the token to generate in seconds
        :param pulumi.Input[str] token_type: The type of token to generate, service or batch
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["policies"] = policies
        __props__["team"] = team
        __props__["token_bound_cidrs"] = token_bound_cidrs
        __props__["token_explicit_max_ttl"] = token_explicit_max_ttl
        __props__["token_max_ttl"] = token_max_ttl
        __props__["token_no_default_policy"] = token_no_default_policy
        __props__["token_num_uses"] = token_num_uses
        __props__["token_period"] = token_period
        __props__["token_policies"] = token_policies
        __props__["token_ttl"] = token_ttl
        __props__["token_type"] = token_type
        return Team(resource_name, opts=opts, __props__=__props__)
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

