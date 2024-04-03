import { router } from "expo-router";
import React from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const aboutGender = require("../../assets/lightningBolt.png");

function AboutGender() {
  const { handleSubmit } = useForm({
    defaultValues: {
      genderPreference: ""
    }
  });
  const onSubmit = (data: Object) => {
    router.push("/AboutMe/AboutInterestedIn");
  };
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
          <Image source={aboutGender} />
          <OnboardingTitle text="I am a..." />
          <View style={scaledStyles.inputWrapper} />
          <View style={scaledStyles.buttonContainer}>
            <View style={scaledStyles.button}>
              <OnboardingButton title="Man" onButtonClick={() => {}} />
            </View>
            <View style={scaledStyles.button}>
              <OnboardingButton title="Woman" onButtonClick={() => {}} />
            </View>
            <View style={scaledStyles.button}>
              <OnboardingButton title="Other" onButtonClick={() => {}} />
            </View>
          </View>
        </View>

        <View>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutGender;

const styles = StyleSheet.create({
  TopUiContainer: {
    alignItems: "center",
    flex: 0.3
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1,
    marginTop: 34,
    marginBottom: 36
  },
  helperContainer: {
    marginTop: 16
  },
  button: {
    marginBottom: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
