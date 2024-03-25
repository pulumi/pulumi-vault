// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class RaftAutopilotState extends com.pulumi.resources.ResourceArgs {

    public static final RaftAutopilotState Empty = new RaftAutopilotState();

    /**
     * Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     * 
     */
    @Import(name="cleanupDeadServers")
    private @Nullable Output<Boolean> cleanupDeadServers;

    /**
     * @return Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     * 
     */
    public Optional<Output<Boolean>> cleanupDeadServers() {
        return Optional.ofNullable(this.cleanupDeadServers);
    }

    /**
     * Limit the amount of time a
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanup_dead_servers` is set.
     * 
     */
    @Import(name="deadServerLastContactThreshold")
    private @Nullable Output<String> deadServerLastContactThreshold;

    /**
     * @return Limit the amount of time a
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanup_dead_servers` is set.
     * 
     */
    public Optional<Output<String>> deadServerLastContactThreshold() {
        return Optional.ofNullable(this.deadServerLastContactThreshold);
    }

    /**
     * Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     * 
     */
    @Import(name="disableUpgradeMigration")
    private @Nullable Output<Boolean> disableUpgradeMigration;

    /**
     * @return Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     * 
     */
    public Optional<Output<Boolean>> disableUpgradeMigration() {
        return Optional.ofNullable(this.disableUpgradeMigration);
    }

    /**
     * Limit the amount of time a server can go
     * without leader contact before being considered unhealthy.
     * 
     */
    @Import(name="lastContactThreshold")
    private @Nullable Output<String> lastContactThreshold;

    /**
     * @return Limit the amount of time a server can go
     * without leader contact before being considered unhealthy.
     * 
     */
    public Optional<Output<String>> lastContactThreshold() {
        return Optional.ofNullable(this.lastContactThreshold);
    }

    /**
     * Maximum number of log entries in the Raft log
     * that a server can be behind its leader before being considered unhealthy.
     * 
     */
    @Import(name="maxTrailingLogs")
    private @Nullable Output<Integer> maxTrailingLogs;

    /**
     * @return Maximum number of log entries in the Raft log
     * that a server can be behind its leader before being considered unhealthy.
     * 
     */
    public Optional<Output<Integer>> maxTrailingLogs() {
        return Optional.ofNullable(this.maxTrailingLogs);
    }

    /**
     * Minimum number of servers allowed in a cluster before
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     * 
     */
    @Import(name="minQuorum")
    private @Nullable Output<Integer> minQuorum;

    /**
     * @return Minimum number of servers allowed in a cluster before
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     * 
     */
    public Optional<Output<Integer>> minQuorum() {
        return Optional.ofNullable(this.minQuorum);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Minimum amount of time a server must be
     * stable in the &#39;healthy&#39; state before being added to the cluster.
     * 
     */
    @Import(name="serverStabilizationTime")
    private @Nullable Output<String> serverStabilizationTime;

    /**
     * @return Minimum amount of time a server must be
     * stable in the &#39;healthy&#39; state before being added to the cluster.
     * 
     */
    public Optional<Output<String>> serverStabilizationTime() {
        return Optional.ofNullable(this.serverStabilizationTime);
    }

    private RaftAutopilotState() {}

    private RaftAutopilotState(RaftAutopilotState $) {
        this.cleanupDeadServers = $.cleanupDeadServers;
        this.deadServerLastContactThreshold = $.deadServerLastContactThreshold;
        this.disableUpgradeMigration = $.disableUpgradeMigration;
        this.lastContactThreshold = $.lastContactThreshold;
        this.maxTrailingLogs = $.maxTrailingLogs;
        this.minQuorum = $.minQuorum;
        this.namespace = $.namespace;
        this.serverStabilizationTime = $.serverStabilizationTime;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(RaftAutopilotState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private RaftAutopilotState $;

        public Builder() {
            $ = new RaftAutopilotState();
        }

        public Builder(RaftAutopilotState defaults) {
            $ = new RaftAutopilotState(Objects.requireNonNull(defaults));
        }

        /**
         * @param cleanupDeadServers Specifies whether to remove dead server nodes
         * periodically or when a new server joins. This requires that `min-quorum` is also set.
         * 
         * @return builder
         * 
         */
        public Builder cleanupDeadServers(@Nullable Output<Boolean> cleanupDeadServers) {
            $.cleanupDeadServers = cleanupDeadServers;
            return this;
        }

        /**
         * @param cleanupDeadServers Specifies whether to remove dead server nodes
         * periodically or when a new server joins. This requires that `min-quorum` is also set.
         * 
         * @return builder
         * 
         */
        public Builder cleanupDeadServers(Boolean cleanupDeadServers) {
            return cleanupDeadServers(Output.of(cleanupDeadServers));
        }

        /**
         * @param deadServerLastContactThreshold Limit the amount of time a
         * server can go without leader contact before being considered failed. This only takes
         * effect when `cleanup_dead_servers` is set.
         * 
         * @return builder
         * 
         */
        public Builder deadServerLastContactThreshold(@Nullable Output<String> deadServerLastContactThreshold) {
            $.deadServerLastContactThreshold = deadServerLastContactThreshold;
            return this;
        }

        /**
         * @param deadServerLastContactThreshold Limit the amount of time a
         * server can go without leader contact before being considered failed. This only takes
         * effect when `cleanup_dead_servers` is set.
         * 
         * @return builder
         * 
         */
        public Builder deadServerLastContactThreshold(String deadServerLastContactThreshold) {
            return deadServerLastContactThreshold(Output.of(deadServerLastContactThreshold));
        }

        /**
         * @param disableUpgradeMigration Disables automatically upgrading Vault using autopilot. (Enterprise-only)
         * 
         * @return builder
         * 
         */
        public Builder disableUpgradeMigration(@Nullable Output<Boolean> disableUpgradeMigration) {
            $.disableUpgradeMigration = disableUpgradeMigration;
            return this;
        }

        /**
         * @param disableUpgradeMigration Disables automatically upgrading Vault using autopilot. (Enterprise-only)
         * 
         * @return builder
         * 
         */
        public Builder disableUpgradeMigration(Boolean disableUpgradeMigration) {
            return disableUpgradeMigration(Output.of(disableUpgradeMigration));
        }

        /**
         * @param lastContactThreshold Limit the amount of time a server can go
         * without leader contact before being considered unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder lastContactThreshold(@Nullable Output<String> lastContactThreshold) {
            $.lastContactThreshold = lastContactThreshold;
            return this;
        }

        /**
         * @param lastContactThreshold Limit the amount of time a server can go
         * without leader contact before being considered unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder lastContactThreshold(String lastContactThreshold) {
            return lastContactThreshold(Output.of(lastContactThreshold));
        }

        /**
         * @param maxTrailingLogs Maximum number of log entries in the Raft log
         * that a server can be behind its leader before being considered unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder maxTrailingLogs(@Nullable Output<Integer> maxTrailingLogs) {
            $.maxTrailingLogs = maxTrailingLogs;
            return this;
        }

        /**
         * @param maxTrailingLogs Maximum number of log entries in the Raft log
         * that a server can be behind its leader before being considered unhealthy.
         * 
         * @return builder
         * 
         */
        public Builder maxTrailingLogs(Integer maxTrailingLogs) {
            return maxTrailingLogs(Output.of(maxTrailingLogs));
        }

        /**
         * @param minQuorum Minimum number of servers allowed in a cluster before
         * autopilot can prune dead servers. This should at least be 3. Applicable only for
         * voting nodes.
         * 
         * @return builder
         * 
         */
        public Builder minQuorum(@Nullable Output<Integer> minQuorum) {
            $.minQuorum = minQuorum;
            return this;
        }

        /**
         * @param minQuorum Minimum number of servers allowed in a cluster before
         * autopilot can prune dead servers. This should at least be 3. Applicable only for
         * voting nodes.
         * 
         * @return builder
         * 
         */
        public Builder minQuorum(Integer minQuorum) {
            return minQuorum(Output.of(minQuorum));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param serverStabilizationTime Minimum amount of time a server must be
         * stable in the &#39;healthy&#39; state before being added to the cluster.
         * 
         * @return builder
         * 
         */
        public Builder serverStabilizationTime(@Nullable Output<String> serverStabilizationTime) {
            $.serverStabilizationTime = serverStabilizationTime;
            return this;
        }

        /**
         * @param serverStabilizationTime Minimum amount of time a server must be
         * stable in the &#39;healthy&#39; state before being added to the cluster.
         * 
         * @return builder
         * 
         */
        public Builder serverStabilizationTime(String serverStabilizationTime) {
            return serverStabilizationTime(Output.of(serverStabilizationTime));
        }

        public RaftAutopilotState build() {
            return $;
        }
    }

}
