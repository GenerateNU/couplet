import React from "react";
import { StyleSheet, Text } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Props = {
  text: string;
};

function OnboardingSmallTitle({ text }: Props) {
  return <Text style={scaledStyles.smallHeadingContainer}>{text}</Text>;
}

const styles = StyleSheet.create({
  smallHeadingContainer: {
    fontFamily: "DMSansBold",
    fontSize: 20,
    lineHeight: 20,
    textAlign: "left",
    color: COLORS.black
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingSmallTitle;
