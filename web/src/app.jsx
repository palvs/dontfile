import React, { Component } from 'react';
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
      <div>
        <div className={this.props.classes.main}>
          <Room />
        </div>
        <div className={this.props.classes.navBottom}>
          <Drop />
        </div>
      </div>
    )
  }
}

export default withStyles(styles)(App);
