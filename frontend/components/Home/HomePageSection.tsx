import { router } from "expo-router";
import React from "react";
import { ScrollView, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
<<<<<<< HEAD
=======
import { getEventById } from "../../api/events";
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
import scaleStyleSheet from "../../scaleStyles";
import HomeEventCard from "./HomeEventCard";

type Event = Awaited<ReturnType<typeof getEventById>>;

export default function HomePageSection({ title, events }: { title: string; events: Event[] }) {
  return (
    <View style={{ marginTop: 32 }}>
      <TouchableOpacity
        onPress={() =>
          router.push({
            pathname: "Collection",
            params: { collectionId: "" }
          })
        }
      >
        <Text style={scaledStyles.sectionTitle}>{`${title} >`} </Text>
      </TouchableOpacity>
      <View style={{ flexDirection: "row" }}>
        <ScrollView horizontal>
          {events.map((event) => (
<<<<<<< HEAD
            <HomeEventCard key={event.id} id={event.id} name={event.name} image={event.images[0]} />
=======
            <HomeEventCard key={event.id} event={event} />
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
          ))}
        </ScrollView>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  sectionTitle: {
    fontSize: 20,
    fontWeight: "700",
    fontFamily: "DMSansMedium",
    letterSpacing: -0.2,
    marginBottom: 8
  }
});

const scaledStyles = scaleStyleSheet(styles);
