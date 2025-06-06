# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
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

__all__ = [
    'SecretBackendRoleVhost',
    'SecretBackendRoleVhostTopic',
    'SecretBackendRoleVhostTopicVhost',
]

@pulumi.output_type
class SecretBackendRoleVhost(dict):
    def __init__(__self__, *,
                 configure: builtins.str,
                 host: builtins.str,
                 read: builtins.str,
                 write: builtins.str):
        """
        :param builtins.str configure: The configure permissions for this vhost.
        :param builtins.str host: The vhost to set permissions for.
        :param builtins.str read: The read permissions for this vhost.
        :param builtins.str write: The write permissions for this vhost.
        """
        pulumi.set(__self__, "configure", configure)
        pulumi.set(__self__, "host", host)
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def configure(self) -> builtins.str:
        """
        The configure permissions for this vhost.
        """
        return pulumi.get(self, "configure")

    @property
    @pulumi.getter
    def host(self) -> builtins.str:
        """
        The vhost to set permissions for.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def read(self) -> builtins.str:
        """
        The read permissions for this vhost.
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter
    def write(self) -> builtins.str:
        """
        The write permissions for this vhost.
        """
        return pulumi.get(self, "write")


@pulumi.output_type
class SecretBackendRoleVhostTopic(dict):
    def __init__(__self__, *,
                 host: builtins.str,
                 vhosts: Optional[Sequence['outputs.SecretBackendRoleVhostTopicVhost']] = None):
        """
        :param builtins.str host: The vhost to set permissions for.
        :param Sequence['SecretBackendRoleVhostTopicVhostArgs'] vhosts: Specifies a map of virtual hosts to permissions.
        """
        pulumi.set(__self__, "host", host)
        if vhosts is not None:
            pulumi.set(__self__, "vhosts", vhosts)

    @property
    @pulumi.getter
    def host(self) -> builtins.str:
        """
        The vhost to set permissions for.
        """
        return pulumi.get(self, "host")

    @property
    @pulumi.getter
    def vhosts(self) -> Optional[Sequence['outputs.SecretBackendRoleVhostTopicVhost']]:
        """
        Specifies a map of virtual hosts to permissions.
        """
        return pulumi.get(self, "vhosts")


@pulumi.output_type
class SecretBackendRoleVhostTopicVhost(dict):
    def __init__(__self__, *,
                 read: builtins.str,
                 topic: builtins.str,
                 write: builtins.str):
        """
        :param builtins.str read: The read permissions for this vhost.
        :param builtins.str topic: The vhost to set permissions for.
        :param builtins.str write: The write permissions for this vhost.
        """
        pulumi.set(__self__, "read", read)
        pulumi.set(__self__, "topic", topic)
        pulumi.set(__self__, "write", write)

    @property
    @pulumi.getter
    def read(self) -> builtins.str:
        """
        The read permissions for this vhost.
        """
        return pulumi.get(self, "read")

    @property
    @pulumi.getter
    def topic(self) -> builtins.str:
        """
        The vhost to set permissions for.
        """
        return pulumi.get(self, "topic")

    @property
    @pulumi.getter
    def write(self) -> builtins.str:
        """
        The write permissions for this vhost.
        """
        return pulumi.get(self, "write")


