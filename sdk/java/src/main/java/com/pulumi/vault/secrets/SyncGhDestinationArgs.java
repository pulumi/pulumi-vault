// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Integer;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SyncGhDestinationArgs extends com.pulumi.resources.ResourceArgs {

    public static final SyncGhDestinationArgs Empty = new SyncGhDestinationArgs();

    /**
     * Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     * 
     */
    @Import(name="accessToken")
    private @Nullable Output<String> accessToken;

    /**
     * @return Fine-grained or personal access token.
     * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
     * variable.
     * 
     */
    public Optional<Output<String>> accessToken() {
        return Optional.ofNullable(this.accessToken);
    }

    /**
     * The user-defined name of the GitHub App configuration. This is a reference to the name used\
     * on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
     * Takes precedence over the `access_token` field.
     * 
     */
    @Import(name="appName")
    private @Nullable Output<String> appName;

    /**
     * @return The user-defined name of the GitHub App configuration. This is a reference to the name used\
     * on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
     * Takes precedence over the `access_token` field.
     * 
     */
    public Optional<Output<String>> appName() {
        return Optional.ofNullable(this.appName);
    }

    /**
     * The ID of the installation generated by GitHub when the app referenced by the `app_name`
     * was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
     * 
     */
    @Import(name="installationId")
    private @Nullable Output<Integer> installationId;

    /**
     * @return The ID of the installation generated by GitHub when the app referenced by the `app_name`
     * was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
     * 
     */
    public Optional<Output<Integer>> installationId() {
        return Optional.ofNullable(this.installationId);
    }

    /**
     * Unique name of the GitHub destination.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name of the GitHub destination.
     * 
     */
    public Optional<Output<String>> name() {
        return Optional.ofNullable(this.name);
    }

    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * 
     */
    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    /**
     * Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     * 
     */
    @Import(name="repositoryName")
    private @Nullable Output<String> repositoryName;

    /**
     * @return Name of the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
     * variable.
     * 
     */
    public Optional<Output<String>> repositoryName() {
        return Optional.ofNullable(this.repositoryName);
    }

    /**
     * GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     * 
     */
    @Import(name="repositoryOwner")
    private @Nullable Output<String> repositoryOwner;

    /**
     * @return GitHub organization or username that owns the repository.
     * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
     * variable.
     * 
     */
    public Optional<Output<String>> repositoryOwner() {
        return Optional.ofNullable(this.repositoryOwner);
    }

    /**
     * Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    @Import(name="secretNameTemplate")
    private @Nullable Output<String> secretNameTemplate;

    /**
     * @return Template describing how to generate external secret names.
     * Supports a subset of the Go Template syntax.
     * 
     */
    public Optional<Output<String>> secretNameTemplate() {
        return Optional.ofNullable(this.secretNameTemplate);
    }

    private SyncGhDestinationArgs() {}

    private SyncGhDestinationArgs(SyncGhDestinationArgs $) {
        this.accessToken = $.accessToken;
        this.appName = $.appName;
        this.installationId = $.installationId;
        this.name = $.name;
        this.namespace = $.namespace;
        this.repositoryName = $.repositoryName;
        this.repositoryOwner = $.repositoryOwner;
        this.secretNameTemplate = $.secretNameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SyncGhDestinationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SyncGhDestinationArgs $;

        public Builder() {
            $ = new SyncGhDestinationArgs();
        }

        public Builder(SyncGhDestinationArgs defaults) {
            $ = new SyncGhDestinationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param accessToken Fine-grained or personal access token.
         * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder accessToken(@Nullable Output<String> accessToken) {
            $.accessToken = accessToken;
            return this;
        }

        /**
         * @param accessToken Fine-grained or personal access token.
         * Can be omitted and directly provided to Vault using the `GITHUB_ACCESS_TOKEN` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder accessToken(String accessToken) {
            return accessToken(Output.of(accessToken));
        }

        /**
         * @param appName The user-defined name of the GitHub App configuration. This is a reference to the name used\
         * on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
         * Takes precedence over the `access_token` field.
         * 
         * @return builder
         * 
         */
        public Builder appName(@Nullable Output<String> appName) {
            $.appName = appName;
            return this;
        }

        /**
         * @param appName The user-defined name of the GitHub App configuration. This is a reference to the name used\
         * on the new endpoint when configuring the GitHub app on the Vault Server. Can be modified.
         * Takes precedence over the `access_token` field.
         * 
         * @return builder
         * 
         */
        public Builder appName(String appName) {
            return appName(Output.of(appName));
        }

        /**
         * @param installationId The ID of the installation generated by GitHub when the app referenced by the `app_name`
         * was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
         * 
         * @return builder
         * 
         */
        public Builder installationId(@Nullable Output<Integer> installationId) {
            $.installationId = installationId;
            return this;
        }

        /**
         * @param installationId The ID of the installation generated by GitHub when the app referenced by the `app_name`
         * was installed in the user’s GitHub account. Can be modified. Necessary if the `app_name` field is also provided.
         * 
         * @return builder
         * 
         */
        public Builder installationId(Integer installationId) {
            return installationId(Output.of(installationId));
        }

        /**
         * @param name Unique name of the GitHub destination.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name of the GitHub destination.
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
         * 
         * @return builder
         * 
         */
        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        /**
         * @param repositoryName Name of the repository.
         * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder repositoryName(@Nullable Output<String> repositoryName) {
            $.repositoryName = repositoryName;
            return this;
        }

        /**
         * @param repositoryName Name of the repository.
         * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_NAME` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder repositoryName(String repositoryName) {
            return repositoryName(Output.of(repositoryName));
        }

        /**
         * @param repositoryOwner GitHub organization or username that owns the repository.
         * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder repositoryOwner(@Nullable Output<String> repositoryOwner) {
            $.repositoryOwner = repositoryOwner;
            return this;
        }

        /**
         * @param repositoryOwner GitHub organization or username that owns the repository.
         * Can be omitted and directly provided to Vault using the `GITHUB_REPOSITORY_OWNER` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder repositoryOwner(String repositoryOwner) {
            return repositoryOwner(Output.of(repositoryOwner));
        }

        /**
         * @param secretNameTemplate Template describing how to generate external secret names.
         * Supports a subset of the Go Template syntax.
         * 
         * @return builder
         * 
         */
        public Builder secretNameTemplate(@Nullable Output<String> secretNameTemplate) {
            $.secretNameTemplate = secretNameTemplate;
            return this;
        }

        /**
         * @param secretNameTemplate Template describing how to generate external secret names.
         * Supports a subset of the Go Template syntax.
         * 
         * @return builder
         * 
         */
        public Builder secretNameTemplate(String secretNameTemplate) {
            return secretNameTemplate(Output.of(secretNameTemplate));
        }

        public SyncGhDestinationArgs build() {
            return $;
        }
    }

}
