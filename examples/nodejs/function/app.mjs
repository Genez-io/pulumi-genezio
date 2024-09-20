export const handler = async (event) => {
  const name = event.queryStringParameters?.name || "World";
  console.log("The environment variable is:", process.env.ENV_VAR);
  return {
    statusCode: 200,
    body: `Hello ${name}!`,
  };
};
