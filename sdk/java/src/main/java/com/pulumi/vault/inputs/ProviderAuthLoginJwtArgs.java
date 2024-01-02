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


public final class ProviderAuthLoginJwtArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginJwtArgs Empty = new ProviderAuthLoginJwtArgs();

    @Import(name="jwt", required=true)
    private Output<String> jwt;

    public Output<String> jwt() {
        return this.jwt;
    }

    @Import(name="mount")
    private @Nullable Output<String> mount;

    public Optional<Output<String>> mount() {
        return Optional.ofNullable(this.mount);
    }

    @Import(name="namespace")
    private @Nullable Output<String> namespace;

    public Optional<Output<String>> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    @Import(name="role", required=true)
    private Output<String> role;

    public Output<String> role() {
        return this.role;
    }

    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    private ProviderAuthLoginJwtArgs() {}

    private ProviderAuthLoginJwtArgs(ProviderAuthLoginJwtArgs $) {
        this.jwt = $.jwt;
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.role = $.role;
        this.useRootNamespace = $.useRootNamespace;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginJwtArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginJwtArgs $;

        public Builder() {
            $ = new ProviderAuthLoginJwtArgs();
        }

        public Builder(ProviderAuthLoginJwtArgs defaults) {
            $ = new ProviderAuthLoginJwtArgs(Objects.requireNonNull(defaults));
        }

        public Builder jwt(Output<String> jwt) {
            $.jwt = jwt;
            return this;
        }

        public Builder jwt(String jwt) {
            return jwt(Output.of(jwt));
        }

        public Builder mount(@Nullable Output<String> mount) {
            $.mount = mount;
            return this;
        }

        public Builder mount(String mount) {
            return mount(Output.of(mount));
        }

        public Builder namespace(@Nullable Output<String> namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder namespace(String namespace) {
            return namespace(Output.of(namespace));
        }

        public Builder role(Output<String> role) {
            $.role = role;
            return this;
        }

        public Builder role(String role) {
            return role(Output.of(role));
        }

        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        public ProviderAuthLoginJwtArgs build() {
            if ($.jwt == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginJwtArgs", "jwt");
            }
            if ($.role == null) {
                throw new MissingRequiredPropertyException("ProviderAuthLoginJwtArgs", "role");
            }
            return $;
        }
    }

}
