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


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:saml/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:saml/authBackendRole:AuthBackendRole":
                return new AuthBackendRole(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "saml/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "saml/authBackendRole", _module)
