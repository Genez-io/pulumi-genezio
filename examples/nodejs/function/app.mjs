export const handler = async (event) => {
  console.log("function was called");
  const name = event.queryStringParameters?.name || "World 23456";
  return {
    statusCode: 200,
    body: JSON.stringify({ message: `Hello, ${name}` }),
  };
};
