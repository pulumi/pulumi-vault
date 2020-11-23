# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['OidcRole']


class OidcRole(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 key: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        ## Import

        The key can be imported with the role name, for example

        ```sh
         $ pulumi import vault:identity/oidcRole:OidcRole role role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] key: A configured named key, the key must already exist
               before tokens can be issued.
        :param pulumi.Input[str] name: Name of the OIDC Role to create.
        :param pulumi.Input[str] template: The template string to use for generating tokens. This may be in
               string-ified JSON or base64 format. See the
               [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
               for the template format.
        :param pulumi.Input[int] ttl: TTL of the tokens generated against the role in number of seconds.
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

            if key is None:
                raise TypeError("Missing required property 'key'")
            __props__['key'] = key
            __props__['name'] = name
            __props__['template'] = template
            __props__['ttl'] = ttl
            __props__['client_id'] = None
        super(OidcRole, __self__).__init__(
            'vault:identity/oidcRole:OidcRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            key: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None,
            ttl: Optional[pulumi.Input[int]] = None) -> 'OidcRole':
        """
        Get an existing OidcRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: The value that will be included in the `aud` field of all the OIDC identity
               tokens issued by this role
        :param pulumi.Input[str] key: A configured named key, the key must already exist
               before tokens can be issued.
        :param pulumi.Input[str] name: Name of the OIDC Role to create.
        :param pulumi.Input[str] template: The template string to use for generating tokens. This may be in
               string-ified JSON or base64 format. See the
               [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
               for the template format.
        :param pulumi.Input[int] ttl: TTL of the tokens generated against the role in number of seconds.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["client_id"] = client_id
        __props__["key"] = key
        __props__["name"] = name
        __props__["template"] = template
        __props__["ttl"] = ttl
        return OidcRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[str]:
        """
        The value that will be included in the `aud` field of all the OIDC identity
        tokens issued by this role
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter
    def key(self) -> pulumi.Output[str]:
        """
        A configured named key, the key must already exist
        before tokens can be issued.
        """
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the OIDC Role to create.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional[str]]:
        """
        The template string to use for generating tokens. This may be in
        string-ified JSON or base64 format. See the
        [documentation](https://www.vaultproject.io/docs/secrets/identity/index.html#token-contents-and-templates)
        for the template format.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[Optional[int]]:
        """
        TTL of the tokens generated against the role in number of seconds.
        """
        return pulumi.get(self, "ttl")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

