import React from "react";
import { StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, View } from "react-native";
import COLORS from "../../colors";

interface ContinueBottonProps extends TouchableOpacityProps {
  title: string;
  isDisabled: boolean;
}

function ContinueButton({title, isDisabled, onPress}: ContinueBottonProps) {
  return (
    <View style={styles.centeringContainer}>
      <TouchableOpacity
        onPress={onPress}
        disabled={isDisabled}
        style={[
          styles.button,
          isDisabled ? styles.buttonDisabled : styles.buttonEnabled
        ]}
      >
        <Text style={styles.text}>{title}</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  centeringContainer: {
    alignItems: "center",
    width: "100%"
  },
  button: {
    width: 330,
    height: 41,
    paddingVertical: 10,
    paddingHorizontal: 130,
    borderRadius: 65,
    borderWidth: 1,
    borderColor: COLORS.disabled,
    backgroundColor: COLORS.disabled,
    shadowColor: "#000",
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
    textAlign: "left",
    color: COLORS.white,
    width: 70,
    height: 21
  }
});

export default ContinueButton;
