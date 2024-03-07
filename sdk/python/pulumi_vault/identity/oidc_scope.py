# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['OidcScopeArgs', 'OidcScope']

@pulumi.input_type
class OidcScopeArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OidcScope resource.
        :param pulumi.Input[str] description: A description of the scope.
        :param pulumi.Input[str] name: The name of the scope. The `openid` scope name is reserved.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] template: The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if template is not None:
            pulumi.set(__self__, "template", template)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the scope.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the scope. The `openid` scope name is reserved.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)


@pulumi.input_type
class _OidcScopeState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OidcScope resources.
        :param pulumi.Input[str] description: A description of the scope.
        :param pulumi.Input[str] name: The name of the scope. The `openid` scope name is reserved.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] template: The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if template is not None:
            pulumi.set(__self__, "template", template)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description of the scope.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the scope. The `openid` scope name is reserved.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)


class OidcScope(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages OIDC Scopes in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-a-scope)
        for more information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        groups = vault.identity.OidcScope("groups",
            description="Vault OIDC Groups Scope",
            template="{\\"groups\\":{{identity.entity.groups.names}}}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OIDC Scopes can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:identity/oidcScope:OidcScope groups groups
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the scope.
        :param pulumi.Input[str] name: The name of the scope. The `openid` scope name is reserved.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] template: The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[OidcScopeArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages OIDC Scopes in a Vault server. See the [Vault documentation](https://www.vaultproject.io/api-docs/secret/identity/oidc-provider#create-or-update-a-scope)
        for more information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        groups = vault.identity.OidcScope("groups",
            description="Vault OIDC Groups Scope",
            template="{\\"groups\\":{{identity.entity.groups.names}}}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        OIDC Scopes can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:identity/oidcScope:OidcScope groups groups
        ```

        :param str resource_name: The name of the resource.
        :param OidcScopeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OidcScopeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OidcScopeArgs.__new__(OidcScopeArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["template"] = template
        super(OidcScope, __self__).__init__(
            'vault:identity/oidcScope:OidcScope',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None) -> 'OidcScope':
        """
        Get an existing OidcScope resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description of the scope.
        :param pulumi.Input[str] name: The name of the scope. The `openid` scope name is reserved.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] template: The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OidcScopeState.__new__(_OidcScopeState)

        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["template"] = template
        return OidcScope(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description of the scope.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the scope. The `openid` scope name is reserved.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def template(self) -> pulumi.Output[Optional[str]]:
        """
        The template string for the scope. This may be provided as escaped JSON or base64 encoded JSON.
        """
        return pulumi.get(self, "template")

