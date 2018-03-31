import React from 'react';
import PropTypes from 'prop-types';
import { withStyles } from 'material-ui/styles';
import Card, { CardActions, CardContent, CardMedia } from 'material-ui/Card';
import Button from 'material-ui/Button';
import Typography from 'material-ui/Typography';

import GetApp from 'material-ui-icons/GetApp';
import Bookmark from 'material-ui-icons/Bookmark';
import Delete from 'material-ui-icons/Delete';

const styles = theme => ({
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
});


function FileCard(props) {
    const { classes } = props;

    return (
        <div>
            <Card className={classes.card}>
                <CardMedia
                    className={classes.media}
                    image="http://via.placeholder.com/220x220"
                    title="Image Placeholder"
                />
                <CardContent>
                    <Typography variant="headline" component="h2">
                        {props.name}
                    </Typography>
                    <Typography className={classes.pos} color="textSecondary">
                        {props.date} | {props.size}
                    </Typography>
                </CardContent>
                <CardActions>
                    <Button size="small" color="primary" className={classes.button}>
                         <GetApp /> 
                    </Button>
                    <Button size="small" color="secondary" className={classes.button}>
                        <Delete /> 
                    </Button>
                    <Button size="small" color="" className={classes.button}>
                        <Bookmark /> 
                    </Button>
                </CardActions>
            </Card>
        </div>
    );
}

FileCard.propTypes = {
    classes: PropTypes.object.isRequired,
};

export default withStyles(styles)(FileCard);