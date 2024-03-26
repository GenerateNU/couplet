import React from "react";
import { StyleSheet, Text, View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import PurpleProfileCard from "./PurpleProfileCard";
import UserDetails from "./UserDetails";

export default function Profile() {
  return (
    <View style={{ flex: 1 }}>
      <Text style={{ fontFamily: "DMSansRegular" }}>ProfileA</Text>
      <UserDetails
        onPress={function (): void {
          console.log("Edit Profile");
        }}
        profileImageSource={undefined}
        name={"Tiffany"}
        editDetailsText={"Edit Account Details"}
        chevronColor={""}
      />
      <View style={styles.purpleCardsContainer}>
        <PurpleProfileCard
          imageUrl={"../assets/ProfilePurple.png"}
          name={"My Profile"}
          detailText={"View and make changes to your profile"}
          onPress={function (): void {
            console.log("My Profile");
          }}
        />
        <PurpleProfileCard
          imageUrl={"../assets/ProfilePurple.png"}
          name={"My Preferences"}
          detailText={"Manage what you are looking for"}
          onPress={function (): void {
            console.log("My Preferences");
          }}
        />
      </View>

      <Navbar />
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
