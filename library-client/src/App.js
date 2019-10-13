import React from 'react';
import logo from './logo.svg';
import './App.css';

class App extends React.Component {
  state = {
    location: 0,
    page: 0,
    text: ''
  }

  handleSubmit(event) {
    event.preventDefault()

    fetch('http://localhost:8081', {
      method: 'POST',
      body: JSON.stringify({
        location: this.state.location,
        page: parseInt(this.state.page)
      })
    })
    .then(response => response.text())
    .then(data => {
      this.setState({text: data})
    });
  }

  handleLocationChange(e) {
    this.setState({
      location: e.target.value.replace(/\s/g, '')
    })
  }

  handlePageChange(e) {
    this.setState({
      page: e.target.value
    })
  }

  onKeydown(e) {
    if (e.key === 'Enter') {
      this.handleSubmit(e)
    }
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

          <form onSubmit={(e) => this.handleSubmit(e)}>
            <label>
              Location:
              <textarea
                type="text"
                name="Location"
                placeholder="(0..9)*"
                onChange={(e) => this.handleLocationChange(e)}
                onKeyDown={(e) => this.onKeydown(e)}
              />
            </label>
            <br></br>
            <label>
              Page:
              <input
                type="number"
                name="Page"
                placeholder="0-409"
                onChange={(e) => this.handlePageChange(e)}
                onKeyDown={(e) => this.onKeydown(e)}
              />
            </label>
            <input type="submit" value="Submit"/>
          </form>

        </header>
        <p>{this.state.text}</p>
      </div>
    );
  }
}

export default App;
