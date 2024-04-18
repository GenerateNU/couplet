import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Props = {
  text: string;
};

function OnboardingTitle({ text }: Props) {
  return (
    <View style={scaledStyles.container}>
      <Text style={scaledStyles.text}>{text}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  text: {
    fontFamily: "DMSansMedium",
    fontSize: 30,
    fontWeight: "bold",
    lineHeight: 32,
    textAlign: "left",
    color: COLORS.black
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingTitle;
