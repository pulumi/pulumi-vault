# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = [
    'AuthBackendTuneArgs',
    'AuthBackendTuneArgsDict',
]

MYPY = False

if not MYPY:
    class AuthBackendTuneArgsDict(TypedDict):
        allowed_response_headers: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        List of headers to whitelist and allowing
        a plugin to include them in the response.
        """
        audit_non_hmac_request_keys: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the request data object.
        """
        audit_non_hmac_response_keys: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        Specifies the list of keys that will
        not be HMAC'd by audit devices in the response data object.
        """
        default_lease_ttl: NotRequired[pulumi.Input[str]]
        """
        Specifies the default time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        listing_visibility: NotRequired[pulumi.Input[str]]
        """
        Specifies whether to show this mount in
        the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
        """
        max_lease_ttl: NotRequired[pulumi.Input[str]]
        """
        Specifies the maximum time-to-live.
        If set, this overrides the global default.
        Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
        """
        passthrough_request_headers: NotRequired[pulumi.Input[Sequence[pulumi.Input[str]]]]
        """
        List of headers to whitelist and
        pass from the request to the backend.
        """
        token_type: NotRequired[pulumi.Input[str]]
        """
        Specifies the type of tokens that should be returned by
        the mount. Valid values are "default-service", "default-batch", "service", "batch".
        """
elif False:
    AuthBackendTuneArgsDict: TypeAlias = Mapping[str, Any]

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


