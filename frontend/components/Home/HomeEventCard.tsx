import { router } from "expo-router";
import React from "react";
import { Image, Pressable, StyleSheet, Text, View } from "react-native";

const PIN = require("../../assets/pin.png");
const COIN = require("../../assets/coin.png");

type HomeEventCardProps = {
  id: string;
  name: string;
  // TODO: we need Location and Cost, but these aren't in the endpoint response yet
};

export default function HomeEventCard({ id, name }: HomeEventCardProps) {
  return (
    <Pressable
      onPress={() => router.push({ pathname: "Event", params: { collectionId: "", eventId: id } })}
    >
      <View style={styles.card}>
        <View style={styles.imageContainer}>
          {/* <Image
            source={{ uri: "https://reactnative.dev/img/tiny_logo.png" }}
            style={{ width: 150, height: 150 }}
          /> */}
        </View>
        <View>
          <Text style={{ padding: 10, fontSize: 14, fontFamily: "DMSansMedium" }}>{name}</Text>
          <View style={styles.row}>
            <Image source={PIN} style={{ width: 20, height: 20 }} />
            <Text style={styles.text}>Frog Pond</Text>
          </View>
          <View style={styles.row}>
            <Image source={COIN} style={{ width: 20, height: 20 }} />
            <Text style={[styles.text, { marginHorizontal: 2 }]}>$$</Text>
          </View>
        </View>
      </View>
    </Pressable>
  );
}

const styles = StyleSheet.create({
  card: {
    width: 175,
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 10,
    marginRight: 10
  },
  imageContainer: {
    width: "100%",
    height: 150,
    backgroundColor: "rgb(200,200,200)",
    borderTopLeftRadius: 10,
    borderTopRightRadius: 10
  },
  row: {
    flexDirection: "row",
    padding: 10,
    borderRadius: 20,
    paddingTop: 0
  },
  text: {
    marginTop: 2,
    fontFamily: "DMSansRegular"
  }
});
