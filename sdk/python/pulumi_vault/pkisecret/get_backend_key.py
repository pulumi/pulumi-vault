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
    'GetBackendKeyResult',
    'AwaitableGetBackendKeyResult',
    'get_backend_key',
    'get_backend_key_output',
]

@pulumi.output_type
class GetBackendKeyResult:
    """
    A collection of values returned by getBackendKey.
    """
    def __init__(__self__, backend=None, id=None, key_id=None, key_name=None, key_ref=None, key_type=None, namespace=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_id and not isinstance(key_id, str):
            raise TypeError("Expected argument 'key_id' to be a str")
        pulumi.set(__self__, "key_id", key_id)
        if key_name and not isinstance(key_name, str):
            raise TypeError("Expected argument 'key_name' to be a str")
        pulumi.set(__self__, "key_name", key_name)
        if key_ref and not isinstance(key_ref, str):
            raise TypeError("Expected argument 'key_ref' to be a str")
        pulumi.set(__self__, "key_ref", key_ref)
        if key_type and not isinstance(key_type, str):
            raise TypeError("Expected argument 'key_type' to be a str")
        pulumi.set(__self__, "key_type", key_type)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> str:
        """
        ID of the key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> str:
        """
        Name of the key.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyRef")
    def key_ref(self) -> str:
        return pulumi.get(self, "key_ref")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> str:
        """
        Type of the key.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")


class AwaitableGetBackendKeyResult(GetBackendKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendKeyResult(
            backend=self.backend,
            id=self.id,
            key_id=self.key_id,
            key_name=self.key_name,
            key_ref=self.key_ref,
            key_type=self.key_type,
            namespace=self.namespace)


def get_backend_key(backend: Optional[str] = None,
                    key_ref: Optional[str] = None,
                    namespace: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackendKeyResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    key = vault.pki_secret.SecretBackendKey("key",
        backend=pki.path,
        type="internal",
        key_name="example",
        key_type="rsa",
        key_bits=4096)
    example = key.key_id.apply(lambda key_id: vault.pkiSecret.get_backend_key_output(backend=key_vault_mount["path"],
        key_ref=key_id))
    ```


    :param str backend: The path to the PKI secret backend to
           read the key from, with no leading or trailing `/`s.
    :param str key_ref: Reference to an existing key.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['keyRef'] = key_ref
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:pkiSecret/getBackendKey:getBackendKey', __args__, opts=opts, typ=GetBackendKeyResult).value

    return AwaitableGetBackendKeyResult(
        backend=pulumi.get(__ret__, 'backend'),
        id=pulumi.get(__ret__, 'id'),
        key_id=pulumi.get(__ret__, 'key_id'),
        key_name=pulumi.get(__ret__, 'key_name'),
        key_ref=pulumi.get(__ret__, 'key_ref'),
        key_type=pulumi.get(__ret__, 'key_type'),
        namespace=pulumi.get(__ret__, 'namespace'))


@_utilities.lift_output_func(get_backend_key)
def get_backend_key_output(backend: Optional[pulumi.Input[str]] = None,
                           key_ref: Optional[pulumi.Input[str]] = None,
                           namespace: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackendKeyResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    key = vault.pki_secret.SecretBackendKey("key",
        backend=pki.path,
        type="internal",
        key_name="example",
        key_type="rsa",
        key_bits=4096)
    example = key.key_id.apply(lambda key_id: vault.pkiSecret.get_backend_key_output(backend=key_vault_mount["path"],
        key_ref=key_id))
    ```


    :param str backend: The path to the PKI secret backend to
           read the key from, with no leading or trailing `/`s.
    :param str key_ref: Reference to an existing key.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    ...
