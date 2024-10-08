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

__all__ = ['FrontendArgs', 'Frontend']

@pulumi.input_type
class FrontendArgs:
    def __init__(__self__, *,
                 path: pulumi.Input[pulumi.Archive],
                 project: pulumi.Input['_domain.ProjectArgs'],
                 publish: pulumi.Input[str],
                 build_commands: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[Sequence[pulumi.Input['_domain.EnvironmentVariableArgs']]]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Frontend resource.
        :param pulumi.Input[pulumi.Archive] path: The path to the frontend files.
        :param pulumi.Input['_domain.ProjectArgs'] project: The project to which the frontend will be deployed.
        :param pulumi.Input[str] publish: The folder in the path that contains the files to be published.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] build_commands: The commands to run before deploying the frontend.
        :param pulumi.Input[Sequence[pulumi.Input['_domain.EnvironmentVariableArgs']]] environment: The environment variables that will be set for the frontend.
        :param pulumi.Input[str] subdomain: The subdomain of the frontend.
        """
        pulumi.set(__self__, "path", path)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "publish", publish)
        if build_commands is not None:
            pulumi.set(__self__, "build_commands", build_commands)
        if environment is not None:
            pulumi.set(__self__, "environment", environment)
        if subdomain is not None:
            pulumi.set(__self__, "subdomain", subdomain)

    @property
    @pulumi.getter
    def path(self) -> pulumi.Input[pulumi.Archive]:
        """
        The path to the frontend files.
        """
        return pulumi.get(self, "path")

    @path.setter
    def path(self, value: pulumi.Input[pulumi.Archive]):
        pulumi.set(self, "path", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input['_domain.ProjectArgs']:
        """
        The project to which the frontend will be deployed.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input['_domain.ProjectArgs']):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def publish(self) -> pulumi.Input[str]:
        """
        The folder in the path that contains the files to be published.
        """
        return pulumi.get(self, "publish")

    @publish.setter
    def publish(self, value: pulumi.Input[str]):
        pulumi.set(self, "publish", value)

    @property
    @pulumi.getter(name="buildCommands")
    def build_commands(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        The commands to run before deploying the frontend.
        """
        return pulumi.get(self, "build_commands")

    @build_commands.setter
    def build_commands(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "build_commands", value)

    @property
    @pulumi.getter
    def environment(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['_domain.EnvironmentVariableArgs']]]]:
        """
        The environment variables that will be set for the frontend.
        """
        return pulumi.get(self, "environment")

    @environment.setter
    def environment(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['_domain.EnvironmentVariableArgs']]]]):
        pulumi.set(self, "environment", value)

    @property
    @pulumi.getter
    def subdomain(self) -> Optional[pulumi.Input[str]]:
        """
        The subdomain of the frontend.
        """
        return pulumi.get(self, "subdomain")

    @subdomain.setter
    def subdomain(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "subdomain", value)


class Frontend(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build_commands: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_domain.EnvironmentVariableArgs']]]]] = None,
                 path: Optional[pulumi.Input[pulumi.Archive]] = None,
                 project: Optional[pulumi.Input[pulumi.InputType['_domain.ProjectArgs']]] = None,
                 publish: Optional[pulumi.Input[str]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
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
        :param pulumi.Input[Sequence[pulumi.Input[str]]] build_commands: The commands to run before deploying the frontend.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_domain.EnvironmentVariableArgs']]]] environment: The environment variables that will be set for the frontend.
        :param pulumi.Input[pulumi.Archive] path: The path to the frontend files.
        :param pulumi.Input[pulumi.InputType['_domain.ProjectArgs']] project: The project to which the frontend will be deployed.
        :param pulumi.Input[str] publish: The folder in the path that contains the files to be published.
        :param pulumi.Input[str] subdomain: The subdomain of the frontend.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FrontendArgs,
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
        :param FrontendArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FrontendArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 build_commands: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 environment: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['_domain.EnvironmentVariableArgs']]]]] = None,
                 path: Optional[pulumi.Input[pulumi.Archive]] = None,
                 project: Optional[pulumi.Input[pulumi.InputType['_domain.ProjectArgs']]] = None,
                 publish: Optional[pulumi.Input[str]] = None,
                 subdomain: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FrontendArgs.__new__(FrontendArgs)

            __props__.__dict__["build_commands"] = build_commands
            __props__.__dict__["environment"] = environment
            if path is None and not opts.urn:
                raise TypeError("Missing required property 'path'")
            __props__.__dict__["path"] = path
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if publish is None and not opts.urn:
                raise TypeError("Missing required property 'publish'")
            __props__.__dict__["publish"] = publish
            __props__.__dict__["subdomain"] = subdomain
            __props__.__dict__["url"] = None
        super(Frontend, __self__).__init__(
            'genezio:index:Frontend',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Frontend':
        """
        Get an existing Frontend resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FrontendArgs.__new__(FrontendArgs)

        __props__.__dict__["build_commands"] = None
        __props__.__dict__["environment"] = None
        __props__.__dict__["path"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["publish"] = None
        __props__.__dict__["subdomain"] = None
        __props__.__dict__["url"] = None
        return Frontend(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="buildCommands")
    def build_commands(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        The commands to run before deploying the frontend.
        """
        return pulumi.get(self, "build_commands")

    @property
    @pulumi.getter
    def environment(self) -> pulumi.Output[Optional[Sequence['_domain.outputs.EnvironmentVariable']]]:
        """
        The environment variables that will be set for the frontend.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter
    def path(self) -> pulumi.Output[pulumi.Archive]:
        """
        The path to the frontend files.
        """
        return pulumi.get(self, "path")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output['_domain.outputs.Project']:
        """
        The project to which the frontend will be deployed.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def publish(self) -> pulumi.Output[str]:
        """
        The folder in the path that contains the files to be published.
        """
        return pulumi.get(self, "publish")

    @property
    @pulumi.getter
    def subdomain(self) -> pulumi.Output[Optional[str]]:
        """
        The subdomain of the frontend.
        """
        return pulumi.get(self, "subdomain")

    @property
    @pulumi.getter
    def url(self) -> pulumi.Output[str]:
        """
        The URL of the frontend.
        """
        return pulumi.get(self, "url")

