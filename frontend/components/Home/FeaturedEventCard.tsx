import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { Button } from "react-native-paper";
import { getEventById } from "../../api/events";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Event = Awaited<ReturnType<typeof getEventById>>;

const FEATURED = require("../../assets/homeScreenBackground.png");

type FeaturedEventCardProps = {
  event: Event;
};

export default function FeaturedEventCard({ event }: FeaturedEventCardProps) {
  return (
    <View style={{ flex: 1, alignItems: "center" }}>
      <Text style={scaledStyles.header}>Don’t miss out on these new events!</Text>
      <TouchableOpacity
        onPress={() =>
          router.push({ pathname: "Event", params: { collectionId: "", eventId: event.id } })
        }
      >
        <View>
          <Image source={FEATURED} style={styles.backgroundImage} />
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
      <Button
        mode="contained"
        buttonColor={COLORS.primary}
        labelStyle={scaledStyles.button}
        onPress={() =>
          router.push({ pathname: "Event", params: { collectionId: "", eventId: event.id } })
        }
      >
        Start swiping
      </Button>
    </View>
  );
}

const styles = StyleSheet.create({
  absoluteImage: {
    position: "absolute",
    zIndex: 1,
    top: 30,
    left: 100,
    width: 200,
    borderRadius: 8,
    shadowColor: "#000000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    marginBottom: 20,
    backgroundColor: "white",
    padding: 12
  },
  textContainer: {
    backgroundColor: "white",
    borderRadius: 8
  },
  image: {
    height: 150,
    borderRadius: 8
  },
  backgroundImage: {
    width: 400,
    height: 350
  },
  titleText: {
    paddingHorizontal: 10,
    paddingTop: 10,
    fontSize: 15,
    fontWeight: "500",
    fontFamily: "DMSansMedium",
    textAlign: "center"
  },
  text: {
    paddingTop: 5,
    paddingHorizontal: 10,
    fontSize: 15,
    fontFamily: "DMSansRegular",
    textAlign: "center"
  },
  header: {
    fontSize: 17,
    fontFamily: "DMSansMedium",
    fontWeight: "500",
    textAlign: "center",
    marginLeft: 24,
    marginVertical: 10
  },
  button: {
    fontFamily: "DMSansMedium",
    fontSize: 15,
    fontWeight: "700"
  }
});

const scaledStyles = scaleStyleSheet(styles);
