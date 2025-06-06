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

__all__ = [
    'BackendConfigCmpv2AuthenticatorsArgs',
    'BackendConfigCmpv2AuthenticatorsArgsDict',
    'BackendConfigEstAuthenticatorsArgs',
    'BackendConfigEstAuthenticatorsArgsDict',
    'SecretBackendRolePolicyIdentifierArgs',
    'SecretBackendRolePolicyIdentifierArgsDict',
]

MYPY = False

if not MYPY:
    class BackendConfigCmpv2AuthenticatorsArgsDict(TypedDict):
        cert: NotRequired[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]
        """
        "The accessor (required) and cert_role (optional) properties for cert auth backends".
        """
elif False:
    BackendConfigCmpv2AuthenticatorsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BackendConfigCmpv2AuthenticatorsArgs:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] cert: "The accessor (required) and cert_role (optional) properties for cert auth backends".
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        "The accessor (required) and cert_role (optional) properties for cert auth backends".
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "cert", value)


if not MYPY:
    class BackendConfigEstAuthenticatorsArgsDict(TypedDict):
        cert: NotRequired[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]
        """
        "The accessor (required) and cert_role (optional) properties for cert auth backends".
        """
        userpass: NotRequired[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]
        """
        "The accessor (required) property for user pass auth backends".
        """
elif False:
    BackendConfigEstAuthenticatorsArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class BackendConfigEstAuthenticatorsArgs:
    def __init__(__self__, *,
                 cert: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None,
                 userpass: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]] = None):
        """
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] cert: "The accessor (required) and cert_role (optional) properties for cert auth backends".
        :param pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]] userpass: "The accessor (required) property for user pass auth backends".
        """
        if cert is not None:
            pulumi.set(__self__, "cert", cert)
        if userpass is not None:
            pulumi.set(__self__, "userpass", userpass)

    @property
    @pulumi.getter
    def cert(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        "The accessor (required) and cert_role (optional) properties for cert auth backends".
        """
        return pulumi.get(self, "cert")

    @cert.setter
    def cert(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "cert", value)

    @property
    @pulumi.getter
    def userpass(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]:
        """
        "The accessor (required) property for user pass auth backends".
        """
        return pulumi.get(self, "userpass")

    @userpass.setter
    def userpass(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "userpass", value)


if not MYPY:
    class SecretBackendRolePolicyIdentifierArgsDict(TypedDict):
        oid: pulumi.Input[builtins.str]
        """
        The OID for the policy identifier
        """
        cps: NotRequired[pulumi.Input[builtins.str]]
        """
        The URL of the CPS for the policy identifier
        """
        notice: NotRequired[pulumi.Input[builtins.str]]
        """
        A notice for the policy identifier
        """
elif False:
    SecretBackendRolePolicyIdentifierArgsDict: TypeAlias = Mapping[str, Any]

@pulumi.input_type
class SecretBackendRolePolicyIdentifierArgs:
    def __init__(__self__, *,
                 oid: pulumi.Input[builtins.str],
                 cps: Optional[pulumi.Input[builtins.str]] = None,
                 notice: Optional[pulumi.Input[builtins.str]] = None):
        """
        :param pulumi.Input[builtins.str] oid: The OID for the policy identifier
        :param pulumi.Input[builtins.str] cps: The URL of the CPS for the policy identifier
        :param pulumi.Input[builtins.str] notice: A notice for the policy identifier
        """
        pulumi.set(__self__, "oid", oid)
        if cps is not None:
            pulumi.set(__self__, "cps", cps)
        if notice is not None:
            pulumi.set(__self__, "notice", notice)

    @property
    @pulumi.getter
    def oid(self) -> pulumi.Input[builtins.str]:
        """
        The OID for the policy identifier
        """
        return pulumi.get(self, "oid")

    @oid.setter
    def oid(self, value: pulumi.Input[builtins.str]):
        pulumi.set(self, "oid", value)

    @property
    @pulumi.getter
    def cps(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The URL of the CPS for the policy identifier
        """
        return pulumi.get(self, "cps")

    @cps.setter
    def cps(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "cps", value)

    @property
    @pulumi.getter
    def notice(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        A notice for the policy identifier
        """
        return pulumi.get(self, "notice")

    @notice.setter
    def notice(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "notice", value)


