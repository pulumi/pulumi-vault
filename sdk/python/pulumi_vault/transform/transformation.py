# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Transformation']


class Transformation(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allowed_roles: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 masking_character: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 template: Optional[pulumi.Input[str]] = None,
                 templates: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 tweak_source: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Transformation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] allowed_roles: The set of roles allowed to perform this transformation.
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] path: The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            __props__['allowed_roles'] = allowed_roles
            __props__['masking_character'] = masking_character
            __props__['name'] = name
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
            __props__['template'] = template
            __props__['templates'] = templates
            __props__['tweak_source'] = tweak_source
            __props__['type'] = type
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
            masking_character: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
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
        :param pulumi.Input[str] masking_character: The character used to replace data when in masking mode
        :param pulumi.Input[str] name: The name of the transformation.
        :param pulumi.Input[str] path: The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
        :param pulumi.Input[str] template: The name of the template to use.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] templates: Templates configured for transformation.
        :param pulumi.Input[str] tweak_source: The source of where the tweak value comes from. Only valid when in FPE mode.
        :param pulumi.Input[str] type: The type of transformation to perform.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["allowed_roles"] = allowed_roles
        __props__["masking_character"] = masking_character
        __props__["name"] = name
        __props__["path"] = path
        __props__["template"] = template
        __props__["templates"] = templates
        __props__["tweak_source"] = tweak_source
        __props__["type"] = type
        return Transformation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowedRoles")
    def allowed_roles(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The set of roles allowed to perform this transformation.
        """
        return pulumi.get(self, "allowed_roles")

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
    def path(self) -> pulumi.Output[str]:
        """
        The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
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

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

