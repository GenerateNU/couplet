import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { getEventById } from "../../api/events";
import scaleStyleSheet from "../../scaleStyles";

type Event = Awaited<ReturnType<typeof getEventById>>;

const FEATURED = require("../../assets/homeScreenBackground.png");

type FeaturedEventCardProps = {
  event: Event;
};

export default function FeaturedEventCard({ event }: FeaturedEventCardProps) {
  return (
    <View>
      <Text style={scaledStyles.header}>Don’t miss out on these new events!</Text>

      <TouchableOpacity
        onPress={() =>
          router.push({ pathname: "Event", params: { collectionId: "", eventId: event.id } })
        }
      >
        <View style={styles.card}>
          <Image source={FEATURED} style={styles.featuredImage} />
          <View style={styles.absoluteImage}>
            <Image source={{ uri: event.images[0] }} style={styles.image} />
            <View style={styles.textContainer}>
              <Text style={styles.titleText}>{event.name}</Text>
              <Text numberOfLines={2} ellipsizeMode="tail" style={styles.text}>
                {event.bio}
              </Text>
            </View>
          </View>
        </View>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  absoluteImage: {
    position: "absolute",
    zIndex: 1,
    top: 30,
    left: 100
  },
  card: {
    width: 200,
    height: 255,
    borderRadius: 8,
    shadowColor: "#000000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    marginBottom: 40
  },
  textContainer: {
    backgroundColor: "white",
    borderRadius: 8,
    height: 100
  },
  image: {
    width: "100%",
    height: 150,
    borderTopLeftRadius: 10,
    borderTopRightRadius: 10
  },
  featuredImage: {
    width: 400,
    height: 350
  },
  titleText: {
    padding: 10,
    fontSize: 15,
    fontWeight: "500",
    fontFamily: "DMSansMedium"
  },
  text: {
    padding: 10,
    fontSize: 15,
    fontFamily: "DMSansRegular"
  },
  header: {
    fontSize: 17,
    fontFamily: "DMSansMedium",
    fontWeight: "500",
    textAlign: "center",
    marginLeft: 20,
    marginVertical: 10
  }
});

const scaledStyles = scaleStyleSheet(styles);
