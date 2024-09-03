# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TemplateArgs', 'Template']

@pulumi.input_type
class TemplateArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 alphabet: Optional[pulumi.Input[str]] = None,
                 decode_formats: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 encode_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Template resource.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] alphabet: The alphabet to use for this template. This is only used during FPE transformations.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] decode_formats: Optional mapping of name to regular expression template, used to customize
               the decoded output. (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] encode_format: The regular expression template used to format encoded values.
               (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] name: The name of the template.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] pattern: The pattern used for matching. Currently, only regular expression pattern is supported.
        :param pulumi.Input[str] type: The pattern type to use for match detection. Currently, only regex is supported.
        """
        pulumi.set(__self__, "path", path)
        if alphabet is not None:
            pulumi.set(__self__, "alphabet", alphabet)
        if decode_formats is not None:
            pulumi.set(__self__, "decode_formats", decode_formats)
        if encode_format is not None:
            pulumi.set(__self__, "encode_format", encode_format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if type is not None:
            pulumi.set(__self__, "type", type)

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
        The alphabet to use for this template. This is only used during FPE transformations.
        """
        return pulumi.get(self, "alphabet")

    @alphabet.setter
    def alphabet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alphabet", value)

    @property
    @pulumi.getter(name="decodeFormats")
    def decode_formats(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional mapping of name to regular expression template, used to customize
        the decoded output. (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "decode_formats")

    @decode_formats.setter
    def decode_formats(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "decode_formats", value)

    @property
    @pulumi.getter(name="encodeFormat")
    def encode_format(self) -> Optional[pulumi.Input[str]]:
        """
        The regular expression template used to format encoded values.
        (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "encode_format")

    @encode_format.setter
    def encode_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encode_format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template.
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
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern used for matching. Currently, only regular expression pattern is supported.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern type to use for match detection. Currently, only regex is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _TemplateState:
    def __init__(__self__, *,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 decode_formats: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 encode_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Template resources.
        :param pulumi.Input[str] alphabet: The alphabet to use for this template. This is only used during FPE transformations.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] decode_formats: Optional mapping of name to regular expression template, used to customize
               the decoded output. (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] encode_format: The regular expression template used to format encoded values.
               (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] name: The name of the template.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] pattern: The pattern used for matching. Currently, only regular expression pattern is supported.
        :param pulumi.Input[str] type: The pattern type to use for match detection. Currently, only regex is supported.
        """
        if alphabet is not None:
            pulumi.set(__self__, "alphabet", alphabet)
        if decode_formats is not None:
            pulumi.set(__self__, "decode_formats", decode_formats)
        if encode_format is not None:
            pulumi.set(__self__, "encode_format", encode_format)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if pattern is not None:
            pulumi.set(__self__, "pattern", pattern)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def alphabet(self) -> Optional[pulumi.Input[str]]:
        """
        The alphabet to use for this template. This is only used during FPE transformations.
        """
        return pulumi.get(self, "alphabet")

    @alphabet.setter
    def alphabet(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alphabet", value)

    @property
    @pulumi.getter(name="decodeFormats")
    def decode_formats(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional mapping of name to regular expression template, used to customize
        the decoded output. (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "decode_formats")

    @decode_formats.setter
    def decode_formats(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "decode_formats", value)

    @property
    @pulumi.getter(name="encodeFormat")
    def encode_format(self) -> Optional[pulumi.Input[str]]:
        """
        The regular expression template used to format encoded values.
        (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "encode_format")

    @encode_format.setter
    def encode_format(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "encode_format", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template.
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

    @property
    @pulumi.getter
    def pattern(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern used for matching. Currently, only regular expression pattern is supported.
        """
        return pulumi.get(self, "pattern")

    @pattern.setter
    def pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "pattern", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The pattern type to use for match detection. Currently, only regex is supported.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Template(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 decode_formats: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 encode_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        This resource supports the `/transform/template/{name}` Vault endpoint.

        It creates or updates a template with the given name. If a template with the name does not exist,
        it will be created. If the template exists, it will be updated with the new attributes.

        > Requires *Vault Enterprise with the Advanced Data Protection Transform Module*.
        See [Transform Secrets Engine](https://www.vaultproject.io/docs/secrets/transform)
        for more information.

        ## Example Usage

        Please note that the `pattern` below holds a regex. The regex shown
        is identical to the one in our [Setup](https://www.vaultproject.io/docs/secrets/transform#setup)
        docs, `(\\d{4})-(\\d{4})-(\\d{4})-(\\d{4})`. However, due to HCL, the
        backslashes must be escaped to appear correctly in Vault. For further
        assistance escaping your own custom regex, see String Literals.

        ```python
        import pulumi
        import pulumi_vault as vault

        transform = vault.Mount("transform",
            path="transform",
            type="transform")
        numerics = vault.transform.Alphabet("numerics",
            path=transform.path,
            name="numerics",
            alphabet="0123456789")
        test = vault.transform.Template("test",
            path=numerics.path,
            name="ccn",
            type="regex",
            pattern="(\\\\d{4})[- ](\\\\d{4})[- ](\\\\d{4})[- ](\\\\d{4})",
            alphabet="numerics",
            encode_format="$1-$2-$3-$4",
            decode_formats={
                "last-four-digits": "$4",
            })
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: The alphabet to use for this template. This is only used during FPE transformations.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] decode_formats: Optional mapping of name to regular expression template, used to customize
               the decoded output. (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] encode_format: The regular expression template used to format encoded values.
               (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] name: The name of the template.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] pattern: The pattern used for matching. Currently, only regular expression pattern is supported.
        :param pulumi.Input[str] type: The pattern type to use for match detection. Currently, only regex is supported.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TemplateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        This resource supports the `/transform/template/{name}` Vault endpoint.

        It creates or updates a template with the given name. If a template with the name does not exist,
        it will be created. If the template exists, it will be updated with the new attributes.

        > Requires *Vault Enterprise with the Advanced Data Protection Transform Module*.
        See [Transform Secrets Engine](https://www.vaultproject.io/docs/secrets/transform)
        for more information.

        ## Example Usage

        Please note that the `pattern` below holds a regex. The regex shown
        is identical to the one in our [Setup](https://www.vaultproject.io/docs/secrets/transform#setup)
        docs, `(\\d{4})-(\\d{4})-(\\d{4})-(\\d{4})`. However, due to HCL, the
        backslashes must be escaped to appear correctly in Vault. For further
        assistance escaping your own custom regex, see String Literals.

        ```python
        import pulumi
        import pulumi_vault as vault

        transform = vault.Mount("transform",
            path="transform",
            type="transform")
        numerics = vault.transform.Alphabet("numerics",
            path=transform.path,
            name="numerics",
            alphabet="0123456789")
        test = vault.transform.Template("test",
            path=numerics.path,
            name="ccn",
            type="regex",
            pattern="(\\\\d{4})[- ](\\\\d{4})[- ](\\\\d{4})[- ](\\\\d{4})",
            alphabet="numerics",
            encode_format="$1-$2-$3-$4",
            decode_formats={
                "last-four-digits": "$4",
            })
        ```

        :param str resource_name: The name of the resource.
        :param TemplateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TemplateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 decode_formats: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 encode_format: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 pattern: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TemplateArgs.__new__(TemplateArgs)

            __props__.__dict__["alphabet"] = alphabet
            __props__.__dict__["decode_formats"] = decode_formats
            __props__.__dict__["encode_format"] = encode_format
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["pattern"] = pattern
            __props__.__dict__["type"] = type
        super(Template, __self__).__init__(
            'vault:transform/template:Template',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alphabet: Optional[pulumi.Input[str]] = None,
            decode_formats: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            encode_format: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            pattern: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Template':
        """
        Get an existing Template resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: The alphabet to use for this template. This is only used during FPE transformations.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] decode_formats: Optional mapping of name to regular expression template, used to customize
               the decoded output. (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] encode_format: The regular expression template used to format encoded values.
               (requires Vault Enterprise 1.9+)
        :param pulumi.Input[str] name: The name of the template.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] pattern: The pattern used for matching. Currently, only regular expression pattern is supported.
        :param pulumi.Input[str] type: The pattern type to use for match detection. Currently, only regex is supported.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TemplateState.__new__(_TemplateState)

        __props__.__dict__["alphabet"] = alphabet
        __props__.__dict__["decode_formats"] = decode_formats
        __props__.__dict__["encode_format"] = encode_format
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["pattern"] = pattern
        __props__.__dict__["type"] = type
        return Template(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def alphabet(self) -> pulumi.Output[Optional[str]]:
        """
        The alphabet to use for this template. This is only used during FPE transformations.
        """
        return pulumi.get(self, "alphabet")

    @property
    @pulumi.getter(name="decodeFormats")
    def decode_formats(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        Optional mapping of name to regular expression template, used to customize
        the decoded output. (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "decode_formats")

    @property
    @pulumi.getter(name="encodeFormat")
    def encode_format(self) -> pulumi.Output[Optional[str]]:
        """
        The regular expression template used to format encoded values.
        (requires Vault Enterprise 1.9+)
        """
        return pulumi.get(self, "encode_format")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the template.
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

    @property
    @pulumi.getter
    def pattern(self) -> pulumi.Output[Optional[str]]:
        """
        The pattern used for matching. Currently, only regular expression pattern is supported.
        """
        return pulumi.get(self, "pattern")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The pattern type to use for match detection. Currently, only regex is supported.
        """
        return pulumi.get(self, "type")

