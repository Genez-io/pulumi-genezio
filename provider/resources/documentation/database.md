The Database Resource allows you to define and deploy databases on the Genezio platform. This resource provides a way to manage databases as part of your infrastructure, making it easier to store and retrieve application data within your project.

Databases are associated with a Project Resource, ensuring that they are organized and managed within the same project environment. When configuring a Database Resource, you specify attributes like the database name and the region in which it should be deployed, aligning it with your project's needs and geographic requirements.

## Example Usage

```typescript
import * as genezio from "@pulumi/genezio";

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "my-database",
  region: "us-east-1",
});
```

## Pulumi Output Reference

Once the database is created, the `url` and `databaseId` are available as outputs.

```typescript
const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "my-database",
  region: "us-east-1",
});

export const databaseUrl = database.url;
export const databaseId = database.databaseId;
```
