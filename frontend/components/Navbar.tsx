import { View } from 'react-native'
import React from 'react'
import NavButton from './NavButton'

export default function Navbar() {
  return (
    <View style={{
        backgroundColor: "black",
        height: 75,
        width: "100%",
        flexDirection: "row",
        position:"absolute",
        bottom: 0,
      }}>
        <NavButton route="Events"/>
        <NavButton route="Matches" />
        <NavButton route="Settings" />
        <NavButton route="Users" />
    </View>
  )
}
