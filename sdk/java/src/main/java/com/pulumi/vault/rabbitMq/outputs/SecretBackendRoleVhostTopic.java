// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.rabbitMq.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.rabbitMq.outputs.SecretBackendRoleVhostTopicVhost;
import java.lang.String;
import java.util.List;
import java.util.Objects;
import javax.annotation.Nullable;

@CustomType
public final class SecretBackendRoleVhostTopic {
    /**
     * @return The vhost to set permissions for.
     * 
     */
    private String host;
    /**
     * @return Specifies a map of virtual hosts to permissions.
     * 
     */
    private @Nullable List<SecretBackendRoleVhostTopicVhost> vhosts;

    private SecretBackendRoleVhostTopic() {}
    /**
     * @return The vhost to set permissions for.
     * 
     */
    public String host() {
        return this.host;
    }
    /**
     * @return Specifies a map of virtual hosts to permissions.
     * 
     */
    public List<SecretBackendRoleVhostTopicVhost> vhosts() {
        return this.vhosts == null ? List.of() : this.vhosts;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(SecretBackendRoleVhostTopic defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String host;
        private @Nullable List<SecretBackendRoleVhostTopicVhost> vhosts;
        public Builder() {}
        public Builder(SecretBackendRoleVhostTopic defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.host = defaults.host;
    	      this.vhosts = defaults.vhosts;
        }

        @CustomType.Setter
        public Builder host(String host) {
            if (host == null) {
              throw new MissingRequiredPropertyException("SecretBackendRoleVhostTopic", "host");
            }
            this.host = host;
            return this;
        }
        @CustomType.Setter
        public Builder vhosts(@Nullable List<SecretBackendRoleVhostTopicVhost> vhosts) {

            this.vhosts = vhosts;
            return this;
        }
        public Builder vhosts(SecretBackendRoleVhostTopicVhost... vhosts) {
            return vhosts(List.of(vhosts));
        }
        public SecretBackendRoleVhostTopic build() {
            final var _resultValue = new SecretBackendRoleVhostTopic();
            _resultValue.host = host;
            _resultValue.vhosts = vhosts;
            return _resultValue;
        }
    }
}
