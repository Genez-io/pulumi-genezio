using System.Collections.Generic;
using System.Linq;
using Pulumi;
using genezio = Pulumi.genezio;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new genezio.Random("myRandomResource", new()
    {
        Length = 24,
    });

    return new Dictionary<string, object?>
    {
        ["output"] = 
        {
            { "value", myRandomResource.Result },
        },
    };
});

