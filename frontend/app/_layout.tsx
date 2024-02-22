import { Slot } from "expo-router";
import React from "react";
import { View } from "react-native";


export default function Layout() {
  return (
    <View>
      <Slot />
    </View>
  );
}
