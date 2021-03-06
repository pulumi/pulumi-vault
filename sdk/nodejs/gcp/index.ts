// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./authBackend";
export * from "./authBackendRole";
export * from "./secretBackend";
export * from "./secretRoleset";

// Import resources to register:
import { AuthBackend } from "./authBackend";
import { AuthBackendRole } from "./authBackendRole";
import { SecretBackend } from "./secretBackend";
import { SecretRoleset } from "./secretRoleset";

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
            case "vault:gcp/secretRoleset:SecretRoleset":
                return new SecretRoleset(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "gcp/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/authBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "gcp/secretRoleset", _module)
