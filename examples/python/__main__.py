import pulumi
import pulumi_genezio as genezio

my_random_resource = genezio.Random("myRandomResource", length=24)
pulumi.export("output", {
    "value": my_random_resource.result,
})
