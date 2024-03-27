import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

type Props = {
  text: string;
};

function OnboardingSmallTitle({text}: Props) {
    return <Text style={styles.smallHeadingContainer}>{text}</Text>;
};

const styles = StyleSheet.create({
  smallHeadingContainer: {
    fontFamily: 'DMSansBold',
    fontSize: 20,
    fontWeight: 'bold',
    lineHeight: 20,
    textAlign: 'left',
    color: COLORS.black,
    paddingLeft: 10,
  },
});
  
  const scaledStyles = scaleStyleSheet(styles);
  
  export default OnboardingSmallTitle;