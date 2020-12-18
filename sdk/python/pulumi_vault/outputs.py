# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from . import _utilities, _tables
from . import outputs

__all__ = [
    'AuthBackendTune',
    'ProviderAuthLogin',
    'ProviderClientAuth',
    'ProviderHeader',
    'GetPolicyDocumentRuleResult',
    'GetPolicyDocumentRuleAllowedParameterResult',
    'GetPolicyDocumentRuleDeniedParameterResult',
]

@pulumi.output_type
class AuthBackendTune(dict):
    def __init__(__self__, *,
                 allowed_response_headers: Optional[Sequence[str]] = None,
                 audit_non_hmac_request_keys: Optional[Sequence[str]] = None,
                 audit_non_hmac_response_keys: Optional[Sequence[str]] = None,
                 default_lease_ttl: Optional[str] = None,
                 listing_visibility: Optional[str] = None,
                 max_lease_ttl: Optional[str] = None,
                 passthrough_request_headers: Optional[Sequence[str]] = None,
                 token_type: Optional[str] = None):
        """
        :param str listing_visibility: Speficies whether to show this mount in the UI-specific listing endpoint.
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
    def allowed_response_headers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "allowed_response_headers")

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @property
    @pulumi.getter(name="defaultLeaseTtl")
    def default_lease_ttl(self) -> Optional[str]:
        return pulumi.get(self, "default_lease_ttl")

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> Optional[str]:
        """
        Speficies whether to show this mount in the UI-specific listing endpoint.
        """
        return pulumi.get(self, "listing_visibility")

    @property
    @pulumi.getter(name="maxLeaseTtl")
    def max_lease_ttl(self) -> Optional[str]:
        return pulumi.get(self, "max_lease_ttl")

    @property
    @pulumi.getter(name="passthroughRequestHeaders")
    def passthrough_request_headers(self) -> Optional[Sequence[str]]:
        return pulumi.get(self, "passthrough_request_headers")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[str]:
        return pulumi.get(self, "token_type")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderAuthLogin(dict):
    def __init__(__self__, *,
                 path: str,
                 namespace: Optional[str] = None,
                 parameters: Optional[Mapping[str, str]] = None):
        pulumi.set(__self__, "path", path)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def parameters(self) -> Optional[Mapping[str, str]]:
        return pulumi.get(self, "parameters")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderClientAuth(dict):
    def __init__(__self__, *,
                 cert_file: str,
                 key_file: str):
        pulumi.set(__self__, "cert_file", cert_file)
        pulumi.set(__self__, "key_file", key_file)

    @property
    @pulumi.getter(name="certFile")
    def cert_file(self) -> str:
        return pulumi.get(self, "cert_file")

    @property
    @pulumi.getter(name="keyFile")
    def key_file(self) -> str:
        return pulumi.get(self, "key_file")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class ProviderHeader(dict):
    def __init__(__self__, *,
                 name: str,
                 value: str):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def value(self) -> str:
        return pulumi.get(self, "value")

    def _translate_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop


@pulumi.output_type
class GetPolicyDocumentRuleResult(dict):
    def __init__(__self__, *,
                 capabilities: Sequence[str],
                 path: str,
                 allowed_parameters: Optional[Sequence['outputs.GetPolicyDocumentRuleAllowedParameterResult']] = None,
                 denied_parameters: Optional[Sequence['outputs.GetPolicyDocumentRuleDeniedParameterResult']] = None,
                 description: Optional[str] = None,
                 max_wrapping_ttl: Optional[str] = None,
                 min_wrapping_ttl: Optional[str] = None,
                 required_parameters: Optional[Sequence[str]] = None):
        """
        :param Sequence[str] capabilities: A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
        :param str path: A path in Vault that this rule applies to.
        :param Sequence['GetPolicyDocumentRuleAllowedParameterArgs'] allowed_parameters: Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        :param Sequence['GetPolicyDocumentRuleDeniedParameterArgs'] denied_parameters: Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        :param str description: Description of the rule. Will be added as a commend to rendered rule.
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

    @property
    @pulumi.getter
    def path(self) -> str:
        """
        A path in Vault that this rule applies to.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="allowedParameters")
    def allowed_parameters(self) -> Optional[Sequence['outputs.GetPolicyDocumentRuleAllowedParameterResult']]:
        """
        Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
        """
        return pulumi.get(self, "allowed_parameters")

    @property
    @pulumi.getter(name="deniedParameters")
    def denied_parameters(self) -> Optional[Sequence['outputs.GetPolicyDocumentRuleDeniedParameterResult']]:
        """
        Blacklists a list of parameter and values. Any values specified here take precedence over `allowed_parameter`. See Parameters below.
        """
        return pulumi.get(self, "denied_parameters")

    @property
    @pulumi.getter
    def description(self) -> Optional[str]:
        """
        Description of the rule. Will be added as a commend to rendered rule.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="maxWrappingTtl")
    def max_wrapping_ttl(self) -> Optional[str]:
        """
        The maximum allowed TTL that clients can specify for a wrapped response.
        """
        return pulumi.get(self, "max_wrapping_ttl")

    @property
    @pulumi.getter(name="minWrappingTtl")
    def min_wrapping_ttl(self) -> Optional[str]:
        """
        The minimum allowed TTL that clients can specify for a wrapped response.
        """
        return pulumi.get(self, "min_wrapping_ttl")

    @property
    @pulumi.getter(name="requiredParameters")
    def required_parameters(self) -> Optional[Sequence[str]]:
        """
        A list of parameters that must be specified.
        """
        return pulumi.get(self, "required_parameters")


@pulumi.output_type
class GetPolicyDocumentRuleAllowedParameterResult(dict):
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

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        list of values what are permitted or denied by policy rule.
        """
        return pulumi.get(self, "values")


@pulumi.output_type
class GetPolicyDocumentRuleDeniedParameterResult(dict):
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

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        list of values what are permitted or denied by policy rule.
        """
        return pulumi.get(self, "values")


