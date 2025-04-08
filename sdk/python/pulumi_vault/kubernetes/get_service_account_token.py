# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
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
    'GetServiceAccountTokenResult',
    'AwaitableGetServiceAccountTokenResult',
    'get_service_account_token',
    'get_service_account_token_output',
]

@pulumi.output_type
class GetServiceAccountTokenResult:
    """
    A collection of values returned by getServiceAccountToken.
    """
    def __init__(__self__, backend=None, cluster_role_binding=None, id=None, kubernetes_namespace=None, lease_duration=None, lease_id=None, lease_renewable=None, namespace=None, role=None, service_account_name=None, service_account_namespace=None, service_account_token=None, ttl=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if cluster_role_binding and not isinstance(cluster_role_binding, bool):
            raise TypeError("Expected argument 'cluster_role_binding' to be a bool")
        pulumi.set(__self__, "cluster_role_binding", cluster_role_binding)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if kubernetes_namespace and not isinstance(kubernetes_namespace, str):
            raise TypeError("Expected argument 'kubernetes_namespace' to be a str")
        pulumi.set(__self__, "kubernetes_namespace", kubernetes_namespace)
        if lease_duration and not isinstance(lease_duration, int):
            raise TypeError("Expected argument 'lease_duration' to be a int")
        pulumi.set(__self__, "lease_duration", lease_duration)
        if lease_id and not isinstance(lease_id, str):
            raise TypeError("Expected argument 'lease_id' to be a str")
        pulumi.set(__self__, "lease_id", lease_id)
        if lease_renewable and not isinstance(lease_renewable, bool):
            raise TypeError("Expected argument 'lease_renewable' to be a bool")
        pulumi.set(__self__, "lease_renewable", lease_renewable)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if service_account_name and not isinstance(service_account_name, str):
            raise TypeError("Expected argument 'service_account_name' to be a str")
        pulumi.set(__self__, "service_account_name", service_account_name)
        if service_account_namespace and not isinstance(service_account_namespace, str):
            raise TypeError("Expected argument 'service_account_namespace' to be a str")
        pulumi.set(__self__, "service_account_namespace", service_account_namespace)
        if service_account_token and not isinstance(service_account_token, str):
            raise TypeError("Expected argument 'service_account_token' to be a str")
        pulumi.set(__self__, "service_account_token", service_account_token)
        if ttl and not isinstance(ttl, str):
            raise TypeError("Expected argument 'ttl' to be a str")
        pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter
    def backend(self) -> builtins.str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="clusterRoleBinding")
    def cluster_role_binding(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "cluster_role_binding")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="kubernetesNamespace")
    def kubernetes_namespace(self) -> builtins.str:
        return pulumi.get(self, "kubernetes_namespace")

    @property
    @pulumi.getter(name="leaseDuration")
    def lease_duration(self) -> builtins.int:
        """
        The duration of the lease in seconds.
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
        """
        True if the duration of this lease can be extended through renewal.
        """
        return pulumi.get(self, "lease_renewable")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def role(self) -> builtins.str:
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="serviceAccountName")
    def service_account_name(self) -> builtins.str:
        """
        The name of the service account associated with the token.
        """
        return pulumi.get(self, "service_account_name")

    @property
    @pulumi.getter(name="serviceAccountNamespace")
    def service_account_namespace(self) -> builtins.str:
        """
        The Kubernetes namespace that the service account resides in.
        """
        return pulumi.get(self, "service_account_namespace")

    @property
    @pulumi.getter(name="serviceAccountToken")
    def service_account_token(self) -> builtins.str:
        """
        The Kubernetes service account token.
        """
        return pulumi.get(self, "service_account_token")

    @property
    @pulumi.getter
    def ttl(self) -> Optional[builtins.str]:
        return pulumi.get(self, "ttl")


class AwaitableGetServiceAccountTokenResult(GetServiceAccountTokenResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceAccountTokenResult(
            backend=self.backend,
            cluster_role_binding=self.cluster_role_binding,
            id=self.id,
            kubernetes_namespace=self.kubernetes_namespace,
            lease_duration=self.lease_duration,
            lease_id=self.lease_id,
            lease_renewable=self.lease_renewable,
            namespace=self.namespace,
            role=self.role,
            service_account_name=self.service_account_name,
            service_account_namespace=self.service_account_namespace,
            service_account_token=self.service_account_token,
            ttl=self.ttl)


def get_service_account_token(backend: Optional[builtins.str] = None,
                              cluster_role_binding: Optional[builtins.bool] = None,
                              kubernetes_namespace: Optional[builtins.str] = None,
                              namespace: Optional[builtins.str] = None,
                              role: Optional[builtins.str] = None,
                              ttl: Optional[builtins.str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceAccountTokenResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_std as std
    import pulumi_vault as vault

    config = vault.kubernetes.SecretBackend("config",
        path="kubernetes",
        description="kubernetes secrets engine description",
        kubernetes_host="https://127.0.0.1:61233",
        kubernetes_ca_cert=std.file(input="/path/to/cert").result,
        service_account_jwt=std.file(input="/path/to/token").result,
        disable_local_ca_jwt=False)
    role = vault.kubernetes.SecretBackendRole("role",
        backend=config.path,
        name="service-account-name-role",
        allowed_kubernetes_namespaces=["*"],
        token_max_ttl=43200,
        token_default_ttl=21600,
        service_account_name="test-service-account-with-generated-token",
        extra_labels={
            "id": "abc123",
            "name": "some_name",
        },
        extra_annotations={
            "env": "development",
            "location": "earth",
        })
    token = vault.kubernetes.get_service_account_token_output(backend=config.path,
        role=role.name,
        kubernetes_namespace="test",
        cluster_role_binding=False,
        ttl="1h")
    ```


    :param builtins.str backend: The Kubernetes secret backend to generate service account 
           tokens from.
    :param builtins.bool cluster_role_binding: If true, generate a ClusterRoleBinding to grant 
           permissions across the whole cluster instead of within a namespace.
    :param builtins.str kubernetes_namespace: The name of the Kubernetes namespace in which to 
           generate the credentials.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param builtins.str role: The name of the Kubernetes secret backend role to generate service 
           account tokens from.
    :param builtins.str ttl: The TTL of the generated Kubernetes service account token, specified in 
           seconds or as a Go duration format string.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['clusterRoleBinding'] = cluster_role_binding
    __args__['kubernetesNamespace'] = kubernetes_namespace
    __args__['namespace'] = namespace
    __args__['role'] = role
    __args__['ttl'] = ttl
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getServiceAccountToken:getServiceAccountToken', __args__, opts=opts, typ=GetServiceAccountTokenResult).value

    return AwaitableGetServiceAccountTokenResult(
        backend=pulumi.get(__ret__, 'backend'),
        cluster_role_binding=pulumi.get(__ret__, 'cluster_role_binding'),
        id=pulumi.get(__ret__, 'id'),
        kubernetes_namespace=pulumi.get(__ret__, 'kubernetes_namespace'),
        lease_duration=pulumi.get(__ret__, 'lease_duration'),
        lease_id=pulumi.get(__ret__, 'lease_id'),
        lease_renewable=pulumi.get(__ret__, 'lease_renewable'),
        namespace=pulumi.get(__ret__, 'namespace'),
        role=pulumi.get(__ret__, 'role'),
        service_account_name=pulumi.get(__ret__, 'service_account_name'),
        service_account_namespace=pulumi.get(__ret__, 'service_account_namespace'),
        service_account_token=pulumi.get(__ret__, 'service_account_token'),
        ttl=pulumi.get(__ret__, 'ttl'))
def get_service_account_token_output(backend: Optional[pulumi.Input[builtins.str]] = None,
                                     cluster_role_binding: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                     kubernetes_namespace: Optional[pulumi.Input[builtins.str]] = None,
                                     namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     role: Optional[pulumi.Input[builtins.str]] = None,
                                     ttl: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetServiceAccountTokenResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_std as std
    import pulumi_vault as vault

    config = vault.kubernetes.SecretBackend("config",
        path="kubernetes",
        description="kubernetes secrets engine description",
        kubernetes_host="https://127.0.0.1:61233",
        kubernetes_ca_cert=std.file(input="/path/to/cert").result,
        service_account_jwt=std.file(input="/path/to/token").result,
        disable_local_ca_jwt=False)
    role = vault.kubernetes.SecretBackendRole("role",
        backend=config.path,
        name="service-account-name-role",
        allowed_kubernetes_namespaces=["*"],
        token_max_ttl=43200,
        token_default_ttl=21600,
        service_account_name="test-service-account-with-generated-token",
        extra_labels={
            "id": "abc123",
            "name": "some_name",
        },
        extra_annotations={
            "env": "development",
            "location": "earth",
        })
    token = vault.kubernetes.get_service_account_token_output(backend=config.path,
        role=role.name,
        kubernetes_namespace="test",
        cluster_role_binding=False,
        ttl="1h")
    ```


    :param builtins.str backend: The Kubernetes secret backend to generate service account 
           tokens from.
    :param builtins.bool cluster_role_binding: If true, generate a ClusterRoleBinding to grant 
           permissions across the whole cluster instead of within a namespace.
    :param builtins.str kubernetes_namespace: The name of the Kubernetes namespace in which to 
           generate the credentials.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    :param builtins.str role: The name of the Kubernetes secret backend role to generate service 
           account tokens from.
    :param builtins.str ttl: The TTL of the generated Kubernetes service account token, specified in 
           seconds or as a Go duration format string.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['clusterRoleBinding'] = cluster_role_binding
    __args__['kubernetesNamespace'] = kubernetes_namespace
    __args__['namespace'] = namespace
    __args__['role'] = role
    __args__['ttl'] = ttl
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:kubernetes/getServiceAccountToken:getServiceAccountToken', __args__, opts=opts, typ=GetServiceAccountTokenResult)
    return __ret__.apply(lambda __response__: GetServiceAccountTokenResult(
        backend=pulumi.get(__response__, 'backend'),
        cluster_role_binding=pulumi.get(__response__, 'cluster_role_binding'),
        id=pulumi.get(__response__, 'id'),
        kubernetes_namespace=pulumi.get(__response__, 'kubernetes_namespace'),
        lease_duration=pulumi.get(__response__, 'lease_duration'),
        lease_id=pulumi.get(__response__, 'lease_id'),
        lease_renewable=pulumi.get(__response__, 'lease_renewable'),
        namespace=pulumi.get(__response__, 'namespace'),
        role=pulumi.get(__response__, 'role'),
        service_account_name=pulumi.get(__response__, 'service_account_name'),
        service_account_namespace=pulumi.get(__response__, 'service_account_namespace'),
        service_account_token=pulumi.get(__response__, 'service_account_token'),
        ttl=pulumi.get(__response__, 'ttl')))
