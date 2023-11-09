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
    'GetDecodeResult',
    'AwaitableGetDecodeResult',
    'get_decode',
    'get_decode_output',
]

@pulumi.output_type
class GetDecodeResult:
    """
    A collection of values returned by getDecode.
    """
    def __init__(__self__, batch_inputs=None, batch_results=None, decoded_value=None, id=None, namespace=None, path=None, role_name=None, transformation=None, tweak=None, value=None):
        if batch_inputs and not isinstance(batch_inputs, list):
            raise TypeError("Expected argument 'batch_inputs' to be a list")
        pulumi.set(__self__, "batch_inputs", batch_inputs)
        if batch_results and not isinstance(batch_results, list):
            raise TypeError("Expected argument 'batch_results' to be a list")
        pulumi.set(__self__, "batch_results", batch_results)
        if decoded_value and not isinstance(decoded_value, str):
            raise TypeError("Expected argument 'decoded_value' to be a str")
        pulumi.set(__self__, "decoded_value", decoded_value)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if role_name and not isinstance(role_name, str):
            raise TypeError("Expected argument 'role_name' to be a str")
        pulumi.set(__self__, "role_name", role_name)
        if transformation and not isinstance(transformation, str):
            raise TypeError("Expected argument 'transformation' to be a str")
        pulumi.set(__self__, "transformation", transformation)
        if tweak and not isinstance(tweak, str):
            raise TypeError("Expected argument 'tweak' to be a str")
        pulumi.set(__self__, "tweak", tweak)
        if value and not isinstance(value, str):
            raise TypeError("Expected argument 'value' to be a str")
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="batchInputs")
    def batch_inputs(self) -> Optional[Sequence[Mapping[str, Any]]]:
        return pulumi.get(self, "batch_inputs")

    @property
    @pulumi.getter(name="batchResults")
    def batch_results(self) -> Sequence[Mapping[str, Any]]:
        return pulumi.get(self, "batch_results")

    @property
    @pulumi.getter(name="decodedValue")
    def decoded_value(self) -> str:
        return pulumi.get(self, "decoded_value")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter(name="roleName")
    def role_name(self) -> str:
        return pulumi.get(self, "role_name")

    @property
    @pulumi.getter
    def transformation(self) -> Optional[str]:
        return pulumi.get(self, "transformation")

    @property
    @pulumi.getter
    def tweak(self) -> Optional[str]:
        return pulumi.get(self, "tweak")

    @property
    @pulumi.getter
    def value(self) -> Optional[str]:
        return pulumi.get(self, "value")


class AwaitableGetDecodeResult(GetDecodeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDecodeResult(
            batch_inputs=self.batch_inputs,
            batch_results=self.batch_results,
            decoded_value=self.decoded_value,
            id=self.id,
            namespace=self.namespace,
            path=self.path,
            role_name=self.role_name,
            transformation=self.transformation,
            tweak=self.tweak,
            value=self.value)


def get_decode(batch_inputs: Optional[Sequence[Mapping[str, Any]]] = None,
               batch_results: Optional[Sequence[Mapping[str, Any]]] = None,
               decoded_value: Optional[str] = None,
               namespace: Optional[str] = None,
               path: Optional[str] = None,
               role_name: Optional[str] = None,
               transformation: Optional[str] = None,
               tweak: Optional[str] = None,
               value: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDecodeResult:
    """
    This data source supports the "/transform/decode/{role_name}" Vault endpoint.

    It decodes the provided value using a named role.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    transform = vault.Mount("transform",
        path="transform",
        type="transform")
    ccn_fpe = vault.transform.Transformation("ccn-fpe",
        path=transform.path,
        type="fpe",
        template="builtin/creditcardnumber",
        tweak_source="internal",
        allowed_roles=["payments"])
    payments = vault.transform.Role("payments",
        path=ccn_fpe.path,
        transformations=["ccn-fpe"])
    test = vault.transform.get_decode_output(path=payments.path,
        role_name="payments",
        value="9300-3376-4943-8903")
    ```


    :param Sequence[Mapping[str, Any]] batch_inputs: Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
    :param Sequence[Mapping[str, Any]] batch_results: The result of decoding a batch.
    :param str decoded_value: The result of decoding a value.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Path to where the back-end is mounted within Vault.
    :param str role_name: The name of the role.
    :param str transformation: The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
    :param str tweak: The tweak value to use. Only applicable for FPE transformations
    :param str value: The value in which to decode.
    """
    __args__ = dict()
    __args__['batchInputs'] = batch_inputs
    __args__['batchResults'] = batch_results
    __args__['decodedValue'] = decoded_value
    __args__['namespace'] = namespace
    __args__['path'] = path
    __args__['roleName'] = role_name
    __args__['transformation'] = transformation
    __args__['tweak'] = tweak
    __args__['value'] = value
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:transform/getDecode:getDecode', __args__, opts=opts, typ=GetDecodeResult).value

    return AwaitableGetDecodeResult(
        batch_inputs=pulumi.get(__ret__, 'batch_inputs'),
        batch_results=pulumi.get(__ret__, 'batch_results'),
        decoded_value=pulumi.get(__ret__, 'decoded_value'),
        id=pulumi.get(__ret__, 'id'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'),
        role_name=pulumi.get(__ret__, 'role_name'),
        transformation=pulumi.get(__ret__, 'transformation'),
        tweak=pulumi.get(__ret__, 'tweak'),
        value=pulumi.get(__ret__, 'value'))


@_utilities.lift_output_func(get_decode)
def get_decode_output(batch_inputs: Optional[pulumi.Input[Optional[Sequence[Mapping[str, Any]]]]] = None,
                      batch_results: Optional[pulumi.Input[Optional[Sequence[Mapping[str, Any]]]]] = None,
                      decoded_value: Optional[pulumi.Input[Optional[str]]] = None,
                      namespace: Optional[pulumi.Input[Optional[str]]] = None,
                      path: Optional[pulumi.Input[str]] = None,
                      role_name: Optional[pulumi.Input[str]] = None,
                      transformation: Optional[pulumi.Input[Optional[str]]] = None,
                      tweak: Optional[pulumi.Input[Optional[str]]] = None,
                      value: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDecodeResult]:
    """
    This data source supports the "/transform/decode/{role_name}" Vault endpoint.

    It decodes the provided value using a named role.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    transform = vault.Mount("transform",
        path="transform",
        type="transform")
    ccn_fpe = vault.transform.Transformation("ccn-fpe",
        path=transform.path,
        type="fpe",
        template="builtin/creditcardnumber",
        tweak_source="internal",
        allowed_roles=["payments"])
    payments = vault.transform.Role("payments",
        path=ccn_fpe.path,
        transformations=["ccn-fpe"])
    test = vault.transform.get_decode_output(path=payments.path,
        role_name="payments",
        value="9300-3376-4943-8903")
    ```


    :param Sequence[Mapping[str, Any]] batch_inputs: Specifies a list of items to be decoded in a single batch. If this parameter is set, the top-level parameters 'value', 'transformation' and 'tweak' will be ignored. Each batch item within the list can specify these parameters instead.
    :param Sequence[Mapping[str, Any]] batch_results: The result of decoding a batch.
    :param str decoded_value: The result of decoding a value.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    :param str path: Path to where the back-end is mounted within Vault.
    :param str role_name: The name of the role.
    :param str transformation: The transformation to perform. If no value is provided and the role contains a single transformation, this value will be inferred from the role.
    :param str tweak: The tweak value to use. Only applicable for FPE transformations
    :param str value: The value in which to decode.
    """
    ...
