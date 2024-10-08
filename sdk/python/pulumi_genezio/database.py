# coding=utf-8
# *** WARNING: this file was generated by pulumi-language-python. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import domain as _domain

__all__ = ['DatabaseArgs', 'Database']

@pulumi.input_type
class DatabaseArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 project: Optional[pulumi.Input['_domain.ProjectArgs']] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Database resource.
        :param pulumi.Input[str] name: The name of the database to be deployed.
        :param pulumi.Input['_domain.ProjectArgs'] project: A database can be used in a project by linking it.
               	Linking the database will expose a connection URL as an environment variable for convenience.
               	The same database can be linked to multiple projects.
        :param pulumi.Input[str] region: The region in which the database will be deployed.
        :param pulumi.Input[str] type: The type of the database to be deployed.
               
               	Supported types are:
               	- postgres-neon
        """
        pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if region is None:
            region = 'us-east-1'
        if region is not None:
            pulumi.set(__self__, "region", region)
        if type is None:
            type = 'postgres-neon'
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        The name of the database to be deployed.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input['_domain.ProjectArgs']]:
        """
        A database can be used in a project by linking it.
        	Linking the database will expose a connection URL as an environment variable for convenience.
        	The same database can be linked to multiple projects.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input['_domain.ProjectArgs']]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> Optional[pulumi.Input[str]]:
        """
        The region in which the database will be deployed.
        """
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of the database to be deployed.

        	Supported types are:
        	- postgres-neon
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Database(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[pulumi.InputType['_domain.ProjectArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.

        The project resource will deploy an empty project on the Genezio platform.

        It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.

        ## Example Usage

        ### Basic Usage

        ### Environment Variables

        ## Pulumi Output Reference

        Once the project is created, the `projectId` and `projectUrl` are available as outputs.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] name: The name of the database to be deployed.
        :param pulumi.Input[pulumi.InputType['_domain.ProjectArgs']] project: A database can be used in a project by linking it.
               	Linking the database will expose a connection URL as an environment variable for convenience.
               	The same database can be linked to multiple projects.
        :param pulumi.Input[str] region: The region in which the database will be deployed.
        :param pulumi.Input[str] type: The type of the database to be deployed.
               
               	Supported types are:
               	- postgres-neon
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        A project resource that will be deployed on the Genezio platform.The project resource is used to group resources together and manage them as a single unit.

        The project resource will deploy an empty project on the Genezio platform.

        It is recommended to create a Project Resource as the first step in your deployment workflow. The output from this resource can then be utilized to provision and configure other resources within the project, ensuring they are properly associated and managed under a unified project.

        ## Example Usage

        ### Basic Usage

        ### Environment Variables

        ## Pulumi Output Reference

        Once the project is created, the `projectId` and `projectUrl` are available as outputs.

        :param str resource_name: The name of the resource.
        :param DatabaseArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[pulumi.InputType['_domain.ProjectArgs']]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseArgs.__new__(DatabaseArgs)

            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            if region is None:
                region = 'us-east-1'
            __props__.__dict__["region"] = region
            if type is None:
                type = 'postgres-neon'
            __props__.__dict__["type"] = type
            __props__.__dict__["database_id"] = None
            __props__.__dict__["url"] = None
        secret_opts = pulumi.ResourceOptions(additional_secret_outputs=["url"])
        opts = pulumi.ResourceOptions.merge(opts, secret_opts)
        super(Database, __self__).__init__(
            'genezio:index:Database',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Database':
        """
        Get an existing Database resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabaseArgs.__new__(DatabaseArgs)

        __props__.__dict__["database_id"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["region"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["url"] = None
        return Database(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Output[str]:
        """
        The database ID.
        """
        return pulumi.get(self, "database_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the database to be deployed.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[Optional['_domain.outputs.Project']]:
        """
        A database can be used in a project by linking it.
        	Linking the database will expose a connection URL as an environment variable for convenience.
        	The same database can be linked to multiple projects.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def region(self) -> pulumi.Output[Optional[str]]:
        """
        The region in which the database will be deployed.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[Optional[str]]:
        """
        The type of the database to be deployed.

        	Supported types are:
        	- postgres-neon
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the database.
        """
        return pulumi.get(self, "url")

