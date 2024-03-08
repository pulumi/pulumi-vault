// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * Resource for configuring the totp MFA method.
 *
 * ## Example Usage
 *
 * <!--Start PulumiCodeChooser -->
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vault from "@pulumi/vault";
 *
 * const example = new vault.identity.MfaTotp("example", {issuer: "issuer1"});
 * ```
 * <!--End PulumiCodeChooser -->
 *
 * ## Import
 *
 * Resource can be imported using its `uuid` field, e.g.
 *
 * ```sh
 * $ pulumi import vault:identity/mfaTotp:MfaTotp example 0d89c36a-4ff5-4d70-8749-bb6a5598aeec
 * ```
 */
export class MfaTotp extends pulumi.CustomResource {
    /**
     * Get an existing MfaTotp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: MfaTotpState, opts?: pulumi.CustomResourceOptions): MfaTotp {
        return new MfaTotp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:identity/mfaTotp:MfaTotp';

    /**
     * Returns true if the given object is an instance of MfaTotp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is MfaTotp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === MfaTotp.__pulumiType;
    }

    /**
     * Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     */
    public readonly algorithm!: pulumi.Output<string | undefined>;
    /**
     * The number of digits in the generated TOTP token. This value can either be 6 or 8
     */
    public readonly digits!: pulumi.Output<number | undefined>;
    /**
     * The name of the key's issuing organization.
     */
    public readonly issuer!: pulumi.Output<string>;
    /**
     * Specifies the size in bytes of the generated key.
     */
    public readonly keySize!: pulumi.Output<number | undefined>;
    /**
     * The maximum number of consecutive failed validation attempts allowed.
     */
    public readonly maxValidationAttempts!: pulumi.Output<number | undefined>;
    /**
     * Method ID.
     */
    public /*out*/ readonly methodId!: pulumi.Output<string>;
    /**
     * Mount accessor.
     */
    public /*out*/ readonly mountAccessor!: pulumi.Output<string>;
    /**
     * Method name.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    public readonly namespace!: pulumi.Output<string | undefined>;
    /**
     * Method's namespace ID.
     */
    public /*out*/ readonly namespaceId!: pulumi.Output<string>;
    /**
     * Method's namespace path.
     */
    public /*out*/ readonly namespacePath!: pulumi.Output<string>;
    /**
     * The length of time in seconds used to generate a counter for the TOTP token calculation.
     */
    public readonly period!: pulumi.Output<number | undefined>;
    /**
     * The pixel size of the generated square QR code.
     */
    public readonly qrSize!: pulumi.Output<number>;
    /**
     * The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     */
    public readonly skew!: pulumi.Output<number | undefined>;
    /**
     * MFA type.
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Resource UUID.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a MfaTotp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: MfaTotpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: MfaTotpArgs | MfaTotpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as MfaTotpState | undefined;
            resourceInputs["algorithm"] = state ? state.algorithm : undefined;
            resourceInputs["digits"] = state ? state.digits : undefined;
            resourceInputs["issuer"] = state ? state.issuer : undefined;
            resourceInputs["keySize"] = state ? state.keySize : undefined;
            resourceInputs["maxValidationAttempts"] = state ? state.maxValidationAttempts : undefined;
            resourceInputs["methodId"] = state ? state.methodId : undefined;
            resourceInputs["mountAccessor"] = state ? state.mountAccessor : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["namespace"] = state ? state.namespace : undefined;
            resourceInputs["namespaceId"] = state ? state.namespaceId : undefined;
            resourceInputs["namespacePath"] = state ? state.namespacePath : undefined;
            resourceInputs["period"] = state ? state.period : undefined;
            resourceInputs["qrSize"] = state ? state.qrSize : undefined;
            resourceInputs["skew"] = state ? state.skew : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as MfaTotpArgs | undefined;
            if ((!args || args.issuer === undefined) && !opts.urn) {
                throw new Error("Missing required property 'issuer'");
            }
            resourceInputs["algorithm"] = args ? args.algorithm : undefined;
            resourceInputs["digits"] = args ? args.digits : undefined;
            resourceInputs["issuer"] = args ? args.issuer : undefined;
            resourceInputs["keySize"] = args ? args.keySize : undefined;
            resourceInputs["maxValidationAttempts"] = args ? args.maxValidationAttempts : undefined;
            resourceInputs["namespace"] = args ? args.namespace : undefined;
            resourceInputs["period"] = args ? args.period : undefined;
            resourceInputs["qrSize"] = args ? args.qrSize : undefined;
            resourceInputs["skew"] = args ? args.skew : undefined;
            resourceInputs["methodId"] = undefined /*out*/;
            resourceInputs["mountAccessor"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["namespaceId"] = undefined /*out*/;
            resourceInputs["namespacePath"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(MfaTotp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering MfaTotp resources.
 */
export interface MfaTotpState {
    /**
     * Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The number of digits in the generated TOTP token. This value can either be 6 or 8
     */
    digits?: pulumi.Input<number>;
    /**
     * The name of the key's issuing organization.
     */
    issuer?: pulumi.Input<string>;
    /**
     * Specifies the size in bytes of the generated key.
     */
    keySize?: pulumi.Input<number>;
    /**
     * The maximum number of consecutive failed validation attempts allowed.
     */
    maxValidationAttempts?: pulumi.Input<number>;
    /**
     * Method ID.
     */
    methodId?: pulumi.Input<string>;
    /**
     * Mount accessor.
     */
    mountAccessor?: pulumi.Input<string>;
    /**
     * Method name.
     */
    name?: pulumi.Input<string>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * Method's namespace ID.
     */
    namespaceId?: pulumi.Input<string>;
    /**
     * Method's namespace path.
     */
    namespacePath?: pulumi.Input<string>;
    /**
     * The length of time in seconds used to generate a counter for the TOTP token calculation.
     */
    period?: pulumi.Input<number>;
    /**
     * The pixel size of the generated square QR code.
     */
    qrSize?: pulumi.Input<number>;
    /**
     * The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     */
    skew?: pulumi.Input<number>;
    /**
     * MFA type.
     */
    type?: pulumi.Input<string>;
    /**
     * Resource UUID.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a MfaTotp resource.
 */
export interface MfaTotpArgs {
    /**
     * Specifies the hashing algorithm used to generate the TOTP code. Options include SHA1, SHA256, SHA512.
     */
    algorithm?: pulumi.Input<string>;
    /**
     * The number of digits in the generated TOTP token. This value can either be 6 or 8
     */
    digits?: pulumi.Input<number>;
    /**
     * The name of the key's issuing organization.
     */
    issuer: pulumi.Input<string>;
    /**
     * Specifies the size in bytes of the generated key.
     */
    keySize?: pulumi.Input<number>;
    /**
     * The maximum number of consecutive failed validation attempts allowed.
     */
    maxValidationAttempts?: pulumi.Input<number>;
    /**
     * Target namespace. (requires Enterprise)
     */
    namespace?: pulumi.Input<string>;
    /**
     * The length of time in seconds used to generate a counter for the TOTP token calculation.
     */
    period?: pulumi.Input<number>;
    /**
     * The pixel size of the generated square QR code.
     */
    qrSize?: pulumi.Input<number>;
    /**
     * The number of delay periods that are allowed when validating a TOTP token. This value can either be 0 or 1.
     */
    skew?: pulumi.Input<number>;
}
