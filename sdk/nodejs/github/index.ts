// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export { AuthBackendArgs, AuthBackendState } from "./authBackend";
export type AuthBackend = import("./authBackend").AuthBackend;
export const AuthBackend: typeof import("./authBackend").AuthBackend = null as any;
utilities.lazyLoad(exports, ["AuthBackend"], () => require("./authBackend"));

export { TeamArgs, TeamState } from "./team";
export type Team = import("./team").Team;
export const Team: typeof import("./team").Team = null as any;
utilities.lazyLoad(exports, ["Team"], () => require("./team"));

export { UserArgs, UserState } from "./user";
export type User = import("./user").User;
export const User: typeof import("./user").User = null as any;
utilities.lazyLoad(exports, ["User"], () => require("./user"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:github/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:github/team:Team":
                return new Team(name, <any>undefined, { urn })
            case "vault:github/user:User":
                return new User(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "github/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "github/team", _module)
pulumi.runtime.registerResourceModule("vault", "github/user", _module)