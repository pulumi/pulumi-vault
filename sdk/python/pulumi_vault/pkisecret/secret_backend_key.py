# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

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

__all__ = ['SecretBackendKeyArgs', 'SecretBackendKey']

@pulumi.input_type
class SecretBackendKeyArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 type: pulumi.Input[str],
                 key_bits: Optional[pulumi.Input[int]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 managed_key_id: Optional[pulumi.Input[str]] = None,
                 managed_key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackendKey resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[str] type: Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        :param pulumi.Input[int] key_bits: Specifies the number of bits to use for the generated keys. 
               Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
               2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
               384, or 521; ignored with `key_type=ed25519`.
        :param pulumi.Input[str] key_name: When a new key is created with this request, optionally specifies the name for this. 
               The global ref `default` may not be used as a name.
        :param pulumi.Input[str] key_type: Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        :param pulumi.Input[str] managed_key_id: The managed key's UUID.
        :param pulumi.Input[str] managed_key_name: The managed key's configured name.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "type", type)
        if key_bits is not None:
            pulumi.set(__self__, "key_bits", key_bits)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if managed_key_id is not None:
            pulumi.set(__self__, "managed_key_id", managed_key_id)
        if managed_key_name is not None:
            pulumi.set(__self__, "managed_key_name", managed_key_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="keyBits")
    def key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of bits to use for the generated keys. 
        Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
        2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
        384, or 521; ignored with `key_type=ed25519`.
        """
        return pulumi.get(self, "key_bits")

    @key_bits.setter
    def key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key_bits", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        When a new key is created with this request, optionally specifies the name for this. 
        The global ref `default` may not be used as a name.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter(name="managedKeyId")
    def managed_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed key's UUID.
        """
        return pulumi.get(self, "managed_key_id")

    @managed_key_id.setter
    def managed_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_key_id", value)

    @property
    @pulumi.getter(name="managedKeyName")
    def managed_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The managed key's configured name.
        """
        return pulumi.get(self, "managed_key_name")

    @managed_key_name.setter
    def managed_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_key_name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _SecretBackendKeyState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 key_bits: Optional[pulumi.Input[int]] = None,
                 key_id: Optional[pulumi.Input[str]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 managed_key_id: Optional[pulumi.Input[str]] = None,
                 managed_key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackendKey resources.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] key_bits: Specifies the number of bits to use for the generated keys. 
               Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
               2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
               384, or 521; ignored with `key_type=ed25519`.
        :param pulumi.Input[str] key_id: ID of the generated key.
        :param pulumi.Input[str] key_name: When a new key is created with this request, optionally specifies the name for this. 
               The global ref `default` may not be used as a name.
        :param pulumi.Input[str] key_type: Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        :param pulumi.Input[str] managed_key_id: The managed key's UUID.
        :param pulumi.Input[str] managed_key_name: The managed key's configured name.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] type: Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if key_bits is not None:
            pulumi.set(__self__, "key_bits", key_bits)
        if key_id is not None:
            pulumi.set(__self__, "key_id", key_id)
        if key_name is not None:
            pulumi.set(__self__, "key_name", key_name)
        if key_type is not None:
            pulumi.set(__self__, "key_type", key_type)
        if managed_key_id is not None:
            pulumi.set(__self__, "managed_key_id", managed_key_id)
        if managed_key_name is not None:
            pulumi.set(__self__, "managed_key_name", managed_key_name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="keyBits")
    def key_bits(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the number of bits to use for the generated keys. 
        Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
        2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
        384, or 521; ignored with `key_type=ed25519`.
        """
        return pulumi.get(self, "key_bits")

    @key_bits.setter
    def key_bits(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "key_bits", value)

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the generated key.
        """
        return pulumi.get(self, "key_id")

    @key_id.setter
    def key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_id", value)

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> Optional[pulumi.Input[str]]:
        """
        When a new key is created with this request, optionally specifies the name for this. 
        The global ref `default` may not be used as a name.
        """
        return pulumi.get(self, "key_name")

    @key_name.setter
    def key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_name", value)

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        """
        return pulumi.get(self, "key_type")

    @key_type.setter
    def key_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_type", value)

    @property
    @pulumi.getter(name="managedKeyId")
    def managed_key_id(self) -> Optional[pulumi.Input[str]]:
        """
        The managed key's UUID.
        """
        return pulumi.get(self, "managed_key_id")

    @managed_key_id.setter
    def managed_key_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_key_id", value)

    @property
    @pulumi.getter(name="managedKeyName")
    def managed_key_name(self) -> Optional[pulumi.Input[str]]:
        """
        The managed key's configured name.
        """
        return pulumi.get(self, "managed_key_name")

    @managed_key_name.setter
    def managed_key_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "managed_key_name", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SecretBackendKey(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 key_bits: Optional[pulumi.Input[int]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 managed_key_id: Optional[pulumi.Input[str]] = None,
                 managed_key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a key on a PKI Secret Backend for Vault.

        ## Import

        PKI secret backend key can be imported using the `id`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendKey:SecretBackendKey key pki/key/bf9b0d48-d0dd-652c-30be-77d04fc7e94d
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] key_bits: Specifies the number of bits to use for the generated keys. 
               Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
               2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
               384, or 521; ignored with `key_type=ed25519`.
        :param pulumi.Input[str] key_name: When a new key is created with this request, optionally specifies the name for this. 
               The global ref `default` may not be used as a name.
        :param pulumi.Input[str] key_type: Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        :param pulumi.Input[str] managed_key_id: The managed key's UUID.
        :param pulumi.Input[str] managed_key_name: The managed key's configured name.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] type: Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendKeyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a key on a PKI Secret Backend for Vault.

        ## Import

        PKI secret backend key can be imported using the `id`, e.g.

        ```sh
        $ pulumi import vault:pkiSecret/secretBackendKey:SecretBackendKey key pki/key/bf9b0d48-d0dd-652c-30be-77d04fc7e94d
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendKeyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendKeyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 key_bits: Optional[pulumi.Input[int]] = None,
                 key_name: Optional[pulumi.Input[str]] = None,
                 key_type: Optional[pulumi.Input[str]] = None,
                 managed_key_id: Optional[pulumi.Input[str]] = None,
                 managed_key_name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendKeyArgs.__new__(SecretBackendKeyArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["key_bits"] = key_bits
            __props__.__dict__["key_name"] = key_name
            __props__.__dict__["key_type"] = key_type
            __props__.__dict__["managed_key_id"] = managed_key_id
            __props__.__dict__["managed_key_name"] = managed_key_name
            __props__.__dict__["namespace"] = namespace
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["key_id"] = None
        super(SecretBackendKey, __self__).__init__(
            'vault:pkiSecret/secretBackendKey:SecretBackendKey',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            key_bits: Optional[pulumi.Input[int]] = None,
            key_id: Optional[pulumi.Input[str]] = None,
            key_name: Optional[pulumi.Input[str]] = None,
            key_type: Optional[pulumi.Input[str]] = None,
            managed_key_id: Optional[pulumi.Input[str]] = None,
            managed_key_name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SecretBackendKey':
        """
        Get an existing SecretBackendKey resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        :param pulumi.Input[int] key_bits: Specifies the number of bits to use for the generated keys. 
               Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
               2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
               384, or 521; ignored with `key_type=ed25519`.
        :param pulumi.Input[str] key_id: ID of the generated key.
        :param pulumi.Input[str] key_name: When a new key is created with this request, optionally specifies the name for this. 
               The global ref `default` may not be used as a name.
        :param pulumi.Input[str] key_type: Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        :param pulumi.Input[str] managed_key_id: The managed key's UUID.
        :param pulumi.Input[str] managed_key_name: The managed key's configured name.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] type: Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendKeyState.__new__(_SecretBackendKeyState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["key_bits"] = key_bits
        __props__.__dict__["key_id"] = key_id
        __props__.__dict__["key_name"] = key_name
        __props__.__dict__["key_type"] = key_type
        __props__.__dict__["managed_key_id"] = managed_key_id
        __props__.__dict__["managed_key_name"] = managed_key_name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["type"] = type
        return SecretBackendKey(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the PKI secret backend is mounted at, with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="keyBits")
    def key_bits(self) -> pulumi.Output[int]:
        """
        Specifies the number of bits to use for the generated keys. 
        Allowed values are 0 (universal default); with `key_type=rsa`, allowed values are:
        2048 (default), 3072, or 4096; with `key_type=ec`, allowed values are: 224, 256 (default),
        384, or 521; ignored with `key_type=ed25519`.
        """
        return pulumi.get(self, "key_bits")

    @property
    @pulumi.getter(name="keyId")
    def key_id(self) -> pulumi.Output[str]:
        """
        ID of the generated key.
        """
        return pulumi.get(self, "key_id")

    @property
    @pulumi.getter(name="keyName")
    def key_name(self) -> pulumi.Output[Optional[str]]:
        """
        When a new key is created with this request, optionally specifies the name for this. 
        The global ref `default` may not be used as a name.
        """
        return pulumi.get(self, "key_name")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> pulumi.Output[str]:
        """
        Specifies the desired key type; must be `rsa`, `ed25519` or `ec`.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter(name="managedKeyId")
    def managed_key_id(self) -> pulumi.Output[Optional[str]]:
        """
        The managed key's UUID.
        """
        return pulumi.get(self, "managed_key_id")

    @property
    @pulumi.getter(name="managedKeyName")
    def managed_key_name(self) -> pulumi.Output[Optional[str]]:
        """
        The managed key's configured name.
        """
        return pulumi.get(self, "managed_key_name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Specifies the type of the key to create. Can be `exported`,`internal` or `kms`.
        """
        return pulumi.get(self, "type")

