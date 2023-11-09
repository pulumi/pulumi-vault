# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'GetRaftAutopilotStateResult',
    'AwaitableGetRaftAutopilotStateResult',
    'get_raft_autopilot_state',
    'get_raft_autopilot_state_output',
]

@pulumi.output_type
class GetRaftAutopilotStateResult:
    """
    A collection of values returned by getRaftAutopilotState.
    """
    def __init__(__self__, failure_tolerance=None, healthy=None, id=None, leader=None, namespace=None, optimistic_failure_tolerance=None, redundancy_zones=None, redundancy_zones_json=None, servers=None, servers_json=None, upgrade_info=None, upgrade_info_json=None, voters=None):
        if failure_tolerance and not isinstance(failure_tolerance, int):
            raise TypeError("Expected argument 'failure_tolerance' to be a int")
        pulumi.set(__self__, "failure_tolerance", failure_tolerance)
        if healthy and not isinstance(healthy, bool):
            raise TypeError("Expected argument 'healthy' to be a bool")
        pulumi.set(__self__, "healthy", healthy)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if leader and not isinstance(leader, str):
            raise TypeError("Expected argument 'leader' to be a str")
        pulumi.set(__self__, "leader", leader)
        if namespace and not isinstance(namespace, str):
            raise TypeError("Expected argument 'namespace' to be a str")
        pulumi.set(__self__, "namespace", namespace)
        if optimistic_failure_tolerance and not isinstance(optimistic_failure_tolerance, int):
            raise TypeError("Expected argument 'optimistic_failure_tolerance' to be a int")
        pulumi.set(__self__, "optimistic_failure_tolerance", optimistic_failure_tolerance)
        if redundancy_zones and not isinstance(redundancy_zones, dict):
            raise TypeError("Expected argument 'redundancy_zones' to be a dict")
        pulumi.set(__self__, "redundancy_zones", redundancy_zones)
        if redundancy_zones_json and not isinstance(redundancy_zones_json, str):
            raise TypeError("Expected argument 'redundancy_zones_json' to be a str")
        pulumi.set(__self__, "redundancy_zones_json", redundancy_zones_json)
        if servers and not isinstance(servers, dict):
            raise TypeError("Expected argument 'servers' to be a dict")
        pulumi.set(__self__, "servers", servers)
        if servers_json and not isinstance(servers_json, str):
            raise TypeError("Expected argument 'servers_json' to be a str")
        pulumi.set(__self__, "servers_json", servers_json)
        if upgrade_info and not isinstance(upgrade_info, dict):
            raise TypeError("Expected argument 'upgrade_info' to be a dict")
        pulumi.set(__self__, "upgrade_info", upgrade_info)
        if upgrade_info_json and not isinstance(upgrade_info_json, str):
            raise TypeError("Expected argument 'upgrade_info_json' to be a str")
        pulumi.set(__self__, "upgrade_info_json", upgrade_info_json)
        if voters and not isinstance(voters, list):
            raise TypeError("Expected argument 'voters' to be a list")
        pulumi.set(__self__, "voters", voters)

    @property
    @pulumi.getter(name="failureTolerance")
    def failure_tolerance(self) -> int:
        """
        How many nodes could fail before the cluster becomes unhealthy.
        """
        return pulumi.get(self, "failure_tolerance")

    @property
    @pulumi.getter
    def healthy(self) -> bool:
        """
        Cluster health status.
        """
        return pulumi.get(self, "healthy")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def leader(self) -> str:
        """
        The current leader of Vault.
        """
        return pulumi.get(self, "leader")

    @property
    @pulumi.getter
    def namespace(self) -> Optional[str]:
        return pulumi.get(self, "namespace")

    @property
    @pulumi.getter(name="optimisticFailureTolerance")
    def optimistic_failure_tolerance(self) -> int:
        """
        The cluster-level optimistic failure tolerance.
        """
        return pulumi.get(self, "optimistic_failure_tolerance")

    @property
    @pulumi.getter(name="redundancyZones")
    def redundancy_zones(self) -> Mapping[str, Any]:
        """
        Additional output related to redundancy zones stored as a serialized map of strings.
        """
        return pulumi.get(self, "redundancy_zones")

    @property
    @pulumi.getter(name="redundancyZonesJson")
    def redundancy_zones_json(self) -> str:
        """
        Additional output related to redundancy zones.
        """
        return pulumi.get(self, "redundancy_zones_json")

    @property
    @pulumi.getter
    def servers(self) -> Mapping[str, Any]:
        """
        Additionaly output related to servers in the cluster stored as a serialized map of strings.
        """
        return pulumi.get(self, "servers")

    @property
    @pulumi.getter(name="serversJson")
    def servers_json(self) -> str:
        """
        Additionaly output related to servers in the cluster.
        """
        return pulumi.get(self, "servers_json")

    @property
    @pulumi.getter(name="upgradeInfo")
    def upgrade_info(self) -> Mapping[str, Any]:
        """
        Additional output related to upgrade information stored as a serialized map of strings.
        """
        return pulumi.get(self, "upgrade_info")

    @property
    @pulumi.getter(name="upgradeInfoJson")
    def upgrade_info_json(self) -> str:
        """
        Additional output related to upgrade information.
        """
        return pulumi.get(self, "upgrade_info_json")

    @property
    @pulumi.getter
    def voters(self) -> Sequence[str]:
        """
        The voters in the Vault cluster.
        """
        return pulumi.get(self, "voters")


class AwaitableGetRaftAutopilotStateResult(GetRaftAutopilotStateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRaftAutopilotStateResult(
            failure_tolerance=self.failure_tolerance,
            healthy=self.healthy,
            id=self.id,
            leader=self.leader,
            namespace=self.namespace,
            optimistic_failure_tolerance=self.optimistic_failure_tolerance,
            redundancy_zones=self.redundancy_zones,
            redundancy_zones_json=self.redundancy_zones_json,
            servers=self.servers,
            servers_json=self.servers_json,
            upgrade_info=self.upgrade_info,
            upgrade_info_json=self.upgrade_info_json,
            voters=self.voters)


def get_raft_autopilot_state(namespace: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRaftAutopilotStateResult:
    """
    Displays the state of the raft cluster under integrated storage as seen by
    autopilot. It shows whether autopilot thinks the cluster is healthy or not, and
    how many nodes could fail before the cluster becomes unhealthy ("Failure
    Tolerance"). For more information, please refer to the
    [Vault documentation](https://developer.hashicorp.com/vault/api-docs/system/storage/raftautopilot#get-cluster-state).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    main = vault.get_raft_autopilot_state()
    pulumi.export("failure-tolerance", main.failure_tolerance)
    ```


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    """
    __args__ = dict()
    __args__['namespace'] = namespace
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vault:index/getRaftAutopilotState:getRaftAutopilotState', __args__, opts=opts, typ=GetRaftAutopilotStateResult).value

    return AwaitableGetRaftAutopilotStateResult(
        failure_tolerance=pulumi.get(__ret__, 'failure_tolerance'),
        healthy=pulumi.get(__ret__, 'healthy'),
        id=pulumi.get(__ret__, 'id'),
        leader=pulumi.get(__ret__, 'leader'),
        namespace=pulumi.get(__ret__, 'namespace'),
        optimistic_failure_tolerance=pulumi.get(__ret__, 'optimistic_failure_tolerance'),
        redundancy_zones=pulumi.get(__ret__, 'redundancy_zones'),
        redundancy_zones_json=pulumi.get(__ret__, 'redundancy_zones_json'),
        servers=pulumi.get(__ret__, 'servers'),
        servers_json=pulumi.get(__ret__, 'servers_json'),
        upgrade_info=pulumi.get(__ret__, 'upgrade_info'),
        upgrade_info_json=pulumi.get(__ret__, 'upgrade_info_json'),
        voters=pulumi.get(__ret__, 'voters'))


@_utilities.lift_output_func(get_raft_autopilot_state)
def get_raft_autopilot_state_output(namespace: Optional[pulumi.Input[Optional[str]]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRaftAutopilotStateResult]:
    """
    Displays the state of the raft cluster under integrated storage as seen by
    autopilot. It shows whether autopilot thinks the cluster is healthy or not, and
    how many nodes could fail before the cluster becomes unhealthy ("Failure
    Tolerance"). For more information, please refer to the
    [Vault documentation](https://developer.hashicorp.com/vault/api-docs/system/storage/raftautopilot#get-cluster-state).

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vault as vault

    main = vault.get_raft_autopilot_state()
    pulumi.export("failure-tolerance", main.failure_tolerance)
    ```


    :param str namespace: The namespace of the target resource.
           The value should not contain leading or trailing forward slashes.
           The `namespace` is always relative to the provider's configured [namespace](https://www.terraform.io/docs/providers/vault#namespace).
           *Available only for Vault Enterprise*.
    """
    ...
