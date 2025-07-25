# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from .. import _utilities

__all__ = ['MfaOktaArgs', 'MfaOkta']

@pulumi.input_type
class MfaOktaArgs:
    def __init__(__self__, *,
                 api_token: pulumi.Input[_builtins.str],
                 org_name: pulumi.Input[_builtins.str],
                 base_url: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 primary_email: Optional[pulumi.Input[_builtins.bool]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a MfaOkta resource.
        :param pulumi.Input[_builtins.str] api_token: Okta API token.
        :param pulumi.Input[_builtins.str] org_name: Name of the organization to be used in the Okta API.
        :param pulumi.Input[_builtins.str] base_url: The base domain to use for API requests.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.bool] primary_email: Only match the primary email for the account.
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        """
        pulumi.set(__self__, "api_token", api_token)
        pulumi.set(__self__, "org_name", org_name)
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if primary_email is not None:
            pulumi.set(__self__, "primary_email", primary_email)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)

    @_builtins.property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Input[_builtins.str]:
        """
        Okta API token.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "api_token", value)

    @_builtins.property
    @pulumi.getter(name="orgName")
    def org_name(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @org_name.setter
    def org_name(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "org_name", value)

    @_builtins.property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The base domain to use for API requests.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "base_url", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Only match the primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @primary_email.setter
    def primary_email(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "primary_email", value)

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_format", value)


@pulumi.input_type
class _MfaOktaState:
    def __init__(__self__, *,
                 api_token: Optional[pulumi.Input[_builtins.str]] = None,
                 base_url: Optional[pulumi.Input[_builtins.str]] = None,
                 method_id: Optional[pulumi.Input[_builtins.str]] = None,
                 mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
                 org_name: Optional[pulumi.Input[_builtins.str]] = None,
                 primary_email: Optional[pulumi.Input[_builtins.bool]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 uuid: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering MfaOkta resources.
        :param pulumi.Input[_builtins.str] api_token: Okta API token.
        :param pulumi.Input[_builtins.str] base_url: The base domain to use for API requests.
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] org_name: Name of the organization to be used in the Okta API.
        :param pulumi.Input[_builtins.bool] primary_email: Only match the primary email for the account.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        if api_token is not None:
            pulumi.set(__self__, "api_token", api_token)
        if base_url is not None:
            pulumi.set(__self__, "base_url", base_url)
        if method_id is not None:
            pulumi.set(__self__, "method_id", method_id)
        if mount_accessor is not None:
            pulumi.set(__self__, "mount_accessor", mount_accessor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if org_name is not None:
            pulumi.set(__self__, "org_name", org_name)
        if primary_email is not None:
            pulumi.set(__self__, "primary_email", primary_email)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if username_format is not None:
            pulumi.set(__self__, "username_format", username_format)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @_builtins.property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Okta API token.
        """
        return pulumi.get(self, "api_token")

    @api_token.setter
    def api_token(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "api_token", value)

    @_builtins.property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The base domain to use for API requests.
        """
        return pulumi.get(self, "base_url")

    @base_url.setter
    def base_url(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "base_url", value)

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @method_id.setter
    def method_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "method_id", value)

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mount_accessor", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_id", value)

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_path", value)

    @_builtins.property
    @pulumi.getter(name="orgName")
    def org_name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @org_name.setter
    def org_name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "org_name", value)

    @_builtins.property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> Optional[pulumi.Input[_builtins.bool]]:
        """
        Only match the primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @primary_email.setter
    def primary_email(self, value: Optional[pulumi.Input[_builtins.bool]]):
        pulumi.set(self, "primary_email", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @username_format.setter
    def username_format(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "username_format", value)

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "uuid", value)


@pulumi.type_token("vault:identity/mfaOkta:MfaOkta")
class MfaOkta(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[_builtins.str]] = None,
                 base_url: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 org_name: Optional[pulumi.Input[_builtins.str]] = None,
                 primary_email: Optional[pulumi.Input[_builtins.bool]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Resource for configuring the okta MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaOkta("example",
            org_name="org1",
            api_token="token1",
            base_url="qux.baz.com")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaOkta:MfaOkta example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] api_token: Okta API token.
        :param pulumi.Input[_builtins.str] base_url: The base domain to use for API requests.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] org_name: Name of the organization to be used in the Okta API.
        :param pulumi.Input[_builtins.bool] primary_email: Only match the primary email for the account.
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaOktaArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for configuring the okta MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaOkta("example",
            org_name="org1",
            api_token="token1",
            base_url="qux.baz.com")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaOkta:MfaOkta example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param MfaOktaArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaOktaArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_token: Optional[pulumi.Input[_builtins.str]] = None,
                 base_url: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 org_name: Optional[pulumi.Input[_builtins.str]] = None,
                 primary_email: Optional[pulumi.Input[_builtins.bool]] = None,
                 username_format: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaOktaArgs.__new__(MfaOktaArgs)

            if api_token is None and not opts.urn:
                raise TypeError("Missing required property 'api_token'")
            __props__.__dict__["api_token"] = None if api_token is None else pulumi.Output.secret(api_token)
            __props__.__dict__["base_url"] = base_url
            __props__.__dict__["namespace"] = namespace
            if org_name is None and not opts.urn:
                raise TypeError("Missing required property 'org_name'")
            __props__.__dict__["org_name"] = org_name
            __props__.__dict__["primary_email"] = primary_email
            __props__.__dict__["username_format"] = username_format
            __props__.__dict__["method_id"] = None
            __props__.__dict__["mount_accessor"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_path"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["uuid"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["apiToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(MfaOkta, __self__).__init__(
            'vault:identity/mfaOkta:MfaOkta',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            api_token: Optional[pulumi.Input[_builtins.str]] = None,
            base_url: Optional[pulumi.Input[_builtins.str]] = None,
            method_id: Optional[pulumi.Input[_builtins.str]] = None,
            mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
            org_name: Optional[pulumi.Input[_builtins.str]] = None,
            primary_email: Optional[pulumi.Input[_builtins.bool]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None,
            username_format: Optional[pulumi.Input[_builtins.str]] = None,
            uuid: Optional[pulumi.Input[_builtins.str]] = None) -> 'MfaOkta':
        """
        Get an existing MfaOkta resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] api_token: Okta API token.
        :param pulumi.Input[_builtins.str] base_url: The base domain to use for API requests.
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] org_name: Name of the organization to be used in the Okta API.
        :param pulumi.Input[_builtins.bool] primary_email: Only match the primary email for the account.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.str] username_format: A template string for mapping Identity names to MFA methods.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaOktaState.__new__(_MfaOktaState)

        __props__.__dict__["api_token"] = api_token
        __props__.__dict__["base_url"] = base_url
        __props__.__dict__["method_id"] = method_id
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["org_name"] = org_name
        __props__.__dict__["primary_email"] = primary_email
        __props__.__dict__["type"] = type
        __props__.__dict__["username_format"] = username_format
        __props__.__dict__["uuid"] = uuid
        return MfaOkta(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="apiToken")
    def api_token(self) -> pulumi.Output[_builtins.str]:
        """
        Okta API token.
        """
        return pulumi.get(self, "api_token")

    @_builtins.property
    @pulumi.getter(name="baseUrl")
    def base_url(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The base domain to use for API requests.
        """
        return pulumi.get(self, "base_url")

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[_builtins.str]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @_builtins.property
    @pulumi.getter(name="orgName")
    def org_name(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the organization to be used in the Okta API.
        """
        return pulumi.get(self, "org_name")

    @_builtins.property
    @pulumi.getter(name="primaryEmail")
    def primary_email(self) -> pulumi.Output[Optional[_builtins.bool]]:
        """
        Only match the primary email for the account.
        """
        return pulumi.get(self, "primary_email")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter(name="usernameFormat")
    def username_format(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        A template string for mapping Identity names to MFA methods.
        """
        return pulumi.get(self, "username_format")

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[_builtins.str]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

