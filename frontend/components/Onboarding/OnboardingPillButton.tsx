import React from "react";
import { StyleSheet, Text, TouchableOpacity } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

interface OnboardPillButtonProps {
  label: string;
  onPress: (label: string | null) => void;
  isSelected: boolean;
}

export default function OnboardingPillButton({
  label,
  onPress,
  isSelected
}: OnboardPillButtonProps) {
  const handlePress = () => {
    if (isSelected) {
      onPress(null);
    } else {
      onPress(label);
    }
  };

  return (
    <TouchableOpacity
      style={[scaledStyles.pillButton, isSelected && scaledStyles.selected]}
      onPress={handlePress}
    >
      <Text style={scaledStyles.buttonText}>{label}</Text>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  pillButton: {
    borderRadius: 100,
    borderWidth: 1.5,
    borderColor: COLORS.secondary,
    backgroundColor: COLORS.white,
    paddingHorizontal: 20,
    paddingVertical: 10,
    margin: 5
  },
  selected: {
    backgroundColor: COLORS.secondary
  },
  buttonText: {
    fontSize: 15
  }
});

const scaledStyles = scaleStyleSheet(styles);
