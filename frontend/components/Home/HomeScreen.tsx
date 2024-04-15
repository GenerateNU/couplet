import React, { useCallback, useEffect, useState } from "react";
import { ScrollView, StyleSheet, View } from "react-native";
import { getEventById, getEvents } from "../../api/events";
import scaleStyleSheet from "../../scaleStyles";
import LabelToggle from "../LabelToggle";
import Header from "../Layout/Header";
import FeaturedEventCard from "./FeaturedEventCard";
import HomePageSection from "./HomePageSection";
import NoLikedEvents from "./NoLikedEvents";

type Event = Awaited<ReturnType<typeof getEventById>>;

const toggles = ["All Events", "Liked Events"];
const sections = [
  "This Weekend in Boston",
  "Food & Drink",
  "Arts & Culture",
  "Nightlife",
  "Live Music & Concerts",
  "Nature & Outdoors"
];

export default function HomeScreen() {
  const [filter, setFilter] = useState(toggles[0]);
  const [events, setEvents] = useState<Event[]>([]);
  const [featuredEvent, setFeaturedEvent] = useState<Event>();

  const setFilterLikedEvents = useCallback((newFilter: string) => {
    setFilter(newFilter);
  }, []);

  useEffect(() => {
    if (!events.length) {
      // only fetch events if they haven't been fetched yet
      getEvents({ limit: 10, offset: 0 })
        .then((fetchedEvents: any) => {
          setEvents(fetchedEvents || []);
          setFeaturedEvent(fetchedEvents[0]);
        })
        .catch((e) => console.log(e));
    }
  }, [events]);

  return (
    <ScrollView stickyHeaderIndices={[0]} style={scaledStyles.scrollView}>
      <Header />
      <View style={scaledStyles.toggleContainer}>
        <LabelToggle labels={toggles} onChange={setFilterLikedEvents} />
      </View>
      {filter === "Liked Events" ? (
        <NoLikedEvents onClick={setFilterLikedEvents} />
      ) : (
        <View style={scaledStyles.sectionContainer}>
          {featuredEvent && (
            <View style={scaledStyles.featuredContainer}>
              <FeaturedEventCard event={featuredEvent} />
            </View>
          )}
          {sections.map((section) => (
            <HomePageSection key={section} title={section} events={events} />
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
  featuredContainer: {
    marginRight: 48
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
  },
  noMatchesImage: {
    width: 300,
    height: 300
  }
});

const scaledStyles = scaleStyleSheet(styles);
