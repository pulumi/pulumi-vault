# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import json
import warnings
import pulumi
import pulumi.runtime
from typing import Union
from .. import utilities, tables

class SecretBackendStaticRole(pulumi.CustomResource):
    backend: pulumi.Output[str]
    """
    The unique name of the Vault mount to configure.
    """
    db_name: pulumi.Output[str]
    """
    The unique name of the database connection to use for the static role.
    """
    name: pulumi.Output[str]
    """
    A unique name to give the static role.
    """
    rotation_period: pulumi.Output[float]
    """
    The amount of time Vault should wait before rotating the password, in seconds.
    """
    rotation_statements: pulumi.Output[list]
    """
    Database statements to execute to rotate the password for the configured database user.
    """
    username: pulumi.Output[str]
    """
    The database username that this static role corresponds to.
    """
    def __init__(__self__, resource_name, opts=None, backend=None, db_name=None, name=None, rotation_period=None, rotation_statements=None, username=None, __props__=None, __name__=None, __opts__=None):
        """
        Creates a Database Secret Backend static role in Vault. Database secret backend
        static roles can be used to manage 1-to-1 mapping of a Vault Role to a user in a
        database for the database.
        
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of the Vault mount to configure.
        :param pulumi.Input[str] db_name: The unique name of the database connection to use for the static role.
        :param pulumi.Input[str] name: A unique name to give the static role.
        :param pulumi.Input[float] rotation_period: The amount of time Vault should wait before rotating the password, in seconds.
        :param pulumi.Input[list] rotation_statements: Database statements to execute to rotate the password for the configured database user.
        :param pulumi.Input[str] username: The database username that this static role corresponds to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/database_secret_backend_static_role.html.markdown.
        """
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
            opts.version = utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = dict()

            if backend is None:
                raise TypeError("Missing required property 'backend'")
            __props__['backend'] = backend
            if db_name is None:
                raise TypeError("Missing required property 'db_name'")
            __props__['db_name'] = db_name
            __props__['name'] = name
            if rotation_period is None:
                raise TypeError("Missing required property 'rotation_period'")
            __props__['rotation_period'] = rotation_period
            __props__['rotation_statements'] = rotation_statements
            if username is None:
                raise TypeError("Missing required property 'username'")
            __props__['username'] = username
        super(SecretBackendStaticRole, __self__).__init__(
            'vault:database/secretBackendStaticRole:SecretBackendStaticRole',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name, id, opts=None, backend=None, db_name=None, name=None, rotation_period=None, rotation_statements=None, username=None):
        """
        Get an existing SecretBackendStaticRole resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.
        
        :param str resource_name: The unique name of the resulting resource.
        :param str id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backend: The unique name of the Vault mount to configure.
        :param pulumi.Input[str] db_name: The unique name of the database connection to use for the static role.
        :param pulumi.Input[str] name: A unique name to give the static role.
        :param pulumi.Input[float] rotation_period: The amount of time Vault should wait before rotating the password, in seconds.
        :param pulumi.Input[list] rotation_statements: Database statements to execute to rotate the password for the configured database user.
        :param pulumi.Input[str] username: The database username that this static role corresponds to.

        > This content is derived from https://github.com/terraform-providers/terraform-provider-vault/blob/master/website/docs/r/database_secret_backend_static_role.html.markdown.
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
    def translate_output_property(self, prop):
        return tables._CAMEL_TO_SNAKE_CASE_TABLE.get(prop) or prop

    def translate_input_property(self, prop):
        return tables._SNAKE_TO_CAMEL_CASE_TABLE.get(prop) or prop
