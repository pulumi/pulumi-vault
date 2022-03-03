# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthBackendConfigArgs', 'AuthBackendConfig']

@pulumi.input_type
class AuthBackendConfigArgs:
    def __init__(__self__, *,
                 kubernetes_host: pulumi.Input[str],
                 backend: Optional[pulumi.Input[str]] = None,
                 disable_iss_validation: Optional[pulumi.Input[bool]] = None,
                 disable_local_ca_jwt: Optional[pulumi.Input[bool]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 pem_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_reviewer_jwt: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AuthBackendConfig resource.
        :param pulumi.Input[str] kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        :param pulumi.Input[str] backend: Unique name of the kubernetes backend to configure.
        :param pulumi.Input[bool] disable_iss_validation: Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[bool] disable_local_ca_jwt: Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[str] issuer: JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        :param pulumi.Input[str] kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] pem_keys: List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        :param pulumi.Input[str] token_reviewer_jwt: A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        pulumi.set(__self__, "kubernetes_host", kubernetes_host)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if disable_iss_validation is not None:
            pulumi.set(__self__, "disable_iss_validation", disable_iss_validation)
        if disable_local_ca_jwt is not None:
            pulumi.set(__self__, "disable_local_ca_jwt", disable_local_ca_jwt)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if kubernetes_ca_cert is not None:
            pulumi.set(__self__, "kubernetes_ca_cert", kubernetes_ca_cert)
        if pem_keys is not None:
            pulumi.set(__self__, "pem_keys", pem_keys)
        if token_reviewer_jwt is not None:
            pulumi.set(__self__, "token_reviewer_jwt", token_reviewer_jwt)

    @property
    @pulumi.getter(name="kubernetesHost")
    def kubernetes_host(self) -> pulumi.Input[str]:
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        return pulumi.get(self, "kubernetes_host")

    @kubernetes_host.setter
    def kubernetes_host(self, value: pulumi.Input[str]):
        pulumi.set(self, "kubernetes_host", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the kubernetes backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="disableIssValidation")
    def disable_iss_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_iss_validation")

    @disable_iss_validation.setter
    def disable_iss_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_iss_validation", value)

    @property
    @pulumi.getter(name="disableLocalCaJwt")
    def disable_local_ca_jwt(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_local_ca_jwt")

    @disable_local_ca_jwt.setter
    def disable_local_ca_jwt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_local_ca_jwt", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @kubernetes_ca_cert.setter
    def kubernetes_ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_ca_cert", value)

    @property
    @pulumi.getter(name="pemKeys")
    def pem_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        return pulumi.get(self, "pem_keys")

    @pem_keys.setter
    def pem_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "pem_keys", value)

    @property
    @pulumi.getter(name="tokenReviewerJwt")
    def token_reviewer_jwt(self) -> Optional[pulumi.Input[str]]:
        """
        A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        return pulumi.get(self, "token_reviewer_jwt")

    @token_reviewer_jwt.setter
    def token_reviewer_jwt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_reviewer_jwt", value)


@pulumi.input_type
class _AuthBackendConfigState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 disable_iss_validation: Optional[pulumi.Input[bool]] = None,
                 disable_local_ca_jwt: Optional[pulumi.Input[bool]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_host: Optional[pulumi.Input[str]] = None,
                 pem_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_reviewer_jwt: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering AuthBackendConfig resources.
        :param pulumi.Input[str] backend: Unique name of the kubernetes backend to configure.
        :param pulumi.Input[bool] disable_iss_validation: Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[bool] disable_local_ca_jwt: Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[str] issuer: JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        :param pulumi.Input[str] kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        :param pulumi.Input[str] kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] pem_keys: List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        :param pulumi.Input[str] token_reviewer_jwt: A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if disable_iss_validation is not None:
            pulumi.set(__self__, "disable_iss_validation", disable_iss_validation)
        if disable_local_ca_jwt is not None:
            pulumi.set(__self__, "disable_local_ca_jwt", disable_local_ca_jwt)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if kubernetes_ca_cert is not None:
            pulumi.set(__self__, "kubernetes_ca_cert", kubernetes_ca_cert)
        if kubernetes_host is not None:
            pulumi.set(__self__, "kubernetes_host", kubernetes_host)
        if pem_keys is not None:
            pulumi.set(__self__, "pem_keys", pem_keys)
        if token_reviewer_jwt is not None:
            pulumi.set(__self__, "token_reviewer_jwt", token_reviewer_jwt)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the kubernetes backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="disableIssValidation")
    def disable_iss_validation(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_iss_validation")

    @disable_iss_validation.setter
    def disable_iss_validation(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_iss_validation", value)

    @property
    @pulumi.getter(name="disableLocalCaJwt")
    def disable_local_ca_jwt(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_local_ca_jwt")

    @disable_local_ca_jwt.setter
    def disable_local_ca_jwt(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_local_ca_jwt", value)

    @property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[str]]:
        """
        JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "issuer", value)

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> Optional[pulumi.Input[str]]:
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @kubernetes_ca_cert.setter
    def kubernetes_ca_cert(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_ca_cert", value)

    @property
    @pulumi.getter(name="kubernetesHost")
    def kubernetes_host(self) -> Optional[pulumi.Input[str]]:
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        return pulumi.get(self, "kubernetes_host")

    @kubernetes_host.setter
    def kubernetes_host(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "kubernetes_host", value)

    @property
    @pulumi.getter(name="pemKeys")
    def pem_keys(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        return pulumi.get(self, "pem_keys")

    @pem_keys.setter
    def pem_keys(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "pem_keys", value)

    @property
    @pulumi.getter(name="tokenReviewerJwt")
    def token_reviewer_jwt(self) -> Optional[pulumi.Input[str]]:
        """
        A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        return pulumi.get(self, "token_reviewer_jwt")

    @token_reviewer_jwt.setter
    def token_reviewer_jwt(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_reviewer_jwt", value)


class AuthBackendConfig(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 disable_iss_validation: Optional[pulumi.Input[bool]] = None,
                 disable_local_ca_jwt: Optional[pulumi.Input[bool]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_host: Optional[pulumi.Input[str]] = None,
                 pem_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_reviewer_jwt: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages an Kubernetes auth backend config in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        kubernetes = vault.AuthBackend("kubernetes", type="kubernetes")
        example = vault.kubernetes.AuthBackendConfig("example",
            backend=kubernetes.path,
            kubernetes_host="http://example.com:443",
            kubernetes_ca_cert=\"\"\"-----BEGIN CERTIFICATE-----
        example
        -----END CERTIFICATE-----\"\"\",
            token_reviewer_jwt="ZXhhbXBsZQo=",
            issuer="api",
            disable_iss_validation=True)
        ```

        ## Import

        Kubernetes authentication backend can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:kubernetes/authBackendConfig:AuthBackendConfig config auth/kubernetes/config
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Unique name of the kubernetes backend to configure.
        :param pulumi.Input[bool] disable_iss_validation: Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[bool] disable_local_ca_jwt: Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[str] issuer: JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        :param pulumi.Input[str] kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        :param pulumi.Input[str] kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] pem_keys: List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        :param pulumi.Input[str] token_reviewer_jwt: A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AuthBackendConfigArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an Kubernetes auth backend config in a Vault server. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/kubernetes.html) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        kubernetes = vault.AuthBackend("kubernetes", type="kubernetes")
        example = vault.kubernetes.AuthBackendConfig("example",
            backend=kubernetes.path,
            kubernetes_host="http://example.com:443",
            kubernetes_ca_cert=\"\"\"-----BEGIN CERTIFICATE-----
        example
        -----END CERTIFICATE-----\"\"\",
            token_reviewer_jwt="ZXhhbXBsZQo=",
            issuer="api",
            disable_iss_validation=True)
        ```

        ## Import

        Kubernetes authentication backend can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:kubernetes/authBackendConfig:AuthBackendConfig config auth/kubernetes/config
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendConfigArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendConfigArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 disable_iss_validation: Optional[pulumi.Input[bool]] = None,
                 disable_local_ca_jwt: Optional[pulumi.Input[bool]] = None,
                 issuer: Optional[pulumi.Input[str]] = None,
                 kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
                 kubernetes_host: Optional[pulumi.Input[str]] = None,
                 pem_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 token_reviewer_jwt: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendConfigArgs.__new__(AuthBackendConfigArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["disable_iss_validation"] = disable_iss_validation
            __props__.__dict__["disable_local_ca_jwt"] = disable_local_ca_jwt
            __props__.__dict__["issuer"] = issuer
            __props__.__dict__["kubernetes_ca_cert"] = kubernetes_ca_cert
            if kubernetes_host is None and not opts.urn:
                raise TypeError("Missing required property 'kubernetes_host'")
            __props__.__dict__["kubernetes_host"] = kubernetes_host
            __props__.__dict__["pem_keys"] = pem_keys
            __props__.__dict__["token_reviewer_jwt"] = token_reviewer_jwt
        super(AuthBackendConfig, __self__).__init__(
            'vault:kubernetes/authBackendConfig:AuthBackendConfig',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            disable_iss_validation: Optional[pulumi.Input[bool]] = None,
            disable_local_ca_jwt: Optional[pulumi.Input[bool]] = None,
            issuer: Optional[pulumi.Input[str]] = None,
            kubernetes_ca_cert: Optional[pulumi.Input[str]] = None,
            kubernetes_host: Optional[pulumi.Input[str]] = None,
            pem_keys: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            token_reviewer_jwt: Optional[pulumi.Input[str]] = None) -> 'AuthBackendConfig':
        """
        Get an existing AuthBackendConfig resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Unique name of the kubernetes backend to configure.
        :param pulumi.Input[bool] disable_iss_validation: Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[bool] disable_local_ca_jwt: Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        :param pulumi.Input[str] issuer: JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        :param pulumi.Input[str] kubernetes_ca_cert: PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        :param pulumi.Input[str] kubernetes_host: Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] pem_keys: List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        :param pulumi.Input[str] token_reviewer_jwt: A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendConfigState.__new__(_AuthBackendConfigState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["disable_iss_validation"] = disable_iss_validation
        __props__.__dict__["disable_local_ca_jwt"] = disable_local_ca_jwt
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["kubernetes_ca_cert"] = kubernetes_ca_cert
        __props__.__dict__["kubernetes_host"] = kubernetes_host
        __props__.__dict__["pem_keys"] = pem_keys
        __props__.__dict__["token_reviewer_jwt"] = token_reviewer_jwt
        return AuthBackendConfig(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Unique name of the kubernetes backend to configure.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="disableIssValidation")
    def disable_iss_validation(self) -> pulumi.Output[bool]:
        """
        Disable JWT issuer validation. Allows to skip ISS validation. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_iss_validation")

    @property
    @pulumi.getter(name="disableLocalCaJwt")
    def disable_local_ca_jwt(self) -> pulumi.Output[bool]:
        """
        Disable defaulting to the local CA cert and service account JWT when running in a Kubernetes pod. Requires Vault `v1.5.4+` or Vault auth kubernetes plugin `v0.7.1+`
        """
        return pulumi.get(self, "disable_local_ca_jwt")

    @property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[Optional[str]]:
        """
        JWT issuer. If no issuer is specified, `kubernetes.io/serviceaccount` will be used as the default issuer.
        """
        return pulumi.get(self, "issuer")

    @property
    @pulumi.getter(name="kubernetesCaCert")
    def kubernetes_ca_cert(self) -> pulumi.Output[str]:
        """
        PEM encoded CA cert for use by the TLS client used to talk with the Kubernetes API.
        """
        return pulumi.get(self, "kubernetes_ca_cert")

    @property
    @pulumi.getter(name="kubernetesHost")
    def kubernetes_host(self) -> pulumi.Output[str]:
        """
        Host must be a host string, a host:port pair, or a URL to the base of the Kubernetes API server.
        """
        return pulumi.get(self, "kubernetes_host")

    @property
    @pulumi.getter(name="pemKeys")
    def pem_keys(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of PEM-formatted public keys or certificates used to verify the signatures of Kubernetes service account JWTs. If a certificate is given, its public key will be extracted. Not every installation of Kubernetes exposes these keys.
        """
        return pulumi.get(self, "pem_keys")

    @property
    @pulumi.getter(name="tokenReviewerJwt")
    def token_reviewer_jwt(self) -> pulumi.Output[Optional[str]]:
        """
        A service account JWT used to access the TokenReview API to validate other JWTs during login. If not set the JWT used for login will be used to access the API.
        """
        return pulumi.get(self, "token_reviewer_jwt")

