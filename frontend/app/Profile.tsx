import { router } from "expo-router";
import * as SecureStorage from "expo-secure-store";
import React from "react";
import { StyleSheet, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import ContinueButton from "../components/Onboarding/ContinueButton";
import CardWrapper from "../components/Profile/CardWrapper";
import SettingsCard from "../components/Profile/SettingsCard";
import PurpleProfileCard from "./PurpleProfileCard";
import UserDetails from "./UserDetails";

/* // eslint-disable global-require */
export default function Profile() {
  const handleLogout = async () => {
    await SecureStorage.deleteItemAsync("appleAuth");
    router.push("/");
  };

  return (
    <View style={{ flex: 1, gap: 15 }}>
      <View style={{ alignSelf: "center" }}>
        <UserDetails
          onPress={() => "Placeholder"}
          profileImageSource={undefined}
          name="TiffanyA"
          editDetailsText="Edit Account Details"
          chevronColor=""
        />
      </View>
      <View style={styles.purpleCardsContainer}>
        <PurpleProfileCard
          imageUrl="../assets/ProfilePurple.png"
          name="My Profile"
          detailText="View and make changes to your profile"
          onPress={() => "Placeholder"}
        />
        <PurpleProfileCard
          imageUrl="../assets/ProfilePurple.png"
          name="My Preferences"
          detailText="Manage what you are looking for"
          onPress={() => "Placeholder"}
        />
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="Settings" img={require("../assets/gear.png")} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="Help and Support" img={require("../assets/help.png")} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="About Couplet" img={require("../assets/info.png")} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="Log Out" img={require("../assets/lougout.png")} last />
        </CardWrapper>
      </View>
      <Navbar activePage="Profile" />
    </View>
  );
}

const styles = StyleSheet.create({
  purpleCardsContainer: {
    display: "flex",
    justifyContent: "center",
    alignItems: "flex-start",
    gap: 20,
    alignSelf: "stretch",
    flexDirection: "row"
  }
});
