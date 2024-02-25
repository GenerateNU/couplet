import React from "react";
import { Text, View } from "react-native";
import Navbar from "../components/Navbar";

export default function Home() {
  console.log("HEY")
  return (
    <View style={{ flex: 1,
    justifyContent: 'space-between'}}>
      <Text>Home</Text>
      <Navbar/>
    </View>
  );
}
