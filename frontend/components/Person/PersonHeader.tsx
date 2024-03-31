import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { PersonHeaderProps } from "./PersonProps";

export default function PersonHeader({ firstName, age, pronouns }: PersonHeaderProps) {
  return (
    <View style={styles.container}>
      <View style={{ flexDirection: "row" }}>
        <Text style={styles.firstName}>{firstName}</Text>
        <Text style={styles.age}>{age}</Text>
      </View>
      {pronouns && <Text style={styles.pronouns}>{pronouns}</Text>}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    margin: 20
  },
  firstName: {
    fontSize: 25,
    fontFamily: "DMSansBold",
    marginRight: 10
  },
  age: {
    fontFamily: "DMSansRegular",
    fontSize: 25
  },
  pronouns: {
    fontFamily: "DMSansRegular",
    color: "gray"
  }
});
