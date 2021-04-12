# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities, _tables

__all__ = ['SecretBackendStaticRoleArgs', 'SecretBackendStaticRole']

@pulumi.input_type
class SecretBackendStaticRoleArgs:
    def __init__(__self__, *,
                 backend: pulumi.Input[str],
                 db_name: pulumi.Input[str],
                 rotation_period: pulumi.Input[int],
                 username: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_statements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a SecretBackendStaticRole resource.
        :param pulumi.Input[str] backend: The unique name of the Vault mount to configure.
        :param pulumi.Input[str] db_name: The unique name of the database connection to use for the static role.
        :param pulumi.Input[int] rotation_period: The amount of time Vault should wait before rotating the password, in seconds.
        :param pulumi.Input[str] username: The database username that this static role corresponds to.
        :param pulumi.Input[str] name: A unique name to give the static role.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rotation_statements: Database statements to execute to rotate the password for the configured database user.
        """
        pulumi.set(__self__, "backend", backend)
        pulumi.set(__self__, "db_name", db_name)
        pulumi.set(__self__, "rotation_period", rotation_period)
        pulumi.set(__self__, "username", username)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rotation_statements is not None:
            pulumi.set(__self__, "rotation_statements", rotation_statements)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Input[str]:
        """
        The unique name of the Vault mount to configure.
        """
        return pulumi.get(self, "backend")

    @backend.setter
    def backend(self, value: pulumi.Input[str]):
        pulumi.set(self, "backend", value)

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Input[str]:
        """
        The unique name of the database connection to use for the static role.
        """
        return pulumi.get(self, "db_name")

    @db_name.setter
    def db_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "db_name", value)

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Input[int]:
        """
        The amount of time Vault should wait before rotating the password, in seconds.
        """
        return pulumi.get(self, "rotation_period")

    @rotation_period.setter
    def rotation_period(self, value: pulumi.Input[int]):
        pulumi.set(self, "rotation_period", value)

    @property
    @pulumi.getter
    def username(self) -> pulumi.Input[str]:
        """
        The database username that this static role corresponds to.
        """
        return pulumi.get(self, "username")

    @username.setter
    def username(self, value: pulumi.Input[str]):
        pulumi.set(self, "username", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name to give the static role.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rotationStatements")
    def rotation_statements(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        Database statements to execute to rotate the password for the configured database user.
        """
        return pulumi.get(self, "rotation_statements")

    @rotation_statements.setter
    def rotation_statements(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "rotation_statements", value)


class SecretBackendStaticRole(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[int]] = None,
                 rotation_statements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        """
        Creates a Database Secret Backend static role in Vault. Database secret backend
        static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
        database for the database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        db = vault.Mount("db",
            path="postgres",
            type="database")
        postgres = vault.database.SecretBackendConnection("postgres",
            allowed_roles=["*"],
            backend=db.path,
            postgresql=vault.database.SecretBackendConnectionPostgresqlArgs(
                connection_url="postgres://username:password@host:port/database",
            ))
        static_role = vault.database.SecretBackendStaticRole("staticRole",
            backend=db.path,
            db_name=postgres.name,
            rotation_period=3600,
            rotation_statements=["ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';"],
            username="example")
        ```

        ## Import

        Database secret backend static roles can be imported using the `backend`, `/static-roles/`, and the `name` e.g.

        ```sh
         $ pulumi import vault:database/secretBackendStaticRole:SecretBackendStaticRole example postgres/static-roles/my-role
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of the Vault mount to configure.
        :param pulumi.Input[str] db_name: The unique name of the database connection to use for the static role.
        :param pulumi.Input[str] name: A unique name to give the static role.
        :param pulumi.Input[int] rotation_period: The amount of time Vault should wait before rotating the password, in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rotation_statements: Database statements to execute to rotate the password for the configured database user.
        :param pulumi.Input[str] username: The database username that this static role corresponds to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SecretBackendStaticRoleArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a Database Secret Backend static role in Vault. Database secret backend
        static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
        database for the database.

        ## Example Usage

        ```python
        import pulumi
        import pulumi_vault as vault

        db = vault.Mount("db",
            path="postgres",
            type="database")
        postgres = vault.database.SecretBackendConnection("postgres",
            allowed_roles=["*"],
            backend=db.path,
            postgresql=vault.database.SecretBackendConnectionPostgresqlArgs(
                connection_url="postgres://username:password@host:port/database",
            ))
        static_role = vault.database.SecretBackendStaticRole("staticRole",
            backend=db.path,
            db_name=postgres.name,
            rotation_period=3600,
            rotation_statements=["ALTER USER \"{{name}}\" WITH PASSWORD '{{password}}';"],
            username="example")
        ```

        ## Import

        Database secret backend static roles can be imported using the `backend`, `/static-roles/`, and the `name` e.g.

        ```sh
         $ pulumi import vault:database/secretBackendStaticRole:SecretBackendStaticRole example postgres/static-roles/my-role
        ```

        :param str resource_name: The name of the resource.
        :param SecretBackendStaticRoleArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SecretBackendStaticRoleArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backend: Optional[pulumi.Input[str]] = None,
                 db_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rotation_period: Optional[pulumi.Input[int]] = None,
                 rotation_statements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 username: Optional[pulumi.Input[str]] = None,
                 __props__=None,
                 __name__=None,
                 __opts__=None):
        if __name__ is not None:
            warnings.warn("explicit use of __name__ is deprecated", DeprecationWarning)
            resource_name = __name__
        if __opts__ is not None:
            warnings.warn("explicit use of __opts__ is deprecated, use 'opts' instead", DeprecationWarning)
            opts = __opts__
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if backend is None and not opts.urn:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if db_name is None and not opts.urn:
                raise TypeError("Missing required property 'db_name'")
            __props__['db_name'] = db_name
            __props__['name'] = name
            if rotation_period is None and not opts.urn:
                raise TypeError("Missing required property 'rotation_period'")
            __props__['rotation_period'] = rotation_period
            __props__['rotation_statements'] = rotation_statements
            if username is None and not opts.urn:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(SecretBackendStaticRole, __self__).__init__(
            'vault:database/secretBackendStaticRole:SecretBackendStaticRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            backend: Optional[pulumi.Input[str]] = None,
            db_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rotation_period: Optional[pulumi.Input[int]] = None,
            rotation_statements: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            username: Optional[pulumi.Input[str]] = None) -> 'SecretBackendStaticRole':
        """
        Get an existing SecretBackendStaticRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of the Vault mount to configure.
        :param pulumi.Input[str] db_name: The unique name of the database connection to use for the static role.
        :param pulumi.Input[str] name: A unique name to give the static role.
        :param pulumi.Input[int] rotation_period: The amount of time Vault should wait before rotating the password, in seconds.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] rotation_statements: Database statements to execute to rotate the password for the configured database user.
        :param pulumi.Input[str] username: The database username that this static role corresponds to.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = dict()

        __props__["backend"] = backend
        __props__["db_name"] = db_name
        __props__["name"] = name
        __props__["rotation_period"] = rotation_period
        __props__["rotation_statements"] = rotation_statements
        __props__["username"] = username
        return SecretBackendStaticRole(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def backend(self) -> pulumi.Output[str]:
        """
        The unique name of the Vault mount to configure.
        """
        return pulumi.get(self, "backend")

    @property
    @pulumi.getter(name="dbName")
    def db_name(self) -> pulumi.Output[str]:
        """
        The unique name of the database connection to use for the static role.
        """
        return pulumi.get(self, "db_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name to give the static role.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rotationPeriod")
    def rotation_period(self) -> pulumi.Output[int]:
        """
        The amount of time Vault should wait before rotating the password, in seconds.
        """
        return pulumi.get(self, "rotation_period")

    @property
    @pulumi.getter(name="rotationStatements")
    def rotation_statements(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        Database statements to execute to rotate the password for the configured database user.
        """
        return pulumi.get(self, "rotation_statements")

    @property
    @pulumi.getter
    def username(self) -> pulumi.Output[str]:
        """
        The database username that this static role corresponds to.
        """
        return pulumi.get(self, "username")

    def translate_output_property(self, prop):
        return _tables.CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return _tables.SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop

