import React from "react";
import { View } from "react-native";
import Navbar from "../components/Layout/Navbar";
import LikesScreen from "../components/Matches/LikesScreen";

export default function Favorites() {
  return (
    <View style={{ flex: 1, justifyContent: "space-between" }}>
      <View style={{ flex: 1, marginBottom: 35 }}>
        <LikesScreen />
      </View>
      <Navbar />
    </View>
  );
}
