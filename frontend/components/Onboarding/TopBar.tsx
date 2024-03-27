import React from "react";
import { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import BackButton from "./BackButton";

type TopBarProps = {
  onBackPress: () => void;
  text: string;
  selectedCount: number;
};

function TopBar({ onBackPress, text, selectedCount }: TopBarProps) {
  return (
    <View style={scaledStyles.container}>
      <BackButton onPress={onBackPress} />
      <View style={scaledStyles.textBarContainer}>
        <Text style={scaledStyles.informationText}>{text}</Text>
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    width: 346,
    height: 69,
    top: 68,
    left: 24,
    position: "absolute",
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
