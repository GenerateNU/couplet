import React from "react";
import { View } from "react-native";
import NavButton from "./NavButton";
import COLORS from "../../colors";

const ACTIVE_HOME = require('../../assets/navbarhomeactive.png')
const INACTIVE_HOME = require('../../assets/navbarhomeinactive.png')
const ACTIVE_LIKES = require('../../assets/navbarlikesactive.png')
const INACTIVE_LIKES = require('../../assets/navbarlikesinactive.png')
const ACTIVE_MATCHES = require('../../assets/navbarmatchesactive.png')
const INACTIVE_MATCHES = require('../../assets/navbarmatchesinactive.png')
const ACTIVE_PROFILE = require('../../assets/navbarprofileactive.png')
const INACTIVE_PROFILE = require('../../assets/navbarprofileinactive.png')

interface NavbarProps {
  activePage: string;
}

export default function Navbar({ activePage }: NavbarProps) {
  return (
    <View
      style={{
        flex: 1,
        width: "100%",
        borderWidth: 1,
        borderTopColor: COLORS.darkGray,
        backgroundColor: COLORS.white,
        flexDirection: "row",
        position: "absolute",
        justifyContent: "space-around",
        bottom: 0
      }}
    >
      <NavButton route="Home" icon={activePage === 'Home' ? ACTIVE_HOME : INACTIVE_HOME} />
      <NavButton route="Favorites" icon={activePage === 'Likes' ? ACTIVE_LIKES : INACTIVE_LIKES} />
      <NavButton route="Matches" icon={activePage === 'Matches' ? ACTIVE_MATCHES : INACTIVE_MATCHES} />
      <NavButton route="Profile" icon={activePage === 'Profile' ? ACTIVE_PROFILE : INACTIVE_PROFILE} />
    </View>
  );
}
