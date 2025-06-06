// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthBackendCertArgs, AuthBackendCertState } from "./authBackendCert";
export type AuthBackendCert = import("./authBackendCert").AuthBackendCert;
export const AuthBackendCert: typeof import("./authBackendCert").AuthBackendCert = null as any;
utilities.lazyLoad(exports, ["AuthBackendCert"], () => require("./authBackendCert"));

export { AuthBackendClientArgs, AuthBackendClientState } from "./authBackendClient";
export type AuthBackendClient = import("./authBackendClient").AuthBackendClient;
export const AuthBackendClient: typeof import("./authBackendClient").AuthBackendClient = null as any;
utilities.lazyLoad(exports, ["AuthBackendClient"], () => require("./authBackendClient"));

export { AuthBackendConfigIdentityArgs, AuthBackendConfigIdentityState } from "./authBackendConfigIdentity";
export type AuthBackendConfigIdentity = import("./authBackendConfigIdentity").AuthBackendConfigIdentity;
export const AuthBackendConfigIdentity: typeof import("./authBackendConfigIdentity").AuthBackendConfigIdentity = null as any;
utilities.lazyLoad(exports, ["AuthBackendConfigIdentity"], () => require("./authBackendConfigIdentity"));

export { AuthBackendIdentityWhitelistArgs, AuthBackendIdentityWhitelistState } from "./authBackendIdentityWhitelist";
export type AuthBackendIdentityWhitelist = import("./authBackendIdentityWhitelist").AuthBackendIdentityWhitelist;
export const AuthBackendIdentityWhitelist: typeof import("./authBackendIdentityWhitelist").AuthBackendIdentityWhitelist = null as any;
utilities.lazyLoad(exports, ["AuthBackendIdentityWhitelist"], () => require("./authBackendIdentityWhitelist"));

export { AuthBackendLoginArgs, AuthBackendLoginState } from "./authBackendLogin";
export type AuthBackendLogin = import("./authBackendLogin").AuthBackendLogin;
export const AuthBackendLogin: typeof import("./authBackendLogin").AuthBackendLogin = null as any;
utilities.lazyLoad(exports, ["AuthBackendLogin"], () => require("./authBackendLogin"));

export { AuthBackendRoleArgs, AuthBackendRoleState } from "./authBackendRole";
export type AuthBackendRole = import("./authBackendRole").AuthBackendRole;
export const AuthBackendRole: typeof import("./authBackendRole").AuthBackendRole = null as any;
utilities.lazyLoad(exports, ["AuthBackendRole"], () => require("./authBackendRole"));

export { AuthBackendRoleTagArgs, AuthBackendRoleTagState } from "./authBackendRoleTag";
export type AuthBackendRoleTag = import("./authBackendRoleTag").AuthBackendRoleTag;
export const AuthBackendRoleTag: typeof import("./authBackendRoleTag").AuthBackendRoleTag = null as any;
utilities.lazyLoad(exports, ["AuthBackendRoleTag"], () => require("./authBackendRoleTag"));

export { AuthBackendRoletagBlacklistArgs, AuthBackendRoletagBlacklistState } from "./authBackendRoletagBlacklist";
export type AuthBackendRoletagBlacklist = import("./authBackendRoletagBlacklist").AuthBackendRoletagBlacklist;
export const AuthBackendRoletagBlacklist: typeof import("./authBackendRoletagBlacklist").AuthBackendRoletagBlacklist = null as any;
utilities.lazyLoad(exports, ["AuthBackendRoletagBlacklist"], () => require("./authBackendRoletagBlacklist"));

export { AuthBackendStsRoleArgs, AuthBackendStsRoleState } from "./authBackendStsRole";
export type AuthBackendStsRole = import("./authBackendStsRole").AuthBackendStsRole;
export const AuthBackendStsRole: typeof import("./authBackendStsRole").AuthBackendStsRole = null as any;
utilities.lazyLoad(exports, ["AuthBackendStsRole"], () => require("./authBackendStsRole"));

export { GetAccessCredentialsArgs, GetAccessCredentialsResult, GetAccessCredentialsOutputArgs } from "./getAccessCredentials";
export const getAccessCredentials: typeof import("./getAccessCredentials").getAccessCredentials = null as any;
export const getAccessCredentialsOutput: typeof import("./getAccessCredentials").getAccessCredentialsOutput = null as any;
utilities.lazyLoad(exports, ["getAccessCredentials","getAccessCredentialsOutput"], () => require("./getAccessCredentials"));

export { GetStaticAccessCredentialsArgs, GetStaticAccessCredentialsResult, GetStaticAccessCredentialsOutputArgs } from "./getStaticAccessCredentials";
export const getStaticAccessCredentials: typeof import("./getStaticAccessCredentials").getStaticAccessCredentials = null as any;
export const getStaticAccessCredentialsOutput: typeof import("./getStaticAccessCredentials").getStaticAccessCredentialsOutput = null as any;
utilities.lazyLoad(exports, ["getStaticAccessCredentials","getStaticAccessCredentialsOutput"], () => require("./getStaticAccessCredentials"));

export { SecretBackendArgs, SecretBackendState } from "./secretBackend";
export type SecretBackend = import("./secretBackend").SecretBackend;
export const SecretBackend: typeof import("./secretBackend").SecretBackend = null as any;
utilities.lazyLoad(exports, ["SecretBackend"], () => require("./secretBackend"));

export { SecretBackendRoleArgs, SecretBackendRoleState } from "./secretBackendRole";
export type SecretBackendRole = import("./secretBackendRole").SecretBackendRole;
export const SecretBackendRole: typeof import("./secretBackendRole").SecretBackendRole = null as any;
utilities.lazyLoad(exports, ["SecretBackendRole"], () => require("./secretBackendRole"));

export { SecretBackendStaticRoleArgs, SecretBackendStaticRoleState } from "./secretBackendStaticRole";
export type SecretBackendStaticRole = import("./secretBackendStaticRole").SecretBackendStaticRole;
export const SecretBackendStaticRole: typeof import("./secretBackendStaticRole").SecretBackendStaticRole = null as any;
utilities.lazyLoad(exports, ["SecretBackendStaticRole"], () => require("./secretBackendStaticRole"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:aws/authBackendCert:AuthBackendCert":
                return new AuthBackendCert(name, <any>undefined, { urn })
            case "vault:aws/authBackendClient:AuthBackendClient":
                return new AuthBackendClient(name, <any>undefined, { urn })
            case "vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity":
                return new AuthBackendConfigIdentity(name, <any>undefined, { urn })
            case "vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist":
                return new AuthBackendIdentityWhitelist(name, <any>undefined, { urn })
            case "vault:aws/authBackendLogin:AuthBackendLogin":
                return new AuthBackendLogin(name, <any>undefined, { urn })
            case "vault:aws/authBackendRole:AuthBackendRole":
                return new AuthBackendRole(name, <any>undefined, { urn })
            case "vault:aws/authBackendRoleTag:AuthBackendRoleTag":
                return new AuthBackendRoleTag(name, <any>undefined, { urn })
            case "vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist":
                return new AuthBackendRoletagBlacklist(name, <any>undefined, { urn })
            case "vault:aws/authBackendStsRole:AuthBackendStsRole":
                return new AuthBackendStsRole(name, <any>undefined, { urn })
            case "vault:aws/secretBackend:SecretBackend":
                return new SecretBackend(name, <any>undefined, { urn })
            case "vault:aws/secretBackendRole:SecretBackendRole":
                return new SecretBackendRole(name, <any>undefined, { urn })
            case "vault:aws/secretBackendStaticRole:SecretBackendStaticRole":
                return new SecretBackendStaticRole(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "aws/authBackendCert", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendClient", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendConfigIdentity", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendIdentityWhitelist", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendLogin", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendRoleTag", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendRoletagBlacklist", _module)
pulumi.runtime.registerResourceModule("vault", "aws/authBackendStsRole", _module)
pulumi.runtime.registerResourceModule("vault", "aws/secretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "aws/secretBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "aws/secretBackendStaticRole", _module)
