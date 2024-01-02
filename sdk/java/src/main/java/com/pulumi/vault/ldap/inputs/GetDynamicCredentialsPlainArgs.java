// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.ldap.inputs;

import com.pulumi.core.annotations.Import;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import java.lang.String;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;


public final class GetDynamicCredentialsPlainArgs extends com.pulumi.resources.InvokeArgs {

    public static final GetDynamicCredentialsPlainArgs Empty = new GetDynamicCredentialsPlainArgs();

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

    private GetDynamicCredentialsPlainArgs() {}

    private GetDynamicCredentialsPlainArgs(GetDynamicCredentialsPlainArgs $) {
        this.mount = $.mount;
        this.namespace = $.namespace;
        this.roleName = $.roleName;
    }

    public static Builder builder() {
        return new Builder();
    }
    public static Builder builder(GetDynamicCredentialsPlainArgs defaults) {
        return new Builder(defaults);
    }

    public static final class Builder {
        private GetDynamicCredentialsPlainArgs $;

        public Builder() {
            $ = new GetDynamicCredentialsPlainArgs();
        }

        public Builder(GetDynamicCredentialsPlainArgs defaults) {
            $ = new GetDynamicCredentialsPlainArgs(Objects.requireNonNull(defaults));
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

        public GetDynamicCredentialsPlainArgs build() {
            if ($.mount == null) {
                throw new MissingRequiredPropertyException("GetDynamicCredentialsPlainArgs", "mount");
            }
            if ($.roleName == null) {
                throw new MissingRequiredPropertyException("GetDynamicCredentialsPlainArgs", "roleName");
            }
            return $;
        }
    }

}
