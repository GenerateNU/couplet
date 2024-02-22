import React from 'react'
import { Text, View, Pressable } from 'react-native'

export default function CallToAction() {
	return (
		<View
        style={{
          padding: 50,
          margin: 25,
          borderStyle: "solid",
          borderWidth: 1,
          backgroundColor: "gray"
        }}
      >
        <Text style={{fontSize: 32}}>Need someone to go with?</Text>
        <Pressable style={{
          padding: 15,
          marginTop: 25,
          borderStyle: "solid",
          borderColor: "black",
          borderWidth: 1,
          borderRadius: 10,
          backgroundColor: "black",
          width: "50%",
        }}>
          <Text style={{
          color: "white",
          textAlign: "center",
          }}>
            Match Now
          </Text>
        </Pressable>
      </View>
	)
}