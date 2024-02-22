import React from "react";
import { ScrollView, Text, View } from "react-native";
import { EventCardProps } from "./EventProps";
import Reaction from "./Reaction";

export default function EventCard({
  handleReact,
  id,
  title,
  description,
  price,
  date,
  location
}: EventCardProps) {
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
      stickyHeaderIndices={[1]}
    >
      <View
        style={{
          flexGrow: 1
        }}
      >
        <Text style={{ fontSize: 72 }}>{title}</Text>
        <Text style={{ fontSize: 24 }}>{date}</Text>
        <View style={{ flexDirection: "row" }}>
          <Text style={{ fontSize: 18, marginRight: 20 }}>{price}</Text>
          <Text style={{ fontSize: 18 }}>{location}</Text>
        </View>

        <Text style={{ fontSize: 18, marginBottom: 10 }}>{description}</Text>
        <Text style={{ fontSize: 24 }}>Location</Text>

        <Text style={{ fontSize: 18 }}>{location}</Text>
      </View>
      <Reaction handleReact={handleReact} />
    </ScrollView>
  );
}
