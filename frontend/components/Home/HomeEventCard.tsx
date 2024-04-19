import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { Icon } from "react-native-paper";
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
          <Text style={styles.titleText} numberOfLines={1}>
            {event.name.length < 20 ? event.name : `${event.name.substring(0, 17)}...`}
          </Text>
          <View style={styles.row}>
            <Icon source="map-marker" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text} numberOfLines={1}>
              {event.address.length < 20 ? event.address : `${event.address.substring(0, 17)}...`}
            </Text>
          </View>
          <View style={styles.row}>
            <Icon source="currency-usd" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>${event.minPrice}</Text>
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
    alignItems: "center",
    paddingHorizontal: 10,
    borderRadius: 20
  },
  titleText: { padding: 10, fontSize: 15, fontWeight: "500", fontFamily: "DMSansMedium" },
  text: {
    marginTop: 2,
    marginLeft: 2,
    fontFamily: "DMSansRegular",
    fontSize: 12,
    fontWeight: "400"
  }
});

const scaledStyles = scaleStyleSheet(styles);
