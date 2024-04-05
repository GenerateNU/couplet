import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setGender } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import { onButtonClick } from "../../../utils/onButtonClick";

const aboutGender = require("../../../assets/lightningBolt.png");

function AboutGender() {
  const dispatch = useAppDispatch();
  const [selectedButton, setSelectedButton] = useState("");
  const { control, handleSubmit } = useForm({
    defaultValues: {
      gender: ""
    }
  });
  const onSubmit = (data: { gender: string }) => {
    dispatch(setGender(data.gender));
    router.push("Onboarding/AboutMe/AboutInterestedIn");
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
          <View style={scaledStyles.inputWrapper}>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="gender"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title="Man"
                      onButtonClick={() => onButtonClick(value, "Man", setSelectedButton, onChange)}
                      isDisabled={Boolean(value && value !== "Man")}
                    />
                  )}
                />
              </View>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="gender"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title="Woman"
                      onButtonClick={() =>
                        onButtonClick(value, "Woman", setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== "Woman")}
                    />
                  )}
                />
              </View>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="gender"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title="Other"
                      onButtonClick={() =>
                        onButtonClick(value, "Other", setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== "Other")}
                    />
                  )}
                />
              </View>
            </View>
          </View>
        </View>
        <View>
          <ContinueButton
            title="Continue"
            isDisabled={!selectedButton}
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

const styles = onboardingStyles;

const scaledStyles = scaleStyleSheet(styles);
