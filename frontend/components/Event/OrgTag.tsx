import React, { StyleSheet, Text, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

export interface OrgTagProps {
  text: string;
}

export default function OrgTag({ text }: OrgTagProps) {
  return (
    <View style={scaledStyles.container}>
      <Text style={scaledStyles.text}>{text}</Text>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    backgroundColor: COLORS.secondary,
    borderRadius: 100,
    paddingVertical: 8,
    paddingHorizontal: 16
  },
  text: {
    fontFamily: "DMSansMedium",
    fontWeight: "700",
    fontSize: 15
  }
});

const scaledStyles = scaleStyleSheet(styles);
