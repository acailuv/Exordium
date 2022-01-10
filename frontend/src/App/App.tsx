import React from 'react';
import logo from '../_assets/logo.svg';
import './App.less';
import { Get } from '../Utils/ApiConnector';
import { Typography } from 'antd';

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
        <p>Don't forget to add <Typography.Text code>@import '~antd/dist/antd.less';</Typography.Text> into the *.less file when making a new component.</p>
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
