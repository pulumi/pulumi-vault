# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['MfaTotpArgs', 'MfaTotp']

@pulumi.input_type
class MfaTotpArgs:
    def __init__(__self__, *,
                 issuer: pulumi.Input[_builtins.str],
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 digits: Optional[pulumi.Input[_builtins.int]] = None,
                 key_size: Optional[pulumi.Input[_builtins.int]] = None,
                 max_validation_attempts: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 period: Optional[pulumi.Input[_builtins.int]] = None,
                 qr_size: Optional[pulumi.Input[_builtins.int]] = None,
                 skew: Optional[pulumi.Input[_builtins.int]] = None):
        """
        The set of arguments for constructing a MfaTotp resource.
        :param pulumi.Input[_builtins.str] issuer: The name of the key's issuing organization.
        :param pulumi.Input[_builtins.str] algorithm: Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        :param pulumi.Input[_builtins.int] digits: The number of digits in the generated TOTP token. This value can either be 6 or 8
        :param pulumi.Input[_builtins.int] key_size: Specifies the size in bytes of the generated key.
        :param pulumi.Input[_builtins.int] max_validation_attempts: The maximum number of consecutive failed validation attempts allowed.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.int] period: The length of time in seconds used to generate a counter for the TOTP token calculation.
        :param pulumi.Input[_builtins.int] qr_size: The pixel size of the generated square QR code.
        :param pulumi.Input[_builtins.int] skew: The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        """
        pulumi.set(__self__, "issuer", issuer)
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if key_size is not None:
            pulumi.set(__self__, "key_size", key_size)
        if max_validation_attempts is not None:
            pulumi.set(__self__, "max_validation_attempts", max_validation_attempts)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if qr_size is not None:
            pulumi.set(__self__, "qr_size", qr_size)
        if skew is not None:
            pulumi.set(__self__, "skew", skew)

    @_builtins.property
    @pulumi.getter
    def issuer(self) -> pulumi.Input[_builtins.str]:
        """
        The name of the key's issuing organization.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "issuer", value)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "algorithm", value)

    @_builtins.property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The number of digits in the generated TOTP token. This value can either be 6 or 8
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "digits", value)

    @_builtins.property
    @pulumi.getter(name="keySize")
    def key_size(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Specifies the size in bytes of the generated key.
        """
        return pulumi.get(self, "key_size")

    @key_size.setter
    def key_size(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "key_size", value)

    @_builtins.property
    @pulumi.getter(name="maxValidationAttempts")
    def max_validation_attempts(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The maximum number of consecutive failed validation attempts allowed.
        """
        return pulumi.get(self, "max_validation_attempts")

    @max_validation_attempts.setter
    def max_validation_attempts(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "max_validation_attempts", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The length of time in seconds used to generate a counter for the TOTP token calculation.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "period", value)

    @_builtins.property
    @pulumi.getter(name="qrSize")
    def qr_size(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The pixel size of the generated square QR code.
        """
        return pulumi.get(self, "qr_size")

    @qr_size.setter
    def qr_size(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "qr_size", value)

    @_builtins.property
    @pulumi.getter
    def skew(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        """
        return pulumi.get(self, "skew")

    @skew.setter
    def skew(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "skew", value)


@pulumi.input_type
class _MfaTotpState:
    def __init__(__self__, *,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 digits: Optional[pulumi.Input[_builtins.int]] = None,
                 issuer: Optional[pulumi.Input[_builtins.str]] = None,
                 key_size: Optional[pulumi.Input[_builtins.int]] = None,
                 max_validation_attempts: Optional[pulumi.Input[_builtins.int]] = None,
                 method_id: Optional[pulumi.Input[_builtins.str]] = None,
                 mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
                 period: Optional[pulumi.Input[_builtins.int]] = None,
                 qr_size: Optional[pulumi.Input[_builtins.int]] = None,
                 skew: Optional[pulumi.Input[_builtins.int]] = None,
                 type: Optional[pulumi.Input[_builtins.str]] = None,
                 uuid: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering MfaTotp resources.
        :param pulumi.Input[_builtins.str] algorithm: Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        :param pulumi.Input[_builtins.int] digits: The number of digits in the generated TOTP token. This value can either be 6 or 8
        :param pulumi.Input[_builtins.str] issuer: The name of the key's issuing organization.
        :param pulumi.Input[_builtins.int] key_size: Specifies the size in bytes of the generated key.
        :param pulumi.Input[_builtins.int] max_validation_attempts: The maximum number of consecutive failed validation attempts allowed.
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.int] period: The length of time in seconds used to generate a counter for the TOTP token calculation.
        :param pulumi.Input[_builtins.int] qr_size: The pixel size of the generated square QR code.
        :param pulumi.Input[_builtins.int] skew: The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        if algorithm is not None:
            pulumi.set(__self__, "algorithm", algorithm)
        if digits is not None:
            pulumi.set(__self__, "digits", digits)
        if issuer is not None:
            pulumi.set(__self__, "issuer", issuer)
        if key_size is not None:
            pulumi.set(__self__, "key_size", key_size)
        if max_validation_attempts is not None:
            pulumi.set(__self__, "max_validation_attempts", max_validation_attempts)
        if method_id is not None:
            pulumi.set(__self__, "method_id", method_id)
        if mount_accessor is not None:
            pulumi.set(__self__, "mount_accessor", mount_accessor)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if namespace_id is not None:
            pulumi.set(__self__, "namespace_id", namespace_id)
        if namespace_path is not None:
            pulumi.set(__self__, "namespace_path", namespace_path)
        if period is not None:
            pulumi.set(__self__, "period", period)
        if qr_size is not None:
            pulumi.set(__self__, "qr_size", qr_size)
        if skew is not None:
            pulumi.set(__self__, "skew", skew)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if uuid is not None:
            pulumi.set(__self__, "uuid", uuid)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        """
        return pulumi.get(self, "algorithm")

    @algorithm.setter
    def algorithm(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "algorithm", value)

    @_builtins.property
    @pulumi.getter
    def digits(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The number of digits in the generated TOTP token. This value can either be 6 or 8
        """
        return pulumi.get(self, "digits")

    @digits.setter
    def digits(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "digits", value)

    @_builtins.property
    @pulumi.getter
    def issuer(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The name of the key's issuing organization.
        """
        return pulumi.get(self, "issuer")

    @issuer.setter
    def issuer(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "issuer", value)

    @_builtins.property
    @pulumi.getter(name="keySize")
    def key_size(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        Specifies the size in bytes of the generated key.
        """
        return pulumi.get(self, "key_size")

    @key_size.setter
    def key_size(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "key_size", value)

    @_builtins.property
    @pulumi.getter(name="maxValidationAttempts")
    def max_validation_attempts(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The maximum number of consecutive failed validation attempts allowed.
        """
        return pulumi.get(self, "max_validation_attempts")

    @max_validation_attempts.setter
    def max_validation_attempts(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "max_validation_attempts", value)

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @method_id.setter
    def method_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "method_id", value)

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @mount_accessor.setter
    def mount_accessor(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "mount_accessor", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @namespace_id.setter
    def namespace_id(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_id", value)

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @namespace_path.setter
    def namespace_path(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace_path", value)

    @_builtins.property
    @pulumi.getter
    def period(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The length of time in seconds used to generate a counter for the TOTP token calculation.
        """
        return pulumi.get(self, "period")

    @period.setter
    def period(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "period", value)

    @_builtins.property
    @pulumi.getter(name="qrSize")
    def qr_size(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The pixel size of the generated square QR code.
        """
        return pulumi.get(self, "qr_size")

    @qr_size.setter
    def qr_size(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "qr_size", value)

    @_builtins.property
    @pulumi.getter
    def skew(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        """
        return pulumi.get(self, "skew")

    @skew.setter
    def skew(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "skew", value)

    @_builtins.property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "type", value)

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

    @uuid.setter
    def uuid(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "uuid", value)


@pulumi.type_token("vault:identity/mfaTotp:MfaTotp")
class MfaTotp(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 digits: Optional[pulumi.Input[_builtins.int]] = None,
                 issuer: Optional[pulumi.Input[_builtins.str]] = None,
                 key_size: Optional[pulumi.Input[_builtins.int]] = None,
                 max_validation_attempts: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 period: Optional[pulumi.Input[_builtins.int]] = None,
                 qr_size: Optional[pulumi.Input[_builtins.int]] = None,
                 skew: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        """
        Resource for configuring the totp MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaTotp("example", issuer="issuer1")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaTotp:MfaTotp example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] algorithm: Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        :param pulumi.Input[_builtins.int] digits: The number of digits in the generated TOTP token. This value can either be 6 or 8
        :param pulumi.Input[_builtins.str] issuer: The name of the key's issuing organization.
        :param pulumi.Input[_builtins.int] key_size: Specifies the size in bytes of the generated key.
        :param pulumi.Input[_builtins.int] max_validation_attempts: The maximum number of consecutive failed validation attempts allowed.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.int] period: The length of time in seconds used to generate a counter for the TOTP token calculation.
        :param pulumi.Input[_builtins.int] qr_size: The pixel size of the generated square QR code.
        :param pulumi.Input[_builtins.int] skew: The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: MfaTotpArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Resource for configuring the totp MFA method.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.identity.MfaTotp("example", issuer="issuer1")
        ```

        ## Import

        Resource can be imported using its `uuid` field, e.g.

        ```sh
        $ pulumi import vault:identity/mfaTotp:MfaTotp example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
        ```

        :param str resource_name: The name of the resource.
        :param MfaTotpArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(MfaTotpArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 algorithm: Optional[pulumi.Input[_builtins.str]] = None,
                 digits: Optional[pulumi.Input[_builtins.int]] = None,
                 issuer: Optional[pulumi.Input[_builtins.str]] = None,
                 key_size: Optional[pulumi.Input[_builtins.int]] = None,
                 max_validation_attempts: Optional[pulumi.Input[_builtins.int]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 period: Optional[pulumi.Input[_builtins.int]] = None,
                 qr_size: Optional[pulumi.Input[_builtins.int]] = None,
                 skew: Optional[pulumi.Input[_builtins.int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = MfaTotpArgs.__new__(MfaTotpArgs)

            __props__.__dict__["algorithm"] = algorithm
            __props__.__dict__["digits"] = digits
            if issuer is None and not opts.urn:
                raise TypeError("Missing required property 'issuer'")
            __props__.__dict__["issuer"] = issuer
            __props__.__dict__["key_size"] = key_size
            __props__.__dict__["max_validation_attempts"] = max_validation_attempts
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["period"] = period
            __props__.__dict__["qr_size"] = qr_size
            __props__.__dict__["skew"] = skew
            __props__.__dict__["method_id"] = None
            __props__.__dict__["mount_accessor"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["namespace_id"] = None
            __props__.__dict__["namespace_path"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["uuid"] = None
        super(MfaTotp, __self__).__init__(
            'vault:identity/mfaTotp:MfaTotp',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            algorithm: Optional[pulumi.Input[_builtins.str]] = None,
            digits: Optional[pulumi.Input[_builtins.int]] = None,
            issuer: Optional[pulumi.Input[_builtins.str]] = None,
            key_size: Optional[pulumi.Input[_builtins.int]] = None,
            max_validation_attempts: Optional[pulumi.Input[_builtins.int]] = None,
            method_id: Optional[pulumi.Input[_builtins.str]] = None,
            mount_accessor: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_id: Optional[pulumi.Input[_builtins.str]] = None,
            namespace_path: Optional[pulumi.Input[_builtins.str]] = None,
            period: Optional[pulumi.Input[_builtins.int]] = None,
            qr_size: Optional[pulumi.Input[_builtins.int]] = None,
            skew: Optional[pulumi.Input[_builtins.int]] = None,
            type: Optional[pulumi.Input[_builtins.str]] = None,
            uuid: Optional[pulumi.Input[_builtins.str]] = None) -> 'MfaTotp':
        """
        Get an existing MfaTotp resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] algorithm: Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        :param pulumi.Input[_builtins.int] digits: The number of digits in the generated TOTP token. This value can either be 6 or 8
        :param pulumi.Input[_builtins.str] issuer: The name of the key's issuing organization.
        :param pulumi.Input[_builtins.int] key_size: Specifies the size in bytes of the generated key.
        :param pulumi.Input[_builtins.int] max_validation_attempts: The maximum number of consecutive failed validation attempts allowed.
        :param pulumi.Input[_builtins.str] method_id: Method ID.
        :param pulumi.Input[_builtins.str] mount_accessor: Mount accessor.
        :param pulumi.Input[_builtins.str] name: Method name.
        :param pulumi.Input[_builtins.str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[_builtins.str] namespace_id: Method's namespace ID.
        :param pulumi.Input[_builtins.str] namespace_path: Method's namespace path.
        :param pulumi.Input[_builtins.int] period: The length of time in seconds used to generate a counter for the TOTP token calculation.
        :param pulumi.Input[_builtins.int] qr_size: The pixel size of the generated square QR code.
        :param pulumi.Input[_builtins.int] skew: The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        :param pulumi.Input[_builtins.str] type: MFA type.
        :param pulumi.Input[_builtins.str] uuid: Resource UUID.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _MfaTotpState.__new__(_MfaTotpState)

        __props__.__dict__["algorithm"] = algorithm
        __props__.__dict__["digits"] = digits
        __props__.__dict__["issuer"] = issuer
        __props__.__dict__["key_size"] = key_size
        __props__.__dict__["max_validation_attempts"] = max_validation_attempts
        __props__.__dict__["method_id"] = method_id
        __props__.__dict__["mount_accessor"] = mount_accessor
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["namespace_id"] = namespace_id
        __props__.__dict__["namespace_path"] = namespace_path
        __props__.__dict__["period"] = period
        __props__.__dict__["qr_size"] = qr_size
        __props__.__dict__["skew"] = skew
        __props__.__dict__["type"] = type
        __props__.__dict__["uuid"] = uuid
        return MfaTotp(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def algorithm(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
        """
        return pulumi.get(self, "algorithm")

    @_builtins.property
    @pulumi.getter
    def digits(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        The number of digits in the generated TOTP token. This value can either be 6 or 8
        """
        return pulumi.get(self, "digits")

    @_builtins.property
    @pulumi.getter
    def issuer(self) -> pulumi.Output[_builtins.str]:
        """
        The name of the key's issuing organization.
        """
        return pulumi.get(self, "issuer")

    @_builtins.property
    @pulumi.getter(name="keySize")
    def key_size(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        Specifies the size in bytes of the generated key.
        """
        return pulumi.get(self, "key_size")

    @_builtins.property
    @pulumi.getter(name="maxValidationAttempts")
    def max_validation_attempts(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        The maximum number of consecutive failed validation attempts allowed.
        """
        return pulumi.get(self, "max_validation_attempts")

    @_builtins.property
    @pulumi.getter(name="methodId")
    def method_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method ID.
        """
        return pulumi.get(self, "method_id")

    @_builtins.property
    @pulumi.getter(name="mountAccessor")
    def mount_accessor(self) -> pulumi.Output[_builtins.str]:
        """
        Mount accessor.
        """
        return pulumi.get(self, "mount_accessor")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        Method name.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter(name="namespaceId")
    def namespace_id(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace ID.
        """
        return pulumi.get(self, "namespace_id")

    @_builtins.property
    @pulumi.getter(name="namespacePath")
    def namespace_path(self) -> pulumi.Output[_builtins.str]:
        """
        Method's namespace path.
        """
        return pulumi.get(self, "namespace_path")

    @_builtins.property
    @pulumi.getter
    def period(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        The length of time in seconds used to generate a counter for the TOTP token calculation.
        """
        return pulumi.get(self, "period")

    @_builtins.property
    @pulumi.getter(name="qrSize")
    def qr_size(self) -> pulumi.Output[_builtins.int]:
        """
        The pixel size of the generated square QR code.
        """
        return pulumi.get(self, "qr_size")

    @_builtins.property
    @pulumi.getter
    def skew(self) -> pulumi.Output[Optional[_builtins.int]]:
        """
        The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
        """
        return pulumi.get(self, "skew")

    @_builtins.property
    @pulumi.getter
    def type(self) -> pulumi.Output[_builtins.str]:
        """
        MFA type.
        """
        return pulumi.get(self, "type")

    @_builtins.property
    @pulumi.getter
    def uuid(self) -> pulumi.Output[_builtins.str]:
        """
        Resource UUID.
        """
        return pulumi.get(self, "uuid")

