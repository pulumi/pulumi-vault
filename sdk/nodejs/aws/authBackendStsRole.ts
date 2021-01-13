// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

/**
 * ## Import
 *
 * AWS auth backend STS roles can be imported using `auth/`, the `backend` path, `/config/sts/`, and the `account_id` e.g.
 *
 * ```sh
 *  $ pulumi import vault:aws/authBackendStsRole:AuthBackendStsRole example auth/aws/config/sts/1234567890
 * ```
 */
export class AuthBackendStsRole extends pulumi.CustomResource {
    /**
     * Get an existing AuthBackendStsRole resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: AuthBackendStsRoleState, opts?: pulumi.CustomResourceOptions): AuthBackendStsRole {
        return new AuthBackendStsRole(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vault:aws/authBackendStsRole:AuthBackendStsRole';

    /**
     * Returns true if the given object is an instance of AuthBackendStsRole.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is AuthBackendStsRole {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === AuthBackendStsRole.__pulumiType;
    }

    /**
     * The AWS account ID to configure the STS role for.
     */
    public readonly accountId!: pulumi.Output<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    public readonly backend!: pulumi.Output<string | undefined>;
    /**
     * The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `accountId`.
     */
    public readonly stsRole!: pulumi.Output<string>;

    /**
     * Create a AuthBackendStsRole resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: AuthBackendStsRoleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: AuthBackendStsRoleArgs | AuthBackendStsRoleState, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        if (opts && opts.id) {
            const state = argsOrState as AuthBackendStsRoleState | undefined;
            inputs["accountId"] = state ? state.accountId : undefined;
            inputs["backend"] = state ? state.backend : undefined;
            inputs["stsRole"] = state ? state.stsRole : undefined;
        } else {
            const args = argsOrState as AuthBackendStsRoleArgs | undefined;
            if ((!args || args.accountId === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'accountId'");
            }
            if ((!args || args.stsRole === undefined) && !(opts && opts.urn)) {
                throw new Error("Missing required property 'stsRole'");
            }
            inputs["accountId"] = args ? args.accountId : undefined;
            inputs["backend"] = args ? args.backend : undefined;
            inputs["stsRole"] = args ? args.stsRole : undefined;
        }
        if (!opts) {
            opts = {}
        }

        if (!opts.version) {
            opts.version = utilities.getVersion();
        }
        super(AuthBackendStsRole.__pulumiType, name, inputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering AuthBackendStsRole resources.
 */
export interface AuthBackendStsRoleState {
    /**
     * The AWS account ID to configure the STS role for.
     */
    readonly accountId?: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `accountId`.
     */
    readonly stsRole?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a AuthBackendStsRole resource.
 */
export interface AuthBackendStsRoleArgs {
    /**
     * The AWS account ID to configure the STS role for.
     */
    readonly accountId: pulumi.Input<string>;
    /**
     * The path the AWS auth backend being configured was
     * mounted at.  Defaults to `aws`.
     */
    readonly backend?: pulumi.Input<string>;
    /**
     * The STS role to assume when verifying requests made
     * by EC2 instances in the account specified by `accountId`.
     */
    readonly stsRole: pulumi.Input<string>;
}
