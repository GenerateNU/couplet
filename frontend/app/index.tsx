import React from "react";
import { SafeAreaView } from "react-native";
import HomeScreen from "../components/HomeScreen";
import Navbar from "../components/Navbar";

export default function index() {
  return (
    <SafeAreaView>
      <SafeAreaView>
        <HomeScreen />
      </SafeAreaView>
      <SafeAreaView>
        <Navbar />
      </SafeAreaView>
    </SafeAreaView>
  );
}
