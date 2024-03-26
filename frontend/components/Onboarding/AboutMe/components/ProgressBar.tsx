import React from "react";
import { Image, StyleSheet, View } from "react-native";

function ProgressBar() {
  return (
    <View style={styles.barContainer}>
      <Image source={require("../../../../assets/ProgressBar.png")} />
    </View>
  );
}
const styles = StyleSheet.create({
  barContainer: {
    alignItems: "center"
  }
});
export default ProgressBar;
