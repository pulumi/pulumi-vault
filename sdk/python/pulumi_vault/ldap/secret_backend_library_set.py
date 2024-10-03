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

__all__ = ['SecretBackendLibrarySetArgs', 'SecretBackendLibrarySet']

@pulumi.input_type
class SecretBackendLibrarySetArgs:
    def __init__(__self__, *,
                 service_account_names: pulumi.Input[Sequence[pulumi.Input[str]]],
                 disable_check_in_enforcement: Optional[pulumi.Input[bool]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        The set of arguments for constructing a SecretBackendLibrarySet resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_account_names: Specifies the slice of service accounts mapped to this set.
        :param pulumi.Input[bool] disable_check_in_enforcement: Disable enforcing that service
               accounts must be checked in by the entity or client token that checked them
               out. Defaults to false.
        :param pulumi.Input[int] max_ttl: The maximum password time-to-live in seconds. Defaults
               to the configuration max_ttl if not provided.
        :param pulumi.Input[str] mount: The path where the LDAP secrets backend is mounted.
        :param pulumi.Input[str] name: The name to identify this set of service accounts.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[int] ttl: The password time-to-live in seconds. Defaults to the configuration
               ttl if not provided.
        """
        pulumi.set(__self__, "service_account_names", service_account_names)
        if disable_check_in_enforcement is not None:
            pulumi.set(__self__, "disable_check_in_enforcement", disable_check_in_enforcement)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="serviceAccountNames")
    def service_account_names(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        Specifies the slice of service accounts mapped to this set.
        """
        return pulumi.get(self, "service_account_names")

    @service_account_names.setter
    def service_account_names(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "service_account_names", value)

    @property
    @pulumi.getter(name="disableCheckInEnforcement")
    def disable_check_in_enforcement(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable enforcing that service
        accounts must be checked in by the entity or client token that checked them
        out. Defaults to false.
        """
        return pulumi.get(self, "disable_check_in_enforcement")

    @disable_check_in_enforcement.setter
    def disable_check_in_enforcement(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_check_in_enforcement", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum password time-to-live in seconds. Defaults
        to the configuration max_ttl if not provided.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_ttl", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        The path where the LDAP secrets backend is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to identify this set of service accounts.
        Must be unique within the backend.
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
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The password time-to-live in seconds. Defaults to the configuration
        ttl if not provided.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


@pulumi.input_type
class _SecretBackendLibrarySetState:
    def __init__(__self__, *,
                 disable_check_in_enforcement: Optional[pulumi.Input[bool]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None):
        """
        Input properties used for looking up and filtering SecretBackendLibrarySet resources.
        :param pulumi.Input[bool] disable_check_in_enforcement: Disable enforcing that service
               accounts must be checked in by the entity or client token that checked them
               out. Defaults to false.
        :param pulumi.Input[int] max_ttl: The maximum password time-to-live in seconds. Defaults
               to the configuration max_ttl if not provided.
        :param pulumi.Input[str] mount: The path where the LDAP secrets backend is mounted.
        :param pulumi.Input[str] name: The name to identify this set of service accounts.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_account_names: Specifies the slice of service accounts mapped to this set.
        :param pulumi.Input[int] ttl: The password time-to-live in seconds. Defaults to the configuration
               ttl if not provided.
        """
        if disable_check_in_enforcement is not None:
            pulumi.set(__self__, "disable_check_in_enforcement", disable_check_in_enforcement)
        if max_ttl is not None:
            pulumi.set(__self__, "max_ttl", max_ttl)
        if mount is not None:
            pulumi.set(__self__, "mount", mount)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if service_account_names is not None:
            pulumi.set(__self__, "service_account_names", service_account_names)
        if ttl is not None:
            pulumi.set(__self__, "ttl", ttl)

    @property
    @pulumi.getter(name="disableCheckInEnforcement")
    def disable_check_in_enforcement(self) -> Optional[pulumi.Input[bool]]:
        """
        Disable enforcing that service
        accounts must be checked in by the entity or client token that checked them
        out. Defaults to false.
        """
        return pulumi.get(self, "disable_check_in_enforcement")

    @disable_check_in_enforcement.setter
    def disable_check_in_enforcement(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "disable_check_in_enforcement", value)

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum password time-to-live in seconds. Defaults
        to the configuration max_ttl if not provided.
        """
        return pulumi.get(self, "max_ttl")

    @max_ttl.setter
    def max_ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_ttl", value)

    @property
    @pulumi.getter
    def mount(self) -> Optional[pulumi.Input[str]]:
        """
        The path where the LDAP secrets backend is mounted.
        """
        return pulumi.get(self, "mount")

    @mount.setter
    def mount(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "mount", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to identify this set of service accounts.
        Must be unique within the backend.
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
    @pulumi.getter(name="serviceAccountNames")
    def service_account_names(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Specifies the slice of service accounts mapped to this set.
        """
        return pulumi.get(self, "service_account_names")

    @service_account_names.setter
    def service_account_names(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "service_account_names", value)

    @property
    @pulumi.getter
    def ttl(self) -> Optional[pulumi.Input[int]]:
        """
        The password time-to-live in seconds. Defaults to the configuration
        ttl if not provided.
        """
        return pulumi.get(self, "ttl")

    @ttl.setter
    def ttl(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "ttl", value)


class SecretBackendLibrarySet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_check_in_enforcement: Optional[pulumi.Input[bool]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.ldap.SecretBackend("config",
            path="ldap",
            binddn="CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
            bindpass="SuperSecretPassw0rd",
            url="ldaps://localhost",
            insecure_tls=True,
            userdn="CN=Users,DC=corp,DC=example,DC=net")
        qa = vault.ldap.SecretBackendLibrarySet("qa",
            mount=config.path,
            name="qa",
            service_account_names=[
                "Bob",
                "Mary",
            ],
            ttl=60,
            disable_check_in_enforcement=True,
            max_ttl=120)
        ```

        ## Import

        LDAP secret backend libraries can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet qa ldap/library/bob
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_check_in_enforcement: Disable enforcing that service
               accounts must be checked in by the entity or client token that checked them
               out. Defaults to false.
        :param pulumi.Input[int] max_ttl: The maximum password time-to-live in seconds. Defaults
               to the configuration max_ttl if not provided.
        :param pulumi.Input[str] mount: The path where the LDAP secrets backend is mounted.
        :param pulumi.Input[str] name: The name to identify this set of service accounts.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_account_names: Specifies the slice of service accounts mapped to this set.
        :param pulumi.Input[int] ttl: The password time-to-live in seconds. Defaults to the configuration
               ttl if not provided.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendLibrarySetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        config = vault.ldap.SecretBackend("config",
            path="ldap",
            binddn="CN=Administrator,CN=Users,DC=corp,DC=example,DC=net",
            bindpass="SuperSecretPassw0rd",
            url="ldaps://localhost",
            insecure_tls=True,
            userdn="CN=Users,DC=corp,DC=example,DC=net")
        qa = vault.ldap.SecretBackendLibrarySet("qa",
            mount=config.path,
            name="qa",
            service_account_names=[
                "Bob",
                "Mary",
            ],
            ttl=60,
            disable_check_in_enforcement=True,
            max_ttl=120)
        ```

        ## Import

        LDAP secret backend libraries can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet qa ldap/library/bob
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendLibrarySetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendLibrarySetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 disable_check_in_enforcement: Optional[pulumi.Input[bool]] = None,
                 max_ttl: Optional[pulumi.Input[int]] = None,
                 mount: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 service_account_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 ttl: Optional[pulumi.Input[int]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendLibrarySetArgs.__new__(SecretBackendLibrarySetArgs)

            __props__.__dict__["disable_check_in_enforcement"] = disable_check_in_enforcement
            __props__.__dict__["max_ttl"] = max_ttl
            __props__.__dict__["mount"] = mount
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if service_account_names is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_names'")
            __props__.__dict__["service_account_names"] = service_account_names
            __props__.__dict__["ttl"] = ttl
        super(SecretBackendLibrarySet, __self__).__init__(
            'vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            disable_check_in_enforcement: Optional[pulumi.Input[bool]] = None,
            max_ttl: Optional[pulumi.Input[int]] = None,
            mount: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            service_account_names: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            ttl: Optional[pulumi.Input[int]] = None) -> 'SecretBackendLibrarySet':
        """
        Get an existing SecretBackendLibrarySet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] disable_check_in_enforcement: Disable enforcing that service
               accounts must be checked in by the entity or client token that checked them
               out. Defaults to false.
        :param pulumi.Input[int] max_ttl: The maximum password time-to-live in seconds. Defaults
               to the configuration max_ttl if not provided.
        :param pulumi.Input[str] mount: The path where the LDAP secrets backend is mounted.
        :param pulumi.Input[str] name: The name to identify this set of service accounts.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] service_account_names: Specifies the slice of service accounts mapped to this set.
        :param pulumi.Input[int] ttl: The password time-to-live in seconds. Defaults to the configuration
               ttl if not provided.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendLibrarySetState.__new__(_SecretBackendLibrarySetState)

        __props__.__dict__["disable_check_in_enforcement"] = disable_check_in_enforcement
        __props__.__dict__["max_ttl"] = max_ttl
        __props__.__dict__["mount"] = mount
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["service_account_names"] = service_account_names
        __props__.__dict__["ttl"] = ttl
        return SecretBackendLibrarySet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="disableCheckInEnforcement")
    def disable_check_in_enforcement(self) -> pulumi.Output[Optional[bool]]:
        """
        Disable enforcing that service
        accounts must be checked in by the entity or client token that checked them
        out. Defaults to false.
        """
        return pulumi.get(self, "disable_check_in_enforcement")

    @property
    @pulumi.getter(name="maxTtl")
    def max_ttl(self) -> pulumi.Output[int]:
        """
        The maximum password time-to-live in seconds. Defaults
        to the configuration max_ttl if not provided.
        """
        return pulumi.get(self, "max_ttl")

    @property
    @pulumi.getter
    def mount(self) -> pulumi.Output[Optional[str]]:
        """
        The path where the LDAP secrets backend is mounted.
        """
        return pulumi.get(self, "mount")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to identify this set of service accounts.
        Must be unique within the backend.
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
    @pulumi.getter(name="serviceAccountNames")
    def service_account_names(self) -> pulumi.Output[Sequence[str]]:
        """
        Specifies the slice of service accounts mapped to this set.
        """
        return pulumi.get(self, "service_account_names")

    @property
    @pulumi.getter
    def ttl(self) -> pulumi.Output[int]:
        """
        The password time-to-live in seconds. Defaults to the configuration
        ttl if not provided.
        """
        return pulumi.get(self, "ttl")

