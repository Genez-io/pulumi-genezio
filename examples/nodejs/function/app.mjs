export const handler = async (event) => {
  console.log("function was called");
  const name = event.queryStringParameters?.name || "sdsadWorld 2364567dasdsad";
  if (process.env.POSTGRES_URL) {
    console.log(process.env.POSTGRES_URL);
  }
  return {
    statusCode: 200,
    body: JSON.stringify({ message: `Hello, ${name}` }),
  };
};
