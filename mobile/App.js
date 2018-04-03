import React from 'react';
import { StyleSheet, Text, View, ScrollView } from 'react-native';

import Room from './src/components/room'

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});

export default class App extends React.Component {
  render() {
    return (
      <View style={styles.container}>
        <Room />
      </View>
    );
  }
}