import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import { useLocalSearchParams, useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { Button } from "react-native-paper";
import { SafeAreaView } from "react-native-safe-area-context";
import EventPage from "../components/Event/EventPage";

export default function Event() {
  const {
    // The route parameter
    eventId
  } = useLocalSearchParams<{ eventId: string }>();
  const router = useRouter();
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <SafeAreaView>
      <View style={styles.buttonContainer}>
        <Button onPress={() => router.back()}>
          <Text style={{ fontFamily: "DMSansRegular" }}>Go Back</Text>
        </Button>
      </View>
      <View style={styles.cardContainer}>
        <EventPage id={eventId || ""} />
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
