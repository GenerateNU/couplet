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

const pronounPicture = require("../../assets/pronouns.png");

function AboutPronouns() {
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
    router.push("/AboutMe/AboutHeight");
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
          <Image source={pronounPicture} />
          <OnboardingTitle text="My pronouns are..." />
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

export default AboutPronouns;

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
