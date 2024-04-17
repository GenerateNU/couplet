import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Image, KeyboardAvoidingView, Platform, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setJob } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import { screenHeight } from "../../../utils/dimensions";

const JOB = require("../../../assets/job.png");

function Career() {
  const dispatch = useAppDispatch();
  // Use Form from React-Hook-Form
  const { control, handleSubmit } = useForm({
    defaultValues: {
      job: ""
    }
  });
  // Watch any changes made to the input form
  const job = useWatch({
    control,
    name: "job",
    defaultValue: ""
  });
  // On submit of the name form
  const onSubmit = (data: { job: string }) => {
    dispatch(setJob(data.job));
    router.push("Onboarding/Lifestyle/LifestyleReligion");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="Education and Career"
          selectedCount={2}
        />
      </View>
      <KeyboardAvoidingView
        style={scaledStyles.avoidContainer}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        keyboardVerticalOffset={screenHeight * 0.1}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={JOB} />
            <OnboardingTitle text="My job is..." />
            <View style={scaledStyles.inputWrapper}>
              <View style={scaledStyles.textInputWrapper}>
                <Controller
                  control={control}
                  render={({ field: { onChange, onBlur, value } }) => (
                    <TextInput
                      style={scaledStyles.textInput}
                      placeholder="Job title"
                      onBlur={onBlur}
                      onChangeText={onChange}
                      value={value}
                    />
                  )}
                  name="job"
                />
              </View>
            </View>
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!job}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default Career;

const scaledStyles = scaleStyleSheet(onboardingStyles);
