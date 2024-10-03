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

__all__ = ['AlphabetArgs', 'Alphabet']

@pulumi.input_type
class AlphabetArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 alphabet: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Alphabet resource.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "path", path)
        if alphabet is not None:
            pulumi.set(__self__, "alphabet", alphabet)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[str]:
        """
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[str]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def alphabet(self) -> Optional[pulumi.Input[str]]:
        """
        A string of characters that contains the alphabet set.
        """
        return pulumi.get(self, "alphabet")

    @alphabet.setter
    def alphabet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alphabet", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alphabet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
class _AlphabetState:
    def __init__(__self__, *,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Alphabet resources.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        """
        if alphabet is not None:
            pulumi.set(__self__, "alphabet", alphabet)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)

    @property
    @pulumi.getter
    def alphabet(self) -> Optional[pulumi.Input[str]]:
        """
        A string of characters that contains the alphabet set.
        """
        return pulumi.get(self, "alphabet")

    @alphabet.setter
    def alphabet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alphabet", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the alphabet.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

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
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)


class Alphabet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports the "/transform/alphabet/{name}" Vault endpoint.

        It queries an existing alphabet by the given name.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mount_transform = vault.Mount("mount_transform",
            path="transform",
            type="transform")
        test = vault.transform.Alphabet("test",
            path=mount_transform.path,
            name="numerics",
            alphabet="0123456789")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AlphabetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports the "/transform/alphabet/{name}" Vault endpoint.

        It queries an existing alphabet by the given name.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        mount_transform = vault.Mount("mount_transform",
            path="transform",
            type="transform")
        test = vault.transform.Alphabet("test",
            path=mount_transform.path,
            name="numerics",
            alphabet="0123456789")
        ```

        :param str resource_name: The name of the resource.
        :param AlphabetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AlphabetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AlphabetArgs.__new__(AlphabetArgs)

            __props__.__dict__["alphabet"] = alphabet
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
        super(Alphabet, __self__).__init__(
            'vault:transform/alphabet:Alphabet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alphabet: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None) -> 'Alphabet':
        """
        Get an existing Alphabet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AlphabetState.__new__(_AlphabetState)

        __props__.__dict__["alphabet"] = alphabet
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        return Alphabet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alphabet(self) -> pulumi.Output[Optional[str]]:
        """
        A string of characters that contains the alphabet set.
        """
        return pulumi.get(self, "alphabet")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the alphabet.
        """
        return pulumi.get(self, "name")

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
    def path(self) -> pulumi.Output[str]:
        """
        Path to where the back-end is mounted within Vault.
        """
        return pulumi.get(self, "path")

