# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['AuthBackendClientArgs', 'AuthBackendClient']

@pulumi.input_type
class AuthBackendClientArgs:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 ec2_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_server_id_header_value: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 sts_region: Optional[pulumi.Input[str]] = None,
                 use_sts_region_from_client: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a AuthBackendClient resource.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
        :param pulumi.Input[str] sts_region: Override the default region when making STS API 
               calls. The `sts_endpoint` argument must be set when using `sts_region`.
        :param pulumi.Input[bool] use_sts_region_from_client: Available in Vault v1.15+. If set, 
               overrides both `sts_endpoint` and `sts_region` to instead use the region
               specified in the client request headers for IAM-based authentication.
               This can be useful when you have client requests coming from different
               regions and want flexibility in which regional STS API is used.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if ec2_endpoint is not None:
            pulumi.set(__self__, "ec2_endpoint", ec2_endpoint)
        if iam_endpoint is not None:
            pulumi.set(__self__, "iam_endpoint", iam_endpoint)
        if iam_server_id_header_value is not None:
            pulumi.set(__self__, "iam_server_id_header_value", iam_server_id_header_value)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sts_endpoint is not None:
            pulumi.set(__self__, "sts_endpoint", sts_endpoint)
        if sts_region is not None:
            pulumi.set(__self__, "sts_region", sts_region)
        if use_sts_region_from_client is not None:
            pulumi.set(__self__, "use_sts_region_from_client", use_sts_region_from_client)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS access key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the AWS auth backend being configured was
        mounted at.  Defaults to `aws`.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="ec2Endpoint")
    def ec2_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making EC2 API
        calls.
        """
        return pulumi.get(self, "ec2_endpoint")

    @ec2_endpoint.setter
    def ec2_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ec2_endpoint", value)

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making IAM API
        calls.
        """
        return pulumi.get(self, "iam_endpoint")

    @iam_endpoint.setter
    def iam_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_endpoint", value)

    @property
    @pulumi.getter(name="iamServerIdHeaderValue")
    def iam_server_id_header_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value to require in the
        `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        that are used in the IAM auth method.
        """
        return pulumi.get(self, "iam_server_id_header_value")

    @iam_server_id_header_value.setter
    def iam_server_id_header_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_server_id_header_value", value)

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
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS secret key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making STS API
        calls.
        """
        return pulumi.get(self, "sts_endpoint")

    @sts_endpoint.setter
    def sts_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_endpoint", value)

    @property
    @pulumi.getter(name="stsRegion")
    def sts_region(self) -> Optional[pulumi.Input[str]]:
        """
        Override the default region when making STS API 
        calls. The `sts_endpoint` argument must be set when using `sts_region`.
        """
        return pulumi.get(self, "sts_region")

    @sts_region.setter
    def sts_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_region", value)

    @property
    @pulumi.getter(name="useStsRegionFromClient")
    def use_sts_region_from_client(self) -> Optional[pulumi.Input[bool]]:
        """
        Available in Vault v1.15+. If set, 
        overrides both `sts_endpoint` and `sts_region` to instead use the region
        specified in the client request headers for IAM-based authentication.
        This can be useful when you have client requests coming from different
        regions and want flexibility in which regional STS API is used.
        """
        return pulumi.get(self, "use_sts_region_from_client")

    @use_sts_region_from_client.setter
    def use_sts_region_from_client(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_sts_region_from_client", value)


@pulumi.input_type
class _AuthBackendClientState:
    def __init__(__self__, *,
                 access_key: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 ec2_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_server_id_header_value: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 sts_region: Optional[pulumi.Input[str]] = None,
                 use_sts_region_from_client: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering AuthBackendClient resources.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
        :param pulumi.Input[str] sts_region: Override the default region when making STS API 
               calls. The `sts_endpoint` argument must be set when using `sts_region`.
        :param pulumi.Input[bool] use_sts_region_from_client: Available in Vault v1.15+. If set, 
               overrides both `sts_endpoint` and `sts_region` to instead use the region
               specified in the client request headers for IAM-based authentication.
               This can be useful when you have client requests coming from different
               regions and want flexibility in which regional STS API is used.
        """
        if access_key is not None:
            pulumi.set(__self__, "access_key", access_key)
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if ec2_endpoint is not None:
            pulumi.set(__self__, "ec2_endpoint", ec2_endpoint)
        if iam_endpoint is not None:
            pulumi.set(__self__, "iam_endpoint", iam_endpoint)
        if iam_server_id_header_value is not None:
            pulumi.set(__self__, "iam_server_id_header_value", iam_server_id_header_value)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if secret_key is not None:
            pulumi.set(__self__, "secret_key", secret_key)
        if sts_endpoint is not None:
            pulumi.set(__self__, "sts_endpoint", sts_endpoint)
        if sts_region is not None:
            pulumi.set(__self__, "sts_region", sts_region)
        if use_sts_region_from_client is not None:
            pulumi.set(__self__, "use_sts_region_from_client", use_sts_region_from_client)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS access key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "access_key")

    @access_key.setter
    def access_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "access_key", value)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        """
        The path the AWS auth backend being configured was
        mounted at.  Defaults to `aws`.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="ec2Endpoint")
    def ec2_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making EC2 API
        calls.
        """
        return pulumi.get(self, "ec2_endpoint")

    @ec2_endpoint.setter
    def ec2_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ec2_endpoint", value)

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making IAM API
        calls.
        """
        return pulumi.get(self, "iam_endpoint")

    @iam_endpoint.setter
    def iam_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_endpoint", value)

    @property
    @pulumi.getter(name="iamServerIdHeaderValue")
    def iam_server_id_header_value(self) -> Optional[pulumi.Input[str]]:
        """
        The value to require in the
        `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        that are used in the IAM auth method.
        """
        return pulumi.get(self, "iam_server_id_header_value")

    @iam_server_id_header_value.setter
    def iam_server_id_header_value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "iam_server_id_header_value", value)

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
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> Optional[pulumi.Input[str]]:
        """
        The AWS secret key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "secret_key")

    @secret_key.setter
    def secret_key(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "secret_key", value)

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> Optional[pulumi.Input[str]]:
        """
        Override the URL Vault uses when making STS API
        calls.
        """
        return pulumi.get(self, "sts_endpoint")

    @sts_endpoint.setter
    def sts_endpoint(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_endpoint", value)

    @property
    @pulumi.getter(name="stsRegion")
    def sts_region(self) -> Optional[pulumi.Input[str]]:
        """
        Override the default region when making STS API 
        calls. The `sts_endpoint` argument must be set when using `sts_region`.
        """
        return pulumi.get(self, "sts_region")

    @sts_region.setter
    def sts_region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "sts_region", value)

    @property
    @pulumi.getter(name="useStsRegionFromClient")
    def use_sts_region_from_client(self) -> Optional[pulumi.Input[bool]]:
        """
        Available in Vault v1.15+. If set, 
        overrides both `sts_endpoint` and `sts_region` to instead use the region
        specified in the client request headers for IAM-based authentication.
        This can be useful when you have client requests coming from different
        regions and want flexibility in which regional STS API is used.
        """
        return pulumi.get(self, "use_sts_region_from_client")

    @use_sts_region_from_client.setter
    def use_sts_region_from_client(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "use_sts_region_from_client", value)


class AuthBackendClient(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 ec2_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_server_id_header_value: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 sts_region: Optional[pulumi.Input[str]] = None,
                 use_sts_region_from_client: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.AuthBackend("example", type="aws")
        example_auth_backend_client = vault.aws.AuthBackendClient("example",
            backend=example.path,
            access_key="INSERT_AWS_ACCESS_KEY",
            secret_key="INSERT_AWS_SECRET_KEY")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.

        ```sh
        $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
        :param pulumi.Input[str] sts_region: Override the default region when making STS API 
               calls. The `sts_endpoint` argument must be set when using `sts_region`.
        :param pulumi.Input[bool] use_sts_region_from_client: Available in Vault v1.15+. If set, 
               overrides both `sts_endpoint` and `sts_region` to instead use the region
               specified in the client request headers for IAM-based authentication.
               This can be useful when you have client requests coming from different
               regions and want flexibility in which regional STS API is used.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AuthBackendClientArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        <!--Start PulumiCodeChooser -->
        ```python
        import pulumi
        import pulumi_vault as vault

        example = vault.AuthBackend("example", type="aws")
        example_auth_backend_client = vault.aws.AuthBackendClient("example",
            backend=example.path,
            access_key="INSERT_AWS_ACCESS_KEY",
            secret_key="INSERT_AWS_SECRET_KEY")
        ```
        <!--End PulumiCodeChooser -->

        ## Import

        AWS auth backend clients can be imported using `auth/`, the `backend` path, and `/config/client` e.g.

        ```sh
        $ pulumi import vault:aws/authBackendClient:AuthBackendClient example auth/aws/config/client
        ```

        :param str resource_name: The name of the resource.
        :param AuthBackendClientArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AuthBackendClientArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 access_key: Optional[pulumi.Input[str]] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 ec2_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_endpoint: Optional[pulumi.Input[str]] = None,
                 iam_server_id_header_value: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 secret_key: Optional[pulumi.Input[str]] = None,
                 sts_endpoint: Optional[pulumi.Input[str]] = None,
                 sts_region: Optional[pulumi.Input[str]] = None,
                 use_sts_region_from_client: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AuthBackendClientArgs.__new__(AuthBackendClientArgs)

            __props__.__dict__["access_key"] = None if access_key is None else pulumi.Output.secret(access_key)
            __props__.__dict__["backend"] = backend
            __props__.__dict__["ec2_endpoint"] = ec2_endpoint
            __props__.__dict__["iam_endpoint"] = iam_endpoint
            __props__.__dict__["iam_server_id_header_value"] = iam_server_id_header_value
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["secret_key"] = None if secret_key is None else pulumi.Output.secret(secret_key)
            __props__.__dict__["sts_endpoint"] = sts_endpoint
            __props__.__dict__["sts_region"] = sts_region
            __props__.__dict__["use_sts_region_from_client"] = use_sts_region_from_client
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["accessKey", "secretKey"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(AuthBackendClient, __self__).__init__(
            'vault:aws/authBackendClient:AuthBackendClient',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            access_key: Optional[pulumi.Input[str]] = None,
            backend: Optional[pulumi.Input[str]] = None,
            ec2_endpoint: Optional[pulumi.Input[str]] = None,
            iam_endpoint: Optional[pulumi.Input[str]] = None,
            iam_server_id_header_value: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            secret_key: Optional[pulumi.Input[str]] = None,
            sts_endpoint: Optional[pulumi.Input[str]] = None,
            sts_region: Optional[pulumi.Input[str]] = None,
            use_sts_region_from_client: Optional[pulumi.Input[bool]] = None) -> 'AuthBackendClient':
        """
        Get an existing AuthBackendClient resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] access_key: The AWS access key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] backend: The path the AWS auth backend being configured was
               mounted at.  Defaults to `aws`.
        :param pulumi.Input[str] ec2_endpoint: Override the URL Vault uses when making EC2 API
               calls.
        :param pulumi.Input[str] iam_endpoint: Override the URL Vault uses when making IAM API
               calls.
        :param pulumi.Input[str] iam_server_id_header_value: The value to require in the
               `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
               that are used in the IAM auth method.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] secret_key: The AWS secret key that Vault should use for the
               auth backend.
        :param pulumi.Input[str] sts_endpoint: Override the URL Vault uses when making STS API
               calls.
        :param pulumi.Input[str] sts_region: Override the default region when making STS API 
               calls. The `sts_endpoint` argument must be set when using `sts_region`.
        :param pulumi.Input[bool] use_sts_region_from_client: Available in Vault v1.15+. If set, 
               overrides both `sts_endpoint` and `sts_region` to instead use the region
               specified in the client request headers for IAM-based authentication.
               This can be useful when you have client requests coming from different
               regions and want flexibility in which regional STS API is used.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AuthBackendClientState.__new__(_AuthBackendClientState)

        __props__.__dict__["access_key"] = access_key
        __props__.__dict__["backend"] = backend
        __props__.__dict__["ec2_endpoint"] = ec2_endpoint
        __props__.__dict__["iam_endpoint"] = iam_endpoint
        __props__.__dict__["iam_server_id_header_value"] = iam_server_id_header_value
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["secret_key"] = secret_key
        __props__.__dict__["sts_endpoint"] = sts_endpoint
        __props__.__dict__["sts_region"] = sts_region
        __props__.__dict__["use_sts_region_from_client"] = use_sts_region_from_client
        return AuthBackendClient(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessKey")
    def access_key(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS access key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "access_key")

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[Optional[str]]:
        """
        The path the AWS auth backend being configured was
        mounted at.  Defaults to `aws`.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="ec2Endpoint")
    def ec2_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Override the URL Vault uses when making EC2 API
        calls.
        """
        return pulumi.get(self, "ec2_endpoint")

    @property
    @pulumi.getter(name="iamEndpoint")
    def iam_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Override the URL Vault uses when making IAM API
        calls.
        """
        return pulumi.get(self, "iam_endpoint")

    @property
    @pulumi.getter(name="iamServerIdHeaderValue")
    def iam_server_id_header_value(self) -> pulumi.Output[Optional[str]]:
        """
        The value to require in the
        `X-Vault-AWS-IAM-Server-ID` header as part of `GetCallerIdentity` requests
        that are used in the IAM auth method.
        """
        return pulumi.get(self, "iam_server_id_header_value")

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
    @pulumi.getter(name="secretKey")
    def secret_key(self) -> pulumi.Output[Optional[str]]:
        """
        The AWS secret key that Vault should use for the
        auth backend.
        """
        return pulumi.get(self, "secret_key")

    @property
    @pulumi.getter(name="stsEndpoint")
    def sts_endpoint(self) -> pulumi.Output[Optional[str]]:
        """
        Override the URL Vault uses when making STS API
        calls.
        """
        return pulumi.get(self, "sts_endpoint")

    @property
    @pulumi.getter(name="stsRegion")
    def sts_region(self) -> pulumi.Output[Optional[str]]:
        """
        Override the default region when making STS API 
        calls. The `sts_endpoint` argument must be set when using `sts_region`.
        """
        return pulumi.get(self, "sts_region")

    @property
    @pulumi.getter(name="useStsRegionFromClient")
    def use_sts_region_from_client(self) -> pulumi.Output[bool]:
        """
        Available in Vault v1.15+. If set, 
        overrides both `sts_endpoint` and `sts_region` to instead use the region
        specified in the client request headers for IAM-based authentication.
        This can be useful when you have client requests coming from different
        regions and want flexibility in which regional STS API is used.
        """
        return pulumi.get(self, "use_sts_region_from_client")

