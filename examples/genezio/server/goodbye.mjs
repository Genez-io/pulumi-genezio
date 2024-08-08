export const handler = async (event) => {
  console.log("Function was called");
  const name = event.queryStringParameters?.name || "test";

  return {
    statusCode: 200,
    body: `Goodbye ${name}!`,
  };
};
