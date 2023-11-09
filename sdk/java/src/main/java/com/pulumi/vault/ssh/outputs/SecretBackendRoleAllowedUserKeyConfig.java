// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ssh.outputs;

import com.pulumi.core.annotations.CustomType;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Objects;

@CustomType
public final class SecretBackendRoleAllowedUserKeyConfig {
    /**
     * @return A list of allowed key lengths as integers.
     * For key types that do not support setting the length a value of `[0]` should be used.
     * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
     * must be set to a single element list.
     * 
     * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
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
     * 
     */
    private List<Integer> lengths;
    /**
     * @return The SSH public key type.\
     * *Supported key types are:*
     * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
     * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
     * 
     */
    private String type;

    private SecretBackendRoleAllowedUserKeyConfig() {}
    /**
     * @return A list of allowed key lengths as integers.
     * For key types that do not support setting the length a value of `[0]` should be used.
     * Setting multiple lengths is only supported on Vault 1.10+. For prior releases `length`
     * must be set to a single element list.
     * 
     * Example configuration blocks that might be included in the `vault.ssh.SecretBackendRole`
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
     * 
     */
    public List<Integer> lengths() {
        return this.lengths;
    }
    /**
     * @return The SSH public key type.\
     * *Supported key types are:*
     * `rsa`, `ecdsa`, `ec`, `dsa`, `ed25519`, `ssh-rsa`, `ssh-dss`, `ssh-ed25519`,
     * `ecdsa-sha2-nistp256`, `ecdsa-sha2-nistp384`, `ecdsa-sha2-nistp521`
     * 
     */
    public String type() {
        return this.type;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRoleAllowedUserKeyConfig defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private List<Integer> lengths;
        private String type;
        public Builder() {}
        public Builder(SecretBackendRoleAllowedUserKeyConfig defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.lengths = defaults.lengths;
    	      this.type = defaults.type;
        }

        @CustomType.Setter
        public Builder lengths(List<Integer> lengths) {
            this.lengths = Objects.requireNonNull(lengths);
            return this;
        }
        public Builder lengths(Integer... lengths) {
            return lengths(List.of(lengths));
        }
        @CustomType.Setter
        public Builder type(String type) {
            this.type = Objects.requireNonNull(type);
            return this;
        }
        public SecretBackendRoleAllowedUserKeyConfig build() {
            final var o = new SecretBackendRoleAllowedUserKeyConfig();
            o.lengths = lengths;
            o.type = type;
            return o;
        }
    }
}
