# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'AuthBackendTuneArgs',
    'ProviderAuthLoginArgs',
    'ProviderClientAuthArgs',
    'ProviderHeaderArgs',
    'GetPolicyDocumentRuleArgs',
    'GetPolicyDocumentRuleAllowedParameterArgs',
    'GetPolicyDocumentRuleDeniedParameterArgs',
]

@pulumi.input_type
class AuthBackendTuneArgs:
    def __init__(__self__, *,
                 allowed_response_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_request_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 audit_non_hmac_response_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 default_lease_ttl: Optional[pulumi.Input[str]] = None,
                 listing_visibility: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl: Optional[pulumi.Input[str]] = None,
                 passthrough_request_headers: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_type: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_response_headers: List of headers to whitelist and allowing
               a plugin to include them in the response.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_request_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the request data object.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] audit_non_hmac_response_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the response data object.
        :param pulumi.Input[str] default_lease_ttl: Specifies the default time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param pulumi.Input[str] listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param pulumi.Input[str] max_lease_ttl: Specifies the maximum time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param pulumi.Input[Sequence[pulumi.Input[str]]] passthrough_request_headers: List of headers to whitelist and
               pass from the request to the backend.
        :param pulumi.Input[str] token_type: Specifies the type of tokens that should be returned by
               the mount. Valid values are "default-service", "default-batch", "service", "batch".
        """
        if allowed_response_headers is not None:
            pulumi.set(__self__, "allowed_response_headers", allowed_response_headers)
        if audit_non_hmac_request_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_request_keys", audit_non_hmac_request_keys)
        if audit_non_hmac_response_keys is not None:
            pulumi.set(__self__, "audit_non_hmac_response_keys", audit_non_hmac_response_keys)
        if default_lease_ttl is not None:
            pulumi.set(__self__, "default_lease_ttl", default_lease_ttl)
        if listing_visibility is not None:
            pulumi.set(__self__, "listing_visibility", listing_visibility)
        if max_lease_ttl is not None:
            pulumi.set(__self__, "max_lease_ttl", max_lease_ttl)
        if passthrough_request_headers is not None:
            pulumi.set(__self__, "passthrough_request_headers", passthrough_request_headers)
        if token_type is not None:
            pulumi.set(__self__, "token_type", token_type)

    @property
    @pulumi.getter(name="allowedResponseHeaders")
    def allowed_response_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of headers to whitelist and allowing
        a plugin to include them in the response.
        """
        return pulumi.get(self, "allowed_response_headers")

    @allowed_response_headers.setter
    def allowed_response_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_response_headers", value)

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @audit_non_hmac_request_keys.setter
    def audit_non_hmac_request_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_request_keys", value)

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @audit_non_hmac_response_keys.setter
    def audit_non_hmac_response_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "audit_non_hmac_response_keys", value)

    @property
    @pulumi.getter(name="defaultLeaseTtl")
    def default_lease_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the default time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "default_lease_ttl")

    @default_lease_ttl.setter
    def default_lease_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_lease_ttl", value)

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        """
        return pulumi.get(self, "listing_visibility")

    @listing_visibility.setter
    def listing_visibility(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "listing_visibility", value)

    @property
    @pulumi.getter(name="maxLeaseTtl")
    def max_lease_ttl(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the maximum time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "max_lease_ttl")

    @max_lease_ttl.setter
    def max_lease_ttl(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "max_lease_ttl", value)

    @property
    @pulumi.getter(name="passthroughRequestHeaders")
    def passthrough_request_headers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of headers to whitelist and
        pass from the request to the backend.
        """
        return pulumi.get(self, "passthrough_request_headers")

    @passthrough_request_headers.setter
    def passthrough_request_headers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "passthrough_request_headers", value)

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of tokens that should be returned by
        the mount. Valid values are "default-service", "default-batch", "service", "batch".
        """
        return pulumi.get(self, "token_type")

    @token_type.setter
    def token_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_type", value)


@pulumi.input_type
class ProviderAuthLoginArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        pulumi.set(__self__, "path", path)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "parameters", value)


@pulumi.input_type
class ProviderClientAuthArgs:
    def __init__(__self__, *,
                 cert_file: pulumi.Input[str],
                 key_file: pulumi.Input[str]):
        pulumi.set(__self__, "cert_file", cert_file)
        pulumi.set(__self__, "key_file", key_file)

    @property
    @pulumi.getter(name="certFile")
    def cert_file(self) -> pulumi.Input[str]:
        return pulumi.get(self, "cert_file")

    @cert_file.setter
    def cert_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "cert_file", value)

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> pulumi.Input[str]:
        return pulumi.get(self, "key_file")

    @key_file.setter
    def key_file(self, value: pulumi.Input[str]):
        pulumi.set(self, "key_file", value)


@pulumi.input_type
class ProviderHeaderArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value: pulumi.Input[str]):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[str]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class GetPolicyDocumentRuleArgs:
    def __init__(__self__, *,
                 capabilities: Sequence[str],
                 path: str,
                 allowed_parameters: Optional[Sequence['GetPolicyDocumentRuleAllowedParameterArgs']] = None,
                 denied_parameters: Optional[Sequence['GetPolicyDocumentRuleDeniedParameterArgs']] = None,
                 description: Optional[str] = None,
                 max_wrapping_ttl: Optional[str] = None,
                 min_wrapping_ttl: Optional[str] = None,
                 required_parameters: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] capabilities: A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        :param str path: A path in Vault that this rule applies to.
        :param Sequence['GetPolicyDocumentRuleAllowedParameterArgs'] allowed_parameters: Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        :param Sequence['GetPolicyDocumentRuleDeniedParameterArgs'] denied_parameters: Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        :param str description: Description of the rule. Will be added as a comment to rendered rule.
        :param str max_wrapping_ttl: The maximum allowed TTL that clients can specify for a wrapped response.
        :param str min_wrapping_ttl: The minimum allowed TTL that clients can specify for a wrapped response.
        :param Sequence[str] required_parameters: A list of parameters that must be specified.
        """
        pulumi.set(__self__, "capabilities", capabilities)
        pulumi.set(__self__, "path", path)
        if allowed_parameters is not None:
            pulumi.set(__self__, "allowed_parameters", allowed_parameters)
        if denied_parameters is not None:
            pulumi.set(__self__, "denied_parameters", denied_parameters)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if max_wrapping_ttl is not None:
            pulumi.set(__self__, "max_wrapping_ttl", max_wrapping_ttl)
        if min_wrapping_ttl is not None:
            pulumi.set(__self__, "min_wrapping_ttl", min_wrapping_ttl)
        if required_parameters is not None:
            pulumi.set(__self__, "required_parameters", required_parameters)

    @property
    @pulumi.getter
    def capabilities(self) -> Sequence[str]:
        """
        A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        """
        return pulumi.get(self, "capabilities")

    @capabilities.setter
    def capabilities(self, value: Sequence[str]):
        pulumi.set(self, "capabilities", value)

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        A path in Vault that this rule applies to.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: str):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter(name="allowedParameters")
    def allowed_parameters(self) -> Optional[Sequence['GetPolicyDocumentRuleAllowedParameterArgs']]:
        """
        Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        """
        return pulumi.get(self, "allowed_parameters")

    @allowed_parameters.setter
    def allowed_parameters(self, value: Optional[Sequence['GetPolicyDocumentRuleAllowedParameterArgs']]):
        pulumi.set(self, "allowed_parameters", value)

    @property
    @pulumi.getter(name="deniedParameters")
    def denied_parameters(self) -> Optional[Sequence['GetPolicyDocumentRuleDeniedParameterArgs']]:
        """
        Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        """
        return pulumi.get(self, "denied_parameters")

    @denied_parameters.setter
    def denied_parameters(self, value: Optional[Sequence['GetPolicyDocumentRuleDeniedParameterArgs']]):
        pulumi.set(self, "denied_parameters", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the rule. Will be added as a comment to rendered rule.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="maxWrappingTtl")
    def max_wrapping_ttl(self) -> Optional[str]:
        """
        The maximum allowed TTL that clients can specify for a wrapped response.
        """
        return pulumi.get(self, "max_wrapping_ttl")

    @max_wrapping_ttl.setter
    def max_wrapping_ttl(self, value: Optional[str]):
        pulumi.set(self, "max_wrapping_ttl", value)

    @property
    @pulumi.getter(name="minWrappingTtl")
    def min_wrapping_ttl(self) -> Optional[str]:
        """
        The minimum allowed TTL that clients can specify for a wrapped response.
        """
        return pulumi.get(self, "min_wrapping_ttl")

    @min_wrapping_ttl.setter
    def min_wrapping_ttl(self, value: Optional[str]):
        pulumi.set(self, "min_wrapping_ttl", value)

    @property
    @pulumi.getter(name="requiredParameters")
    def required_parameters(self) -> Optional[Sequence[str]]:
        """
        A list of parameters that must be specified.
        """
        return pulumi.get(self, "required_parameters")

    @required_parameters.setter
    def required_parameters(self, value: Optional[Sequence[str]]):
        pulumi.set(self, "required_parameters", value)


@pulumi.input_type
class GetPolicyDocumentRuleAllowedParameterArgs:
    def __init__(__self__, *,
                 key: str,
                 values: Sequence[str]):
        """
        :param str key: name of permitted or denied parameter.
        :param Sequence[str] values: list of values what are permitted or denied by policy rule.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        name of permitted or denied parameter.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        list of values what are permitted or denied by policy rule.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class GetPolicyDocumentRuleDeniedParameterArgs:
    def __init__(__self__, *,
                 key: str,
                 values: Sequence[str]):
        """
        :param str key: name of permitted or denied parameter.
        :param Sequence[str] values: list of values what are permitted or denied by policy rule.
        """
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        """
        name of permitted or denied parameter.
        """
        return pulumi.get(self, "key")

    @key.setter
    def key(self, value: str):
        pulumi.set(self, "key", value)

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        list of values what are permitted or denied by policy rule.
        """
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Sequence[str]):
        pulumi.set(self, "values", value)


