export const handler = async (event) => {
  const name = event.queryStringParameters?.name || "World";
  console.log(JSON.stringify(event, null, 2));
  return {
    statusCode: 200,
    body: JSON.stringify({ message: `Hello ${name}` }),
  };
};
