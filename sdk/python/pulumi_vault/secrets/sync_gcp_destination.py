# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncGcpDestinationArgs', 'SyncGcpDestination']

@pulumi.input_type
class SyncGcpDestinationArgs:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SyncGcpDestination resource.
        :param pulumi.Input[str] credentials: JSON-encoded credentials to use to connect to GCP.
               Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
               variable.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] name: Unique name of the GCP destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] project_id: The target project to manage secrets in. If set,
               overrides the project ID derived from the service account JSON credentials or application
               default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
               to perform Secret Manager actions in the target project.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if custom_tags is not None:
            pulumi.set(__self__, "custom_tags", custom_tags)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[str]]:
        """
        JSON-encoded credentials to use to connect to GCP.
        Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
        variable.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="customTags")
    def custom_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Custom tags to set on the secret managed at the destination.
        """
        return pulumi.get(self, "custom_tags")

    @custom_tags.setter
    def custom_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "custom_tags", value)

    @property
    @pulumi.getter
    def granularity(self) -> Optional[pulumi.Input[str]]:
        """
        Determines what level of information is synced as a distinct resource
        at the destination. Supports `secret-path` and `secret-key`.
        """
        return pulumi.get(self, "granularity")

    @granularity.setter
    def granularity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "granularity", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the GCP destination.
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
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target project to manage secrets in. If set,
        overrides the project ID derived from the service account JSON credentials or application
        default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
        to perform Secret Manager actions in the target project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

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
class _SyncGcpDestinationState:
    def __init__(__self__, *,
                 credentials: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyncGcpDestination resources.
        :param pulumi.Input[str] credentials: JSON-encoded credentials to use to connect to GCP.
               Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
               variable.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] name: Unique name of the GCP destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] project_id: The target project to manage secrets in. If set,
               overrides the project ID derived from the service account JSON credentials or application
               default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
               to perform Secret Manager actions in the target project.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] type: The type of the secrets destination (`gcp-sm`).
        """
        if credentials is not None:
            pulumi.set(__self__, "credentials", credentials)
        if custom_tags is not None:
            pulumi.set(__self__, "custom_tags", custom_tags)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if project_id is not None:
            pulumi.set(__self__, "project_id", project_id)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def credentials(self) -> Optional[pulumi.Input[str]]:
        """
        JSON-encoded credentials to use to connect to GCP.
        Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
        variable.
        """
        return pulumi.get(self, "credentials")

    @credentials.setter
    def credentials(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "credentials", value)

    @property
    @pulumi.getter(name="customTags")
    def custom_tags(self) -> Optional[pulumi.Input[Mapping[str, Any]]]:
        """
        Custom tags to set on the secret managed at the destination.
        """
        return pulumi.get(self, "custom_tags")

    @custom_tags.setter
    def custom_tags(self, value: Optional[pulumi.Input[Mapping[str, Any]]]):
        pulumi.set(self, "custom_tags", value)

    @property
    @pulumi.getter
    def granularity(self) -> Optional[pulumi.Input[str]]:
        """
        Determines what level of information is synced as a distinct resource
        at the destination. Supports `secret-path` and `secret-key`.
        """
        return pulumi.get(self, "granularity")

    @granularity.setter
    def granularity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "granularity", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the GCP destination.
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
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> Optional[pulumi.Input[str]]:
        """
        The target project to manage secrets in. If set,
        overrides the project ID derived from the service account JSON credentials or application
        default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
        to perform Secret Manager actions in the target project.
        """
        return pulumi.get(self, "project_id")

    @project_id.setter
    def project_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project_id", value)

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
        The type of the secrets destination (`gcp-sm`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SyncGcpDestination(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        gcp = vault.secrets.SyncGcpDestination("gcp",
            name="gcp-dest",
            project_id="gcp-project-id",
            credentials=std.file(input=credentials_file).result,
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
            custom_tags={
                "foo": "bar",
            })
        ```

        ## Import

        GCP Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGcpDestination:SyncGcpDestination gcp gcp-dest
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] credentials: JSON-encoded credentials to use to connect to GCP.
               Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
               variable.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] name: Unique name of the GCP destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] project_id: The target project to manage secrets in. If set,
               overrides the project ID derived from the service account JSON credentials or application
               default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
               to perform Secret Manager actions in the target project.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SyncGcpDestinationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_std as std
        import pulumi_vault as vault

        gcp = vault.secrets.SyncGcpDestination("gcp",
            name="gcp-dest",
            project_id="gcp-project-id",
            credentials=std.file(input=credentials_file).result,
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
            custom_tags={
                "foo": "bar",
            })
        ```

        ## Import

        GCP Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncGcpDestination:SyncGcpDestination gcp gcp-dest
        ```

        :param str resource_name: The name of the resource.
        :param SyncGcpDestinationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncGcpDestinationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 credentials: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 project_id: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncGcpDestinationArgs.__new__(SyncGcpDestinationArgs)

            __props__.__dict__["credentials"] = None if credentials is None else pulumi.Output.secret(credentials)
            __props__.__dict__["custom_tags"] = custom_tags
            __props__.__dict__["granularity"] = granularity
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["project_id"] = project_id
            __props__.__dict__["secret_name_template"] = secret_name_template
            __props__.__dict__["type"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["credentials"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyncGcpDestination, __self__).__init__(
            'vault:secrets/syncGcpDestination:SyncGcpDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            credentials: Optional[pulumi.Input[str]] = None,
            custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            granularity: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            project_id: Optional[pulumi.Input[str]] = None,
            secret_name_template: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SyncGcpDestination':
        """
        Get an existing SyncGcpDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] credentials: JSON-encoded credentials to use to connect to GCP.
               Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
               variable.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] name: Unique name of the GCP destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] project_id: The target project to manage secrets in. If set,
               overrides the project ID derived from the service account JSON credentials or application
               default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
               to perform Secret Manager actions in the target project.
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] type: The type of the secrets destination (`gcp-sm`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncGcpDestinationState.__new__(_SyncGcpDestinationState)

        __props__.__dict__["credentials"] = credentials
        __props__.__dict__["custom_tags"] = custom_tags
        __props__.__dict__["granularity"] = granularity
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["project_id"] = project_id
        __props__.__dict__["secret_name_template"] = secret_name_template
        __props__.__dict__["type"] = type
        return SyncGcpDestination(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def credentials(self) -> pulumi.Output[Optional[str]]:
        """
        JSON-encoded credentials to use to connect to GCP.
        Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
        variable.
        """
        return pulumi.get(self, "credentials")

    @property
    @pulumi.getter(name="customTags")
    def custom_tags(self) -> pulumi.Output[Optional[Mapping[str, Any]]]:
        """
        Custom tags to set on the secret managed at the destination.
        """
        return pulumi.get(self, "custom_tags")

    @property
    @pulumi.getter
    def granularity(self) -> pulumi.Output[Optional[str]]:
        """
        Determines what level of information is synced as a distinct resource
        at the destination. Supports `secret-path` and `secret-key`.
        """
        return pulumi.get(self, "granularity")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the GCP destination.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def namespace(self) -> pulumi.Output[Optional[str]]:
        """
        The namespace to provision the resource in.
        The value should not contain leading or trailing forward slashes.
        The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        """
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> pulumi.Output[Optional[str]]:
        """
        The target project to manage secrets in. If set,
        overrides the project ID derived from the service account JSON credentials or application
        default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
        to perform Secret Manager actions in the target project.
        """
        return pulumi.get(self, "project_id")

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
        The type of the secrets destination (`gcp-sm`).
        """
        return pulumi.get(self, "type")

