import React from "react";
import { StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, View } from "react-native";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

interface ContinueBottonProps extends TouchableOpacityProps {
  title: string;
  isDisabled: boolean;
}

function ContinueButton({ title, isDisabled, onPress }: ContinueBottonProps) {
  return (
    <View style={scaledStyles.centeringContainer}>
      <TouchableOpacity
        onPress={onPress}
        disabled={isDisabled}
        style={[
          scaledStyles.button,
          isDisabled ? scaledStyles.buttonDisabled : scaledStyles.buttonEnabled
        ]}
      >
        <Text style={scaledStyles.text}>{title}</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  centeringContainer: {
    alignItems: "center",
    width: "100%"
  },
  button: {
    width: 330,
    height: 41,
    borderRadius: 65,
    borderWidth: 1,
    borderColor: COLORS.disabled,
    backgroundColor: COLORS.disabled,
    shadowColor: COLORS.black,
    shadowOffset: {
      width: 0,
      height: 2
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
    justifyContent: "center",
    alignItems: "center"
  },
  buttonEnabled: {
    backgroundColor: COLORS.primary,
    borderColor: COLORS.primary
  },
  buttonDisabled: {},
  text: {
    fontFamily: "DMSansMedium",
    fontSize: 16,
    fontWeight: "500",
    lineHeight: 21,
    color: COLORS.white,
    height: 21
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default ContinueButton;
