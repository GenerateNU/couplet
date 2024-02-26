import React from "react";
import { Text, View } from "react-native";
import { Icon } from "react-native-paper";
import { EventCardProps } from "./EventProps";

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
    <View
      style={{
        flexGrow: 1,
        marginHorizontal: "10%"
      }}
    >
      <Text style={{ fontSize: 32, marginBottom: 10 }}>{title}</Text>
      <View style={{ flexDirection: "row" }}>
        <Icon source="calendar" size={24} />
        <Text style={{ fontSize: 18 }}>{date}</Text>
      </View>

      <View style={{ flexDirection: "row" }}>
        <Icon source="pin-outline" size={24} />
        <Text style={{ fontSize: 18, marginRight: 24 }}>{location}</Text>
        <Icon source="cash" size={24} />
        <Text style={{ fontSize: 18, marginRight: 20 }}>{`$${price}`}</Text>
      </View>

      <Text style={{ fontSize: 18, marginVertical: 10 }}>{description}</Text>
      <Text style={{ fontSize: 18, marginVertical: 10 }}>{description}</Text>

      <Text style={{ fontSize: 24 }}>Location</Text>
      <Text style={{ fontSize: 18 }}>{location}</Text>
    </View>
    // </ScrollView>
  );
}
