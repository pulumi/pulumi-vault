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
from . import _utilities

__all__ = ['EgpPolicyArgs', 'EgpPolicy']

@pulumi.input_type
class EgpPolicyArgs:
    def __init__(__self__, *,
                 enforcement_level: pulumi.Input[str],
                 paths: pulumi.Input[Sequence[pulumi.Input[str]]],
                 policy: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a EgpPolicy resource.
        :param pulumi.Input[str] enforcement_level: Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to which the policy will be applied to
        :param pulumi.Input[str] policy: String containing a Sentinel policy
        :param pulumi.Input[str] name: The name of the policy
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "enforcement_level", enforcement_level)
        pulumi.set(__self__, "paths", paths)
        pulumi.set(__self__, "policy", policy)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="enforcementLevel")
    def enforcement_level(self) -> pulumi.Input[str]:
        """
        Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        """
        return pulumi.get(self, "enforcement_level")

    @enforcement_level.setter
    def enforcement_level(self, value: pulumi.Input[str]):
        pulumi.set(self, "enforcement_level", value)

    @property
    @pulumi.getter
    def paths(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        List of paths to which the policy will be applied to
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Input[str]:
        """
        String containing a Sentinel policy
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: pulumi.Input[str]):
        pulumi.set(self, "policy", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy
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
class _EgpPolicyState:
    def __init__(__self__, *,
                 enforcement_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering EgpPolicy resources.
        :param pulumi.Input[str] enforcement_level: Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        :param pulumi.Input[str] name: The name of the policy
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to which the policy will be applied to
        :param pulumi.Input[str] policy: String containing a Sentinel policy
        """
        if enforcement_level is not None:
            pulumi.set(__self__, "enforcement_level", enforcement_level)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if paths is not None:
            pulumi.set(__self__, "paths", paths)
        if policy is not None:
            pulumi.set(__self__, "policy", policy)

    @property
    @pulumi.getter(name="enforcementLevel")
    def enforcement_level(self) -> Optional[pulumi.Input[str]]:
        """
        Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        """
        return pulumi.get(self, "enforcement_level")

    @enforcement_level.setter
    def enforcement_level(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "enforcement_level", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the policy
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
    def paths(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of paths to which the policy will be applied to
        """
        return pulumi.get(self, "paths")

    @paths.setter
    def paths(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "paths", value)

    @property
    @pulumi.getter
    def policy(self) -> Optional[pulumi.Input[str]]:
        """
        String containing a Sentinel policy
        """
        return pulumi.get(self, "policy")

    @policy.setter
    def policy(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "policy", value)


class EgpPolicy(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforcement_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        allow_all = vault.EgpPolicy("allow-all",
            name="allow-all",
            paths=["*"],
            enforcement_level="soft-mandatory",
            policy=\"\"\"main = rule {
          true
        }
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enforcement_level: Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        :param pulumi.Input[str] name: The name of the policy
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to which the policy will be applied to
        :param pulumi.Input[str] policy: String containing a Sentinel policy
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: EgpPolicyArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a resource to manage Endpoint Governing Policy (EGP) via [Sentinel](https://www.vaultproject.io/docs/enterprise/sentinel/index.html).

        **Note** this feature is available only with Vault Enterprise.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        allow_all = vault.EgpPolicy("allow-all",
            name="allow-all",
            paths=["*"],
            enforcement_level="soft-mandatory",
            policy=\"\"\"main = rule {
          true
        }
        \"\"\")
        ```

        :param str resource_name: The name of the resource.
        :param EgpPolicyArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(EgpPolicyArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 enforcement_level: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 policy: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = EgpPolicyArgs.__new__(EgpPolicyArgs)

            if enforcement_level is None and not opts.urn:
                raise TypeError("Missing required property 'enforcement_level'")
            __props__.__dict__["enforcement_level"] = enforcement_level
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if paths is None and not opts.urn:
                raise TypeError("Missing required property 'paths'")
            __props__.__dict__["paths"] = paths
            if policy is None and not opts.urn:
                raise TypeError("Missing required property 'policy'")
            __props__.__dict__["policy"] = policy
        super(EgpPolicy, __self__).__init__(
            'vault:index/egpPolicy:EgpPolicy',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            enforcement_level: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            paths: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            policy: Optional[pulumi.Input[str]] = None) -> 'EgpPolicy':
        """
        Get an existing EgpPolicy resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] enforcement_level: Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        :param pulumi.Input[str] name: The name of the policy
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] paths: List of paths to which the policy will be applied to
        :param pulumi.Input[str] policy: String containing a Sentinel policy
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _EgpPolicyState.__new__(_EgpPolicyState)

        __props__.__dict__["enforcement_level"] = enforcement_level
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["paths"] = paths
        __props__.__dict__["policy"] = policy
        return EgpPolicy(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="enforcementLevel")
    def enforcement_level(self) -> pulumi.Output[str]:
        """
        Enforcement level of Sentinel policy. Can be either `advisory` or `soft-mandatory` or `hard-mandatory`
        """
        return pulumi.get(self, "enforcement_level")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the policy
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
    def paths(self) -> pulumi.Output[Sequence[str]]:
        """
        List of paths to which the policy will be applied to
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def policy(self) -> pulumi.Output[str]:
        """
        String containing a Sentinel policy
        """
        return pulumi.get(self, "policy")

