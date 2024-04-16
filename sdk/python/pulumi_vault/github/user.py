# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['UserArgs', 'User']

@pulumi.input_type
class UserArgs:
    def __init__(__self__, *,
                 user: pulumi.Input[str],
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a User resource.
        :param pulumi.Input[str] user: GitHub user name.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens issued
               using this role.
        """
        pulumi.set(__self__, "user", user)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)

    @property
    @pulumi.getter
    def user(self) -> pulumi.Input[str]:
        """
        GitHub user name.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: pulumi.Input[str]):
        pulumi.set(self, "user", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

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
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings specifying the policies to be set on tokens issued
        using this role.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)


@pulumi.input_type
class _UserState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering User resources.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens issued
               using this role.
        :param pulumi.Input[str] user: GitHub user name.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if policies is not None:
            pulumi.set(__self__, "policies", policies)
        if user is not None:
            pulumi.set(__self__, "user", user)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

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
    def policies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        An array of strings specifying the policies to be set on tokens issued
        using this role.
        """
        return pulumi.get(self, "policies")

    @policies.setter
    def policies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "policies", value)

    @property
    @pulumi.getter
    def user(self) -> Optional[pulumi.Input[str]]:
        """
        GitHub user name.
        """
        return pulumi.get(self, "user")

    @user.setter
    def user(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "user", value)


class User(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Manages policy mappings for Github Users authenticated via Github. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/github/) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.github.AuthBackend("example", organization="myorg")
        tf_user = vault.github.User("tf_user",
            backend=example.id,
            user="john.doe",
            policies=[
                "developer",
                "read-only",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Github user mappings can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:github/user:User tf_user auth/github/map/users/john.doe
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens issued
               using this role.
        :param pulumi.Input[str] user: GitHub user name.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: UserArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages policy mappings for Github Users authenticated via Github. See the [Vault
        documentation](https://www.vaultproject.io/docs/auth/github/) for more
        information.

        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.github.AuthBackend("example", organization="myorg")
        tf_user = vault.github.User("tf_user",
            backend=example.id,
            user="john.doe",
            policies=[
                "developer",
                "read-only",
            ])
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Github user mappings can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:github/user:User tf_user auth/github/map/users/john.doe
        ```

        :param str resource_name: The name of the resource.
        :param UserArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(UserArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 user: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = UserArgs.__new__(UserArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["policies"] = policies
            if user is None and not opts.urn:
                raise TypeError("Missing required property 'user'")
            __props__.__dict__["user"] = user
        super(User, __self__).__init__(
            'vault:github/user:User',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            policies: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            user: Optional[pulumi.Input[str]] = None) -> 'User':
        """
        Get an existing User resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the github auth backend is mounted. Defaults to `github`
               if not specified.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] policies: An array of strings specifying the policies to be set on tokens issued
               using this role.
        :param pulumi.Input[str] user: GitHub user name.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _UserState.__new__(_UserState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["policies"] = policies
        __props__.__dict__["user"] = user
        return User(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        Path where the github auth backend is mounted. Defaults to `github`
        if not specified.
        """
        return pulumi.get(self, "backend")

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
    def policies(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        An array of strings specifying the policies to be set on tokens issued
        using this role.
        """
        return pulumi.get(self, "policies")

    @property
    @pulumi.getter
    def user(self) -> pulumi.Output[str]:
        """
        GitHub user name.
        """
        return pulumi.get(self, "user")

