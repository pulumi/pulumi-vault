# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
import copy
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
from . import outputs
from ._inputs import *

__all__ = ['BackendConfigCmpv2Args', 'BackendConfigCmpv2']

@pulumi.input_type
class BackendConfigCmpv2Args:
    def __init__(__self__, *,
                 backend: pulumi.Input[builtins.str],
                 audit_fields: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 authenticators: Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']] = None,
                 default_path_policy: Optional[pulumi.Input[builtins.str]] = None,
                 disabled_validations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 enable_sentinel_parsing: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a BackendConfigCmpv2 resource.
        :param pulumi.Input[builtins.str] backend: The path to the PKI secret backend to
               read the CMPv2 configuration from, with no leading or trailing `/`s.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] audit_fields: Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        :param pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs'] authenticators: Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        :param pulumi.Input[builtins.str] default_path_policy: Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] disabled_validations: A comma-separated list of validations not to perform on CMPv2 messages.
               
               <a id="nestedatt--authenticators"></a>
        :param pulumi.Input[builtins.bool] enable_sentinel_parsing: If set, parse out fields from the provided CSR making them available for Sentinel policies.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether CMPv2 is enabled.
        :param pulumi.Input[builtins.str] namespace: The namespace of the target resource.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "backend", backend)
        if audit_fields is not None:
            pulumi.set(__self__, "audit_fields", audit_fields)
        if authenticators is not None:
            pulumi.set(__self__, "authenticators", authenticators)
        if default_path_policy is not None:
            pulumi.set(__self__, "default_path_policy", default_path_policy)
        if disabled_validations is not None:
            pulumi.set(__self__, "disabled_validations", disabled_validations)
        if enable_sentinel_parsing is not None:
            pulumi.set(__self__, "enable_sentinel_parsing", enable_sentinel_parsing)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[builtins.str]:
        """
        The path to the PKI secret backend to
        read the CMPv2 configuration from, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="auditFields")
    def audit_fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        """
        return pulumi.get(self, "audit_fields")

    @audit_fields.setter
    def audit_fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "audit_fields", value)

    @property
    @pulumi.getter
    def authenticators(self) -> Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']]:
        """
        Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        """
        return pulumi.get(self, "authenticators")

    @authenticators.setter
    def authenticators(self, value: Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']]):
        pulumi.set(self, "authenticators", value)

    @property
    @pulumi.getter(name="defaultPathPolicy")
    def default_path_policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        """
        return pulumi.get(self, "default_path_policy")

    @default_path_policy.setter
    def default_path_policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_path_policy", value)

    @property
    @pulumi.getter(name="disabledValidations")
    def disabled_validations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A comma-separated list of validations not to perform on CMPv2 messages.

        <a id="nestedatt--authenticators"></a>
        """
        return pulumi.get(self, "disabled_validations")

    @disabled_validations.setter
    def disabled_validations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "disabled_validations", value)

    @property
    @pulumi.getter(name="enableSentinelParsing")
    def enable_sentinel_parsing(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set, parse out fields from the provided CSR making them available for Sentinel policies.
        """
        return pulumi.get(self, "enable_sentinel_parsing")

    @enable_sentinel_parsing.setter
    def enable_sentinel_parsing(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_sentinel_parsing", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether CMPv2 is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace of the target resource.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _BackendConfigCmpv2State:
    def __init__(__self__, *,
                 audit_fields: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 authenticators: Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 default_path_policy: Optional[pulumi.Input[builtins.str]] = None,
                 disabled_validations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 enable_sentinel_parsing: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 last_updated: Optional[pulumi.Input[builtins.str]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering BackendConfigCmpv2 resources.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] audit_fields: Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        :param pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs'] authenticators: Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        :param pulumi.Input[builtins.str] backend: The path to the PKI secret backend to
               read the CMPv2 configuration from, with no leading or trailing `/`s.
        :param pulumi.Input[builtins.str] default_path_policy: Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] disabled_validations: A comma-separated list of validations not to perform on CMPv2 messages.
               
               <a id="nestedatt--authenticators"></a>
        :param pulumi.Input[builtins.bool] enable_sentinel_parsing: If set, parse out fields from the provided CSR making them available for Sentinel policies.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether CMPv2 is enabled.
        :param pulumi.Input[builtins.str] last_updated: A read-only timestamp representing the last time the configuration was updated.
        :param pulumi.Input[builtins.str] namespace: The namespace of the target resource.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if audit_fields is not None:
            pulumi.set(__self__, "audit_fields", audit_fields)
        if authenticators is not None:
            pulumi.set(__self__, "authenticators", authenticators)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if default_path_policy is not None:
            pulumi.set(__self__, "default_path_policy", default_path_policy)
        if disabled_validations is not None:
            pulumi.set(__self__, "disabled_validations", disabled_validations)
        if enable_sentinel_parsing is not None:
            pulumi.set(__self__, "enable_sentinel_parsing", enable_sentinel_parsing)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if last_updated is not None:
            pulumi.set(__self__, "last_updated", last_updated)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="auditFields")
    def audit_fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        """
        return pulumi.get(self, "audit_fields")

    @audit_fields.setter
    def audit_fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "audit_fields", value)

    @property
    @pulumi.getter
    def authenticators(self) -> Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']]:
        """
        Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        """
        return pulumi.get(self, "authenticators")

    @authenticators.setter
    def authenticators(self, value: Optional[pulumi.Input['BackendConfigCmpv2AuthenticatorsArgs']]):
        pulumi.set(self, "authenticators", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The path to the PKI secret backend to
        read the CMPv2 configuration from, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="defaultPathPolicy")
    def default_path_policy(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        """
        return pulumi.get(self, "default_path_policy")

    @default_path_policy.setter
    def default_path_policy(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "default_path_policy", value)

    @property
    @pulumi.getter(name="disabledValidations")
    def disabled_validations(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        A comma-separated list of validations not to perform on CMPv2 messages.

        <a id="nestedatt--authenticators"></a>
        """
        return pulumi.get(self, "disabled_validations")

    @disabled_validations.setter
    def disabled_validations(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "disabled_validations", value)

    @property
    @pulumi.getter(name="enableSentinelParsing")
    def enable_sentinel_parsing(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        If set, parse out fields from the provided CSR making them available for Sentinel policies.
        """
        return pulumi.get(self, "enable_sentinel_parsing")

    @enable_sentinel_parsing.setter
    def enable_sentinel_parsing(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enable_sentinel_parsing", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[builtins.bool]]:
        """
        Specifies whether CMPv2 is enabled.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[builtins.bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A read-only timestamp representing the last time the configuration was updated.
        """
        return pulumi.get(self, "last_updated")

    @last_updated.setter
    def last_updated(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "last_updated", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace of the target resource.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)


class BackendConfigCmpv2(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_fields: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 authenticators: Optional[pulumi.Input[Union['BackendConfigCmpv2AuthenticatorsArgs', 'BackendConfigCmpv2AuthenticatorsArgsDict']]] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 default_path_policy: Optional[pulumi.Input[builtins.str]] = None,
                 disabled_validations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 enable_sentinel_parsing: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Allows setting the CMPv2 configuration on a PKI Secret Backend

        ## Import

        The PKI config cluster can be imported using the resource's `id`.
        In the case of the example above the `id` would be `pki-root/config/cmpv2`,
        where the `pki-root` component is the resource's `backend`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/backendConfigCmpv2:BackendConfigCmpv2 example pki-root/config/cmpv2
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] audit_fields: Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        :param pulumi.Input[Union['BackendConfigCmpv2AuthenticatorsArgs', 'BackendConfigCmpv2AuthenticatorsArgsDict']] authenticators: Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        :param pulumi.Input[builtins.str] backend: The path to the PKI secret backend to
               read the CMPv2 configuration from, with no leading or trailing `/`s.
        :param pulumi.Input[builtins.str] default_path_policy: Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] disabled_validations: A comma-separated list of validations not to perform on CMPv2 messages.
               
               <a id="nestedatt--authenticators"></a>
        :param pulumi.Input[builtins.bool] enable_sentinel_parsing: If set, parse out fields from the provided CSR making them available for Sentinel policies.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether CMPv2 is enabled.
        :param pulumi.Input[builtins.str] namespace: The namespace of the target resource.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: BackendConfigCmpv2Args,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Allows setting the CMPv2 configuration on a PKI Secret Backend

        ## Import

        The PKI config cluster can be imported using the resource's `id`.
        In the case of the example above the `id` would be `pki-root/config/cmpv2`,
        where the `pki-root` component is the resource's `backend`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/backendConfigCmpv2:BackendConfigCmpv2 example pki-root/config/cmpv2
        ```

        :param str resource_name: The name of the resource.
        :param BackendConfigCmpv2Args args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(BackendConfigCmpv2Args, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 audit_fields: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 authenticators: Optional[pulumi.Input[Union['BackendConfigCmpv2AuthenticatorsArgs', 'BackendConfigCmpv2AuthenticatorsArgsDict']]] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 default_path_policy: Optional[pulumi.Input[builtins.str]] = None,
                 disabled_validations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 enable_sentinel_parsing: Optional[pulumi.Input[builtins.bool]] = None,
                 enabled: Optional[pulumi.Input[builtins.bool]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = BackendConfigCmpv2Args.__new__(BackendConfigCmpv2Args)

            __props__.__dict__["audit_fields"] = audit_fields
            __props__.__dict__["authenticators"] = authenticators
            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["default_path_policy"] = default_path_policy
            __props__.__dict__["disabled_validations"] = disabled_validations
            __props__.__dict__["enable_sentinel_parsing"] = enable_sentinel_parsing
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["last_updated"] = None
        super(BackendConfigCmpv2, __self__).__init__(
            'vault:pkiSecret/backendConfigCmpv2:BackendConfigCmpv2',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            audit_fields: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            authenticators: Optional[pulumi.Input[Union['BackendConfigCmpv2AuthenticatorsArgs', 'BackendConfigCmpv2AuthenticatorsArgsDict']]] = None,
            backend: Optional[pulumi.Input[builtins.str]] = None,
            default_path_policy: Optional[pulumi.Input[builtins.str]] = None,
            disabled_validations: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            enable_sentinel_parsing: Optional[pulumi.Input[builtins.bool]] = None,
            enabled: Optional[pulumi.Input[builtins.bool]] = None,
            last_updated: Optional[pulumi.Input[builtins.str]] = None,
            namespace: Optional[pulumi.Input[builtins.str]] = None) -> 'BackendConfigCmpv2':
        """
        Get an existing BackendConfigCmpv2 resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] audit_fields: Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        :param pulumi.Input[Union['BackendConfigCmpv2AuthenticatorsArgs', 'BackendConfigCmpv2AuthenticatorsArgsDict']] authenticators: Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        :param pulumi.Input[builtins.str] backend: The path to the PKI secret backend to
               read the CMPv2 configuration from, with no leading or trailing `/`s.
        :param pulumi.Input[builtins.str] default_path_policy: Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] disabled_validations: A comma-separated list of validations not to perform on CMPv2 messages.
               
               <a id="nestedatt--authenticators"></a>
        :param pulumi.Input[builtins.bool] enable_sentinel_parsing: If set, parse out fields from the provided CSR making them available for Sentinel policies.
        :param pulumi.Input[builtins.bool] enabled: Specifies whether CMPv2 is enabled.
        :param pulumi.Input[builtins.str] last_updated: A read-only timestamp representing the last time the configuration was updated.
        :param pulumi.Input[builtins.str] namespace: The namespace of the target resource.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _BackendConfigCmpv2State.__new__(_BackendConfigCmpv2State)

        __props__.__dict__["audit_fields"] = audit_fields
        __props__.__dict__["authenticators"] = authenticators
        __props__.__dict__["backend"] = backend
        __props__.__dict__["default_path_policy"] = default_path_policy
        __props__.__dict__["disabled_validations"] = disabled_validations
        __props__.__dict__["enable_sentinel_parsing"] = enable_sentinel_parsing
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["last_updated"] = last_updated
        __props__.__dict__["namespace"] = namespace
        return BackendConfigCmpv2(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="auditFields")
    def audit_fields(self) -> pulumi.Output[Sequence[builtins.str]]:
        """
        Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        """
        return pulumi.get(self, "audit_fields")

    @property
    @pulumi.getter
    def authenticators(self) -> pulumi.Output['outputs.BackendConfigCmpv2Authenticators']:
        """
        Lists the mount accessors CMPv2 should delegate authentication requests towards (see below for nested schema).
        """
        return pulumi.get(self, "authenticators")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[builtins.str]:
        """
        The path to the PKI secret backend to
        read the CMPv2 configuration from, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="defaultPathPolicy")
    def default_path_policy(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Specifies the behavior for requests using the non-role-qualified CMPv2 requests. Can be sign-verbatim or a role given by role:<role_name>.
        """
        return pulumi.get(self, "default_path_policy")

    @property
    @pulumi.getter(name="disabledValidations")
    def disabled_validations(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        A comma-separated list of validations not to perform on CMPv2 messages.

        <a id="nestedatt--authenticators"></a>
        """
        return pulumi.get(self, "disabled_validations")

    @property
    @pulumi.getter(name="enableSentinelParsing")
    def enable_sentinel_parsing(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        If set, parse out fields from the provided CSR making them available for Sentinel policies.
        """
        return pulumi.get(self, "enable_sentinel_parsing")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[builtins.bool]]:
        """
        Specifies whether CMPv2 is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> pulumi.Output[builtins.str]:
        """
        A read-only timestamp representing the last time the configuration was updated.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The namespace of the target resource.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

