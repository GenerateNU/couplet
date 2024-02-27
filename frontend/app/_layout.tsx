import { Slot } from "expo-router";
import React from "react";
import { View } from "react-native";

export default function Layout() {
  return (
    <View style={{ height: "100%" }}>
      <Slot />
    </View>
  );
}
