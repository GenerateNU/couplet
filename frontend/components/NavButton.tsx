import { Button } from 'react-native-paper'

import { View, Text } from 'react-native'
import React from 'react'
import { Link } from 'expo-router'

export default function NavButton({route}: {route: string}) {
  return (
      <View style={{width:"100%", flex:1, justifyContent: "center"}}>
        <Link href={`/${route}`} style={{
                    width: "auto", justifyContent: "center",}}>
            <Button style={{
                width: "auto",
                height: "100%",
                justifyContent: "center",
            }}>

                <Text style={{
                    color: "white",
                    width: "auto",
                    textAlign: "center",
                    textAlignVertical: "center",
                }}>
                    {route}
                </Text>
            </Button>
        </Link>
      </View>
  )
}