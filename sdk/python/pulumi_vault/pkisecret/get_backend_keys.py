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
    'GetBackendKeysResult',
    'AwaitableGetBackendKeysResult',
    'get_backend_keys',
    'get_backend_keys_output',
]

@pulumi.output_type
class GetBackendKeysResult:
    """
    A collection of values returned by getBackendKeys.
    """
    def __init__(__self__, backend=None, id=None, key_info=None, key_info_json=None, keys=None, namespace=None):
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if key_info and not isinstance(key_info, dict):
            raise TypeError("Expected argument 'key_info' to be a dict")
        pulumi.set(__self__, "key_info", key_info)
        if key_info_json and not isinstance(key_info_json, str):
            raise TypeError("Expected argument 'key_info_json' to be a str")
        pulumi.set(__self__, "key_info_json", key_info_json)
        if keys and not isinstance(keys, list):
            raise TypeError("Expected argument 'keys' to be a list")
        pulumi.set(__self__, "keys", keys)
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
    @pulumi.getter(name="keyInfo")
    def key_info(self) -> Mapping[str, str]:
        """
        Map of key strings read from Vault.
        """
        return pulumi.get(self, "key_info")

    @property
    @pulumi.getter(name="keyInfoJson")
    def key_info_json(self) -> str:
        """
        JSON-encoded key data read from Vault.
        """
        return pulumi.get(self, "key_info_json")

    @property
    @pulumi.getter
    def keys(self) -> Sequence[str]:
        """
        Keys used under the backend path.
        """
        return pulumi.get(self, "keys")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")


class AwaitableGetBackendKeysResult(GetBackendKeysResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendKeysResult(
            backend=self.backend,
            id=self.id,
            key_info=self.key_info,
            key_info_json=self.key_info_json,
            keys=self.keys,
            namespace=self.namespace)


def get_backend_keys(backend: Optional[str] = None,
                     namespace: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackendKeysResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    root = vault.pki_secret.SecretBackendRootCert("root",
        backend=pki.path,
        type="internal",
        common_name="example",
        ttl="86400",
        key_name="example")
    example = vault.pkiSecret.get_backend_keys_output(backend=root.backend)
    ```


    :param str backend: The path to the PKI secret backend to
           read the keys from, with no leading or trailing `/`s.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:pkiSecret/getBackendKeys:getBackendKeys', __args__, opts=opts, typ=GetBackendKeysResult).value

    return AwaitableGetBackendKeysResult(
        backend=pulumi.get(__ret__, 'backend'),
        id=pulumi.get(__ret__, 'id'),
        key_info=pulumi.get(__ret__, 'key_info'),
        key_info_json=pulumi.get(__ret__, 'key_info_json'),
        keys=pulumi.get(__ret__, 'keys'),
        namespace=pulumi.get(__ret__, 'namespace'))


@_utilities.lift_output_func(get_backend_keys)
def get_backend_keys_output(backend: Optional[pulumi.Input[str]] = None,
                            namespace: Optional[pulumi.Input[Optional[str]]] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackendKeysResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    root = vault.pki_secret.SecretBackendRootCert("root",
        backend=pki.path,
        type="internal",
        common_name="example",
        ttl="86400",
        key_name="example")
    example = vault.pkiSecret.get_backend_keys_output(backend=root.backend)
    ```


    :param str backend: The path to the PKI secret backend to
           read the keys from, with no leading or trailing `/`s.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    ...
