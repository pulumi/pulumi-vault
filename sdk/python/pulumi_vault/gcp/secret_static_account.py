# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecretStaticAccountArgs', 'SecretStaticAccount']

@pulumi.input_type
class SecretStaticAccountArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 service_account_email: pulumi.Input[str],
                 static_account: pulumi.Input[str],
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecretStaticAccount resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to manage.
        :param pulumi.Input[str] static_account: Name of the Static Account to create
        :param pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]] bindings: Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] secret_type: Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "service_account_email", service_account_email)
        pulumi.set(__self__, "static_account", static_account)
        if bindings is not None:
            pulumi.set(__self__, "bindings", bindings)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Input[str]:
        """
        Email of the GCP service account to manage.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter(name="staticAccount")
    def static_account(self) -> pulumi.Input[str]:
        """
        Name of the Static Account to create
        """
        return pulumi.get(self, "static_account")

    @static_account.setter
    def static_account(self, value: pulumi.Input[str]):
        pulumi.set(self, "static_account", value)

    @property
    @pulumi.getter
    def bindings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]]:
        """
        Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @bindings.setter
    def bindings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]]):
        pulumi.set(self, "bindings", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_type", value)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_scopes", value)


@pulumi.input_type
class _SecretStaticAccountState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 service_account_project: Optional[pulumi.Input[str]] = None,
                 static_account: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering SecretStaticAccount resources.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]] bindings: Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] secret_type: Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to manage.
        :param pulumi.Input[str] service_account_project: Project the service account belongs to.
        :param pulumi.Input[str] static_account: Name of the Static Account to create
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if bindings is not None:
            pulumi.set(__self__, "bindings", bindings)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)
        if service_account_project is not None:
            pulumi.set(__self__, "service_account_project", service_account_project)
        if static_account is not None:
            pulumi.set(__self__, "static_account", static_account)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def bindings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]]:
        """
        Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @bindings.setter
    def bindings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretStaticAccountBindingArgs']]]]):
        pulumi.set(self, "bindings", value)

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[str]]:
        """
        Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_type", value)

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[str]]:
        """
        Email of the GCP service account to manage.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_email", value)

    @property
    @pulumi.getter(name="serviceAccountProject")
    def service_account_project(self) -> Optional[pulumi.Input[str]]:
        """
        Project the service account belongs to.
        """
        return pulumi.get(self, "service_account_project")

    @service_account_project.setter
    def service_account_project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "service_account_project", value)

    @property
    @pulumi.getter(name="staticAccount")
    def static_account(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the Static Account to create
        """
        return pulumi.get(self, "static_account")

    @static_account.setter
    def static_account(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "static_account", value)

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "token_scopes", value)


class SecretStaticAccount(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretStaticAccountBindingArgs']]]]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 static_account: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Creates a Static Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each [static account](https://www.vaultproject.io/docs/secrets/gcp/index.html#static-accounts) is tied to a separately managed
        Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings) associated with it.

        ## Import

        A static account can be imported using its Vault Path. For example, referencing the example above,

        ```sh
         $ pulumi import vault:gcp/secretStaticAccount:SecretStaticAccount static_account gcp/static-account/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretStaticAccountBindingArgs']]]] bindings: Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] secret_type: Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to manage.
        :param pulumi.Input[str] static_account: Name of the Static Account to create
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretStaticAccountArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Static Account in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each [static account](https://www.vaultproject.io/docs/secrets/gcp/index.html#static-accounts) is tied to a separately managed
        Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings) associated with it.

        ## Import

        A static account can be imported using its Vault Path. For example, referencing the example above,

        ```sh
         $ pulumi import vault:gcp/secretStaticAccount:SecretStaticAccount static_account gcp/static-account/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param SecretStaticAccountArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretStaticAccountArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretStaticAccountBindingArgs']]]]] = None,
                 secret_type: Optional[pulumi.Input[str]] = None,
                 service_account_email: Optional[pulumi.Input[str]] = None,
                 static_account: Optional[pulumi.Input[str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretStaticAccountArgs.__new__(SecretStaticAccountArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["bindings"] = bindings
            __props__.__dict__["secret_type"] = secret_type
            if service_account_email is None and not opts.urn:
                raise TypeError("Missing required property 'service_account_email'")
            __props__.__dict__["service_account_email"] = service_account_email
            if static_account is None and not opts.urn:
                raise TypeError("Missing required property 'static_account'")
            __props__.__dict__["static_account"] = static_account
            __props__.__dict__["token_scopes"] = token_scopes
            __props__.__dict__["service_account_project"] = None
        super(SecretStaticAccount, __self__).__init__(
            'vault:gcp/secretStaticAccount:SecretStaticAccount',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            bindings: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretStaticAccountBindingArgs']]]]] = None,
            secret_type: Optional[pulumi.Input[str]] = None,
            service_account_email: Optional[pulumi.Input[str]] = None,
            service_account_project: Optional[pulumi.Input[str]] = None,
            static_account: Optional[pulumi.Input[str]] = None,
            token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None) -> 'SecretStaticAccount':
        """
        Get an existing SecretStaticAccount resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretStaticAccountBindingArgs']]]] bindings: Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[str] secret_type: Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[str] service_account_email: Email of the GCP service account to manage.
        :param pulumi.Input[str] service_account_project: Project the service account belongs to.
        :param pulumi.Input[str] static_account: Name of the Static Account to create
        :param pulumi.Input[Sequence[pulumi.Input[str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretStaticAccountState.__new__(_SecretStaticAccountState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["bindings"] = bindings
        __props__.__dict__["secret_type"] = secret_type
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["service_account_project"] = service_account_project
        __props__.__dict__["static_account"] = static_account
        __props__.__dict__["token_scopes"] = token_scopes
        return SecretStaticAccount(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def bindings(self) -> pulumi.Output[Optional[Sequence['outputs.SecretStaticAccountBinding']]]:
        """
        Bindings to create for this static account. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> pulumi.Output[str]:
        """
        Type of secret generated for this static account. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[str]:
        """
        Email of the GCP service account to manage.
        """
        return pulumi.get(self, "service_account_email")

    @property
    @pulumi.getter(name="serviceAccountProject")
    def service_account_project(self) -> pulumi.Output[str]:
        """
        Project the service account belongs to.
        """
        return pulumi.get(self, "service_account_project")

    @property
    @pulumi.getter(name="staticAccount")
    def static_account(self) -> pulumi.Output[str]:
        """
        Name of the Static Account to create
        """
        return pulumi.get(self, "static_account")

    @property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this static account (`access_token` static accounts only).
        """
        return pulumi.get(self, "token_scopes")

