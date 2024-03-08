// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.pkiSecret.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendRolePolicyIdentifier {
    /**
     * @return The URL of the CPS for the policy identifier
     * 
     * Example usage:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.Mount;
     * import com.pulumi.vault.MountArgs;
     * import com.pulumi.vault.pkiSecret.SecretBackendRole;
     * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
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
     *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
     *             .path(&#34;pki&#34;)
     *             .type(&#34;pki&#34;)
     *             .defaultLeaseTtlSeconds(3600)
     *             .maxLeaseTtlSeconds(86400)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(pki.path())
     *             .ttl(3600)
     *             .allowIpSans(true)
     *             .keyType(&#34;rsa&#34;)
     *             .keyBits(4096)
     *             .allowedDomains(            
     *                 &#34;example.com&#34;,
     *                 &#34;my.domain&#34;)
     *             .allowSubdomains(true)
     *             .policyIdentifiers(            
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
     *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
     *                 ),
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
     *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
     *                 ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    private @Nullable String cps;
    /**
     * @return A notice for the policy identifier
     * 
     */
    private @Nullable String notice;
    /**
     * @return The OID for the policy identifier
     * 
     */
    private String oid;

    private SecretBackendRolePolicyIdentifier() {}
    /**
     * @return The URL of the CPS for the policy identifier
     * 
     * Example usage:
     * &lt;!--Start PulumiCodeChooser --&gt;
     * ```java
     * package generated_program;
     * 
     * import com.pulumi.Context;
     * import com.pulumi.Pulumi;
     * import com.pulumi.core.Output;
     * import com.pulumi.vault.Mount;
     * import com.pulumi.vault.MountArgs;
     * import com.pulumi.vault.pkiSecret.SecretBackendRole;
     * import com.pulumi.vault.pkiSecret.SecretBackendRoleArgs;
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
     *         var pki = new Mount(&#34;pki&#34;, MountArgs.builder()        
     *             .path(&#34;pki&#34;)
     *             .type(&#34;pki&#34;)
     *             .defaultLeaseTtlSeconds(3600)
     *             .maxLeaseTtlSeconds(86400)
     *             .build());
     * 
     *         var role = new SecretBackendRole(&#34;role&#34;, SecretBackendRoleArgs.builder()        
     *             .backend(pki.path())
     *             .ttl(3600)
     *             .allowIpSans(true)
     *             .keyType(&#34;rsa&#34;)
     *             .keyBits(4096)
     *             .allowedDomains(            
     *                 &#34;example.com&#34;,
     *                 &#34;my.domain&#34;)
     *             .allowSubdomains(true)
     *             .policyIdentifiers(            
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.7.8&#34;),
     *                     Map.entry(&#34;notice&#34;, &#34;I am a user Notice&#34;)
     *                 ),
     *                 Map.ofEntries(
     *                     Map.entry(&#34;oid&#34;, &#34;1.3.6.1.4.1.44947.1.2.4&#34;),
     *                     Map.entry(&#34;cps&#34;, &#34;https://example.com&#34;)
     *                 ))
     *             .build());
     * 
     *     }
     * }
     * ```
     * &lt;!--End PulumiCodeChooser --&gt;
     * 
     */
    public Optional<String> cps() {
        return Optional.ofNullable(this.cps);
    }
    /**
     * @return A notice for the policy identifier
     * 
     */
    public Optional<String> notice() {
        return Optional.ofNullable(this.notice);
    }
    /**
     * @return The OID for the policy identifier
     * 
     */
    public String oid() {
        return this.oid;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRolePolicyIdentifier defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String cps;
        private @Nullable String notice;
        private String oid;
        public Builder() {}
        public Builder(SecretBackendRolePolicyIdentifier defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.cps = defaults.cps;
    	      this.notice = defaults.notice;
    	      this.oid = defaults.oid;
        }

        @CustomType.Setter
        public Builder cps(@Nullable String cps) {

            this.cps = cps;
            return this;
        }
        @CustomType.Setter
        public Builder notice(@Nullable String notice) {

            this.notice = notice;
            return this;
        }
        @CustomType.Setter
        public Builder oid(String oid) {
            if (oid == null) {
              throw new MissingRequiredPropertyException("SecretBackendRolePolicyIdentifier", "oid");
            }
            this.oid = oid;
            return this;
        }
        public SecretBackendRolePolicyIdentifier build() {
            final var _resultValue = new SecretBackendRolePolicyIdentifier();
            _resultValue.cps = cps;
            _resultValue.notice = notice;
            _resultValue.oid = oid;
            return _resultValue;
        }
    }
}
