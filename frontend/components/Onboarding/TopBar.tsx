import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import BackButton from "./BackButton";
import OnboardingBar from "./OnboardingBar";

type TopBarProps = {
  onBackPress: () => void;
  text: string;
  selectedCount: number;
};

function TopBar({ onBackPress, text, selectedCount }: TopBarProps) {
  return (
    <View style={styles.container}>
      <BackButton onPress={onBackPress} />
      <View style={styles.textBarContainer}>
        <Text style={styles.informationText}>{text}</Text>
        <OnboardingBar selectedCount={selectedCount} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: 346,
    height: 69,
    flexDirection: "column",
    justifyContent: "space-between"
  },
  textBarContainer: {
    paddingTop: 8,
    width: 346,
    height: 21,
    justifyContent: "flex-end"
  },
  informationText: {
    width: 46,
    height: 10,
    fontFamily: "DMSansRegular",
    fontSize: 10,
    fontWeight: "400",
    lineHeight: 10,
    textAlign: "left",
    color: COLORS.darkGray,
    marginBottom: 2
  }
});

export default TopBar;
