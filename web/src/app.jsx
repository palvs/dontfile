import React, { Component } from 'react';
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import Typography from 'material-ui/Typography';

import File from './components/file'
import Room from './components/room'

class App extends Component {
  render() {
    return (
      <Room />
    )
  }
}

export default App;
