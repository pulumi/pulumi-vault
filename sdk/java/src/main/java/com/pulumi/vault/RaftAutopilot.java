// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.RaftAutopilotArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.RaftAutopilotState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Autopilot enables automated workflows for managing Raft clusters. The
 * current feature set includes 3 main features: Server Stabilization, Dead
 * Server Cleanup and State API. **These three features are introduced in
 * Vault 1.7.**
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * <pre>
 * {@code
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.RaftAutopilot;
 * import com.pulumi.vault.RaftAutopilotArgs;
 * import java.util.List;
 * import java.util.ArrayList;
 * import java.util.Map;
 * import java.io.File;
 * import java.nio.file.Files;
 * import java.nio.file.Paths;
 * 
 * public class App {
 *     public static void main(String[] args) {
 *         Pulumi.run(App::stack);
 *     }
 * 
 *     public static void stack(Context ctx) {
 *         var autopilot = new RaftAutopilot("autopilot", RaftAutopilotArgs.builder()        
 *             .cleanupDeadServers(true)
 *             .deadServerLastContactThreshold("24h0m0s")
 *             .lastContactThreshold("10s")
 *             .maxTrailingLogs(1000)
 *             .minQuorum(3)
 *             .serverStabilizationTime("10s")
 *             .build());
 * 
 *     }
 * }
 * }
 * </pre>
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * Raft Autopilot config can be imported using the ID, e.g.
 * 
 * ```sh
 * $ pulumi import vault:index/raftAutopilot:RaftAutopilot autopilot sys/storage/raft/autopilot/configuration
 * ```
 * 
 */
@ResourceType(type="vault:index/raftAutopilot:RaftAutopilot")
public class RaftAutopilot extends com.pulumi.resources.CustomResource {
    /**
     * Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     * 
     */
    @Export(name="cleanupDeadServers", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> cleanupDeadServers;

    /**
     * @return Specifies whether to remove dead server nodes
     * periodically or when a new server joins. This requires that `min-quorum` is also set.
     * 
     */
    public Output<Optional<Boolean>> cleanupDeadServers() {
        return Codegen.optional(this.cleanupDeadServers);
    }
    /**
     * Limit the amount of time a
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanup_dead_servers` is set.
     * 
     */
    @Export(name="deadServerLastContactThreshold", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> deadServerLastContactThreshold;

    /**
     * @return Limit the amount of time a
     * server can go without leader contact before being considered failed. This only takes
     * effect when `cleanup_dead_servers` is set.
     * 
     */
    public Output<Optional<String>> deadServerLastContactThreshold() {
        return Codegen.optional(this.deadServerLastContactThreshold);
    }
    /**
     * Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     * 
     */
    @Export(name="disableUpgradeMigration", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disableUpgradeMigration;

    /**
     * @return Disables automatically upgrading Vault using autopilot. (Enterprise-only)
     * 
     */
    public Output<Optional<Boolean>> disableUpgradeMigration() {
        return Codegen.optional(this.disableUpgradeMigration);
    }
    /**
     * Limit the amount of time a server can go
     * without leader contact before being considered unhealthy.
     * 
     */
    @Export(name="lastContactThreshold", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> lastContactThreshold;

    /**
     * @return Limit the amount of time a server can go
     * without leader contact before being considered unhealthy.
     * 
     */
    public Output<Optional<String>> lastContactThreshold() {
        return Codegen.optional(this.lastContactThreshold);
    }
    /**
     * Maximum number of log entries in the Raft log
     * that a server can be behind its leader before being considered unhealthy.
     * 
     */
    @Export(name="maxTrailingLogs", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> maxTrailingLogs;

    /**
     * @return Maximum number of log entries in the Raft log
     * that a server can be behind its leader before being considered unhealthy.
     * 
     */
    public Output<Optional<Integer>> maxTrailingLogs() {
        return Codegen.optional(this.maxTrailingLogs);
    }
    /**
     * Minimum number of servers allowed in a cluster before
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     * 
     */
    @Export(name="minQuorum", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> minQuorum;

    /**
     * @return Minimum number of servers allowed in a cluster before
     * autopilot can prune dead servers. This should at least be 3. Applicable only for
     * voting nodes.
     * 
     */
    public Output<Optional<Integer>> minQuorum() {
        return Codegen.optional(this.minQuorum);
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Minimum amount of time a server must be
     * stable in the &#39;healthy&#39; state before being added to the cluster.
     * 
     */
    @Export(name="serverStabilizationTime", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> serverStabilizationTime;

    /**
     * @return Minimum amount of time a server must be
     * stable in the &#39;healthy&#39; state before being added to the cluster.
     * 
     */
    public Output<Optional<String>> serverStabilizationTime() {
        return Codegen.optional(this.serverStabilizationTime);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public RaftAutopilot(String name) {
        this(name, RaftAutopilotArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public RaftAutopilot(String name, @Nullable RaftAutopilotArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public RaftAutopilot(String name, @Nullable RaftAutopilotArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/raftAutopilot:RaftAutopilot", name, args == null ? RaftAutopilotArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private RaftAutopilot(String name, Output<String> id, @Nullable RaftAutopilotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/raftAutopilot:RaftAutopilot", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static RaftAutopilot get(String name, Output<String> id, @Nullable RaftAutopilotState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new RaftAutopilot(name, id, state, options);
    }
}
