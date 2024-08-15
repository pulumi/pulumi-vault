// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.secrets;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class SyncGcpDestinationArgs extends com.pulumi.resources.ResourceArgs {

    public static final SyncGcpDestinationArgs Empty = new SyncGcpDestinationArgs();

    /**
     * JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     * 
     */
    @Import(name="credentials")
    private @Nullable Output<String> credentials;

    /**
     * @return JSON-encoded credentials to use to connect to GCP.
     * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
     * variable.
     * 
     */
    public Optional<Output<String>> credentials() {
        return Optional.ofNullable(this.credentials);
    }

    /**
     * Custom tags to set on the secret managed at the destination.
     * 
     */
    @Import(name="customTags")
    private @Nullable Output<Map<String,String>> customTags;

    /**
     * @return Custom tags to set on the secret managed at the destination.
     * 
     */
    public Optional<Output<Map<String,String>>> customTags() {
        return Optional.ofNullable(this.customTags);
    }

    /**
     * Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     * 
     */
    @Import(name="granularity")
    private @Nullable Output<String> granularity;

    /**
     * @return Determines what level of information is synced as a distinct resource
     * at the destination. Supports `secret-path` and `secret-key`.
     * 
     */
    public Optional<Output<String>> granularity() {
        return Optional.ofNullable(this.granularity);
    }

    /**
     * Unique name of the GCP destination.
     * 
     */
    @Import(name="name")
    private @Nullable Output<String> name;

    /**
     * @return Unique name of the GCP destination.
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
     * The target project to manage secrets in. If set,
     * overrides the project ID derived from the service account JSON credentials or application
     * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
     * to perform Secret Manager actions in the target project.
     * 
     */
    @Import(name="projectId")
    private @Nullable Output<String> projectId;

    /**
     * @return The target project to manage secrets in. If set,
     * overrides the project ID derived from the service account JSON credentials or application
     * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
     * to perform Secret Manager actions in the target project.
     * 
     */
    public Optional<Output<String>> projectId() {
        return Optional.ofNullable(this.projectId);
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

    private SyncGcpDestinationArgs() {}

    private SyncGcpDestinationArgs(SyncGcpDestinationArgs $) {
        this.credentials = $.credentials;
        this.customTags = $.customTags;
        this.granularity = $.granularity;
        this.name = $.name;
        this.namespace = $.namespace;
        this.projectId = $.projectId;
        this.secretNameTemplate = $.secretNameTemplate;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(SyncGcpDestinationArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private SyncGcpDestinationArgs $;

        public Builder() {
            $ = new SyncGcpDestinationArgs();
        }

        public Builder(SyncGcpDestinationArgs defaults) {
            $ = new SyncGcpDestinationArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param credentials JSON-encoded credentials to use to connect to GCP.
         * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder credentials(@Nullable Output<String> credentials) {
            $.credentials = credentials;
            return this;
        }

        /**
         * @param credentials JSON-encoded credentials to use to connect to GCP.
         * Can be omitted and directly provided to Vault using the `GOOGLE_APPLICATION_CREDENTIALS` environment
         * variable.
         * 
         * @return builder
         * 
         */
        public Builder credentials(String credentials) {
            return credentials(Output.of(credentials));
        }

        /**
         * @param customTags Custom tags to set on the secret managed at the destination.
         * 
         * @return builder
         * 
         */
        public Builder customTags(@Nullable Output<Map<String,String>> customTags) {
            $.customTags = customTags;
            return this;
        }

        /**
         * @param customTags Custom tags to set on the secret managed at the destination.
         * 
         * @return builder
         * 
         */
        public Builder customTags(Map<String,String> customTags) {
            return customTags(Output.of(customTags));
        }

        /**
         * @param granularity Determines what level of information is synced as a distinct resource
         * at the destination. Supports `secret-path` and `secret-key`.
         * 
         * @return builder
         * 
         */
        public Builder granularity(@Nullable Output<String> granularity) {
            $.granularity = granularity;
            return this;
        }

        /**
         * @param granularity Determines what level of information is synced as a distinct resource
         * at the destination. Supports `secret-path` and `secret-key`.
         * 
         * @return builder
         * 
         */
        public Builder granularity(String granularity) {
            return granularity(Output.of(granularity));
        }

        /**
         * @param name Unique name of the GCP destination.
         * 
         * @return builder
         * 
         */
        public Builder name(@Nullable Output<String> name) {
            $.name = name;
            return this;
        }

        /**
         * @param name Unique name of the GCP destination.
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
         * @param projectId The target project to manage secrets in. If set,
         * overrides the project ID derived from the service account JSON credentials or application
         * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
         * to perform Secret Manager actions in the target project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(@Nullable Output<String> projectId) {
            $.projectId = projectId;
            return this;
        }

        /**
         * @param projectId The target project to manage secrets in. If set,
         * overrides the project ID derived from the service account JSON credentials or application
         * default credentials. The service account must be [authorized](https://cloud.google.com/iam/docs/service-account-overview#locations)
         * to perform Secret Manager actions in the target project.
         * 
         * @return builder
         * 
         */
        public Builder projectId(String projectId) {
            return projectId(Output.of(projectId));
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

        public SyncGcpDestinationArgs build() {
            return $;
        }
    }

}
