// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault.identity.outputs;

import com.pulumi.core.annotations.CustomType;
import com.pulumi.exceptions.MissingRequiredPropertyException;
import com.pulumi.vault.identity.outputs.GetEntityAlias;
import java.lang.Boolean;
import java.lang.String;
import java.util.List;
import java.util.Map;
import java.util.Objects;
import java.util.Optional;
import javax.annotation.Nullable;

@CustomType
public final class GetEntityResult {
    private String aliasId;
    private String aliasMountAccessor;
    private String aliasName;
    /**
     * @return A list of entity alias. Structure is documented below.
     * 
     */
    private List<GetEntityAlias> aliases;
    /**
     * @return Creation time of the Alias
     * 
     */
    private String creationTime;
    /**
     * @return A string containing the full data payload retrieved from
     * Vault, serialized in JSON format.
     * 
     */
    private String dataJson;
    /**
     * @return List of Group IDs of which the entity is directly a member of
     * 
     */
    private List<String> directGroupIds;
    /**
     * @return Whether the entity is disabled
     * 
     */
    private Boolean disabled;
    private String entityId;
    private String entityName;
    /**
     * @return List of all Group IDs of which the entity is a member of
     * 
     */
    private List<String> groupIds;
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    private String id;
    /**
     * @return List of all Group IDs of which the entity is a member of transitively
     * 
     */
    private List<String> inheritedGroupIds;
    /**
     * @return Last update time of the alias
     * 
     */
    private String lastUpdateTime;
    /**
     * @return Other entity IDs which is merged with this entity
     * 
     */
    private List<String> mergedEntityIds;
    /**
     * @return Arbitrary metadata
     * 
     */
    private Map<String,String> metadata;
    private @Nullable String namespace;
    /**
     * @return Namespace of which the entity is part of
     * 
     */
    private String namespaceId;
    /**
     * @return List of policies attached to the entity
     * 
     */
    private List<String> policies;

    private GetEntityResult() {}
    public String aliasId() {
        return this.aliasId;
    }
    public String aliasMountAccessor() {
        return this.aliasMountAccessor;
    }
    public String aliasName() {
        return this.aliasName;
    }
    /**
     * @return A list of entity alias. Structure is documented below.
     * 
     */
    public List<GetEntityAlias> aliases() {
        return this.aliases;
    }
    /**
     * @return Creation time of the Alias
     * 
     */
    public String creationTime() {
        return this.creationTime;
    }
    /**
     * @return A string containing the full data payload retrieved from
     * Vault, serialized in JSON format.
     * 
     */
    public String dataJson() {
        return this.dataJson;
    }
    /**
     * @return List of Group IDs of which the entity is directly a member of
     * 
     */
    public List<String> directGroupIds() {
        return this.directGroupIds;
    }
    /**
     * @return Whether the entity is disabled
     * 
     */
    public Boolean disabled() {
        return this.disabled;
    }
    public String entityId() {
        return this.entityId;
    }
    public String entityName() {
        return this.entityName;
    }
    /**
     * @return List of all Group IDs of which the entity is a member of
     * 
     */
    public List<String> groupIds() {
        return this.groupIds;
    }
    /**
     * @return The provider-assigned unique ID for this managed resource.
     * 
     */
    public String id() {
        return this.id;
    }
    /**
     * @return List of all Group IDs of which the entity is a member of transitively
     * 
     */
    public List<String> inheritedGroupIds() {
        return this.inheritedGroupIds;
    }
    /**
     * @return Last update time of the alias
     * 
     */
    public String lastUpdateTime() {
        return this.lastUpdateTime;
    }
    /**
     * @return Other entity IDs which is merged with this entity
     * 
     */
    public List<String> mergedEntityIds() {
        return this.mergedEntityIds;
    }
    /**
     * @return Arbitrary metadata
     * 
     */
    public Map<String,String> metadata() {
        return this.metadata;
    }
    public Optional<String> namespace() {
        return Optional.ofNullable(this.namespace);
    }
    /**
     * @return Namespace of which the entity is part of
     * 
     */
    public String namespaceId() {
        return this.namespaceId;
    }
    /**
     * @return List of policies attached to the entity
     * 
     */
    public List<String> policies() {
        return this.policies;
    }

    public static Builder builder() {
        return new Builder();
    }

    public static Builder builder(GetEntityResult defaults) {
        return new Builder(defaults);
    }
    @CustomType.Builder
    public static final class Builder {
        private String aliasId;
        private String aliasMountAccessor;
        private String aliasName;
        private List<GetEntityAlias> aliases;
        private String creationTime;
        private String dataJson;
        private List<String> directGroupIds;
        private Boolean disabled;
        private String entityId;
        private String entityName;
        private List<String> groupIds;
        private String id;
        private List<String> inheritedGroupIds;
        private String lastUpdateTime;
        private List<String> mergedEntityIds;
        private Map<String,String> metadata;
        private @Nullable String namespace;
        private String namespaceId;
        private List<String> policies;
        public Builder() {}
        public Builder(GetEntityResult defaults) {
    	      Objects.requireNonNull(defaults);
    	      this.aliasId = defaults.aliasId;
    	      this.aliasMountAccessor = defaults.aliasMountAccessor;
    	      this.aliasName = defaults.aliasName;
    	      this.aliases = defaults.aliases;
    	      this.creationTime = defaults.creationTime;
    	      this.dataJson = defaults.dataJson;
    	      this.directGroupIds = defaults.directGroupIds;
    	      this.disabled = defaults.disabled;
    	      this.entityId = defaults.entityId;
    	      this.entityName = defaults.entityName;
    	      this.groupIds = defaults.groupIds;
    	      this.id = defaults.id;
    	      this.inheritedGroupIds = defaults.inheritedGroupIds;
    	      this.lastUpdateTime = defaults.lastUpdateTime;
    	      this.mergedEntityIds = defaults.mergedEntityIds;
    	      this.metadata = defaults.metadata;
    	      this.namespace = defaults.namespace;
    	      this.namespaceId = defaults.namespaceId;
    	      this.policies = defaults.policies;
        }

        @CustomType.Setter
        public Builder aliasId(String aliasId) {
            if (aliasId == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "aliasId");
            }
            this.aliasId = aliasId;
            return this;
        }
        @CustomType.Setter
        public Builder aliasMountAccessor(String aliasMountAccessor) {
            if (aliasMountAccessor == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "aliasMountAccessor");
            }
            this.aliasMountAccessor = aliasMountAccessor;
            return this;
        }
        @CustomType.Setter
        public Builder aliasName(String aliasName) {
            if (aliasName == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "aliasName");
            }
            this.aliasName = aliasName;
            return this;
        }
        @CustomType.Setter
        public Builder aliases(List<GetEntityAlias> aliases) {
            if (aliases == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "aliases");
            }
            this.aliases = aliases;
            return this;
        }
        public Builder aliases(GetEntityAlias... aliases) {
            return aliases(List.of(aliases));
        }
        @CustomType.Setter
        public Builder creationTime(String creationTime) {
            if (creationTime == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "creationTime");
            }
            this.creationTime = creationTime;
            return this;
        }
        @CustomType.Setter
        public Builder dataJson(String dataJson) {
            if (dataJson == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "dataJson");
            }
            this.dataJson = dataJson;
            return this;
        }
        @CustomType.Setter
        public Builder directGroupIds(List<String> directGroupIds) {
            if (directGroupIds == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "directGroupIds");
            }
            this.directGroupIds = directGroupIds;
            return this;
        }
        public Builder directGroupIds(String... directGroupIds) {
            return directGroupIds(List.of(directGroupIds));
        }
        @CustomType.Setter
        public Builder disabled(Boolean disabled) {
            if (disabled == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "disabled");
            }
            this.disabled = disabled;
            return this;
        }
        @CustomType.Setter
        public Builder entityId(String entityId) {
            if (entityId == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "entityId");
            }
            this.entityId = entityId;
            return this;
        }
        @CustomType.Setter
        public Builder entityName(String entityName) {
            if (entityName == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "entityName");
            }
            this.entityName = entityName;
            return this;
        }
        @CustomType.Setter
        public Builder groupIds(List<String> groupIds) {
            if (groupIds == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "groupIds");
            }
            this.groupIds = groupIds;
            return this;
        }
        public Builder groupIds(String... groupIds) {
            return groupIds(List.of(groupIds));
        }
        @CustomType.Setter
        public Builder id(String id) {
            if (id == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "id");
            }
            this.id = id;
            return this;
        }
        @CustomType.Setter
        public Builder inheritedGroupIds(List<String> inheritedGroupIds) {
            if (inheritedGroupIds == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "inheritedGroupIds");
            }
            this.inheritedGroupIds = inheritedGroupIds;
            return this;
        }
        public Builder inheritedGroupIds(String... inheritedGroupIds) {
            return inheritedGroupIds(List.of(inheritedGroupIds));
        }
        @CustomType.Setter
        public Builder lastUpdateTime(String lastUpdateTime) {
            if (lastUpdateTime == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "lastUpdateTime");
            }
            this.lastUpdateTime = lastUpdateTime;
            return this;
        }
        @CustomType.Setter
        public Builder mergedEntityIds(List<String> mergedEntityIds) {
            if (mergedEntityIds == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "mergedEntityIds");
            }
            this.mergedEntityIds = mergedEntityIds;
            return this;
        }
        public Builder mergedEntityIds(String... mergedEntityIds) {
            return mergedEntityIds(List.of(mergedEntityIds));
        }
        @CustomType.Setter
        public Builder metadata(Map<String,String> metadata) {
            if (metadata == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "metadata");
            }
            this.metadata = metadata;
            return this;
        }
        @CustomType.Setter
        public Builder namespace(@Nullable String namespace) {

            this.namespace = namespace;
            return this;
        }
        @CustomType.Setter
        public Builder namespaceId(String namespaceId) {
            if (namespaceId == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "namespaceId");
            }
            this.namespaceId = namespaceId;
            return this;
        }
        @CustomType.Setter
        public Builder policies(List<String> policies) {
            if (policies == null) {
              throw new MissingRequiredPropertyException("GetEntityResult", "policies");
            }
            this.policies = policies;
            return this;
        }
        public Builder policies(String... policies) {
            return policies(List.of(policies));
        }
        public GetEntityResult build() {
            final var _resultValue = new GetEntityResult();
            _resultValue.aliasId = aliasId;
            _resultValue.aliasMountAccessor = aliasMountAccessor;
            _resultValue.aliasName = aliasName;
            _resultValue.aliases = aliases;
            _resultValue.creationTime = creationTime;
            _resultValue.dataJson = dataJson;
            _resultValue.directGroupIds = directGroupIds;
            _resultValue.disabled = disabled;
            _resultValue.entityId = entityId;
            _resultValue.entityName = entityName;
            _resultValue.groupIds = groupIds;
            _resultValue.id = id;
            _resultValue.inheritedGroupIds = inheritedGroupIds;
            _resultValue.lastUpdateTime = lastUpdateTime;
            _resultValue.mergedEntityIds = mergedEntityIds;
            _resultValue.metadata = metadata;
            _resultValue.namespace = namespace;
            _resultValue.namespaceId = namespaceId;
            _resultValue.policies = policies;
            return _resultValue;
        }
    }
}