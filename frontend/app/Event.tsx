import { useLocalSearchParams, useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { Button } from "react-native-paper";
import { SafeAreaView } from "react-native-safe-area-context";
import CardStack from "../components/Event/CardStack";

export default function Event() {
  const { collectionId, eventId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();
  const router = useRouter();

  // I'm thinking a collectionID can be how we separate events into HomePageSections (rows)
  // example "This weekend in Boston"
  // I don't think this is a concept we have yet though
  // ultimately we will probably want to pass collectionID to the CardStack so it can fetch that collection's items
  // for now I'm logging it to make the linter happy
  console.log(collectionId);

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
