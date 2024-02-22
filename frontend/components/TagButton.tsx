import React from "react";
import { Pressable, Text } from "react-native";

export default function TagButton({ text }: { text: string }) {
  return (
    <Pressable
      style={{
        borderStyle: "solid",
        borderColor: "black",
        borderWidth: 1,
        padding: 10,
        paddingLeft: 20,
        paddingRight: 20,
        borderRadius: 100,
      }}
      onPress={() => console.log("Favorites")}
    >
      <Text>{text}</Text>
    </Pressable>
  );
}
