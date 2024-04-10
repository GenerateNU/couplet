import React from "react";
import { StyleSheet, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import CardWrapper from "../components/Profile/CardWrapper";
import SettingsCard from "../components/Profile/SettingsCard";
import PurpleProfileCard from "./PurpleProfileCard";
import UserDetails from "./UserDetails";

import GEAR from "../assets/gear.png";
import HELP from "../assets/help.png";
import INFO from "../assets/info.png";
import LOGOUT from "../assets/lougout.png";

export default function Profile() {
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
          <SettingsCard text="Settings" img={GEAR} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="Help and Support" img={HELP} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="About Couplet" img={INFO} />
          {/* //eslint-disable-next-line global-require */}
          <SettingsCard text="Log Out" img={LOGOUT} last />
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
