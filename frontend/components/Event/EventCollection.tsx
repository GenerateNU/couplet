import { router } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { getEventById } from "../../api/events";
import scaleStyleSheet from "../../scaleStyles";
import HomeEventCard from "../Home/HomeEventCard";

type Event = Awaited<ReturnType<typeof getEventById>>;

export type EventCollectionProps = {
  name: string;
  events: Event[];
};

export default function EventCollection({ name, events }: EventCollectionProps) {
  return (
    <View style={scaledStyles.container}>
      <Text onPress={() => router.back()} style={scaledStyles.title}>{`< ${name}`}</Text>
      <Text style={scaledStyles.subtitle}>Most Liked</Text>
      <View style={scaledStyles.likedContainer}>
        {events.map((event) => (
          <View style={scaledStyles.likedEvent} key={event.id}>
            <HomeEventCard key={event.id} id={event.id} name={event.name} image={event.images[0]} />
          </View>
        ))}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { marginHorizontal: 24 },
  title: {
    fontFamily: "DMSansMedium",
    fontSize: 20,
    fontWeight: "700",
    lineHeight: 32
  },
  subtitle: { fontFamily: "DMSansMedium", fontSize: 20, fontWeight: "500", marginVertical: 16 },

  likedContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "space-between" // Distribute cards evenly in the row
  },
  likedEvent: {
    width: "45%",
    margin: 5
  }
});

const scaledStyles = scaleStyleSheet(styles);
