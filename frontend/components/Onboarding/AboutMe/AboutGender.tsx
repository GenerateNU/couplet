import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import BackButton from "./components/BackButton";
import ContinueButton from "./components/ContinueButton";
import ProgressBar from "./components/ProgressBar";

function AboutGender() {
  return (
    <View style={styles.container}>
      <View style={styles.mainContainer}>
        <View style={styles.TopUiContainer}>
          <BackButton />
          <ProgressBar />
        </View>
        <View>
          <Image source={require("../../../assets/lightningBolt.png")} />
          <View>
            <Text style={styles.headerContainer}>I am a...</Text>
          </View>
          <View>
            <View style={styles.buttonContainer}>
              <TouchableOpacity style={styles.button}>
                <Text style={styles.buttonText}>Man</Text>
              </TouchableOpacity>
              <TouchableOpacity style={styles.button}>
                <Text style={styles.buttonText}>Woman</Text>
              </TouchableOpacity>
              <TouchableOpacity style={styles.button}>
                <Text style={styles.buttonText}>Other</Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
        <View style={styles.ContinueButtonContainer}>
          <ContinueButton />
        </View>
      </View>
    </View>
  );
}

export default AboutGender;

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
  },
  button: {
    backgroundColor: "#FFF",
    borderWidth: 1,
    borderColor: "#D1AAF6",
    borderRadius: 25,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 16
  },
  buttonText: {
    color: "black",
    fontSize: 17,
    fontWeight: "500",
    letterSpacing: -0.17
  },
  buttonContainer: {}
});
