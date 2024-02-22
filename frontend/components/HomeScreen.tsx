import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { ScrollView } from "react-native-gesture-handler";
import Header from "./Header";
import HomePageSection from "./HomePageSection";
import LinkButton from "./LinkButton";

const DUMMY_IMAGE = require("../assets/blankProfile.jpg");

export default function HomeScreen() {
  return (
    <ScrollView style={styles.scrollView}>
      <View>
        {/* Header View */}
        <View>
          <Header />
        </View>

        {/* Match Container View */}
        <View style={styles.matchContainer}>
          <View style={styles.imageContainer}>
            <Image source={DUMMY_IMAGE} style={styles.image} />
            <Image source={DUMMY_IMAGE} style={styles.image} />
            <Image source={DUMMY_IMAGE} style={styles.image} />
            <Image source={DUMMY_IMAGE} style={styles.image} />
          </View>
          <Text>Need Someone to Go With?</Text>
          <View style={styles.buttonContainer}>
            <LinkButton text="Match Now"/>
          </View>
        </View>

        {/* Pintrestesque Section Views */}
        <View style={styles.sectionContainer}>
          <HomePageSection title="This weekend in Boston" events={[1, 2, 3, 4, 5].map((n) => ({id: n}))} />
          <HomePageSection title="Live music and concerts" events={[1, 2, 3].map((n) => ({id: n}))} />
          <HomePageSection title="Other events" events={[1, 2, 3, 4, 5].map((n) => ({id: n}))} />
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  matchContainer: {
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
    paddingBottom: 30
  },
  imageContainer: {
    flexDirection: "row",
    paddingBottom: 10
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    paddingTop: 10
  },
  sectionContainer: {
    margin: 10
  },
  scrollableSection: {
    marginVertical: 20
  },
  scrollView: {
    marginBottom: 40
  }
});
