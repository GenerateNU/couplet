import { useLocalSearchParams } from "expo-router";
import React, { useEffect, useState } from "react";
import { SafeAreaView, StyleSheet, View } from "react-native";
import { getEventById, getEvents } from "../api/events";
import EventCollection from "../components/Event/EventCollection";
import Navbar from "../components/Layout/Navbar";
import scaleStyleSheet from "../scaleStyles";

type Event = Awaited<ReturnType<typeof getEventById>>;

export default function Collection() {
  const [events, setEvents] = useState<Event[]>([]);
  const { collectionId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();

  useEffect(() => {
<<<<<<< HEAD
    getEvents({ limit: 10, offset: 10 }).then((event) => {
=======
    getEvents({ limit: 20, offset: 0 }).then((event) => {
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      setEvents(event);
    });
  }, []);

  return (
    <View style={scaledStyles.container}>
      <SafeAreaView style={{ flex: 1 }}>
        <EventCollection name={collectionId || "This Weekend in Boston"} events={events} />
      </SafeAreaView>
      <Navbar activePage="Home" />
    </View>
  );
}

const styles = StyleSheet.create({ container: { flex: 1, justifyContent: "space-between" } });

const scaledStyles = scaleStyleSheet(styles);
