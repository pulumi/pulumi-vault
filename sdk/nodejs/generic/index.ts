// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./endpoint";
export * from "./getSecret";
export * from "./secret";

// Import resources to register:
import { Endpoint } from "./endpoint";
import { Secret } from "./secret";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:generic/endpoint:Endpoint":
                return new Endpoint(name, <any>undefined, { urn })
            case "vault:generic/secret:Secret":
                return new Secret(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "generic/endpoint", _module)
pulumi.runtime.registerResourceModule("vault", "generic/secret", _module)
