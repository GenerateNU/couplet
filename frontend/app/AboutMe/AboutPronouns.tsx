import { router } from "expo-router";
import React from "react";
import { Image, SafeAreaView, StyleSheet, Text, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const pronounPicture = require("../../assets/pronouns.png");

function AboutPronouns() {
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
          <Image source={pronounPicture} />
          <View>
            <Text style={scaledStyles.headerContainer}>My pronouns are...</Text>
          </View>
          <View>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <OnboardingButton title="He/Him" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="She/Her" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="They/them" onButtonClick={() => {}} />
              </View>
            </View>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <OnboardingButton title="He/They" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="She/They" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Xe/Xem" onButtonClick={() => {}} />
              </View>
            </View>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Ze/Zir" onButtonClick={() => {}} />
              </View>
            </View>
          </View>
        </View>
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/AboutMe/AboutHeight");
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutPronouns;

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
    marginRight: 8,
    marginBottom: 8
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start"
  }
});

const scaledStyles = scaleStyleSheet(styles);
