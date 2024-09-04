The Authentication Resource enables you to integrate authentication mechanisms into your project on the Genezio platform. This resource allows you to configure authentication providers, such as email-based authentication, and connect them to the appropriate database where user information is stored.

When creating an Authentication Resource, you can specify the project and database it will use, as well as define the authentication providers, such as email, that your application will support. This makes it easy to manage secure user authentication across your application infrastructure.

## Example Usage

```typescript
const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: database.url,
  provider: {
    email: true,
    web3: true,
    google: {
      id: "my-client-id",
      secret: "my-client",
    }
  },
});
```

## Pulumi Output Reference

Once the authentication resource is created, the `authId` is available as an output.

```typescript
const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: database.url,
  provider: {
    email: true,
    web3: true,
    google: {
      id: "my-client-id",
      secret: "my-client",
    }
  },
});

export const authId = auth.authId;
```
