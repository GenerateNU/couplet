import React from "react";
import { Image, ScrollView, Text, View } from "react-native";
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
      stickyHeaderIndices={[0]}
    >
      <Reaction handleReact={handleReact} />
      <View
        style={{
          flexGrow: 1,
          justifyContent: "center",
          alignItems: "center"
        }}
      >
        <Text
          style={{
            fontSize: 96,
            paddingTop: "10%"
          }}
        />
        <Text style={{ fontSize: 96 }}>{id}</Text>
        <Image source={logo} />
        <Text style={{ fontSize: 96 }}>{title}</Text>
        <Text style={{ fontSize: 48 }}>{description}</Text>
        <Text style={{ fontSize: 48 }}>{date}</Text>
        <Text style={{ fontSize: 48 }}>{location}</Text>
      </View>
    </ScrollView>
  );
}
