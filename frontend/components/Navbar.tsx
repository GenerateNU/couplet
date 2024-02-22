import React from "react";
import { View } from "react-native";
import NavButton from "./NavButton";

export default function Navbar() {
  return (
    <View
      style={{
        backgroundColor: "black",
        flexDirection: "row",
        position: "absolute",
        bottom: 0
      }}
    >
      <NavButton route="People" />
      <NavButton route="Favorites" />
      <NavButton route="Matches" />
      <NavButton route="Profile" />
    </View>
  );
}
