# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['SecretBackendRoleArgs', 'SecretBackendRole']

@pulumi.input_type
class SecretBackendRoleArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[str]] = None,
                 vhost_topics: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]] = None,
                 vhosts: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]] = None):
        """
        The set of arguments for constructing a SecretBackendRole resource.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        :param pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]] vhost_topics: Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        :param pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]] vhosts: Specifies a map of virtual hosts to permissions.
        """
        pulumi.set(__self__, "backend", backend)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vhost_topics is not None:
            pulumi.set(__self__, "vhost_topics", vhost_topics)
        if vhosts is not None:
            pulumi.set(__self__, "vhosts", vhosts)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The path the RabbitMQ secret backend is mounted at,
        with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to identify this role within the backend.
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
    def tags(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a comma-separated RabbitMQ management tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vhostTopics")
    def vhost_topics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]]:
        """
        Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        """
        return pulumi.get(self, "vhost_topics")

    @vhost_topics.setter
    def vhost_topics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]]):
        pulumi.set(self, "vhost_topics", value)

    @property
    @pulumi.getter
    def vhosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]]:
        """
        Specifies a map of virtual hosts to permissions.
        """
        return pulumi.get(self, "vhosts")

    @vhosts.setter
    def vhosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]]):
        pulumi.set(self, "vhosts", value)


@pulumi.input_type
class _SecretBackendRoleState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[str]] = None,
                 vhost_topics: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]] = None,
                 vhosts: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]] = None):
        """
        Input properties used for looking up and filtering SecretBackendRole resources.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        :param pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]] vhost_topics: Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        :param pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]] vhosts: Specifies a map of virtual hosts to permissions.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if vhost_topics is not None:
            pulumi.set(__self__, "vhost_topics", vhost_topics)
        if vhosts is not None:
            pulumi.set(__self__, "vhosts", vhosts)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the RabbitMQ secret backend is mounted at,
        with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name to identify this role within the backend.
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
    def tags(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a comma-separated RabbitMQ management tags.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="vhostTopics")
    def vhost_topics(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]]:
        """
        Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        """
        return pulumi.get(self, "vhost_topics")

    @vhost_topics.setter
    def vhost_topics(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostTopicArgs']]]]):
        pulumi.set(self, "vhost_topics", value)

    @property
    @pulumi.getter
    def vhosts(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]]:
        """
        Specifies a map of virtual hosts to permissions.
        """
        return pulumi.get(self, "vhosts")

    @vhosts.setter
    def vhosts(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SecretBackendRoleVhostArgs']]]]):
        pulumi.set(self, "vhosts", value)


class SecretBackendRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[str]] = None,
                 vhost_topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostTopicArgs']]]]] = None,
                 vhosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostArgs']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        rabbitmq = vault.rabbit_mq.SecretBackend("rabbitmq",
            connection_uri="https://.....",
            username="user",
            password="password")
        role = vault.rabbit_mq.SecretBackendRole("role",
            backend=rabbitmq.path,
            tags="tag1,tag2",
            vhosts=[vault.rabbit_mq.SecretBackendRoleVhostArgs(
                host="/",
                configure="",
                read=".*",
                write="",
            )],
            vhost_topics=[vault.rabbit_mq.SecretBackendRoleVhostTopicArgs(
                vhosts=[vault.rabbit_mq.SecretBackendRoleVhostTopicVhostArgs(
                    topic="amq.topic",
                    read=".*",
                    write="",
                )],
                host="/",
            )])
        ```

        ## Import

        RabbitMQ secret backend roles can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:rabbitMq/secretBackendRole:SecretBackendRole role rabbitmq/roles/deploy
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostTopicArgs']]]] vhost_topics: Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostArgs']]]] vhosts: Specifies a map of virtual hosts to permissions.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        rabbitmq = vault.rabbit_mq.SecretBackend("rabbitmq",
            connection_uri="https://.....",
            username="user",
            password="password")
        role = vault.rabbit_mq.SecretBackendRole("role",
            backend=rabbitmq.path,
            tags="tag1,tag2",
            vhosts=[vault.rabbit_mq.SecretBackendRoleVhostArgs(
                host="/",
                configure="",
                read=".*",
                write="",
            )],
            vhost_topics=[vault.rabbit_mq.SecretBackendRoleVhostTopicArgs(
                vhosts=[vault.rabbit_mq.SecretBackendRoleVhostTopicVhostArgs(
                    topic="amq.topic",
                    read=".*",
                    write="",
                )],
                host="/",
            )])
        ```

        ## Import

        RabbitMQ secret backend roles can be imported using the `path`, e.g.

        ```sh
        $ pulumi import vault:rabbitMq/secretBackendRole:SecretBackendRole role rabbitmq/roles/deploy
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[str]] = None,
                 vhost_topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostTopicArgs']]]]] = None,
                 vhosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostArgs']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretBackendRoleArgs.__new__(SecretBackendRoleArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["tags"] = tags
            __props__.__dict__["vhost_topics"] = vhost_topics
            __props__.__dict__["vhosts"] = vhosts
        super(SecretBackendRole, __self__).__init__(
            'vault:rabbitMq/secretBackendRole:SecretBackendRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[str]] = None,
            vhost_topics: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostTopicArgs']]]]] = None,
            vhosts: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostArgs']]]]] = None) -> 'SecretBackendRole':
        """
        Get an existing SecretBackendRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The path the RabbitMQ secret backend is mounted at,
               with no leading or trailing `/`s.
        :param pulumi.Input[str] name: The name to identify this role within the backend.
               Must be unique within the backend.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] tags: Specifies a comma-separated RabbitMQ management tags.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostTopicArgs']]]] vhost_topics: Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['SecretBackendRoleVhostArgs']]]] vhosts: Specifies a map of virtual hosts to permissions.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendRoleState.__new__(_SecretBackendRoleState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["name"] = name
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["tags"] = tags
        __props__.__dict__["vhost_topics"] = vhost_topics
        __props__.__dict__["vhosts"] = vhosts
        return SecretBackendRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The path the RabbitMQ secret backend is mounted at,
        with no leading or trailing `/`s.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name to identify this role within the backend.
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
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a comma-separated RabbitMQ management tags.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="vhostTopics")
    def vhost_topics(self) -> pulumi.Output[Optional[Sequence['outputs.SecretBackendRoleVhostTopic']]]:
        """
        Specifies a map of virtual hosts and exchanges to topic permissions. This option requires RabbitMQ 3.7.0 or later.
        """
        return pulumi.get(self, "vhost_topics")

    @property
    @pulumi.getter
    def vhosts(self) -> pulumi.Output[Optional[Sequence['outputs.SecretBackendRoleVhost']]]:
        """
        Specifies a map of virtual hosts to permissions.
        """
        return pulumi.get(self, "vhosts")

