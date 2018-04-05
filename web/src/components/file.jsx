import React, { Component } from 'react';

import { withStyles } from 'material-ui/styles';
import Card, { CardActions, CardContent, CardMedia } from 'material-ui/Card';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';

import GetApp from 'material-ui-icons/GetApp';
import Bookmark from 'material-ui-icons/Bookmark';
import Delete from 'material-ui-icons/Delete';

const styles = {
    card: {
        width: 220
    },
    media: {
        height: 220,
    },
    title: {
        marginBottom: 16,
        fontSize: 14,
    },
    pos: {
        marginBottom: 12,
    },
}

class File extends Component {

    renderName(name) {
        if(name.length > 18) {
            name = `${name.substring(0, 0.1*name.length)}...${name.substring(0.8*name.length, name.length)}`
        }
        return name
    }

    renderDate(date) {
        return new Date(date).toDateString()
    }

    renderSize(size) {
        const i = Math.floor(Math.log(size) / Math.log(1024));
        return (size / Math.pow(1024, i)).toFixed(2) * 1 + ' ' + ['B', 'kB', 'MB', 'GB', 'TB'][i];
    }

    render() {
        console.log(this.props)
        const { name, date, size, classes, source } = this.props

        return (
            <div>
                <Card className={classes.card}>
                    <CardMedia
                        className={classes.media}
                        image={source}
                        title="Image Placeholder"
                    />
                    <CardContent>
                        <Typography variant="body2">
                            {this.renderName(name)}
                            <Typography className={classes.pos}
                                color="textSecondary">
                                {this.renderSize(size)}
                            </Typography>   
                        </Typography>
                        <Typography className={classes.pos} 
                                    color="textSecondary">
                            {this.renderDate(date)}
                        </Typography>
                    </CardContent>
                    <CardActions>
                        <Button href={source}
                                size="small" 
                                color="primary" 
                                className={classes.button}>
                            <GetApp />
                        </Button>
                        <Button size="small" 
                                color="secondary" 
                                className={classes.button}
                                onClick={this.props.delete}>
                            <Delete />
                        </Button>
                        <Button size="small" 
                                color="default" 
                                className={classes.button}>
                            <Bookmark />
                        </Button>
                    </CardActions>
                </Card>
            </div>
        )
    }
}

export default withStyles(styles)(File);