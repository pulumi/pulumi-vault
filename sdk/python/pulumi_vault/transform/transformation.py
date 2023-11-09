# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['TransformationArgs', 'Transformation']

@pulumi.input_type
class TransformationArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[str],
                 allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_allowed: Optional[pulumi.Input[bool]] = None,
                 masking_character: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tweak_source: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Transformation resource.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_roles: The set of roles allowed to perform this transformation.
        :param pulumi.Input[bool] deletion_allowed: If true, this transform can be deleted.
               Otherwise, deletion is blocked while this value remains false. Default: `false`
               *Only supported on vault-1.12+*
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        pulumi.set(__self__, "path", path)
        if allowed_roles is not None:
            pulumi.set(__self__, "allowed_roles", allowed_roles)
        if deletion_allowed is not None:
            pulumi.set(__self__, "deletion_allowed", deletion_allowed)
        if masking_character is not None:
            pulumi.set(__self__, "masking_character", masking_character)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if templates is not None:
            pulumi.set(__self__, "templates", templates)
        if tweak_source is not None:
            pulumi.set(__self__, "tweak_source", tweak_source)
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
    @pulumi.getter(name="allowedRoles")
    def allowed_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of roles allowed to perform this transformation.
        """
        return pulumi.get(self, "allowed_roles")

    @allowed_roles.setter
    def allowed_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_roles", value)

    @property
    @pulumi.getter(name="deletionAllowed")
    def deletion_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, this transform can be deleted.
        Otherwise, deletion is blocked while this value remains false. Default: `false`
        *Only supported on vault-1.12+*
        """
        return pulumi.get(self, "deletion_allowed")

    @deletion_allowed.setter
    def deletion_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_allowed", value)

    @property
    @pulumi.getter(name="maskingCharacter")
    def masking_character(self) -> Optional[pulumi.Input[str]]:
        """
        The character used to replace data when in masking mode
        """
        return pulumi.get(self, "masking_character")

    @masking_character.setter
    def masking_character(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "masking_character", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transformation.
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
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template to use.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter
    def templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Templates configured for transformation.
        """
        return pulumi.get(self, "templates")

    @templates.setter
    def templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "templates", value)

    @property
    @pulumi.getter(name="tweakSource")
    def tweak_source(self) -> Optional[pulumi.Input[str]]:
        """
        The source of where the tweak value comes from. Only valid when in FPE mode.
        """
        return pulumi.get(self, "tweak_source")

    @tweak_source.setter
    def tweak_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tweak_source", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of transformation to perform.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _TransformationState:
    def __init__(__self__, *,
                 allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_allowed: Optional[pulumi.Input[bool]] = None,
                 masking_character: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tweak_source: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Transformation resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_roles: The set of roles allowed to perform this transformation.
        :param pulumi.Input[bool] deletion_allowed: If true, this transform can be deleted.
               Otherwise, deletion is blocked while this value remains false. Default: `false`
               *Only supported on vault-1.12+*
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        if allowed_roles is not None:
            pulumi.set(__self__, "allowed_roles", allowed_roles)
        if deletion_allowed is not None:
            pulumi.set(__self__, "deletion_allowed", deletion_allowed)
        if masking_character is not None:
            pulumi.set(__self__, "masking_character", masking_character)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if template is not None:
            pulumi.set(__self__, "template", template)
        if templates is not None:
            pulumi.set(__self__, "templates", templates)
        if tweak_source is not None:
            pulumi.set(__self__, "tweak_source", tweak_source)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="allowedRoles")
    def allowed_roles(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The set of roles allowed to perform this transformation.
        """
        return pulumi.get(self, "allowed_roles")

    @allowed_roles.setter
    def allowed_roles(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "allowed_roles", value)

    @property
    @pulumi.getter(name="deletionAllowed")
    def deletion_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        If true, this transform can be deleted.
        Otherwise, deletion is blocked while this value remains false. Default: `false`
        *Only supported on vault-1.12+*
        """
        return pulumi.get(self, "deletion_allowed")

    @deletion_allowed.setter
    def deletion_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "deletion_allowed", value)

    @property
    @pulumi.getter(name="maskingCharacter")
    def masking_character(self) -> Optional[pulumi.Input[str]]:
        """
        The character used to replace data when in masking mode
        """
        return pulumi.get(self, "masking_character")

    @masking_character.setter
    def masking_character(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "masking_character", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the transformation.
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
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
    def template(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the template to use.
        """
        return pulumi.get(self, "template")

    @template.setter
    def template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "template", value)

    @property
    @pulumi.getter
    def templates(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Templates configured for transformation.
        """
        return pulumi.get(self, "templates")

    @templates.setter
    def templates(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "templates", value)

    @property
    @pulumi.getter(name="tweakSource")
    def tweak_source(self) -> Optional[pulumi.Input[str]]:
        """
        The source of where the tweak value comes from. Only valid when in FPE mode.
        """
        return pulumi.get(self, "tweak_source")

    @tweak_source.setter
    def tweak_source(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tweak_source", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of transformation to perform.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Transformation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_allowed: Optional[pulumi.Input[bool]] = None,
                 masking_character: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tweak_source: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Transformation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_roles: The set of roles allowed to perform this transformation.
        :param pulumi.Input[bool] deletion_allowed: If true, this transform can be deleted.
               Otherwise, deletion is blocked while this value remains false. Default: `false`
               *Only supported on vault-1.12+*
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TransformationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Transformation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param TransformationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TransformationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 deletion_allowed: Optional[pulumi.Input[bool]] = None,
                 masking_character: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tweak_source: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TransformationArgs.__new__(TransformationArgs)

            __props__.__dict__["allowed_roles"] = allowed_roles
            __props__.__dict__["deletion_allowed"] = deletion_allowed
            __props__.__dict__["masking_character"] = masking_character
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            __props__.__dict__["template"] = template
            __props__.__dict__["templates"] = templates
            __props__.__dict__["tweak_source"] = tweak_source
            __props__.__dict__["type"] = type
        super(Transformation, __self__).__init__(
            'vault:transform/transformation:Transformation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            deletion_allowed: Optional[pulumi.Input[bool]] = None,
            masking_character: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            template: Optional[pulumi.Input[str]] = None,
            templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            tweak_source: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Transformation':
        """
        Get an existing Transformation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_roles: The set of roles allowed to perform this transformation.
        :param pulumi.Input[bool] deletion_allowed: If true, this transform can be deleted.
               Otherwise, deletion is blocked while this value remains false. Default: `false`
               *Only supported on vault-1.12+*
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path to where the back-end is mounted within Vault.
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _TransformationState.__new__(_TransformationState)

        __props__.__dict__["allowed_roles"] = allowed_roles
        __props__.__dict__["deletion_allowed"] = deletion_allowed
        __props__.__dict__["masking_character"] = masking_character
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["template"] = template
        __props__.__dict__["templates"] = templates
        __props__.__dict__["tweak_source"] = tweak_source
        __props__.__dict__["type"] = type
        return Transformation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedRoles")
    def allowed_roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of roles allowed to perform this transformation.
        """
        return pulumi.get(self, "allowed_roles")

    @property
    @pulumi.getter(name="deletionAllowed")
    def deletion_allowed(self) -> pulumi.Output[Optional[bool]]:
        """
        If true, this transform can be deleted.
        Otherwise, deletion is blocked while this value remains false. Default: `false`
        *Only supported on vault-1.12+*
        """
        return pulumi.get(self, "deletion_allowed")

    @property
    @pulumi.getter(name="maskingCharacter")
    def masking_character(self) -> pulumi.Output[Optional[str]]:
        """
        The character used to replace data when in masking mode
        """
        return pulumi.get(self, "masking_character")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the transformation.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
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
    def template(self) -> pulumi.Output[Optional[str]]:
        """
        The name of the template to use.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter
    def templates(self) -> pulumi.Output[Sequence[str]]:
        """
        Templates configured for transformation.
        """
        return pulumi.get(self, "templates")

    @property
    @pulumi.getter(name="tweakSource")
    def tweak_source(self) -> pulumi.Output[Optional[str]]:
        """
        The source of where the tweak value comes from. Only valid when in FPE mode.
        """
        return pulumi.get(self, "tweak_source")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of transformation to perform.
        """
        return pulumi.get(self, "type")

