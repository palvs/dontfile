import React, { Component } from 'react'
import { withStyles } from 'material-ui/styles'
;
import Dropzone from 'react-dropzone'

import Grid from 'material-ui/Grid';
import Typography from 'material-ui/Typography';
import BottomNavigation from 'material-ui/BottomNavigation';

const styles = theme => ({
    dropzone: {
        flexGrow: 1,
        width: "100%",
        height: "30%",
        border: "dashed 1px #000",
        margin: "0"
    }
});

class Drop extends Component {
    constructor() {
        super()
        this.state = { files: [] }
    }

    onDrop(files) {
        this.setState({ files });
    }

    render() {
        return (            
            <div>
                <Dropzone onDrop={this.onDrop.bind(this)} 
                          className={this.props.classes.dropzone}>
                    <p>Arraste arquivos aqui!</p>
                </Dropzone>
            </div>    
        );
    }
}

export default withStyles(styles)(Drop)