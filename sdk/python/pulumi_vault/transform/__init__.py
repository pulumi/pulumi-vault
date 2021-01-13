# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

# Export this package's modules as members:
from .alphabet import *
from .get_decode import *
from .get_encode import *
from .role import *
from .template import *
from .transformation import *

def _register_module():
    import pulumi
    from .. import _utilities


    class Module(pulumi.runtime.ResourceModule):
        _version = _utilities.get_semver_version()

        def version(self):
            return Module._version

        def construct(self, name: str, typ: str, urn: str) -> pulumi.Resource:
            if typ == "vault:transform/alphabet:Alphabet":
                return Alphabet(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "vault:transform/role:Role":
                return Role(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "vault:transform/template:Template":
                return Template(name, pulumi.ResourceOptions(urn=urn))
            elif typ == "vault:transform/transformation:Transformation":
                return Transformation(name, pulumi.ResourceOptions(urn=urn))
            else:
                raise Exception(f"unknown resource type {typ}")


    _module_instance = Module()
    pulumi.runtime.register_resource_module("vault", "transform/alphabet", _module_instance)
    pulumi.runtime.register_resource_module("vault", "transform/role", _module_instance)
    pulumi.runtime.register_resource_module("vault", "transform/template", _module_instance)
    pulumi.runtime.register_resource_module("vault", "transform/transformation", _module_instance)

_register_module()
