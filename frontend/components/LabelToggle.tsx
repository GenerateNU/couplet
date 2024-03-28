import React, { useState } from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import COLORS from "../colors";
import scaleStyleSheet from "../scaleStyles";

interface LabelToggleProps {
  labels: string[];
}

export default function LabelToggle({ labels }: LabelToggleProps) {
  const [chosen, setChosen] = useState<number>(0);

  return (
    <View style={scaledStyles.container}>
      {labels.map((label, i) => (
        <TouchableOpacity
          onPress={() => setChosen(i)}
          style={chosen === i ? scaledStyles.chosenLabel : scaledStyles.unchosenLabel}
        >
          {chosen === i ? (
            <View style={scaledStyles.dropShadowContainer}>
              <Text style={scaledStyles.textStyle}>{label}</Text>
            </View>
          ) : (
            <Text style={scaledStyles.textStyle}>{label}</Text>
          )}
        </TouchableOpacity>
      ))}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    alignSelf: "flex-start",
    flexBasis: "auto",
    flexDirection: "row",
    padding: 2,
    backgroundColor: COLORS.primary,
    borderRadius: 30
  },
  textStyle: {
    fontSize: 16,
    fontFamily: "DMSansMedium",
    color: COLORS.white
  },
  dropShadowContainer: {
    padding: 10,
    margin: 1,
    borderColor: "#F95D5D",
    borderWidth: 1,
    borderRadius: 24,
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.5,
    shadowRadius: 2.5,
    backgroundColor: "#F95D5D",
  },
  chosenLabel: {
    borderWidth: 1,
    borderColor: COLORS.primary,
    borderRadius: 24,
    margin: 2,
    shadowColor: "#000",
    shadowOffset: { width: 0, height: 1 },
    shadowOpacity: 0.2,
    shadowRadius: 4,
    elevation: 3,
    backgroundColor: COLORS.primary
  },
  unchosenLabel: {
    padding: 12,
    borderWidth: 1,
    borderColor: COLORS.primary,
    borderRadius: 20,
    margin: 2,
    opacity: 0.6
  }
});

const scaledStyles = scaleStyleSheet(styles);
