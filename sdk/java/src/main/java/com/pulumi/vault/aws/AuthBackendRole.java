// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.aws.AuthBackendRoleArgs;
import com.pulumi.vault.aws.inputs.AuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Manages an AWS auth backend role in a Vault server. Roles constrain the
 * instances or principals that can perform the login operation against the
 * backend. See the [Vault
 * documentation](https://www.vaultproject.io/docs/auth/aws.html) for more
 * information.
 * 
 * ## Example Usage
 * 
 * &lt;!--Start PulumiCodeChooser --&gt;
 * ```java
 * package generated_program;
 * 
 * import com.pulumi.Context;
 * import com.pulumi.Pulumi;
 * import com.pulumi.core.Output;
 * import com.pulumi.vault.AuthBackend;
 * import com.pulumi.vault.AuthBackendArgs;
 * import com.pulumi.vault.aws.AuthBackendRole;
 * import com.pulumi.vault.aws.AuthBackendRoleArgs;
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
 *         var aws = new AuthBackend(&#34;aws&#34;, AuthBackendArgs.builder()        
 *             .type(&#34;aws&#34;)
 *             .build());
 * 
 *         var example = new AuthBackendRole(&#34;example&#34;, AuthBackendRoleArgs.builder()        
 *             .backend(aws.path())
 *             .role(&#34;test-role&#34;)
 *             .authType(&#34;iam&#34;)
 *             .boundAmiIds(&#34;ami-8c1be5f6&#34;)
 *             .boundAccountIds(&#34;123456789012&#34;)
 *             .boundVpcIds(&#34;vpc-b61106d4&#34;)
 *             .boundSubnetIds(&#34;vpc-133128f1&#34;)
 *             .boundIamRoleArns(&#34;arn:aws:iam::123456789012:role/MyRole&#34;)
 *             .boundIamInstanceProfileArns(&#34;arn:aws:iam::123456789012:instance-profile/MyProfile&#34;)
 *             .inferredEntityType(&#34;ec2_instance&#34;)
 *             .inferredAwsRegion(&#34;us-east-1&#34;)
 *             .tokenTtl(60)
 *             .tokenMaxTtl(120)
 *             .tokenPolicies(            
 *                 &#34;default&#34;,
 *                 &#34;dev&#34;,
 *                 &#34;prod&#34;)
 *             .build());
 * 
 *     }
 * }
 * ```
 * &lt;!--End PulumiCodeChooser --&gt;
 * 
 * ## Import
 * 
 * AWS auth backend roles can be imported using `auth/`, the `backend` path, `/role/`, and the `role` name e.g.
 * 
 * ```sh
 * $ pulumi import vault:aws/authBackendRole:AuthBackendRole example auth/aws/role/test-role
 * ```
 * 
 */
@ResourceType(type="vault:aws/authBackendRole:AuthBackendRole")
public class AuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * If set to `true`, allows migration of
     * the underlying instance where the client resides.
     * 
     */
    @Export(name="allowInstanceMigration", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> allowInstanceMigration;

    /**
     * @return If set to `true`, allows migration of
     * the underlying instance where the client resides.
     * 
     */
    public Output<Optional<Boolean>> allowInstanceMigration() {
        return Codegen.optional(this.allowInstanceMigration);
    }
    /**
     * The auth type permitted for this role. Valid choices
     * are `ec2` and `iam`. Defaults to `iam`.
     * 
     */
    @Export(name="authType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> authType;

    /**
     * @return The auth type permitted for this role. Valid choices
     * are `ec2` and `iam`. Defaults to `iam`.
     * 
     */
    public Output<Optional<String>> authType() {
        return Codegen.optional(this.authType);
    }
    /**
     * Path to the mounted aws auth backend.
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Path to the mounted aws auth backend.
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they should be using the
     * account ID specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    @Export(name="boundAccountIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundAccountIds;

    /**
     * @return If set, defines a constraint on the EC2
     * instances that can perform the login operation that they should be using the
     * account ID specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    public Output<Optional<List<String>>> boundAccountIds() {
        return Codegen.optional(this.boundAccountIds);
    }
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they should be using the AMI ID
     * specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    @Export(name="boundAmiIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundAmiIds;

    /**
     * @return If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they should be using the AMI ID
     * specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    public Output<Optional<List<String>>> boundAmiIds() {
        return Codegen.optional(this.boundAmiIds);
    }
    /**
     * Only EC2 instances that match this instance ID will be permitted to log in.
     * 
     */
    @Export(name="boundEc2InstanceIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundEc2InstanceIds;

    /**
     * @return Only EC2 instances that match this instance ID will be permitted to log in.
     * 
     */
    public Output<Optional<List<String>>> boundEc2InstanceIds() {
        return Codegen.optional(this.boundEc2InstanceIds);
    }
    /**
     * If set, defines a constraint on
     * the EC2 instances that can perform the login operation that they must be
     * associated with an IAM instance profile ARN which has a prefix that matches
     * the value specified by this field. The value is prefix-matched as though it
     * were a glob ending in `*`. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    @Export(name="boundIamInstanceProfileArns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundIamInstanceProfileArns;

    /**
     * @return If set, defines a constraint on
     * the EC2 instances that can perform the login operation that they must be
     * associated with an IAM instance profile ARN which has a prefix that matches
     * the value specified by this field. The value is prefix-matched as though it
     * were a glob ending in `*`. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    public Output<Optional<List<String>>> boundIamInstanceProfileArns() {
        return Codegen.optional(this.boundIamInstanceProfileArns);
    }
    /**
     * If set, defines the IAM principal that
     * must be authenticated when `auth_type` is set to `iam`. Wildcards are
     * supported at the end of the ARN.
     * 
     */
    @Export(name="boundIamPrincipalArns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundIamPrincipalArns;

    /**
     * @return If set, defines the IAM principal that
     * must be authenticated when `auth_type` is set to `iam`. Wildcards are
     * supported at the end of the ARN.
     * 
     */
    public Output<Optional<List<String>>> boundIamPrincipalArns() {
        return Codegen.optional(this.boundIamPrincipalArns);
    }
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they must match the IAM
     * role ARN specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    @Export(name="boundIamRoleArns", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundIamRoleArns;

    /**
     * @return If set, defines a constraint on the EC2
     * instances that can perform the login operation that they must match the IAM
     * role ARN specified by this field. `auth_type` must be set to `ec2` or
     * `inferred_entity_type` must be set to `ec2_instance` to use this constraint.
     * 
     */
    public Output<Optional<List<String>>> boundIamRoleArns() {
        return Codegen.optional(this.boundIamRoleArns);
    }
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that the region in their identity
     * document must match the one specified by this field. `auth_type` must be set
     * to `ec2` or `inferred_entity_type` must be set to `ec2_instance` to use this
     * constraint.
     * 
     */
    @Export(name="boundRegions", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundRegions;

    /**
     * @return If set, defines a constraint on the EC2 instances
     * that can perform the login operation that the region in their identity
     * document must match the one specified by this field. `auth_type` must be set
     * to `ec2` or `inferred_entity_type` must be set to `ec2_instance` to use this
     * constraint.
     * 
     */
    public Output<Optional<List<String>>> boundRegions() {
        return Codegen.optional(this.boundRegions);
    }
    /**
     * If set, defines a constraint on the EC2
     * instances that can perform the login operation that they be associated with
     * the subnet ID that matches the value specified by this field. `auth_type`
     * must be set to `ec2` or `inferred_entity_type` must be set to `ec2_instance`
     * to use this constraint.
     * 
     */
    @Export(name="boundSubnetIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundSubnetIds;

    /**
     * @return If set, defines a constraint on the EC2
     * instances that can perform the login operation that they be associated with
     * the subnet ID that matches the value specified by this field. `auth_type`
     * must be set to `ec2` or `inferred_entity_type` must be set to `ec2_instance`
     * to use this constraint.
     * 
     */
    public Output<Optional<List<String>>> boundSubnetIds() {
        return Codegen.optional(this.boundSubnetIds);
    }
    /**
     * If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they be associated with the VPC ID
     * that matches the value specified by this field. `auth_type` must be set to
     * `ec2` or `inferred_entity_type` must be set to `ec2_instance` to use this
     * constraint.
     * 
     */
    @Export(name="boundVpcIds", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> boundVpcIds;

    /**
     * @return If set, defines a constraint on the EC2 instances
     * that can perform the login operation that they be associated with the VPC ID
     * that matches the value specified by this field. `auth_type` must be set to
     * `ec2` or `inferred_entity_type` must be set to `ec2_instance` to use this
     * constraint.
     * 
     */
    public Output<Optional<List<String>>> boundVpcIds() {
        return Codegen.optional(this.boundVpcIds);
    }
    /**
     * IF set to `true`, only allows a
     * single token to be granted per instance ID. This can only be set when
     * `auth_type` is set to `ec2`.
     * 
     */
    @Export(name="disallowReauthentication", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> disallowReauthentication;

    /**
     * @return IF set to `true`, only allows a
     * single token to be granted per instance ID. This can only be set when
     * `auth_type` is set to `ec2`.
     * 
     */
    public Output<Optional<Boolean>> disallowReauthentication() {
        return Codegen.optional(this.disallowReauthentication);
    }
    /**
     * When `inferred_entity_type` is set, this
     * is the region to search for the inferred entities. Required if
     * `inferred_entity_type` is set. This only applies when `auth_type` is set to
     * `iam`.
     * 
     */
    @Export(name="inferredAwsRegion", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inferredAwsRegion;

    /**
     * @return When `inferred_entity_type` is set, this
     * is the region to search for the inferred entities. Required if
     * `inferred_entity_type` is set. This only applies when `auth_type` is set to
     * `iam`.
     * 
     */
    public Output<Optional<String>> inferredAwsRegion() {
        return Codegen.optional(this.inferredAwsRegion);
    }
    /**
     * If set, instructs Vault to turn on
     * inferencing. The only valid value is `ec2_instance`, which instructs Vault to
     * infer that the role comes from an EC2 instance in an IAM instance profile.
     * This only applies when `auth_type` is set to `iam`.
     * 
     */
    @Export(name="inferredEntityType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> inferredEntityType;

    /**
     * @return If set, instructs Vault to turn on
     * inferencing. The only valid value is `ec2_instance`, which instructs Vault to
     * infer that the role comes from an EC2 instance in an IAM instance profile.
     * This only applies when `auth_type` is set to `iam`.
     * 
     */
    public Output<Optional<String>> inferredEntityType() {
        return Codegen.optional(this.inferredEntityType);
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
     * Only valid when
     * `auth_type` is `iam`. If set to `true`, the `bound_iam_principal_arns` are
     * resolved to [AWS Unique
     * IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids)
     * for the bound principal ARN. This field is ignored when a
     * `bound_iam_principal_arn` ends in a wildcard. Resolving to unique IDs more
     * closely mimics the behavior of AWS services in that if an IAM user or role is
     * deleted and a new one is recreated with the same name, those new users or
     * roles won&#39;t get access to roles in Vault that were permissioned to the prior
     * principals of the same name. Defaults to `true`.
     * Once set to `true`, this cannot be changed to `false` without recreating the role.
     * 
     */
    @Export(name="resolveAwsUniqueIds", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> resolveAwsUniqueIds;

    /**
     * @return Only valid when
     * `auth_type` is `iam`. If set to `true`, the `bound_iam_principal_arns` are
     * resolved to [AWS Unique
     * IDs](http://docs.aws.amazon.com/IAM/latest/UserGuide/reference_identifiers.html#identifiers-unique-ids)
     * for the bound principal ARN. This field is ignored when a
     * `bound_iam_principal_arn` ends in a wildcard. Resolving to unique IDs more
     * closely mimics the behavior of AWS services in that if an IAM user or role is
     * deleted and a new one is recreated with the same name, those new users or
     * roles won&#39;t get access to roles in Vault that were permissioned to the prior
     * principals of the same name. Defaults to `true`.
     * Once set to `true`, this cannot be changed to `false` without recreating the role.
     * 
     */
    public Output<Optional<Boolean>> resolveAwsUniqueIds() {
        return Codegen.optional(this.resolveAwsUniqueIds);
    }
    /**
     * The name of the role.
     * 
     */
    @Export(name="role", refs={String.class}, tree="[0]")
    private Output<String> role;

    /**
     * @return The name of the role.
     * 
     */
    public Output<String> role() {
        return this.role;
    }
    /**
     * The Vault generated role ID.
     * 
     */
    @Export(name="roleId", refs={String.class}, tree="[0]")
    private Output<String> roleId;

    /**
     * @return The Vault generated role ID.
     * 
     */
    public Output<String> roleId() {
        return this.roleId;
    }
    /**
     * If set, enable role tags for this role. The value set
     * for this field should be the key of the tag on the EC2 instance. `auth_type`
     * must be set to `ec2` or `inferred_entity_type` must be set to `ec2_instance`
     * to use this constraint.
     * 
     */
    @Export(name="roleTag", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> roleTag;

    /**
     * @return If set, enable role tags for this role. The value set
     * for this field should be the key of the tag on the EC2 instance. `auth_type`
     * must be set to `ec2` or `inferred_entity_type` must be set to `ec2_instance`
     * to use this constraint.
     * 
     */
    public Output<Optional<String>> roleTag() {
        return Codegen.optional(this.roleTag);
    }
    /**
     * Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    @Export(name="tokenBoundCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return Specifies the blocks of IP addresses which are allowed to use the generated token
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    @Export(name="tokenExplicitMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return Generated Token&#39;s Explicit Maximum TTL in seconds
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * The maximum lifetime of the generated token
     * 
     */
    @Export(name="tokenMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime of the generated token
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    @Export(name="tokenNoDefaultPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return If true, the &#39;default&#39; policy will not automatically be added to generated tokens
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    @Export(name="tokenNumUses", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The maximum number of times a token may be used, a value of zero means unlimited
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * Generated Token&#39;s Period
     * 
     */
    @Export(name="tokenPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return Generated Token&#39;s Period
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * Generated Token&#39;s Policies
     * 
     */
    @Export(name="tokenPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return Generated Token&#39;s Policies
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The initial ttl of the token to generate in seconds
     * 
     */
    @Export(name="tokenTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The initial ttl of the token to generate in seconds
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * The type of token to generate, service or batch
     * 
     */
    @Export(name="tokenType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return The type of token to generate, service or batch
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public AuthBackendRole(String name) {
        this(name, AuthBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public AuthBackendRole(String name, AuthBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public AuthBackendRole(String name, AuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendRole:AuthBackendRole", name, args == null ? AuthBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private AuthBackendRole(String name, Output<String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:aws/authBackendRole:AuthBackendRole", name, state, makeResourceOptions(options, id));
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
    public static AuthBackendRole get(String name, Output<String> id, @Nullable AuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new AuthBackendRole(name, id, state, options);
    }
}
