import { router } from "expo-router";
import * as SecureStorage from "expo-secure-store";
import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import ContinueButton from "../components/Onboarding/ContinueButton";

export default function Profile() {
  const handleLogout = async () => {
    await SecureStorage.deleteItemAsync("appleAuth");
    router.push("/");
  };

  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>Profile</Text>
      <View
        style={{ height: "100%", width: "100%", justifyContent: "center", alignSelf: "center" }}
      >
        <ContinueButton title="Logout" isDisabled={false} onPress={handleLogout} />
      </View>
      <Navbar activePage="Profile" />
    </View>
  );
}
