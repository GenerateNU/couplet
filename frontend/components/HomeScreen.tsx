import React from "react";
import CardStack from "./CardStack";
import Header from "./Header";
import { View, Text, Pressable } from "react-native";
import TagButton from "./TagButton";



export default function HomeScreen() {
  return <View>
    <View style={{height:"35%"}}>
      <Header />
    </View>
    <View style={{flexDirection: "row", padding:"4%", marginTop: "4%"}}>
      <TagButton text="All Events" />
      <TagButton text="Liked Events" selected/>
    </View>
  </View>;
}
