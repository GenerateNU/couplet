import React from "react";
import { View } from "react-native";
import HomeScreen from "../components/HomeScreen";
import Navbar from "../components/Navbar";

export default function Home() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <View style={{ flex: 1, marginBottom: 35 }}>
        <HomeScreen />
      </View>
      <Navbar />
    </View>
  );
}
