import React, { Component } from 'react'
import {
    Dimensions,
    Linking,
    StyleSheet,
    Text,
    View
} from 'react-native'

import Icon from 'react-native-vector-icons/Ionicons';

const width = Dimensions.get('screen').width;

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: 'center',
        alignItems: 'center',
    },
    title: {
        fontSize: 20,
        textAlign: 'center',
        margin: 10,
    },
    subtitle: {
        textAlign: 'center',
        color: '#333333',
        marginBottom: 5,
    },
});

export default class File extends Component {
    constructor(props) {
        super(props)
        this.state = { file: this.props.file }
    }

    render() {
        const { file } = this.state

        return (
            <View style={styles.container}>
                <Text style={styles.title}>
                    {this.props.name}
                </Text>
                <Text style={styles.subtitle}>
                    {this.props.date}
                    {this.props.size}
                </Text>
                <Text style={styles.subtitle}>
                    <Icon name="ios-download" 
                          size={30} 
                          color="#4F8EF7"
                        onPress={() => Linking.openURL(this.props.source)}   />
                </Text>
            </View>
        )
    }
}

