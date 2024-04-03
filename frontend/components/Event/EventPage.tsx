import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import { ScrollView, StyleSheet, Text, View } from "react-native";
import { IconButton } from "react-native-paper";
import { getEventById } from "../../api/events";
import type { components } from "../../api/schema";
import scaleStyleSheet from "../../scaleStyles";
import Reaction from "../Reaction/Reaction";
import EventCard from "./EventCard";
import EventImageCarousel from "./EventImageCarousel";

type Event = components["schemas"]["Event"];

interface EventPageProps {
  id: string;
  handleReact: (like: boolean) => void;
}

export default function EventPage({ id, handleReact }: EventPageProps) {
  const [event, setEvent] = useState<Event>();

  useEffect(() => {
    getEventById(id)
      .then((fetchedEvent) => setEvent(fetchedEvent))
      .catch((e) => console.log(e));
  }, [id]);

  const dummyImages: string[] = [
    "https://marvel-b1-cdn.bc0a.com/f00000000283318/home.dartmouth.edu/sites/home/files/styles/max_width_720px/public/2023-12/20220127_around_campus_eb_157.jpg?itok=bJJ9L7nZ",
    "https://www.lawnstarter.com/blog/wp-content/uploads/2022/12/iStock-1423384637-2-feature-image-1.jpg",
    "https://www.flightonice.com/wp-content/uploads/2022/10/e4d4996c-da07-403e-a1c9-17696615d7ea_750x422.jpg",
    "https://www.novaparks.com/sites/default/files/styles/scale_1440/public/2024-01/IceSkating202312190151_NP.jpg?itok=a6ScPTLd"
  ];

  return (
    <View>
      <View style={scaledStyles.eventContentContainer}>
        <IconButton
          icon="arrow-left-drop-circle"
          style={{ marginLeft: -10 }}
          size={24}
          onPress={() => router.push("/Home")}
        />
        <Text style={{ fontSize: 32, marginBottom: 10, fontFamily: "DMSansMedium" }}>
          {event?.name}
        </Text>
        <ScrollView showsVerticalScrollIndicator={false}>
          <View style={scaledStyles.eventImageContainer}>
            <EventImageCarousel images={dummyImages} />
          </View>
          <View>{event && <EventCard event={event} handleReact={handleReact} />}</View>
        </ScrollView>
      </View>
      <View style={scaledStyles.reactionContainer}>
        <Reaction handleReact={handleReact} />
      </View>
    </View>
  );
}
const styles = StyleSheet.create({
  eventContentContainer: {
    paddingHorizontal: 20,
    height: "100%",
    width: "100%"
  },
  eventImageContainer: {
    marginBottom: 10
  },
  reactionContainer: {
    position: "absolute",
    width: "100%",
    bottom: 0
  },
  viewShare: {
    flexDirection: "row",
    justifyContent: "space-between",
    marginVertical: 10,
    paddingBottom: 50
  },
  buttonLabel: {
    fontFamily: "DMSansMedium",
    fontSize: 16,
    paddingHorizontal: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
