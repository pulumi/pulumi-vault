// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.azure;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.azure.inputs.GetAccessCredentialsArgs;
import com.pulumi.vault.azure.inputs.GetAccessCredentialsPlainArgs;
import com.pulumi.vault.azure.outputs.GetAccessCredentialsResult;
import java.util.concurrent.CompletableFuture;

public final class AzureFunctions {
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
     * import com.pulumi.vault.azure.AzureFunctions;
     * import com.pulumi.vault.azure.inputs.GetAccessCredentialsArgs;
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
     *         final var creds = AzureFunctions.getAccessCredentials(GetAccessCredentialsArgs.builder()
     *             .role(&#34;my-role&#34;)
     *             .validateCreds(true)
     *             .numSequentialSuccesses(8)
     *             .numSecondsBetweenTests(1)
     *             .maxCredValidationSeconds(300)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Caveats
     * 
     * The `validate_creds` option requires read-access to the `backend` config endpoint.
     * If the effective Vault role does not have the required permissions then valid values
     * are required to be set for: `subscription_id`, `tenant_id`, `environment`.
     * 
     */
    public static Output<GetAccessCredentialsResult> getAccessCredentials(GetAccessCredentialsArgs args) {
        return getAccessCredentials(args, InvokeOptions.Empty);
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
     * import com.pulumi.vault.azure.AzureFunctions;
     * import com.pulumi.vault.azure.inputs.GetAccessCredentialsArgs;
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
     *         final var creds = AzureFunctions.getAccessCredentials(GetAccessCredentialsArgs.builder()
     *             .role(&#34;my-role&#34;)
     *             .validateCreds(true)
     *             .numSequentialSuccesses(8)
     *             .numSecondsBetweenTests(1)
     *             .maxCredValidationSeconds(300)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Caveats
     * 
     * The `validate_creds` option requires read-access to the `backend` config endpoint.
     * If the effective Vault role does not have the required permissions then valid values
     * are required to be set for: `subscription_id`, `tenant_id`, `environment`.
     * 
     */
    public static CompletableFuture<GetAccessCredentialsResult> getAccessCredentialsPlain(GetAccessCredentialsPlainArgs args) {
        return getAccessCredentialsPlain(args, InvokeOptions.Empty);
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
     * import com.pulumi.vault.azure.AzureFunctions;
     * import com.pulumi.vault.azure.inputs.GetAccessCredentialsArgs;
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
     *         final var creds = AzureFunctions.getAccessCredentials(GetAccessCredentialsArgs.builder()
     *             .role(&#34;my-role&#34;)
     *             .validateCreds(true)
     *             .numSequentialSuccesses(8)
     *             .numSecondsBetweenTests(1)
     *             .maxCredValidationSeconds(300)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Caveats
     * 
     * The `validate_creds` option requires read-access to the `backend` config endpoint.
     * If the effective Vault role does not have the required permissions then valid values
     * are required to be set for: `subscription_id`, `tenant_id`, `environment`.
     * 
     */
    public static Output<GetAccessCredentialsResult> getAccessCredentials(GetAccessCredentialsArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:azure/getAccessCredentials:getAccessCredentials", TypeShape.of(GetAccessCredentialsResult.class), args, Utilities.withVersion(options));
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
     * import com.pulumi.vault.azure.AzureFunctions;
     * import com.pulumi.vault.azure.inputs.GetAccessCredentialsArgs;
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
     *         final var creds = AzureFunctions.getAccessCredentials(GetAccessCredentialsArgs.builder()
     *             .role(&#34;my-role&#34;)
     *             .validateCreds(true)
     *             .numSequentialSuccesses(8)
     *             .numSecondsBetweenTests(1)
     *             .maxCredValidationSeconds(300)
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Caveats
     * 
     * The `validate_creds` option requires read-access to the `backend` config endpoint.
     * If the effective Vault role does not have the required permissions then valid values
     * are required to be set for: `subscription_id`, `tenant_id`, `environment`.
     * 
     */
    public static CompletableFuture<GetAccessCredentialsResult> getAccessCredentialsPlain(GetAccessCredentialsPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:azure/getAccessCredentials:getAccessCredentials", TypeShape.of(GetAccessCredentialsResult.class), args, Utilities.withVersion(options));
    }
}
