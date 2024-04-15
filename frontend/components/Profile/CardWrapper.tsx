import React from "react";
import { StyleSheet, View } from "react-native";

export default function CardWrapper({ children }: { children: React.ReactNode }) {
  return <View style={styles.container}>{children}</View>;
}

const styles = StyleSheet.create({
  container: {
    padding: 5,
    borderRadius: 20,
    width: "90%",
    alignSelf: "center"
  }
});
