import { useRouter } from "expo-router";
import React from "react";
import { Button, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import EventPage from "../components/EventPage";
import { StyleSheet } from "react-native";

export default function DummyEventDetails() {
  const router = useRouter();

  return (
    <SafeAreaView>
      <View style = { styles.buttonContainer}>
      <Button title="Go back" onPress={() => router.back()} />
      </View>
      <View style = { styles.cardContainer}>
      <EventPage />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  buttonContainer : {
    transform: [{ translateY: -50 }]
  },
  cardContainer : {
    transform : [{ translateY : -50}]
  }
})
