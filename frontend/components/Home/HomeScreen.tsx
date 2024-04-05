import React, { useCallback, useEffect, useState } from "react";
import { ScrollView, StyleSheet, View } from "react-native";
import { getEvents } from "../../api/events";
import scaleStyleSheet from "../../scaleStyles";
import LabelToggle from "../LabelToggle";
import Header from "../Layout/Header";
import HomePageSection from "./HomePageSection";

// type Event = components["schemas"]["Event"];

const toggles = ["All events", "Liked events"];
const sections = [
  "This weekend in Boston",
  "Food & Drink",
  "Arts & Culture",
  "Nightlife",
  "Live Music & Concerts",
  "Nature & Outdoors"
];
const MemoizedHomePageSection = React.memo(HomePageSection);

export default function HomeScreen() {
  const [filter, setFilter] = useState(toggles[0]);
  const [events, setEvents] = useState<any[]>([]);

  const setFilterLikedEvents = useCallback((newFilter: string) => {
    setFilter(newFilter);
  }, []);

  useEffect(() => {
    getEvents({ limit: 20, offset: 0 })
      .then((fetchedEvents: any) => {
        setEvents(fetchedEvents || []);
      })
      .catch((e) => console.log(e));
  }, []);

  return (
    <ScrollView stickyHeaderIndices={[0]} style={scaledStyles.scrollView}>
      <Header />
      <View style={scaledStyles.toggleContainer}>
        <LabelToggle labels={toggles} onChange={setFilterLikedEvents} />
      </View>
      <View style={scaledStyles.matchContainer} />
      {filter === "Liked events" ? (
        <View />
      ) : (
        <View style={scaledStyles.sectionContainer}>
          {sections.map((section) => (
            <MemoizedHomePageSection key={section} title={section} events={events} />
          ))}
        </View>
      )}
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  toggleContainer: {
    display: "flex",
    flexDirection: "row",
    paddingBottom: 16
  },
  matchContainer: {
    flexDirection: "column",
    alignItems: "flex-start",
    borderStyle: "solid",
    borderWidth: 1,
    flex: 1,
    backgroundColor: "grey",
    padding: 20,
    marginRight: 24
  },
  image: {
    borderRadius: 50,
    borderWidth: 1,
    paddingBottom: 30
  },
  imageContainer: {
    flexDirection: "row",
    paddingBottom: 10
  },
  scrollView: {
    marginBottom: 40,
    paddingLeft: 24
  }
});

const scaledStyles = scaleStyleSheet(styles);
