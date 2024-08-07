// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.aws.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SecretBackendRoleState extends com.pulumi.resources.ResourceArgs {

    public static final SecretBackendRoleState Empty = new SecretBackendRoleState();

    /**
     * The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    @Import(name="backend")
    private @Nullable Output<String> backend;

    /**
     * @return The path the AWS secret backend is mounted at,
     * with no leading or trailing `/`s.
     * 
     */
    public Optional<Output<String>> backend() {
        return Optional.ofNullable(this.backend);
    }

    /**
     * Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
     * `federation_token`.
     * 
     */
    @Import(name="credentialType")
    private @Nullable Output<String> credentialType;

    /**
     * @return Specifies the type of credential to be used when
     * retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
     * `federation_token`.
     * 
     */
    public Optional<Output<String>> credentialType() {
        return Optional.ofNullable(this.credentialType);
    }

    /**
     * The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credential_type` is one of
     * `assumed_role` or `federation_token`.
     * 
     */
    @Import(name="defaultStsTtl")
    private @Nullable Output<Integer> defaultStsTtl;

    /**
     * @return The default TTL in seconds for STS credentials.
     * When a TTL is not specified when STS credentials are requested,
     * and a default TTL is specified on the role,
     * then this default TTL will be used. Valid only when `credential_type` is one of
     * `assumed_role` or `federation_token`.
     * 
     */
    public Optional<Output<Integer>> defaultStsTtl() {
        return Optional.ofNullable(this.defaultStsTtl);
    }

    /**
     * External ID to set for assume role creds.
     * Valid only when `credential_type` is set to `assumed_role`.
     * 
     */
    @Import(name="externalId")
    private @Nullable Output<String> externalId;

    /**
     * @return External ID to set for assume role creds.
     * Valid only when `credential_type` is set to `assumed_role`.
     * 
     */
    public Optional<Output<String>> externalId() {
        return Optional.ofNullable(this.externalId);
    }

    /**
     * A list of IAM group names. IAM users generated
     * against this vault role will be added to these IAM Groups. For a credential
     * type of `assumed_role` or `federation_token`, the policies sent to the
     * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
     * policies from each group in `iam_groups` combined with the `policy_document`
     * and `policy_arns` parameters.
     * 
     */
    @Import(name="iamGroups")
    private @Nullable Output<List<String>> iamGroups;

    /**
     * @return A list of IAM group names. IAM users generated
     * against this vault role will be added to these IAM Groups. For a credential
     * type of `assumed_role` or `federation_token`, the policies sent to the
     * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
     * policies from each group in `iam_groups` combined with the `policy_document`
     * and `policy_arns` parameters.
     * 
     */
    public Optional<Output<List<String>>> iamGroups() {
        return Optional.ofNullable(this.iamGroups);
    }

    /**
     * A map of strings representing key/value pairs
     * to be used as tags for any IAM user that is created by this role.
     * 
     */
    @Import(name="iamTags")
    private @Nullable Output<Map<String,String>> iamTags;

    /**
     * @return A map of strings representing key/value pairs
     * to be used as tags for any IAM user that is created by this role.
     * 
     */
    public Optional<Output<Map<String,String>>> iamTags() {
        return Optional.ofNullable(this.iamTags);
    }

    /**
     * The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
     * one of `assumed_role` or `federation_token`.
     * 
     */
    @Import(name="maxStsTtl")
    private @Nullable Output<Integer> maxStsTtl;

    /**
     * @return The max allowed TTL in seconds for STS credentials
     * (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
     * one of `assumed_role` or `federation_token`.
     * 
     */
    public Optional<Output<Integer>> maxStsTtl() {
        return Optional.ofNullable(this.maxStsTtl);
    }

    /**
     * The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return The name to identify this role within the backend.
     * Must be unique within the backend.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The ARN of the AWS Permissions
     * Boundary to attach to IAM users created in the role. Valid only when
     * `credential_type` is `iam_user`. If not specified, then no permissions boundary
     * policy will be attached.
     * 
     */
    @Import(name="permissionsBoundaryArn")
    private @Nullable Output<String> permissionsBoundaryArn;

    /**
     * @return The ARN of the AWS Permissions
     * Boundary to attach to IAM users created in the role. Valid only when
     * `credential_type` is `iam_user`. If not specified, then no permissions boundary
     * policy will be attached.
     * 
     */
    public Optional<Output<String>> permissionsBoundaryArn() {
        return Optional.ofNullable(this.permissionsBoundaryArn);
    }

    /**
     * Specifies a list of AWS managed policy ARNs. The
     * behavior depends on the credential type. With `iam_user`, the policies will be
     * attached to IAM users when they are requested. With `assumed_role` and
     * `federation_token`, the policy ARNs will act as a filter on what the credentials
     * can do, similar to `policy_document`. When `credential_type` is `iam_user` or
     * `federation_token`, at least one of `policy_document` or `policy_arns` must
     * be specified.
     * 
     */
    @Import(name="policyArns")
    private @Nullable Output<List<String>> policyArns;

    /**
     * @return Specifies a list of AWS managed policy ARNs. The
     * behavior depends on the credential type. With `iam_user`, the policies will be
     * attached to IAM users when they are requested. With `assumed_role` and
     * `federation_token`, the policy ARNs will act as a filter on what the credentials
     * can do, similar to `policy_document`. When `credential_type` is `iam_user` or
     * `federation_token`, at least one of `policy_document` or `policy_arns` must
     * be specified.
     * 
     */
    public Optional<Output<List<String>>> policyArns() {
        return Optional.ofNullable(this.policyArns);
    }

    /**
     * The IAM policy document for the role. The
     * behavior depends on the credential type. With `iam_user`, the policy document
     * will be attached to the IAM user generated and augment the permissions the IAM
     * user has. With `assumed_role` and `federation_token`, the policy document will
     * act as a filter on what the credentials can do, similar to `policy_arns`.
     * 
     */
    @Import(name="policyDocument")
    private @Nullable Output<String> policyDocument;

    /**
     * @return The IAM policy document for the role. The
     * behavior depends on the credential type. With `iam_user`, the policy document
     * will be attached to the IAM user generated and augment the permissions the IAM
     * user has. With `assumed_role` and `federation_token`, the policy document will
     * act as a filter on what the credentials can do, similar to `policy_arns`.
     * 
     */
    public Optional<Output<String>> policyDocument() {
        return Optional.ofNullable(this.policyDocument);
    }

    /**
     * Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credential_type` is `assumed_role` and
     * prohibited otherwise.
     * 
     */
    @Import(name="roleArns")
    private @Nullable Output<List<String>> roleArns;

    /**
     * @return Specifies the ARNs of the AWS roles this Vault role
     * is allowed to assume. Required when `credential_type` is `assumed_role` and
     * prohibited otherwise.
     * 
     */
    public Optional<Output<List<String>>> roleArns() {
        return Optional.ofNullable(this.roleArns);
    }

    /**
     * A map of strings representing key/value pairs to be set
     * during assume role creds creation. Valid only when `credential_type` is set to
     * `assumed_role`.
     * 
     */
    @Import(name="sessionTags")
    private @Nullable Output<Map<String,String>> sessionTags;

    /**
     * @return A map of strings representing key/value pairs to be set
     * during assume role creds creation. Valid only when `credential_type` is set to
     * `assumed_role`.
     * 
     */
    public Optional<Output<Map<String,String>>> sessionTags() {
        return Optional.ofNullable(this.sessionTags);
    }

    /**
     * The path for the user name. Valid only when
     * `credential_type` is `iam_user`. Default is `/`.
     * 
     */
    @Import(name="userPath")
    private @Nullable Output<String> userPath;

    /**
     * @return The path for the user name. Valid only when
     * `credential_type` is `iam_user`. Default is `/`.
     * 
     */
    public Optional<Output<String>> userPath() {
        return Optional.ofNullable(this.userPath);
    }

    private SecretBackendRoleState() {}

    private SecretBackendRoleState(SecretBackendRoleState $) {
        this.backend = $.backend;
        this.credentialType = $.credentialType;
        this.defaultStsTtl = $.defaultStsTtl;
        this.externalId = $.externalId;
        this.iamGroups = $.iamGroups;
        this.iamTags = $.iamTags;
        this.maxStsTtl = $.maxStsTtl;
        this.name = $.name;
        this.namespace = $.namespace;
        this.permissionsBoundaryArn = $.permissionsBoundaryArn;
        this.policyArns = $.policyArns;
        this.policyDocument = $.policyDocument;
        this.roleArns = $.roleArns;
        this.sessionTags = $.sessionTags;
        this.userPath = $.userPath;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SecretBackendRoleState defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SecretBackendRoleState $;

        public Builder() {
            $ = new SecretBackendRoleState();
        }

        public Builder(SecretBackendRoleState defaults) {
            $ = new SecretBackendRoleState(Objects.requireNonNull(defaults));
        }

        /**
         * @param backend The path the AWS secret backend is mounted at,
         * with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(@Nullable Output<String> backend) {
            $.backend = backend;
            return this;
        }

        /**
         * @param backend The path the AWS secret backend is mounted at,
         * with no leading or trailing `/`s.
         * 
         * @return builder
         * 
         */
        public Builder backend(String backend) {
            return backend(Output.of(backend));
        }

        /**
         * @param credentialType Specifies the type of credential to be used when
         * retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
         * `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(@Nullable Output<String> credentialType) {
            $.credentialType = credentialType;
            return this;
        }

        /**
         * @param credentialType Specifies the type of credential to be used when
         * retrieving credentials from the role. Must be one of `iam_user`, `assumed_role`, or
         * `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder credentialType(String credentialType) {
            return credentialType(Output.of(credentialType));
        }

        /**
         * @param defaultStsTtl The default TTL in seconds for STS credentials.
         * When a TTL is not specified when STS credentials are requested,
         * and a default TTL is specified on the role,
         * then this default TTL will be used. Valid only when `credential_type` is one of
         * `assumed_role` or `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder defaultStsTtl(@Nullable Output<Integer> defaultStsTtl) {
            $.defaultStsTtl = defaultStsTtl;
            return this;
        }

        /**
         * @param defaultStsTtl The default TTL in seconds for STS credentials.
         * When a TTL is not specified when STS credentials are requested,
         * and a default TTL is specified on the role,
         * then this default TTL will be used. Valid only when `credential_type` is one of
         * `assumed_role` or `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder defaultStsTtl(Integer defaultStsTtl) {
            return defaultStsTtl(Output.of(defaultStsTtl));
        }

        /**
         * @param externalId External ID to set for assume role creds.
         * Valid only when `credential_type` is set to `assumed_role`.
         * 
         * @return builder
         * 
         */
        public Builder externalId(@Nullable Output<String> externalId) {
            $.externalId = externalId;
            return this;
        }

        /**
         * @param externalId External ID to set for assume role creds.
         * Valid only when `credential_type` is set to `assumed_role`.
         * 
         * @return builder
         * 
         */
        public Builder externalId(String externalId) {
            return externalId(Output.of(externalId));
        }

        /**
         * @param iamGroups A list of IAM group names. IAM users generated
         * against this vault role will be added to these IAM Groups. For a credential
         * type of `assumed_role` or `federation_token`, the policies sent to the
         * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
         * policies from each group in `iam_groups` combined with the `policy_document`
         * and `policy_arns` parameters.
         * 
         * @return builder
         * 
         */
        public Builder iamGroups(@Nullable Output<List<String>> iamGroups) {
            $.iamGroups = iamGroups;
            return this;
        }

        /**
         * @param iamGroups A list of IAM group names. IAM users generated
         * against this vault role will be added to these IAM Groups. For a credential
         * type of `assumed_role` or `federation_token`, the policies sent to the
         * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
         * policies from each group in `iam_groups` combined with the `policy_document`
         * and `policy_arns` parameters.
         * 
         * @return builder
         * 
         */
        public Builder iamGroups(List<String> iamGroups) {
            return iamGroups(Output.of(iamGroups));
        }

        /**
         * @param iamGroups A list of IAM group names. IAM users generated
         * against this vault role will be added to these IAM Groups. For a credential
         * type of `assumed_role` or `federation_token`, the policies sent to the
         * corresponding AWS call (sts:AssumeRole or sts:GetFederation) will be the
         * policies from each group in `iam_groups` combined with the `policy_document`
         * and `policy_arns` parameters.
         * 
         * @return builder
         * 
         */
        public Builder iamGroups(String... iamGroups) {
            return iamGroups(List.of(iamGroups));
        }

        /**
         * @param iamTags A map of strings representing key/value pairs
         * to be used as tags for any IAM user that is created by this role.
         * 
         * @return builder
         * 
         */
        public Builder iamTags(@Nullable Output<Map<String,String>> iamTags) {
            $.iamTags = iamTags;
            return this;
        }

        /**
         * @param iamTags A map of strings representing key/value pairs
         * to be used as tags for any IAM user that is created by this role.
         * 
         * @return builder
         * 
         */
        public Builder iamTags(Map<String,String> iamTags) {
            return iamTags(Output.of(iamTags));
        }

        /**
         * @param maxStsTtl The max allowed TTL in seconds for STS credentials
         * (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
         * one of `assumed_role` or `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder maxStsTtl(@Nullable Output<Integer> maxStsTtl) {
            $.maxStsTtl = maxStsTtl;
            return this;
        }

        /**
         * @param maxStsTtl The max allowed TTL in seconds for STS credentials
         * (credentials TTL are capped to `max_sts_ttl`). Valid only when `credential_type` is
         * one of `assumed_role` or `federation_token`.
         * 
         * @return builder
         * 
         */
        public Builder maxStsTtl(Integer maxStsTtl) {
            return maxStsTtl(Output.of(maxStsTtl));
        }

        /**
         * @param name The name to identify this role within the backend.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name The name to identify this role within the backend.
         * Must be unique within the backend.
         * 
         * @return builder
         * 
         */
        public Builder name(String name) {
            return name(Output.of(name));
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The namespace to provision the resource in.
         * The value should not contain leading or trailing forward slashes.
         * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
         * *Available only for Vault Enterprise*.
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param permissionsBoundaryArn The ARN of the AWS Permissions
         * Boundary to attach to IAM users created in the role. Valid only when
         * `credential_type` is `iam_user`. If not specified, then no permissions boundary
         * policy will be attached.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundaryArn(@Nullable Output<String> permissionsBoundaryArn) {
            $.permissionsBoundaryArn = permissionsBoundaryArn;
            return this;
        }

        /**
         * @param permissionsBoundaryArn The ARN of the AWS Permissions
         * Boundary to attach to IAM users created in the role. Valid only when
         * `credential_type` is `iam_user`. If not specified, then no permissions boundary
         * policy will be attached.
         * 
         * @return builder
         * 
         */
        public Builder permissionsBoundaryArn(String permissionsBoundaryArn) {
            return permissionsBoundaryArn(Output.of(permissionsBoundaryArn));
        }

        /**
         * @param policyArns Specifies a list of AWS managed policy ARNs. The
         * behavior depends on the credential type. With `iam_user`, the policies will be
         * attached to IAM users when they are requested. With `assumed_role` and
         * `federation_token`, the policy ARNs will act as a filter on what the credentials
         * can do, similar to `policy_document`. When `credential_type` is `iam_user` or
         * `federation_token`, at least one of `policy_document` or `policy_arns` must
         * be specified.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(@Nullable Output<List<String>> policyArns) {
            $.policyArns = policyArns;
            return this;
        }

        /**
         * @param policyArns Specifies a list of AWS managed policy ARNs. The
         * behavior depends on the credential type. With `iam_user`, the policies will be
         * attached to IAM users when they are requested. With `assumed_role` and
         * `federation_token`, the policy ARNs will act as a filter on what the credentials
         * can do, similar to `policy_document`. When `credential_type` is `iam_user` or
         * `federation_token`, at least one of `policy_document` or `policy_arns` must
         * be specified.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(List<String> policyArns) {
            return policyArns(Output.of(policyArns));
        }

        /**
         * @param policyArns Specifies a list of AWS managed policy ARNs. The
         * behavior depends on the credential type. With `iam_user`, the policies will be
         * attached to IAM users when they are requested. With `assumed_role` and
         * `federation_token`, the policy ARNs will act as a filter on what the credentials
         * can do, similar to `policy_document`. When `credential_type` is `iam_user` or
         * `federation_token`, at least one of `policy_document` or `policy_arns` must
         * be specified.
         * 
         * @return builder
         * 
         */
        public Builder policyArns(String... policyArns) {
            return policyArns(List.of(policyArns));
        }

        /**
         * @param policyDocument The IAM policy document for the role. The
         * behavior depends on the credential type. With `iam_user`, the policy document
         * will be attached to the IAM user generated and augment the permissions the IAM
         * user has. With `assumed_role` and `federation_token`, the policy document will
         * act as a filter on what the credentials can do, similar to `policy_arns`.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(@Nullable Output<String> policyDocument) {
            $.policyDocument = policyDocument;
            return this;
        }

        /**
         * @param policyDocument The IAM policy document for the role. The
         * behavior depends on the credential type. With `iam_user`, the policy document
         * will be attached to the IAM user generated and augment the permissions the IAM
         * user has. With `assumed_role` and `federation_token`, the policy document will
         * act as a filter on what the credentials can do, similar to `policy_arns`.
         * 
         * @return builder
         * 
         */
        public Builder policyDocument(String policyDocument) {
            return policyDocument(Output.of(policyDocument));
        }

        /**
         * @param roleArns Specifies the ARNs of the AWS roles this Vault role
         * is allowed to assume. Required when `credential_type` is `assumed_role` and
         * prohibited otherwise.
         * 
         * @return builder
         * 
         */
        public Builder roleArns(@Nullable Output<List<String>> roleArns) {
            $.roleArns = roleArns;
            return this;
        }

        /**
         * @param roleArns Specifies the ARNs of the AWS roles this Vault role
         * is allowed to assume. Required when `credential_type` is `assumed_role` and
         * prohibited otherwise.
         * 
         * @return builder
         * 
         */
        public Builder roleArns(List<String> roleArns) {
            return roleArns(Output.of(roleArns));
        }

        /**
         * @param roleArns Specifies the ARNs of the AWS roles this Vault role
         * is allowed to assume. Required when `credential_type` is `assumed_role` and
         * prohibited otherwise.
         * 
         * @return builder
         * 
         */
        public Builder roleArns(String... roleArns) {
            return roleArns(List.of(roleArns));
        }

        /**
         * @param sessionTags A map of strings representing key/value pairs to be set
         * during assume role creds creation. Valid only when `credential_type` is set to
         * `assumed_role`.
         * 
         * @return builder
         * 
         */
        public Builder sessionTags(@Nullable Output<Map<String,String>> sessionTags) {
            $.sessionTags = sessionTags;
            return this;
        }

        /**
         * @param sessionTags A map of strings representing key/value pairs to be set
         * during assume role creds creation. Valid only when `credential_type` is set to
         * `assumed_role`.
         * 
         * @return builder
         * 
         */
        public Builder sessionTags(Map<String,String> sessionTags) {
            return sessionTags(Output.of(sessionTags));
        }

        /**
         * @param userPath The path for the user name. Valid only when
         * `credential_type` is `iam_user`. Default is `/`.
         * 
         * @return builder
         * 
         */
        public Builder userPath(@Nullable Output<String> userPath) {
            $.userPath = userPath;
            return this;
        }

        /**
         * @param userPath The path for the user name. Valid only when
         * `credential_type` is `iam_user`. Default is `/`.
         * 
         * @return builder
         * 
         */
        public Builder userPath(String userPath) {
            return userPath(Output.of(userPath));
        }

        public SecretBackendRoleState build() {
            return $;
        }
    }

}
