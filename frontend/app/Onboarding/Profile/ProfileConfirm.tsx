import { LinearGradient } from "expo-linear-gradient";
import { router } from "expo-router";
import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import COLORS from "../../../colors";
import scaleStyleSheet from "../../../scaleStyles";

export default function ProfileConfirm() {
  return (
    <LinearGradient
      colors={[COLORS.primary, "orange", COLORS.primary]}
      style={styles.linearGradient}
      end={{ x: 0, y: 0 }}
      start={{ x: 1, y: 1 }}
    >
      <View style={scaledStyles.container}>
        <View style={{ flexDirection: "column", alignItems: "center" }}>
          <Text style={styles.welcomeHeader}>Welcome to Couplet</Text>
          <Text style={styles.welcomeText}>Like 5 Events you want to go to</Text>
          <TouchableOpacity
            onPress={() => {
              router.push("Profile");
            }}
            disabled={false}
            style={styles.button}
          >
            <Text style={styles.buttonText}>Let&apos;s Get Started</Text>
          </TouchableOpacity>
        </View>
      </View>
    </LinearGradient>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    flexDirection: "column",
    justifyContent: "center",
    alignItems: "center"
  },
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
    height: 21
  },
  linearGradient: {
    alignItems: "center",
    justifyContent: "center",
    height: "100%",
    width: "100%"
  }
});

const scaledStyles = scaleStyleSheet(styles);
