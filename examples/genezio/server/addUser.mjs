// Import the postgres driver library at the top of your file
import pg from "pg";
const { Pool } = pg;
import dotenv from "dotenv";

dotenv.config();

export const handler = async (event) => {
  if (!process.env.CORE_DATABASE_DATABASE_URL) {
    return {
      statusCode: 500,
      body: "Internal Server Error: CORE_DATABASE_URL environment variable is not set",
    };
  }
  const pool = new Pool({
    connectionString: process.env.CORE_DATABASE_DATABASE_URL,
    ssl: true,
  });

  await pool.query(
    "CREATE TABLE IF NOT EXISTS users (id SERIAL PRIMARY KEY, name VARCHAR(255))"
  );

  const data = JSON.parse(event.body);
  const name = data.name;
  await pool.query("INSERT INTO users (name) VALUES ($1)", [name]);
  const result = await pool.query("SELECT * FROM users");

  return {
    statusCode: 200,
    body: `User successfully added. There are now ${result.rows.length} users in the database.`,
  };
};
