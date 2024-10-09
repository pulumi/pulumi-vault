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
    'AuthBackendTune',
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


