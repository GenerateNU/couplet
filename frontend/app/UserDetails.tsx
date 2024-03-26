import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";

interface UserDetailsProps {
  onPress: () => void;
  profileImageSource: any; // Adjust the type as needed
  name: string;
  editDetailsText: string;
  chevronColor: string;
}

export default function UserDetails(props: UserDetailsProps) {
  const { onPress, profileImageSource, name, editDetailsText, chevronColor } = props;
  console.log(profileImageSource); // Make this the source of the Image component when images are available

  return (
    <TouchableOpacity onPress={onPress} style={styles.container}>
      <View style={styles.profileNameContainer}>
        <Image source={require("../assets/stockProPic.png")} style={styles.imageContainer} />
        <View style={styles.nameAndEditContainer}>
          <Text style={styles.name}>{name}</Text>
          <Text style={styles.edit}>{editDetailsText}</Text>
        </View>
      </View>
      <View style={[styles.chevronContainer, { backgroundColor: chevronColor }]}>
        <Text style={styles.chevron}>&gt;</Text>
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    display: "flex",
    width: 349,
    padding: 16,
    justifyContent: "space-between",
    alignContent: "center",
    elevation: 5,
    borderRadius: 8,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 2 },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    backgroundColor: "#FFF",
    flexDirection: "row"
  },
  profileNameContainer: {
    width: 164,
    height: 54,
    flexShrink: 0,
    flexDirection: "row",
    alignItems: "center",
    gap: 16
  },
  imageContainer: {
    width: 54,
    height: 54,
    borderRadius: 27,
    backgroundColor: "#0F0"
  },
  nameAndEditContainer: {
    display: "flex",
    flexDirection: "column",
    alignItems: "flex-start"
  },
  name: {
    fontSize: 16,
    fontFamily: "DMSansRegular",
    color: "#222",
    fontStyle: "normal",
    fontWeight: "700",
    lineHeight: 21
  },
  edit: {
    fontSize: 10,
    fontFamily: "DMSansRegular",
    color: "#222",
    fontStyle: "normal",
    fontWeight: "400",
    lineHeight: 13
  },
  chevronContainer: {
    height: "auto",
    backgroundColor: "#0FF",
    display: "flex",
    flexDirection: "row",
    justifyContent: "center",
    alignItems: "center"
  },
  chevron: {
    fontSize: 16,
    color: "#222",
    lineHeight: 21
  }
});
