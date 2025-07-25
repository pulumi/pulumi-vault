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

__all__ = ['MfaLoginEnforcementArgs', 'MfaLoginEnforcement']

@pulumi.input_type
class MfaLoginEnforcementArgs:
    def __init__(__self__, *,
                 mfa_method_ids: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]],
                 auth_method_accessors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 auth_method_types: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a MfaLoginEnforcement resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] mfa_method_ids: Set of MFA method UUIDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_accessors: Set of auth method accessor IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_types: Set of auth method types.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_entity_ids: Set of identity entity IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_group_ids: Set of identity group IDs.
        :param pulumi.Input[_builtins.str] name: Login enforcement name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        """
        pulumi.set(__self__, "mfa_method_ids", mfa_method_ids)
        if auth_method_accessors is not None:
            pulumi.set(__self__, "auth_method_accessors", auth_method_accessors)
        if auth_method_types is not None:
            pulumi.set(__self__, "auth_method_types", auth_method_types)
        if identity_entity_ids is not None:
            pulumi.set(__self__, "identity_entity_ids", identity_entity_ids)
        if identity_group_ids is not None:
            pulumi.set(__self__, "identity_group_ids", identity_group_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @_builtins.property
    @pulumi.getter(name="mfaMethodIds")
    def mfa_method_ids(self) -> pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]:
        """
        Set of MFA method UUIDs.
        """
        return pulumi.get(self, "mfa_method_ids")

    @mfa_method_ids.setter
    def mfa_method_ids(self, value: pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]):
        pulumi.set(self, "mfa_method_ids", value)

    @_builtins.property
    @pulumi.getter(name="authMethodAccessors")
    def auth_method_accessors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of auth method accessor IDs.
        """
        return pulumi.get(self, "auth_method_accessors")

    @auth_method_accessors.setter
    def auth_method_accessors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "auth_method_accessors", value)

    @_builtins.property
    @pulumi.getter(name="authMethodTypes")
    def auth_method_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of auth method types.
        """
        return pulumi.get(self, "auth_method_types")

    @auth_method_types.setter
    def auth_method_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "auth_method_types", value)

    @_builtins.property
    @pulumi.getter(name="identityEntityIds")
    def identity_entity_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of identity entity IDs.
        """
        return pulumi.get(self, "identity_entity_ids")

    @identity_entity_ids.setter
    def identity_entity_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "identity_entity_ids", value)

    @_builtins.property
    @pulumi.getter(name="identityGroupIds")
    def identity_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of identity group IDs.
        """
        return pulumi.get(self, "identity_group_ids")

    @identity_group_ids.setter
    def identity_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "identity_group_ids", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Login enforcement name.
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


@pulumi.input_type
class _MfaLoginEnforcementState:
    def __init__(__self__, *,
                 auth_method_accessors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 auth_method_types: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mfa_method_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
                 uuid: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering MfaLoginEnforcement resources.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_accessors: Set of auth method accessor IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_types: Set of auth method types.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_entity_ids: Set of identity entity IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_group_ids: Set of identity group IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] mfa_method_ids: Set of MFA method UUIDs.
        :param pulumi.Input[_builtins.str] name: Login enforcement name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        if auth_method_accessors is not None:
            pulumi.set(__self__, "auth_method_accessors", auth_method_accessors)
        if auth_method_types is not None:
            pulumi.set(__self__, "auth_method_types", auth_method_types)
        if identity_entity_ids is not None:
            pulumi.set(__self__, "identity_entity_ids", identity_entity_ids)
        if identity_group_ids is not None:
            pulumi.set(__self__, "identity_group_ids", identity_group_ids)
        if mfa_method_ids is not None:
            pulumi.set(__self__, "mfa_method_ids", mfa_method_ids)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @_builtins.property
    @pulumi.getter(name="authMethodAccessors")
    def auth_method_accessors(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of auth method accessor IDs.
        """
        return pulumi.get(self, "auth_method_accessors")

    @auth_method_accessors.setter
    def auth_method_accessors(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "auth_method_accessors", value)

    @_builtins.property
    @pulumi.getter(name="authMethodTypes")
    def auth_method_types(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of auth method types.
        """
        return pulumi.get(self, "auth_method_types")

    @auth_method_types.setter
    def auth_method_types(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "auth_method_types", value)

    @_builtins.property
    @pulumi.getter(name="identityEntityIds")
    def identity_entity_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of identity entity IDs.
        """
        return pulumi.get(self, "identity_entity_ids")

    @identity_entity_ids.setter
    def identity_entity_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "identity_entity_ids", value)

    @_builtins.property
    @pulumi.getter(name="identityGroupIds")
    def identity_group_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of identity group IDs.
        """
        return pulumi.get(self, "identity_group_ids")

    @identity_group_ids.setter
    def identity_group_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "identity_group_ids", value)

    @_builtins.property
    @pulumi.getter(name="mfaMethodIds")
    def mfa_method_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        Set of MFA method UUIDs.
        """
        return pulumi.get(self, "mfa_method_ids")

    @mfa_method_ids.setter
    def mfa_method_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "mfa_method_ids", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Login enforcement name.
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
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "uuid", value)


@pulumi.type_token("vault:identity/mfaLoginEnforcement:MfaLoginEnforcement")
class MfaLoginEnforcement(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_accessors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 auth_method_types: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mfa_method_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        Resource for configuring MFA login-enforcement

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaDuo("example",
            secret_key="secret-key",
            integration_key="int-key",
            api_hostname="foo.baz",
            push_info="push-info")
        example_mfa_login_enforcement = vault.identity.MfaLoginEnforcement("example",
            name="default",
            mfa_method_ids=[example.method_id])
        ```

        ## Import

        Resource can be imported using its `name` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaLoginEnforcement:MfaLoginEnforcement example default
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_accessors: Set of auth method accessor IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_types: Set of auth method types.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_entity_ids: Set of identity entity IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_group_ids: Set of identity group IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] mfa_method_ids: Set of MFA method UUIDs.
        :param pulumi.Input[_builtins.str] name: Login enforcement name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaLoginEnforcementArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for configuring MFA login-enforcement

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaDuo("example",
            secret_key="secret-key",
            integration_key="int-key",
            api_hostname="foo.baz",
            push_info="push-info")
        example_mfa_login_enforcement = vault.identity.MfaLoginEnforcement("example",
            name="default",
            mfa_method_ids=[example.method_id])
        ```

        ## Import

        Resource can be imported using its `name` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaLoginEnforcement:MfaLoginEnforcement example default
        ```

        :param str resource_name: The name of the resource.
        :param MfaLoginEnforcementArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaLoginEnforcementArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_method_accessors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 auth_method_types: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 identity_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 mfa_method_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaLoginEnforcementArgs.__new__(MfaLoginEnforcementArgs)

            __props__.__dict__["auth_method_accessors"] = auth_method_accessors
            __props__.__dict__["auth_method_types"] = auth_method_types
            __props__.__dict__["identity_entity_ids"] = identity_entity_ids
            __props__.__dict__["identity_group_ids"] = identity_group_ids
            if mfa_method_ids is None and not opts.urn:
                raise TypeError("Missing required property 'mfa_method_ids'")
            __props__.__dict__["mfa_method_ids"] = mfa_method_ids
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_path"] = None
            __props__.__dict__["uuid"] = None
        super(MfaLoginEnforcement, __self__).__init__(
            'vault:identity/mfaLoginEnforcement:MfaLoginEnforcement',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            auth_method_accessors: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            auth_method_types: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            identity_entity_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            identity_group_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            mfa_method_ids: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
            uuid: Optional[pulumi.Input[_builtins.str]] = None) -> 'MfaLoginEnforcement':
        """
        Get an existing MfaLoginEnforcement resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_accessors: Set of auth method accessor IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] auth_method_types: Set of auth method types.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_entity_ids: Set of identity entity IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] identity_group_ids: Set of identity group IDs.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] mfa_method_ids: Set of MFA method UUIDs.
        :param pulumi.Input[_builtins.str] name: Login enforcement name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaLoginEnforcementState.__new__(_MfaLoginEnforcementState)

        __props__.__dict__["auth_method_accessors"] = auth_method_accessors
        __props__.__dict__["auth_method_types"] = auth_method_types
        __props__.__dict__["identity_entity_ids"] = identity_entity_ids
        __props__.__dict__["identity_group_ids"] = identity_group_ids
        __props__.__dict__["mfa_method_ids"] = mfa_method_ids
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["uuid"] = uuid
        return MfaLoginEnforcement(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="authMethodAccessors")
    def auth_method_accessors(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Set of auth method accessor IDs.
        """
        return pulumi.get(self, "auth_method_accessors")

    @_builtins.property
    @pulumi.getter(name="authMethodTypes")
    def auth_method_types(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Set of auth method types.
        """
        return pulumi.get(self, "auth_method_types")

    @_builtins.property
    @pulumi.getter(name="identityEntityIds")
    def identity_entity_ids(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Set of identity entity IDs.
        """
        return pulumi.get(self, "identity_entity_ids")

    @_builtins.property
    @pulumi.getter(name="identityGroupIds")
    def identity_group_ids(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        Set of identity group IDs.
        """
        return pulumi.get(self, "identity_group_ids")

    @_builtins.property
    @pulumi.getter(name="mfaMethodIds")
    def mfa_method_ids(self) -> pulumi.Output[Sequence[_builtins.str]]:
        """
        Set of MFA method UUIDs.
        """
        return pulumi.get(self, "mfa_method_ids")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Login enforcement name.
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
    @pulumi.getter
    def uuid(self) -> pulumi.Output[_builtins.str]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

