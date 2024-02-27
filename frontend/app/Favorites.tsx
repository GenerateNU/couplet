import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Navbar";

export default function Favorites() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text>Favorites</Text>
      <Navbar />
    </View>
  );
}
