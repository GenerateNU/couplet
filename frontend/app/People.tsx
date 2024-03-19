import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";

export default function People() {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>People</Text>
      <Navbar />
    </View>
  );
}
