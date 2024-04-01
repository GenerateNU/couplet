import React from "react";
import { StyleSheet, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import CardWrapper from "../components/Profile/CardWrapper";
import SettingsCard from "../components/Profile/SettingsCard";
import PurpleProfileCard from "./PurpleProfileCard";
import UserDetails from "./UserDetails";

export default function Profile() {
  return (
    <View style={{ flex: 1, gap: 15 }}>
      <View style={{ alignSelf: "center" }}>
        <UserDetails
          onPress={() => console.log("Edit Account Details")}
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
          onPress={() => console.log("Edit Account Details")}
        />
        <PurpleProfileCard
          imageUrl="../assets/ProfilePurple.png"
          name="My Preferences"
          detailText="Manage what you are looking for"
          onPress={() => console.log("Edit Account Details")}
        />
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          <SettingsCard text="Settings" img={require("../assets/gear.png")} />
          <SettingsCard text="Help and Support" img={require("../assets/help.png")} />
          <SettingsCard text="About Couplet" img={require("../assets/info.png")} />
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
