import { useState } from "react";
import reactLogo from "./assets/react.svg";
import "./App.css";

import axios from "axios";

export default function App() {
  const [name, setName] = useState("");
  const [response, setResponse] = useState("");

  async function sayHello() {
    const url = `${import.meta.env.VITE_HELLO_WORLD_FUNCTION_URL}?name=${name}`;
    const response = await axios.get(url);
    const data = response.data;
    setResponse(data);
  }

  async function sayGoodbye() {
    const url = `${import.meta.env.VITE_GOODBYE_FUNCTION_URL}?name=${name}`;
    const response = await axios.get(url);
    const data = response.data;
    setResponse(data);
  }

  async function addToDatabase() {
    const url = `${import.meta.env.VITE_ADD_USER_FUNCTION_URL}`;
    const response = await axios.post(url, { name: name }, {
      headers: {
        "Content-Type": "application/json",
      },
    });
    const data = response.data;
    setResponse(data);
  }

  return (
    <>
      <div>
        <a href="https://genezio.com" target="_blank">
          <img
            src="https://raw.githubusercontent.com/Genez-io/graphics/main/svg/Logo_Genezio_White.svg"
            className="logo genezio light"
            alt="Genezio Logo"
          />
          <img
            src="https://raw.githubusercontent.com/Genez-io/graphics/main/svg/Logo_Genezio_Black.svg"
            className="logo genezio dark"
            alt="Genezio Logo"
          />
        </a>
        <a href="https://react.dev" target="_blank">``
          <img src={reactLogo} className="logo react" alt="React logo" />
        </a>
      </div>
      <h1>Genezio + React = ❤️</h1>
      <div className="card">
        <input
          type="text"
          className="input-box"
          onChange={(e) => setName(e.target.value)}
          placeholder="Enter your name"
        />
        <br />
        <br />

        <button onClick={() => sayHello()}>Say Hello</button>
        <button onClick={() => sayGoodbye()}>Say Goodbye</button>
        <button onClick={() => addToDatabase()}>Add to Database</button>
        <p className="read-the-docs">{response}</p>
      </div>
    </>
  );
}
