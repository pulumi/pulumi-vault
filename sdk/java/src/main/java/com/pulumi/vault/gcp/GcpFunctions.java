// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
import com.pulumi.vault.gcp.inputs.GetAuthBackendRolePlainArgs;
import com.pulumi.vault.gcp.outputs.GetAuthBackendRoleResult;
import java.util.concurrent.CompletableFuture;

public final class GcpFunctions {
    /**
     * Reads a GCP auth role from a Vault server.
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
     * import com.pulumi.vault.gcp.GcpFunctions;
     * import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
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
     *         final var role = GcpFunctions.getAuthBackendRole(GetAuthBackendRoleArgs.builder()
     *             .backend("my-gcp-backend")
     *             .roleName("my-role")
     *             .build());
     * 
     *         ctx.export("role-id", role.roleId());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAuthBackendRoleResult> getAuthBackendRole(GetAuthBackendRoleArgs args) {
        return getAuthBackendRole(args, InvokeOptions.Empty);
    }
    /**
     * Reads a GCP auth role from a Vault server.
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
     * import com.pulumi.vault.gcp.GcpFunctions;
     * import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
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
     *         final var role = GcpFunctions.getAuthBackendRole(GetAuthBackendRoleArgs.builder()
     *             .backend("my-gcp-backend")
     *             .roleName("my-role")
     *             .build());
     * 
     *         ctx.export("role-id", role.roleId());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAuthBackendRoleResult> getAuthBackendRolePlain(GetAuthBackendRolePlainArgs args) {
        return getAuthBackendRolePlain(args, InvokeOptions.Empty);
    }
    /**
     * Reads a GCP auth role from a Vault server.
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
     * import com.pulumi.vault.gcp.GcpFunctions;
     * import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
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
     *         final var role = GcpFunctions.getAuthBackendRole(GetAuthBackendRoleArgs.builder()
     *             .backend("my-gcp-backend")
     *             .roleName("my-role")
     *             .build());
     * 
     *         ctx.export("role-id", role.roleId());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAuthBackendRoleResult> getAuthBackendRole(GetAuthBackendRoleArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:gcp/getAuthBackendRole:getAuthBackendRole", TypeShape.of(GetAuthBackendRoleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads a GCP auth role from a Vault server.
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
     * import com.pulumi.vault.gcp.GcpFunctions;
     * import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
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
     *         final var role = GcpFunctions.getAuthBackendRole(GetAuthBackendRoleArgs.builder()
     *             .backend("my-gcp-backend")
     *             .roleName("my-role")
     *             .build());
     * 
     *         ctx.export("role-id", role.roleId());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetAuthBackendRoleResult> getAuthBackendRole(GetAuthBackendRoleArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("vault:gcp/getAuthBackendRole:getAuthBackendRole", TypeShape.of(GetAuthBackendRoleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads a GCP auth role from a Vault server.
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
     * import com.pulumi.vault.gcp.GcpFunctions;
     * import com.pulumi.vault.gcp.inputs.GetAuthBackendRoleArgs;
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
     *         final var role = GcpFunctions.getAuthBackendRole(GetAuthBackendRoleArgs.builder()
     *             .backend("my-gcp-backend")
     *             .roleName("my-role")
     *             .build());
     * 
     *         ctx.export("role-id", role.roleId());
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetAuthBackendRoleResult> getAuthBackendRolePlain(GetAuthBackendRolePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:gcp/getAuthBackendRole:getAuthBackendRole", TypeShape.of(GetAuthBackendRoleResult.class), args, Utilities.withVersion(options));
    }
}
