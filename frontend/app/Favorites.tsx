import React from "react";
import LikesScreen from "../components/LikesScreen";
import { View } from "react-native";
import Navbar from "../components/Navbar";


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
