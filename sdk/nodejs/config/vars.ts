// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../types/input";
import * as outputs from "../types/output";
import * as utilities from "../utilities";

declare var exports: any;
const __config = new pulumi.Config("vault");

/**
 * If true, adds the value of the `address` argument to the Terraform process environment.
 */
export declare const addAddressToEnv: string | undefined;
Object.defineProperty(exports, "addAddressToEnv", {
    get() {
        return __config.get("addAddressToEnv");
    },
    enumerable: true,
});

/**
 * URL of the root of the target Vault server.
 */
export declare const address: string | undefined;
Object.defineProperty(exports, "address", {
    get() {
        return __config.get("address");
    },
    enumerable: true,
});

/**
 * Login to vault with an existing auth method using auth/<mount>/login
 */
export declare const authLogin: outputs.config.AuthLogin | undefined;
Object.defineProperty(exports, "authLogin", {
    get() {
        return __config.getObject<outputs.config.AuthLogin>("authLogin");
    },
    enumerable: true,
});

/**
 * Login to vault using the AWS method
 */
export declare const authLoginAws: outputs.config.AuthLoginAws | undefined;
Object.defineProperty(exports, "authLoginAws", {
    get() {
        return __config.getObject<outputs.config.AuthLoginAws>("authLoginAws");
    },
    enumerable: true,
});

/**
 * Login to vault using the azure method
 */
export declare const authLoginAzure: outputs.config.AuthLoginAzure | undefined;
Object.defineProperty(exports, "authLoginAzure", {
    get() {
        return __config.getObject<outputs.config.AuthLoginAzure>("authLoginAzure");
    },
    enumerable: true,
});

/**
 * Login to vault using the cert method
 */
export declare const authLoginCert: outputs.config.AuthLoginCert | undefined;
Object.defineProperty(exports, "authLoginCert", {
    get() {
        return __config.getObject<outputs.config.AuthLoginCert>("authLoginCert");
    },
    enumerable: true,
});

/**
 * Login to vault using the gcp method
 */
export declare const authLoginGcp: outputs.config.AuthLoginGcp | undefined;
Object.defineProperty(exports, "authLoginGcp", {
    get() {
        return __config.getObject<outputs.config.AuthLoginGcp>("authLoginGcp");
    },
    enumerable: true,
});

/**
 * Login to vault using the jwt method
 */
export declare const authLoginJwt: outputs.config.AuthLoginJwt | undefined;
Object.defineProperty(exports, "authLoginJwt", {
    get() {
        return __config.getObject<outputs.config.AuthLoginJwt>("authLoginJwt");
    },
    enumerable: true,
});

/**
 * Login to vault using the kerberos method
 */
export declare const authLoginKerberos: outputs.config.AuthLoginKerberos | undefined;
Object.defineProperty(exports, "authLoginKerberos", {
    get() {
        return __config.getObject<outputs.config.AuthLoginKerberos>("authLoginKerberos");
    },
    enumerable: true,
});

/**
 * Login to vault using the OCI method
 */
export declare const authLoginOci: outputs.config.AuthLoginOci | undefined;
Object.defineProperty(exports, "authLoginOci", {
    get() {
        return __config.getObject<outputs.config.AuthLoginOci>("authLoginOci");
    },
    enumerable: true,
});

/**
 * Login to vault using the oidc method
 */
export declare const authLoginOidc: outputs.config.AuthLoginOidc | undefined;
Object.defineProperty(exports, "authLoginOidc", {
    get() {
        return __config.getObject<outputs.config.AuthLoginOidc>("authLoginOidc");
    },
    enumerable: true,
});

/**
 * Login to vault using the radius method
 */
export declare const authLoginRadius: outputs.config.AuthLoginRadius | undefined;
Object.defineProperty(exports, "authLoginRadius", {
    get() {
        return __config.getObject<outputs.config.AuthLoginRadius>("authLoginRadius");
    },
    enumerable: true,
});

/**
 * Login to vault using
 */
export declare const authLoginTokenFile: outputs.config.AuthLoginTokenFile | undefined;
Object.defineProperty(exports, "authLoginTokenFile", {
    get() {
        return __config.getObject<outputs.config.AuthLoginTokenFile>("authLoginTokenFile");
    },
    enumerable: true,
});

/**
 * Login to vault using the userpass method
 */
export declare const authLoginUserpass: outputs.config.AuthLoginUserpass | undefined;
Object.defineProperty(exports, "authLoginUserpass", {
    get() {
        return __config.getObject<outputs.config.AuthLoginUserpass>("authLoginUserpass");
    },
    enumerable: true,
});

/**
 * Path to directory containing CA certificate files to validate the server's certificate.
 */
export declare const caCertDir: string | undefined;
Object.defineProperty(exports, "caCertDir", {
    get() {
        return __config.get("caCertDir");
    },
    enumerable: true,
});

/**
 * Path to a CA certificate file to validate the server's certificate.
 */
export declare const caCertFile: string | undefined;
Object.defineProperty(exports, "caCertFile", {
    get() {
        return __config.get("caCertFile");
    },
    enumerable: true,
});

/**
 * Client authentication credentials.
 */
export declare const clientAuth: outputs.config.ClientAuth | undefined;
Object.defineProperty(exports, "clientAuth", {
    get() {
        return __config.getObject<outputs.config.ClientAuth>("clientAuth");
    },
    enumerable: true,
});

/**
 * The headers to send with each Vault request.
 */
export declare const headers: outputs.config.Headers[] | undefined;
Object.defineProperty(exports, "headers", {
    get() {
        return __config.getObject<outputs.config.Headers[]>("headers");
    },
    enumerable: true,
});

/**
 * Maximum TTL for secret leases requested by this provider.
 */
export declare const maxLeaseTtlSeconds: number;
Object.defineProperty(exports, "maxLeaseTtlSeconds", {
    get() {
        return __config.getObject<number>("maxLeaseTtlSeconds") ?? (utilities.getEnvNumber("TERRAFORM_VAULT_MAX_TTL") || 1200);
    },
    enumerable: true,
});

/**
 * Maximum number of retries when a 5xx error code is encountered.
 */
export declare const maxRetries: number;
Object.defineProperty(exports, "maxRetries", {
    get() {
        return __config.getObject<number>("maxRetries") ?? (utilities.getEnvNumber("VAULT_MAX_RETRIES") || 2);
    },
    enumerable: true,
});

/**
 * Maximum number of retries for Client Controlled Consistency related operations
 */
export declare const maxRetriesCcc: number | undefined;
Object.defineProperty(exports, "maxRetriesCcc", {
    get() {
        return __config.getObject<number>("maxRetriesCcc");
    },
    enumerable: true,
});

/**
 * The namespace to use. Available only for Vault Enterprise.
 */
export declare const namespace: string | undefined;
Object.defineProperty(exports, "namespace", {
    get() {
        return __config.get("namespace");
    },
    enumerable: true,
});

/**
 * Set this to true to prevent the creation of ephemeral child token used by this provider.
 */
export declare const skipChildToken: boolean | undefined;
Object.defineProperty(exports, "skipChildToken", {
    get() {
        return __config.getObject<boolean>("skipChildToken");
    },
    enumerable: true,
});

/**
 * Skip the dynamic fetching of the Vault server version.
 */
export declare const skipGetVaultVersion: boolean | undefined;
Object.defineProperty(exports, "skipGetVaultVersion", {
    get() {
        return __config.getObject<boolean>("skipGetVaultVersion");
    },
    enumerable: true,
});

/**
 * Set this to true only if the target Vault server is an insecure development instance.
 */
export declare const skipTlsVerify: boolean | undefined;
Object.defineProperty(exports, "skipTlsVerify", {
    get() {
        return __config.getObject<boolean>("skipTlsVerify") ?? utilities.getEnvBoolean("VAULT_SKIP_VERIFY");
    },
    enumerable: true,
});

/**
 * Name to use as the SNI host when connecting via TLS.
 */
export declare const tlsServerName: string | undefined;
Object.defineProperty(exports, "tlsServerName", {
    get() {
        return __config.get("tlsServerName");
    },
    enumerable: true,
});

/**
 * Token to use to authenticate to Vault.
 */
export declare const token: string | undefined;
Object.defineProperty(exports, "token", {
    get() {
        return __config.get("token");
    },
    enumerable: true,
});

/**
 * Token name to use for creating the Vault child token.
 */
export declare const tokenName: string | undefined;
Object.defineProperty(exports, "tokenName", {
    get() {
        return __config.get("tokenName");
    },
    enumerable: true,
});

/**
 * Override the Vault server version, which is normally determined dynamically from the target Vault server
 */
export declare const vaultVersionOverride: string | undefined;
Object.defineProperty(exports, "vaultVersionOverride", {
    get() {
        return __config.get("vaultVersionOverride");
    },
    enumerable: true,
});

