import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import LifestyleReligion from "./Onboarding/LifestyleReligion";

export default function Profile() {
  return (
    // <View style={{ flex: 1, justifyContent: "space-between" }}>
    //   <Text style={{ fontFamily: "DMSansRegular" }}>Profile</Text>
    //   <Navbar />
    // </View>
    <LifestyleReligion />
  );
}
