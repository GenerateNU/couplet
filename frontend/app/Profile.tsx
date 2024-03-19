import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import PhotoPicker from "../components/PhotoPicker";

export default function Profile() {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>Profile</Text>
      <PhotoPicker/>
      <Navbar />
    </View>
  );
}
