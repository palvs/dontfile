import React, { Component } from 'react';
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import Typography from 'material-ui/Typography';

import File from './file'

const styles = theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    height: 140,
    width: 100,
  },
  control: {
    padding: theme.spacing.unit * 2,
  },
});

class Room extends Component {
  constructor() {
    super()
    this.state = {
      Directory: '',
      Files: []
    }
  }

  componentDidMount() {
    fetch('http://localhost:3001/room')
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
        <Grid container className={this.props.classes.root} spacing={16}>
          <Grid item xs={12}>
            <Grid container justify="center" spacing={16}>
              {this.state.Files.map(file => (
                  <Grid item>
                    <File source={'http://localhost:3001/'+this.state.Directory+'/'+file.Name}
                      name={file.Name}
                      date={file.ModTime}
                      size={file.Size}  
                    />
                  </Grid>
              ))}
            </Grid>
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default withStyles(styles)(Room);
