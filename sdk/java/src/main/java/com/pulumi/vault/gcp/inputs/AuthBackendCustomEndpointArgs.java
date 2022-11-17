// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.gcp.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class AuthBackendCustomEndpointArgs extends com.pulumi.resources.ResourceArgs {

    public static final AuthBackendCustomEndpointArgs Empty = new AuthBackendCustomEndpointArgs();

    /**
     * Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
     * 
     */
    @Import(name="api")
    private @Nullable Output<String> api;

    /**
     * @return Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
     * 
     */
    public Optional<Output<String>> api() {
        return Optional.ofNullable(this.api);
    }

    /**
     * Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
     * 
     */
    @Import(name="compute")
    private @Nullable Output<String> compute;

    /**
     * @return Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
     * 
     */
    public Optional<Output<String>> compute() {
        return Optional.ofNullable(this.compute);
    }

    /**
     * Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
     * 
     */
    @Import(name="crm")
    private @Nullable Output<String> crm;

    /**
     * @return Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
     * 
     */
    public Optional<Output<String>> crm() {
        return Optional.ofNullable(this.crm);
    }

    /**
     * Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
     * 
     */
    @Import(name="iam")
    private @Nullable Output<String> iam;

    /**
     * @return Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
     * 
     */
    public Optional<Output<String>> iam() {
        return Optional.ofNullable(this.iam);
    }

    private AuthBackendCustomEndpointArgs() {}

    private AuthBackendCustomEndpointArgs(AuthBackendCustomEndpointArgs $) {
        this.api = $.api;
        this.compute = $.compute;
        this.crm = $.crm;
        this.iam = $.iam;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(AuthBackendCustomEndpointArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private AuthBackendCustomEndpointArgs $;

        public Builder() {
            $ = new AuthBackendCustomEndpointArgs();
        }

        public Builder(AuthBackendCustomEndpointArgs defaults) {
            $ = new AuthBackendCustomEndpointArgs(Objects.requireNonNull(defaults));
        }

        /**
         * @param api Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder api(@Nullable Output<String> api) {
            $.api = api;
            return this;
        }

        /**
         * @param api Replaces the service endpoint used in API requests to `https://www.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder api(String api) {
            return api(Output.of(api));
        }

        /**
         * @param compute Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder compute(@Nullable Output<String> compute) {
            $.compute = compute;
            return this;
        }

        /**
         * @param compute Replaces the service endpoint used in API requests to `https://compute.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder compute(String compute) {
            return compute(Output.of(compute));
        }

        /**
         * @param crm Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder crm(@Nullable Output<String> crm) {
            $.crm = crm;
            return this;
        }

        /**
         * @param crm Replaces the service endpoint used in API requests to `https://cloudresourcemanager.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder crm(String crm) {
            return crm(Output.of(crm));
        }

        /**
         * @param iam Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder iam(@Nullable Output<String> iam) {
            $.iam = iam;
            return this;
        }

        /**
         * @param iam Replaces the service endpoint used in API requests to `https://iam.googleapis.com`.
         * 
         * @return builder
         * 
         */
        public Builder iam(String iam) {
            return iam(Output.of(iam));
        }

        public AuthBackendCustomEndpointArgs build() {
            return $;
        }
    }

}