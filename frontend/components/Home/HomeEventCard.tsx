import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { Icon } from "react-native-paper";
<<<<<<< HEAD
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type HomeEventCardProps = {
  id: string;
  name: string;
  image: string;
  // TODO: we need Location and Cost, but these aren't in the endpoint response yet
};

export default function HomeEventCard({ id, name, image }: HomeEventCardProps) {
  return (
    <TouchableOpacity
      onPress={() => router.push({ pathname: "Event", params: { collectionId: "", eventId: id } })}
    >
      <View style={styles.card}>
        <View style={styles.imageContainer}>
          <Image source={{ uri: image }} style={scaledStyles.image} />
        </View>
        <View style={scaledStyles.textContainer}>
          <Text style={scaledStyles.titleText} numberOfLines={2} ellipsizeMode="tail">
            {name}
          </Text>
          <View style={styles.row}>
            <Icon source="map-marker" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>Frog Pond</Text>
          </View>
          <View style={styles.row}>
            <Icon source="currency-usd" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>$5</Text>
=======
import { getEventById } from "../../api/events";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Event = Awaited<ReturnType<typeof getEventById>>;

type HomeEventCardProps = {
  event: Event;
};

export default function HomeEventCard({ event }: HomeEventCardProps) {
  return (
    <TouchableOpacity
      onPress={() =>
        router.push({ pathname: "Event", params: { collectionId: "", eventId: event.id } })
      }
    >
      <View style={styles.card}>
        <View style={styles.imageContainer}>
          <Image source={{ uri: event.images[0] }} style={scaledStyles.image} />
        </View>
        <View style={scaledStyles.textContainer}>
          <Text style={scaledStyles.titleText} numberOfLines={2} ellipsizeMode="tail">
            {event.name}
          </Text>
          <View style={styles.row}>
            <Icon source="map-marker" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>{event.address}</Text>
          </View>
          <View style={styles.row}>
            <Icon source="currency-usd" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>${event.minPrice}</Text>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
          </View>
        </View>
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  card: {
    width: 166,
    marginRight: 10,
    paddingBottom: 5,
    marginBottom: 10,
    backgroundColor: "#fff",
    borderRadius: 8,
    shadowColor: "#000000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4
  },
  imageContainer: {
    width: "100%",
    height: 150,
    backgroundColor: "rgb(200,200,200)",
    borderTopLeftRadius: 10,
    borderTopRightRadius: 10
  },
  textContainer: {
    height: 96
  },
  image: { width: 166, height: 150, borderTopLeftRadius: 10, borderTopRightRadius: 10 },
  row: {
    flexDirection: "row",
<<<<<<< HEAD
=======
    alignItems: "center",
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
    paddingHorizontal: 10,
    borderRadius: 20
  },
  titleText: { padding: 10, fontSize: 15, fontWeight: "500", fontFamily: "DMSansMedium" },
  text: {
    marginTop: 2,
<<<<<<< HEAD
=======
    marginLeft: 2,
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
    fontFamily: "DMSansRegular",
    fontSize: 12,
    fontWeight: "400"
  }
});

const scaledStyles = scaleStyleSheet(styles);
