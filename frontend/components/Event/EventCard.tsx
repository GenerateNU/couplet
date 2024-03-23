import React, { useEffect, useState } from "react";
import { Text, View } from "react-native";
import { Icon } from "react-native-paper";
import { getEventById } from "../../api/events";
import type { components } from "../../api/schema";

type Event = components["schemas"]["Event"];

export type EventCardProps = {
  id: string;
  handleReact: (like: boolean) => void;
};

export default function EventCard({ handleReact, id }: EventCardProps) {
  const [event, setEvent] = useState<Event>();
  useEffect(() => {
    if (id) {
      getEventById(id)
        .then((fetchedEvent) => {
          if (Array.isArray(fetchedEvent)) {
            console.log("Unexpected never[] array returned");
          } else {
            setEvent(fetchedEvent);
          }
        })
        .catch((e) => {
          console.log(e);
        });
    }
  }, [id]);

  return (
    <View
      style={{
        flexGrow: 1,
        marginHorizontal: "10%",
        paddingTop: 10
      }}
    >
      <Text style={{ fontSize: 32, marginBottom: 10, fontFamily: "DMSansMedium" }}>
        {event?.name}
      </Text>
      <View style={{ flexDirection: "row" }}>
        <Icon source="calendar" size={24} />
        <Text style={{ fontSize: 18 }}>DATE</Text>
      </View>

      <View style={{ flexDirection: "row" }}>
        <Icon source="pin-outline" size={24} />
        <Text style={{ fontSize: 18, marginRight: 24, fontFamily: "DMSansRegular" }}>Boston</Text>
        <Icon source="cash" size={24} />
        <Text style={{ fontSize: 18, marginRight: 20, fontFamily: "DMSansRegular" }}>$20</Text>
      </View>

      <Text style={{ fontSize: 18, marginVertical: 10, fontFamily: "DMSansRegular" }}>
        {event?.bio}
      </Text>
    </View>
  );
}
