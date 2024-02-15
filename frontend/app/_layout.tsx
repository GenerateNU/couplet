import { Slot } from "expo-router";
import React from "react";
import { View } from "react-native";
import Navbar from "../components/Navbar";

export default function Layout() {
  return (
    <View style={{ height: "100%" }}>
      <Slot />
      <Navbar />
    </View>
  );
}
