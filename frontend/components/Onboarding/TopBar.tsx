import { router } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import BackButton from "./BackButton";
import OnboardingBar from "./OnboardingBar";

type TopBarProps = {
  onBackPress: () => void;
  text: string;
  selectedCount: number;
  skipToRoute?: string; // optional
};

function TopBar({ onBackPress, text, selectedCount, skipToRoute }: TopBarProps) {
  return (
    <View style={scaledStyles.container}>
      <View style={scaledStyles.topContainer}>
        <BackButton onPress={onBackPress} />
        {skipToRoute && (
          <Text style={scaledStyles.skipText} onPress={() => router.push(skipToRoute)}>
            Skip
          </Text>
        )}
      </View>
      <View style={scaledStyles.textBarContainer}>
        <Text style={scaledStyles.informationText}>{text}</Text>
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
  topContainer: {
    flexDirection: "row",
    justifyContent: "space-between"
  },
  skipText: {
    fontFamily: "DMSansRegular",
    fontSize: 17,
    fontWeight: "500",
    color: COLORS.primary,
    marginTop: 10
  },
  textBarContainer: {
    paddingTop: 8,
    width: 346,
    height: 21,
    justifyContent: "flex-end"
  },
  informationText: {
    height: 18,
    fontFamily: "DMSansMedium",
    fontSize: 14,
    fontWeight: "500",
    lineHeight: 18.23,
    textAlign: "center",
    color: COLORS.darkGray,
    marginBottom: 2
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default TopBar;

TopBar.defaultProps = {
  skipToRoute: ""
};
