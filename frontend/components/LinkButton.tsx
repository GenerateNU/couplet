import { Link } from "expo-router";
import React from "react";
import { Pressable, Text } from "react-native";

export default function LinkButton({ text }: { text: string }) {
  return (
    
      <Pressable
        style={{
          borderStyle: "solid",
          borderColor: "black",
          backgroundColor: "black",
          borderWidth: 1,
          padding: "5%",
          borderRadius: 100
        }}
      >
        <Link href="/People">
          <Text style={{ color: "white" }}>{text}</Text>
        </Link>
      </Pressable>
  );
}
