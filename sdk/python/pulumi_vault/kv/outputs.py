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
    'SecretV2CustomMetadata',
]

@pulumi.output_type
class SecretV2CustomMetadata(dict):
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "casRequired":
            suggest = "cas_required"
        elif key == "deleteVersionAfter":
            suggest = "delete_version_after"
        elif key == "maxVersions":
            suggest = "max_versions"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in SecretV2CustomMetadata. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        SecretV2CustomMetadata.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        SecretV2CustomMetadata.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 cas_required: Optional[bool] = None,
                 data: Optional[Mapping[str, str]] = None,
                 delete_version_after: Optional[int] = None,
                 max_versions: Optional[int] = None):
        """
        :param bool cas_required: If true, all keys will require the cas parameter to be set on all write requests.
        :param Mapping[str, str] data: A mapping whose keys are the top-level data keys returned from
               Vault and whose values are the corresponding values. This map can only
               represent string data, so any non-string values returned from Vault are
               serialized as JSON.
        :param int delete_version_after: If set, specifies the length of time before a version is deleted.
        :param int max_versions: The number of versions to keep per key.
        """
        if cas_required is not None:
            pulumi.set(__self__, "cas_required", cas_required)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if delete_version_after is not None:
            pulumi.set(__self__, "delete_version_after", delete_version_after)
        if max_versions is not None:
            pulumi.set(__self__, "max_versions", max_versions)

    @property
    @pulumi.getter(name="casRequired")
    def cas_required(self) -> Optional[bool]:
        """
        If true, all keys will require the cas parameter to be set on all write requests.
        """
        return pulumi.get(self, "cas_required")

    @property
    @pulumi.getter
    def data(self) -> Optional[Mapping[str, str]]:
        """
        A mapping whose keys are the top-level data keys returned from
        Vault and whose values are the corresponding values. This map can only
        represent string data, so any non-string values returned from Vault are
        serialized as JSON.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="deleteVersionAfter")
    def delete_version_after(self) -> Optional[int]:
        """
        If set, specifies the length of time before a version is deleted.
        """
        return pulumi.get(self, "delete_version_after")

    @property
    @pulumi.getter(name="maxVersions")
    def max_versions(self) -> Optional[int]:
        """
        The number of versions to keep per key.
        """
        return pulumi.get(self, "max_versions")


