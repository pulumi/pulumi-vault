# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncGhDestinationArgs', 'SyncGhDestination']

@pulumi.input_type
class SyncGhDestinationArgs:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_owner: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SyncGhDestination resource.
        :param pulumi.Input[str] access_token: Fine-grained or personal access token.
               Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the GitHub destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        :param pulumi.Input[str] repository_name: Name of the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
               variable.
        :param pulumi.Input[str] repository_owner: GitHub organization or username that owns the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
               variable.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if repository_name is not None:
            pulumi.set(__self__, "repository_name", repository_name)
        if repository_owner is not None:
            pulumi.set(__self__, "repository_owner", repository_owner)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        Fine-grained or personal access token.
        Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the GitHub destination.
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
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        variable.
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter(name="repositoryOwner")
    def repository_owner(self) -> Optional[pulumi.Input[str]]:
        """
        GitHub organization or username that owns the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        variable.
        """
        return pulumi.get(self, "repository_owner")

    @repository_owner.setter
    def repository_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_owner", value)

    @property
    @pulumi.getter(name="secretNameTemplate")
    def secret_name_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template describing how to generate external secret names.
        Supports a subset of the Go Template syntax.
        """
        return pulumi.get(self, "secret_name_template")

    @secret_name_template.setter
    def secret_name_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name_template", value)


@pulumi.input_type
class _SyncGhDestinationState:
    def __init__(__self__, *,
                 access_token: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_owner: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyncGhDestination resources.
        :param pulumi.Input[str] access_token: Fine-grained or personal access token.
               Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the GitHub destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        :param pulumi.Input[str] repository_name: Name of the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
               variable.
        :param pulumi.Input[str] repository_owner: GitHub organization or username that owns the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
               variable.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] type: The type of the secrets destination (`gh`).
        """
        if access_token is not None:
            pulumi.set(__self__, "access_token", access_token)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if repository_name is not None:
            pulumi.set(__self__, "repository_name", repository_name)
        if repository_owner is not None:
            pulumi.set(__self__, "repository_owner", repository_owner)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> Optional[pulumi.Input[str]]:
        """
        Fine-grained or personal access token.
        Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "access_token")

    @access_token.setter
    def access_token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_token", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the GitHub destination.
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
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        variable.
        """
        return pulumi.get(self, "repository_name")

    @repository_name.setter
    def repository_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_name", value)

    @property
    @pulumi.getter(name="repositoryOwner")
    def repository_owner(self) -> Optional[pulumi.Input[str]]:
        """
        GitHub organization or username that owns the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        variable.
        """
        return pulumi.get(self, "repository_owner")

    @repository_owner.setter
    def repository_owner(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "repository_owner", value)

    @property
    @pulumi.getter(name="secretNameTemplate")
    def secret_name_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template describing how to generate external secret names.
        Supports a subset of the Go Template syntax.
        """
        return pulumi.get(self, "secret_name_template")

    @secret_name_template.setter
    def secret_name_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_name_template", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the secrets destination (`gh`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SyncGhDestination(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_owner: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        gh = vault.secrets.SyncGhDestination("gh",
            access_token=var["access_token"],
            repository_owner=var["repo_owner"],
            repository_name="repo-name-example",
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGhDestination:SyncGhDestination gh gh-dest
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: Fine-grained or personal access token.
               Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the GitHub destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        :param pulumi.Input[str] repository_name: Name of the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
               variable.
        :param pulumi.Input[str] repository_owner: GitHub organization or username that owns the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
               variable.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SyncGhDestinationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        gh = vault.secrets.SyncGhDestination("gh",
            access_token=var["access_token"],
            repository_owner=var["repo_owner"],
            repository_name="repo-name-example",
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        GitHub Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGhDestination:SyncGhDestination gh gh-dest
        ```

        :param str resource_name: The name of the resource.
        :param SyncGhDestinationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncGhDestinationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_token: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 repository_name: Optional[pulumi.Input[str]] = None,
                 repository_owner: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncGhDestinationArgs.__new__(SyncGhDestinationArgs)

            __props__.__dict__["access_token"] = None if access_token is None else pulumi.Output.secret(access_token)
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["repository_name"] = repository_name
            __props__.__dict__["repository_owner"] = repository_owner
            __props__.__dict__["secret_name_template"] = secret_name_template
            __props__.__dict__["type"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessToken"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyncGhDestination, __self__).__init__(
            'vault:secrets/syncGhDestination:SyncGhDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_token: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            repository_name: Optional[pulumi.Input[str]] = None,
            repository_owner: Optional[pulumi.Input[str]] = None,
            secret_name_template: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SyncGhDestination':
        """
        Get an existing SyncGhDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_token: Fine-grained or personal access token.
               Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the GitHub destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        :param pulumi.Input[str] repository_name: Name of the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
               variable.
        :param pulumi.Input[str] repository_owner: GitHub organization or username that owns the repository.
               Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
               variable.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] type: The type of the secrets destination (`gh`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncGhDestinationState.__new__(_SyncGhDestinationState)

        __props__.__dict__["access_token"] = access_token
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["repository_name"] = repository_name
        __props__.__dict__["repository_owner"] = repository_owner
        __props__.__dict__["secret_name_template"] = secret_name_template
        __props__.__dict__["type"] = type
        return SyncGhDestination(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessToken")
    def access_token(self) -> pulumi.Output[Optional[str]]:
        """
        Fine-grained or personal access token.
        Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
        variable.
        """
        return pulumi.get(self, "access_token")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the GitHub destination.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="repositoryName")
    def repository_name(self) -> pulumi.Output[Optional[str]]:
        """
        Name of the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
        variable.
        """
        return pulumi.get(self, "repository_name")

    @property
    @pulumi.getter(name="repositoryOwner")
    def repository_owner(self) -> pulumi.Output[Optional[str]]:
        """
        GitHub organization or username that owns the repository.
        Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
        variable.
        """
        return pulumi.get(self, "repository_owner")

    @property
    @pulumi.getter(name="secretNameTemplate")
    def secret_name_template(self) -> pulumi.Output[str]:
        """
        Template describing how to generate external secret names.
        Supports a subset of the Go Template syntax.
        """
        return pulumi.get(self, "secret_name_template")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the secrets destination (`gh`).
        """
        return pulumi.get(self, "type")

