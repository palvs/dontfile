import React, { Component } from 'react';

import { BrowserRouter as Router, Route, DefaultRoute } from 'react-router-dom'

import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import Typography from 'material-ui/Typography';

import File from './components/file'
import Room from './components/room'
import Drop from './components/drop'

const styles = theme => ({
  navBottom: {
    background: "#f5f5f5",
    overflow: "hidden",
    position: "fixed",
    bottom: "0",
    width: "100%"
  }
});

class App extends Component {
  render() {

    return (
      <Router>
        <div>
          <Route exact path="/" component={Drop} />
          <Route path="*" component={Room} />
        </div>
      </Router>
    )
  }
}

export default withStyles(styles)(App);
