import React from "react";
import { View } from "react-native";
import HomeScreen from "../components/Home/HomeScreen";
import Navbar from "../components/Layout/Navbar";

export default function Home() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <View style={{ flex: 1 }}>
        <HomeScreen />
      </View>
      <Navbar />
    </View>
  );
}
