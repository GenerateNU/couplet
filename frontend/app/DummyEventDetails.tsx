import { useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { Button } from "react-native-paper";
import { SafeAreaView } from "react-native-safe-area-context";
import EventPage from "../components/Event/EventPage";

export default function DummyEventDetails() {
  const router = useRouter();

  return (
    <SafeAreaView>
      <View style={styles.buttonContainer}>
        <Button onPress={() => router.back()}>
          <Text style={{ fontFamily: "DMSansRegular" }}>Go Back</Text>
        </Button>
      </View>
      <View style={styles.cardContainer}>
        <EventPage />
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
