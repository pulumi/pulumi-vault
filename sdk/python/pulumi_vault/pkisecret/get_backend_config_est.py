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

__all__ = [
    'GetBackendConfigEstResult',
    'AwaitableGetBackendConfigEstResult',
    'get_backend_config_est',
    'get_backend_config_est_output',
]

@pulumi.output_type
class GetBackendConfigEstResult:
    """
    A collection of values returned by getBackendConfigEst.
    """
    def __init__(__self__, audit_fields=None, authenticators=None, backend=None, default_mount=None, default_path_policy=None, enable_sentinel_parsing=None, enabled=None, id=None, label_to_path_policy=None, last_updated=None, namespace=None):
        if audit_fields and not isinstance(audit_fields, list):
            raise TypeError("Expected argument 'audit_fields' to be a list")
        pulumi.set(__self__, "audit_fields", audit_fields)
        if authenticators and not isinstance(authenticators, list):
            raise TypeError("Expected argument 'authenticators' to be a list")
        pulumi.set(__self__, "authenticators", authenticators)
        if backend and not isinstance(backend, str):
            raise TypeError("Expected argument 'backend' to be a str")
        pulumi.set(__self__, "backend", backend)
        if default_mount and not isinstance(default_mount, bool):
            raise TypeError("Expected argument 'default_mount' to be a bool")
        pulumi.set(__self__, "default_mount", default_mount)
        if default_path_policy and not isinstance(default_path_policy, str):
            raise TypeError("Expected argument 'default_path_policy' to be a str")
        pulumi.set(__self__, "default_path_policy", default_path_policy)
        if enable_sentinel_parsing and not isinstance(enable_sentinel_parsing, bool):
            raise TypeError("Expected argument 'enable_sentinel_parsing' to be a bool")
        pulumi.set(__self__, "enable_sentinel_parsing", enable_sentinel_parsing)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if label_to_path_policy and not isinstance(label_to_path_policy, dict):
            raise TypeError("Expected argument 'label_to_path_policy' to be a dict")
        pulumi.set(__self__, "label_to_path_policy", label_to_path_policy)
        if last_updated and not isinstance(last_updated, str):
            raise TypeError("Expected argument 'last_updated' to be a str")
        pulumi.set(__self__, "last_updated", last_updated)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)

    @property
    @pulumi.getter(name="auditFields")
    def audit_fields(self) -> Sequence[str]:
        """
        Fields parsed from the CSR that appear in the audit and can be used by sentinel policies.
        """
        return pulumi.get(self, "audit_fields")

    @property
    @pulumi.getter
    def authenticators(self) -> Sequence['outputs.GetBackendConfigEstAuthenticatorResult']:
        """
        Lists the mount accessors EST should delegate authentication requests towards (see below for nested schema).
        """
        return pulumi.get(self, "authenticators")

    @property
    @pulumi.getter
    def backend(self) -> str:
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="defaultMount")
    def default_mount(self) -> bool:
        """
        If set, this mount is registered as the default `.well-known/est` URL path. Only a single mount can enable this across a Vault cluster.
        """
        return pulumi.get(self, "default_mount")

    @property
    @pulumi.getter(name="defaultPathPolicy")
    def default_path_policy(self) -> str:
        """
        Required to be set if default_mount is enabled. Specifies the behavior for requests using the default EST label. Can be sign-verbatim or a role given by role:<role_name>.
        """
        return pulumi.get(self, "default_path_policy")

    @property
    @pulumi.getter(name="enableSentinelParsing")
    def enable_sentinel_parsing(self) -> bool:
        """
        If set, parse out fields from the provided CSR making them available for Sentinel policies.
        """
        return pulumi.get(self, "enable_sentinel_parsing")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        Specifies whether EST is enabled.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="labelToPathPolicy")
    def label_to_path_policy(self) -> Mapping[str, Any]:
        """
        A pairing of an EST label with the redirected behavior for requests hitting that role. The path policy can be sign-verbatim or a role given by role:<role_name>. Labels must be unique across Vault cluster, and will register .well-known/est/<label> URL paths.
        """
        return pulumi.get(self, "label_to_path_policy")

    @property
    @pulumi.getter(name="lastUpdated")
    def last_updated(self) -> str:
        """
        A read-only timestamp representing the last time the configuration was updated.
        """
        return pulumi.get(self, "last_updated")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")


class AwaitableGetBackendConfigEstResult(GetBackendConfigEstResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBackendConfigEstResult(
            audit_fields=self.audit_fields,
            authenticators=self.authenticators,
            backend=self.backend,
            default_mount=self.default_mount,
            default_path_policy=self.default_path_policy,
            enable_sentinel_parsing=self.enable_sentinel_parsing,
            enabled=self.enabled,
            id=self.id,
            label_to_path_policy=self.label_to_path_policy,
            last_updated=self.last_updated,
            namespace=self.namespace)


def get_backend_config_est(backend: Optional[str] = None,
                           namespace: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBackendConfigEstResult:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    est_config = vault.pkiSecret.get_backend_config_est_output(backend=pki.path)
    ```


    :param str backend: The path to the PKI secret backend to
           read the EST configuration from, with no leading or trailing `/`s.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['backend'] = backend
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:pkiSecret/getBackendConfigEst:getBackendConfigEst', __args__, opts=opts, typ=GetBackendConfigEstResult).value

    return AwaitableGetBackendConfigEstResult(
        audit_fields=pulumi.get(__ret__, 'audit_fields'),
        authenticators=pulumi.get(__ret__, 'authenticators'),
        backend=pulumi.get(__ret__, 'backend'),
        default_mount=pulumi.get(__ret__, 'default_mount'),
        default_path_policy=pulumi.get(__ret__, 'default_path_policy'),
        enable_sentinel_parsing=pulumi.get(__ret__, 'enable_sentinel_parsing'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        label_to_path_policy=pulumi.get(__ret__, 'label_to_path_policy'),
        last_updated=pulumi.get(__ret__, 'last_updated'),
        namespace=pulumi.get(__ret__, 'namespace'))


@_utilities.lift_output_func(get_backend_config_est)
def get_backend_config_est_output(backend: Optional[pulumi.Input[str]] = None,
                                  namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBackendConfigEstResult]:
    """
    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    pki = vault.Mount("pki",
        path="pki",
        type="pki",
        description="PKI secret engine mount")
    est_config = vault.pkiSecret.get_backend_config_est_output(backend=pki.path)
    ```


    :param str backend: The path to the PKI secret backend to
           read the EST configuration from, with no leading or trailing `/`s.
    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
           *Available only for Vault Enterprise*.
    """
    ...
