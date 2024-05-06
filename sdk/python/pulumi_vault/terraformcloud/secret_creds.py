# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = ['SecretCredsArgs', 'SecretCreds']

@pulumi.input_type
class SecretCredsArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 role: pulumi.Input[str],
                 namespace: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SecretCreds resource.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "role", role)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter
    def role(self) -> pulumi.Input[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: pulumi.Input[str]):
        pulumi.set(self, "role", value)

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


@pulumi.input_type
class _SecretCredsState:
    def __init__(__self__, *,
                 backend: Optional[pulumi.Input[str]] = None,
                 lease_id: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 organization: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 team_id: Optional[pulumi.Input[str]] = None,
                 token: Optional[pulumi.Input[str]] = None,
                 token_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering SecretCreds resources.
        :param pulumi.Input[str] lease_id: The lease associated with the token. Only user tokens will have a 
               Vault lease associated with them.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization: The organization associated with the token provided.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] team_id: The team id associated with the token provided.
        :param pulumi.Input[str] token: The actual token that was generated and can be used with API calls
               to identify the user of the call.
        :param pulumi.Input[str] token_id: The public identifier for a specific token. It can be used 
               to look up information about a token or to revoke a token.
        """
        if backend is not None:
            pulumi.set(__self__, "backend", backend)
        if lease_id is not None:
            pulumi.set(__self__, "lease_id", lease_id)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if organization is not None:
            pulumi.set(__self__, "organization", organization)
        if role is not None:
            pulumi.set(__self__, "role", role)
        if team_id is not None:
            pulumi.set(__self__, "team_id", team_id)
        if token is not None:
            pulumi.set(__self__, "token", token)
        if token_id is not None:
            pulumi.set(__self__, "token_id", token_id)

    @property
    @pulumi.getter
    def backend(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> Optional[pulumi.Input[str]]:
        """
        The lease associated with the token. Only user tokens will have a 
        Vault lease associated with them.
        """
        return pulumi.get(self, "lease_id")

    @lease_id.setter
    def lease_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "lease_id", value)

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
    def organization(self) -> Optional[pulumi.Input[str]]:
        """
        The organization associated with the token provided.
        """
        return pulumi.get(self, "organization")

    @organization.setter
    def organization(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "organization", value)

    @property
    @pulumi.getter
    def role(self) -> Optional[pulumi.Input[str]]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role")

    @role.setter
    def role(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "role", value)

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> Optional[pulumi.Input[str]]:
        """
        The team id associated with the token provided.
        """
        return pulumi.get(self, "team_id")

    @team_id.setter
    def team_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "team_id", value)

    @property
    @pulumi.getter
    def token(self) -> Optional[pulumi.Input[str]]:
        """
        The actual token that was generated and can be used with API calls
        to identify the user of the call.
        """
        return pulumi.get(self, "token")

    @token.setter
    def token(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token", value)

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> Optional[pulumi.Input[str]]:
        """
        The public identifier for a specific token. It can be used 
        to look up information about a token or to revoke a token.
        """
        return pulumi.get(self, "token_id")

    @token_id.setter
    def token_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "token_id", value)


class SecretCreds(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.terraformcloud.SecretBackend("test",
            backend="terraform",
            description="Manages the Terraform Cloud backend",
            token="V0idfhi2iksSDU234ucdbi2nidsi...")
        example = vault.terraformcloud.SecretRole("example",
            backend=test.backend,
            name="test-role",
            organization="example-organization-name",
            team_id="team-ieF4isC...")
        token = vault.terraformcloud.SecretCreds("token",
            backend=test.backend,
            role=example.name)
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] role: Name of the role.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretCredsArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        test = vault.terraformcloud.SecretBackend("test",
            backend="terraform",
            description="Manages the Terraform Cloud backend",
            token="V0idfhi2iksSDU234ucdbi2nidsi...")
        example = vault.terraformcloud.SecretRole("example",
            backend=test.backend,
            name="test-role",
            organization="example-organization-name",
            team_id="team-ieF4isC...")
        token = vault.terraformcloud.SecretCreds("token",
            backend=test.backend,
            role=example.name)
        ```

        :param str resource_name: The name of the resource.
        :param SecretCredsArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretCredsArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 role: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SecretCredsArgs.__new__(SecretCredsArgs)

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__.__dict__["backend"] = backend
            __props__.__dict__["namespace"] = namespace
            if role is None and not opts.urn:
                raise TypeError("Missing required property 'role'")
            __props__.__dict__["role"] = role
            __props__.__dict__["lease_id"] = None
            __props__.__dict__["organization"] = None
            __props__.__dict__["team_id"] = None
            __props__.__dict__["token"] = None
            __props__.__dict__["token_id"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["leaseId", "token"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(SecretCreds, __self__).__init__(
            'vault:terraformcloud/secretCreds:SecretCreds',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            lease_id: Optional[pulumi.Input[str]] = None,
            namespace: Optional[pulumi.Input[str]] = None,
            organization: Optional[pulumi.Input[str]] = None,
            role: Optional[pulumi.Input[str]] = None,
            team_id: Optional[pulumi.Input[str]] = None,
            token: Optional[pulumi.Input[str]] = None,
            token_id: Optional[pulumi.Input[str]] = None) -> 'SecretCreds':
        """
        Get an existing SecretCreds resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] lease_id: The lease associated with the token. Only user tokens will have a 
               Vault lease associated with them.
        :param pulumi.Input[str] namespace: The namespace to provision the resource in.
               The value should not contain leading or trailing forward slashes.
               The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
               *Available only for Vault Enterprise*.
        :param pulumi.Input[str] organization: The organization associated with the token provided.
        :param pulumi.Input[str] role: Name of the role.
        :param pulumi.Input[str] team_id: The team id associated with the token provided.
        :param pulumi.Input[str] token: The actual token that was generated and can be used with API calls
               to identify the user of the call.
        :param pulumi.Input[str] token_id: The public identifier for a specific token. It can be used 
               to look up information about a token or to revoke a token.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SecretCredsState.__new__(_SecretCredsState)

        __props__.__dict__["backend"] = backend
        __props__.__dict__["lease_id"] = lease_id
        __props__.__dict__["namespace"] = namespace
        __props__.__dict__["organization"] = organization
        __props__.__dict__["role"] = role
        __props__.__dict__["team_id"] = team_id
        __props__.__dict__["token"] = token
        __props__.__dict__["token_id"] = token_id
        return SecretCreds(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="leaseId")
    def lease_id(self) -> pulumi.Output[str]:
        """
        The lease associated with the token. Only user tokens will have a 
        Vault lease associated with them.
        """
        return pulumi.get(self, "lease_id")

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
    def organization(self) -> pulumi.Output[str]:
        """
        The organization associated with the token provided.
        """
        return pulumi.get(self, "organization")

    @property
    @pulumi.getter
    def role(self) -> pulumi.Output[str]:
        """
        Name of the role.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="teamId")
    def team_id(self) -> pulumi.Output[str]:
        """
        The team id associated with the token provided.
        """
        return pulumi.get(self, "team_id")

    @property
    @pulumi.getter
    def token(self) -> pulumi.Output[str]:
        """
        The actual token that was generated and can be used with API calls
        to identify the user of the call.
        """
        return pulumi.get(self, "token")

    @property
    @pulumi.getter(name="tokenId")
    def token_id(self) -> pulumi.Output[str]:
        """
        The public identifier for a specific token. It can be used 
        to look up information about a token or to revoke a token.
        """
        return pulumi.get(self, "token_id")

