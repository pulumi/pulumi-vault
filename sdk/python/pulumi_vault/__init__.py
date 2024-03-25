# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .audit import *
from .audit_request_header import *
from .auth_backend import *
from .cert_auth_backend_role import *
from .egp_policy import *
from .get_auth_backend import *
from .get_auth_backends import *
from .get_nomad_access_token import *
from .get_policy_document import *
from .get_raft_autopilot_state import *
from .mfa_duo import *
from .mfa_okta import *
from .mfa_pingid import *
from .mfa_totp import *
from .mount import *
from .namespace import *
from .nomad_secret_backend import *
from .nomad_secret_role import *
from .password_policy import *
from .policy import *
from .provider import *
from .quota_lease_count import *
from .quota_rate_limit import *
from .raft_autopilot import *
from .raft_snapshot_agent_config import *
from .rgp_policy import *
from .token import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_vault.ad as __ad
    ad = __ad
    import pulumi_vault.alicloud as __alicloud
    alicloud = __alicloud
    import pulumi_vault.approle as __approle
    approle = __approle
    import pulumi_vault.aws as __aws
    aws = __aws
    import pulumi_vault.azure as __azure
    azure = __azure
    import pulumi_vault.config as __config
    config = __config
    import pulumi_vault.consul as __consul
    consul = __consul
    import pulumi_vault.database as __database
    database = __database
    import pulumi_vault.gcp as __gcp
    gcp = __gcp
    import pulumi_vault.generic as __generic
    generic = __generic
    import pulumi_vault.github as __github
    github = __github
    import pulumi_vault.identity as __identity
    identity = __identity
    import pulumi_vault.jwt as __jwt
    jwt = __jwt
    import pulumi_vault.kmip as __kmip
    kmip = __kmip
    import pulumi_vault.kubernetes as __kubernetes
    kubernetes = __kubernetes
    import pulumi_vault.kv as __kv
    kv = __kv
    import pulumi_vault.ldap as __ldap
    ldap = __ldap
    import pulumi_vault.managed as __managed
    managed = __managed
    import pulumi_vault.mongodbatlas as __mongodbatlas
    mongodbatlas = __mongodbatlas
    import pulumi_vault.okta as __okta
    okta = __okta
    import pulumi_vault.pkisecret as __pkisecret
    pkisecret = __pkisecret
    import pulumi_vault.rabbitmq as __rabbitmq
    rabbitmq = __rabbitmq
    import pulumi_vault.saml as __saml
    saml = __saml
    import pulumi_vault.secrets as __secrets
    secrets = __secrets
    import pulumi_vault.ssh as __ssh
    ssh = __ssh
    import pulumi_vault.terraformcloud as __terraformcloud
    terraformcloud = __terraformcloud
    import pulumi_vault.tokenauth as __tokenauth
    tokenauth = __tokenauth
    import pulumi_vault.transform as __transform
    transform = __transform
    import pulumi_vault.transit as __transit
    transit = __transit
else:
    ad = _utilities.lazy_import('pulumi_vault.ad')
    alicloud = _utilities.lazy_import('pulumi_vault.alicloud')
    approle = _utilities.lazy_import('pulumi_vault.approle')
    aws = _utilities.lazy_import('pulumi_vault.aws')
    azure = _utilities.lazy_import('pulumi_vault.azure')
    config = _utilities.lazy_import('pulumi_vault.config')
    consul = _utilities.lazy_import('pulumi_vault.consul')
    database = _utilities.lazy_import('pulumi_vault.database')
    gcp = _utilities.lazy_import('pulumi_vault.gcp')
    generic = _utilities.lazy_import('pulumi_vault.generic')
    github = _utilities.lazy_import('pulumi_vault.github')
    identity = _utilities.lazy_import('pulumi_vault.identity')
    jwt = _utilities.lazy_import('pulumi_vault.jwt')
    kmip = _utilities.lazy_import('pulumi_vault.kmip')
    kubernetes = _utilities.lazy_import('pulumi_vault.kubernetes')
    kv = _utilities.lazy_import('pulumi_vault.kv')
    ldap = _utilities.lazy_import('pulumi_vault.ldap')
    managed = _utilities.lazy_import('pulumi_vault.managed')
    mongodbatlas = _utilities.lazy_import('pulumi_vault.mongodbatlas')
    okta = _utilities.lazy_import('pulumi_vault.okta')
    pkisecret = _utilities.lazy_import('pulumi_vault.pkisecret')
    rabbitmq = _utilities.lazy_import('pulumi_vault.rabbitmq')
    saml = _utilities.lazy_import('pulumi_vault.saml')
    secrets = _utilities.lazy_import('pulumi_vault.secrets')
    ssh = _utilities.lazy_import('pulumi_vault.ssh')
    terraformcloud = _utilities.lazy_import('pulumi_vault.terraformcloud')
    tokenauth = _utilities.lazy_import('pulumi_vault.tokenauth')
    transform = _utilities.lazy_import('pulumi_vault.transform')
    transit = _utilities.lazy_import('pulumi_vault.transit')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "vault",
  "mod": "ad/secretBackend",
  "fqn": "pulumi_vault.ad",
  "classes": {
   "vault:ad/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "ad/secretLibrary",
  "fqn": "pulumi_vault.ad",
  "classes": {
   "vault:ad/secretLibrary:SecretLibrary": "SecretLibrary"
  }
 },
 {
  "pkg": "vault",
  "mod": "ad/secretRole",
  "fqn": "pulumi_vault.ad",
  "classes": {
   "vault:ad/secretRole:SecretRole": "SecretRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "alicloud/authBackendRole",
  "fqn": "pulumi_vault.alicloud",
  "classes": {
   "vault:alicloud/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "appRole/authBackendLogin",
  "fqn": "pulumi_vault.approle",
  "classes": {
   "vault:appRole/authBackendLogin:AuthBackendLogin": "AuthBackendLogin"
  }
 },
 {
  "pkg": "vault",
  "mod": "appRole/authBackendRole",
  "fqn": "pulumi_vault.approle",
  "classes": {
   "vault:appRole/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "appRole/authBackendRoleSecretId",
  "fqn": "pulumi_vault.approle",
  "classes": {
   "vault:appRole/authBackendRoleSecretId:AuthBackendRoleSecretId": "AuthBackendRoleSecretId"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendCert",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendCert:AuthBackendCert": "AuthBackendCert"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendClient",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendClient:AuthBackendClient": "AuthBackendClient"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendConfigIdentity",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendConfigIdentity:AuthBackendConfigIdentity": "AuthBackendConfigIdentity"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendIdentityWhitelist",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendIdentityWhitelist:AuthBackendIdentityWhitelist": "AuthBackendIdentityWhitelist"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendLogin",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendLogin:AuthBackendLogin": "AuthBackendLogin"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendRole",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendRoleTag",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendRoleTag:AuthBackendRoleTag": "AuthBackendRoleTag"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendRoletagBlacklist",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendRoletagBlacklist:AuthBackendRoletagBlacklist": "AuthBackendRoletagBlacklist"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/authBackendStsRole",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/authBackendStsRole:AuthBackendStsRole": "AuthBackendStsRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/secretBackend",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/secretBackendRole",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "aws/secretBackendStaticRole",
  "fqn": "pulumi_vault.aws",
  "classes": {
   "vault:aws/secretBackendStaticRole:SecretBackendStaticRole": "SecretBackendStaticRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "azure/authBackendConfig",
  "fqn": "pulumi_vault.azure",
  "classes": {
   "vault:azure/authBackendConfig:AuthBackendConfig": "AuthBackendConfig"
  }
 },
 {
  "pkg": "vault",
  "mod": "azure/authBackendRole",
  "fqn": "pulumi_vault.azure",
  "classes": {
   "vault:azure/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "azure/backend",
  "fqn": "pulumi_vault.azure",
  "classes": {
   "vault:azure/backend:Backend": "Backend"
  }
 },
 {
  "pkg": "vault",
  "mod": "azure/backendRole",
  "fqn": "pulumi_vault.azure",
  "classes": {
   "vault:azure/backendRole:BackendRole": "BackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "config/uiCustomMessage",
  "fqn": "pulumi_vault.config",
  "classes": {
   "vault:config/uiCustomMessage:UiCustomMessage": "UiCustomMessage"
  }
 },
 {
  "pkg": "vault",
  "mod": "consul/secretBackend",
  "fqn": "pulumi_vault.consul",
  "classes": {
   "vault:consul/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "consul/secretBackendRole",
  "fqn": "pulumi_vault.consul",
  "classes": {
   "vault:consul/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "database/secretBackendConnection",
  "fqn": "pulumi_vault.database",
  "classes": {
   "vault:database/secretBackendConnection:SecretBackendConnection": "SecretBackendConnection"
  }
 },
 {
  "pkg": "vault",
  "mod": "database/secretBackendRole",
  "fqn": "pulumi_vault.database",
  "classes": {
   "vault:database/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "database/secretBackendStaticRole",
  "fqn": "pulumi_vault.database",
  "classes": {
   "vault:database/secretBackendStaticRole:SecretBackendStaticRole": "SecretBackendStaticRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "database/secretsMount",
  "fqn": "pulumi_vault.database",
  "classes": {
   "vault:database/secretsMount:SecretsMount": "SecretsMount"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/authBackend",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/authBackendRole",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/secretBackend",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/secretImpersonatedAccount",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/secretImpersonatedAccount:SecretImpersonatedAccount": "SecretImpersonatedAccount"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/secretRoleset",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/secretRoleset:SecretRoleset": "SecretRoleset"
  }
 },
 {
  "pkg": "vault",
  "mod": "gcp/secretStaticAccount",
  "fqn": "pulumi_vault.gcp",
  "classes": {
   "vault:gcp/secretStaticAccount:SecretStaticAccount": "SecretStaticAccount"
  }
 },
 {
  "pkg": "vault",
  "mod": "generic/endpoint",
  "fqn": "pulumi_vault.generic",
  "classes": {
   "vault:generic/endpoint:Endpoint": "Endpoint"
  }
 },
 {
  "pkg": "vault",
  "mod": "generic/secret",
  "fqn": "pulumi_vault.generic",
  "classes": {
   "vault:generic/secret:Secret": "Secret"
  }
 },
 {
  "pkg": "vault",
  "mod": "github/authBackend",
  "fqn": "pulumi_vault.github",
  "classes": {
   "vault:github/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "github/team",
  "fqn": "pulumi_vault.github",
  "classes": {
   "vault:github/team:Team": "Team"
  }
 },
 {
  "pkg": "vault",
  "mod": "github/user",
  "fqn": "pulumi_vault.github",
  "classes": {
   "vault:github/user:User": "User"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/entity",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/entity:Entity": "Entity"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/entityAlias",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/entityAlias:EntityAlias": "EntityAlias"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/entityPolicies",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/entityPolicies:EntityPolicies": "EntityPolicies"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/group",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/group:Group": "Group"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/groupAlias",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/groupAlias:GroupAlias": "GroupAlias"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/groupMemberEntityIds",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/groupMemberEntityIds:GroupMemberEntityIds": "GroupMemberEntityIds"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/groupMemberGroupIds",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/groupMemberGroupIds:GroupMemberGroupIds": "GroupMemberGroupIds"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/groupPolicies",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/groupPolicies:GroupPolicies": "GroupPolicies"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/mfaDuo",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/mfaDuo:MfaDuo": "MfaDuo"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/mfaLoginEnforcement",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/mfaLoginEnforcement:MfaLoginEnforcement": "MfaLoginEnforcement"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/mfaOkta",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/mfaOkta:MfaOkta": "MfaOkta"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/mfaPingid",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/mfaPingid:MfaPingid": "MfaPingid"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/mfaTotp",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/mfaTotp:MfaTotp": "MfaTotp"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidc",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidc:Oidc": "Oidc"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcAssignment",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcAssignment:OidcAssignment": "OidcAssignment"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcClient",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcClient:OidcClient": "OidcClient"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcKey",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcKey:OidcKey": "OidcKey"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcKeyAllowedClientID",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcKeyAllowedClientID:OidcKeyAllowedClientID": "OidcKeyAllowedClientID"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcProvider",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcProvider:OidcProvider": "OidcProvider"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcRole",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcRole:OidcRole": "OidcRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "identity/oidcScope",
  "fqn": "pulumi_vault.identity",
  "classes": {
   "vault:identity/oidcScope:OidcScope": "OidcScope"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/audit",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/audit:Audit": "Audit"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/auditRequestHeader",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/auditRequestHeader:AuditRequestHeader": "AuditRequestHeader"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/authBackend",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/certAuthBackendRole",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/certAuthBackendRole:CertAuthBackendRole": "CertAuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/egpPolicy",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/egpPolicy:EgpPolicy": "EgpPolicy"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/mfaDuo",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/mfaDuo:MfaDuo": "MfaDuo"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/mfaOkta",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/mfaOkta:MfaOkta": "MfaOkta"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/mfaPingid",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/mfaPingid:MfaPingid": "MfaPingid"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/mfaTotp",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/mfaTotp:MfaTotp": "MfaTotp"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/mount",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/mount:Mount": "Mount"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/namespace",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/namespace:Namespace": "Namespace"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/nomadSecretBackend",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/nomadSecretBackend:NomadSecretBackend": "NomadSecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/nomadSecretRole",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/nomadSecretRole:NomadSecretRole": "NomadSecretRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/passwordPolicy",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/passwordPolicy:PasswordPolicy": "PasswordPolicy"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/policy",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/policy:Policy": "Policy"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/quotaLeaseCount",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/quotaLeaseCount:QuotaLeaseCount": "QuotaLeaseCount"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/quotaRateLimit",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/quotaRateLimit:QuotaRateLimit": "QuotaRateLimit"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/raftAutopilot",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/raftAutopilot:RaftAutopilot": "RaftAutopilot"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/raftSnapshotAgentConfig",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/raftSnapshotAgentConfig:RaftSnapshotAgentConfig": "RaftSnapshotAgentConfig"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/rgpPolicy",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/rgpPolicy:RgpPolicy": "RgpPolicy"
  }
 },
 {
  "pkg": "vault",
  "mod": "index/token",
  "fqn": "pulumi_vault",
  "classes": {
   "vault:index/token:Token": "Token"
  }
 },
 {
  "pkg": "vault",
  "mod": "jwt/authBackend",
  "fqn": "pulumi_vault.jwt",
  "classes": {
   "vault:jwt/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "jwt/authBackendRole",
  "fqn": "pulumi_vault.jwt",
  "classes": {
   "vault:jwt/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "kmip/secretBackend",
  "fqn": "pulumi_vault.kmip",
  "classes": {
   "vault:kmip/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "kmip/secretRole",
  "fqn": "pulumi_vault.kmip",
  "classes": {
   "vault:kmip/secretRole:SecretRole": "SecretRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "kmip/secretScope",
  "fqn": "pulumi_vault.kmip",
  "classes": {
   "vault:kmip/secretScope:SecretScope": "SecretScope"
  }
 },
 {
  "pkg": "vault",
  "mod": "kubernetes/authBackendConfig",
  "fqn": "pulumi_vault.kubernetes",
  "classes": {
   "vault:kubernetes/authBackendConfig:AuthBackendConfig": "AuthBackendConfig"
  }
 },
 {
  "pkg": "vault",
  "mod": "kubernetes/authBackendRole",
  "fqn": "pulumi_vault.kubernetes",
  "classes": {
   "vault:kubernetes/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "kubernetes/secretBackend",
  "fqn": "pulumi_vault.kubernetes",
  "classes": {
   "vault:kubernetes/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "kubernetes/secretBackendRole",
  "fqn": "pulumi_vault.kubernetes",
  "classes": {
   "vault:kubernetes/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "kv/secret",
  "fqn": "pulumi_vault.kv",
  "classes": {
   "vault:kv/secret:Secret": "Secret"
  }
 },
 {
  "pkg": "vault",
  "mod": "kv/secretBackendV2",
  "fqn": "pulumi_vault.kv",
  "classes": {
   "vault:kv/secretBackendV2:SecretBackendV2": "SecretBackendV2"
  }
 },
 {
  "pkg": "vault",
  "mod": "kv/secretV2",
  "fqn": "pulumi_vault.kv",
  "classes": {
   "vault:kv/secretV2:SecretV2": "SecretV2"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/authBackend",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/authBackendGroup",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/authBackendGroup:AuthBackendGroup": "AuthBackendGroup"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/authBackendUser",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/authBackendUser:AuthBackendUser": "AuthBackendUser"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/secretBackend",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/secretBackendDynamicRole",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/secretBackendDynamicRole:SecretBackendDynamicRole": "SecretBackendDynamicRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/secretBackendLibrarySet",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/secretBackendLibrarySet:SecretBackendLibrarySet": "SecretBackendLibrarySet"
  }
 },
 {
  "pkg": "vault",
  "mod": "ldap/secretBackendStaticRole",
  "fqn": "pulumi_vault.ldap",
  "classes": {
   "vault:ldap/secretBackendStaticRole:SecretBackendStaticRole": "SecretBackendStaticRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "managed/keys",
  "fqn": "pulumi_vault.managed",
  "classes": {
   "vault:managed/keys:Keys": "Keys"
  }
 },
 {
  "pkg": "vault",
  "mod": "mongodbatlas/secretBackend",
  "fqn": "pulumi_vault.mongodbatlas",
  "classes": {
   "vault:mongodbatlas/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "mongodbatlas/secretRole",
  "fqn": "pulumi_vault.mongodbatlas",
  "classes": {
   "vault:mongodbatlas/secretRole:SecretRole": "SecretRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "okta/authBackend",
  "fqn": "pulumi_vault.okta",
  "classes": {
   "vault:okta/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "okta/authBackendGroup",
  "fqn": "pulumi_vault.okta",
  "classes": {
   "vault:okta/authBackendGroup:AuthBackendGroup": "AuthBackendGroup"
  }
 },
 {
  "pkg": "vault",
  "mod": "okta/authBackendUser",
  "fqn": "pulumi_vault.okta",
  "classes": {
   "vault:okta/authBackendUser:AuthBackendUser": "AuthBackendUser"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/backendConfigCluster",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/backendConfigCluster:BackendConfigCluster": "BackendConfigCluster"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendCert",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendCert:SecretBackendCert": "SecretBackendCert"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendConfigCa",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendConfigCa:SecretBackendConfigCa": "SecretBackendConfigCa"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendConfigIssuers",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendConfigIssuers:SecretBackendConfigIssuers": "SecretBackendConfigIssuers"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendConfigUrls",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendConfigUrls:SecretBackendConfigUrls": "SecretBackendConfigUrls"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendCrlConfig",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendCrlConfig:SecretBackendCrlConfig": "SecretBackendCrlConfig"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendIntermediateCertRequest",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendIntermediateCertRequest:SecretBackendIntermediateCertRequest": "SecretBackendIntermediateCertRequest"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendIntermediateSetSigned",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendIntermediateSetSigned:SecretBackendIntermediateSetSigned": "SecretBackendIntermediateSetSigned"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendIssuer",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendIssuer:SecretBackendIssuer": "SecretBackendIssuer"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendKey",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendKey:SecretBackendKey": "SecretBackendKey"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendRole",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendRootCert",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendRootCert:SecretBackendRootCert": "SecretBackendRootCert"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendRootSignIntermediate",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendRootSignIntermediate:SecretBackendRootSignIntermediate": "SecretBackendRootSignIntermediate"
  }
 },
 {
  "pkg": "vault",
  "mod": "pkiSecret/secretBackendSign",
  "fqn": "pulumi_vault.pkisecret",
  "classes": {
   "vault:pkiSecret/secretBackendSign:SecretBackendSign": "SecretBackendSign"
  }
 },
 {
  "pkg": "vault",
  "mod": "rabbitMq/secretBackend",
  "fqn": "pulumi_vault.rabbitmq",
  "classes": {
   "vault:rabbitMq/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "rabbitMq/secretBackendRole",
  "fqn": "pulumi_vault.rabbitmq",
  "classes": {
   "vault:rabbitMq/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "saml/authBackend",
  "fqn": "pulumi_vault.saml",
  "classes": {
   "vault:saml/authBackend:AuthBackend": "AuthBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "saml/authBackendRole",
  "fqn": "pulumi_vault.saml",
  "classes": {
   "vault:saml/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncAssociation",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncAssociation:SyncAssociation": "SyncAssociation"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncAwsDestination",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncAwsDestination:SyncAwsDestination": "SyncAwsDestination"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncAzureDestination",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncAzureDestination:SyncAzureDestination": "SyncAzureDestination"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncConfig",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncConfig:SyncConfig": "SyncConfig"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncGcpDestination",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncGcpDestination:SyncGcpDestination": "SyncGcpDestination"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncGhDestination",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncGhDestination:SyncGhDestination": "SyncGhDestination"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncGithubApps",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncGithubApps:SyncGithubApps": "SyncGithubApps"
  }
 },
 {
  "pkg": "vault",
  "mod": "secrets/syncVercelDestination",
  "fqn": "pulumi_vault.secrets",
  "classes": {
   "vault:secrets/syncVercelDestination:SyncVercelDestination": "SyncVercelDestination"
  }
 },
 {
  "pkg": "vault",
  "mod": "ssh/secretBackendCa",
  "fqn": "pulumi_vault.ssh",
  "classes": {
   "vault:ssh/secretBackendCa:SecretBackendCa": "SecretBackendCa"
  }
 },
 {
  "pkg": "vault",
  "mod": "ssh/secretBackendRole",
  "fqn": "pulumi_vault.ssh",
  "classes": {
   "vault:ssh/secretBackendRole:SecretBackendRole": "SecretBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "terraformcloud/secretBackend",
  "fqn": "pulumi_vault.terraformcloud",
  "classes": {
   "vault:terraformcloud/secretBackend:SecretBackend": "SecretBackend"
  }
 },
 {
  "pkg": "vault",
  "mod": "terraformcloud/secretCreds",
  "fqn": "pulumi_vault.terraformcloud",
  "classes": {
   "vault:terraformcloud/secretCreds:SecretCreds": "SecretCreds"
  }
 },
 {
  "pkg": "vault",
  "mod": "terraformcloud/secretRole",
  "fqn": "pulumi_vault.terraformcloud",
  "classes": {
   "vault:terraformcloud/secretRole:SecretRole": "SecretRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "tokenauth/authBackendRole",
  "fqn": "pulumi_vault.tokenauth",
  "classes": {
   "vault:tokenauth/authBackendRole:AuthBackendRole": "AuthBackendRole"
  }
 },
 {
  "pkg": "vault",
  "mod": "transform/alphabet",
  "fqn": "pulumi_vault.transform",
  "classes": {
   "vault:transform/alphabet:Alphabet": "Alphabet"
  }
 },
 {
  "pkg": "vault",
  "mod": "transform/role",
  "fqn": "pulumi_vault.transform",
  "classes": {
   "vault:transform/role:Role": "Role"
  }
 },
 {
  "pkg": "vault",
  "mod": "transform/template",
  "fqn": "pulumi_vault.transform",
  "classes": {
   "vault:transform/template:Template": "Template"
  }
 },
 {
  "pkg": "vault",
  "mod": "transform/transformation",
  "fqn": "pulumi_vault.transform",
  "classes": {
   "vault:transform/transformation:Transformation": "Transformation"
  }
 },
 {
  "pkg": "vault",
  "mod": "transit/secretBackendKey",
  "fqn": "pulumi_vault.transit",
  "classes": {
   "vault:transit/secretBackendKey:SecretBackendKey": "SecretBackendKey"
  }
 },
 {
  "pkg": "vault",
  "mod": "transit/secretCacheConfig",
  "fqn": "pulumi_vault.transit",
  "classes": {
   "vault:transit/secretCacheConfig:SecretCacheConfig": "SecretCacheConfig"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "vault",
  "token": "pulumi:providers:vault",
  "fqn": "pulumi_vault",
  "class": "Provider"
 }
]
"""
)
