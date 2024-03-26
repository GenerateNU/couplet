import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";

function ContinueButton() {
  return (
    <View>
      <TouchableOpacity style={styles.button}>
        <Text style={styles.text}>Continue</Text>
      </TouchableOpacity>
    </View>
  );
}

const styles = StyleSheet.create({
  button: {
    backgroundColor: "#FBA4A4",
    borderRadius: 25,
    height: 50,
    justifyContent: "center",
    alignItems: "center"
  },
  text: {
    color: "white", // choose your color
    fontSize: 16 // choose your size
  }
});

export default ContinueButton;
