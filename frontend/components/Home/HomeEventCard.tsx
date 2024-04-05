import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { TouchableOpacity } from "react-native-gesture-handler";
import { Icon } from "react-native-paper";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type HomeEventCardProps = {
  id: string;
  name: string;
  // TODO: we need Location and Cost, but these aren't in the endpoint response yet
};

export default function HomeEventCard({ id, name }: HomeEventCardProps) {
  return (
    <TouchableOpacity
      onPress={() => router.push({ pathname: "Event", params: { collectionId: "", eventId: id } })}
    >
      <View style={styles.card}>
        <View style={styles.imageContainer}>
          <Image
            source={{
              uri: "https://images.fineartamerica.com/images/artworkimages/mediumlarge/1/beautiful-winter-wonderland-in-the-boston-public-garden-boston-ma-wonderland-toby-mcguire.jpg"
            }}
            style={scaledStyles.image}
          />
        </View>
        <View>
          <Text style={scaledStyles.titleText}>{name}</Text>
          <View style={styles.row}>
            <Icon source="map-marker" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>Frog Pond</Text>
          </View>
          <View style={styles.row}>
            <Icon source="currency-usd" size={20} color={COLORS.darkPurple} />
            <Text style={styles.text}>$5</Text>
          </View>
        </View>
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  card: {
    width: 175,
    marginRight: 10,
    paddingBottom: 5,
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
  image: { width: 150, height: 150, borderTopLeftRadius: 10, borderTopRightRadius: 10 },
  row: {
    flexDirection: "row",
    paddingHorizontal: 10,
    borderRadius: 20
  },
  titleText: { padding: 10, fontSize: 15, fontWeight: "500", fontFamily: "DMSansMedium" },
  text: {
    marginTop: 2,
    fontFamily: "DMSansRegular",
    fontSize: 12,
    fontWeight: "400"
  }
});

const scaledStyles = scaleStyleSheet(styles);
