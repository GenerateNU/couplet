import { router } from "expo-router";
import React from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";

const aboutLookingPicture = require("../../assets/lookingfor.png");

function AboutLooking() {
  const { control, handleSubmit } = useForm({
    defaultValues: {
      name: ""
    }
  });
  const name = useWatch({
    control,
    name: "name",
    defaultValue: ""
  });
  const onSubmit = (data: Object) => {
    console.log(name);
    router.push("/AboutMe/AboutPronouns");
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
          <Image source={aboutLookingPicture} />
          <OnboardingTitle text="I&apos;m looking for..."/>
          <View style={scaledStyles.inputWrapper} />
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

export default AboutLooking;

const styles = StyleSheet.create({
  TopUiContainer: {
    alignItems: "center",
    flex: 0.2
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
    marginBottom: 14
  }
});

const scaledStyles = scaleStyleSheet(styles);
