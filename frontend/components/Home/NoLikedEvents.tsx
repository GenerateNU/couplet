import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import { Button } from "react-native-paper";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

const FEATURED = require("../../assets/noLikesYet.png");

export type NoLikedEventsProps = {
  onClick: (newFilter: string) => void;
};

export default function NoLikedEvents({ onClick }: NoLikedEventsProps) {
  return (
    <View style={scaledStyles.container}>
      <Image source={FEATURED} style={styles.noMatchesImage} />
      <View style={scaledStyles.bodyContainer}>
        <View style={scaledStyles.textContainer}>
          <Text style={scaledStyles.titleText}>{`You haven't liked any events!`}</Text>
          <Text style={scaledStyles.bodyText}>
            Liking an event helps you share your interests with potential matches.
          </Text>
        </View>
        <Button
          mode="contained"
          buttonColor={COLORS.primary}
          style={{ marginTop: 20, paddingHorizontal: 60, marginRight: 24 }}
          onPress={() => onClick("All Events")}
        >
          Start exploring events
        </Button>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: "column",
    alignItems: "center",
    justifyContent: "space-between",
    padding: 40
  },
  noMatchesImage: {
    width: 300,
    height: 300
  },
  bodyContainer: {
    marginTop: 40
  },
  textContainer: {
    width: 350,
    paddingRight: 24
  },
  titleText: {
    fontFamily: "DMSansMedium",
    fontWeight: "700",
    fontSize: 20,
    textAlign: "center"
  },
  bodyText: {
    fontFamily: "DMSansRegular",
    fontWeight: "400",
    fontSize: 17,
    textAlign: "center"
  }
});

const scaledStyles = scaleStyleSheet(styles);
