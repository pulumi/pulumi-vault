# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = [
    'GetAuthBackendConfigResult',
    'AwaitableGetAuthBackendConfigResult',
    'get_auth_backend_config',
]

@pulumi.output_type
class GetAuthBackendConfigResult:
    """
    A collection of values returned by getAuthBackendConfig.
    """
    def __init__(__self__, backend=None, disable_iss_validation=None, disable_local_ca_jwt=None, id=None, issuer=None, kubernetes_ca_cert=None, kubernetes_host=None, pem_keys=None):
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
        if pem_keys and not isinstance(pem_keys, list):
            raise TypeError("Expected argument 'pem_keys' to be a list")
        pulumi.set(__self__, "pem_keys", pem_keys)

    @property
    @pulumi.getter
    def backend(self) -> Optional[str]:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="disableIssValidation")
    def disable_iss_validation(self) -> bool:
        return pulumi.get(self, "disable_iss_validation")

    @property
    @pulumi.getter(name="disableLocalCaJwt")
    def disable_local_ca_jwt(self) -> bool:
        return pulumi.get(self, "disable_local_ca_jwt")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def issuer(self) -> str:
        """
        Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> str:
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @property
    @pulumi.getter(name="kubernetesHost")
    def kubernetes_host(self) -> str:
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        return pulumi.get(self, "kubernetes_host")

    @property
    @pulumi.getter(name="pemKeys")
    def pem_keys(self) -> Sequence[str]:
        """
        Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        return pulumi.get(self, "pem_keys")


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
            pem_keys=self.pem_keys)


def get_auth_backend_config(backend: Optional[str] = None,
                            disable_iss_validation: Optional[bool] = None,
                            disable_local_ca_jwt: Optional[bool] = None,
                            issuer: Optional[str] = None,
                            kubernetes_ca_cert: Optional[str] = None,
                            kubernetes_host: Optional[str] = None,
                            pem_keys: Optional[Sequence[str]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAuthBackendConfigResult:
    """
    Reads the Role of an Kubernetes from a Vault server. See the [Vault
    documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
    information.


    :param str backend: The unique name for the Kubernetes backend the config to
           retrieve Role attributes for resides in. Defaults to "kubernetes".
    :param str issuer: Optional JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
    :param str kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
    :param str kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
    :param Sequence[str] pem_keys: Optional list of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['disableIssValidation'] = disable_iss_validation
    __args__['disableLocalCaJwt'] = disable_local_ca_jwt
    __args__['issuer'] = issuer
    __args__['kubernetesCaCert'] = kubernetes_ca_cert
    __args__['kubernetesHost'] = kubernetes_host
    __args__['pemKeys'] = pem_keys
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig', __args__, opts=opts, typ=GetAuthBackendConfigResult).value

    return AwaitableGetAuthBackendConfigResult(
        backend=__ret__.backend,
        disable_iss_validation=__ret__.disable_iss_validation,
        disable_local_ca_jwt=__ret__.disable_local_ca_jwt,
        id=__ret__.id,
        issuer=__ret__.issuer,
        kubernetes_ca_cert=__ret__.kubernetes_ca_cert,
        kubernetes_host=__ret__.kubernetes_host,
        pem_keys=__ret__.pem_keys)
