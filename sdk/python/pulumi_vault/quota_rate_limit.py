# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = ['QuotaRateLimitArgs', 'QuotaRateLimit']

@pulumi.input_type
class QuotaRateLimitArgs:
    def __init__(__self__, *,
                 rate: pulumi.Input[float],
                 block_interval: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a QuotaRateLimit resource.
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
        :param pulumi.Input[int] block_interval: If set, when a client reaches a rate limit threshold, the client will
               be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        :param pulumi.Input[int] interval: The duration in seconds to enforce rate limiting for.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[str] role: If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        pulumi.set(__self__, "rate", rate)
        if block_interval is not None:
            pulumi.set(__self__, "block_interval", block_interval)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Input[float]:
        """
        The maximum number of requests at any given second to be allowed by the quota
        rule. The `rate` must be positive.
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: pulumi.Input[float]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter(name="blockInterval")
    def block_interval(self) -> Optional[pulumi.Input[int]]:
        """
        If set, when a client reaches a rate limit threshold, the client will
        be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        """
        return pulumi.get(self, "block_interval")

    @block_interval.setter
    def block_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "block_interval", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The duration in seconds to enforce rate limiting for.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the rate limit quota
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
        The `namespace` is always relative to the provider's configured namespace.
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
        Path of the mount or namespace to apply the quota. A blank path configures a
        global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        Updating this field on an existing quota can have "moving" effects. For example, updating
        `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


@pulumi.input_type
class _QuotaRateLimitState:
    def __init__(__self__, *,
                 block_interval: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[float]] = None,
                 role: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering QuotaRateLimit resources.
        :param pulumi.Input[int] block_interval: If set, when a client reaches a rate limit threshold, the client will
               be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        :param pulumi.Input[int] interval: The duration in seconds to enforce rate limiting for.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
        :param pulumi.Input[str] role: If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        if block_interval is not None:
            pulumi.set(__self__, "block_interval", block_interval)
        if interval is not None:
            pulumi.set(__self__, "interval", interval)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if rate is not None:
            pulumi.set(__self__, "rate", rate)
        if role is not None:
            pulumi.set(__self__, "role", role)

    @property
    @pulumi.getter(name="blockInterval")
    def block_interval(self) -> Optional[pulumi.Input[int]]:
        """
        If set, when a client reaches a rate limit threshold, the client will
        be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        """
        return pulumi.get(self, "block_interval")

    @block_interval.setter
    def block_interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "block_interval", value)

    @property
    @pulumi.getter
    def interval(self) -> Optional[pulumi.Input[int]]:
        """
        The duration in seconds to enforce rate limiting for.
        """
        return pulumi.get(self, "interval")

    @interval.setter
    def interval(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "interval", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the rate limit quota
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
        The `namespace` is always relative to the provider's configured namespace.
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
        Path of the mount or namespace to apply the quota. A blank path configures a
        global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        Updating this field on an existing quota can have "moving" effects. For example, updating
        `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def rate(self) -> Optional[pulumi.Input[float]]:
        """
        The maximum number of requests at any given second to be allowed by the quota
        rule. The `rate` must be positive.
        """
        return pulumi.get(self, "rate")

    @rate.setter
    def rate(self, value: Optional[pulumi.Input[float]]):
        pulumi.set(self, "rate", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)


class QuotaRateLimit(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_interval: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[float]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
        A rate limit quota can be created at the root level or defined on a namespace or mount by
        specifying a path when creating the quota.

        See [Vault's Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        global_ = vault.QuotaRateLimit("global",
            name="global",
            path="",
            rate=100)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Rate limit quotas can be imported using their names

        ```sh
        $ pulumi import vault:index/quotaRateLimit:QuotaRateLimit global global
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] block_interval: If set, when a client reaches a rate limit threshold, the client will
               be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        :param pulumi.Input[int] interval: The duration in seconds to enforce rate limiting for.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
        :param pulumi.Input[str] role: If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: QuotaRateLimitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manage rate limit quotas which enforce API rate limiting using a token bucket algorithm.
        A rate limit quota can be created at the root level or defined on a namespace or mount by
        specifying a path when creating the quota.

        See [Vault's Documentation](https://www.vaultproject.io/docs/concepts/resource-quotas) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        global_ = vault.QuotaRateLimit("global",
            name="global",
            path="",
            rate=100)
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Rate limit quotas can be imported using their names

        ```sh
        $ pulumi import vault:index/quotaRateLimit:QuotaRateLimit global global
        ```

        :param str resource_name: The name of the resource.
        :param QuotaRateLimitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(QuotaRateLimitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 block_interval: Optional[pulumi.Input[int]] = None,
                 interval: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 rate: Optional[pulumi.Input[float]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = QuotaRateLimitArgs.__new__(QuotaRateLimitArgs)

            __props__.__dict__["block_interval"] = block_interval
            __props__.__dict__["interval"] = interval
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["path"] = path
            if rate is None and not opts.urn:
                raise TypeError("Missing required property 'rate'")
            __props__.__dict__["rate"] = rate
            __props__.__dict__["role"] = role
        super(QuotaRateLimit, __self__).__init__(
            'vault:index/quotaRateLimit:QuotaRateLimit',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            block_interval: Optional[pulumi.Input[int]] = None,
            interval: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            path: Optional[pulumi.Input[str]] = None,
            rate: Optional[pulumi.Input[float]] = None,
            role: Optional[pulumi.Input[str]] = None) -> 'QuotaRateLimit':
        """
        Get an existing QuotaRateLimit resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] block_interval: If set, when a client reaches a rate limit threshold, the client will
               be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        :param pulumi.Input[int] interval: The duration in seconds to enforce rate limiting for.
        :param pulumi.Input[str] name: Name of the rate limit quota
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured namespace.
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] path: Path of the mount or namespace to apply the quota. A blank path configures a
               global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
               `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
               Updating this field on an existing quota can have "moving" effects. For example, updating
               `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
               a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        :param pulumi.Input[float] rate: The maximum number of requests at any given second to be allowed by the quota
               rule. The `rate` must be positive.
        :param pulumi.Input[str] role: If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _QuotaRateLimitState.__new__(_QuotaRateLimitState)

        __props__.__dict__["block_interval"] = block_interval
        __props__.__dict__["interval"] = interval
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["path"] = path
        __props__.__dict__["rate"] = rate
        __props__.__dict__["role"] = role
        return QuotaRateLimit(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="blockInterval")
    def block_interval(self) -> pulumi.Output[Optional[int]]:
        """
        If set, when a client reaches a rate limit threshold, the client will
        be prohibited from any further requests until after the 'block_interval' in seconds has elapsed.
        """
        return pulumi.get(self, "block_interval")

    @property
    @pulumi.getter
    def interval(self) -> pulumi.Output[int]:
        """
        The duration in seconds to enforce rate limiting for.
        """
        return pulumi.get(self, "interval")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name of the rate limit quota
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured namespace.
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        Path of the mount or namespace to apply the quota. A blank path configures a
        global rate limit quota. For example `namespace1/` adds a quota to a full namespace,
        `namespace1/auth/userpass` adds a `quota` to `userpass` in `namespace1`.
        Updating this field on an existing quota can have "moving" effects. For example, updating
        `auth/userpass` to `namespace1/auth/userpass` moves this quota from being a global mount quota to
        a namespace specific mount quota. **Note, namespaces are supported in Enterprise only.**
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def rate(self) -> pulumi.Output[float]:
        """
        The maximum number of requests at any given second to be allowed by the quota
        rule. The `rate` must be positive.
        """
        return pulumi.get(self, "rate")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[Optional[str]]:
        """
        If set on a quota where `path` is set to an auth mount with a concept of roles (such as /auth/approle/), this will make the quota restrict login requests to that mount that are made with the specified role.
        """
        return pulumi.get(self, "role")

