---
title: Genezio Installation & Configuration
meta_desc: Provides an overview on how to configure credentials for the Pulumi Genezio Provider.
layout: installation
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).s

The Genezio provider is available as a package in the following languages:

* JavaScript/TypeScript: [`@pulumi/genezio`](https://www.npmjs.com/package/@pulumi/genezio)
* Python: [`pulumi-genezio`](https://pypi.org/project/pulumi-genezio/)
* Go: [`github.com/pulumi/pulumi-genezio/sdk/v3/go/genezio`](https://github.com/pulumi/pulumi-genezio)

## Credentials

The Pulumi Genezio Provider needs to be configured with a Genezio [`Personal Access Token`](https://app.genez.io/settings/tokens).

> If you don't have an `Personal Access Token`, you can create one [here](https://app.genez.io/settings/tokens).

Once you generated the `Personal Access Token` you can configure the Genezio provider using `pulumi config` command:

```bash
$ pulumi config set --secret genezio:authToken <token>
```

> Remember to pass `--secret` when setting `genezio:authToken` so it is properly encrypted.
