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
    'GetAccessCredentialsResult',
    'AwaitableGetAccessCredentialsResult',
    'get_access_credentials',
    'get_access_credentials_output',
]

@pulumi.output_type
class GetAccessCredentialsResult:
    """
    A collection of values returned by getAccessCredentials.
    """
    def __init__(__self__, access_key=None, backend=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, lease_start_time=None, namespace=None, region=None, role=None, role_arn=None, secret_key=None, security_token=None, ttl=None, type=None):
        if access_key and not isinstance(access_key, str):
            raise TypeError("Expected argument 'access_key' to be a str")
        pulumi.set(__self__, "access_key", access_key)
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if lease_duration and not isinstance(lease_duration, int):
            raise TypeError("Expected argument 'lease_duration' to be a int")
        pulumi.set(__self__, "lease_duration", lease_duration)
        if lease_id and not isinstance(lease_id, str):
            raise TypeError("Expected argument 'lease_id' to be a str")
        pulumi.set(__self__, "lease_id", lease_id)
        if lease_renewable and not isinstance(lease_renewable, bool):
            raise TypeError("Expected argument 'lease_renewable' to be a bool")
        pulumi.set(__self__, "lease_renewable", lease_renewable)
        if lease_start_time and not isinstance(lease_start_time, str):
            raise TypeError("Expected argument 'lease_start_time' to be a str")
        pulumi.set(__self__, "lease_start_time", lease_start_time)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if role_arn and not isinstance(role_arn, str):
            raise TypeError("Expected argument 'role_arn' to be a str")
        pulumi.set(__self__, "role_arn", role_arn)
        if secret_key and not isinstance(secret_key, str):
            raise TypeError("Expected argument 'secret_key' to be a str")
        pulumi.set(__self__, "secret_key", secret_key)
        if security_token and not isinstance(security_token, str):
            raise TypeError("Expected argument 'security_token' to be a str")
        pulumi.set(__self__, "security_token", security_token)
        if ttl and not isinstance(ttl, str):
            raise TypeError("Expected argument 'ttl' to be a str")
        pulumi.set(__self__, "ttl", ttl)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> builtins.str:
        """
        The AWS Access Key ID returned by Vault.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def backend(self) -> builtins.str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> builtins.int:
        """
        The duration of the secret lease, in seconds relative
        to the time the data was requested. Once this time has passed any plan
        generated with this data may fail to apply.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> builtins.str:
        """
        The lease identifier assigned by Vault.
        """
        return pulumi.get(self, "lease_id")

    @property
    @pulumi.getter(name="leaseRenewable")
    def lease_renewable(self) -> builtins.bool:
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter(name="leaseStartTime")
    def lease_start_time(self) -> builtins.str:
        return pulumi.get(self, "lease_start_time")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def region(self) -> Optional[builtins.str]:
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def role(self) -> builtins.str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="roleArn")
    def role_arn(self) -> Optional[builtins.str]:
        return pulumi.get(self, "role_arn")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> builtins.str:
        """
        The AWS Secret Key returned by Vault.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="securityToken")
    def security_token(self) -> builtins.str:
        """
        The STS token returned by Vault, if any.
        """
        return pulumi.get(self, "security_token")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[builtins.str]:
        return pulumi.get(self, "ttl")

    @property
    @pulumi.getter
    def type(self) -> Optional[builtins.str]:
        return pulumi.get(self, "type")


class AwaitableGetAccessCredentialsResult(GetAccessCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessCredentialsResult(
            access_key=self.access_key,
            backend=self.backend,
            id=self.id,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            lease_start_time=self.lease_start_time,
            namespace=self.namespace,
            region=self.region,
            role=self.role,
            role_arn=self.role_arn,
            secret_key=self.secret_key,
            security_token=self.security_token,
            ttl=self.ttl,
            type=self.type)


def get_access_credentials(backend: Optional[builtins.str] = None,
                           namespace: Optional[builtins.str] = None,
                           region: Optional[builtins.str] = None,
                           role: Optional[builtins.str] = None,
                           role_arn: Optional[builtins.str] = None,
                           ttl: Optional[builtins.str] = None,
                           type: Optional[builtins.str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessCredentialsResult:
    """
    ## Example Usage


    :param builtins.str backend: The path to the AWS secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param builtins.str region: The region the read credentials belong to.
    :param builtins.str role: The name of the AWS secret backend role to read
           credentials from, with no leading or trailing `/`s.
    :param builtins.str role_arn: The specific AWS ARN to use
           from the configured role. If the role does not have multiple ARNs, this does
           not need to be specified.
    :param builtins.str ttl: Specifies the TTL for the use of the STS token. This
           is specified as a string with a duration suffix. Valid only when
           `credential_type` of the connected `aws.SecretBackendRole` resource is `assumed_role` or `federation_token`
    :param builtins.str type: The type of credentials to read. Defaults
           to `"creds"`, which just returns an AWS Access Key ID and Secret
           Key. Can also be set to `"sts"`, which will return a security token
           in addition to the keys.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['region'] = region
    __args__['role'] = role
    __args__['roleArn'] = role_arn
    __args__['ttl'] = ttl
    __args__['type'] = type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:aws/getAccessCredentials:getAccessCredentials', __args__, opts=opts, typ=GetAccessCredentialsResult).value

    return AwaitableGetAccessCredentialsResult(
        access_key=pulumi.get(__ret__, 'access_key'),
        backend=pulumi.get(__ret__, 'backend'),
        id=pulumi.get(__ret__, 'id'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_id=pulumi.get(__ret__, 'lease_id'),
        lease_renewable=pulumi.get(__ret__, 'lease_renewable'),
        lease_start_time=pulumi.get(__ret__, 'lease_start_time'),
        namespace=pulumi.get(__ret__, 'namespace'),
        region=pulumi.get(__ret__, 'region'),
        role=pulumi.get(__ret__, 'role'),
        role_arn=pulumi.get(__ret__, 'role_arn'),
        secret_key=pulumi.get(__ret__, 'secret_key'),
        security_token=pulumi.get(__ret__, 'security_token'),
        ttl=pulumi.get(__ret__, 'ttl'),
        type=pulumi.get(__ret__, 'type'))
def get_access_credentials_output(backend: Optional[pulumi.Input[builtins.str]] = None,
                                  namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  region: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  role: Optional[pulumi.Input[builtins.str]] = None,
                                  role_arn: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  ttl: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  type: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                  opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAccessCredentialsResult]:
    """
    ## Example Usage


    :param builtins.str backend: The path to the AWS secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param builtins.str region: The region the read credentials belong to.
    :param builtins.str role: The name of the AWS secret backend role to read
           credentials from, with no leading or trailing `/`s.
    :param builtins.str role_arn: The specific AWS ARN to use
           from the configured role. If the role does not have multiple ARNs, this does
           not need to be specified.
    :param builtins.str ttl: Specifies the TTL for the use of the STS token. This
           is specified as a string with a duration suffix. Valid only when
           `credential_type` of the connected `aws.SecretBackendRole` resource is `assumed_role` or `federation_token`
    :param builtins.str type: The type of credentials to read. Defaults
           to `"creds"`, which just returns an AWS Access Key ID and Secret
           Key. Can also be set to `"sts"`, which will return a security token
           in addition to the keys.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    __args__['region'] = region
    __args__['role'] = role
    __args__['roleArn'] = role_arn
    __args__['ttl'] = ttl
    __args__['type'] = type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:aws/getAccessCredentials:getAccessCredentials', __args__, opts=opts, typ=GetAccessCredentialsResult)
    return __ret__.apply(lambda __response__: GetAccessCredentialsResult(
        access_key=pulumi.get(__response__, 'access_key'),
        backend=pulumi.get(__response__, 'backend'),
        id=pulumi.get(__response__, 'id'),
        lease_duration=pulumi.get(__response__, 'lease_duration'),
        lease_id=pulumi.get(__response__, 'lease_id'),
        lease_renewable=pulumi.get(__response__, 'lease_renewable'),
        lease_start_time=pulumi.get(__response__, 'lease_start_time'),
        namespace=pulumi.get(__response__, 'namespace'),
        region=pulumi.get(__response__, 'region'),
        role=pulumi.get(__response__, 'role'),
        role_arn=pulumi.get(__response__, 'role_arn'),
        secret_key=pulumi.get(__response__, 'secret_key'),
        security_token=pulumi.get(__response__, 'security_token'),
        ttl=pulumi.get(__response__, 'ttl'),
        type=pulumi.get(__response__, 'type')))
