# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import builtins as _builtins
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

__all__ = ['SyncGithubAppsArgs', 'SyncGithubApps']

@pulumi.input_type
class SyncGithubAppsArgs:
    def __init__(__self__, *,
                 app_id: pulumi.Input[_builtins.int],
                 private_key: pulumi.Input[_builtins.str],
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None):
        """
        The set of arguments for constructing a SyncGithubApps resource.
        :param pulumi.Input[_builtins.int] app_id: The GitHub application ID.
        :param pulumi.Input[_builtins.str] private_key: The content of a PEM formatted private key generated on GitHub for the app.
        :param pulumi.Input[_builtins.str] name: The user-defined name of the GitHub App configuration.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        pulumi.set(__self__, "app_id", app_id)
        pulumi.set(__self__, "private_key", private_key)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @_builtins.property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Input[_builtins.int]:
        """
        The GitHub application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: pulumi.Input[_builtins.int]):
        pulumi.set(self, "app_id", value)

    @_builtins.property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Input[_builtins.str]:
        """
        The content of a PEM formatted private key generated on GitHub for the app.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "private_key", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The user-defined name of the GitHub App configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)


@pulumi.input_type
class _SyncGithubAppsState:
    def __init__(__self__, *,
                 app_id: Optional[pulumi.Input[_builtins.int]] = None,
                 fingerprint: Optional[pulumi.Input[_builtins.str]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key: Optional[pulumi.Input[_builtins.str]] = None):
        """
        Input properties used for looking up and filtering SyncGithubApps resources.
        :param pulumi.Input[_builtins.int] app_id: The GitHub application ID.
        :param pulumi.Input[_builtins.str] fingerprint: A fingerprint of a private key.
        :param pulumi.Input[_builtins.str] name: The user-defined name of the GitHub App configuration.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[_builtins.str] private_key: The content of a PEM formatted private key generated on GitHub for the app.
        """
        if app_id is not None:
            pulumi.set(__self__, "app_id", app_id)
        if fingerprint is not None:
            pulumi.set(__self__, "fingerprint", fingerprint)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if private_key is not None:
            pulumi.set(__self__, "private_key", private_key)

    @_builtins.property
    @pulumi.getter(name="appId")
    def app_id(self) -> Optional[pulumi.Input[_builtins.int]]:
        """
        The GitHub application ID.
        """
        return pulumi.get(self, "app_id")

    @app_id.setter
    def app_id(self, value: Optional[pulumi.Input[_builtins.int]]):
        pulumi.set(self, "app_id", value)

    @_builtins.property
    @pulumi.getter
    def fingerprint(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        A fingerprint of a private key.
        """
        return pulumi.get(self, "fingerprint")

    @fingerprint.setter
    def fingerprint(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "fingerprint", value)

    @_builtins.property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The user-defined name of the GitHub App configuration.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "name", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The content of a PEM formatted private key generated on GitHub for the app.
        """
        return pulumi.get(self, "private_key")

    @private_key.setter
    def private_key(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "private_key", value)


@pulumi.type_token("vault:secrets/syncGithubApps:SyncGithubApps")
class SyncGithubApps(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        github_apps = vault.secrets.SyncGithubApps("github-apps",
            name="gh-apps",
            app_id=app_id,
            private_key=std.file(input=privatekey_file).result)
        ```

        ## Import

        GitHub Apps Secrets sync configuration endpoint can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGithubApps:SyncGithubApps gh github-apps
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.int] app_id: The GitHub application ID.
        :param pulumi.Input[_builtins.str] name: The user-defined name of the GitHub App configuration.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[_builtins.str] private_key: The content of a PEM formatted private key generated on GitHub for the app.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SyncGithubAppsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        github_apps = vault.secrets.SyncGithubApps("github-apps",
            name="gh-apps",
            app_id=app_id,
            private_key=std.file(input=privatekey_file).result)
        ```

        ## Import

        GitHub Apps Secrets sync configuration endpoint can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGithubApps:SyncGithubApps gh github-apps
        ```

        :param str resource_name: The name of the resource.
        :param SyncGithubAppsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncGithubAppsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 app_id: Optional[pulumi.Input[_builtins.int]] = None,
                 name: Optional[pulumi.Input[_builtins.str]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 private_key: Optional[pulumi.Input[_builtins.str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncGithubAppsArgs.__new__(SyncGithubAppsArgs)

            if app_id is None and not opts.urn:
                raise TypeError("Missing required property 'app_id'")
            __props__.__dict__["app_id"] = app_id
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            if private_key is None and not opts.urn:
                raise TypeError("Missing required property 'private_key'")
            __props__.__dict__["private_key"] = None if private_key is None else pulumi.Output.secret(private_key)
            __props__.__dict__["fingerprint"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["privateKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyncGithubApps, __self__).__init__(
            'vault:secrets/syncGithubApps:SyncGithubApps',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            app_id: Optional[pulumi.Input[_builtins.int]] = None,
            fingerprint: Optional[pulumi.Input[_builtins.str]] = None,
            name: Optional[pulumi.Input[_builtins.str]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            private_key: Optional[pulumi.Input[_builtins.str]] = None) -> 'SyncGithubApps':
        """
        Get an existing SyncGithubApps resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.int] app_id: The GitHub application ID.
        :param pulumi.Input[_builtins.str] fingerprint: A fingerprint of a private key.
        :param pulumi.Input[_builtins.str] name: The user-defined name of the GitHub App configuration.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[_builtins.str] private_key: The content of a PEM formatted private key generated on GitHub for the app.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncGithubAppsState.__new__(_SyncGithubAppsState)

        __props__.__dict__["app_id"] = app_id
        __props__.__dict__["fingerprint"] = fingerprint
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["private_key"] = private_key
        return SyncGithubApps(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter(name="appId")
    def app_id(self) -> pulumi.Output[_builtins.int]:
        """
        The GitHub application ID.
        """
        return pulumi.get(self, "app_id")

    @_builtins.property
    @pulumi.getter
    def fingerprint(self) -> pulumi.Output[_builtins.str]:
        """
        A fingerprint of a private key.
        """
        return pulumi.get(self, "fingerprint")

    @_builtins.property
    @pulumi.getter
    def name(self) -> pulumi.Output[_builtins.str]:
        """
        The user-defined name of the GitHub App configuration.
        """
        return pulumi.get(self, "name")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter(name="privateKey")
    def private_key(self) -> pulumi.Output[_builtins.str]:
        """
        The content of a PEM formatted private key generated on GitHub for the app.
        """
        return pulumi.get(self, "private_key")

