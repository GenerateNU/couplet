import { useLocalSearchParams, useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { Button } from "react-native-paper";
import { SafeAreaView } from "react-native-safe-area-context";
import CardStack from "../components/Event/CardStack";

export default function Event() {
  const { eventId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();
  const router = useRouter();

  // TODO: I think we need a notion of collectionId, which can be how we separate events into HomePageSections (rows)
  // We probably want to pass collectionId to the CardStack so it can fetch that collection's items
  // console.log(collectionId);

  return (
    <SafeAreaView>
      <View style={styles.buttonContainer}>
        <Button onPress={() => router.back()}>
          <Text style={{ fontFamily: "DMSansRegular" }}>Go Back</Text>
        </Button>
      </View>
      <View style={styles.cardContainer}>
        <CardStack startingEventId={eventId || ""} />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  buttonContainer: {
    transform: [{ translateY: -50 }]
  },
  cardContainer: {
    transform: [{ translateY: -50 }]
  }
});
