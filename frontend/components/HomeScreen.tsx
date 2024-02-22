import React from "react";
import { Text, Image, StyleSheet, View } from "react-native";
import Header from "./Header";
import TagButton from "./TagButton";

export default function HomeScreen() {
  return (
    <View>
      <View style={{ height: 80 }}>
        <Header />
      </View>

      <View style={{ flexDirection: "row", padding: 20 }}>
        <TagButton text="All Events" />
        <TagButton text="Liked Events" />
      </View>

      <View style={styles.container}>
        <View style = {styles.imageContainer}>
        <Image source={{ uri: "../assets/blankProfile.jpg" }} style={styles.image} />
        <Image source={{ uri: "../assets/blankProfile.jpg" }} style={styles.image} />
        <Image source={{ uri: "../assets/blankProfile.jpg" }} style={styles.image} />
        <Image source={{ uri: "../assets/blankProfile.jpg" }} style={styles.image} />
        </View>
        <Text>Need Someone to Go With?</Text>
        <View style= {styles.buttonContainer}>
          <TagButton text="Match"/>
        </View>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flexDirection: "column",
    alignItems: "flex-start",
    borderStyle: "solid",
    borderWidth: 1,
    flex: 1,
    backgroundColor: "grey",
    padding: 20,
    margin: 10
  },
  image: {
    height: 50,
    width: 50,
    borderRadius: 50,
    borderWidth: 1,
    paddingBottom: 30,
  },
  imageContainer: {
    flexDirection: "row",
    paddingBottom: 10,
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    paddingTop: 10
  }
});
