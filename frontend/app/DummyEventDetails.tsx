import { useRouter } from "expo-router";
import React from "react";
import { Button } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import EventPage from "../components/EventPage";

export default function DummyEventDetails() {
  const router = useRouter();

  return (
    <SafeAreaView>
      <Button title="Go back" onPress={() => router.back()}/>
      <EventPage />
    </SafeAreaView>
  );
}
