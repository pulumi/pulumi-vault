// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export { AuditArgs, AuditState } from "./audit";
export type Audit = import("./audit").Audit;
export const Audit: typeof import("./audit").Audit = null as any;
utilities.lazyLoad(exports, ["Audit"], () => require("./audit"));

export { AuditRequestHeaderArgs, AuditRequestHeaderState } from "./auditRequestHeader";
export type AuditRequestHeader = import("./auditRequestHeader").AuditRequestHeader;
export const AuditRequestHeader: typeof import("./auditRequestHeader").AuditRequestHeader = null as any;
utilities.lazyLoad(exports, ["AuditRequestHeader"], () => require("./auditRequestHeader"));

export { AuthBackendArgs, AuthBackendState } from "./authBackend";
export type AuthBackend = import("./authBackend").AuthBackend;
export const AuthBackend: typeof import("./authBackend").AuthBackend = null as any;
utilities.lazyLoad(exports, ["AuthBackend"], () => require("./authBackend"));

export { CertAuthBackendRoleArgs, CertAuthBackendRoleState } from "./certAuthBackendRole";
export type CertAuthBackendRole = import("./certAuthBackendRole").CertAuthBackendRole;
export const CertAuthBackendRole: typeof import("./certAuthBackendRole").CertAuthBackendRole = null as any;
utilities.lazyLoad(exports, ["CertAuthBackendRole"], () => require("./certAuthBackendRole"));

export { EgpPolicyArgs, EgpPolicyState } from "./egpPolicy";
export type EgpPolicy = import("./egpPolicy").EgpPolicy;
export const EgpPolicy: typeof import("./egpPolicy").EgpPolicy = null as any;
utilities.lazyLoad(exports, ["EgpPolicy"], () => require("./egpPolicy"));

export { GetAuthBackendArgs, GetAuthBackendResult, GetAuthBackendOutputArgs } from "./getAuthBackend";
export const getAuthBackend: typeof import("./getAuthBackend").getAuthBackend = null as any;
export const getAuthBackendOutput: typeof import("./getAuthBackend").getAuthBackendOutput = null as any;
utilities.lazyLoad(exports, ["getAuthBackend","getAuthBackendOutput"], () => require("./getAuthBackend"));

export { GetAuthBackendsArgs, GetAuthBackendsResult, GetAuthBackendsOutputArgs } from "./getAuthBackends";
export const getAuthBackends: typeof import("./getAuthBackends").getAuthBackends = null as any;
export const getAuthBackendsOutput: typeof import("./getAuthBackends").getAuthBackendsOutput = null as any;
utilities.lazyLoad(exports, ["getAuthBackends","getAuthBackendsOutput"], () => require("./getAuthBackends"));

export { GetNamespaceArgs, GetNamespaceResult, GetNamespaceOutputArgs } from "./getNamespace";
export const getNamespace: typeof import("./getNamespace").getNamespace = null as any;
export const getNamespaceOutput: typeof import("./getNamespace").getNamespaceOutput = null as any;
utilities.lazyLoad(exports, ["getNamespace","getNamespaceOutput"], () => require("./getNamespace"));

export { GetNamespacesArgs, GetNamespacesResult, GetNamespacesOutputArgs } from "./getNamespaces";
export const getNamespaces: typeof import("./getNamespaces").getNamespaces = null as any;
export const getNamespacesOutput: typeof import("./getNamespaces").getNamespacesOutput = null as any;
utilities.lazyLoad(exports, ["getNamespaces","getNamespacesOutput"], () => require("./getNamespaces"));

export { GetNomadAccessTokenArgs, GetNomadAccessTokenResult, GetNomadAccessTokenOutputArgs } from "./getNomadAccessToken";
export const getNomadAccessToken: typeof import("./getNomadAccessToken").getNomadAccessToken = null as any;
export const getNomadAccessTokenOutput: typeof import("./getNomadAccessToken").getNomadAccessTokenOutput = null as any;
utilities.lazyLoad(exports, ["getNomadAccessToken","getNomadAccessTokenOutput"], () => require("./getNomadAccessToken"));

export { GetPolicyDocumentArgs, GetPolicyDocumentResult, GetPolicyDocumentOutputArgs } from "./getPolicyDocument";
export const getPolicyDocument: typeof import("./getPolicyDocument").getPolicyDocument = null as any;
export const getPolicyDocumentOutput: typeof import("./getPolicyDocument").getPolicyDocumentOutput = null as any;
utilities.lazyLoad(exports, ["getPolicyDocument","getPolicyDocumentOutput"], () => require("./getPolicyDocument"));

export { GetRaftAutopilotStateArgs, GetRaftAutopilotStateResult, GetRaftAutopilotStateOutputArgs } from "./getRaftAutopilotState";
export const getRaftAutopilotState: typeof import("./getRaftAutopilotState").getRaftAutopilotState = null as any;
export const getRaftAutopilotStateOutput: typeof import("./getRaftAutopilotState").getRaftAutopilotStateOutput = null as any;
utilities.lazyLoad(exports, ["getRaftAutopilotState","getRaftAutopilotStateOutput"], () => require("./getRaftAutopilotState"));

export { MfaDuoArgs, MfaDuoState } from "./mfaDuo";
export type MfaDuo = import("./mfaDuo").MfaDuo;
export const MfaDuo: typeof import("./mfaDuo").MfaDuo = null as any;
utilities.lazyLoad(exports, ["MfaDuo"], () => require("./mfaDuo"));

export { MfaOktaArgs, MfaOktaState } from "./mfaOkta";
export type MfaOkta = import("./mfaOkta").MfaOkta;
export const MfaOkta: typeof import("./mfaOkta").MfaOkta = null as any;
utilities.lazyLoad(exports, ["MfaOkta"], () => require("./mfaOkta"));

export { MfaPingidArgs, MfaPingidState } from "./mfaPingid";
export type MfaPingid = import("./mfaPingid").MfaPingid;
export const MfaPingid: typeof import("./mfaPingid").MfaPingid = null as any;
utilities.lazyLoad(exports, ["MfaPingid"], () => require("./mfaPingid"));

export { MfaTotpArgs, MfaTotpState } from "./mfaTotp";
export type MfaTotp = import("./mfaTotp").MfaTotp;
export const MfaTotp: typeof import("./mfaTotp").MfaTotp = null as any;
utilities.lazyLoad(exports, ["MfaTotp"], () => require("./mfaTotp"));

export { MountArgs, MountState } from "./mount";
export type Mount = import("./mount").Mount;
export const Mount: typeof import("./mount").Mount = null as any;
utilities.lazyLoad(exports, ["Mount"], () => require("./mount"));

export { NamespaceArgs, NamespaceState } from "./namespace";
export type Namespace = import("./namespace").Namespace;
export const Namespace: typeof import("./namespace").Namespace = null as any;
utilities.lazyLoad(exports, ["Namespace"], () => require("./namespace"));

export { NomadSecretBackendArgs, NomadSecretBackendState } from "./nomadSecretBackend";
export type NomadSecretBackend = import("./nomadSecretBackend").NomadSecretBackend;
export const NomadSecretBackend: typeof import("./nomadSecretBackend").NomadSecretBackend = null as any;
utilities.lazyLoad(exports, ["NomadSecretBackend"], () => require("./nomadSecretBackend"));

export { NomadSecretRoleArgs, NomadSecretRoleState } from "./nomadSecretRole";
export type NomadSecretRole = import("./nomadSecretRole").NomadSecretRole;
export const NomadSecretRole: typeof import("./nomadSecretRole").NomadSecretRole = null as any;
utilities.lazyLoad(exports, ["NomadSecretRole"], () => require("./nomadSecretRole"));

export { PasswordPolicyArgs, PasswordPolicyState } from "./passwordPolicy";
export type PasswordPolicy = import("./passwordPolicy").PasswordPolicy;
export const PasswordPolicy: typeof import("./passwordPolicy").PasswordPolicy = null as any;
utilities.lazyLoad(exports, ["PasswordPolicy"], () => require("./passwordPolicy"));

export { PluginArgs, PluginState } from "./plugin";
export type Plugin = import("./plugin").Plugin;
export const Plugin: typeof import("./plugin").Plugin = null as any;
utilities.lazyLoad(exports, ["Plugin"], () => require("./plugin"));

export { PluginPinnedVersionArgs, PluginPinnedVersionState } from "./pluginPinnedVersion";
export type PluginPinnedVersion = import("./pluginPinnedVersion").PluginPinnedVersion;
export const PluginPinnedVersion: typeof import("./pluginPinnedVersion").PluginPinnedVersion = null as any;
utilities.lazyLoad(exports, ["PluginPinnedVersion"], () => require("./pluginPinnedVersion"));

export { PolicyArgs, PolicyState } from "./policy";
export type Policy = import("./policy").Policy;
export const Policy: typeof import("./policy").Policy = null as any;
utilities.lazyLoad(exports, ["Policy"], () => require("./policy"));

export { ProviderArgs } from "./provider";
export type Provider = import("./provider").Provider;
export const Provider: typeof import("./provider").Provider = null as any;
utilities.lazyLoad(exports, ["Provider"], () => require("./provider"));

export { QuotaLeaseCountArgs, QuotaLeaseCountState } from "./quotaLeaseCount";
export type QuotaLeaseCount = import("./quotaLeaseCount").QuotaLeaseCount;
export const QuotaLeaseCount: typeof import("./quotaLeaseCount").QuotaLeaseCount = null as any;
utilities.lazyLoad(exports, ["QuotaLeaseCount"], () => require("./quotaLeaseCount"));

export { QuotaRateLimitArgs, QuotaRateLimitState } from "./quotaRateLimit";
export type QuotaRateLimit = import("./quotaRateLimit").QuotaRateLimit;
export const QuotaRateLimit: typeof import("./quotaRateLimit").QuotaRateLimit = null as any;
utilities.lazyLoad(exports, ["QuotaRateLimit"], () => require("./quotaRateLimit"));

export { RaftAutopilotArgs, RaftAutopilotState } from "./raftAutopilot";
export type RaftAutopilot = import("./raftAutopilot").RaftAutopilot;
export const RaftAutopilot: typeof import("./raftAutopilot").RaftAutopilot = null as any;
utilities.lazyLoad(exports, ["RaftAutopilot"], () => require("./raftAutopilot"));

export { RaftSnapshotAgentConfigArgs, RaftSnapshotAgentConfigState } from "./raftSnapshotAgentConfig";
export type RaftSnapshotAgentConfig = import("./raftSnapshotAgentConfig").RaftSnapshotAgentConfig;
export const RaftSnapshotAgentConfig: typeof import("./raftSnapshotAgentConfig").RaftSnapshotAgentConfig = null as any;
utilities.lazyLoad(exports, ["RaftSnapshotAgentConfig"], () => require("./raftSnapshotAgentConfig"));

export { RgpPolicyArgs, RgpPolicyState } from "./rgpPolicy";
export type RgpPolicy = import("./rgpPolicy").RgpPolicy;
export const RgpPolicy: typeof import("./rgpPolicy").RgpPolicy = null as any;
utilities.lazyLoad(exports, ["RgpPolicy"], () => require("./rgpPolicy"));

export { TokenArgs, TokenState } from "./token";
export type Token = import("./token").Token;
export const Token: typeof import("./token").Token = null as any;
utilities.lazyLoad(exports, ["Token"], () => require("./token"));


// Export sub-modules:
import * as ad from "./ad";
import * as alicloud from "./alicloud";
import * as approle from "./approle";
import * as aws from "./aws";
import * as azure from "./azure";
import * as config from "./config";
import * as consul from "./consul";
import * as database from "./database";
import * as gcp from "./gcp";
import * as generic from "./generic";
import * as github from "./github";
import * as identity from "./identity";
import * as jwt from "./jwt";
import * as kmip from "./kmip";
import * as kubernetes from "./kubernetes";
import * as kv from "./kv";
import * as ldap from "./ldap";
import * as managed from "./managed";
import * as mongodbatlas from "./mongodbatlas";
import * as okta from "./okta";
import * as pkisecret from "./pkisecret";
import * as rabbitmq from "./rabbitmq";
import * as saml from "./saml";
import * as secrets from "./secrets";
import * as ssh from "./ssh";
import * as terraformcloud from "./terraformcloud";
import * as tokenauth from "./tokenauth";
import * as transform from "./transform";
import * as transit from "./transit";
import * as types from "./types";

export {
    ad,
    alicloud,
    approle,
    aws,
    azure,
    config,
    consul,
    database,
    gcp,
    generic,
    github,
    identity,
    jwt,
    kmip,
    kubernetes,
    kv,
    ldap,
    managed,
    mongodbatlas,
    okta,
    pkisecret,
    rabbitmq,
    saml,
    secrets,
    ssh,
    terraformcloud,
    tokenauth,
    transform,
    transit,
    types,
};

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "vault:index/audit:Audit":
                return new Audit(name, <any>undefined, { urn })
            case "vault:index/auditRequestHeader:AuditRequestHeader":
                return new AuditRequestHeader(name, <any>undefined, { urn })
            case "vault:index/authBackend:AuthBackend":
                return new AuthBackend(name, <any>undefined, { urn })
            case "vault:index/certAuthBackendRole:CertAuthBackendRole":
                return new CertAuthBackendRole(name, <any>undefined, { urn })
            case "vault:index/egpPolicy:EgpPolicy":
                return new EgpPolicy(name, <any>undefined, { urn })
            case "vault:index/mfaDuo:MfaDuo":
                return new MfaDuo(name, <any>undefined, { urn })
            case "vault:index/mfaOkta:MfaOkta":
                return new MfaOkta(name, <any>undefined, { urn })
            case "vault:index/mfaPingid:MfaPingid":
                return new MfaPingid(name, <any>undefined, { urn })
            case "vault:index/mfaTotp:MfaTotp":
                return new MfaTotp(name, <any>undefined, { urn })
            case "vault:index/mount:Mount":
                return new Mount(name, <any>undefined, { urn })
            case "vault:index/namespace:Namespace":
                return new Namespace(name, <any>undefined, { urn })
            case "vault:index/nomadSecretBackend:NomadSecretBackend":
                return new NomadSecretBackend(name, <any>undefined, { urn })
            case "vault:index/nomadSecretRole:NomadSecretRole":
                return new NomadSecretRole(name, <any>undefined, { urn })
            case "vault:index/passwordPolicy:PasswordPolicy":
                return new PasswordPolicy(name, <any>undefined, { urn })
            case "vault:index/plugin:Plugin":
                return new Plugin(name, <any>undefined, { urn })
            case "vault:index/pluginPinnedVersion:PluginPinnedVersion":
                return new PluginPinnedVersion(name, <any>undefined, { urn })
            case "vault:index/policy:Policy":
                return new Policy(name, <any>undefined, { urn })
            case "vault:index/quotaLeaseCount:QuotaLeaseCount":
                return new QuotaLeaseCount(name, <any>undefined, { urn })
            case "vault:index/quotaRateLimit:QuotaRateLimit":
                return new QuotaRateLimit(name, <any>undefined, { urn })
            case "vault:index/raftAutopilot:RaftAutopilot":
                return new RaftAutopilot(name, <any>undefined, { urn })
            case "vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig":
                return new RaftSnapshotAgentConfig(name, <any>undefined, { urn })
            case "vault:index/rgpPolicy:RgpPolicy":
                return new RgpPolicy(name, <any>undefined, { urn })
            case "vault:index/token:Token":
                return new Token(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("vault", "index/audit", _module)
pulumi.runtime.registerResourceModule("vault", "index/auditRequestHeader", _module)
pulumi.runtime.registerResourceModule("vault", "index/authBackend", _module)
pulumi.runtime.registerResourceModule("vault", "index/certAuthBackendRole", _module)
pulumi.runtime.registerResourceModule("vault", "index/egpPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/mfaDuo", _module)
pulumi.runtime.registerResourceModule("vault", "index/mfaOkta", _module)
pulumi.runtime.registerResourceModule("vault", "index/mfaPingid", _module)
pulumi.runtime.registerResourceModule("vault", "index/mfaTotp", _module)
pulumi.runtime.registerResourceModule("vault", "index/mount", _module)
pulumi.runtime.registerResourceModule("vault", "index/namespace", _module)
pulumi.runtime.registerResourceModule("vault", "index/nomadSecretBackend", _module)
pulumi.runtime.registerResourceModule("vault", "index/nomadSecretRole", _module)
pulumi.runtime.registerResourceModule("vault", "index/passwordPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/plugin", _module)
pulumi.runtime.registerResourceModule("vault", "index/pluginPinnedVersion", _module)
pulumi.runtime.registerResourceModule("vault", "index/policy", _module)
pulumi.runtime.registerResourceModule("vault", "index/quotaLeaseCount", _module)
pulumi.runtime.registerResourceModule("vault", "index/quotaRateLimit", _module)
pulumi.runtime.registerResourceModule("vault", "index/raftAutopilot", _module)
pulumi.runtime.registerResourceModule("vault", "index/raftSnapshotAgentConfig", _module)
pulumi.runtime.registerResourceModule("vault", "index/rgpPolicy", _module)
pulumi.runtime.registerResourceModule("vault", "index/token", _module)
pulumi.runtime.registerResourcePackage("vault", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:vault") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
