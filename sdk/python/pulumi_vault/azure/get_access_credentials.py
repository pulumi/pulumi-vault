# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
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
    def __init__(__self__, backend=None, client_id=None, client_secret=None, environment=None, id=None, lease_duration=None, lease_id=None, lease_renewable=None, lease_start_time=None, max_cred_validation_seconds=None, namespace=None, num_seconds_between_tests=None, num_sequential_successes=None, role=None, subscription_id=None, tenant_id=None, validate_creds=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if client_id and not isinstance(client_id, str):
            raise TypeError("Expected argument 'client_id' to be a str")
        pulumi.set(__self__, "client_id", client_id)
        if client_secret and not isinstance(client_secret, str):
            raise TypeError("Expected argument 'client_secret' to be a str")
        pulumi.set(__self__, "client_secret", client_secret)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
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
        if max_cred_validation_seconds and not isinstance(max_cred_validation_seconds, int):
            raise TypeError("Expected argument 'max_cred_validation_seconds' to be a int")
        pulumi.set(__self__, "max_cred_validation_seconds", max_cred_validation_seconds)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if num_seconds_between_tests and not isinstance(num_seconds_between_tests, int):
            raise TypeError("Expected argument 'num_seconds_between_tests' to be a int")
        pulumi.set(__self__, "num_seconds_between_tests", num_seconds_between_tests)
        if num_sequential_successes and not isinstance(num_sequential_successes, int):
            raise TypeError("Expected argument 'num_sequential_successes' to be a int")
        pulumi.set(__self__, "num_sequential_successes", num_sequential_successes)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if subscription_id and not isinstance(subscription_id, str):
            raise TypeError("Expected argument 'subscription_id' to be a str")
        pulumi.set(__self__, "subscription_id", subscription_id)
        if tenant_id and not isinstance(tenant_id, str):
            raise TypeError("Expected argument 'tenant_id' to be a str")
        pulumi.set(__self__, "tenant_id", tenant_id)
        if validate_creds and not isinstance(validate_creds, bool):
            raise TypeError("Expected argument 'validate_creds' to be a bool")
        pulumi.set(__self__, "validate_creds", validate_creds)

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> str:
        """
        The client id for credentials to query the Azure APIs.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> str:
        """
        The client secret for credentials to query the Azure APIs.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def environment(self) -> Optional[str]:
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> int:
        """
        The duration of the secret lease, in seconds relative
        to the time the data was requested. Once this time has passed any plan
        generated with this data may fail to apply.
        """
        return pulumi.get(self, "lease_duration")

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> str:
        """
        The lease identifier assigned by Vault.
        """
        return pulumi.get(self, "lease_id")

    @property
    @pulumi.getter(name="leaseRenewable")
    def lease_renewable(self) -> bool:
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter(name="leaseStartTime")
    def lease_start_time(self) -> str:
        return pulumi.get(self, "lease_start_time")

    @property
    @pulumi.getter(name="maxCredValidationSeconds")
    def max_cred_validation_seconds(self) -> Optional[int]:
        return pulumi.get(self, "max_cred_validation_seconds")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="numSecondsBetweenTests")
    def num_seconds_between_tests(self) -> Optional[int]:
        return pulumi.get(self, "num_seconds_between_tests")

    @property
    @pulumi.getter(name="numSequentialSuccesses")
    def num_sequential_successes(self) -> Optional[int]:
        return pulumi.get(self, "num_sequential_successes")

    @property
    @pulumi.getter
    def role(self) -> str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="subscriptionId")
    def subscription_id(self) -> Optional[str]:
        return pulumi.get(self, "subscription_id")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[str]:
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter(name="validateCreds")
    def validate_creds(self) -> Optional[bool]:
        return pulumi.get(self, "validate_creds")


class AwaitableGetAccessCredentialsResult(GetAccessCredentialsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAccessCredentialsResult(
            backend=self.backend,
            client_id=self.client_id,
            client_secret=self.client_secret,
            environment=self.environment,
            id=self.id,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            lease_start_time=self.lease_start_time,
            max_cred_validation_seconds=self.max_cred_validation_seconds,
            namespace=self.namespace,
            num_seconds_between_tests=self.num_seconds_between_tests,
            num_sequential_successes=self.num_sequential_successes,
            role=self.role,
            subscription_id=self.subscription_id,
            tenant_id=self.tenant_id,
            validate_creds=self.validate_creds)


def get_access_credentials(backend: Optional[str] = None,
                           environment: Optional[str] = None,
                           max_cred_validation_seconds: Optional[int] = None,
                           namespace: Optional[str] = None,
                           num_seconds_between_tests: Optional[int] = None,
                           num_sequential_successes: Optional[int] = None,
                           role: Optional[str] = None,
                           subscription_id: Optional[str] = None,
                           tenant_id: Optional[str] = None,
                           validate_creds: Optional[bool] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAccessCredentialsResult:
    """
    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_vault as vault

    creds = vault.azure.get_access_credentials(role="my-role",
        validate_creds=True,
        num_sequential_successes=8,
        num_seconds_between_tests=1,
        max_cred_validation_seconds=300)
    ```
    <!--End PulumiCodeChooser -->

    ## Caveats

    The `validate_creds` option requires read-access to the `backend` config endpoint.
    If the effective Vault role does not have the required permissions then valid values
    are required to be set for: `subscription_id`, `tenant_id`, `environment`.


    :param str backend: The path to the Azure secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param str environment: The Azure environment to use during credential validation.
           Defaults to the environment configured in the Vault backend.
           Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
           *See the caveats section for more information on this field.*
    :param int max_cred_validation_seconds: If 'validate_creds' is true, 
           the number of seconds after which to give up validating credentials. Defaults
           to 300.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param int num_seconds_between_tests: If 'validate_creds' is true, 
           the number of seconds to wait between each test of generated credentials.
           Defaults to 1.
    :param int num_sequential_successes: If 'validate_creds' is true, 
           the number of sequential successes required to validate generated
           credentials. Defaults to 8.
    :param str role: The name of the Azure secret backend role to read
           credentials from, with no leading or trailing `/`s.
    :param str subscription_id: The subscription ID to use during credential
           validation. Defaults to the subscription ID configured in the Vault `backend`.
           *See the caveats section for more information on this field.*
    :param str tenant_id: The tenant ID to use during credential validation.
           Defaults to the tenant ID configured in the Vault `backend`.
           *See the caveats section for more information on this field.*
    :param bool validate_creds: Whether generated credentials should be 
           validated before being returned. Defaults to `false`, which returns
           credentials without checking whether they have fully propagated throughout
           Azure Active Directory. Designating `true` activates testing.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['environment'] = environment
    __args__['maxCredValidationSeconds'] = max_cred_validation_seconds
    __args__['namespace'] = namespace
    __args__['numSecondsBetweenTests'] = num_seconds_between_tests
    __args__['numSequentialSuccesses'] = num_sequential_successes
    __args__['role'] = role
    __args__['subscriptionId'] = subscription_id
    __args__['tenantId'] = tenant_id
    __args__['validateCreds'] = validate_creds
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:azure/getAccessCredentials:getAccessCredentials', __args__, opts=opts, typ=GetAccessCredentialsResult).value

    return AwaitableGetAccessCredentialsResult(
        backend=pulumi.get(__ret__, 'backend'),
        client_id=pulumi.get(__ret__, 'client_id'),
        client_secret=pulumi.get(__ret__, 'client_secret'),
        environment=pulumi.get(__ret__, 'environment'),
        id=pulumi.get(__ret__, 'id'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_id=pulumi.get(__ret__, 'lease_id'),
        lease_renewable=pulumi.get(__ret__, 'lease_renewable'),
        lease_start_time=pulumi.get(__ret__, 'lease_start_time'),
        max_cred_validation_seconds=pulumi.get(__ret__, 'max_cred_validation_seconds'),
        namespace=pulumi.get(__ret__, 'namespace'),
        num_seconds_between_tests=pulumi.get(__ret__, 'num_seconds_between_tests'),
        num_sequential_successes=pulumi.get(__ret__, 'num_sequential_successes'),
        role=pulumi.get(__ret__, 'role'),
        subscription_id=pulumi.get(__ret__, 'subscription_id'),
        tenant_id=pulumi.get(__ret__, 'tenant_id'),
        validate_creds=pulumi.get(__ret__, 'validate_creds'))


@_utilities.lift_output_func(get_access_credentials)
def get_access_credentials_output(backend: Optional[pulumi.Input[str]] = None,
                                  environment: Optional[pulumi.Input[Optional[str]]] = None,
                                  max_cred_validation_seconds: Optional[pulumi.Input[Optional[int]]] = None,
                                  namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                  num_seconds_between_tests: Optional[pulumi.Input[Optional[int]]] = None,
                                  num_sequential_successes: Optional[pulumi.Input[Optional[int]]] = None,
                                  role: Optional[pulumi.Input[str]] = None,
                                  subscription_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  tenant_id: Optional[pulumi.Input[Optional[str]]] = None,
                                  validate_creds: Optional[pulumi.Input[Optional[bool]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAccessCredentialsResult]:
    """
    ## Example Usage

    <!--Start PulumiCodeChooser -->
    ```python
    import pulumi
    import pulumi_vault as vault

    creds = vault.azure.get_access_credentials(role="my-role",
        validate_creds=True,
        num_sequential_successes=8,
        num_seconds_between_tests=1,
        max_cred_validation_seconds=300)
    ```
    <!--End PulumiCodeChooser -->

    ## Caveats

    The `validate_creds` option requires read-access to the `backend` config endpoint.
    If the effective Vault role does not have the required permissions then valid values
    are required to be set for: `subscription_id`, `tenant_id`, `environment`.


    :param str backend: The path to the Azure secret backend to
           read credentials from, with no leading or trailing `/`s.
    :param str environment: The Azure environment to use during credential validation.
           Defaults to the environment configured in the Vault backend.
           Some possible values: `AzurePublicCloud`, `AzureGovernmentCloud`
           *See the caveats section for more information on this field.*
    :param int max_cred_validation_seconds: If 'validate_creds' is true, 
           the number of seconds after which to give up validating credentials. Defaults
           to 300.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param int num_seconds_between_tests: If 'validate_creds' is true, 
           the number of seconds to wait between each test of generated credentials.
           Defaults to 1.
    :param int num_sequential_successes: If 'validate_creds' is true, 
           the number of sequential successes required to validate generated
           credentials. Defaults to 8.
    :param str role: The name of the Azure secret backend role to read
           credentials from, with no leading or trailing `/`s.
    :param str subscription_id: The subscription ID to use during credential
           validation. Defaults to the subscription ID configured in the Vault `backend`.
           *See the caveats section for more information on this field.*
    :param str tenant_id: The tenant ID to use during credential validation.
           Defaults to the tenant ID configured in the Vault `backend`.
           *See the caveats section for more information on this field.*
    :param bool validate_creds: Whether generated credentials should be 
           validated before being returned. Defaults to `false`, which returns
           credentials without checking whether they have fully propagated throughout
           Azure Active Directory. Designating `true` activates testing.
    """
    ...
