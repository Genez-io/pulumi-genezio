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

__all__ = ['ProviderArgs', 'Provider']

@pulumi.input_type
class ProviderArgs:
    def __init__(__self__, *,
                 auth_token: pulumi.Input[str],
                 stage: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Provider resource.
        """
        pulumi.set(__self__, "auth_token", auth_token)
        if stage is not None:
            pulumi.set(__self__, "stage", stage)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> pulumi.Input[str]:
        return pulumi.get(self, "auth_token")

    @auth_token.setter
    def auth_token(self, value: pulumi.Input[str]):
        pulumi.set(self, "auth_token", value)

    @property
    @pulumi.getter
    def stage(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "stage")

    @stage.setter
    def stage(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "stage", value)


class Provider(pulumi.ProviderResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a Genezio resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ProviderArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a Genezio resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ProviderArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ProviderArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 auth_token: Optional[pulumi.Input[str]] = None,
                 stage: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ProviderArgs.__new__(ProviderArgs)

            if auth_token is None and not opts.urn:
                raise TypeError("Missing required property 'auth_token'")
            __props__.__dict__["auth_token"] = auth_token
            __props__.__dict__["stage"] = stage
        super(Provider, __self__).__init__(
            'genezio',
            resource_name,
            __props__,
            opts)

    @property
    @pulumi.getter(name="authToken")
    def auth_token(self) -> pulumi.Output[str]:
        return pulumi.get(self, "auth_token")

    @property
    @pulumi.getter
    def stage(self) -> pulumi.Output[Optional[str]]:
        return pulumi.get(self, "stage")

