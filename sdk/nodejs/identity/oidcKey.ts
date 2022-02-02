// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * The key can be imported with the key name, for example
 *
 * ```sh
 *  $ pulumi import vault:identity/oidcKey:OidcKey key key
 * ```
 */
export class OidcKey extends pulumi.CustomResource {
    /**
     * Get an existing OidcKey resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OidcKeyState, opts?: pulumi.CustomResourceOptions): OidcKey {
        return new OidcKey(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/oidcKey:OidcKey';

    /**
     * Returns true if the given object is an instance of OidcKey.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OidcKey {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OidcKey.__pulumiType;
    }

    /**
     * Signing algorithm to use. Signing algorithm to use.
     * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
     */
    public readonly algorithm!: pulumi.Output<string | undefined>;
    /**
     * Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
     * allowed.
     */
    public readonly allowedClientIds!: pulumi.Output<string[]>;
    /**
     * Name of the OIDC Key to create.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * How often to generate a new signing key in number of seconds
     */
    public readonly rotationPeriod!: pulumi.Output<number | undefined>;
    /**
     * "Controls how long the public portion of a signing key will be
     * available for verification after being rotated in seconds.
     */
    public readonly verificationTtl!: pulumi.Output<number | undefined>;

    /**
     * Create a OidcKey resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OidcKeyArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OidcKeyArgs | OidcKeyState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OidcKeyState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["allowedClientIds"] = state ? state.allowedClientIds : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rotationPeriod"] = state ? state.rotationPeriod : undefined;
            resourceInputs["verificationTtl"] = state ? state.verificationTtl : undefined;
        } else {
            const args = argsOrState as OidcKeyArgs | undefined;
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["allowedClientIds"] = args ? args.allowedClientIds : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rotationPeriod"] = args ? args.rotationPeriod : undefined;
            resourceInputs["verificationTtl"] = args ? args.verificationTtl : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OidcKey.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OidcKey resources.
 */
export interface OidcKeyState {
    /**
     * Signing algorithm to use. Signing algorithm to use.
     * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
     * allowed.
     */
    allowedClientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the OIDC Key to create.
     */
    name?: pulumi.Input<string>;
    /**
     * How often to generate a new signing key in number of seconds
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * "Controls how long the public portion of a signing key will be
     * available for verification after being rotated in seconds.
     */
    verificationTtl?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OidcKey resource.
 */
export interface OidcKeyArgs {
    /**
     * Signing algorithm to use. Signing algorithm to use.
     * Allowed values are: RS256 (default), RS384, RS512, ES256, ES384, ES512, EdDSA.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * Array of role client ids allowed to use this key for signing. If empty, no roles are allowed. If "*", all roles are
     * allowed.
     */
    allowedClientIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Name of the OIDC Key to create.
     */
    name?: pulumi.Input<string>;
    /**
     * How often to generate a new signing key in number of seconds
     */
    rotationPeriod?: pulumi.Input<number>;
    /**
     * "Controls how long the public portion of a signing key will be
     * available for verification after being rotated in seconds.
     */
    verificationTtl?: pulumi.Input<number>;
}
