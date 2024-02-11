import { View, Text } from 'react-native'
import React from 'react'
import { Slot } from 'expo-router';
import Navbar from '../components/Navbar';

export default function _layout() {
  return (
    <View style={{height:"100%"}}>
      <Slot />
      <Navbar />
    </View>
  )
}