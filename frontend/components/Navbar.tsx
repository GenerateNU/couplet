import React from "react";
import { View } from "react-native";
import NavButton from "./NavButton";

const CHAT = require("../assets/chat.png");
const HEART = require("../assets/heart.png");
const HOME = require("../assets/home.png");
const PERSON = require("../assets/person.png");

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
      <NavButton route="Home" icon={HOME} />
      <NavButton route="Favorites" icon={HEART} />
      <NavButton route="Matches" icon={CHAT} />
      <NavButton route="Profile" icon={PERSON} />
    </View>
  );
}
