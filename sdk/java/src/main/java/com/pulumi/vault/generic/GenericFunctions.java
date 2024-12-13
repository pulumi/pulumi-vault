// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.generic;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.generic.inputs.GetSecretArgs;
import com.pulumi.vault.generic.inputs.GetSecretPlainArgs;
import com.pulumi.vault.generic.outputs.GetSecretResult;
import java.util.concurrent.CompletableFuture;

public final class GenericFunctions {
    /**
     * ## Example Usage
     * 
     * ### Generic secret
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.generic.GenericFunctions;
     * import com.pulumi.vault.generic.inputs.GetSecretArgs;
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
     *         final var rundeckAuth = GenericFunctions.getSecret(GetSecretArgs.builder()
     *             .path("secret/rundeck_auth")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ### KV
     * 
     * For this example, consider `example` as a path for a KV engine.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Required Vault Capabilities
     * 
     * Use of this resource requires the `read` capability on the given path.
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args) {
        return getSecret(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * 
     * ### Generic secret
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.generic.GenericFunctions;
     * import com.pulumi.vault.generic.inputs.GetSecretArgs;
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
     *         final var rundeckAuth = GenericFunctions.getSecret(GetSecretArgs.builder()
     *             .path("secret/rundeck_auth")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ### KV
     * 
     * For this example, consider `example` as a path for a KV engine.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Required Vault Capabilities
     * 
     * Use of this resource requires the `read` capability on the given path.
     * 
     */
    public static CompletableFuture<GetSecretResult> getSecretPlain(GetSecretPlainArgs args) {
        return getSecretPlain(args, InvokeOptions.Empty);
    }
    /**
     * ## Example Usage
     * 
     * ### Generic secret
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.generic.GenericFunctions;
     * import com.pulumi.vault.generic.inputs.GetSecretArgs;
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
     *         final var rundeckAuth = GenericFunctions.getSecret(GetSecretArgs.builder()
     *             .path("secret/rundeck_auth")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ### KV
     * 
     * For this example, consider `example` as a path for a KV engine.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Required Vault Capabilities
     * 
     * Use of this resource requires the `read` capability on the given path.
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:generic/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * 
     * ### Generic secret
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.generic.GenericFunctions;
     * import com.pulumi.vault.generic.inputs.GetSecretArgs;
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
     *         final var rundeckAuth = GenericFunctions.getSecret(GetSecretArgs.builder()
     *             .path("secret/rundeck_auth")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ### KV
     * 
     * For this example, consider `example` as a path for a KV engine.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Required Vault Capabilities
     * 
     * Use of this resource requires the `read` capability on the given path.
     * 
     */
    public static Output<GetSecretResult> getSecret(GetSecretArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("vault:generic/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
    /**
     * ## Example Usage
     * 
     * ### Generic secret
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * <pre>
     * {@code
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.generic.GenericFunctions;
     * import com.pulumi.vault.generic.inputs.GetSecretArgs;
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
     *         final var rundeckAuth = GenericFunctions.getSecret(GetSecretArgs.builder()
     *             .path("secret/rundeck_auth")
     *             .build());
     * 
     *     }
     * }
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ### KV
     * 
     * For this example, consider `example` as a path for a KV engine.
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     * ## Required Vault Capabilities
     * 
     * Use of this resource requires the `read` capability on the given path.
     * 
     */
    public static CompletableFuture<GetSecretResult> getSecretPlain(GetSecretPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:generic/getSecret:getSecret", TypeShape.of(GetSecretResult.class), args, Utilities.withVersion(options));
    }
}
