import React, { useCallback, useEffect, useState } from "react";
import { ActivityIndicator, View } from "react-native";
import { eventSwipe, getAllEvents } from "../../api/events";
import EventPage from "./EventPage";

export type CardStackProps = {
  startingEventId: string;
};

export default function CardStack({ startingEventId }: CardStackProps) {
  const [events, setEvents] = useState<any[]>([]);
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [isLoading, setIsLoading] = useState(true);

  const handleReact = useCallback(
    (like: boolean) => {
      const userId = "c69626f1-f73d-4045-87d8-40e28f136c62"; // HARDCODED FROM MY DB. TODO: switch to logged-in user
      const currentEventId = events[currentCardIndex]?.id;
      eventSwipe(userId, currentEventId, like).then();
      const nextIndex = (currentCardIndex + 1) % events.length;
      setCurrentCardIndex(nextIndex);
    },
    [events, currentCardIndex]
  );

  useEffect(() => {
    getAllEvents().then((fetchedEvents: any) => {
      setEvents(fetchedEvents || []);
      const index = fetchedEvents.findIndex((event: any) => event.id === startingEventId);

      if (index !== -1) {
        setCurrentCardIndex(index);
      } else {
        console.log(`No event found with ID ${startingEventId}`);
        setCurrentCardIndex(0);
      }

      setIsLoading(false);
    });
  }, [startingEventId]);

  if (isLoading) {
    return (
      <View style={{ flex: 1, justifyContent: "center", alignItems: "center" }}>
        <ActivityIndicator size="large" />
      </View>
    );
  }

  return (
    <View>
      <EventPage id={events[currentCardIndex]?.id} handleReact={handleReact} />
    </View>
  );
}
