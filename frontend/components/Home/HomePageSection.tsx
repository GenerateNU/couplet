import React from "react";
import { ScrollView, Text, View } from "react-native";
import HomeEventCard from "./HomeEventCard";

export default function HomePageSection({ title, events }: { title: string; events: any[] }) {
  return (
    <View style={{ marginVertical: 10 }}>
      <Text style={{ fontSize: 20, fontFamily: "DMSansRegular" }}>{title} </Text>
      <View style={{ flexDirection: "row" }}>
        <ScrollView horizontal>
          {events.map((event) => (
            <HomeEventCard key={event.id} id={event.id} name={event.name} />
          ))}
        </ScrollView>
      </View>
    </View>
  );
}
