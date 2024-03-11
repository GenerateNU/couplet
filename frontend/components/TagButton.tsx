import React from "react";
import { Pressable, Text } from "react-native";

export default function TagButton({
  text,
  selected,
  onPress
}: {
  text: string;
  selected: boolean;
  onPress: () => void;
}) {
  return (
    <Pressable
      style={{
        borderStyle: "solid",
        borderColor: "black",
        backgroundColor: selected ? "black" : "white",
        borderWidth: 1,
        marginHorizontal: 10,
        padding: 15,
        borderRadius: 50
      }}
      onPress={onPress}
    >
      <Text style={{ color: selected ? "white" : "black" }}>{text}</Text>
    </Pressable>
  );
}
