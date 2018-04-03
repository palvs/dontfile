import React, { Component } from 'react'
import {
    StyleSheet,
    Text,
    View

} from 'react-native'

const width = Dimensions.get('screen').width;

export default class File extends Component {
    constructor(props) {
        super(props)
        this.state = { file: this.props.file }
    }

    render() {
        const { file } = this.state

        return (
            <View>
                
            </View>
        )
    }
}

