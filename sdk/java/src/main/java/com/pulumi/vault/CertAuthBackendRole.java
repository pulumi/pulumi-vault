// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.Output;
import com.pulumi.core.annotations.Export;
import com.pulumi.core.annotations.ResourceType;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.CertAuthBackendRoleArgs;
import com.pulumi.vault.Utilities;
import com.pulumi.vault.inputs.CertAuthBackendRoleState;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;
import javax.annotation.Nullable;

/**
 * Provides a resource to create a role in an [Cert auth backend within Vault](https://www.vaultproject.io/docs/auth/cert.html).
 * 
 */
@ResourceType(type="vault:index/certAuthBackendRole:CertAuthBackendRole")
public class CertAuthBackendRole extends com.pulumi.resources.CustomResource {
    /**
     * Allowed the common names for authenticated client certificates
     * 
     */
    @Export(name="allowedCommonNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedCommonNames;

    /**
     * @return Allowed the common names for authenticated client certificates
     * 
     */
    public Output<List<String>> allowedCommonNames() {
        return this.allowedCommonNames;
    }
    /**
     * Allowed alternative dns names for authenticated client certificates
     * 
     */
    @Export(name="allowedDnsSans", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedDnsSans;

    /**
     * @return Allowed alternative dns names for authenticated client certificates
     * 
     */
    public Output<List<String>> allowedDnsSans() {
        return this.allowedDnsSans;
    }
    /**
     * Allowed emails for authenticated client certificates
     * 
     */
    @Export(name="allowedEmailSans", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedEmailSans;

    /**
     * @return Allowed emails for authenticated client certificates
     * 
     */
    public Output<List<String>> allowedEmailSans() {
        return this.allowedEmailSans;
    }
    /**
     * DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
     * 
     */
    @Export(name="allowedNames", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedNames;

    /**
     * @return DEPRECATED: Please use the individual `allowed_X_sans` parameters instead. Allowed subject names for authenticated client certificates
     * 
     */
    public Output<List<String>> allowedNames() {
        return this.allowedNames;
    }
    /**
     * Allowed organization units for authenticated client certificates.
     * 
     */
    @Export(name="allowedOrganizationalUnits", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> allowedOrganizationalUnits;

    /**
     * @return Allowed organization units for authenticated client certificates.
     * 
     */
    public Output<Optional<List<String>>> allowedOrganizationalUnits() {
        return Codegen.optional(this.allowedOrganizationalUnits);
    }
    /**
     * Allowed URIs for authenticated client certificates
     * 
     */
    @Export(name="allowedUriSans", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> allowedUriSans;

    /**
     * @return Allowed URIs for authenticated client certificates
     * 
     */
    public Output<List<String>> allowedUriSans() {
        return this.allowedUriSans;
    }
    /**
     * Path to the mounted Cert auth backend
     * 
     */
    @Export(name="backend", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> backend;

    /**
     * @return Path to the mounted Cert auth backend
     * 
     */
    public Output<Optional<String>> backend() {
        return Codegen.optional(this.backend);
    }
    /**
     * CA certificate used to validate client certificates
     * 
     */
    @Export(name="certificate", refs={String.class}, tree="[0]")
    private Output<String> certificate;

    /**
     * @return CA certificate used to validate client certificates
     * 
     */
    public Output<String> certificate() {
        return this.certificate;
    }
    /**
     * The name to display on tokens issued under this role.
     * 
     */
    @Export(name="displayName", refs={String.class}, tree="[0]")
    private Output<String> displayName;

    /**
     * @return The name to display on tokens issued under this role.
     * 
     */
    public Output<String> displayName() {
        return this.displayName;
    }
    /**
     * Name of the role
     * 
     */
    @Export(name="name", refs={String.class}, tree="[0]")
    private Output<String> name;

    /**
     * @return Name of the role
     * 
     */
    public Output<String> name() {
        return this.name;
    }
    /**
     * The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    @Export(name="namespace", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> namespace;

    /**
     * @return The namespace to provision the resource in.
     * The value should not contain leading or trailing forward slashes.
     * The `namespace` is always relative to the provider&#39;s configured [namespace](https://www.terraform.io/docs/providers/vault/index.html#namespace).
     * *Available only for Vault Enterprise*.
     * 
     */
    public Output<Optional<String>> namespace() {
        return Codegen.optional(this.namespace);
    }
    /**
     * Any additional CA certificates
     * needed to verify OCSP responses. Provided as base64 encoded PEM data.
     * Requires Vault version 1.13+.
     * 
     */
    @Export(name="ocspCaCertificates", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> ocspCaCertificates;

    /**
     * @return Any additional CA certificates
     * needed to verify OCSP responses. Provided as base64 encoded PEM data.
     * Requires Vault version 1.13+.
     * 
     */
    public Output<Optional<String>> ocspCaCertificates() {
        return Codegen.optional(this.ocspCaCertificates);
    }
    /**
     * If enabled, validate certificates&#39;
     * revocation status using OCSP. Requires Vault version 1.13+.
     * 
     */
    @Export(name="ocspEnabled", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ocspEnabled;

    /**
     * @return If enabled, validate certificates&#39;
     * revocation status using OCSP. Requires Vault version 1.13+.
     * 
     */
    public Output<Boolean> ocspEnabled() {
        return this.ocspEnabled;
    }
    /**
     * If true and an OCSP response cannot
     * be fetched or is of an unknown status, the login will proceed as if the
     * certificate has not been revoked.
     * Requires Vault version 1.13+.
     * 
     */
    @Export(name="ocspFailOpen", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ocspFailOpen;

    /**
     * @return If true and an OCSP response cannot
     * be fetched or is of an unknown status, the login will proceed as if the
     * certificate has not been revoked.
     * Requires Vault version 1.13+.
     * 
     */
    public Output<Boolean> ocspFailOpen() {
        return this.ocspFailOpen;
    }
    /**
     * If set to true, rather than
     * accepting the first successful OCSP response, query all servers and consider
     * the certificate valid only if all servers agree.
     * Requires Vault version 1.13+.
     * 
     */
    @Export(name="ocspQueryAllServers", refs={Boolean.class}, tree="[0]")
    private Output<Boolean> ocspQueryAllServers;

    /**
     * @return If set to true, rather than
     * accepting the first successful OCSP response, query all servers and consider
     * the certificate valid only if all servers agree.
     * Requires Vault version 1.13+.
     * 
     */
    public Output<Boolean> ocspQueryAllServers() {
        return this.ocspQueryAllServers;
    }
    /**
     * : A comma-separated list of OCSP
     * server addresses. If unset, the OCSP server is determined from the
     * AuthorityInformationAccess extension on the certificate being inspected.
     * Requires Vault version 1.13+.
     * 
     */
    @Export(name="ocspServersOverrides", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> ocspServersOverrides;

    /**
     * @return : A comma-separated list of OCSP
     * server addresses. If unset, the OCSP server is determined from the
     * AuthorityInformationAccess extension on the certificate being inspected.
     * Requires Vault version 1.13+.
     * 
     */
    public Output<Optional<List<String>>> ocspServersOverrides() {
        return Codegen.optional(this.ocspServersOverrides);
    }
    /**
     * TLS extensions required on
     * client certificates
     * 
     */
    @Export(name="requiredExtensions", refs={List.class,String.class}, tree="[0,1]")
    private Output<List<String>> requiredExtensions;

    /**
     * @return TLS extensions required on
     * client certificates
     * 
     */
    public Output<List<String>> requiredExtensions() {
        return this.requiredExtensions;
    }
    /**
     * List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    @Export(name="tokenBoundCidrs", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenBoundCidrs;

    /**
     * @return List of CIDR blocks; if set, specifies blocks of IP
     * addresses which can authenticate successfully, and ties the resulting token to these blocks
     * as well.
     * 
     */
    public Output<Optional<List<String>>> tokenBoundCidrs() {
        return Codegen.optional(this.tokenBoundCidrs);
    }
    /**
     * If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    @Export(name="tokenExplicitMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenExplicitMaxTtl;

    /**
     * @return If set, will encode an
     * [explicit max TTL](https://www.vaultproject.io/docs/concepts/tokens.html#token-time-to-live-periodic-tokens-and-explicit-max-ttls)
     * onto the token in number of seconds. This is a hard cap even if `token_ttl` and
     * `token_max_ttl` would otherwise allow a renewal.
     * 
     */
    public Output<Optional<Integer>> tokenExplicitMaxTtl() {
        return Codegen.optional(this.tokenExplicitMaxTtl);
    }
    /**
     * The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenMaxTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenMaxTtl;

    /**
     * @return The maximum lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenMaxTtl() {
        return Codegen.optional(this.tokenMaxTtl);
    }
    /**
     * If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    @Export(name="tokenNoDefaultPolicy", refs={Boolean.class}, tree="[0]")
    private Output</* @Nullable */ Boolean> tokenNoDefaultPolicy;

    /**
     * @return If set, the default policy will not be set on
     * generated tokens; otherwise it will be added to the policies set in token_policies.
     * 
     */
    public Output<Optional<Boolean>> tokenNoDefaultPolicy() {
        return Codegen.optional(this.tokenNoDefaultPolicy);
    }
    /**
     * The [maximum number](https://www.vaultproject.io/api-docs/auth/cert#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    @Export(name="tokenNumUses", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenNumUses;

    /**
     * @return The [maximum number](https://www.vaultproject.io/api-docs/auth/cert#token_num_uses)
     * of times a generated token may be used (within its lifetime); 0 means unlimited.
     * 
     */
    public Output<Optional<Integer>> tokenNumUses() {
        return Codegen.optional(this.tokenNumUses);
    }
    /**
     * If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    @Export(name="tokenPeriod", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenPeriod;

    /**
     * @return If set, indicates that the
     * token generated using this role should never expire. The token should be renewed within the
     * duration specified by this value. At each renewal, the token&#39;s TTL will be set to the
     * value of this field. Specified in seconds.
     * 
     */
    public Output<Optional<Integer>> tokenPeriod() {
        return Codegen.optional(this.tokenPeriod);
    }
    /**
     * List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    @Export(name="tokenPolicies", refs={List.class,String.class}, tree="[0,1]")
    private Output</* @Nullable */ List<String>> tokenPolicies;

    /**
     * @return List of policies to encode onto generated tokens. Depending
     * on the auth method, this list may be supplemented by user/group/other values.
     * 
     */
    public Output<Optional<List<String>>> tokenPolicies() {
        return Codegen.optional(this.tokenPolicies);
    }
    /**
     * The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    @Export(name="tokenTtl", refs={Integer.class}, tree="[0]")
    private Output</* @Nullable */ Integer> tokenTtl;

    /**
     * @return The incremental lifetime for generated tokens in number of seconds.
     * Its current value will be referenced at renewal time.
     * 
     */
    public Output<Optional<Integer>> tokenTtl() {
        return Codegen.optional(this.tokenTtl);
    }
    /**
     * The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     * For more details on the usage of each argument consult the [Vault Cert API documentation](https://www.vaultproject.io/api-docs/auth/cert).
     * 
     */
    @Export(name="tokenType", refs={String.class}, tree="[0]")
    private Output</* @Nullable */ String> tokenType;

    /**
     * @return The type of token that should be generated. Can be `service`,
     * `batch`, or `default` to use the mount&#39;s tuned default (which unless changed will be
     * `service` tokens). For token store roles, there are two additional possibilities:
     * `default-service` and `default-batch` which specify the type to return unless the client
     * requests a different type at generation time.
     * 
     * For more details on the usage of each argument consult the [Vault Cert API documentation](https://www.vaultproject.io/api-docs/auth/cert).
     * 
     */
    public Output<Optional<String>> tokenType() {
        return Codegen.optional(this.tokenType);
    }

    /**
     *
     * @param name The _unique_ name of the resulting resource.
     */
    public CertAuthBackendRole(String name) {
        this(name, CertAuthBackendRoleArgs.Empty);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     */
    public CertAuthBackendRole(String name, CertAuthBackendRoleArgs args) {
        this(name, args, null);
    }
    /**
     *
     * @param name The _unique_ name of the resulting resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param options A bag of options that control this resource's behavior.
     */
    public CertAuthBackendRole(String name, CertAuthBackendRoleArgs args, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/certAuthBackendRole:CertAuthBackendRole", name, args == null ? CertAuthBackendRoleArgs.Empty : args, makeResourceOptions(options, Codegen.empty()));
    }

    private CertAuthBackendRole(String name, Output<String> id, @Nullable CertAuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        super("vault:index/certAuthBackendRole:CertAuthBackendRole", name, state, makeResourceOptions(options, id));
    }

    private static com.pulumi.resources.CustomResourceOptions makeResourceOptions(@Nullable com.pulumi.resources.CustomResourceOptions options, @Nullable Output<String> id) {
        var defaultOptions = com.pulumi.resources.CustomResourceOptions.builder()
            .version(Utilities.getVersion())
            .build();
        return com.pulumi.resources.CustomResourceOptions.merge(defaultOptions, options, id);
    }

    /**
     * Get an existing Host resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state
     * @param options Optional settings to control the behavior of the CustomResource.
     */
    public static CertAuthBackendRole get(String name, Output<String> id, @Nullable CertAuthBackendRoleState state, @Nullable com.pulumi.resources.CustomResourceOptions options) {
        return new CertAuthBackendRole(name, id, state, options);
    }
}
