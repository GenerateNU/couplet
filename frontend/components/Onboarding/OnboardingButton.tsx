import React, { useState } from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

interface ButtonProps {
  title: string;
  onButtonClick: () => void;
  isDisabled: boolean;
}

function OnboardingButton({ title, onButtonClick, isDisabled }: ButtonProps) {
  const [isPressed, setIsPressed] = useState(false);

  const handlePress = () => {
    onButtonClick();
    setIsPressed(!isPressed);
  };
  return (
    <View>
      <TouchableOpacity
        onPress={handlePress}
        style={[scaledStyles.button, isPressed ? scaledStyles.buttonPressed : null]}
        disabled={isDisabled}
      >
        <Text style={scaledStyles.text}>{title}</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  text: {
    fontFamily: "DMSansRegular",
    fontSize: 12,
    fontWeight: "bold",
    lineHeight: 15.62,
    textAlign: "center",
    paddingHorizontal: 8
  },
  button: {
    height: 36,
    alignItems: "center",
    justifyContent: "center",
    paddingVertical: 4,
    paddingHorizontal: 12,
    borderRadius: 100,
    borderWidth: 1,
    borderColor: COLORS.secondary,
    backgroundColor: COLORS.white
  },
  buttonPressed: {
    backgroundColor: COLORS.secondary
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingButton;
