import { Slot } from "expo-router";
import React from "react";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Layout() {
  return (
    <SafeAreaView style={{ height: "100%" }}>
      <Slot />
    </SafeAreaView>
  );
}
