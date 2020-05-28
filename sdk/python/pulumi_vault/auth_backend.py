# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from . import utilities, tables

class AuthBackend(pulumi.CustomResource):
    accessor: pulumi.Output[str]
    """
    The accessor for this auth method
    """
    default_lease_ttl_seconds: pulumi.Output[float]
    """
    (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
    """
    description: pulumi.Output[str]
    """
    A description of the auth method
    """
    listing_visibility: pulumi.Output[str]
    """
    Specifies whether to show this mount in
    the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
    """
    local: pulumi.Output[bool]
    """
    Specifies if the auth method is local only.
    """
    max_lease_ttl_seconds: pulumi.Output[float]
    """
    (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
    """
    path: pulumi.Output[str]
    """
    The path to mount the auth method — this defaults to the name of the type
    """
    tune: pulumi.Output[dict]
    """
    Extra configuration block. Structure is documented below.

      * `allowedResponseHeaders` (`list`) - List of headers to whitelist and allowing
        a plugin to include them in the response.
      * `auditNonHmacRequestKeys` (`list`) - Specifies the list of keys that will
        not be HMAC'd by audit devices in the request data object.
      * `auditNonHmacResponseKeys` (`list`) - Specifies the list of keys that will
        not be HMAC'd by audit devices in the response data object.
      * `defaultLeaseTtl` (`str`) - Specifies the default time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
      * `listing_visibility` (`str`) - Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
      * `maxLeaseTtl` (`str`) - Specifies the maximum time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
      * `passthroughRequestHeaders` (`list`) - List of headers to whitelist and
        pass from the request to the backend.
      * `token_type` (`str`) - Specifies the type of tokens that should be returned by
        the mount. Valid values are "default-service", "default-batch", "service", "batch".
    """
    type: pulumi.Output[str]
    """
    The name of the auth method type
    """
    def __init__(__self__, resource_name, opts=None, default_lease_ttl_seconds=None, description=None, listing_visibility=None, local=None, max_lease_ttl_seconds=None, path=None, tune=None, type=None, __props__=None, __name__=None, __opts__=None):
        """
        ## Example Usage



        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.AuthBackend("example",
            tune={
                "listing_visibility": "unauth",
                "maxLeaseTtl": "90000s",
            },
            type="github")
        ```


        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[float] default_lease_ttl_seconds: (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
        :param pulumi.Input[str] description: A description of the auth method
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[bool] local: Specifies if the auth method is local only.
        :param pulumi.Input[float] max_lease_ttl_seconds: (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
        :param pulumi.Input[str] path: The path to mount the auth method — this defaults to the name of the type
        :param pulumi.Input[dict] tune: Extra configuration block. Structure is documented below.
        :param pulumi.Input[str] type: The name of the auth method type

        The **tune** object supports the following:

          * `allowedResponseHeaders` (`pulumi.Input[list]`) - List of headers to whitelist and allowing
            a plugin to include them in the response.
          * `auditNonHmacRequestKeys` (`pulumi.Input[list]`) - Specifies the list of keys that will
            not be HMAC'd by audit devices in the request data object.
          * `auditNonHmacResponseKeys` (`pulumi.Input[list]`) - Specifies the list of keys that will
            not be HMAC'd by audit devices in the response data object.
          * `defaultLeaseTtl` (`pulumi.Input[str]`) - Specifies the default time-to-live.
            If set, this overrides the global default.
            Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
          * `listing_visibility` (`pulumi.Input[str]`) - Specifies whether to show this mount in
            the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
          * `maxLeaseTtl` (`pulumi.Input[str]`) - Specifies the maximum time-to-live.
            If set, this overrides the global default.
            Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
          * `passthroughRequestHeaders` (`pulumi.Input[list]`) - List of headers to whitelist and
            pass from the request to the backend.
          * `token_type` (`pulumi.Input[str]`) - Specifies the type of tokens that should be returned by
            the mount. Valid values are "default-service", "default-batch", "service", "batch".
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

            __props__['default_lease_ttl_seconds'] = default_lease_ttl_seconds
            __props__['description'] = description
            __props__['listing_visibility'] = listing_visibility
            __props__['local'] = local
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
    def get(resource_name, id, opts=None, accessor=None, default_lease_ttl_seconds=None, description=None, listing_visibility=None, local=None, max_lease_ttl_seconds=None, path=None, tune=None, type=None):
        """
        Get an existing AuthBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] accessor: The accessor for this auth method
        :param pulumi.Input[float] default_lease_ttl_seconds: (Optional; Deprecated, use `tune.default_lease_ttl` if you are using Vault provider version >= 1.8) The default lease duration in seconds.
        :param pulumi.Input[str] description: A description of the auth method
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[bool] local: Specifies if the auth method is local only.
        :param pulumi.Input[float] max_lease_ttl_seconds: (Optional; Deprecated, use `tune.max_lease_ttl` if you are using Vault provider version >= 1.8) The maximum lease duration in seconds.
        :param pulumi.Input[str] path: The path to mount the auth method — this defaults to the name of the type
        :param pulumi.Input[dict] tune: Extra configuration block. Structure is documented below.
        :param pulumi.Input[str] type: The name of the auth method type

        The **tune** object supports the following:

          * `allowedResponseHeaders` (`pulumi.Input[list]`) - List of headers to whitelist and allowing
            a plugin to include them in the response.
          * `auditNonHmacRequestKeys` (`pulumi.Input[list]`) - Specifies the list of keys that will
            not be HMAC'd by audit devices in the request data object.
          * `auditNonHmacResponseKeys` (`pulumi.Input[list]`) - Specifies the list of keys that will
            not be HMAC'd by audit devices in the response data object.
          * `defaultLeaseTtl` (`pulumi.Input[str]`) - Specifies the default time-to-live.
            If set, this overrides the global default.
            Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
          * `listing_visibility` (`pulumi.Input[str]`) - Specifies whether to show this mount in
            the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
          * `maxLeaseTtl` (`pulumi.Input[str]`) - Specifies the maximum time-to-live.
            If set, this overrides the global default.
            Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
          * `passthroughRequestHeaders` (`pulumi.Input[list]`) - List of headers to whitelist and
            pass from the request to the backend.
          * `token_type` (`pulumi.Input[str]`) - Specifies the type of tokens that should be returned by
            the mount. Valid values are "default-service", "default-batch", "service", "batch".
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

