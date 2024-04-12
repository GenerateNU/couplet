import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { PillProps } from "./PersonProps";
import scaleStyleSheet from "../../scaleStyles";

export default function InfoChips({ items, textColor, backgroundColor }: PillProps) {
  return (
    <View style={scaledStyles.container}>
      {items.map((item) => (
        <View style={{ ...scaledStyles.chipItemContainer, backgroundColor }}>
          <Text style={{ ...scaledStyles.chipItemText, color: textColor }}>{item}</Text>
        </View>
      ))}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    display: "flex",
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "flex-start"
  },
  chipItemContainer: {
    border: "black",
    borderRadius: 20,
    width: "auto",
    padding: 10,
    paddingLeft: 17,
    paddingRight: 17,
    margin: 5
  },
  chipItemText: {
    fontFamily: "DMSansBold",
    fontSize: 13
  }
});

const scaledStyles = scaleStyleSheet(styles);