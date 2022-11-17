// *** WARNING: this file was generated by pulumi-java-gen. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package com.pulumi.vault;

import com.pulumi.core.TypeShape;
import com.pulumi.core.internal.Codegen;
import com.pulumi.vault.config.inputs.AuthLogin;
import com.pulumi.vault.config.inputs.AuthLoginAws;
import com.pulumi.vault.config.inputs.AuthLoginAzure;
import com.pulumi.vault.config.inputs.AuthLoginCert;
import com.pulumi.vault.config.inputs.AuthLoginGcp;
import com.pulumi.vault.config.inputs.AuthLoginJwt;
import com.pulumi.vault.config.inputs.AuthLoginKerberos;
import com.pulumi.vault.config.inputs.AuthLoginOci;
import com.pulumi.vault.config.inputs.AuthLoginOidc;
import com.pulumi.vault.config.inputs.AuthLoginRadius;
import com.pulumi.vault.config.inputs.AuthLoginUserpass;
import com.pulumi.vault.config.inputs.ClientAuth;
import com.pulumi.vault.config.inputs.Headers;
import java.lang.Boolean;
import java.lang.Integer;
import java.lang.String;
import java.util.List;
import java.util.Optional;

public final class Config {

    private static final com.pulumi.Config config = com.pulumi.Config.of("vault");
/**
 * If true, adds the value of the `address` argument to the Terraform process environment.
 * 
 */
    public Optional<String> addAddressToEnv() {
        return Codegen.stringProp("addAddressToEnv").config(config).get();
    }
/**
 * URL of the root of the target Vault server.
 * 
 */
    public String address() {
        return Codegen.stringProp("address").config(config).require();
    }
/**
 * Login to vault with an existing auth method using auth/&lt;mount&gt;/login
 * 
 */
    public Optional<AuthLogin> authLogin() {
        return Codegen.objectProp("authLogin", AuthLogin.class).config(config).get();
    }
/**
 * Login to vault using the AWS method
 * 
 */
    public Optional<AuthLoginAws> authLoginAws() {
        return Codegen.objectProp("authLoginAws", AuthLoginAws.class).config(config).get();
    }
/**
 * Login to vault using the azure method
 * 
 */
    public Optional<AuthLoginAzure> authLoginAzure() {
        return Codegen.objectProp("authLoginAzure", AuthLoginAzure.class).config(config).get();
    }
/**
 * Login to vault using the cert method
 * 
 */
    public Optional<AuthLoginCert> authLoginCert() {
        return Codegen.objectProp("authLoginCert", AuthLoginCert.class).config(config).get();
    }
/**
 * Login to vault using the gcp method
 * 
 */
    public Optional<AuthLoginGcp> authLoginGcp() {
        return Codegen.objectProp("authLoginGcp", AuthLoginGcp.class).config(config).get();
    }
/**
 * Login to vault using the jwt method
 * 
 */
    public Optional<AuthLoginJwt> authLoginJwt() {
        return Codegen.objectProp("authLoginJwt", AuthLoginJwt.class).config(config).get();
    }
/**
 * Login to vault using the kerberos method
 * 
 */
    public Optional<AuthLoginKerberos> authLoginKerberos() {
        return Codegen.objectProp("authLoginKerberos", AuthLoginKerberos.class).config(config).get();
    }
/**
 * Login to vault using the OCI method
 * 
 */
    public Optional<AuthLoginOci> authLoginOci() {
        return Codegen.objectProp("authLoginOci", AuthLoginOci.class).config(config).get();
    }
/**
 * Login to vault using the oidc method
 * 
 */
    public Optional<AuthLoginOidc> authLoginOidc() {
        return Codegen.objectProp("authLoginOidc", AuthLoginOidc.class).config(config).get();
    }
/**
 * Login to vault using the radius method
 * 
 */
    public Optional<AuthLoginRadius> authLoginRadius() {
        return Codegen.objectProp("authLoginRadius", AuthLoginRadius.class).config(config).get();
    }
/**
 * Login to vault using the userpass method
 * 
 */
    public Optional<AuthLoginUserpass> authLoginUserpass() {
        return Codegen.objectProp("authLoginUserpass", AuthLoginUserpass.class).config(config).get();
    }
/**
 * Path to directory containing CA certificate files to validate the server&#39;s certificate.
 * 
 */
    public Optional<String> caCertDir() {
        return Codegen.stringProp("caCertDir").config(config).get();
    }
/**
 * Path to a CA certificate file to validate the server&#39;s certificate.
 * 
 */
    public Optional<String> caCertFile() {
        return Codegen.stringProp("caCertFile").config(config).get();
    }
/**
 * Client authentication credentials.
 * 
 */
    public Optional<ClientAuth> clientAuth() {
        return Codegen.objectProp("clientAuth", ClientAuth.class).config(config).get();
    }
/**
 * The headers to send with each Vault request.
 * 
 */
    public Optional<List<Headers>> headers() {
        return Codegen.objectProp("headers", TypeShape.<List<Headers>>builder(List.class).addParameter(Headers.class).build()).config(config).get();
    }
/**
 * Maximum TTL for secret leases requested by this provider.
 * 
 */
    public Optional<Integer> maxLeaseTtlSeconds() {
        return Codegen.integerProp("maxLeaseTtlSeconds").config(config).env("TERRAFORM_VAULT_MAX_TTL").def(1200).get();
    }
/**
 * Maximum number of retries when a 5xx error code is encountered.
 * 
 */
    public Optional<Integer> maxRetries() {
        return Codegen.integerProp("maxRetries").config(config).env("VAULT_MAX_RETRIES").def(2).get();
    }
/**
 * Maximum number of retries for Client Controlled Consistency related operations
 * 
 */
    public Optional<Integer> maxRetriesCcc() {
        return Codegen.integerProp("maxRetriesCcc").config(config).get();
    }
/**
 * The namespace to use. Available only for Vault Enterprise.
 * 
 */
    public Optional<String> namespace() {
        return Codegen.stringProp("namespace").config(config).get();
    }
/**
 * Set this to true to prevent the creation of ephemeral child token used by this provider.
 * 
 */
    public Optional<Boolean> skipChildToken() {
        return Codegen.booleanProp("skipChildToken").config(config).get();
    }
/**
 * Set this to true only if the target Vault server is an insecure development instance.
 * 
 */
    public Optional<Boolean> skipTlsVerify() {
        return Codegen.booleanProp("skipTlsVerify").config(config).env("VAULT_SKIP_VERIFY").get();
    }
/**
 * Name to use as the SNI host when connecting via TLS.
 * 
 */
    public Optional<String> tlsServerName() {
        return Codegen.stringProp("tlsServerName").config(config).get();
    }
/**
 * Token to use to authenticate to Vault.
 * 
 */
    public String token() {
        return Codegen.stringProp("token").config(config).require();
    }
/**
 * Token name to use for creating the Vault child token.
 * 
 */
    public Optional<String> tokenName() {
        return Codegen.stringProp("tokenName").config(config).get();
    }
}