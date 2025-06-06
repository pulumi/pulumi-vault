// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthBackendArgs, AuthBackendState } from "./authBackend";
export type AuthBackend = import("./authBackend").AuthBackend;
export const AuthBackend: typeof import("./authBackend").AuthBackend = null as any;
utilities.lazyLoad(exports, ["AuthBackend"], () => require("./authBackend"));

export { AuthBackendRoleArgs, AuthBackendRoleState } from "./authBackendRole";
export type AuthBackendRole = import("./authBackendRole").AuthBackendRole;
export const AuthBackendRole: typeof import("./authBackendRole").AuthBackendRole = null as any;
utilities.lazyLoad(exports, ["AuthBackendRole"], () => require("./authBackendRole"));

export { GetAuthBackendRoleArgs, GetAuthBackendRoleResult, GetAuthBackendRoleOutputArgs } from "./getAuthBackendRole";
export const getAuthBackendRole: typeof import("./getAuthBackendRole").getAuthBackendRole = null as any;
export const getAuthBackendRoleOutput: typeof import("./getAuthBackendRole").getAuthBackendRoleOutput = null as any;
utilities.lazyLoad(exports, ["getAuthBackendRole","getAuthBackendRoleOutput"], () => require("./getAuthBackendRole"));

export { SecretBackendArgs, SecretBackendState } from "./secretBackend";
export type SecretBackend = import("./secretBackend").SecretBackend;
export const SecretBackend: typeof import("./secretBackend").SecretBackend = null as any;
utilities.lazyLoad(exports, ["SecretBackend"], () => require("./secretBackend"));

export { SecretImpersonatedAccountArgs, SecretImpersonatedAccountState } from "./secretImpersonatedAccount";
export type SecretImpersonatedAccount = import("./secretImpersonatedAccount").SecretImpersonatedAccount;
export const SecretImpersonatedAccount: typeof import("./secretImpersonatedAccount").SecretImpersonatedAccount = null as any;
utilities.lazyLoad(exports, ["SecretImpersonatedAccount"], () => require("./secretImpersonatedAccount"));

export { SecretRolesetArgs, SecretRolesetState } from "./secretRoleset";
export type SecretRoleset = import("./secretRoleset").SecretRoleset;
export const SecretRoleset: typeof import("./secretRoleset").SecretRoleset = null as any;
utilities.lazyLoad(exports, ["SecretRoleset"], () => require("./secretRoleset"));

export { SecretStaticAccountArgs, SecretStaticAccountState } from "./secretStaticAccount";
export type SecretStaticAccount = import("./secretStaticAccount").SecretStaticAccount;
export const SecretStaticAccount: typeof import("./secretStaticAccount").SecretStaticAccount = null as any;
utilities.lazyLoad(exports, ["SecretStaticAccount"], () => require("./secretStaticAccount"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:gcp/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:gcp/authBackendRole:AuthBackendRole":
                return new AuthBackendRole(name, <any>undefined, { urn })
            case "vault:gcp/secretBackend:SecretBackend":
                return new SecretBackend(name, <any>undefined, { urn })
            case "vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount":
                return new SecretImpersonatedAccount(name, <any>undefined, { urn })
            case "vault:gcp/secretRoleset:SecretRoleset":
                return new SecretRoleset(name, <any>undefined, { urn })
            case "vault:gcp/secretStaticAccount:SecretStaticAccount":
                return new SecretStaticAccount(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "gcp/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/authBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretImpersonatedAccount", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretRoleset", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretStaticAccount", _module)
