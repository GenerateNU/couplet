import React from "react";
import { View } from "react-native";
import HomeScreen from "../components/HomeScreen";
import Navbar from "../components/Navbar";

export default function index() {
  return (
    <View>
      <HomeScreen />
      <View>
        <Navbar/>
      </View>
    </View>
  );
}
