/* eslint-disable */
import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";

type UserDetailsProps = {
  onPress: () => void;
<<<<<<< HEAD
  profileImageSource: any; // Adjust the type as needed
=======
  profileImageSource: {
    uri: string;
  } | null;
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
  name: string;
  editDetailsText: string;
  chevronColor: string;
};

export default function UserDetails(props: UserDetailsProps) {
  const { onPress, profileImageSource, name, editDetailsText, chevronColor } = props;
<<<<<<< HEAD
  //console.log(profileImageSource); // Make this the source of the Image component when images are available
=======
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb

  return (
    <TouchableOpacity onPress={onPress} style={styles.container}>
      <View style={styles.profileNameContainer}>
        {/* // eslint-disable-next-line global-require */}
<<<<<<< HEAD
        <Image source={require("../../assets/stockProPic.png")} style={styles.imageContainer} />
=======
        {profileImageSource && <Image source={profileImageSource} style={styles.imageContainer} />}
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
        <View style={styles.nameAndEditContainer}>
          <Text style={styles.name}>{name}</Text>
          <Text style={styles.edit}>{editDetailsText}</Text>
        </View>
      </View>
      <View style={[styles.chevronContainer, { backgroundColor: chevronColor }]}>
        {/* <Text style={styles.chevron}>&gt;</Text> */}
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    display: "flex",
    width: 349,
    paddingVertical: 16,
    justifyContent: "space-between",
    alignContent: "center",
    borderRadius: 8,
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
    fontSize: 24,
    fontFamily: "DMSansRegular",
    color: "#222",
    fontStyle: "normal",
    fontWeight: "700",
    lineHeight: 32
  },
  edit: {
    fontSize: 12,
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
