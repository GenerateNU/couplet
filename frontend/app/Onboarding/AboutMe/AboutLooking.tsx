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
import { setLooking } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import { onButtonClick } from "../../../utils/onButtonClick";

const aboutLookingPicture = require("../../../assets/lookingfor.png");

function AboutLooking() {
  const prompt1 = "Long term relationship";
  const prompt2 = "Short term relationship";
  const prompt3 = "Seeing where things go";
  const prompt4 = "Friends";
  const dispatch = useAppDispatch();
  const [selectedButton, setSelectedButton] = useState("");
  const { control, handleSubmit } = useForm({
    defaultValues: {
      looking: ""
    }
  });
  const onSubmit = (data: { looking: string }) => {
    dispatch(setLooking(data.looking));
    router.push("Onboarding/AboutMe/AboutPronouns");
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
          <OnboardingTitle text="I'm looking for..." />
          <View style={scaledStyles.inputWrapper}>
            <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="looking"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title={prompt1}
                      onButtonClick={() =>
                        onButtonClick(value, prompt1, setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== prompt1)}
                    />
                  )}
                />
              </View>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="looking"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title={prompt2}
                      onButtonClick={() =>
                        onButtonClick(value, prompt2, setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== prompt2)}
                    />
                  )}
                />
              </View>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="looking"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title={prompt3}
                      onButtonClick={() =>
                        onButtonClick(value, prompt3, setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== prompt3)}
                    />
                  )}
                />
              </View>
              <View style={scaledStyles.button}>
                <Controller
                  control={control}
                  name="looking"
                  render={({ field: { onChange, value } }) => (
                    <OnboardingButton
                      title={prompt4}
                      onButtonClick={() =>
                        onButtonClick(value, prompt4, setSelectedButton, onChange)
                      }
                      isDisabled={Boolean(value && value !== prompt4)}
                    />
                  )}
                />
              </View>
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
    </SafeAreaView>
  );
}

export default AboutLooking;

const overrideStyles = {
  button: {
    marginBottom: 14
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
