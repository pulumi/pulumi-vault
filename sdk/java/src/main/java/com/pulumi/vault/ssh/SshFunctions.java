// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ssh;

import com.pulumi.core.Output;
import com.pulumi.core.TypeShape;
import com.pulumi.deployment.Deployment;
import com.pulumi.deployment.InvokeOptions;
import com.pulumi.deployment.InvokeOutputOptions;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
import com.pulumi.vault.ssh.inputs.GetSecretBackendSignPlainArgs;
import com.pulumi.vault.ssh.outputs.GetSecretBackendSignResult;
import java.util.concurrent.CompletableFuture;

public final class SshFunctions {
    /**
     * This is a data source which can be used to sign an SSH public key
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
     * import com.pulumi.vault.ssh.SshFunctions;
     * import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         final var test = SshFunctions.getSecretBackendSign(GetSecretBackendSignArgs.builder()
     *             .path("ssh")
     *             .publicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user}{@literal @}{@code example.com")
     *             .name("test")
     *             .validPrincipals("my-user")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretBackendSignResult> getSecretBackendSign(GetSecretBackendSignArgs args) {
        return getSecretBackendSign(args, InvokeOptions.Empty);
    }
    /**
     * This is a data source which can be used to sign an SSH public key
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
     * import com.pulumi.vault.ssh.SshFunctions;
     * import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         final var test = SshFunctions.getSecretBackendSign(GetSecretBackendSignArgs.builder()
     *             .path("ssh")
     *             .publicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user}{@literal @}{@code example.com")
     *             .name("test")
     *             .validPrincipals("my-user")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSecretBackendSignResult> getSecretBackendSignPlain(GetSecretBackendSignPlainArgs args) {
        return getSecretBackendSignPlain(args, InvokeOptions.Empty);
    }
    /**
     * This is a data source which can be used to sign an SSH public key
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
     * import com.pulumi.vault.ssh.SshFunctions;
     * import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         final var test = SshFunctions.getSecretBackendSign(GetSecretBackendSignArgs.builder()
     *             .path("ssh")
     *             .publicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user}{@literal @}{@code example.com")
     *             .name("test")
     *             .validPrincipals("my-user")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretBackendSignResult> getSecretBackendSign(GetSecretBackendSignArgs args, InvokeOptions options) {
        return Deployment.getInstance().invoke("vault:ssh/getSecretBackendSign:getSecretBackendSign", TypeShape.of(GetSecretBackendSignResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This is a data source which can be used to sign an SSH public key
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
     * import com.pulumi.vault.ssh.SshFunctions;
     * import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         final var test = SshFunctions.getSecretBackendSign(GetSecretBackendSignArgs.builder()
     *             .path("ssh")
     *             .publicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user}{@literal @}{@code example.com")
     *             .name("test")
     *             .validPrincipals("my-user")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static Output<GetSecretBackendSignResult> getSecretBackendSign(GetSecretBackendSignArgs args, InvokeOutputOptions options) {
        return Deployment.getInstance().invoke("vault:ssh/getSecretBackendSign:getSecretBackendSign", TypeShape.of(GetSecretBackendSignResult.class), args, Utilities.withVersion(options));
    }
    /**
     * This is a data source which can be used to sign an SSH public key
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
     * import com.pulumi.vault.ssh.SshFunctions;
     * import com.pulumi.vault.ssh.inputs.GetSecretBackendSignArgs;
     * import java.util.List;
     * import java.util.ArrayList;
     * import java.util.Map;
     * import java.io.File;
     * import java.nio.file.Files;
     * import java.nio.file.Paths;
     * 
     * public class App }{{@code
     *     public static void main(String[] args) }{{@code
     *         Pulumi.run(App::stack);
     *     }}{@code
     * 
     *     public static void stack(Context ctx) }{{@code
     *         final var test = SshFunctions.getSecretBackendSign(GetSecretBackendSignArgs.builder()
     *             .path("ssh")
     *             .publicKey("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDR6q4PTcuIkpdGEqaCaxnR8/REqlbSiEIKaRZkVSjiTXOaiSfUsy9cY2+7+oO9fLMUrhylImerjzEoagX1IjYvc9IeUBaRnfacN7QwUDfstgp2jknbg7rNX9j9nFxwltV/jYQPcRq8Ud0wn1nb4qixq+diM7+Up+xJOeaKxbpjEUJH5dcvaBB+Aa24tJpjOQxtFyQ6dUxlgJu0tcygZR92kKYCVjZDohlSED3i/Ak2KFwqCKx2IZWq9z1vMEgmRzv++4Qt1OsbpW8itiCyWn6lmV33eDCdjMrr9TEThQNnMinPrHdmVUnPZ/OomP+rLDRE9lQR16uaSvKhg5TWOFIXRPyEhX9arEATrE4KSWeQN2qgHOb6P24YqgEm1ZdHJq25q/nBBAa1x0tFMiWqZwOsGeJ9nTeOeyiqFKH5YRBo6DIy3ag3taFsfQSve6oqjnrudUd1hJ8/bNSz8amECfP0ULvAEAgpiurj3eCPc3OcXl4tAld9F6KwabEJV5eelcs= user}{@literal @}{@code example.com")
     *             .name("test")
     *             .validPrincipals("my-user")
     *             .build());
     * 
     *     }}{@code
     * }}{@code
     * }
     * </pre>
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public static CompletableFuture<GetSecretBackendSignResult> getSecretBackendSignPlain(GetSecretBackendSignPlainArgs args, InvokeOptions options) {
        return Deployment.getInstance().invokeAsync("vault:ssh/getSecretBackendSign:getSecretBackendSign", TypeShape.of(GetSecretBackendSignResult.class), args, Utilities.withVersion(options));
    }
}
