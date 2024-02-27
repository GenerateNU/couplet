import React from "react";
import { ScrollView, View } from "react-native";
import EventCard from "./EventCard";
import EventImageCarousel from "./EventImageCarousel";
import Reaction from "./Reaction";
import { StyleSheet } from "react-native";

function EventPage() {
  const dummyImages: string[] = [
    "https://marvel-b1-cdn.bc0a.com/f00000000283318/home.dartmouth.edu/sites/home/files/styles/max_width_720px/public/2023-12/20220127_around_campus_eb_157.jpg?itok=bJJ9L7nZ",
    "https://www.lawnstarter.com/blog/wp-content/uploads/2022/12/iStock-1423384637-2-feature-image-1.jpg",
    "https://www.flightonice.com/wp-content/uploads/2022/10/e4d4996c-da07-403e-a1c9-17696615d7ea_750x422.jpg",
    "https://www.novaparks.com/sites/default/files/styles/scale_1440/public/2024-01/IceSkating202312190151_NP.jpg?itok=a6ScPTLd"
  ];
  const handleReact = () => {
    // TODO: Add logic to update the database with the user's reaction to the event
  };

  return (
    <View>
      <ScrollView
        style={{
          width: "100%",
          alignSelf: "center",
          height: "90%",
        }}
        contentContainerStyle={{}}
      >
        <View >
        <EventImageCarousel images={dummyImages} />
        </View>
        <View style = { styles.EventCardContainer}>
        <EventCard
          id={1}
          title="Winter Ice Skating"
          description="The best place to skate outdoors in Boston. Whether you are a first-time skater or an aspiring Olympian, Frog Pond enables you to take pleasure in the sport of ice skating. Frog Pond offers public ice skating, skate rentals, and skate sharpening."
          date="1/2/24"
          price={20}
          location="Frog Pond"
          handleReact={handleReact}
        />
        </View>
      </ScrollView>
      <Reaction handleReact={handleReact} />
    </View>
  );
}
const styles = StyleSheet.create({
  EventCardContainer : {
    flex : 1,
    transform : [{ translateY : 100}]
  }
})
export default EventPage;
