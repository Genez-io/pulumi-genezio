using System.Collections.Generic;
using System.Linq;
using Pulumi;
using Genezio = Pulumi.Genezio;

return await Deployment.RunAsync(() => 
{
    var myRandomResource = new Genezio.Random("myRandomResource", new()
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

