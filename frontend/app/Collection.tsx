import { useLocalSearchParams } from "expo-router";
import React, { useEffect, useState } from "react";
import { SafeAreaView } from "react-native";
import { getEventById, getEvents } from "../api/events";
import EventCollection from "../components/Event/EventCollection";

type Event = Awaited<ReturnType<typeof getEventById>>;

export default function Collection() {
  const [events, setEvents] = useState<Event[]>([]);
  const { collectionId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();

  useEffect(() => {
    getEvents({ limit: 10, offset: 10 }).then((event) => {
      setEvents(event);
    });
  }, []);

  return (
    <SafeAreaView>
      <EventCollection name={collectionId || "This Weekend in Boston"} events={events} />
    </SafeAreaView>
  );
}

// const styles = StyleSheet.create({});
