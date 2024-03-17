import React from "react";
import { ScrollView, StyleSheet, View } from "react-native";
import Reaction from "../Reaction/Reaction";
import EventCard from "./EventCard";
import EventImageCarousel from "./EventImageCarousel";

interface EventPageProps {
  id: string;
  handleReact: (like: boolean) => void;
}

function EventPage({ id, handleReact }: EventPageProps) {
  const dummyImages: string[] = [
    "https://marvel-b1-cdn.bc0a.com/f00000000283318/home.dartmouth.edu/sites/home/files/styles/max_width_720px/public/2023-12/20220127_around_campus_eb_157.jpg?itok=bJJ9L7nZ",
    "https://www.lawnstarter.com/blog/wp-content/uploads/2022/12/iStock-1423384637-2-feature-image-1.jpg",
    "https://www.flightonice.com/wp-content/uploads/2022/10/e4d4996c-da07-403e-a1c9-17696615d7ea_750x422.jpg",
    "https://www.novaparks.com/sites/default/files/styles/scale_1440/public/2024-01/IceSkating202312190151_NP.jpg?itok=a6ScPTLd"
  ];

  return (
    <View>
      <ScrollView
        style={{
          width: "100%",
          alignSelf: "center",
          height: "90%"
        }}
        contentContainerStyle={{}}
      >
        <View>
          <EventImageCarousel images={dummyImages} />
        </View>
        <View style={styles.EventCardContainer}>
          <EventCard id={id} handleReact={handleReact} />
        </View>
      </ScrollView>
      <Reaction handleReact={handleReact} />
    </View>
  );
}
const styles = StyleSheet.create({
  EventCardContainer: {
    flex: 1,
    transform: [{ translateY: 100 }]
  }
});
export default EventPage;
