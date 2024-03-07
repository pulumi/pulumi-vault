// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.kubernetes;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.kubernetes.inputs.GetAuthBackendConfigArgs;
import com.pulumi.vault.kubernetes.inputs.GetAuthBackendConfigPlainArgs;
import com.pulumi.vault.kubernetes.inputs.GetAuthBackendRoleArgs;
import com.pulumi.vault.kubernetes.inputs.GetAuthBackendRolePlainArgs;
import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenArgs;
import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenPlainArgs;
import com.pulumi.vault.kubernetes.outputs.GetAuthBackendConfigResult;
import com.pulumi.vault.kubernetes.outputs.GetAuthBackendRoleResult;
import com.pulumi.vault.kubernetes.outputs.GetServiceAccountTokenResult;
import java.util.concurrent.CompletableFuture;

public final class KubernetesFunctions {
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static Output<GetAuthBackendConfigResult> getAuthBackendConfig() {
        return getAuthBackendConfig(GetAuthBackendConfigArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static CompletableFuture<GetAuthBackendConfigResult> getAuthBackendConfigPlain() {
        return getAuthBackendConfigPlain(GetAuthBackendConfigPlainArgs.Empty, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static Output<GetAuthBackendConfigResult> getAuthBackendConfig(GetAuthBackendConfigArgs args) {
        return getAuthBackendConfig(args, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static CompletableFuture<GetAuthBackendConfigResult> getAuthBackendConfigPlain(GetAuthBackendConfigPlainArgs args) {
        return getAuthBackendConfigPlain(args, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static Output<GetAuthBackendConfigResult> getAuthBackendConfig(GetAuthBackendConfigArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", TypeShape.of(GetAuthBackendConfigResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-config) for more
     * information.
     * 
     */
    public static CompletableFuture<GetAuthBackendConfigResult> getAuthBackendConfigPlain(GetAuthBackendConfigPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:kubernetes/getAuthBackendConfig:getAuthBackendConfig", TypeShape.of(GetAuthBackendConfigResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
     * information.
     * 
     */
    public static Output<GetAuthBackendRoleResult> getAuthBackendRole(GetAuthBackendRoleArgs args) {
        return getAuthBackendRole(args, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
     * information.
     * 
     */
    public static CompletableFuture<GetAuthBackendRoleResult> getAuthBackendRolePlain(GetAuthBackendRolePlainArgs args) {
        return getAuthBackendRolePlain(args, InvokeOptions.Empty);
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
     * information.
     * 
     */
    public static Output<GetAuthBackendRoleResult> getAuthBackendRole(GetAuthBackendRoleArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", TypeShape.of(GetAuthBackendRoleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * Reads the Role of an Kubernetes from a Vault server. See the [Vault
     * documentation](https://www.vaultproject.io/api-docs/auth/kubernetes#read-role) for more
     * information.
     * 
     */
    public static CompletableFuture<GetAuthBackendRoleResult> getAuthBackendRolePlain(GetAuthBackendRolePlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:kubernetes/getAuthBackendRole:getAuthBackendRole", TypeShape.of(GetAuthBackendRoleResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.kubernetes.SecretBackend;
     * import com.pulumi.vault.kubernetes.SecretBackendArgs;
     * import com.pulumi.vault.kubernetes.SecretBackendRole;
     * import com.pulumi.vault.kubernetes.SecretBackendRoleArgs;
     * import com.pulumi.vault.kubernetes.KubernetesFunctions;
     * import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenArgs;
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
     *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
     *             .path(&#34;kubernetes&#34;)
     *             .description(&#34;kubernetes secrets engine description&#34;)
     *             .kubernetesHost(&#34;https://127.0.0.1:61233&#34;)
     *             .kubernetesCaCert(Files.readString(Paths.get(&#34;/path/to/cert&#34;)))
     *             .serviceAccountJwt(Files.readString(Paths.get(&#34;/path/to/token&#34;)))
     *             .disableLocalCaJwt(false)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(config.path())
     *             .allowedKubernetesNamespaces(&#34;*&#34;)
     *             .tokenMaxTtl(43200)
     *             .tokenDefaultTtl(21600)
     *             .serviceAccountName(&#34;test-service-account-with-generated-token&#34;)
     *             .extraLabels(Map.ofEntries(
     *                 Map.entry(&#34;id&#34;, &#34;abc123&#34;),
     *                 Map.entry(&#34;name&#34;, &#34;some_name&#34;)
     *             ))
     *             .extraAnnotations(Map.ofEntries(
     *                 Map.entry(&#34;env&#34;, &#34;development&#34;),
     *                 Map.entry(&#34;location&#34;, &#34;earth&#34;)
     *             ))
     *             .build());
     * 
     *         final var token = KubernetesFunctions.getServiceAccountToken(GetServiceAccountTokenArgs.builder()
     *             .backend(config.path())
     *             .role(role.name())
     *             .kubernetesNamespace(&#34;test&#34;)
     *             .clusterRoleBinding(false)
     *             .ttl(&#34;1h&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceAccountTokenResult> getServiceAccountToken(GetServiceAccountTokenArgs args) {
        return getServiceAccountToken(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.kubernetes.SecretBackend;
     * import com.pulumi.vault.kubernetes.SecretBackendArgs;
     * import com.pulumi.vault.kubernetes.SecretBackendRole;
     * import com.pulumi.vault.kubernetes.SecretBackendRoleArgs;
     * import com.pulumi.vault.kubernetes.KubernetesFunctions;
     * import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenArgs;
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
     *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
     *             .path(&#34;kubernetes&#34;)
     *             .description(&#34;kubernetes secrets engine description&#34;)
     *             .kubernetesHost(&#34;https://127.0.0.1:61233&#34;)
     *             .kubernetesCaCert(Files.readString(Paths.get(&#34;/path/to/cert&#34;)))
     *             .serviceAccountJwt(Files.readString(Paths.get(&#34;/path/to/token&#34;)))
     *             .disableLocalCaJwt(false)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(config.path())
     *             .allowedKubernetesNamespaces(&#34;*&#34;)
     *             .tokenMaxTtl(43200)
     *             .tokenDefaultTtl(21600)
     *             .serviceAccountName(&#34;test-service-account-with-generated-token&#34;)
     *             .extraLabels(Map.ofEntries(
     *                 Map.entry(&#34;id&#34;, &#34;abc123&#34;),
     *                 Map.entry(&#34;name&#34;, &#34;some_name&#34;)
     *             ))
     *             .extraAnnotations(Map.ofEntries(
     *                 Map.entry(&#34;env&#34;, &#34;development&#34;),
     *                 Map.entry(&#34;location&#34;, &#34;earth&#34;)
     *             ))
     *             .build());
     * 
     *         final var token = KubernetesFunctions.getServiceAccountToken(GetServiceAccountTokenArgs.builder()
     *             .backend(config.path())
     *             .role(role.name())
     *             .kubernetesNamespace(&#34;test&#34;)
     *             .clusterRoleBinding(false)
     *             .ttl(&#34;1h&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceAccountTokenResult> getServiceAccountTokenPlain(GetServiceAccountTokenPlainArgs args) {
        return getServiceAccountTokenPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.kubernetes.SecretBackend;
     * import com.pulumi.vault.kubernetes.SecretBackendArgs;
     * import com.pulumi.vault.kubernetes.SecretBackendRole;
     * import com.pulumi.vault.kubernetes.SecretBackendRoleArgs;
     * import com.pulumi.vault.kubernetes.KubernetesFunctions;
     * import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenArgs;
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
     *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
     *             .path(&#34;kubernetes&#34;)
     *             .description(&#34;kubernetes secrets engine description&#34;)
     *             .kubernetesHost(&#34;https://127.0.0.1:61233&#34;)
     *             .kubernetesCaCert(Files.readString(Paths.get(&#34;/path/to/cert&#34;)))
     *             .serviceAccountJwt(Files.readString(Paths.get(&#34;/path/to/token&#34;)))
     *             .disableLocalCaJwt(false)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(config.path())
     *             .allowedKubernetesNamespaces(&#34;*&#34;)
     *             .tokenMaxTtl(43200)
     *             .tokenDefaultTtl(21600)
     *             .serviceAccountName(&#34;test-service-account-with-generated-token&#34;)
     *             .extraLabels(Map.ofEntries(
     *                 Map.entry(&#34;id&#34;, &#34;abc123&#34;),
     *                 Map.entry(&#34;name&#34;, &#34;some_name&#34;)
     *             ))
     *             .extraAnnotations(Map.ofEntries(
     *                 Map.entry(&#34;env&#34;, &#34;development&#34;),
     *                 Map.entry(&#34;location&#34;, &#34;earth&#34;)
     *             ))
     *             .build());
     * 
     *         final var token = KubernetesFunctions.getServiceAccountToken(GetServiceAccountTokenArgs.builder()
     *             .backend(config.path())
     *             .role(role.name())
     *             .kubernetesNamespace(&#34;test&#34;)
     *             .clusterRoleBinding(false)
     *             .ttl(&#34;1h&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetServiceAccountTokenResult> getServiceAccountToken(GetServiceAccountTokenArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", TypeShape.of(GetServiceAccountTokenResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.kubernetes.SecretBackend;
     * import com.pulumi.vault.kubernetes.SecretBackendArgs;
     * import com.pulumi.vault.kubernetes.SecretBackendRole;
     * import com.pulumi.vault.kubernetes.SecretBackendRoleArgs;
     * import com.pulumi.vault.kubernetes.KubernetesFunctions;
     * import com.pulumi.vault.kubernetes.inputs.GetServiceAccountTokenArgs;
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
     *         var config = new SecretBackend(&#34;config&#34;, SecretBackendArgs.builder()        
     *             .path(&#34;kubernetes&#34;)
     *             .description(&#34;kubernetes secrets engine description&#34;)
     *             .kubernetesHost(&#34;https://127.0.0.1:61233&#34;)
     *             .kubernetesCaCert(Files.readString(Paths.get(&#34;/path/to/cert&#34;)))
     *             .serviceAccountJwt(Files.readString(Paths.get(&#34;/path/to/token&#34;)))
     *             .disableLocalCaJwt(false)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(config.path())
     *             .allowedKubernetesNamespaces(&#34;*&#34;)
     *             .tokenMaxTtl(43200)
     *             .tokenDefaultTtl(21600)
     *             .serviceAccountName(&#34;test-service-account-with-generated-token&#34;)
     *             .extraLabels(Map.ofEntries(
     *                 Map.entry(&#34;id&#34;, &#34;abc123&#34;),
     *                 Map.entry(&#34;name&#34;, &#34;some_name&#34;)
     *             ))
     *             .extraAnnotations(Map.ofEntries(
     *                 Map.entry(&#34;env&#34;, &#34;development&#34;),
     *                 Map.entry(&#34;location&#34;, &#34;earth&#34;)
     *             ))
     *             .build());
     * 
     *         final var token = KubernetesFunctions.getServiceAccountToken(GetServiceAccountTokenArgs.builder()
     *             .backend(config.path())
     *             .role(role.name())
     *             .kubernetesNamespace(&#34;test&#34;)
     *             .clusterRoleBinding(false)
     *             .ttl(&#34;1h&#34;)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetServiceAccountTokenResult> getServiceAccountTokenPlain(GetServiceAccountTokenPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:kubernetes/getServiceAccountToken:getServiceAccountToken", TypeShape.of(GetServiceAccountTokenResult.class), args, Utilities.withVersion(options));
    }
}
