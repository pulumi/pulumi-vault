// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./alphabet";
export * from "./getDecode";
export * from "./getEncode";
export * from "./role";
export * from "./template";
export * from "./transformation";

// Import resources to register:
import { Alphabet } from "./alphabet";
import { Role } from "./role";
import { Template } from "./template";
import { Transformation } from "./transformation";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:transform/alphabet:Alphabet":
                return new Alphabet(name, <any>undefined, { urn })
            case "vault:transform/role:Role":
                return new Role(name, <any>undefined, { urn })
            case "vault:transform/template:Template":
                return new Template(name, <any>undefined, { urn })
            case "vault:transform/transformation:Transformation":
                return new Transformation(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "transform/alphabet", _module)
pulumi.runtime.registerResourceModule("vault", "transform/role", _module)
pulumi.runtime.registerResourceModule("vault", "transform/template", _module)
pulumi.runtime.registerResourceModule("vault", "transform/transformation", _module)
