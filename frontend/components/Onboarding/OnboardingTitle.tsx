import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Props = {
  text: string;
};

function OnboardingTitle({ text }: Props) {
  return (
    <View style={scaledStyles.centeringContainer}>
      <View style={scaledStyles.container}>
        <Text style={scaledStyles.text}>{text}</Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  centeringContainer: {
    alignItems: "center",
    width: "100%"
  },
  container: {
    width: 346,
    height: 64,
    justifyContent: "center",
    alignItems: "flex-start",
    opacity: 1
  },
  text: {
    fontFamily: "DMSansMedium",
    fontSize: 32,
    fontWeight: "bold",
    lineHeight: 32,
    textAlign: "left",
    color: COLORS.black
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingTitle;
