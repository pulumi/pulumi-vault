# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['UiCustomMessageArgs', 'UiCustomMessage']

@pulumi.input_type
class UiCustomMessageArgs:
    def __init__(__self__, *,
                 message_base64: pulumi.Input[str],
                 start_time: pulumi.Input[str],
                 title: pulumi.Input[str],
                 authenticated: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 link: Optional[pulumi.Input['UiCustomMessageLinkArgs']] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a UiCustomMessage resource.
        :param pulumi.Input[str] message_base64: The base64-encoded content of the custom message
        :param pulumi.Input[str] start_time: The starting time of the active period of the custom message
        :param pulumi.Input[str] title: The title of the custom message
        :param pulumi.Input[bool] authenticated: A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        :param pulumi.Input[str] end_time: The ending time of the active period of the custom message. Can be omitted for non-expiring message
        :param pulumi.Input['UiCustomMessageLinkArgs'] link: A block containing a hyperlink associated with the custom message
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: A map containing additional options for the custom message
        :param pulumi.Input[str] type: The display type of custom message. Allowed values are banner and modal
        """
        pulumi.set(__self__, "message_base64", message_base64)
        pulumi.set(__self__, "start_time", start_time)
        pulumi.set(__self__, "title", title)
        if authenticated is not None:
            pulumi.set(__self__, "authenticated", authenticated)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if link is not None:
            pulumi.set(__self__, "link", link)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="messageBase64")
    def message_base64(self) -> pulumi.Input[str]:
        """
        The base64-encoded content of the custom message
        """
        return pulumi.get(self, "message_base64")

    @message_base64.setter
    def message_base64(self, value: pulumi.Input[str]):
        pulumi.set(self, "message_base64", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Input[str]:
        """
        The starting time of the active period of the custom message
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: pulumi.Input[str]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def title(self) -> pulumi.Input[str]:
        """
        The title of the custom message
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: pulumi.Input[str]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def authenticated(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        """
        return pulumi.get(self, "authenticated")

    @authenticated.setter
    def authenticated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "authenticated", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The ending time of the active period of the custom message. Can be omitted for non-expiring message
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def link(self) -> Optional[pulumi.Input['UiCustomMessageLinkArgs']]:
        """
        A block containing a hyperlink associated with the custom message
        """
        return pulumi.get(self, "link")

    @link.setter
    def link(self, value: Optional[pulumi.Input['UiCustomMessageLinkArgs']]):
        pulumi.set(self, "link", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map containing additional options for the custom message
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The display type of custom message. Allowed values are banner and modal
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class _UiCustomMessageState:
    def __init__(__self__, *,
                 authenticated: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 link: Optional[pulumi.Input['UiCustomMessageLinkArgs']] = None,
                 message_base64: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering UiCustomMessage resources.
        :param pulumi.Input[bool] authenticated: A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        :param pulumi.Input[str] end_time: The ending time of the active period of the custom message. Can be omitted for non-expiring message
        :param pulumi.Input['UiCustomMessageLinkArgs'] link: A block containing a hyperlink associated with the custom message
        :param pulumi.Input[str] message_base64: The base64-encoded content of the custom message
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: A map containing additional options for the custom message
        :param pulumi.Input[str] start_time: The starting time of the active period of the custom message
        :param pulumi.Input[str] title: The title of the custom message
        :param pulumi.Input[str] type: The display type of custom message. Allowed values are banner and modal
        """
        if authenticated is not None:
            pulumi.set(__self__, "authenticated", authenticated)
        if end_time is not None:
            pulumi.set(__self__, "end_time", end_time)
        if link is not None:
            pulumi.set(__self__, "link", link)
        if message_base64 is not None:
            pulumi.set(__self__, "message_base64", message_base64)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if options is not None:
            pulumi.set(__self__, "options", options)
        if start_time is not None:
            pulumi.set(__self__, "start_time", start_time)
        if title is not None:
            pulumi.set(__self__, "title", title)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def authenticated(self) -> Optional[pulumi.Input[bool]]:
        """
        A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        """
        return pulumi.get(self, "authenticated")

    @authenticated.setter
    def authenticated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "authenticated", value)

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> Optional[pulumi.Input[str]]:
        """
        The ending time of the active period of the custom message. Can be omitted for non-expiring message
        """
        return pulumi.get(self, "end_time")

    @end_time.setter
    def end_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "end_time", value)

    @property
    @pulumi.getter
    def link(self) -> Optional[pulumi.Input['UiCustomMessageLinkArgs']]:
        """
        A block containing a hyperlink associated with the custom message
        """
        return pulumi.get(self, "link")

    @link.setter
    def link(self, value: Optional[pulumi.Input['UiCustomMessageLinkArgs']]):
        pulumi.set(self, "link", value)

    @property
    @pulumi.getter(name="messageBase64")
    def message_base64(self) -> Optional[pulumi.Input[str]]:
        """
        The base64-encoded content of the custom message
        """
        return pulumi.get(self, "message_base64")

    @message_base64.setter
    def message_base64(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "message_base64", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def options(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map containing additional options for the custom message
        """
        return pulumi.get(self, "options")

    @options.setter
    def options(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "options", value)

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> Optional[pulumi.Input[str]]:
        """
        The starting time of the active period of the custom message
        """
        return pulumi.get(self, "start_time")

    @start_time.setter
    def start_time(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "start_time", value)

    @property
    @pulumi.getter
    def title(self) -> Optional[pulumi.Input[str]]:
        """
        The title of the custom message
        """
        return pulumi.get(self, "title")

    @title.setter
    def title(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "title", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The display type of custom message. Allowed values are banner and modal
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class UiCustomMessage(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticated: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 link: Optional[pulumi.Input[Union['UiCustomMessageLinkArgs', 'UiCustomMessageLinkArgsDict']]] = None,
                 message_base64: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a UiCustomMessage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authenticated: A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        :param pulumi.Input[str] end_time: The ending time of the active period of the custom message. Can be omitted for non-expiring message
        :param pulumi.Input[Union['UiCustomMessageLinkArgs', 'UiCustomMessageLinkArgsDict']] link: A block containing a hyperlink associated with the custom message
        :param pulumi.Input[str] message_base64: The base64-encoded content of the custom message
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: A map containing additional options for the custom message
        :param pulumi.Input[str] start_time: The starting time of the active period of the custom message
        :param pulumi.Input[str] title: The title of the custom message
        :param pulumi.Input[str] type: The display type of custom message. Allowed values are banner and modal
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UiCustomMessageArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a UiCustomMessage resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param UiCustomMessageArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UiCustomMessageArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 authenticated: Optional[pulumi.Input[bool]] = None,
                 end_time: Optional[pulumi.Input[str]] = None,
                 link: Optional[pulumi.Input[Union['UiCustomMessageLinkArgs', 'UiCustomMessageLinkArgsDict']]] = None,
                 message_base64: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 start_time: Optional[pulumi.Input[str]] = None,
                 title: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UiCustomMessageArgs.__new__(UiCustomMessageArgs)

            __props__.__dict__["authenticated"] = authenticated
            __props__.__dict__["end_time"] = end_time
            __props__.__dict__["link"] = link
            if message_base64 is None and not opts.urn:
                raise TypeError("Missing required property 'message_base64'")
            __props__.__dict__["message_base64"] = message_base64
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["options"] = options
            if start_time is None and not opts.urn:
                raise TypeError("Missing required property 'start_time'")
            __props__.__dict__["start_time"] = start_time
            if title is None and not opts.urn:
                raise TypeError("Missing required property 'title'")
            __props__.__dict__["title"] = title
            __props__.__dict__["type"] = type
        super(UiCustomMessage, __self__).__init__(
            'vault:config/uiCustomMessage:UiCustomMessage',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            authenticated: Optional[pulumi.Input[bool]] = None,
            end_time: Optional[pulumi.Input[str]] = None,
            link: Optional[pulumi.Input[Union['UiCustomMessageLinkArgs', 'UiCustomMessageLinkArgsDict']]] = None,
            message_base64: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            options: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            start_time: Optional[pulumi.Input[str]] = None,
            title: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'UiCustomMessage':
        """
        Get an existing UiCustomMessage resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] authenticated: A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        :param pulumi.Input[str] end_time: The ending time of the active period of the custom message. Can be omitted for non-expiring message
        :param pulumi.Input[Union['UiCustomMessageLinkArgs', 'UiCustomMessageLinkArgsDict']] link: A block containing a hyperlink associated with the custom message
        :param pulumi.Input[str] message_base64: The base64-encoded content of the custom message
        :param pulumi.Input[str] namespace: Target namespace. (requires Enterprise)
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] options: A map containing additional options for the custom message
        :param pulumi.Input[str] start_time: The starting time of the active period of the custom message
        :param pulumi.Input[str] title: The title of the custom message
        :param pulumi.Input[str] type: The display type of custom message. Allowed values are banner and modal
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UiCustomMessageState.__new__(_UiCustomMessageState)

        __props__.__dict__["authenticated"] = authenticated
        __props__.__dict__["end_time"] = end_time
        __props__.__dict__["link"] = link
        __props__.__dict__["message_base64"] = message_base64
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["options"] = options
        __props__.__dict__["start_time"] = start_time
        __props__.__dict__["title"] = title
        __props__.__dict__["type"] = type
        return UiCustomMessage(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def authenticated(self) -> pulumi.Output[Optional[bool]]:
        """
        A flag indicating whether the custom message is displayed pre-login (false) or post-login (true)
        """
        return pulumi.get(self, "authenticated")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[Optional[str]]:
        """
        The ending time of the active period of the custom message. Can be omitted for non-expiring message
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def link(self) -> pulumi.Output[Optional['outputs.UiCustomMessageLink']]:
        """
        A block containing a hyperlink associated with the custom message
        """
        return pulumi.get(self, "link")

    @property
    @pulumi.getter(name="messageBase64")
    def message_base64(self) -> pulumi.Output[str]:
        """
        The base64-encoded content of the custom message
        """
        return pulumi.get(self, "message_base64")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        Target namespace. (requires Enterprise)
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def options(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map containing additional options for the custom message
        """
        return pulumi.get(self, "options")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        The starting time of the active period of the custom message
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def title(self) -> pulumi.Output[str]:
        """
        The title of the custom message
        """
        return pulumi.get(self, "title")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The display type of custom message. Allowed values are banner and modal
        """
        return pulumi.get(self, "type")

