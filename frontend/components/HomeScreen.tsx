import React from "react";
import { View } from "react-native";
import CallToAction from "./CallToAction";
import Header from "./Header";
import TagButton from "./TagButton";

export default function HomeScreen() {
  return (
    <View>
      <View style={{ height: 80 }}>
        <Header />
      </View>
      <View style={{ flexDirection: "row", padding: 20 }}>
        <TagButton text="All Events" />
        <TagButton text="Liked Events" />
      </View>
      <CallToAction />
    </View>
  );
}
