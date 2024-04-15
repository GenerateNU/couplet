import { router } from "expo-router";
import React from "react";
import {
  Image,
  ImageBackground,
  SafeAreaView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View
} from "react-native";
import gradient from "../../../assets/gradient.png";
import logo from "../../../assets/logo.png";
import COLORS from "../../../colors";
import scaleStyleSheet from "../../../scaleStyles";

export default function ProfileConfirm() {
  return (
    <ImageBackground source={gradient} style={{ flex: 1 }} resizeMode="cover">
      <SafeAreaView style={scaledStyles.outerView}>
        <View>
          <View style={scaledStyles.titleImageView}>
            <Text style={scaledStyles.coupletText}>Your profile is set.</Text>
            <Text style={scaledStyles.helperText}>Now onto the fun part, lets get you a date!</Text>
            <Image style={scaledStyles.coupletLogo} source={logo} />
          </View>
          {/* Buttons */}
          <View>
            <TouchableOpacity
              style={scaledStyles.button}
              onPress={() => {
                router.push("/Home");
              }}
            >
              <Text style={scaledStyles.helperText}>Let&apos;s check out some events</Text>
            </TouchableOpacity>
          </View>
        </View>
      </SafeAreaView>
    </ImageBackground>
  );
}

const styles = StyleSheet.create({
  outerView: {
    justifyContent: "center",
    alignItems: "center",
    flex: 1
  },
  coupletLogo: {
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 0 },
    shadowRadius: 10,
    shadowOpacity: 0.5
  },
  titleImageView: {
    alignItems: "center",
    marginBottom: 20
  },
  coupletText: {
    fontFamily: "DMSansBold",
    fontSize: 32,
    color: COLORS.white,
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 0 },
    shadowRadius: 25,
    shadowOpacity: 1
  },
  helperText: {
    fontSize: 17,
    fontWeight: "400",
    lineHeight: 30,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium",
    color: COLORS.white
  },
  button: {
    width: 330,
    height: 41,
    borderRadius: 65,
    borderWidth: 1,
    borderColor: COLORS.primary,
    backgroundColor: COLORS.primary,
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
  }
});

const scaledStyles = scaleStyleSheet(styles);
