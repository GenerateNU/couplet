import React from "react";
import { View } from "react-native";
import { Title } from "react-native-paper";

export default function HomeScreen() {
  return (
    <View
      style={{
        paddingTop: "100%",
        width: "100%",
        borderRadius: 12,
        alignSelf: "center",
        marginBottom: 10,
        minHeight: "100%",
        paddingBottom: "30%",
      }}
    >
      <Title
        style={{
          width: "100%",
          borderRadius: 12,
          alignSelf: "center",
          marginBottom: 10,
          minHeight: "100%",
          paddingBottom: "30%",
          textAlign: "center",
        }}
      >
        Welcome to Couplet!
      </Title>
    </View>
  );
}
