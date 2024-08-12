import pulumi
import pulumi_genezio as genezio

my_project = genezio.Project("myProject", region="us-east-1", name="my-project")