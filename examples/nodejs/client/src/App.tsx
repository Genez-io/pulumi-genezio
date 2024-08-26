import { useState } from 'react'
import reactLogo from './assets/react.svg'
import viteLogo from '/vite.svg'
import './App.css'

import axios from 'axios'

function App() {
  const [serverString, setServerString] = useState('')
  const [name, setName] = useState('')

  const envVar = import.meta.env.VITE_HELLO_WORLD_FUNCTION_URL

  const makeRequest = async () => {
    const url = `${envVar}?name=${name}`
    const response = await axios.get(url)
    const data = response.data
    setServerString(data.message)
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
          <div style={{marginBottom:"1em"}}>
          <label>Name:</label>
          <input
            style={{marginLeft:"5px"}} 
            type="text"
            value={name}
            onChange={(e) => setName(e.target.value)}
          ></input>
          </div>
        <button onClick={() => makeRequest()}>
          Call Serverless Function
        </button>
        <p>
          Edit <code>src/App.tsx</code> and save to test HMR
        </p>
        {
          serverString && (
            <div>
              The serverless function says: <strong>{serverString}</strong>
            </div>
          )
        }
        
      </div>
      <p className="read-the-docs">
        Click on the Vite and React logos to learn more
      </p>
    </>
  )
}

export default App
