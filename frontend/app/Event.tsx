import { useLocalSearchParams } from "expo-router";
import React from "react";
import { SafeAreaView } from "react-native";
import CardStack from "../components/Event/CardStack";

export default function Event() {
  const { eventId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();
  // TODO: I think we need a notion of collectionId, which can be how we separate events into HomePageSections (rows)
  // We probably want to pass collectionId to the CardStack so it can fetch that collection's items

  if (!fontsLoaded) {
    return null;
  }

  return (
    <SafeAreaView>
      <CardStack startingEventId={eventId || ""} />
    </SafeAreaView>
  );
}

// const styles = StyleSheet.create({});
