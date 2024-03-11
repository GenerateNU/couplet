import React from "react";
import { Text, View } from "react-native";
import { Icon } from "react-native-paper";
import {
  DMSans_400Regular as DMSansRegular,
  DMSans_500Medium as DMSansMedium
} from '@expo-google-fonts/dm-sans';
import { useFonts } from "expo-font";
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
  const [fontsLoaded] = useFonts({
    DMSansRegular, 
    DMSansMedium
  });

  if (!fontsLoaded) {
    return null; 
  }
  return (
    <View
      style={{
        flexGrow: 1,
        marginHorizontal: "10%"
      }}
    >
      <Text style={{ fontSize: 32, marginBottom: 10, fontFamily: "DMSansMedium" }}>{title}</Text>
      <View style={{ flexDirection: "row" }}>
        <Icon source="calendar" size={24} />
        <Text style={{ fontSize: 18 }}>{date}</Text>
      </View>

      <View style={{ flexDirection: "row" }}>
        <Icon source="pin-outline" size={24} />
        <Text style={{ fontSize: 18, marginRight: 24, fontFamily: "DMSansRegular" }}>{location}</Text>
        <Icon source="cash" size={24} />
        <Text style={{ fontSize: 18, marginRight: 20, fontFamily: "DMSansRegular" }}>{`$${price}`}</Text>
      </View>

      <Text style={{ fontSize: 18, marginVertical: 10, fontFamily: "DMSansRegular" }}>{description}</Text>
      <Text style={{ fontSize: 18, marginVertical: 10, fontFamily: "DMSansRegular" }}>{description}</Text>

      <Text style={{ fontSize: 24, fontFamily: "DMSansRegular" }}>Location</Text>
      <Text style={{ fontSize: 18, fontFamily: "DMSansRegular" }}>{location}</Text>
    </View>
    // </ScrollView>
  );
}
