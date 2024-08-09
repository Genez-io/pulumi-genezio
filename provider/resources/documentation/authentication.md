An authentication resource that will be deployed to the genezio platform.

This resource requires an existing project to be deployed. The authentication resource will be used to provide out of the box authetication to your genezio project. Right now, the authentication feature works if you use genezio with typesafe classes. More info about the genezio authetication feature can be found here: https://genezio.com/docs/genezio-typesafe/authentication/

The authentication resource will require a database url to be passed in. This database url will be used to store the user information. By default, database type is `postgresql`. However, you can also use `mongodb` as the database type. The database url should work with the database type you choose.

## Example Usage

### Using PostgreSQL

This example creates a new project and an authentication resource. The authentication resource will use a postgresql database to store the user information.

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: "postgresql://username:password@host:port/database",
});
```

### Using MongoDB

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseType: "mongodb",
  databaseUrl: "mongodb://username:password@host:port/database",
});
```

### Using Providers

To actually use the authentication feature, you will need to specify what providers you want to use. The authentication resource supports the following providers:

- Email
- Google
- web3

#### Using The Email Provider

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: "postgresql://username:password@host:port/database",
  provider: {
    email: true,
  },
});
```

#### Using The Web3 Provider

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: "postgresql://username:password@host:port/database",
  provider: {
    web3: true,
  },
});
```

#### Using The Google Provider

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: "postgresql://username:password@host:port/database",
  provider: {
    google: {
      id: "google-client-id",
      secret: "google-client-secret",
    },
  },
});
```

#### Using Multiple Providers

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const auth = new genezio.Authentication("MyAuth", {
  project: {
    name: project.name,
    region: project.region,
  },
  databaseUrl: "postgresql://username:password@host:port/database",
  provider: {
    email: true,
    web3: true,
    google: {
      id: "google-client-id",
      secret: "google-client-secret",
    },
  },
});
```

### Use Databases to get your database url

```typescript
const project = new genezio.Project("MyProject", {
  name: "my-project",
  region: "us-east-1",
});

const database = new genezio.Database("MyDatabase", {
  project: {
    name: project.name,
    region: project.region,
  },
  name: "my-database",
});

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
      id: "google-client-id",
      secret: "google-client-secret",
    },
  },
});
```
