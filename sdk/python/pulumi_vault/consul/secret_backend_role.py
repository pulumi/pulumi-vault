# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Dict, List, Mapping, Optional, Tuple, Union
from .. import _utilities, _tables

__all__ = ['SecretBackendRole']


class SecretBackendRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_ttl: Optional[pulumi.Input[float]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
                 token_type: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[float]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Manages a Consul secrets role for a Consul secrets engine in Vault. Consul secret backends can then issue Consul tokens.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.consul.SecretBackend("test",
            path="consul",
            description="Manages the Consul backend",
            address="127.0.0.1:8500",
            token="4240861b-ce3d-8530-115a-521ff070dd29")
        example = vault.consul.SecretBackendRole("example",
            backend=test.path,
            policies=["example-policy"])
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        :param pulumi.Input[bool] local: Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        :param pulumi.Input[float] max_ttl: Maximum TTL for leases associated with this role, in seconds.
        :param pulumi.Input[str] name: The name of the Consul secrets engine role to create.
        :param pulumi.Input[str] path: The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. **Deprecated**
        :param pulumi.Input[List[pulumi.Input[str]]] policies: The list of Consul ACL policies to associate with these roles.
        :param pulumi.Input[str] token_type: Specifies the type of token to create when using this role. Valid values are "client" or "management".
        :param pulumi.Input[float] ttl: Specifies the TTL for this role.
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

            __props__['backend'] = backend
            __props__['local'] = local
            __props__['max_ttl'] = max_ttl
            __props__['name'] = name
            if path is not None:
                warnings.warn("use `backend` instead", DeprecationWarning)
                pulumi.log.warn("path is deprecated: use `backend` instead")
            __props__['path'] = path
            if policies is None:
                raise TypeError("Missing required property 'policies'")
            __props__['policies'] = policies
            __props__['token_type'] = token_type
            __props__['ttl'] = ttl
        super(SecretBackendRole, __self__).__init__(
            'vault:consul/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            local: Optional[pulumi.Input[bool]] = None,
            max_ttl: Optional[pulumi.Input[float]] = None,
            name: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[List[pulumi.Input[str]]]] = None,
            token_type: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[float]] = None) -> 'SecretBackendRole':
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        :param pulumi.Input[bool] local: Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        :param pulumi.Input[float] max_ttl: Maximum TTL for leases associated with this role, in seconds.
        :param pulumi.Input[str] name: The name of the Consul secrets engine role to create.
        :param pulumi.Input[str] path: The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. **Deprecated**
        :param pulumi.Input[List[pulumi.Input[str]]] policies: The list of Consul ACL policies to associate with these roles.
        :param pulumi.Input[str] token_type: Specifies the type of token to create when using this role. Valid values are "client" or "management".
        :param pulumi.Input[float] ttl: Specifies the TTL for this role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["local"] = local
        __props__["max_ttl"] = max_ttl
        __props__["name"] = name
        __props__["path"] = path
        __props__["policies"] = policies
        __props__["token_type"] = token_type
        __props__["ttl"] = ttl
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. One of `path` or `backend` is required.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def local(self) -> pulumi.Output[Optional[bool]]:
        """
        Indicates that the token should not be replicated globally and instead be local to the current datacenter.
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[Optional[float]]:
        """
        Maximum TTL for leases associated with this role, in seconds.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the Consul secrets engine role to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The unique name of an existing Consul secrets backend mount. Must not begin or end with a `/`. **Deprecated**
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[List[str]]:
        """
        The list of Consul ACL policies to associate with these roles.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies the type of token to create when using this role. Valid values are "client" or "management".
        """
        return pulumi.get(self, "token_type")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[float]]:
        """
        Specifies the TTL for this role.
        """
        return pulumi.get(self, "ttl")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

