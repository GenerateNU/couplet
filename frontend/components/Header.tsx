import React from "react";
import { Text, View } from "react-native";

export default function Header() {
  return (
    <View
      style={{
        backgroundColor: "black",
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        padding: 20
      }}
    >
      <Text
        style={{
          color: "white",
          justifyContent: "center",
          textAlign: "center",
          width: "100%",
          fontSize: 20
        }}
      >
        Couplet
      </Text>
    </View>
  );
}
