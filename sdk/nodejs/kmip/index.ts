// *** WARNING: this file was generated by pulumi-language-nodejs. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { SecretBackendArgs, SecretBackendState } from "./secretBackend";
export type SecretBackend = import("./secretBackend").SecretBackend;
export const SecretBackend: typeof import("./secretBackend").SecretBackend = null as any;
utilities.lazyLoad(exports, ["SecretBackend"], () => require("./secretBackend"));

export { SecretRoleArgs, SecretRoleState } from "./secretRole";
export type SecretRole = import("./secretRole").SecretRole;
export const SecretRole: typeof import("./secretRole").SecretRole = null as any;
utilities.lazyLoad(exports, ["SecretRole"], () => require("./secretRole"));

export { SecretScopeArgs, SecretScopeState } from "./secretScope";
export type SecretScope = import("./secretScope").SecretScope;
export const SecretScope: typeof import("./secretScope").SecretScope = null as any;
utilities.lazyLoad(exports, ["SecretScope"], () => require("./secretScope"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:kmip/secretBackend:SecretBackend":
                return new SecretBackend(name, <any>undefined, { urn })
            case "vault:kmip/secretRole:SecretRole":
                return new SecretRole(name, <any>undefined, { urn })
            case "vault:kmip/secretScope:SecretScope":
                return new SecretScope(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "kmip/secretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "kmip/secretRole", _module)
pulumi.runtime.registerResourceModule("vault", "kmip/secretScope", _module)
