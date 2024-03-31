import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";

export default function Profile() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>Profile</Text>
      <Navbar activePage="Profile" />
    </View>
  );
}
