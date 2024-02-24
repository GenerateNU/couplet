import React from "react";
import { View } from "react-native";
import NavButton from "./NavButton";

export default function Navbar() {
  return (
    <View
      style={{
        backgroundColor: "black",
        height: 75,
        width: "100%",
        flexDirection: "row",
        position: "absolute",
        bottom: 0
      }}
    >
      <NavButton route="routes/Home" />
      <NavButton route="routes/People" />
      <NavButton route="routes/Favorites" />
      <NavButton route="routes/Matches" />
      <NavButton route="routes/Profile" />
    </View>
  );
}
