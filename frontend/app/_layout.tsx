import {
  DMSans_700Bold as DMSansBold,
  DMSans_500Medium as DMSansMedium,
  DMSans_400Regular as DMSansRegular
} from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import { Slot } from "expo-router";
import React from "react";
import { SafeAreaView } from "react-native-safe-area-context";

export default function Layout() {
  const [fontsLoaded] = useFonts({
    DMSansRegular,
    DMSansMedium,
    DMSansBold
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <SafeAreaView style={{ height: "100%" }}>
      <Slot />
    </SafeAreaView>
  );
}
