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
    'GetAuthBackendConfigResult',
    'AwaitableGetAuthBackendConfigResult',
    'get_auth_backend_config',
    'get_auth_backend_config_output',
]

@pulumi.output_type
class GetAuthBackendConfigResult:
    """
    A collection of values returned by getAuthBackendConfig.
    """
    def __init__(__self__, backend=None, disable_iss_validation=None, disable_local_ca_jwt=None, id=None, issuer=None, kubernetes_ca_cert=None, kubernetes_host=None, namespace=None, pem_keys=None, use_annotations_as_alias_metadata=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if disable_iss_validation and not isinstance(disable_iss_validation, bool):
            raise TypeError("Expected argument 'disable_iss_validation' to be a bool")
        pulumi.set(__self__, "disable_iss_validation", disable_iss_validation)
        if disable_local_ca_jwt and not isinstance(disable_local_ca_jwt, bool):
            raise TypeError("Expected argument 'disable_local_ca_jwt' to be a bool")
        pulumi.set(__self__, "disable_local_ca_jwt", disable_local_ca_jwt)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if issuer and not isinstance(issuer, str):
            raise TypeError("Expected argument 'issuer' to be a str")
        pulumi.set(__self__, "issuer", issuer)
        if kubernetes_ca_cert and not isinstance(kubernetes_ca_cert, str):
            raise TypeError("Expected argument 'kubernetes_ca_cert' to be a str")
        pulumi.set(__self__, "kubernetes_ca_cert", kubernetes_ca_cert)
        if kubernetes_host and not isinstance(kubernetes_host, str):
            raise TypeError("Expected argument 'kubernetes_host' to be a str")
        pulumi.set(__self__, "kubernetes_host", kubernetes_host)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if pem_keys and not isinstance(pem_keys, list):
            raise TypeError("Expected argument 'pem_keys' to be a list")
        pulumi.set(__self__, "pem_keys", pem_keys)
        if use_annotations_as_alias_metadata and not isinstance(use_annotations_as_alias_metadata, bool):
            raise TypeError("Expected argument 'use_annotations_as_alias_metadata' to be a bool")
        pulumi.set(__self__, "use_annotations_as_alias_metadata", use_annotations_as_alias_metadata)

    @property
    @pulumi.getter
    def backend(self) -> Optional[builtins.str]:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="disableIssValidation")
    def disable_iss_validation(self) -> builtins.bool:
        """
        (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_iss_validation")

    @property
    @pulumi.getter(name="disableLocalCaJwt")
    def disable_local_ca_jwt(self) -> builtins.bool:
        """
        (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_local_ca_jwt")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def issuer(self) -> builtins.str:
        """
        Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> builtins.str:
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @property
    @pulumi.getter(name="kubernetesHost")
    def kubernetes_host(self) -> builtins.str:
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        return pulumi.get(self, "kubernetes_host")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="pemKeys")
    def pem_keys(self) -> Sequence[builtins.str]:
        """
        Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        return pulumi.get(self, "pem_keys")

    @property
    @pulumi.getter(name="useAnnotationsAsAliasMetadata")
    def use_annotations_as_alias_metadata(self) -> builtins.bool:
        """
        (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
        """
        return pulumi.get(self, "use_annotations_as_alias_metadata")


class AwaitableGetAuthBackendConfigResult(GetAuthBackendConfigResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAuthBackendConfigResult(
            backend=self.backend,
            disable_iss_validation=self.disable_iss_validation,
            disable_local_ca_jwt=self.disable_local_ca_jwt,
            id=self.id,
            issuer=self.issuer,
            kubernetes_ca_cert=self.kubernetes_ca_cert,
            kubernetes_host=self.kubernetes_host,
            namespace=self.namespace,
            pem_keys=self.pem_keys,
            use_annotations_as_alias_metadata=self.use_annotations_as_alias_metadata)


def get_auth_backend_config(backend: Optional[builtins.str] = None,
                            disable_iss_validation: Optional[builtins.bool] = None,
                            disable_local_ca_jwt: Optional[builtins.bool] = None,
                            issuer: Optional[builtins.str] = None,
                            kubernetes_ca_cert: Optional[builtins.str] = None,
                            kubernetes_host: Optional[builtins.str] = None,
                            namespace: Optional[builtins.str] = None,
                            pem_keys: Optional[Sequence[builtins.str]] = None,
                            use_annotations_as_alias_metadata: Optional[builtins.bool] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthBackendConfigResult:
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
    information.


    :param builtins.str backend: The unique name for the Kubernetes backend the config to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param builtins.bool disable_iss_validation: (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
    :param builtins.bool disable_local_ca_jwt: (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
    :param builtins.str issuer: Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
    :param builtins.str kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
    :param builtins.str kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured namespace.
           *Available only for Vault Enterprise*.
    :param Sequence[builtins.str] pem_keys: Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
    :param builtins.bool use_annotations_as_alias_metadata: (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['disableIssValidation'] = disable_iss_validation
    __args__['disableLocalCaJwt'] = disable_local_ca_jwt
    __args__['issuer'] = issuer
    __args__['kubernetesCaCert'] = kubernetes_ca_cert
    __args__['kubernetesHost'] = kubernetes_host
    __args__['namespace'] = namespace
    __args__['pemKeys'] = pem_keys
    __args__['useAnnotationsAsAliasMetadata'] = use_annotations_as_alias_metadata
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig', __args__, opts=opts, typ=GetAuthBackendConfigResult).value

    return AwaitableGetAuthBackendConfigResult(
        backend=pulumi.get(__ret__, 'backend'),
        disable_iss_validation=pulumi.get(__ret__, 'disable_iss_validation'),
        disable_local_ca_jwt=pulumi.get(__ret__, 'disable_local_ca_jwt'),
        id=pulumi.get(__ret__, 'id'),
        issuer=pulumi.get(__ret__, 'issuer'),
        kubernetes_ca_cert=pulumi.get(__ret__, 'kubernetes_ca_cert'),
        kubernetes_host=pulumi.get(__ret__, 'kubernetes_host'),
        namespace=pulumi.get(__ret__, 'namespace'),
        pem_keys=pulumi.get(__ret__, 'pem_keys'),
        use_annotations_as_alias_metadata=pulumi.get(__ret__, 'use_annotations_as_alias_metadata'))
def get_auth_backend_config_output(backend: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   disable_iss_validation: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                   disable_local_ca_jwt: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                   issuer: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   kubernetes_ca_cert: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   kubernetes_host: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                                   pem_keys: Optional[pulumi.Input[Optional[Sequence[builtins.str]]]] = None,
                                   use_annotations_as_alias_metadata: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetAuthBackendConfigResult]:
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
    information.


    :param builtins.str backend: The unique name for the Kubernetes backend the config to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param builtins.bool disable_iss_validation: (Optional) Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
    :param builtins.bool disable_local_ca_jwt: (Optional) Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
    :param builtins.str issuer: Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
    :param builtins.str kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
    :param builtins.str kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
    :param builtins.str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured namespace.
           *Available only for Vault Enterprise*.
    :param Sequence[builtins.str] pem_keys: Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
    :param builtins.bool use_annotations_as_alias_metadata: (Optional) Use annotations from the client token's associated service account as alias metadata for the Vault entity. Requires Vault `v1.16+` or Vault auth kubernetes plugin `v0.18.0+`
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['disableIssValidation'] = disable_iss_validation
    __args__['disableLocalCaJwt'] = disable_local_ca_jwt
    __args__['issuer'] = issuer
    __args__['kubernetesCaCert'] = kubernetes_ca_cert
    __args__['kubernetesHost'] = kubernetes_host
    __args__['namespace'] = namespace
    __args__['pemKeys'] = pem_keys
    __args__['useAnnotationsAsAliasMetadata'] = use_annotations_as_alias_metadata
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig', __args__, opts=opts, typ=GetAuthBackendConfigResult)
    return __ret__.apply(lambda __response__: GetAuthBackendConfigResult(
        backend=pulumi.get(__response__, 'backend'),
        disable_iss_validation=pulumi.get(__response__, 'disable_iss_validation'),
        disable_local_ca_jwt=pulumi.get(__response__, 'disable_local_ca_jwt'),
        id=pulumi.get(__response__, 'id'),
        issuer=pulumi.get(__response__, 'issuer'),
        kubernetes_ca_cert=pulumi.get(__response__, 'kubernetes_ca_cert'),
        kubernetes_host=pulumi.get(__response__, 'kubernetes_host'),
        namespace=pulumi.get(__response__, 'namespace'),
        pem_keys=pulumi.get(__response__, 'pem_keys'),
        use_annotations_as_alias_metadata=pulumi.get(__response__, 'use_annotations_as_alias_metadata')))
