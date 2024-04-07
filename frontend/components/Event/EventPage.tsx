import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import { ScrollView, StyleSheet, Text, View } from "react-native";
import { IconButton } from "react-native-paper";
import { getEventById } from "../../api/events";
import scaleStyleSheet from "../../scaleStyles";
import Reaction from "../Reaction/Reaction";
import EventCard from "./EventCard";
import EventImageCarousel from "./EventImageCarousel";

type Event = Awaited<ReturnType<typeof getEventById>>;

type EventPageProps = {
  id: string;
  handleReact: (like: boolean) => void;
};

export default function EventPage({ id, handleReact }: EventPageProps) {
  const [event, setEvent] = useState<Event>();

  useEffect(() => {
    getEventById({ id }).then((fetchedEvent) => {
      console.log(fetchedEvent);
      setEvent(fetchedEvent);
    });
  }, [id]);

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
            <EventImageCarousel images={event?.images || []} />
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
