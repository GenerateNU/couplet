import React from "react";
import { Image, Text, View } from "react-native";
import { EventCardProps } from "./EventProps";
import Reaction from "./Reaction";

const logo = {
  uri: "https://reactnative.dev/img/tiny_logo.png",
  width: 250,
  height: 250
};

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
    <View>
      <Reaction handleReact={handleReact} />
      <View
        style={{
          flexGrow: 1
        }}
      >
        <Image source={logo} />
        <Text style={{ fontSize: 72 }}>{title}</Text>
        <Text style={{ fontSize: 24 }}>{date}</Text>
        <Text style={{ fontSize: 18 }}>{price}</Text>

        <Text style={{ fontSize: 48 }}>{description}</Text>
        <Text style={{ fontSize: 48 }}>{location}</Text>
      </View>
    </View>
  );
}
