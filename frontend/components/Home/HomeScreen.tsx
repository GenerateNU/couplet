import React, { useEffect, useState } from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { ScrollView } from "react-native-gesture-handler";
import { getEvents } from "../../api/events";
import Header from "../Layout/Header";
import LinkButton from "../Layout/LinkButton";
import TagButton from "../Layout/TagButton";
import HomePageSection from "./HomePageSection";

const DUMMY_IMAGE = require("../../assets/blankProfile.jpg");

export default function HomeScreen() {
  const [filter, setFilter] = useState(0);
  const [events, setEvents] = useState<Event[]>([]);

  useEffect(() => {
    getEvents().then((fetchedEvents: any) => {
      setEvents(fetchedEvents || []);
    });
  }, []);

  return (
    <ScrollView stickyHeaderIndices={[0]} style={styles.scrollView}>
      {/* Header View */}
      <View>
        <Header />
        <View
          style={{
            display: "flex",
            flexDirection: "row",
            paddingBottom: 10
          }}
        >
          <TagButton text="All Events" selected={filter === 0} onPress={() => setFilter(0)} />
          <TagButton text="Liked Events" selected={filter === 1} onPress={() => setFilter(1)} />
        </View>
      </View>

      {/* Match Container View */}
      <View style={styles.matchContainer}>
        <View style={styles.imageContainer}>
          <Image source={DUMMY_IMAGE} style={styles.image} />
          <Image source={DUMMY_IMAGE} style={styles.image} />
          <Image source={DUMMY_IMAGE} style={styles.image} />
          <Image source={DUMMY_IMAGE} style={styles.image} />
        </View>
        <Text style={{ fontFamily: "DMSansRegular" }}>Need Someone to Go With?</Text>
        <View style={styles.buttonContainer}>
          <LinkButton text="Match Now" />
        </View>
      </View>

      {/* Pintrestesque Section Views */}
      {filter === 0 ? (
        <View style={styles.sectionContainer}>
          <HomePageSection title="This weekend in Boston" events={events} />
          <HomePageSection title="Live music and concerts" events={events} />
          <HomePageSection title="Other events" events={events} />
        </View>
      ) : (
        // Eventually replace this with a different event browsing screen
        <View style={styles.sectionContainer}>
          <HomePageSection
            title="This weekend in Boston"
            events={[1, 2, 3, 4, 5].map((n) => ({ id: n }))}
          />
          <HomePageSection
            title="Live music and concerts"
            events={[1, 2, 3].map((n) => ({ id: n }))}
          />
          <HomePageSection title="Other events" events={[1, 2, 3, 4, 5].map((n) => ({ id: n }))} />
        </View>
      )}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  matchContainer: {
    flexDirection: "column",
    alignItems: "flex-start",
    borderStyle: "solid",
    borderWidth: 1,
    flex: 1,
    backgroundColor: "grey",
    padding: 20,
    margin: 10
  },
  image: {
    height: 50,
    width: 50,
    borderRadius: 50,
    borderWidth: 1,
    paddingBottom: 30
  },
  imageContainer: {
    flexDirection: "row",
    paddingBottom: 10
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    paddingTop: 10
  },
  sectionContainer: {
    margin: 10
  },
  scrollableSection: {
    marginVertical: 20
  },
  scrollView: {
    marginBottom: 40
  }
});
