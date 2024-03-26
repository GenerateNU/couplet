import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import COLORS from "../../colors";

export default function ProfileConfirm() {
  return (
    <View style={{ flex: 1, margin: 0, top: "40%" }}>
      {/* <LinearGradient
          colors={['red', 'yellow', 'green' ]}
          style={styles.linearGradient}
        > */}
      <View style={{ alignItems: "center", width: "100%" }}>
        <Text style={styles.welcomeHeader}>Welcome to Couplet</Text>
        <Text style={styles.welcomeText}>Like 5 Events you want to go to</Text>
        <TouchableOpacity onPress={() => {}} disabled={false} style={styles.button}>
          <Text style={styles.buttonText}>Let`&apos;`s Get Started</Text>
        </TouchableOpacity>
      </View>
      {/* </LinearGradient> */}
    </View>
  );
}

const styles = StyleSheet.create({
  welcomeHeader: {
    fontFamily: "DMSansBold",
    textAlign: "center",
    fontSize: 24,
    marginBottom: 10
  },
  welcomeText: {
    fontFamily: "DMSansRegular",
    textAlign: "center",
    fontSize: 16,
    marginBottom: 10
  },
  button: {
    width: 330,
    height: 41,
    paddingVertical: 10,
    paddingHorizontal: 130,
    borderRadius: 65,
    borderWidth: 1,
    borderColor: COLORS.white,
    backgroundColor: COLORS.white,
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
  buttonText: {
    fontFamily: "DMSansMedium",
    fontSize: 16,
    fontWeight: "500",
    lineHeight: 21,
    textAlign: "left",
    color: COLORS.primary,
    width: 200,
    height: 21
  },
  linearGradient: {
    alignItems: "center",
    justifyContent: "center",
    borderRadius: 5,
    height: 200,
    width: 350
  }
});
