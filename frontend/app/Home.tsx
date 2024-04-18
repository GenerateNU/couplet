import React from "react";
import { SafeAreaView, View } from "react-native";
import HomeScreen from "../components/Home/HomeScreen";
import Navbar from "../components/Layout/Navbar";

export default function Home() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <SafeAreaView style={{ flex: 1 }}>
        <HomeScreen />
      </SafeAreaView>
      <Navbar activePage="Home" />
    </View>
  );
}
