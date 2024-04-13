/* eslint-disable */
import { router } from "expo-router";
import React from "react";
import { SafeAreaView, StyleSheet, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import CardWrapper from "../components/Profile/CardWrapper";
import PurpleProfileCard from "../components/Profile/PurpleProfileCard";
import SettingsCard from "../components/Profile/SettingsCard";
import UserDetails from "../components/Profile/UserDetails";

/* // eslint-disable global-require */
export default function Profile() {
  return (
    <SafeAreaView style={{ flex: 1, gap: 15 }}>
      <View style={{ alignSelf: "center" }}>
        <UserDetails
          onPress={() => router.push("ProfileScreens/AccountDetails")}
          profileImageSource={undefined}
          name="TiffanyA"
          editDetailsText="Edit Account Details"
          chevronColor=""
        />
      </View>
      <View style={styles.purpleCardsContainer}>
        <PurpleProfileCard
          imageUrl={require("../assets/pencil.png")}
          name="Profile"
          detailText="View or change your profile"
          onPress={() => router.push("ProfileScreens/AccountDetails")}
        />
        <PurpleProfileCard
          imageUrl={require("../assets/magnifyingGlass.png")}
          name="Preferences"
          detailText="Manage what you are looking for"
          onPress={() => router.push("ProfileScreens/AccountDetails")}
        />
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard
            text="Account Details"
            img={require("../assets/redProfile.png")}
            onPress={() => router.push("ProfileScreens/AccountDetails")}
          />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard
            text="About Couplet"
            img={require("../assets/redAbout.png")}
            onPress={() => router.push("ProfileScreens/AboutCouplet")}
          />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard
            text="Help & Support"
            img={require("../assets/redHelp.png")}
            onPress={() => router.push("ProfileScreens/HelpAndSupport")}
          />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard
            text="Settings"
            img={require("../assets/redSettings.png")}
            last
            onPress={() => router.push("ProfileScreens/Settings")}
          />
        </CardWrapper>
      </View>
      <Navbar activePage="Profile" />
    </SafeAreaView>
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
