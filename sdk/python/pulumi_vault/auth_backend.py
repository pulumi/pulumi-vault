# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs
from ._inputs import *

__all__ = ['AuthBackend']


class AuthBackend(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 listing_visibility: Optional[pulumi.Input[str]] = None,
                 local: Optional[pulumi.Input[bool]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 tune: Optional[pulumi.Input[pulumi.InputType['AuthBackendTuneArgs']]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.AuthBackend("example",
            tune=vault.AuthBackendTuneArgs(
                listing_visibility="unauth",
                max_lease_ttl="90000s",
            ),
            type="github")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] default_lease_ttl_seconds: (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
        :param pulumi.Input[str] description: A description of the auth method
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[bool] local: Specifies if the auth method is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
        :param pulumi.Input[str] path: The path to mount the auth method — this defaults to the name of the type
        :param pulumi.Input[pulumi.InputType['AuthBackendTuneArgs']] tune: Extra configuration block. Structure is documented below.
        :param pulumi.Input[str] type: The name of the auth method type
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

            if default_lease_ttl_seconds is not None:
                warnings.warn("Use the tune configuration block to avoid forcing creation of new resource on an update", DeprecationWarning)
                pulumi.log.warn("default_lease_ttl_seconds is deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update")
            __props__['default_lease_ttl_seconds'] = default_lease_ttl_seconds
            __props__['description'] = description
            if listing_visibility is not None:
                warnings.warn("Use the tune configuration block to avoid forcing creation of new resource on an update", DeprecationWarning)
                pulumi.log.warn("listing_visibility is deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update")
            __props__['listing_visibility'] = listing_visibility
            __props__['local'] = local
            if max_lease_ttl_seconds is not None:
                warnings.warn("Use the tune configuration block to avoid forcing creation of new resource on an update", DeprecationWarning)
                pulumi.log.warn("max_lease_ttl_seconds is deprecated: Use the tune configuration block to avoid forcing creation of new resource on an update")
            __props__['max_lease_ttl_seconds'] = max_lease_ttl_seconds
            __props__['path'] = path
            __props__['tune'] = tune
            if type is None:
                raise TypeError("Missing required property 'type'")
            __props__['type'] = type
            __props__['accessor'] = None
        super(AuthBackend, __self__).__init__(
            'vault:index/authBackend:AuthBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accessor: Optional[pulumi.Input[str]] = None,
            default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            listing_visibility: Optional[pulumi.Input[str]] = None,
            local: Optional[pulumi.Input[bool]] = None,
            max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            path: Optional[pulumi.Input[str]] = None,
            tune: Optional[pulumi.Input[pulumi.InputType['AuthBackendTuneArgs']]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'AuthBackend':
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor for this auth method
        :param pulumi.Input[int] default_lease_ttl_seconds: (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
        :param pulumi.Input[str] description: A description of the auth method
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[bool] local: Specifies if the auth method is local only.
        :param pulumi.Input[int] max_lease_ttl_seconds: (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
        :param pulumi.Input[str] path: The path to mount the auth method — this defaults to the name of the type
        :param pulumi.Input[pulumi.InputType['AuthBackendTuneArgs']] tune: Extra configuration block. Structure is documented below.
        :param pulumi.Input[str] type: The name of the auth method type
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["accessor"] = accessor
        __props__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
        __props__["description"] = description
        __props__["listing_visibility"] = listing_visibility
        __props__["local"] = local
        __props__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
        __props__["path"] = path
        __props__["tune"] = tune
        __props__["type"] = type
        return AuthBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def accessor(self) -> pulumi.Output[str]:
        """
        The accessor for this auth method
        """
        return pulumi.get(self, "accessor")

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the auth method
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> pulumi.Output[str]:
        """
        Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        """
        return pulumi.get(self, "listing_visibility")

    @property
    @pulumi.getter
    def local(self) -> pulumi.Output[Optional[bool]]:
        """
        Specifies if the auth method is local only.
        """
        return pulumi.get(self, "local")

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[str]:
        """
        The path to mount the auth method — this defaults to the name of the type
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def tune(self) -> pulumi.Output['outputs.AuthBackendTune']:
        """
        Extra configuration block. Structure is documented below.
        """
        return pulumi.get(self, "tune")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The name of the auth method type
        """
        return pulumi.get(self, "type")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

