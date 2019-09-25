import React from 'react';
import logo from './logo.svg';
import './App.css';

class App extends React.Component {
  state = {
    bookNum: 0
  }

  handleSubmit(event) {
    console.log('what a submit')
    event.preventDefault()

    fetch('http://localhost:8081/book/100/page/0', {
      method: 'GET',
    })
  }

  render() {
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

          <form onSubmit={this.handleSubmit}>
            <label>
              Name:
              <textarea type="text" name="name" placeholder="asdgoij"/>
            </label>
            <input type="submit" value="Submit"/>
          </form>

        </header>
      </div>
    );
  }
}

export default App;
