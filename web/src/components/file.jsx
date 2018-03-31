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
    render() {
        return (
            <div>
                <Card className={this.props.classes.card}>
                    <CardMedia
                        className={this.props.classes.media}
                        image={this.props.source}
                        title="Image Placeholder"
                    />
                    <CardContent>
                        <Typography variant="headline" 
                                    component="h2">
                            {this.props.name}
                        </Typography>
                        <Typography className={this.props.classes.pos} 
                                    color="textSecondary">
                            {this.props.date} | {this.props.size}
                        </Typography>
                    </CardContent>
                    <CardActions>
                        <Button href={this.props.source}
                                size="small" 
                                color="primary" 
                                className={this.props.classes.button}>
                            <GetApp />
                        </Button>
                        <Button size="small" 
                                color="secondary" 
                                className={this.props.classes.button}>
                            <Delete />
                        </Button>
                        <Button size="small" 
                                color="default" 
                                className={this.props.classes.button}>
                            <Bookmark />
                        </Button>
                    </CardActions>
                </Card>
            </div>
        )
    }
}

export default withStyles(styles)(File);