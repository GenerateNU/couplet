import React from "react";
import { Pressable, Text } from "react-native";


export default function TagButton({ text, selected }: { text: string; selected: boolean}) {
  return (
    <Pressable
      style={{
        borderStyle: "solid",
        borderColor: "black",
        backgroundColor: selected ? "black" : "white",
        borderWidth: 1,
        padding: "5%",
        borderRadius: 100
      }}
      onPress={() => console.log(text)}
    >
      <Text style={{ color: selected ? "white" : "black" }}>{text}</Text>
    </Pressable>
  );
}

