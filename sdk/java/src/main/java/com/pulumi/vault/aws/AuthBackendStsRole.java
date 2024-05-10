// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.AuthBackendStsRoleArgs;
import com.pulumi.vault.aws.inputs.AuthBackendStsRoleState;
import java.lang.String;
import java.util.Optional;
import javax.annotation.Nullable;

/**
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
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendStsRole;
 * import com.pulumi.vault.aws.AuthBackendStsRoleArgs;
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
 *         var aws = new AuthBackend("aws", AuthBackendArgs.builder()        
 *             .type("aws")
 *             .build());
 * 
 *         var role = new AuthBackendStsRole("role", AuthBackendStsRoleArgs.builder()        
 *             .backend(aws.path())
 *             .accountId("1234567890")
 *             .stsRole("arn:aws:iam::1234567890:role/my-role")
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
 * AWS auth backend STS roles can be imported using `auth/`, the `backend` path, `/config/sts/`, and the `account_id` e.g.
 * 
 * ```sh
 * $ pulumi import vault:aws/authBackendStsRole:AuthBackendStsRole example auth/aws/config/sts/1234567890
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendStsRole:AuthBackendStsRole")
public class AuthBackendStsRole extends com.pulumi.resources.CustomResource {
    /**
     * The AWS account ID to configure the STS role for.
     * 
     */
    @Export(name="accountId", refs={String.class}, tree="[0]")
    private Output<String> accountId;

    /**
     * @return The AWS account ID to configure the STS role for.
     * 
     */
    public Output<String> accountId() {
        return this.accountId;
    }
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
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
     * The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `account_id`.
     * 
     */
    @Export(name="stsRole", refs={String.class}, tree="[0]")
    private Output<String> stsRole;

    /**
     * @return The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `account_id`.
     * 
     */
    public Output<String> stsRole() {
        return this.stsRole;
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendStsRole(String name) {
        this(name, AuthBackendStsRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendStsRole(String name, AuthBackendStsRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendStsRole(String name, AuthBackendStsRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendStsRole:AuthBackendStsRole", name, args == null ? AuthBackendStsRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendStsRole(String name, Output<String> id, @Nullable AuthBackendStsRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendStsRole:AuthBackendStsRole", name, state, makeResourceOptions(options, id));
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
    public static AuthBackendStsRole get(String name, Output<String> id, @Nullable AuthBackendStsRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendStsRole(name, id, state, options);
    }
}
