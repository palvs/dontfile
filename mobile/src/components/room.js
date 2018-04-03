import React, { Component } from 'react';
import { StyleSheet, Text, View, ScrollView } from 'react-native';

import File from './file'

const styles = StyleSheet.create({
    container: {
        flex: 1,
        backgroundColor: '#fff',
        alignItems: 'center',
        justifyContent: 'center',
    },
});

export default class Room extends Component {
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
            <View>
                <Text>Room Name</Text>
                <ScrollView>
                    {this.state.Files.map(file => (
                        <File source={'http://localhost:3001/' + this.state.Directory + '/' + file.Name}
                              name={file.Name}
                              date={file.ModTime}
                              size={file.Size}
                        />
                    ))}
                </ScrollView>
            </View>
        );
    }
}