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
from . import outputs
from ._inputs import *

__all__ = ['SecretRolesetArgs', 'SecretRoleset']

@pulumi.input_type
class SecretRolesetArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[_builtins.str],
                 bindings: pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]],
                 project: pulumi.Input[_builtins.str],
                 roleset: pulumi.Input[_builtins.str],
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_type: Optional[pulumi.Input[_builtins.str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        The set of arguments for constructing a SecretRoleset resource.
        :param pulumi.Input[_builtins.str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[_builtins.str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[_builtins.str] roleset: Name of the Roleset to create
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "bindings", bindings)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "roleset", roleset)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> pulumi.Input[_builtins.str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "backend", value)

    @_builtins.property
    @pulumi.getter
    def bindings(self) -> pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]]:
        """
        Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @bindings.setter
    def bindings(self, value: pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]]):
        pulumi.set(self, "bindings", value)

    @_builtins.property
    @pulumi.getter
    def project(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the GCP project that this roleset's service account will belong to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "project", value)

    @_builtins.property
    @pulumi.getter
    def roleset(self) -> pulumi.Input[_builtins.str]:
        """
        Name of the Roleset to create
        """
        return pulumi.get(self, "roleset")

    @roleset.setter
    def roleset(self, value: pulumi.Input[_builtins.str]):
        pulumi.set(self, "roleset", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "secret_type", value)

    @_builtins.property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "token_scopes", value)


@pulumi.input_type
class _SecretRolesetState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 roleset: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_type: Optional[pulumi.Input[_builtins.str]] = None,
                 service_account_email: Optional[pulumi.Input[_builtins.str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None):
        """
        Input properties used for looking up and filtering SecretRoleset resources.
        :param pulumi.Input[_builtins.str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[_builtins.str] roleset: Name of the Roleset to create
        :param pulumi.Input[_builtins.str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[_builtins.str] service_account_email: Email of the service account created by Vault for this Roleset.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if bindings is not None:
            pulumi.set(__self__, "bindings", bindings)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if roleset is not None:
            pulumi.set(__self__, "roleset", roleset)
        if secret_type is not None:
            pulumi.set(__self__, "secret_type", secret_type)
        if service_account_email is not None:
            pulumi.set(__self__, "service_account_email", service_account_email)
        if token_scopes is not None:
            pulumi.set(__self__, "token_scopes", token_scopes)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "backend", value)

    @_builtins.property
    @pulumi.getter
    def bindings(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]]]:
        """
        Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @bindings.setter
    def bindings(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretRolesetBindingArgs']]]]):
        pulumi.set(self, "bindings", value)

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "namespace", value)

    @_builtins.property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the GCP project that this roleset's service account will belong to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "project", value)

    @_builtins.property
    @pulumi.getter
    def roleset(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Name of the Roleset to create
        """
        return pulumi.get(self, "roleset")

    @roleset.setter
    def roleset(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "roleset", value)

    @_builtins.property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @secret_type.setter
    def secret_type(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "secret_type", value)

    @_builtins.property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> Optional[pulumi.Input[_builtins.str]]:
        """
        Email of the service account created by Vault for this Roleset.
        """
        return pulumi.get(self, "service_account_email")

    @service_account_email.setter
    def service_account_email(self, value: Optional[pulumi.Input[_builtins.str]]):
        pulumi.set(self, "service_account_email", value)

    @_builtins.property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        return pulumi.get(self, "token_scopes")

    @token_scopes.setter
    def token_scopes(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]]):
        pulumi.set(self, "token_scopes", value)


@pulumi.type_token("vault:gcp/secretRoleset:SecretRoleset")
class SecretRoleset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretRolesetBindingArgs', 'SecretRolesetBindingArgsDict']]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 roleset: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_type: Optional[pulumi.Input[_builtins.str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        """
        Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        project = "my-awesome-project"
        gcp = vault.gcp.SecretBackend("gcp",
            path="gcp",
            credentials=std.file(input="credentials.json").result)
        roleset = vault.gcp.SecretRoleset("roleset",
            backend=gcp.path,
            roleset="project_viewer",
            secret_type="access_token",
            project=project,
            token_scopes=["https://www.googleapis.com/auth/cloud-platform"],
            bindings=[{
                "resource": f"//cloudresourcemanager.googleapis.com/projects/{project}",
                "roles": ["roles/viewer"],
            }])
        ```

        ## Import

        A roleset can be imported using its Vault Path. For example, referencing the example above,

        ```sh
        $ pulumi import vault:gcp/secretRoleset:SecretRoleset roleset gcp/roleset/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[Union['SecretRolesetBindingArgs', 'SecretRolesetBindingArgsDict']]]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[_builtins.str] roleset: Name of the Roleset to create
        :param pulumi.Input[_builtins.str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretRolesetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Roleset in the [GCP Secrets Engine](https://www.vaultproject.io/docs/secrets/gcp/index.html) for Vault.

        Each Roleset is [tied](https://www.vaultproject.io/docs/secrets/gcp/index.html#service-accounts-are-tied-to-rolesets) to a Service Account, and can have one or more [bindings](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings) associated with it.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        project = "my-awesome-project"
        gcp = vault.gcp.SecretBackend("gcp",
            path="gcp",
            credentials=std.file(input="credentials.json").result)
        roleset = vault.gcp.SecretRoleset("roleset",
            backend=gcp.path,
            roleset="project_viewer",
            secret_type="access_token",
            project=project,
            token_scopes=["https://www.googleapis.com/auth/cloud-platform"],
            bindings=[{
                "resource": f"//cloudresourcemanager.googleapis.com/projects/{project}",
                "roles": ["roles/viewer"],
            }])
        ```

        ## Import

        A roleset can be imported using its Vault Path. For example, referencing the example above,

        ```sh
        $ pulumi import vault:gcp/secretRoleset:SecretRoleset roleset gcp/roleset/project_viewer
        ```

        :param str resource_name: The name of the resource.
        :param SecretRolesetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretRolesetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[_builtins.str]] = None,
                 bindings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretRolesetBindingArgs', 'SecretRolesetBindingArgsDict']]]]] = None,
                 namespace: Optional[pulumi.Input[_builtins.str]] = None,
                 project: Optional[pulumi.Input[_builtins.str]] = None,
                 roleset: Optional[pulumi.Input[_builtins.str]] = None,
                 secret_type: Optional[pulumi.Input[_builtins.str]] = None,
                 token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretRolesetArgs.__new__(SecretRolesetArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            if bindings is None and not opts.urn:
                raise TypeError("Missing required property 'bindings'")
            __props__.__dict__["bindings"] = bindings
            __props__.__dict__["namespace"] = namespace
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if roleset is None and not opts.urn:
                raise TypeError("Missing required property 'roleset'")
            __props__.__dict__["roleset"] = roleset
            __props__.__dict__["secret_type"] = secret_type
            __props__.__dict__["token_scopes"] = token_scopes
            __props__.__dict__["service_account_email"] = None
        super(SecretRoleset, __self__).__init__(
            'vault:gcp/secretRoleset:SecretRoleset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[_builtins.str]] = None,
            bindings: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SecretRolesetBindingArgs', 'SecretRolesetBindingArgsDict']]]]] = None,
            namespace: Optional[pulumi.Input[_builtins.str]] = None,
            project: Optional[pulumi.Input[_builtins.str]] = None,
            roleset: Optional[pulumi.Input[_builtins.str]] = None,
            secret_type: Optional[pulumi.Input[_builtins.str]] = None,
            service_account_email: Optional[pulumi.Input[_builtins.str]] = None,
            token_scopes: Optional[pulumi.Input[Sequence[pulumi.Input[_builtins.str]]]] = None) -> 'SecretRoleset':
        """
        Get an existing SecretRoleset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[_builtins.str] backend: Path where the GCP Secrets Engine is mounted
        :param pulumi.Input[Sequence[pulumi.Input[Union['SecretRolesetBindingArgs', 'SecretRolesetBindingArgsDict']]]] bindings: Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        :param pulumi.Input[_builtins.str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[_builtins.str] project: Name of the GCP project that this roleset's service account will belong to.
        :param pulumi.Input[_builtins.str] roleset: Name of the Roleset to create
        :param pulumi.Input[_builtins.str] secret_type: Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        :param pulumi.Input[_builtins.str] service_account_email: Email of the service account created by Vault for this Roleset.
        :param pulumi.Input[Sequence[pulumi.Input[_builtins.str]]] token_scopes: List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretRolesetState.__new__(_SecretRolesetState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["bindings"] = bindings
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["project"] = project
        __props__.__dict__["roleset"] = roleset
        __props__.__dict__["secret_type"] = secret_type
        __props__.__dict__["service_account_email"] = service_account_email
        __props__.__dict__["token_scopes"] = token_scopes
        return SecretRoleset(resource_name, opts=opts, __props__=__props__)

    @_builtins.property
    @pulumi.getter
    def backend(self) -> pulumi.Output[_builtins.str]:
        """
        Path where the GCP Secrets Engine is mounted
        """
        return pulumi.get(self, "backend")

    @_builtins.property
    @pulumi.getter
    def bindings(self) -> pulumi.Output[Sequence['outputs.SecretRolesetBinding']]:
        """
        Bindings to create for this roleset. This can be specified multiple times for multiple bindings. Structure is documented below.
        """
        return pulumi.get(self, "bindings")

    @_builtins.property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[_builtins.str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        *Available only for Vault Enterprise*.
        """
        return pulumi.get(self, "namespace")

    @_builtins.property
    @pulumi.getter
    def project(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the GCP project that this roleset's service account will belong to.
        """
        return pulumi.get(self, "project")

    @_builtins.property
    @pulumi.getter
    def roleset(self) -> pulumi.Output[_builtins.str]:
        """
        Name of the Roleset to create
        """
        return pulumi.get(self, "roleset")

    @_builtins.property
    @pulumi.getter(name="secretType")
    def secret_type(self) -> pulumi.Output[_builtins.str]:
        """
        Type of secret generated for this role set. Accepted values: `access_token`, `service_account_key`. Defaults to `access_token`.
        """
        return pulumi.get(self, "secret_type")

    @_builtins.property
    @pulumi.getter(name="serviceAccountEmail")
    def service_account_email(self) -> pulumi.Output[_builtins.str]:
        """
        Email of the service account created by Vault for this Roleset.
        """
        return pulumi.get(self, "service_account_email")

    @_builtins.property
    @pulumi.getter(name="tokenScopes")
    def token_scopes(self) -> pulumi.Output[Optional[Sequence[_builtins.str]]]:
        """
        List of OAuth scopes to assign to `access_token` secrets generated under this role set (`access_token` role sets only).
        """
        return pulumi.get(self, "token_scopes")

