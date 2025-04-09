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
    'GetVerifyResult',
    'AwaitableGetVerifyResult',
    'get_verify',
    'get_verify_output',
]

@pulumi.output_type
class GetVerifyResult:
    """
    A collection of values returned by getVerify.
    """
    def __init__(__self__, batch_inputs=None, batch_results=None, cmac=None, context=None, hash_algorithm=None, hmac=None, id=None, input=None, marshaling_algorithm=None, name=None, namespace=None, path=None, prehashed=None, reference=None, salt_length=None, signature=None, signature_algorithm=None, signature_context=None, valid=None):
        if batch_inputs and not isinstance(batch_inputs, list):
            raise TypeError("Expected argument 'batch_inputs' to be a list")
        pulumi.set(__self__, "batch_inputs", batch_inputs)
        if batch_results and not isinstance(batch_results, list):
            raise TypeError("Expected argument 'batch_results' to be a list")
        pulumi.set(__self__, "batch_results", batch_results)
        if cmac and not isinstance(cmac, str):
            raise TypeError("Expected argument 'cmac' to be a str")
        pulumi.set(__self__, "cmac", cmac)
        if context and not isinstance(context, str):
            raise TypeError("Expected argument 'context' to be a str")
        pulumi.set(__self__, "context", context)
        if hash_algorithm and not isinstance(hash_algorithm, str):
            raise TypeError("Expected argument 'hash_algorithm' to be a str")
        pulumi.set(__self__, "hash_algorithm", hash_algorithm)
        if hmac and not isinstance(hmac, str):
            raise TypeError("Expected argument 'hmac' to be a str")
        pulumi.set(__self__, "hmac", hmac)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if input and not isinstance(input, str):
            raise TypeError("Expected argument 'input' to be a str")
        pulumi.set(__self__, "input", input)
        if marshaling_algorithm and not isinstance(marshaling_algorithm, str):
            raise TypeError("Expected argument 'marshaling_algorithm' to be a str")
        pulumi.set(__self__, "marshaling_algorithm", marshaling_algorithm)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if path and not isinstance(path, str):
            raise TypeError("Expected argument 'path' to be a str")
        pulumi.set(__self__, "path", path)
        if prehashed and not isinstance(prehashed, bool):
            raise TypeError("Expected argument 'prehashed' to be a bool")
        pulumi.set(__self__, "prehashed", prehashed)
        if reference and not isinstance(reference, str):
            raise TypeError("Expected argument 'reference' to be a str")
        pulumi.set(__self__, "reference", reference)
        if salt_length and not isinstance(salt_length, str):
            raise TypeError("Expected argument 'salt_length' to be a str")
        pulumi.set(__self__, "salt_length", salt_length)
        if signature and not isinstance(signature, str):
            raise TypeError("Expected argument 'signature' to be a str")
        pulumi.set(__self__, "signature", signature)
        if signature_algorithm and not isinstance(signature_algorithm, str):
            raise TypeError("Expected argument 'signature_algorithm' to be a str")
        pulumi.set(__self__, "signature_algorithm", signature_algorithm)
        if signature_context and not isinstance(signature_context, str):
            raise TypeError("Expected argument 'signature_context' to be a str")
        pulumi.set(__self__, "signature_context", signature_context)
        if valid and not isinstance(valid, bool):
            raise TypeError("Expected argument 'valid' to be a bool")
        pulumi.set(__self__, "valid", valid)

    @property
    @pulumi.getter(name="batchInputs")
    def batch_inputs(self) -> Optional[Sequence[Mapping[str, builtins.str]]]:
        return pulumi.get(self, "batch_inputs")

    @property
    @pulumi.getter(name="batchResults")
    def batch_results(self) -> Sequence[Mapping[str, builtins.str]]:
        """
        The results returned from Vault if using `batch_input`
        """
        return pulumi.get(self, "batch_results")

    @property
    @pulumi.getter
    def cmac(self) -> Optional[builtins.str]:
        return pulumi.get(self, "cmac")

    @property
    @pulumi.getter
    def context(self) -> Optional[builtins.str]:
        return pulumi.get(self, "context")

    @property
    @pulumi.getter(name="hashAlgorithm")
    def hash_algorithm(self) -> Optional[builtins.str]:
        return pulumi.get(self, "hash_algorithm")

    @property
    @pulumi.getter
    def hmac(self) -> Optional[builtins.str]:
        return pulumi.get(self, "hmac")

    @property
    @pulumi.getter
    def id(self) -> builtins.str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def input(self) -> Optional[builtins.str]:
        return pulumi.get(self, "input")

    @property
    @pulumi.getter(name="marshalingAlgorithm")
    def marshaling_algorithm(self) -> Optional[builtins.str]:
        return pulumi.get(self, "marshaling_algorithm")

    @property
    @pulumi.getter
    def name(self) -> builtins.str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[builtins.str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> builtins.str:
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def prehashed(self) -> Optional[builtins.bool]:
        return pulumi.get(self, "prehashed")

    @property
    @pulumi.getter
    def reference(self) -> Optional[builtins.str]:
        return pulumi.get(self, "reference")

    @property
    @pulumi.getter(name="saltLength")
    def salt_length(self) -> Optional[builtins.str]:
        return pulumi.get(self, "salt_length")

    @property
    @pulumi.getter
    def signature(self) -> Optional[builtins.str]:
        return pulumi.get(self, "signature")

    @property
    @pulumi.getter(name="signatureAlgorithm")
    def signature_algorithm(self) -> Optional[builtins.str]:
        return pulumi.get(self, "signature_algorithm")

    @property
    @pulumi.getter(name="signatureContext")
    def signature_context(self) -> Optional[builtins.str]:
        return pulumi.get(self, "signature_context")

    @property
    @pulumi.getter
    def valid(self) -> builtins.bool:
        """
        Returns `true` if the signature verification succeeded and `false` otherwise
        """
        return pulumi.get(self, "valid")


class AwaitableGetVerifyResult(GetVerifyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVerifyResult(
            batch_inputs=self.batch_inputs,
            batch_results=self.batch_results,
            cmac=self.cmac,
            context=self.context,
            hash_algorithm=self.hash_algorithm,
            hmac=self.hmac,
            id=self.id,
            input=self.input,
            marshaling_algorithm=self.marshaling_algorithm,
            name=self.name,
            namespace=self.namespace,
            path=self.path,
            prehashed=self.prehashed,
            reference=self.reference,
            salt_length=self.salt_length,
            signature=self.signature,
            signature_algorithm=self.signature_algorithm,
            signature_context=self.signature_context,
            valid=self.valid)


def get_verify(batch_inputs: Optional[Sequence[Mapping[str, builtins.str]]] = None,
               batch_results: Optional[Sequence[Mapping[str, builtins.str]]] = None,
               cmac: Optional[builtins.str] = None,
               context: Optional[builtins.str] = None,
               hash_algorithm: Optional[builtins.str] = None,
               hmac: Optional[builtins.str] = None,
               input: Optional[builtins.str] = None,
               marshaling_algorithm: Optional[builtins.str] = None,
               name: Optional[builtins.str] = None,
               namespace: Optional[builtins.str] = None,
               path: Optional[builtins.str] = None,
               prehashed: Optional[builtins.bool] = None,
               reference: Optional[builtins.str] = None,
               salt_length: Optional[builtins.str] = None,
               signature: Optional[builtins.str] = None,
               signature_algorithm: Optional[builtins.str] = None,
               signature_context: Optional[builtins.str] = None,
               valid: Optional[builtins.bool] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVerifyResult:
    """
    This is a data source which can be used to verify a signature using a Vault Transit key.


    :param Sequence[Mapping[str, builtins.str]] batch_results: The results returned from Vault if using `batch_input`
    :param builtins.bool valid: Returns `true` if the signature verification succeeded and `false` otherwise
    """
    __args__ = dict()
    __args__['batchInputs'] = batch_inputs
    __args__['batchResults'] = batch_results
    __args__['cmac'] = cmac
    __args__['context'] = context
    __args__['hashAlgorithm'] = hash_algorithm
    __args__['hmac'] = hmac
    __args__['input'] = input
    __args__['marshalingAlgorithm'] = marshaling_algorithm
    __args__['name'] = name
    __args__['namespace'] = namespace
    __args__['path'] = path
    __args__['prehashed'] = prehashed
    __args__['reference'] = reference
    __args__['saltLength'] = salt_length
    __args__['signature'] = signature
    __args__['signatureAlgorithm'] = signature_algorithm
    __args__['signatureContext'] = signature_context
    __args__['valid'] = valid
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:transit/getVerify:getVerify', __args__, opts=opts, typ=GetVerifyResult).value

    return AwaitableGetVerifyResult(
        batch_inputs=pulumi.get(__ret__, 'batch_inputs'),
        batch_results=pulumi.get(__ret__, 'batch_results'),
        cmac=pulumi.get(__ret__, 'cmac'),
        context=pulumi.get(__ret__, 'context'),
        hash_algorithm=pulumi.get(__ret__, 'hash_algorithm'),
        hmac=pulumi.get(__ret__, 'hmac'),
        id=pulumi.get(__ret__, 'id'),
        input=pulumi.get(__ret__, 'input'),
        marshaling_algorithm=pulumi.get(__ret__, 'marshaling_algorithm'),
        name=pulumi.get(__ret__, 'name'),
        namespace=pulumi.get(__ret__, 'namespace'),
        path=pulumi.get(__ret__, 'path'),
        prehashed=pulumi.get(__ret__, 'prehashed'),
        reference=pulumi.get(__ret__, 'reference'),
        salt_length=pulumi.get(__ret__, 'salt_length'),
        signature=pulumi.get(__ret__, 'signature'),
        signature_algorithm=pulumi.get(__ret__, 'signature_algorithm'),
        signature_context=pulumi.get(__ret__, 'signature_context'),
        valid=pulumi.get(__ret__, 'valid'))
def get_verify_output(batch_inputs: Optional[pulumi.Input[Optional[Sequence[Mapping[str, builtins.str]]]]] = None,
                      batch_results: Optional[pulumi.Input[Optional[Sequence[Mapping[str, builtins.str]]]]] = None,
                      cmac: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      context: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      hash_algorithm: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      hmac: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      input: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      marshaling_algorithm: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      name: Optional[pulumi.Input[builtins.str]] = None,
                      namespace: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      path: Optional[pulumi.Input[builtins.str]] = None,
                      prehashed: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                      reference: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      salt_length: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      signature: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      signature_algorithm: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      signature_context: Optional[pulumi.Input[Optional[builtins.str]]] = None,
                      valid: Optional[pulumi.Input[Optional[builtins.bool]]] = None,
                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetVerifyResult]:
    """
    This is a data source which can be used to verify a signature using a Vault Transit key.


    :param Sequence[Mapping[str, builtins.str]] batch_results: The results returned from Vault if using `batch_input`
    :param builtins.bool valid: Returns `true` if the signature verification succeeded and `false` otherwise
    """
    __args__ = dict()
    __args__['batchInputs'] = batch_inputs
    __args__['batchResults'] = batch_results
    __args__['cmac'] = cmac
    __args__['context'] = context
    __args__['hashAlgorithm'] = hash_algorithm
    __args__['hmac'] = hmac
    __args__['input'] = input
    __args__['marshalingAlgorithm'] = marshaling_algorithm
    __args__['name'] = name
    __args__['namespace'] = namespace
    __args__['path'] = path
    __args__['prehashed'] = prehashed
    __args__['reference'] = reference
    __args__['saltLength'] = salt_length
    __args__['signature'] = signature
    __args__['signatureAlgorithm'] = signature_algorithm
    __args__['signatureContext'] = signature_context
    __args__['valid'] = valid
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vault:transit/getVerify:getVerify', __args__, opts=opts, typ=GetVerifyResult)
    return __ret__.apply(lambda __response__: GetVerifyResult(
        batch_inputs=pulumi.get(__response__, 'batch_inputs'),
        batch_results=pulumi.get(__response__, 'batch_results'),
        cmac=pulumi.get(__response__, 'cmac'),
        context=pulumi.get(__response__, 'context'),
        hash_algorithm=pulumi.get(__response__, 'hash_algorithm'),
        hmac=pulumi.get(__response__, 'hmac'),
        id=pulumi.get(__response__, 'id'),
        input=pulumi.get(__response__, 'input'),
        marshaling_algorithm=pulumi.get(__response__, 'marshaling_algorithm'),
        name=pulumi.get(__response__, 'name'),
        namespace=pulumi.get(__response__, 'namespace'),
        path=pulumi.get(__response__, 'path'),
        prehashed=pulumi.get(__response__, 'prehashed'),
        reference=pulumi.get(__response__, 'reference'),
        salt_length=pulumi.get(__response__, 'salt_length'),
        signature=pulumi.get(__response__, 'signature'),
        signature_algorithm=pulumi.get(__response__, 'signature_algorithm'),
        signature_context=pulumi.get(__response__, 'signature_context'),
        valid=pulumi.get(__response__, 'valid')))
