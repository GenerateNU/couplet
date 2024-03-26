/* eslint-disable */
import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity } from "react-native";

interface PurpleProfileCardProps {
  imageUrl: string;
  name: string;
  detailText: string;
  onPress: () => void;
}

export default function PurpleProfileCard(props: PurpleProfileCardProps) {
  const { imageUrl, name, detailText, onPress } = props;
  console.log(imageUrl); // Make this the source of the Image component when images are available
  return (
    <TouchableOpacity style={styles.card} onPress={onPress}>
      {/* // eslint-disable-next-line global-require */}
      <Image source={require("../assets/ProfilePurple.png")} style={styles.imageContainer} />
      <Text style={styles.myProfile}>{name}</Text>
      <Text style={styles.description}>{detailText}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  card: {
    backgroundColor: "#D1AAF6",
    borderRadius: 8,
    paddingBottom: 27,
    paddingHorizontal: 7,
    paddingTop: 14,
    display: "flex",
    flexDirection: "column",
    justifyContent: "center",
    gap: 3,
    alignItems: "center",
    maxWidth: 165,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    flex: 1
  },
  imageContainer: {
    width: 100,
    height: 100,
    borderRadius: 8
  },
  myProfile: {
    fontSize: 17,
    fontFamily: "DMSansRegular",
    color: "#222",
    fontStyle: "normal",
    fontWeight: "500",
    lineHeight: 22
  },
  description: {
    fontSize: 10,
    fontFamily: "DMSansRegular",
    color: "#222",
    fontStyle: "normal",
    fontWeight: "400",
    lineHeight: 13,
    maxWidth: 109,
    textAlign: "center"
  }
});
