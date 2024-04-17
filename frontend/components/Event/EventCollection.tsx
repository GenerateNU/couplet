import { router } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { ScrollView } from "react-native-gesture-handler";
import { getEventById } from "../../api/events";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import HomeEventCard from "../Home/HomeEventCard";

type Event = Awaited<ReturnType<typeof getEventById>>;

export type EventCollectionProps = {
  name: string;
  events: Event[];
};

export default function EventCollection({ name, events }: EventCollectionProps) {
  return (
    <ScrollView stickyHeaderIndices={[0]} style={scaledStyles.scrollView}>
      <View>
        <Text onPress={() => router.back()} style={scaledStyles.title}>{`< ${name}`}</Text>
      </View>
      <View style={scaledStyles.container}>
        <Text style={scaledStyles.subtitle}>Most Liked</Text>
        <View style={scaledStyles.likedContainer}>
          {events.map((event) => (
            <View style={scaledStyles.likedEvent} key={event.id}>
              <HomeEventCard key={event.id} event={event} />
            </View>
          ))}
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  scrollView: { marginTop: 16, marginBottom: 40 },
  container: { paddingHorizontal: 24 },
  title: {
    fontFamily: "DMSansMedium",
    fontSize: 24,
    fontWeight: "700",
    lineHeight: 32,
    backgroundColor: COLORS.white,
    paddingBottom: 16,
    marginLeft: 24
  },
  subtitle: { fontFamily: "DMSansMedium", fontSize: 20, fontWeight: "500", marginBottom: 16 },
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
