import React from "react";
import { View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import MatchesScreen from "../components/Matches/MatchesScreen";

export default function Matches() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <View style={{ flex: 1, marginBottom: 35 }}>
        <MatchesScreen />
      </View>
      <Navbar activePage="Matches" />
    </View>
  );
}
