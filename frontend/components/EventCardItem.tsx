import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";

type EventCardItemProps = {
  title: string;
  description: string;
  imageUrl: string;
};

export function EventCardItem({ title, description, imageUrl }: EventCardItemProps) {
  return (
    <View style={styles.cardContainer}>
      <Image source={{ uri: imageUrl }} style={styles.cardImage} />
      <Text style={styles.cardTitle}>{title}</Text>
      <Text style={styles.cardDescription}>{description}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    alignItems: "center",
    backgroundColor: "#fff"
  },
  cardContainer: {
    borderWidth: 1,
    borderColor: "#000",
    borderRadius: 10,
    maxWidth: 150,
    alignItems: "flex-start"
  },
  cardImage: {
    width: "100%",
    aspectRatio: 0.93
  },
  cardTitle: {
    marginTop: 6,
    fontSize: 18,
    fontWeight: "bold",
    paddingLeft: 6
  },
  cardDescription: {
    fontSize: 14,
    paddingLeft: 6
  }
});

export default EventCardItem;
