import { router } from "expo-router";
import React from "react";
import { Image, SafeAreaView, StyleSheet, Text, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const aboutLookingPicture = require("../../assets/lookingfor.png");

function AboutLooking() {
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="About Me"
          selectedCount={1}
        />
      </View>
      <View style={scaledStyles.mainContainer}>
        <View>
          <Image source={aboutLookingPicture} />
          <View>
            <Text style={scaledStyles.headerContainer}>I'm looking for...</Text>
          </View>
          <View>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Long term relationship" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Short term relationship" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Seeing where things go" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Friends" onButtonClick={() => {}} />
              </View>
            </View>
          </View>
        </View>
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/AboutMe/AboutPronouns");
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutLooking;

const styles = StyleSheet.create({
  TopUiContainer: {
    flex: 0.3,
    alignItems: "center"
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
    marginBottom: 16,
    fontFamily: "DMSansMedium"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1
  },
  ContinueButtonContainer: {
    marginBottom: 10
  },
  button: {
    marginBottom: 16
  },
  buttonText: {
    color: "black",
    fontSize: 17,
    fontWeight: "500",
    letterSpacing: -0.17,
    fontFamily: "DMSansMedium"
  },
  buttonContainer: {}
});

const scaledStyles = scaleStyleSheet(styles);
