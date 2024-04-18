<<<<<<< HEAD
import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import { useLocalSearchParams, useRouter } from "expo-router";
=======
import { useLocalSearchParams } from "expo-router";
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
import React from "react";
import { SafeAreaView } from "react-native";
import CardStack from "../components/Event/CardStack";

export default function Event() {
<<<<<<< HEAD
  const { collectionId } = useLocalSearchParams<{ collectionId: string }>();
  const router = useRouter();
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  // I'm thinking a collectionID can be how we separate events into HomePageSections (rows)
  // example "This weekend in Boston"
  // I don't think this is a concept we have yet though
  // ultimately we will probably want to pass collectionID to the CardStack so it can fetch that collection's items
  // for now I'm logging it to make the linter happy
  console.log(collectionId);
=======
  const { eventId } = useLocalSearchParams<{
    collectionId: string;
    eventId: string;
  }>();
  // TODO: I think we need a notion of collectionId, which can be how we separate events into HomePageSections (rows)
  // We probably want to pass collectionId to the CardStack so it can fetch that collection's items
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb

  if (!fontsLoaded) {
    return null;
  }

  return (
    <SafeAreaView>
<<<<<<< HEAD
      <View style={styles.buttonContainer}>
        <Button onPress={() => router.back()}>
          <Text style={{ fontFamily: "DMSansRegular" }}>Go Back</Text>
        </Button>
      </View>
      <View style={styles.cardContainer}>
        <CardStack startingEventId="" />
      </View>
=======
      <CardStack startingEventId={eventId || ""} />
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
    </SafeAreaView>
  );
}

// const styles = StyleSheet.create({});
