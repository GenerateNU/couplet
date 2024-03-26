import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";

export default function People() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>People</Text>
      <Navbar />
    </View>
  );
}
