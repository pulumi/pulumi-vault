# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union
from .. import _utilities, _tables

__all__ = ['Alphabet']


class Alphabet(pulumi.CustomResource):
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alphabet: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Create a Alphabet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] path: The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
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

            __props__['alphabet'] = alphabet
            __props__['name'] = name
            if path is None:
                raise TypeError("Missing required property 'path'")
            __props__['path'] = path
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
            path: Optional[pulumi.Input[str]] = None) -> 'Alphabet':
        """
        Get an existing Alphabet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alphabet: A string of characters that contains the alphabet set.
        :param pulumi.Input[str] name: The name of the alphabet.
        :param pulumi.Input[str] path: The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["alphabet"] = alphabet
        __props__["name"] = name
        __props__["path"] = path
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
    def path(self) -> pulumi.Output[str]:
        """
        The mount path for a back-end, for example, the path given in "$ vault auth enable -path=my-aws aws".
        """
        return pulumi.get(self, "path")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

