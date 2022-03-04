// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

export interface AuthBackendTune {
    /**
     * List of headers to whitelist and allowing
     * a plugin to include them in the response.
     */
    allowedResponseHeaders?: string[];
    /**
     * Specifies the list of keys that will
     * not be HMAC'd by audit devices in the request data object.
     */
    auditNonHmacRequestKeys?: string[];
    /**
     * Specifies the list of keys that will
     * not be HMAC'd by audit devices in the response data object.
     */
    auditNonHmacResponseKeys?: string[];
    /**
     * Specifies the default time-to-live.
     * If set, this overrides the global default.
     * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
     */
    defaultLeaseTtl?: string;
    /**
     * Specifies whether to show this mount in
     * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
     */
    listingVisibility?: string;
    /**
     * Specifies the maximum time-to-live.
     * If set, this overrides the global default.
     * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
     */
    maxLeaseTtl?: string;
    /**
     * List of headers to whitelist and
     * pass from the request to the backend.
     */
    passthroughRequestHeaders?: string[];
    /**
     * Specifies the type of tokens that should be returned by
     * the mount. Valid values are "default-service", "default-batch", "service", "batch".
     */
    tokenType?: string;
}

export interface GetPolicyDocumentRule {
    /**
     * Whitelists a list of keys and values that are permitted on the given path. See Parameters below.
     */
    allowedParameters?: outputs.GetPolicyDocumentRuleAllowedParameter[];
    /**
     * A list of capabilities that this rule apply to `path`. For example, ["read", "write"].
     */
    capabilities: string[];
    /**
     * Blacklists a list of parameter and values. Any values specified here take precedence over `allowedParameter`. See Parameters below.
     */
    deniedParameters?: outputs.GetPolicyDocumentRuleDeniedParameter[];
    /**
     * Description of the rule. Will be added as a comment to rendered rule.
     */
    description?: string;
    /**
     * The maximum allowed TTL that clients can specify for a wrapped response.
     */
    maxWrappingTtl?: string;
    /**
     * The minimum allowed TTL that clients can specify for a wrapped response.
     */
    minWrappingTtl?: string;
    /**
     * A path in Vault that this rule applies to.
     */
    path: string;
    /**
     * A list of parameters that must be specified.
     */
    requiredParameters?: string[];
}

export interface GetPolicyDocumentRuleAllowedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export interface GetPolicyDocumentRuleDeniedParameter {
    /**
     * name of permitted or denied parameter.
     */
    key: string;
    /**
     * list of values what are permitted or denied by policy rule.
     */
    values: string[];
}

export namespace azure {
    export interface BackendRoleAzureGroup {
        groupName: string;
        objectId: string;
    }

    export interface BackendRoleAzureRole {
        roleId: string;
        roleName: string;
        scope: string;
    }

}

export namespace config {
    export interface AuthLogins {
        method?: string;
        namespace?: string;
        parameters?: {[key: string]: string};
        path: string;
    }

    export interface ClientAuths {
        certFile: string;
        keyFile: string;
    }

    export interface Headers {
        name: string;
        value: string;
    }

}

export namespace database {
    export interface SecretBackendConnectionCassandra {
        /**
         * The number of seconds to use as a connection
         * timeout.
         */
        connectTimeout?: number;
        /**
         * A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
         */
        hosts?: string[];
        /**
         * Whether to skip verification of the server
         * certificate when using TLS.
         */
        insecureTls?: boolean;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * Concatenated PEM blocks configuring the certificate
         * chain.
         */
        pemBundle?: string;
        /**
         * A JSON structure configuring the certificate chain.
         */
        pemJson?: string;
        /**
         * The default port to connect to if no port is specified as
         * part of the host.
         */
        port?: number;
        /**
         * The CQL protocol version to use.
         */
        protocolVersion?: number;
        /**
         * Whether to use TLS when connecting to Cassandra.
         */
        tls?: boolean;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
    }

    export interface SecretBackendConnectionCouchbase {
        /**
         * Required if `tls` is `true`. Specifies the certificate authority of the Couchbase server, as a PEM certificate that has been base64 encoded.
         */
        base64Pem?: string;
        /**
         * Required for Couchbase versions prior to 6.5.0. This is only used to verify vault's connection to the server.
         */
        bucketName?: string;
        /**
         * A set of Couchbase URIs to connect to. Must use `couchbases://` scheme if `tls` is `true`.
         */
        hosts: string[];
        /**
         * Whether to skip verification of the server
         * certificate when using TLS.
         */
        insecureTls?: boolean;
        /**
         * The root credential password used in the connection URL.
         */
        password: string;
        /**
         * Whether to use TLS when connecting to Cassandra.
         */
        tls?: boolean;
        /**
         * The root credential username used in the connection URL.
         */
        username: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionElasticsearch {
        /**
         * The root credential password used in the connection URL.
         */
        password: string;
        /**
         * The URL for Elasticsearch's API. https requires certificate
         * by trusted CA if used.
         */
        url: string;
        /**
         * The root credential username used in the connection URL.
         */
        username: string;
    }

    export interface SecretBackendConnectionHana {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
    }

    export interface SecretBackendConnectionInfluxdb {
        /**
         * The number of seconds to use as a connection
         * timeout.
         */
        connectTimeout?: number;
        /**
         * The host to connect to.
         */
        host: string;
        /**
         * Whether to skip verification of the server
         * certificate when using TLS.
         */
        insecureTls?: boolean;
        /**
         * The root credential password used in the connection URL.
         */
        password: string;
        /**
         * Concatenated PEM blocks configuring the certificate
         * chain.
         */
        pemBundle?: string;
        /**
         * A JSON structure configuring the certificate chain.
         */
        pemJson?: string;
        /**
         * The default port to connect to if no port is specified as
         * part of the host.
         */
        port?: number;
        /**
         * Whether to use TLS when connecting to Cassandra.
         */
        tls?: boolean;
        /**
         * The root credential username used in the connection URL.
         */
        username: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMongodb {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMongodbatlas {
        /**
         * The Private Programmatic API Key used to connect with MongoDB Atlas API.
         */
        privateKey: string;
        /**
         * The Project ID the Database User should be created within.
         */
        projectId: string;
        /**
         * The Public Programmatic API Key used to authenticate with the MongoDB Atlas API.
         */
        publicKey: string;
    }

    export interface SecretBackendConnectionMssql {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * For Vault v1.9+. Set to true when the target is a
         * Contained Database, e.g. AzureSQL.
         * See the [Vault
         * docs](https://www.vaultproject.io/api/secret/databases/mssql#contained_db)
         */
        containedDb?: boolean;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMysql {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * x509 CA file for validating the certificate presented by the MySQL server. Must be PEM encoded.
         */
        tlsCa?: string;
        /**
         * x509 certificate for connecting to the database. This must be a PEM encoded version of the private key and the certificate combined.
         */
        tlsCertificateKey?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMysqlAurora {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMysqlLegacy {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionMysqlRds {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionOracle {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionPostgresql {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionRedshift {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

    export interface SecretBackendConnectionSnowflake {
        /**
         * Specifies the Redshift DSN. See
         * the [Vault
         * docs](https://www.vaultproject.io/api-docs/secret/databases/redshift#sample-payload)
         * for an example.
         */
        connectionUrl?: string;
        /**
         * The maximum amount of time a connection may be reused.
         */
        maxConnectionLifetime?: number;
        /**
         * The maximum number of idle connections to
         * the database.
         */
        maxIdleConnections?: number;
        /**
         * The maximum number of open connections to
         * the database.
         */
        maxOpenConnections?: number;
        /**
         * The root credential password used in the connection URL.
         */
        password?: string;
        /**
         * The root credential username used in the connection URL.
         */
        username?: string;
        /**
         * - [Template](https://www.vaultproject.io/docs/concepts/username-templating) describing how dynamic usernames are generated.
         */
        usernameTemplate?: string;
    }

}

export namespace gcp {
    export interface SecretRolesetBinding {
        /**
         * Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#roleset-bindings).
         */
        resource: string;
        /**
         * List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         */
        roles: string[];
    }

    export interface SecretStaticAccountBinding {
        /**
         * Resource or resource path for which IAM policy information will be bound. The resource path may be specified in a few different [formats](https://www.vaultproject.io/docs/secrets/gcp/index.html#bindings).
         */
        resource: string;
        /**
         * List of [GCP IAM roles](https://cloud.google.com/iam/docs/understanding-roles) for the resource.
         */
        roles: string[];
    }

}

export namespace github {
    export interface AuthBackendTune {
        /**
         * List of headers to whitelist and allowing
         * a plugin to include them in the response.
         */
        allowedResponseHeaders?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the request data object.
         */
        auditNonHmacRequestKeys?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the response data object.
         */
        auditNonHmacResponseKeys?: string[];
        /**
         * Specifies the default time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        defaultLeaseTtl?: string;
        /**
         * Specifies whether to show this mount in
         * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
         */
        listingVisibility?: string;
        /**
         * Specifies the maximum time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        maxLeaseTtl?: string;
        /**
         * List of headers to whitelist and
         * pass from the request to the backend.
         */
        passthroughRequestHeaders?: string[];
        /**
         * Specifies the type of tokens that should be returned by
         * the mount. Valid values are "default-service", "default-batch", "service", "batch".
         */
        tokenType?: string;
    }

}

export namespace identity {
    export interface GetEntityAlias {
        /**
         * Canonical ID of the Alias
         */
        canonicalId: string;
        /**
         * Creation time of the Alias
         */
        creationTime: string;
        /**
         * ID of the alias
         */
        id: string;
        /**
         * Last update time of the alias
         */
        lastUpdateTime: string;
        /**
         * List of canonical IDs merged with this alias
         */
        mergedFromCanonicalIds: string[];
        /**
         * Arbitrary metadata
         */
        metadata: {[key: string]: any};
        /**
         * Authentication mount acccessor which this alias belongs to
         */
        mountAccessor: string;
        /**
         * Authentication mount path which this alias belongs to
         */
        mountPath: string;
        /**
         * Authentication mount type which this alias belongs to
         */
        mountType: string;
        /**
         * Name of the alias
         */
        name: string;
    }

}

export namespace jwt {
    export interface AuthBackendTune {
        /**
         * List of headers to whitelist and allowing
         * a plugin to include them in the response.
         */
        allowedResponseHeaders?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the request data object.
         */
        auditNonHmacRequestKeys?: string[];
        /**
         * Specifies the list of keys that will
         * not be HMAC'd by audit devices in the response data object.
         */
        auditNonHmacResponseKeys?: string[];
        /**
         * Specifies the default time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        defaultLeaseTtl?: string;
        /**
         * Specifies whether to show this mount in
         * the UI-specific listing endpoint. Valid values are "unauth" or "hidden".
         */
        listingVisibility?: string;
        /**
         * Specifies the maximum time-to-live.
         * If set, this overrides the global default.
         * Must be a valid [duration string](https://golang.org/pkg/time/#ParseDuration)
         */
        maxLeaseTtl?: string;
        /**
         * List of headers to whitelist and
         * pass from the request to the backend.
         */
        passthroughRequestHeaders?: string[];
        /**
         * Specifies the type of tokens that should be returned by
         * the mount. Valid values are "default-service", "default-batch", "service", "batch".
         */
        tokenType?: string;
    }

}

export namespace okta {
    export interface AuthBackendGroup {
        /**
         * Name of the group within the Okta
         */
        groupName: string;
        /**
         * List of Vault policies to associate with this user
         */
        policies: string[];
    }

    export interface AuthBackendUser {
        /**
         * List of Okta groups to associate with this user
         */
        groups: string[];
        /**
         * List of Vault policies to associate with this user
         */
        policies?: string[];
        /**
         * Name of the user within Okta
         */
        username: string;
    }

}

export namespace rabbitMq {
    export interface SecretBackendRoleVhost {
        configure: string;
        host: string;
        read: string;
        write: string;
    }

    export interface SecretBackendRoleVhostTopic {
        host: string;
        /**
         * Specifies a map of virtual hosts to permissions.
         */
        vhosts?: outputs.rabbitMq.SecretBackendRoleVhostTopicVhost[];
    }

    export interface SecretBackendRoleVhostTopicVhost {
        read: string;
        topic: string;
        write: string;
    }

}
