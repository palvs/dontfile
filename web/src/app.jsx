import React, { Component } from 'react';

import { BrowserRouter as Router, Route, DefaultRoute } from 'react-router-dom'

import { withStyles } from 'material-ui/styles'
import Grid from 'material-ui/Grid'
import Typography from 'material-ui/Typography'

import Home from './components/home'
import File from './components/file'
import Room from './components/room'

class App extends Component {
  render() {

    return (
      <Router>
        <div>
          <Route exact path="/" component={Home} />
          <Route path="*" component={Room} />
        </div>
      </Router>
    )
  }
}

export default App
