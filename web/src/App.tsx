import React from 'react';
import logo from './logo.svg';
import './App.css';
import Websocket from './components/Websocket';

  export function wsURL(room: string, reqPath: string = "ws"): string {
    const protocol = window.location.protocol === "https:" ? "wss" : "ws";
    return `${protocol}://${
      process.env.NODE_ENV === "development"
        ? "localhost:3002"
        : window.location.host
    }/${reqPath}?room=${room}`;
  }

function App() {

  return (
    <div className="App">
      <Websocket
        url={wsURL('websocketTest')}
        onMessage={(msg: any) => {
          const obj = JSON.parse(msg);
          switch (obj.type) {
            case "test":
              console.info('test success')
              break;
            default:
              break;
          }
        }}
      />
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.tsx</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>
      </header>
    </div>
  );
}

export default App;
