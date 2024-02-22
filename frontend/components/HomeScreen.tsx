import React from "react";
import CallToAction from "./CallToAction";
import Header from "./Header";
import { View, Text, Pressable, ScrollView } from "react-native";
import TagButton from "./TagButton";
import HomeEventCard from "./HomeEventCard";
import HomePageSection from "./HomePageSection";

export default function HomeScreen() {
  return <View>
    <View style={{height:"10%"}}>
      <Header />
    </View>
    <View style={{flexDirection: "row", padding:"4%", marginTop: "4%"}}>
      <TagButton text="All Events" />
      <TagButton text="Liked Events" selected/>
    </View>

    <HomePageSection title="This weekend in Boston" events={[1,2,3,4,5]} />
    <HomePageSection title="Live music and concerts" events={[1,2,3]} />
    <HomePageSection title="Other events" events={[1,2,3,4,5]} />

  </View>;
}
