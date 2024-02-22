import React from "react";
import { Image, ScrollView } from "react-native";
import EventCard from "./EventCard";

function EventPage() {
  const logo = {
    uri: "https://reactnative.dev/img/tiny_logo.png",
    width: 250,
    height: 250
  };
  const handleReact = () => {
    // TODO: Add logic to update the database with the user's reaction to the event
  };

  return (
    <ScrollView
      style={{
        width: "100%",
        borderRadius: 12,
        alignSelf: "center",
        marginBottom: 10,
        minHeight: "100%"
      }}
      contentContainerStyle={{}}
      stickyHeaderIndices={[0]}
    >
      <Image source={logo} />

      <EventCard
        id={1}
        title="Winter Ice Skating"
        description="The best place to skate outdoors in Boston. Whether you are a first-time skater or an aspiring Olympian, Frog Pond enables you to take pleasure in the sport of ice skating. Frog Pond offers public ice skating, skate rentals, and skate sharpening."
        date="1/2/24"
        price={20}
        location="Frog Pond"
        handleReact={handleReact}
      />
    </ScrollView>
  );
}

export default EventPage;
