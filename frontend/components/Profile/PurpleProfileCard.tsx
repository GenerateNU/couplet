/* eslint-disable */
import { LinearGradient } from "expo-linear-gradient";
import React from "react";
import { Image, ImageSourcePropType, StyleSheet, Text, TouchableOpacity } from "react-native";

type PurpleProfileCardProps = {
  imageUrl: ImageSourcePropType;
  name: string;
  detailText: string;
  onPress: () => void;
};

export default function PurpleProfileCard(props: PurpleProfileCardProps) {
  const { imageUrl, name, detailText, onPress } = props;
  return (
    <TouchableOpacity style={styles.card} onPress={onPress}>
      <LinearGradient
        colors={["#E7D4FA", "#6B5DBE"]}
        style={styles.gradient}
        start={[0, 0]}
        end={[1, 1]}
      >
        <Image source={imageUrl} style={styles.imageContainer} />
        <Text style={styles.myProfile}>{name}</Text>
        <Text style={styles.description}>{detailText}</Text>
      </LinearGradient>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  card: {
    borderRadius: 8,
    display: "flex",
    flexDirection: "column",
    gap: 3,
    alignItems: "center",
    maxWidth: 165,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    flex: 1,
    minHeight: 220,
    maxHeight: 220
  },
  gradient: {
    borderRadius: 8,
    alignItems: "center",
    width: "100%",
    height: "100%",
    display: "flex",
    flexDirection: "column",
    gap: 3,
    paddingTop: 15
  },
  imageContainer: {
    minWidth: 100,
    minHeight: 110,
    borderRadius: 8
  },
  myProfile: {
    fontSize: 20,
    fontFamily: "DMSansRegular",
    color: "#FFF",
    fontStyle: "normal",
    fontWeight: "700",
    lineHeight: 26
  },
  description: {
    marginTop: 5,
    fontSize: 15,
    fontFamily: "DMSansRegular",
    color: "#FFF",
    fontStyle: "normal",
    fontWeight: "400",
    lineHeight: 13,
    maxWidth: 109,
    textAlign: "center"
  }
});
