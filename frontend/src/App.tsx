import React from 'react';
import logo from './logo.svg';
import './App.css';
import { Get } from './Utils/ApiConnector';

function App() {
  const [status, setStatus] = React.useState("");

  React.useEffect(() => {
    Get('http://localhost:8080/healthcheck')
      .then(data => setStatus(data.data))
      .catch(() => setStatus("Backend System is down?"));
  });

  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          {status}
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
