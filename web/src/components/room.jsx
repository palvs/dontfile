import React, { Component } from 'react';
import { withStyles } from 'material-ui/styles';
import Grid from 'material-ui/Grid';
import Typography from 'material-ui/Typography';

import File from './file'
import Dropzone from 'react-dropzone'


const styles = theme => ({
  root: {
    flexGrow: 1,
  },
  control: {
    padding: theme.spacing.unit * 2,
  },
  drop: {
    background: "#f5f5f5",
    bottom: "0",
    padding: "1em",
    position: "fixed",
  }
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
    const path = window.location.pathname

    fetch('http://localhost:3002'+path)
    .then(resp => {
      resp.json()
      .then((resp) => {
        this.setState(resp)
      })
    })
  }


  deleteFile(file) {
    fetch('http://localhost:3002/'+this.state.Directory+'/'+file.Name, { method: 'DELETE' })
    .then(resp => {
      const files = this.state.Files.filter(f => f.Name !== file.Name)
      this.setState({...this.state.Directory, Files: files})
    })
  }

  uploadFile(files) {
    console.log('upload')
    const uploads = files.map(file => {
      const formData = new FormData();
      formData.append("file", file);

      fetch('http://localhost:3002/'+this.state.Directory, {
        method: 'POST',
        body: formData
      })
    })
  }
  
  render() {
    return (
      <div>
        {console.log('render')}
        <Grid container 
              className={this.props.classes.root} 
              spacing={16}>
          <Grid item 
                xs={12}>
            <Grid container 
                  justify="center" 
                  spacing={16}>
              {this.state.Files.map(file => (
                  <Grid item>
                    <File source={'http://localhost:3002/'+this.state.Directory+'/'+file.Name}
                      name={file.Name}
                      date={file.ModTime}
                      size={file.Size}
                      delete={() => { this.deleteFile(file) }}
                    />
                  </Grid>
              ))}
            </Grid>
            <Grid container 
                  className={this.props.classes.drop}
                  justify="flex-end"
                  alignItems="stretch"
                  direction="column" 
                  spacing={16}>
              <Dropzone onDrop={this.uploadFile.bind(this)}>
                <p>Arraste arquivos aqui!</p>
              </Dropzone>
            </Grid>
          </Grid>
        </Grid>
      </div>
    );
  }
}

export default withStyles(styles)(Room);
