# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SyncAzureDestinationArgs', 'SyncAzureDestination']

@pulumi.input_type
class SyncAzureDestinationArgs:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 cloud: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SyncAzureDestination resource.
        :param pulumi.Input[str] client_id: Client ID of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
               variable.
        :param pulumi.Input[str] client_secret: Client Secret of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
               variable.
        :param pulumi.Input[str] cloud: Specifies a cloud for the client. The default is Azure Public Cloud.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] key_vault_uri: URI of an existing Azure Key Vault instance.
               Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the Azure destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] tenant_id: ID of the target Azure tenant.
               Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
               variable.
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if cloud is not None:
            pulumi.set(__self__, "cloud", cloud)
        if custom_tags is not None:
            pulumi.set(__self__, "custom_tags", custom_tags)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client ID of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        variable.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        Client Secret of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        variable.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def cloud(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a cloud for the client. The default is Azure Public Cloud.
        """
        return pulumi.get(self, "cloud")

    @cloud.setter
    def cloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud", value)

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
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI of an existing Azure Key Vault instance.
        Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        variable.
        """
        return pulumi.get(self, "key_vault_uri")

    @key_vault_uri.setter
    def key_vault_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_uri", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the Azure destination.
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
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the target Azure tenant.
        Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        variable.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)


@pulumi.input_type
class _SyncAzureDestinationState:
    def __init__(__self__, *,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 cloud: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SyncAzureDestination resources.
        :param pulumi.Input[str] client_id: Client ID of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
               variable.
        :param pulumi.Input[str] client_secret: Client Secret of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
               variable.
        :param pulumi.Input[str] cloud: Specifies a cloud for the client. The default is Azure Public Cloud.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] key_vault_uri: URI of an existing Azure Key Vault instance.
               Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the Azure destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] tenant_id: ID of the target Azure tenant.
               Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
               variable.
        :param pulumi.Input[str] type: The type of the secrets destination (`azure-kv`).
        """
        if client_id is not None:
            pulumi.set(__self__, "client_id", client_id)
        if client_secret is not None:
            pulumi.set(__self__, "client_secret", client_secret)
        if cloud is not None:
            pulumi.set(__self__, "cloud", cloud)
        if custom_tags is not None:
            pulumi.set(__self__, "custom_tags", custom_tags)
        if granularity is not None:
            pulumi.set(__self__, "granularity", granularity)
        if key_vault_uri is not None:
            pulumi.set(__self__, "key_vault_uri", key_vault_uri)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_name_template is not None:
            pulumi.set(__self__, "secret_name_template", secret_name_template)
        if tenant_id is not None:
            pulumi.set(__self__, "tenant_id", tenant_id)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> Optional[pulumi.Input[str]]:
        """
        Client ID of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        variable.
        """
        return pulumi.get(self, "client_id")

    @client_id.setter
    def client_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_id", value)

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> Optional[pulumi.Input[str]]:
        """
        Client Secret of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        variable.
        """
        return pulumi.get(self, "client_secret")

    @client_secret.setter
    def client_secret(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "client_secret", value)

    @property
    @pulumi.getter
    def cloud(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a cloud for the client. The default is Azure Public Cloud.
        """
        return pulumi.get(self, "cloud")

    @cloud.setter
    def cloud(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "cloud", value)

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
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> Optional[pulumi.Input[str]]:
        """
        URI of an existing Azure Key Vault instance.
        Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        variable.
        """
        return pulumi.get(self, "key_vault_uri")

    @key_vault_uri.setter
    def key_vault_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "key_vault_uri", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique name of the Azure destination.
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
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the target Azure tenant.
        Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        variable.
        """
        return pulumi.get(self, "tenant_id")

    @tenant_id.setter
    def tenant_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tenant_id", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the secrets destination (`azure-kv`).
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class SyncAzureDestination(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 cloud: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        az = vault.secrets.SyncAzureDestination("az",
            key_vault_uri=var["key_vault_uri"],
            client_id=var["client_id"],
            client_secret=var["client_secret"],
            tenant_id=var["tenant_id"],
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
            custom_tags={
                "foo": "bar",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Azure Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncAzureDestination:SyncAzureDestination az az-dest
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: Client ID of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
               variable.
        :param pulumi.Input[str] client_secret: Client Secret of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
               variable.
        :param pulumi.Input[str] cloud: Specifies a cloud for the client. The default is Azure Public Cloud.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] key_vault_uri: URI of an existing Azure Key Vault instance.
               Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the Azure destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] tenant_id: ID of the target Azure tenant.
               Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
               variable.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SyncAzureDestinationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        az = vault.secrets.SyncAzureDestination("az",
            key_vault_uri=var["key_vault_uri"],
            client_id=var["client_id"],
            client_secret=var["client_secret"],
            tenant_id=var["tenant_id"],
            secret_name_template="vault_{{ .MountAccessor | lowercase }}_{{ .SecretPath | lowercase }}",
            custom_tags={
                "foo": "bar",
            })
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        Azure Secrets sync destinations can be imported using the `name`, e.g.

        ```sh
        $ pulumi import vault:secrets/syncAzureDestination:SyncAzureDestination az az-dest
        ```

        :param str resource_name: The name of the resource.
        :param SyncAzureDestinationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SyncAzureDestinationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 client_id: Optional[pulumi.Input[str]] = None,
                 client_secret: Optional[pulumi.Input[str]] = None,
                 cloud: Optional[pulumi.Input[str]] = None,
                 custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
                 granularity: Optional[pulumi.Input[str]] = None,
                 key_vault_uri: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_name_template: Optional[pulumi.Input[str]] = None,
                 tenant_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SyncAzureDestinationArgs.__new__(SyncAzureDestinationArgs)

            __props__.__dict__["client_id"] = client_id
            __props__.__dict__["client_secret"] = None if client_secret is None else pulumi.Output.secret(client_secret)
            __props__.__dict__["cloud"] = cloud
            __props__.__dict__["custom_tags"] = custom_tags
            __props__.__dict__["granularity"] = granularity
            __props__.__dict__["key_vault_uri"] = key_vault_uri
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["secret_name_template"] = secret_name_template
            __props__.__dict__["tenant_id"] = tenant_id
            __props__.__dict__["type"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["clientSecret"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SyncAzureDestination, __self__).__init__(
            'vault:secrets/syncAzureDestination:SyncAzureDestination',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            client_id: Optional[pulumi.Input[str]] = None,
            client_secret: Optional[pulumi.Input[str]] = None,
            cloud: Optional[pulumi.Input[str]] = None,
            custom_tags: Optional[pulumi.Input[Mapping[str, Any]]] = None,
            granularity: Optional[pulumi.Input[str]] = None,
            key_vault_uri: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            secret_name_template: Optional[pulumi.Input[str]] = None,
            tenant_id: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'SyncAzureDestination':
        """
        Get an existing SyncAzureDestination resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] client_id: Client ID of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
               variable.
        :param pulumi.Input[str] client_secret: Client Secret of an Azure app registration.
               Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
               variable.
        :param pulumi.Input[str] cloud: Specifies a cloud for the client. The default is Azure Public Cloud.
        :param pulumi.Input[Mapping[str, Any]] custom_tags: Custom tags to set on the secret managed at the destination.
        :param pulumi.Input[str] granularity: Determines what level of information is synced as a distinct resource
               at the destination. Supports `secret-path` and `secret-key`.
        :param pulumi.Input[str] key_vault_uri: URI of an existing Azure Key Vault instance.
               Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
               variable.
        :param pulumi.Input[str] name: Unique name of the Azure destination.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
        :param pulumi.Input[str] secret_name_template: Template describing how to generate external secret names.
               Supports a subset of the Go Template syntax.
        :param pulumi.Input[str] tenant_id: ID of the target Azure tenant.
               Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
               variable.
        :param pulumi.Input[str] type: The type of the secrets destination (`azure-kv`).
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SyncAzureDestinationState.__new__(_SyncAzureDestinationState)

        __props__.__dict__["client_id"] = client_id
        __props__.__dict__["client_secret"] = client_secret
        __props__.__dict__["cloud"] = cloud
        __props__.__dict__["custom_tags"] = custom_tags
        __props__.__dict__["granularity"] = granularity
        __props__.__dict__["key_vault_uri"] = key_vault_uri
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["secret_name_template"] = secret_name_template
        __props__.__dict__["tenant_id"] = tenant_id
        __props__.__dict__["type"] = type
        return SyncAzureDestination(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clientId")
    def client_id(self) -> pulumi.Output[Optional[str]]:
        """
        Client ID of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_ID` environment
        variable.
        """
        return pulumi.get(self, "client_id")

    @property
    @pulumi.getter(name="clientSecret")
    def client_secret(self) -> pulumi.Output[Optional[str]]:
        """
        Client Secret of an Azure app registration.
        Can be omitted and directly provided to Vault using the `AZURE_CLIENT_SECRET` environment
        variable.
        """
        return pulumi.get(self, "client_secret")

    @property
    @pulumi.getter
    def cloud(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a cloud for the client. The default is Azure Public Cloud.
        """
        return pulumi.get(self, "cloud")

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
    @pulumi.getter(name="keyVaultUri")
    def key_vault_uri(self) -> pulumi.Output[Optional[str]]:
        """
        URI of an existing Azure Key Vault instance.
        Can be omitted and directly provided to Vault using the `KEY_VAULT_URI` environment
        variable.
        """
        return pulumi.get(self, "key_vault_uri")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique name of the Azure destination.
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
    @pulumi.getter(name="secretNameTemplate")
    def secret_name_template(self) -> pulumi.Output[str]:
        """
        Template describing how to generate external secret names.
        Supports a subset of the Go Template syntax.
        """
        return pulumi.get(self, "secret_name_template")

    @property
    @pulumi.getter(name="tenantId")
    def tenant_id(self) -> pulumi.Output[Optional[str]]:
        """
        ID of the target Azure tenant.
        Can be omitted and directly provided to Vault using the `AZURE_TENANT_ID` environment
        variable.
        """
        return pulumi.get(self, "tenant_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of the secrets destination (`azure-kv`).
        """
        return pulumi.get(self, "type")

