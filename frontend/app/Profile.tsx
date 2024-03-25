import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import CardWrapper from "../components/Profile/CardWrapper";
import SettingsCard from "../components/Profile/SettingsCard";
export default function Profile() {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }
  return (
    <View style={{ flex: 1 }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>Profile</Text>
      <CardWrapper>
        <SettingsCard text="Settings" img={require("../assets/gear.png")} />
        <SettingsCard text="Help and Support" img={require("../assets/help.png")} />
        <SettingsCard text="About Couplet" img={require("../assets/info.png")} />
        <SettingsCard text="Log Out" img={require("../assets/lougout.png")} last />
      </CardWrapper>
      {/* <PhotoPicker /> */}
      <Navbar />
    </View>
  );
}
