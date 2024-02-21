import React from "react";
import Header from "./Header";
import { View, Text } from "react-native";
import TagButton from "./TagButton";



export default function HomeScreen() {
  return <View>
    <View style={{height:80}}>
      <Header />
    </View>
    <View style={{flexDirection: "row", padding:20}}>
      <TagButton text="All Events" />
      <TagButton text="Liked Events" />
    </View>
    <View style={{padding: 50, margin: 25, borderStyle: "solid", borderWidth: 1, backgroundColor: "gray"}}>
      <Text>Need someone to go with?</Text>
      
    </View>
  </View>;
}
