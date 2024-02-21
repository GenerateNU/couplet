import { View, Text, Pressable } from 'react-native'
import React from 'react'

export default function TagButton({text, selected}: {text: string, selected?: boolean}) {
    return (
            <Pressable style={{ 
                borderStyle:"solid", 
                borderColor: "black", 
                backgroundColor: selected ? "black" : "white",
                borderWidth: 1,
                padding: "3%", 
                paddingLeft: "5%", 
                paddingRight: "5%", 
                borderRadius: 100, 
                marginLeft: "2%"}} onPress={() => console.log(text)}>
                <Text style={{color: selected ? "white":"black"}}>{text}</Text>
            </Pressable>
    )
}