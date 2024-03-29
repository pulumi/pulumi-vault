// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetStaticCredentialsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetStaticCredentialsPlainArgs Empty = new GetStaticCredentialsPlainArgs();

    @Import(name="mount", required=true)
    private String mount;

    public String mount() {
        return this.mount;
    }

    @Import(name="namespace")
    private @Nullable String namespace;

    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }

    @Import(name="roleName", required=true)
    private String roleName;

    public String roleName() {
        return this.roleName;
    }

    private GetStaticCredentialsPlainArgs() {}

    private GetStaticCredentialsPlainArgs(GetStaticCredentialsPlainArgs $) {
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.roleName = $.roleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetStaticCredentialsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetStaticCredentialsPlainArgs $;

        public Builder() {
            $ = new GetStaticCredentialsPlainArgs();
        }

        public Builder(GetStaticCredentialsPlainArgs defaults) {
            $ = new GetStaticCredentialsPlainArgs(Objects.requireNonNull(defaults));
        }

        public Builder mount(String mount) {
            $.mount = mount;
            return this;
        }

        public Builder namespace(@Nullable String namespace) {
            $.namespace = namespace;
            return this;
        }

        public Builder roleName(String roleName) {
            $.roleName = roleName;
            return this;
        }

        public GetStaticCredentialsPlainArgs build() {
            if ($.mount == null) {
                throw new MissingRequiredPropertyException("GetStaticCredentialsPlainArgs", "mount");
            }
            if ($.roleName == null) {
                throw new MissingRequiredPropertyException("GetStaticCredentialsPlainArgs", "roleName");
            }
            return $;
        }
    }

}
