// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.config.inputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class AuthLoginAws {
    /**
     * @return The AWS access key ID.
     * 
     */
    private @Nullable String awsAccessKeyId;
    /**
     * @return The IAM endpoint URL.
     * 
     */
    private @Nullable String awsIamEndpoint;
    /**
     * @return The name of the AWS profile.
     * 
     */
    private @Nullable String awsProfile;
    /**
     * @return The AWS region.
     * 
     */
    private @Nullable String awsRegion;
    /**
     * @return The ARN of the AWS Role to assume.Used during STS AssumeRole
     * 
     */
    private @Nullable String awsRoleArn;
    /**
     * @return Specifies the name to attach to the AWS role session. Used during STS AssumeRole
     * 
     */
    private @Nullable String awsRoleSessionName;
    /**
     * @return The AWS secret access key.
     * 
     */
    private @Nullable String awsSecretAccessKey;
    /**
     * @return The AWS session token.
     * 
     */
    private @Nullable String awsSessionToken;
    /**
     * @return Path to the AWS shared credentials file.
     * 
     */
    private @Nullable String awsSharedCredentialsFile;
    /**
     * @return The STS endpoint URL.
     * 
     */
    private @Nullable String awsStsEndpoint;
    /**
     * @return Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
     * 
     */
    private @Nullable String awsWebIdentityTokenFile;
    /**
     * @return The Vault header value to include in the STS signing request.
     * 
     */
    private @Nullable String headerValue;
    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    private @Nullable String mount;
    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    private @Nullable String namespace;
    /**
     * @return The Vault role to use when logging into Vault.
     * 
     */
    private String role;
    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    private @Nullable Boolean useRootNamespace;

    private AuthLoginAws() {}
    /**
     * @return The AWS access key ID.
     * 
     */
    public Optional<String> awsAccessKeyId() {
        return Optional.ofNullable(this.awsAccessKeyId);
    }
    /**
     * @return The IAM endpoint URL.
     * 
     */
    public Optional<String> awsIamEndpoint() {
        return Optional.ofNullable(this.awsIamEndpoint);
    }
    /**
     * @return The name of the AWS profile.
     * 
     */
    public Optional<String> awsProfile() {
        return Optional.ofNullable(this.awsProfile);
    }
    /**
     * @return The AWS region.
     * 
     */
    public Optional<String> awsRegion() {
        return Optional.ofNullable(this.awsRegion);
    }
    /**
     * @return The ARN of the AWS Role to assume.Used during STS AssumeRole
     * 
     */
    public Optional<String> awsRoleArn() {
        return Optional.ofNullable(this.awsRoleArn);
    }
    /**
     * @return Specifies the name to attach to the AWS role session. Used during STS AssumeRole
     * 
     */
    public Optional<String> awsRoleSessionName() {
        return Optional.ofNullable(this.awsRoleSessionName);
    }
    /**
     * @return The AWS secret access key.
     * 
     */
    public Optional<String> awsSecretAccessKey() {
        return Optional.ofNullable(this.awsSecretAccessKey);
    }
    /**
     * @return The AWS session token.
     * 
     */
    public Optional<String> awsSessionToken() {
        return Optional.ofNullable(this.awsSessionToken);
    }
    /**
     * @return Path to the AWS shared credentials file.
     * 
     */
    public Optional<String> awsSharedCredentialsFile() {
        return Optional.ofNullable(this.awsSharedCredentialsFile);
    }
    /**
     * @return The STS endpoint URL.
     * 
     */
    public Optional<String> awsStsEndpoint() {
        return Optional.ofNullable(this.awsStsEndpoint);
    }
    /**
     * @return Path to the file containing an OAuth 2.0 access token or OpenID Connect ID token.
     * 
     */
    public Optional<String> awsWebIdentityTokenFile() {
        return Optional.ofNullable(this.awsWebIdentityTokenFile);
    }
    /**
     * @return The Vault header value to include in the STS signing request.
     * 
     */
    public Optional<String> headerValue() {
        return Optional.ofNullable(this.headerValue);
    }
    /**
     * @return The path where the authentication engine is mounted.
     * 
     */
    public Optional<String> mount() {
        return Optional.ofNullable(this.mount);
    }
    /**
     * @return The authentication engine&#39;s namespace. Conflicts with use_root_namespace
     * 
     */
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return The Vault role to use when logging into Vault.
     * 
     */
    public String role() {
        return this.role;
    }
    /**
     * @return Authenticate to the root Vault namespace. Conflicts with namespace
     * 
     */
    public Optional<Boolean> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(AuthLoginAws defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private @Nullable String awsAccessKeyId;
        private @Nullable String awsIamEndpoint;
        private @Nullable String awsProfile;
        private @Nullable String awsRegion;
        private @Nullable String awsRoleArn;
        private @Nullable String awsRoleSessionName;
        private @Nullable String awsSecretAccessKey;
        private @Nullable String awsSessionToken;
        private @Nullable String awsSharedCredentialsFile;
        private @Nullable String awsStsEndpoint;
        private @Nullable String awsWebIdentityTokenFile;
        private @Nullable String headerValue;
        private @Nullable String mount;
        private @Nullable String namespace;
        private String role;
        private @Nullable Boolean useRootNamespace;
        public Builder() {}
        public Builder(AuthLoginAws defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.awsAccessKeyId = defaults.awsAccessKeyId;
    	      this.awsIamEndpoint = defaults.awsIamEndpoint;
    	      this.awsProfile = defaults.awsProfile;
    	      this.awsRegion = defaults.awsRegion;
    	      this.awsRoleArn = defaults.awsRoleArn;
    	      this.awsRoleSessionName = defaults.awsRoleSessionName;
    	      this.awsSecretAccessKey = defaults.awsSecretAccessKey;
    	      this.awsSessionToken = defaults.awsSessionToken;
    	      this.awsSharedCredentialsFile = defaults.awsSharedCredentialsFile;
    	      this.awsStsEndpoint = defaults.awsStsEndpoint;
    	      this.awsWebIdentityTokenFile = defaults.awsWebIdentityTokenFile;
    	      this.headerValue = defaults.headerValue;
    	      this.mount = defaults.mount;
    	      this.namespace = defaults.namespace;
    	      this.role = defaults.role;
    	      this.useRootNamespace = defaults.useRootNamespace;
        }

        @CustomType.Setter
        public Builder awsAccessKeyId(@Nullable String awsAccessKeyId) {

            this.awsAccessKeyId = awsAccessKeyId;
            return this;
        }
        @CustomType.Setter
        public Builder awsIamEndpoint(@Nullable String awsIamEndpoint) {

            this.awsIamEndpoint = awsIamEndpoint;
            return this;
        }
        @CustomType.Setter
        public Builder awsProfile(@Nullable String awsProfile) {

            this.awsProfile = awsProfile;
            return this;
        }
        @CustomType.Setter
        public Builder awsRegion(@Nullable String awsRegion) {

            this.awsRegion = awsRegion;
            return this;
        }
        @CustomType.Setter
        public Builder awsRoleArn(@Nullable String awsRoleArn) {

            this.awsRoleArn = awsRoleArn;
            return this;
        }
        @CustomType.Setter
        public Builder awsRoleSessionName(@Nullable String awsRoleSessionName) {

            this.awsRoleSessionName = awsRoleSessionName;
            return this;
        }
        @CustomType.Setter
        public Builder awsSecretAccessKey(@Nullable String awsSecretAccessKey) {

            this.awsSecretAccessKey = awsSecretAccessKey;
            return this;
        }
        @CustomType.Setter
        public Builder awsSessionToken(@Nullable String awsSessionToken) {

            this.awsSessionToken = awsSessionToken;
            return this;
        }
        @CustomType.Setter
        public Builder awsSharedCredentialsFile(@Nullable String awsSharedCredentialsFile) {

            this.awsSharedCredentialsFile = awsSharedCredentialsFile;
            return this;
        }
        @CustomType.Setter
        public Builder awsStsEndpoint(@Nullable String awsStsEndpoint) {

            this.awsStsEndpoint = awsStsEndpoint;
            return this;
        }
        @CustomType.Setter
        public Builder awsWebIdentityTokenFile(@Nullable String awsWebIdentityTokenFile) {

            this.awsWebIdentityTokenFile = awsWebIdentityTokenFile;
            return this;
        }
        @CustomType.Setter
        public Builder headerValue(@Nullable String headerValue) {

            this.headerValue = headerValue;
            return this;
        }
        @CustomType.Setter
        public Builder mount(@Nullable String mount) {

            this.mount = mount;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder role(String role) {
            if (role == null) {
              throw new MissingRequiredPropertyException("AuthLoginAws", "role");
            }
            this.role = role;
            return this;
        }
        @CustomType.Setter
        public Builder useRootNamespace(@Nullable Boolean useRootNamespace) {

            this.useRootNamespace = useRootNamespace;
            return this;
        }
        public AuthLoginAws build() {
            final var _resultValue = new AuthLoginAws();
            _resultValue.awsAccessKeyId = awsAccessKeyId;
            _resultValue.awsIamEndpoint = awsIamEndpoint;
            _resultValue.awsProfile = awsProfile;
            _resultValue.awsRegion = awsRegion;
            _resultValue.awsRoleArn = awsRoleArn;
            _resultValue.awsRoleSessionName = awsRoleSessionName;
            _resultValue.awsSecretAccessKey = awsSecretAccessKey;
            _resultValue.awsSessionToken = awsSessionToken;
            _resultValue.awsSharedCredentialsFile = awsSharedCredentialsFile;
            _resultValue.awsStsEndpoint = awsStsEndpoint;
            _resultValue.awsWebIdentityTokenFile = awsWebIdentityTokenFile;
            _resultValue.headerValue = headerValue;
            _resultValue.mount = mount;
            _resultValue.namespace = namespace;
            _resultValue.role = role;
            _resultValue.useRootNamespace = useRootNamespace;
            return _resultValue;
        }
    }
}
