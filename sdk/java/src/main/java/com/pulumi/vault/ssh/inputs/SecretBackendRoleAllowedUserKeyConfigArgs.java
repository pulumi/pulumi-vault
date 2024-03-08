// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ssh.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;


public final class SecretBackendRoleAllowedUserKeyConfigArgs extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRoleAllowedUserKeyConfigArgs Empty = new SecretBackendRoleAllowedUserKeyConfigArgs();

    /**
     * A list of allowed key lengths as integers.
     * For key types that do not support setting the length a value of `[0]` should be used.
     * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
     * must be set to a single element list.
     * 
     * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    @Import(name="lengths", required=true)
    private Output<List<Integer>> lengths;

    /**
     * @return A list of allowed key lengths as integers.
     * For key types that do not support setting the length a value of `[0]` should be used.
     * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
     * must be set to a single element list.
     * 
     * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
     * 
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
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
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Output<List<Integer>> lengths() {
        return this.lengths;
    }

    /**
     * The SSH public key type.\
     * *Supported key types are:*
     * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
     * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
     * 
     */
    @Import(name="type", required=true)
    private Output<String> type;

    /**
     * @return The SSH public key type.\
     * *Supported key types are:*
     * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
     * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
     * 
     */
    public Output<String> type() {
        return this.type;
    }

    private SecretBackendRoleAllowedUserKeyConfigArgs() {}

    private SecretBackendRoleAllowedUserKeyConfigArgs(SecretBackendRoleAllowedUserKeyConfigArgs $) {
        this.lengths = $.lengths;
        this.type = $.type;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRoleAllowedUserKeyConfigArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRoleAllowedUserKeyConfigArgs $;

        public Builder() {
            $ = new SecretBackendRoleAllowedUserKeyConfigArgs();
        }

        public Builder(SecretBackendRoleAllowedUserKeyConfigArgs defaults) {
            $ = new SecretBackendRoleAllowedUserKeyConfigArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param lengths A list of allowed key lengths as integers.
         * For key types that do not support setting the length a value of `[0]` should be used.
         * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
         * must be set to a single element list.
         * 
         * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
         * 
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
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
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder lengths(Output<List<Integer>> lengths) {
            $.lengths = lengths;
            return this;
        }

        /**
         * @param lengths A list of allowed key lengths as integers.
         * For key types that do not support setting the length a value of `[0]` should be used.
         * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
         * must be set to a single element list.
         * 
         * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
         * 
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
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
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder lengths(List<Integer> lengths) {
            return lengths(Output.of(lengths));
        }

        /**
         * @param lengths A list of allowed key lengths as integers.
         * For key types that do not support setting the length a value of `[0]` should be used.
         * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
         * must be set to a single element list.
         * 
         * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
         * 
         * &lt;!--Start PulumiCodeChooser --&gt;
         * ```java
         * package generated_program;
         * 
         * import com.pulumi.Context;
         * import com.pulumi.Pulumi;
         * import com.pulumi.core.Output;
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
         *     }
         * }
         * ```
         * &lt;!--End PulumiCodeChooser --&gt;
         * 
         * @return builder
         * 
         */
        public Builder lengths(Integer... lengths) {
            return lengths(List.of(lengths));
        }

        /**
         * @param type The SSH public key type.\
         * *Supported key types are:*
         * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
         * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
         * 
         * @return builder
         * 
         */
        public Builder type(Output<String> type) {
            $.type = type;
            return this;
        }

        /**
         * @param type The SSH public key type.\
         * *Supported key types are:*
         * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
         * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
         * 
         * @return builder
         * 
         */
        public Builder type(String type) {
            return type(Output.of(type));
        }

        public SecretBackendRoleAllowedUserKeyConfigArgs build() {
            if ($.lengths == null) {
                throw new MissingRequiredPropertyException("SecretBackendRoleAllowedUserKeyConfigArgs", "lengths");
            }
            if ($.type == null) {
                throw new MissingRequiredPropertyException("SecretBackendRoleAllowedUserKeyConfigArgs", "type");
            }
            return $;
        }
    }

}
