# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs

__all__ = [
    'AuthBackendTune',
    'GetPolicyDocumentRuleResult',
    'GetPolicyDocumentRuleAllowedParameterResult',
    'GetPolicyDocumentRuleDeniedParameterResult',
]

@pulumi.output_type
class AuthBackendTune(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "allowedResponseHeaders":
            suggest = "allowed_response_headers"
        elif key == "auditNonHmacRequestKeys":
            suggest = "audit_non_hmac_request_keys"
        elif key == "auditNonHmacResponseKeys":
            suggest = "audit_non_hmac_response_keys"
        elif key == "defaultLeaseTtl":
            suggest = "default_lease_ttl"
        elif key == "listingVisibility":
            suggest = "listing_visibility"
        elif key == "maxLeaseTtl":
            suggest = "max_lease_ttl"
        elif key == "passthroughRequestHeaders":
            suggest = "passthrough_request_headers"
        elif key == "tokenType":
            suggest = "token_type"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in AuthBackendTune. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        AuthBackendTune.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        AuthBackendTune.__key_warning(key)
        return super().get(key, default)

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
        :param Sequence[str] allowed_response_headers: List of headers to whitelist and allowing
               a plugin to include them in the response.
        :param Sequence[str] audit_non_hmac_request_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the request data object.
        :param Sequence[str] audit_non_hmac_response_keys: Specifies the list of keys that will
               not be HMAC'd by audit devices in the response data object.
        :param str default_lease_ttl: Specifies the default time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param str listing_visibility: Specifies whether to show this mount in
               the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        :param str max_lease_ttl: Specifies the maximum time-to-live.
               If set, this overrides the global default.
               Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        :param Sequence[str] passthrough_request_headers: List of headers to whitelist and
               pass from the request to the backend.
        :param str token_type: Specifies the type of tokens that should be returned by
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
    def allowed_response_headers(self) -> Optional[Sequence[str]]:
        """
        List of headers to whitelist and allowing
        a plugin to include them in the response.
        """
        return pulumi.get(self, "allowed_response_headers")

    @property
    @pulumi.getter(name="auditNonHmacRequestKeys")
    def audit_non_hmac_request_keys(self) -> Optional[Sequence[str]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the request data object.
        """
        return pulumi.get(self, "audit_non_hmac_request_keys")

    @property
    @pulumi.getter(name="auditNonHmacResponseKeys")
    def audit_non_hmac_response_keys(self) -> Optional[Sequence[str]]:
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the response data object.
        """
        return pulumi.get(self, "audit_non_hmac_response_keys")

    @property
    @pulumi.getter(name="defaultLeaseTtl")
    def default_lease_ttl(self) -> Optional[str]:
        """
        Specifies the default time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "default_lease_ttl")

    @property
    @pulumi.getter(name="listingVisibility")
    def listing_visibility(self) -> Optional[str]:
        """
        Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        """
        return pulumi.get(self, "listing_visibility")

    @property
    @pulumi.getter(name="maxLeaseTtl")
    def max_lease_ttl(self) -> Optional[str]:
        """
        Specifies the maximum time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        return pulumi.get(self, "max_lease_ttl")

    @property
    @pulumi.getter(name="passthroughRequestHeaders")
    def passthrough_request_headers(self) -> Optional[Sequence[str]]:
        """
        List of headers to whitelist and
        pass from the request to the backend.
        """
        return pulumi.get(self, "passthrough_request_headers")

    @property
    @pulumi.getter(name="tokenType")
    def token_type(self) -> Optional[str]:
        """
        Specifies the type of tokens that should be returned by
        the mount. Valid values are "default-service", "default-batch", "service", "batch".
        """
        return pulumi.get(self, "token_type")


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
        Description of the rule. Will be added as a comment to rendered rule.
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
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")


@pulumi.output_type
class GetPolicyDocumentRuleDeniedParameterResult(dict):
    def __init__(__self__, *,
                 key: str,
                 values: Sequence[str]):
        pulumi.set(__self__, "key", key)
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def key(self) -> str:
        return pulumi.get(self, "key")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        return pulumi.get(self, "values")


