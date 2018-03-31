import React, { Component } from 'react';

import File from './components/file'

class App extends Component {
  constructor() {
    super()
    this.state = {
      Directory: '',
      Files: []
    }
  }

  componentDidMount() {
    fetch('http://localhost:3001/test')
    .then(resp => {
      resp.json()
      .then((resp) => {
        this.setState(resp)
      })
    })
  }
  
  render() {
    return (
      <div>
        <h1>{this.state.Directory}</h1>
        {this.state.Files.map((file) => {
          
        })}
      </div>
    );
  }
}

export default App;
