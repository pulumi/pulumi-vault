# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins
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

__all__ = ['AuthBackendConfigIdentityArgs', 'AuthBackendConfigIdentity']

@pulumi.input_type
class AuthBackendConfigIdentityArgs:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_alias: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 iam_alias: Optional[pulumi.Input[builtins.str]] = None,
                 iam_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None):
        """
        The set of arguments for constructing a AuthBackendConfigIdentity resource.
        :param pulumi.Input[builtins.str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[builtins.str] ec2_alias: How to generate the identity alias when using the ec2 auth method. Valid choices are
               `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] ec2_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `ec2_alias`
        :param pulumi.Input[builtins.str] iam_alias: How to generate the identity alias when using the iam auth method. Valid choices are
               `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] iam_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `iam_alias`
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if ec2_alias is not None:
            pulumi.set(__self__, "ec2_alias", ec2_alias)
        if ec2_metadatas is not None:
            pulumi.set(__self__, "ec2_metadatas", ec2_metadatas)
        if iam_alias is not None:
            pulumi.set(__self__, "iam_alias", iam_alias)
        if iam_metadatas is not None:
            pulumi.set(__self__, "iam_metadatas", iam_metadatas)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="ec2Alias")
    def ec2_alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        How to generate the identity alias when using the ec2 auth method. Valid choices are
        `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        """
        return pulumi.get(self, "ec2_alias")

    @ec2_alias.setter
    def ec2_alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ec2_alias", value)

    @property
    @pulumi.getter(name="ec2Metadatas")
    def ec2_metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `ec2_alias`
        """
        return pulumi.get(self, "ec2_metadatas")

    @ec2_metadatas.setter
    def ec2_metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "ec2_metadatas", value)

    @property
    @pulumi.getter(name="iamAlias")
    def iam_alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        How to generate the identity alias when using the iam auth method. Valid choices are
        `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        """
        return pulumi.get(self, "iam_alias")

    @iam_alias.setter
    def iam_alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "iam_alias", value)

    @property
    @pulumi.getter(name="iamMetadatas")
    def iam_metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `iam_alias`
        """
        return pulumi.get(self, "iam_metadatas")

    @iam_metadatas.setter
    def iam_metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "iam_metadatas", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _AuthBackendConfigIdentityState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_alias: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 iam_alias: Optional[pulumi.Input[builtins.str]] = None,
                 iam_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None):
        """
        Input properties used for looking up and filtering AuthBackendConfigIdentity resources.
        :param pulumi.Input[builtins.str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[builtins.str] ec2_alias: How to generate the identity alias when using the ec2 auth method. Valid choices are
               `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] ec2_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `ec2_alias`
        :param pulumi.Input[builtins.str] iam_alias: How to generate the identity alias when using the iam auth method. Valid choices are
               `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] iam_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `iam_alias`
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if ec2_alias is not None:
            pulumi.set(__self__, "ec2_alias", ec2_alias)
        if ec2_metadatas is not None:
            pulumi.set(__self__, "ec2_metadatas", ec2_metadatas)
        if iam_alias is not None:
            pulumi.set(__self__, "iam_alias", iam_alias)
        if iam_metadatas is not None:
            pulumi.set(__self__, "iam_metadatas", iam_metadatas)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="ec2Alias")
    def ec2_alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        How to generate the identity alias when using the ec2 auth method. Valid choices are
        `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        """
        return pulumi.get(self, "ec2_alias")

    @ec2_alias.setter
    def ec2_alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "ec2_alias", value)

    @property
    @pulumi.getter(name="ec2Metadatas")
    def ec2_metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `ec2_alias`
        """
        return pulumi.get(self, "ec2_metadatas")

    @ec2_metadatas.setter
    def ec2_metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "ec2_metadatas", value)

    @property
    @pulumi.getter(name="iamAlias")
    def iam_alias(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        How to generate the identity alias when using the iam auth method. Valid choices are
        `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        """
        return pulumi.get(self, "iam_alias")

    @iam_alias.setter
    def iam_alias(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "iam_alias", value)

    @property
    @pulumi.getter(name="iamMetadatas")
    def iam_metadatas(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `iam_alias`
        """
        return pulumi.get(self, "iam_metadatas")

    @iam_metadatas.setter
    def iam_metadatas(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]]):
        pulumi.set(self, "iam_metadatas", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[builtins.str]]):
        pulumi.set(self, "namespace", value)


class AuthBackendConfigIdentity(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_alias: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 iam_alias: Optional[pulumi.Input[builtins.str]] = None,
                 iam_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        """
        Manages an AWS auth backend identity configuration in a Vault server. This configuration defines how Vault interacts
        with the identity store. See the [Vault documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        aws = vault.AuthBackend("aws", type="aws")
        example = vault.aws.AuthBackendConfigIdentity("example",
            backend=aws.path,
            iam_alias="full_arn",
            iam_metadatas=[
                "canonical_arn",
                "account_id",
            ])
        ```

        ## Import

        AWS auth backend identity config can be imported using `auth/`, the `backend` path, and `/config/identity` e.g.

        ```sh
        $ pulumi import vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity example auth/aws/config/identity
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[builtins.str] ec2_alias: How to generate the identity alias when using the ec2 auth method. Valid choices are
               `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] ec2_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `ec2_alias`
        :param pulumi.Input[builtins.str] iam_alias: How to generate the identity alias when using the iam auth method. Valid choices are
               `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] iam_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `iam_alias`
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AuthBackendConfigIdentityArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Manages an AWS auth backend identity configuration in a Vault server. This configuration defines how Vault interacts
        with the identity store. See the [Vault documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
        information.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        aws = vault.AuthBackend("aws", type="aws")
        example = vault.aws.AuthBackendConfigIdentity("example",
            backend=aws.path,
            iam_alias="full_arn",
            iam_metadatas=[
                "canonical_arn",
                "account_id",
            ])
        ```

        ## Import

        AWS auth backend identity config can be imported using `auth/`, the `backend` path, and `/config/identity` e.g.

        ```sh
        $ pulumi import vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity example auth/aws/config/identity
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendConfigIdentityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendConfigIdentityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_alias: Optional[pulumi.Input[builtins.str]] = None,
                 ec2_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 iam_alias: Optional[pulumi.Input[builtins.str]] = None,
                 iam_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
                 namespace: Optional[pulumi.Input[builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendConfigIdentityArgs.__new__(AuthBackendConfigIdentityArgs)

            __props__.__dict__["backend"] = backend
            __props__.__dict__["ec2_alias"] = ec2_alias
            __props__.__dict__["ec2_metadatas"] = ec2_metadatas
            __props__.__dict__["iam_alias"] = iam_alias
            __props__.__dict__["iam_metadatas"] = iam_metadatas
            __props__.__dict__["namespace"] = namespace
        super(AuthBackendConfigIdentity, __self__).__init__(
            'vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[builtins.str]] = None,
            ec2_alias: Optional[pulumi.Input[builtins.str]] = None,
            ec2_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            iam_alias: Optional[pulumi.Input[builtins.str]] = None,
            iam_metadatas: Optional[pulumi.Input[Sequence[pulumi.Input[builtins.str]]]] = None,
            namespace: Optional[pulumi.Input[builtins.str]] = None) -> 'AuthBackendConfigIdentity':
        """
        Get an existing AuthBackendConfigIdentity resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[builtins.str] backend: Unique name of the auth backend to configure.
        :param pulumi.Input[builtins.str] ec2_alias: How to generate the identity alias when using the ec2 auth method. Valid choices are
               `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] ec2_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `ec2_alias`
        :param pulumi.Input[builtins.str] iam_alias: How to generate the identity alias when using the iam auth method. Valid choices are
               `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        :param pulumi.Input[Sequence[pulumi.Input[builtins.str]]] iam_metadatas: The metadata to include on the token returned by the `login` endpoint. This metadata will be
               added to both audit logs, and on the `iam_alias`
        :param pulumi.Input[builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendConfigIdentityState.__new__(_AuthBackendConfigIdentityState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["ec2_alias"] = ec2_alias
        __props__.__dict__["ec2_metadatas"] = ec2_metadatas
        __props__.__dict__["iam_alias"] = iam_alias
        __props__.__dict__["iam_metadatas"] = iam_metadatas
        __props__.__dict__["namespace"] = namespace
        return AuthBackendConfigIdentity(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        Unique name of the auth backend to configure.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="ec2Alias")
    def ec2_alias(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        How to generate the identity alias when using the ec2 auth method. Valid choices are
        `role_id`, `instance_id`, and `image_id`. Defaults to `role_id`
        """
        return pulumi.get(self, "ec2_alias")

    @property
    @pulumi.getter(name="ec2Metadatas")
    def ec2_metadatas(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `ec2_alias`
        """
        return pulumi.get(self, "ec2_metadatas")

    @property
    @pulumi.getter(name="iamAlias")
    def iam_alias(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        How to generate the identity alias when using the iam auth method. Valid choices are
        `role_id`, `unique_id`, and `full_arn`. Defaults to `role_id`
        """
        return pulumi.get(self, "iam_alias")

    @property
    @pulumi.getter(name="iamMetadatas")
    def iam_metadatas(self) -> pulumi.Output[Optional[Sequence[builtins.str]]]:
        """
        The metadata to include on the token returned by the `login` endpoint. This metadata will be
        added to both audit logs, and on the `iam_alias`
        """
        return pulumi.get(self, "iam_metadatas")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

