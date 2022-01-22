import React from 'react';
import './App.css';

class App extends React.Component {
  state = {
    book: 0,
    page: 0,
    textLines: []
  }

  handleSubmit(event) {
    event.preventDefault()

    fetch(`http://localhost:5000/book`, {
      method: 'POST',
      body: JSON.stringify({
        book: this.state.book,
        page: this.state.page
      }),
      headers: {
        "Content-Type": "application/json"
      }
    })
    .then(response => response.json())
    .then(data => {
      this.setState({textLines: data.text})
    });
  }

  handleBookChange(e) {
    this.setState({
      book: e.target.value.replace(/\s/g, '')
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
        <form onSubmit={(e) => this.handleSubmit(e)}>
          <label>
            Location:
            <textarea
              type="text"
              name="Location"
              placeholder="(0..9)*"
              onChange={(e) => this.handleBookChange(e)}
              onKeyDown={(e) => this.onKeydown(e)}
            />
          </label>
          <br></br>
          <label>
            Page:
            <input
              type="number"
              name="Page"
              placeholder="1-410"
              onChange={(e) => this.handlePageChange(e)}
              onKeyDown={(e) => this.onKeydown(e)}
            />
          </label>
          <input type="submit" value="Submit"/>
        </form>

        <div className="book-text">
          {this.state.textLines.map((value)=> {
            return <p className='textLine'>{value}</p>
          })}
        </div>
      </div>
    );
  }
}

export default App;
