import React, { useCallback, useEffect, useState } from "react";
import { View } from "react-native";
import { eventSwipe, getAllEvents } from "../../api/events";
import EventPage from "./EventPage";

function CardStack() {
  const [currentCardIndex, setCurrentCardIndex] = useState(0);
  const [events, setEvents] = useState<any[]>([]);
  const [currentEventId, setCurrentEventId] = useState("");

  const handleReact = useCallback(
    (like: boolean) => {
      const userId = "c69626f1-f73d-4045-87d8-40e28f136c62"; // HARDCODED FROM MY DB
      eventSwipe(userId, currentEventId, like).then();
      const nextIndex = (currentCardIndex + 1) % events.length;
      setCurrentCardIndex(nextIndex);
      setCurrentEventId(events[nextIndex]?.id);
    },
    [events, currentCardIndex, currentEventId]
  );

  useEffect(() => {
    getAllEvents().then((fetchedEvents: any) => {
      setEvents(fetchedEvents || []);
      console.log(fetchedEvents);
      setCurrentEventId(fetchedEvents[0].id);
    });
  }, []);

  return (
    <View>
      <EventPage id={currentEventId} handleReact={handleReact} />
    </View>
  );
}

export default CardStack;
