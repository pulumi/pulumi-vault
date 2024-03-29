// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginAwsArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginAwsArgs Empty = new ProviderAuthLoginAwsArgs();

    /**
     * The AWS access key ID.
     * 
     */
    @Import(name="awsAccessKeyId")
    private @Nullable Output<String> awsAccessKeyId;

    /**
     * @return The AWS access key ID.
     * 
     */
    public Optional<Output<String>> awsAccessKeyId() {
        return Optional.ofNullable(this.awsAccessKeyId);
    }

    /**
     * The IAM endpoint URL.
     * 
     */
    @Import(name="awsIamEndpoint")
    private @Nullable Output<String> awsIamEndpoint;

    /**
     * @return The IAM endpoint URL.
     * 
     */
    public Optional<Output<String>> awsIamEndpoint() {
        return Optional.ofNullable(this.awsIamEndpoint);
    }

    /**
     * The name of the AWS profile.
     * 
     */
    @Import(name="awsProfile")
    private @Nullable Output<String> awsProfile;

    /**
     * @return The name of the AWS profile.
     * 
     */
    public Optional<Output<String>> awsProfile() {
        return Optional.ofNullable(this.awsProfile);
    }

    /**
     * The AWS region.
     * 
     */
    @Import(name="awsRegion")
    private @Nullable Output<String> awsRegion;

    /**
     * @return The AWS region.
     * 
     */
    public Optional<Output<String>> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }

    /**
     * The ARN of the AWS Role to assume.Used during STS AssumeRole
     * 
     */
    @Import(name="awsRoleArn")
    private @Nullable Output<String> awsRoleArn;

    /**
     * @return The ARN of the AWS Role to assume.Used during STS AssumeRole
     * 
     */
    public Optional<Output<String>> awsRoleArn() {
        return Optional.ofNullable(this.awsRoleArn);
    }

    /**
     * Specifies the name to attach to the AWS role session. Used during STS AssumeRole
     * 
     */
    @Import(name="awsRoleSessionName")
    private @Nullable Output<String> awsRoleSessionName;

    /**
     * @return Specifies the name to attach to the AWS role session. Used during STS AssumeRole
     * 
     */
    public Optional<Output<String>> awsRoleSessionName() {
        return Optional.ofNullable(this.awsRoleSessionName);
    }

    /**
     * The AWS secret access key.
     * 
     */
    @Import(name="awsSecretAccessKey")
    private @Nullable Output<String> awsSecretAccessKey;

    /**
     * @return The AWS secret access key.
     * 
     */
    public Optional<Output<String>> awsSecretAccessKey() {
        return Optional.ofNullable(this.awsSecretAccessKey);
    }

    /**
     * The AWS session token.
     * 
     */
    @Import(name="awsSessionToken")
    private @Nullable Output<String> awsSessionToken;

    /**
     * @return The AWS session token.
     * 
     */
    public Optional<Output<String>> awsSessionToken() {
        return Optional.ofNullable(this.awsSessionToken);
    }

    /**
     * Path to the AWS shared credentials file.
     * 
     */
    @Import(name="awsSharedCredentialsFile")
    private @Nullable Output<String> awsSharedCredentialsFile;

    /**
     * @return Path to the AWS shared credentials file.
     * 
     */
    public Optional<Output<String>> awsSharedCredentialsFile() {
        return Optional.ofNullable(this.awsSharedCredentialsFile);
    }

    /**
     * The STS endpoint URL.
     * 
     */
    @Import(name="awsStsEndpoint")
    private @Nullable Output<String> awsStsEndpoint;

    /**
     * @return The STS endpoint URL.
     * 
     */
    public Optional<Output<String>> awsStsEndpoint() {
        return Optional.ofNullable(this.awsStsEndpoint);
    }

    /**
     * Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
     * 
     */
    @Import(name="awsWebIdentityTokenFile")
    private @Nullable Output<String> awsWebIdentityTokenFile;

    /**
     * @return Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
     * 
     */
    public Optional<Output<String>> awsWebIdentityTokenFile() {
        return Optional.ofNullable(this.awsWebIdentityTokenFile);
    }

    /**
     * The Vault header value to include in the STS signing request.
     * 
     */
    @Import(name="headerValue")
    private @Nullable Output<String> headerValue;

    /**
     * @return The Vault header value to include in the STS signing request.
     * 
     */
    public Optional<Output<String>> headerValue() {
        return Optional.ofNullable(this.headerValue);
    }

    /**
     * The path where the authentication engine is mounted.
     * 
     */
    @Import(name="mount")
    private @Nullable Output<String> mount;

    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    /**
     * The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * The Vault role to use when logging into Vault.
     * 
     */
    @Import(name="role", required=true)
    private Output<String> role;

    /**
     * @return The Vault role to use when logging into Vault.
     * 
     */
    public Output<String> role() {
        return this.role;
    }

    /**
     * Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    private ProviderAuthLoginAwsArgs() {}

    private ProviderAuthLoginAwsArgs(ProviderAuthLoginAwsArgs $) {
        this.awsAccessKeyId = $.awsAccessKeyId;
        this.awsIamEndpoint = $.awsIamEndpoint;
        this.awsProfile = $.awsProfile;
        this.awsRegion = $.awsRegion;
        this.awsRoleArn = $.awsRoleArn;
        this.awsRoleSessionName = $.awsRoleSessionName;
        this.awsSecretAccessKey = $.awsSecretAccessKey;
        this.awsSessionToken = $.awsSessionToken;
        this.awsSharedCredentialsFile = $.awsSharedCredentialsFile;
        this.awsStsEndpoint = $.awsStsEndpoint;
        this.awsWebIdentityTokenFile = $.awsWebIdentityTokenFile;
        this.headerValue = $.headerValue;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.role = $.role;
        this.useRootNamespace = $.useRootNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginAwsArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginAwsArgs $;

        public Builder() {
            $ = new ProviderAuthLoginAwsArgs();
        }

        public Builder(ProviderAuthLoginAwsArgs defaults) {
            $ = new ProviderAuthLoginAwsArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param awsAccessKeyId The AWS access key ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccessKeyId(@Nullable Output<String> awsAccessKeyId) {
            $.awsAccessKeyId = awsAccessKeyId;
            return this;
        }

        /**
         * @param awsAccessKeyId The AWS access key ID.
         * 
         * @return builder
         * 
         */
        public Builder awsAccessKeyId(String awsAccessKeyId) {
            return awsAccessKeyId(Output.of(awsAccessKeyId));
        }

        /**
         * @param awsIamEndpoint The IAM endpoint URL.
         * 
         * @return builder
         * 
         */
        public Builder awsIamEndpoint(@Nullable Output<String> awsIamEndpoint) {
            $.awsIamEndpoint = awsIamEndpoint;
            return this;
        }

        /**
         * @param awsIamEndpoint The IAM endpoint URL.
         * 
         * @return builder
         * 
         */
        public Builder awsIamEndpoint(String awsIamEndpoint) {
            return awsIamEndpoint(Output.of(awsIamEndpoint));
        }

        /**
         * @param awsProfile The name of the AWS profile.
         * 
         * @return builder
         * 
         */
        public Builder awsProfile(@Nullable Output<String> awsProfile) {
            $.awsProfile = awsProfile;
            return this;
        }

        /**
         * @param awsProfile The name of the AWS profile.
         * 
         * @return builder
         * 
         */
        public Builder awsProfile(String awsProfile) {
            return awsProfile(Output.of(awsProfile));
        }

        /**
         * @param awsRegion The AWS region.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(@Nullable Output<String> awsRegion) {
            $.awsRegion = awsRegion;
            return this;
        }

        /**
         * @param awsRegion The AWS region.
         * 
         * @return builder
         * 
         */
        public Builder awsRegion(String awsRegion) {
            return awsRegion(Output.of(awsRegion));
        }

        /**
         * @param awsRoleArn The ARN of the AWS Role to assume.Used during STS AssumeRole
         * 
         * @return builder
         * 
         */
        public Builder awsRoleArn(@Nullable Output<String> awsRoleArn) {
            $.awsRoleArn = awsRoleArn;
            return this;
        }

        /**
         * @param awsRoleArn The ARN of the AWS Role to assume.Used during STS AssumeRole
         * 
         * @return builder
         * 
         */
        public Builder awsRoleArn(String awsRoleArn) {
            return awsRoleArn(Output.of(awsRoleArn));
        }

        /**
         * @param awsRoleSessionName Specifies the name to attach to the AWS role session. Used during STS AssumeRole
         * 
         * @return builder
         * 
         */
        public Builder awsRoleSessionName(@Nullable Output<String> awsRoleSessionName) {
            $.awsRoleSessionName = awsRoleSessionName;
            return this;
        }

        /**
         * @param awsRoleSessionName Specifies the name to attach to the AWS role session. Used during STS AssumeRole
         * 
         * @return builder
         * 
         */
        public Builder awsRoleSessionName(String awsRoleSessionName) {
            return awsRoleSessionName(Output.of(awsRoleSessionName));
        }

        /**
         * @param awsSecretAccessKey The AWS secret access key.
         * 
         * @return builder
         * 
         */
        public Builder awsSecretAccessKey(@Nullable Output<String> awsSecretAccessKey) {
            $.awsSecretAccessKey = awsSecretAccessKey;
            return this;
        }

        /**
         * @param awsSecretAccessKey The AWS secret access key.
         * 
         * @return builder
         * 
         */
        public Builder awsSecretAccessKey(String awsSecretAccessKey) {
            return awsSecretAccessKey(Output.of(awsSecretAccessKey));
        }

        /**
         * @param awsSessionToken The AWS session token.
         * 
         * @return builder
         * 
         */
        public Builder awsSessionToken(@Nullable Output<String> awsSessionToken) {
            $.awsSessionToken = awsSessionToken;
            return this;
        }

        /**
         * @param awsSessionToken The AWS session token.
         * 
         * @return builder
         * 
         */
        public Builder awsSessionToken(String awsSessionToken) {
            return awsSessionToken(Output.of(awsSessionToken));
        }

        /**
         * @param awsSharedCredentialsFile Path to the AWS shared credentials file.
         * 
         * @return builder
         * 
         */
        public Builder awsSharedCredentialsFile(@Nullable Output<String> awsSharedCredentialsFile) {
            $.awsSharedCredentialsFile = awsSharedCredentialsFile;
            return this;
        }

        /**
         * @param awsSharedCredentialsFile Path to the AWS shared credentials file.
         * 
         * @return builder
         * 
         */
        public Builder awsSharedCredentialsFile(String awsSharedCredentialsFile) {
            return awsSharedCredentialsFile(Output.of(awsSharedCredentialsFile));
        }

        /**
         * @param awsStsEndpoint The STS endpoint URL.
         * 
         * @return builder
         * 
         */
        public Builder awsStsEndpoint(@Nullable Output<String> awsStsEndpoint) {
            $.awsStsEndpoint = awsStsEndpoint;
            return this;
        }

        /**
         * @param awsStsEndpoint The STS endpoint URL.
         * 
         * @return builder
         * 
         */
        public Builder awsStsEndpoint(String awsStsEndpoint) {
            return awsStsEndpoint(Output.of(awsStsEndpoint));
        }

        /**
         * @param awsWebIdentityTokenFile Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
         * 
         * @return builder
         * 
         */
        public Builder awsWebIdentityTokenFile(@Nullable Output<String> awsWebIdentityTokenFile) {
            $.awsWebIdentityTokenFile = awsWebIdentityTokenFile;
            return this;
        }

        /**
         * @param awsWebIdentityTokenFile Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
         * 
         * @return builder
         * 
         */
        public Builder awsWebIdentityTokenFile(String awsWebIdentityTokenFile) {
            return awsWebIdentityTokenFile(Output.of(awsWebIdentityTokenFile));
        }

        /**
         * @param headerValue The Vault header value to include in the STS signing request.
         * 
         * @return builder
         * 
         */
        public Builder headerValue(@Nullable Output<String> headerValue) {
            $.headerValue = headerValue;
            return this;
        }

        /**
         * @param headerValue The Vault header value to include in the STS signing request.
         * 
         * @return builder
         * 
         */
        public Builder headerValue(String headerValue) {
            return headerValue(Output.of(headerValue));
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        /**
         * @param mount The path where the authentication engine is mounted.
         * 
         * @return builder
         * 
         */
        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        /**
         * @param namespace The authentication engine&#39;s namespace. Conflicts with use_root_namespace
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param role The Vault role to use when logging into Vault.
         * 
         * @return builder
         * 
         */
        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        /**
         * @param role The Vault role to use when logging into Vault.
         * 
         * @return builder
         * 
         */
        public Builder role(String role) {
            return role(Output.of(role));
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        /**
         * @param useRootNamespace Authenticate to the root Vault namespace. Conflicts with namespace
         * 
         * @return builder
         * 
         */
        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        public ProviderAuthLoginAwsArgs build() {
            if ($.role == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginAwsArgs", "role");
            }
            return $;
        }
    }

}
