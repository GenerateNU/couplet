import React from "react";
import { StyleSheet, Text, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";

export default function Header() {
  return (
    <View style={scaledStyles.header}>
      <Text style={scaledStyles.text}>Couplet</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  header: {
    backgroundColor: "white",
    flex: 1,
    flexDirection: "row",
    justifyContent: "center",
    paddingTop: 16,
    paddingBottom: 12
  },
  text: {
    color: "black",
    justifyContent: "center",
    textAlign: "left",
    width: "100%",
    fontSize: 32,
    fontFamily: "DMSansMedium",
    fontWeight: "700"
  }
});

const scaledStyles = scaleStyleSheet(styles);
