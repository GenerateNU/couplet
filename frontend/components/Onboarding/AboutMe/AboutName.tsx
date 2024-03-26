import React from "react";
import { Image, StyleSheet, Text, TextInput, View } from "react-native";
import AboutBirthday from "./AboutBirthday";
import BackButton from "./components/BackButton";
import ContinueButton from "./components/ContinueButton";
import ProgressBar from "./components/ProgressBar";

function AboutName() {
  return <View style={styles.container}>
    <View style={styles.mainContainer}>
      <View style={styles.TopUiContainer}>
        <BackButton />
        <ProgressBar />
      </View>
      <View>
        <Image source={require("../../../assets/3dicons.png")} />
        <View>
          <Text style={styles.headerContainer}>My first name is...</Text>
        </View>
        <View style={styles.inputWrapper}>
          <TextInput style={styles.inputContainer} placeholder="First Name" />
        </View>
        <Text style={styles.textHelper}>
          This is how it will permanently appear on your profile
        </Text>
      </View>
      <View style={styles.ContinueButtonContainer}>
        <ContinueButton
          route="/components/Onboarding/AboutMe/AboutBirthday"
          Component={AboutBirthday}
        />
      </View>
    </View>
  </View>
}

export default AboutName;

const styles = StyleSheet.create({
  TopUiContainer: {
    flex: 0.15,
    marginTop: 10
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  headerContainer: {
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    letterSpacing: -0.32,
    marginTop: 16,
    marginBottom: 16
  },
  inputContainer: {
    padding: 8
  },
  inputWrapper: {
    borderRadius: 6,
    borderWidth: 1,
    borderColor: "grey",
    marginBottom: 8
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12
  },
  container: {
    flex: 1
  },
  ContinueButtonContainer: {
    marginBottom: 10
  }
});
