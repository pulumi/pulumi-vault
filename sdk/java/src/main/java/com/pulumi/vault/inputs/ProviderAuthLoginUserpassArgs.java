// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.inputs;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Import;
import java.lang.Boolean;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class ProviderAuthLoginUserpassArgs extends com.pulumi.resources.ResourceArgs {

    public static final ProviderAuthLoginUserpassArgs Empty = new ProviderAuthLoginUserpassArgs();

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

    @Import(name="password")
    private @Nullable Output<String> password;

    public Optional<Output<String>> password() {
        return Optional.ofNullable(this.password);
    }

    @Import(name="passwordFile")
    private @Nullable Output<String> passwordFile;

    public Optional<Output<String>> passwordFile() {
        return Optional.ofNullable(this.passwordFile);
    }

    @Import(name="useRootNamespace")
    private @Nullable Output<Boolean> useRootNamespace;

    public Optional<Output<Boolean>> useRootNamespace() {
        return Optional.ofNullable(this.useRootNamespace);
    }

    @Import(name="username", required=true)
    private Output<String> username;

    public Output<String> username() {
        return this.username;
    }

    private ProviderAuthLoginUserpassArgs() {}

    private ProviderAuthLoginUserpassArgs(ProviderAuthLoginUserpassArgs $) {
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.password = $.password;
        this.passwordFile = $.passwordFile;
        this.useRootNamespace = $.useRootNamespace;
        this.username = $.username;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(ProviderAuthLoginUserpassArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private ProviderAuthLoginUserpassArgs $;

        public Builder() {
            $ = new ProviderAuthLoginUserpassArgs();
        }

        public Builder(ProviderAuthLoginUserpassArgs defaults) {
            $ = new ProviderAuthLoginUserpassArgs(Objects.requireNonNull(defaults));
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

        public Builder password(@Nullable Output<String> password) {
            $.password = password;
            return this;
        }

        public Builder password(String password) {
            return password(Output.of(password));
        }

        public Builder passwordFile(@Nullable Output<String> passwordFile) {
            $.passwordFile = passwordFile;
            return this;
        }

        public Builder passwordFile(String passwordFile) {
            return passwordFile(Output.of(passwordFile));
        }

        public Builder useRootNamespace(@Nullable Output<Boolean> useRootNamespace) {
            $.useRootNamespace = useRootNamespace;
            return this;
        }

        public Builder useRootNamespace(Boolean useRootNamespace) {
            return useRootNamespace(Output.of(useRootNamespace));
        }

        public Builder username(Output<String> username) {
            $.username = username;
            return this;
        }

        public Builder username(String username) {
            return username(Output.of(username));
        }

        public ProviderAuthLoginUserpassArgs build() {
            $.username = Objects.requireNonNull($.username, "expected parameter 'username' to be non-null");
            return $;
        }
    }

}
