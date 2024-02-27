import React from "react";
import { View } from "react-native";
import MatchesScreen from "../components/MatchesScreen";
import Navbar from "../components/Navbar";

export default function Matches() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <View style={{ flex: 1, marginBottom: 35 }}>
        <MatchesScreen />
      </View>
      <Navbar />
    </View>
  );
}
