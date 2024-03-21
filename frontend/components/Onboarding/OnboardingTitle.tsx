import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { COLORS } from "../../colors";

type Props = {
  text: string;
};

const OnboardingTitle: React.FC<Props> = ({ text }) => {
  return (
    <View style={styles.centeringContainer}>
      <View style={styles.container}>
        <Text style={styles.text}>{text}</Text>
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  centeringContainer: {
    alignItems: "center", 
    width: "100%" 
  },
  container: {
    width: 346,
    height: 61,
    justifyContent: "center", 
    alignItems: "flex-start", 
    opacity: 1 
  },
  text: {
    fontFamily: "DMSansBold",
    fontSize: 32,
    fontWeight: "bold",
    lineHeight: 32,
    textAlign: "left",
    color: COLORS.black
  }
});

export default OnboardingTitle;
