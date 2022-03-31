# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretBackendArgs', 'SecretBackend']

@pulumi.input_type
class SecretBackendArgs:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 username_template: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretBackend resource.
        :param pulumi.Input[str] access_key: The AWS Access Key ID this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[str] iam_endpoint: Specifies a custom HTTP IAM endpoint to use.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `aws`.
        :param pulumi.Input[str] region: The AWS region for API calls. Defaults to `us-east-1`.
        :param pulumi.Input[str] secret_key: The AWS Secret Key this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[str] sts_endpoint: Specifies a custom HTTP STS endpoint to use.
        :param pulumi.Input[str] username_template: Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if iam_endpoint is not None:
            pulumi.set(__self__, "iam_endpoint", iam_endpoint)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sts_endpoint is not None:
            pulumi.set(__self__, "sts_endpoint", sts_endpoint)
        if username_template is not None:
            pulumi.set(__self__, "username_template", username_template)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Access Key ID this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom HTTP IAM endpoint to use.
        """
        return pulumi.get(self, "iam_endpoint")

    @iam_endpoint.setter
    def iam_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_endpoint", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `aws`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region for API calls. Defaults to `us-east-1`.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Secret Key this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom HTTP STS endpoint to use.
        """
        return pulumi.get(self, "sts_endpoint")

    @sts_endpoint.setter
    def sts_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_endpoint", value)

    @property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        return pulumi.get(self, "username_template")

    @username_template.setter
    def username_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_template", value)


@pulumi.input_type
class _SecretBackendState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 username_template: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretBackend resources.
        :param pulumi.Input[str] access_key: The AWS Access Key ID this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[str] iam_endpoint: Specifies a custom HTTP IAM endpoint to use.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `aws`.
        :param pulumi.Input[str] region: The AWS region for API calls. Defaults to `us-east-1`.
        :param pulumi.Input[str] secret_key: The AWS Secret Key this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[str] sts_endpoint: Specifies a custom HTTP STS endpoint to use.
        :param pulumi.Input[str] username_template: Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if default_lease_ttl_seconds is not None:
            pulumi.set(__self__, "default_lease_ttl_seconds", default_lease_ttl_seconds)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if iam_endpoint is not None:
            pulumi.set(__self__, "iam_endpoint", iam_endpoint)
        if max_lease_ttl_seconds is not None:
            pulumi.set(__self__, "max_lease_ttl_seconds", max_lease_ttl_seconds)
        if path is not None:
            pulumi.set(__self__, "path", path)
        if region is not None:
            pulumi.set(__self__, "region", region)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sts_endpoint is not None:
            pulumi.set(__self__, "sts_endpoint", sts_endpoint)
        if username_template is not None:
            pulumi.set(__self__, "username_template", username_template)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Access Key ID this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @default_lease_ttl_seconds.setter
    def default_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "default_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom HTTP IAM endpoint to use.
        """
        return pulumi.get(self, "iam_endpoint")

    @iam_endpoint.setter
    def iam_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_endpoint", value)

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> Optional[pulumi.Input[int]]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @max_lease_ttl_seconds.setter
    def max_lease_ttl_seconds(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_lease_ttl_seconds", value)

    @property
    @pulumi.getter
    def path(self) -> Optional[pulumi.Input[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `aws`.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS region for API calls. Defaults to `us-east-1`.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS Secret Key this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies a custom HTTP STS endpoint to use.
        """
        return pulumi.get(self, "sts_endpoint")

    @sts_endpoint.setter
    def sts_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_endpoint", value)

    @property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> Optional[pulumi.Input[str]]:
        """
        Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        return pulumi.get(self, "username_template")

    @username_template.setter
    def username_template(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "username_template", value)


class SecretBackend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 username_template: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        aws = vault.aws.SecretBackend("aws",
            access_key="AKIA.....",
            secret_key="AWS secret key")
        ```

        ## Import

        AWS secret backends can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS Access Key ID this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[str] iam_endpoint: Specifies a custom HTTP IAM endpoint to use.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `aws`.
        :param pulumi.Input[str] region: The AWS region for API calls. Defaults to `us-east-1`.
        :param pulumi.Input[str] secret_key: The AWS Secret Key this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[str] sts_endpoint: Specifies a custom HTTP STS endpoint to use.
        :param pulumi.Input[str] username_template: Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[SecretBackendArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        aws = vault.aws.SecretBackend("aws",
            access_key="AKIA.....",
            secret_key="AWS secret key")
        ```

        ## Import

        AWS secret backends can be imported using the `path`, e.g.

        ```sh
         $ pulumi import vault:aws/secretBackend:SecretBackend aws aws
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
                 path: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 username_template: Optional[pulumi.Input[str]] = None,
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
            __props__ = SecretBackendArgs.__new__(SecretBackendArgs)

            __props__.__dict__["access_key"] = access_key
            __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
            __props__.__dict__["description"] = description
            __props__.__dict__["iam_endpoint"] = iam_endpoint
            __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
            __props__.__dict__["path"] = path
            __props__.__dict__["region"] = region
            __props__.__dict__["secret_key"] = secret_key
            __props__.__dict__["sts_endpoint"] = sts_endpoint
            __props__.__dict__["username_template"] = username_template
        super(SecretBackend, __self__).__init__(
            'vault:aws/secretBackend:SecretBackend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            default_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            iam_endpoint: Optional[pulumi.Input[str]] = None,
            max_lease_ttl_seconds: Optional[pulumi.Input[int]] = None,
            path: Optional[pulumi.Input[str]] = None,
            region: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            sts_endpoint: Optional[pulumi.Input[str]] = None,
            username_template: Optional[pulumi.Input[str]] = None) -> 'SecretBackend':
        """
        Get an existing SecretBackend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS Access Key ID this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[int] default_lease_ttl_seconds: The default TTL for credentials
               issued by this backend.
        :param pulumi.Input[str] description: A human-friendly description for this backend.
        :param pulumi.Input[str] iam_endpoint: Specifies a custom HTTP IAM endpoint to use.
        :param pulumi.Input[int] max_lease_ttl_seconds: The maximum TTL that can be requested
               for credentials issued by this backend.
        :param pulumi.Input[str] path: The unique path this backend should be mounted at. Must
               not begin or end with a `/`. Defaults to `aws`.
        :param pulumi.Input[str] region: The AWS region for API calls. Defaults to `us-east-1`.
        :param pulumi.Input[str] secret_key: The AWS Secret Key this backend should use to
               issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        :param pulumi.Input[str] sts_endpoint: Specifies a custom HTTP STS endpoint to use.
        :param pulumi.Input[str] username_template: Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretBackendState.__new__(_SecretBackendState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["default_lease_ttl_seconds"] = default_lease_ttl_seconds
        __props__.__dict__["description"] = description
        __props__.__dict__["iam_endpoint"] = iam_endpoint
        __props__.__dict__["max_lease_ttl_seconds"] = max_lease_ttl_seconds
        __props__.__dict__["path"] = path
        __props__.__dict__["region"] = region
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["sts_endpoint"] = sts_endpoint
        __props__.__dict__["username_template"] = username_template
        return SecretBackend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Access Key ID this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter(name="defaultLeaseTtlSeconds")
    def default_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        The default TTL for credentials
        issued by this backend.
        """
        return pulumi.get(self, "default_lease_ttl_seconds")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A human-friendly description for this backend.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a custom HTTP IAM endpoint to use.
        """
        return pulumi.get(self, "iam_endpoint")

    @property
    @pulumi.getter(name="maxLeaseTtlSeconds")
    def max_lease_ttl_seconds(self) -> pulumi.Output[int]:
        """
        The maximum TTL that can be requested
        for credentials issued by this backend.
        """
        return pulumi.get(self, "max_lease_ttl_seconds")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[Optional[str]]:
        """
        The unique path this backend should be mounted at. Must
        not begin or end with a `/`. Defaults to `aws`.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[str]:
        """
        The AWS region for API calls. Defaults to `us-east-1`.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS Secret Key this backend should use to
        issue new credentials. Vault uses the official AWS SDK to authenticate, and thus can also use standard AWS environment credentials, shared file credentials or IAM role/ECS task credentials.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Specifies a custom HTTP STS endpoint to use.
        """
        return pulumi.get(self, "sts_endpoint")

    @property
    @pulumi.getter(name="usernameTemplate")
    def username_template(self) -> pulumi.Output[str]:
        """
        Template describing how dynamic usernames are generated. The username template is used to generate both IAM usernames (capped at 64 characters) and STS usernames (capped at 32 characters). If no template is provided the field defaults to the template:
        """
        return pulumi.get(self, "username_template")

