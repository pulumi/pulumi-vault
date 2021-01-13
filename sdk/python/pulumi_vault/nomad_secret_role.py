# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables

__all__ = ['NomadSecretRole']


class NomadSecretRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 global_: Optional[pulumi.Input[bool]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        Nomad secret role can be imported using the `backend`, e.g.

        ```sh
         $ pulumi import vault:index/nomadSecretRole:NomadSecretRole bob nomad/role/bob
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `nomad`.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
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

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            __props__['global_'] = global_
            __props__['policies'] = policies
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__['role'] = role
            __props__['type'] = type
        super(NomadSecretRole, __self__).__init__(
            'vault:index/nomadSecretRole:NomadSecretRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            global_: Optional[pulumi.Input[bool]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            role: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'NomadSecretRole':
        """
        Get an existing NomadSecretRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `nomad`.
        :param pulumi.Input[bool] global_: Specifies if the generated token should be global. Defaults to 
               false.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: List of policies attached to the generated token. This setting is only used 
               when `type` is 'client'.
        :param pulumi.Input[str] role: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] type: Specifies the type of token to create when using this role. Valid 
               settings are 'client' and 'management'. Defaults to 'client'.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["global_"] = global_
        __props__["policies"] = policies
        __props__["role"] = role
        __props__["type"] = type
        return NomadSecretRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `nomad`.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="global")
    def global_(self) -> pulumi.Output[bool]:
        """
        Specifies if the generated token should be global. Defaults to 
        false.
        """
        return pulumi.get(self, "global_")

    @property
    @pulumi.getter
    def policies(self) -> pulumi.Output[Sequence[str]]:
        """
        List of policies attached to the generated token. This setting is only used 
        when `type` is 'client'.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        The name to identify this role within the backend.
        Must be unique within the backend.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the type of token to create when using this role. Valid 
        settings are 'client' and 'management'. Defaults to 'client'.
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

