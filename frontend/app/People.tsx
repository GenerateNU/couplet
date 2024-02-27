import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Navbar";

export default function People() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text>People</Text>
      <Navbar />
    </View>
  );
}
