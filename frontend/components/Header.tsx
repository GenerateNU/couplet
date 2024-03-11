import React from "react";
import { Text, View } from "react-native";

export default function Header() {
  return (
    <View
      style={{
        backgroundColor: "white",
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        padding: 20
      }}
    >
      <Text
        style={{
          color: "black",
          justifyContent: "center",
          textAlign: "left",
          width: "100%",
          fontSize: 20
        }}
      >
        Couplet
      </Text>
    </View>
  );
}
