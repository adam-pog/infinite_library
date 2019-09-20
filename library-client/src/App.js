import React from 'react';
import logo from './logo.svg';
import './App.css';

function App() {
  return (
    <div className="App">
      <header className="App-header">
        <img src={logo} className="App-logo" alt="logo" />
        <p>
          Edit <code>src/App.js</code> and save to reload.
        </p>
        <a
          className="App-link"
          href="https://reactjs.org"
          target="_blank"
          rel="noopener noreferrer"
        >
          Learn React
        </a>

        <form>
          <label>
            Name:
            <textarea type="text" name="name" placeholder="asdgoij"/>
          </label>
          <input type="button" value="Submit" />
        </form>

      </header>
    </div>
  );
}

export default App;
