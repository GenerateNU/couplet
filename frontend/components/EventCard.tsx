import React from "react";
import { Image, ScrollView, Text, View } from "react-native";
import { EventCardProps } from "./EventProps";
import Reaction from "./Reaction";

const logo = {
  uri: "https://reactnative.dev/img/tiny_logo.png",
  width: 250,
  height: 250
};

const EventCard: React.FC<EventCardProps> = (props) => {
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
      <Reaction handleReact={props.handleReact} />
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
        ></Text>
        <Text style={{ fontSize: 96 }}>{props.id}</Text>
        <Image source={logo} />
        <Text style={{ fontSize: 96 }}>{props.title}</Text>
        <Text style={{ fontSize: 48 }}>{props.description}</Text>
        <Text style={{ fontSize: 48 }}>{props.date}</Text>
        <Text style={{ fontSize: 48 }}>{props.location}</Text>
      </View>
    </ScrollView>
  );
};

export default EventCard;
