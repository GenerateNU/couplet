import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Image, KeyboardAvoidingView, Platform, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setSchool } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const SCHOOL = require("../../../assets/school.png");

function Education() {
  const dispatch = useAppDispatch();
  // Use Form from React-Hook-Form
  const { control, handleSubmit } = useForm({
    defaultValues: {
      school: ""
    }
  });
  // Watch any changes made to the input form
  const school = useWatch({
    control,
    name: "school",
    defaultValue: ""
  });
  // On submit of the name form
  const onSubmit = (data: { school: string }) => {
    dispatch(setSchool(data.school));
    router.push("Onboarding/Education&Career/Career");
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
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={SCHOOL} />
            <OnboardingTitle text="My school is..." />
            <View style={scaledStyles.inputWrapper}>
              <View style={scaledStyles.textInputWrapper}>
                <Controller
                  control={control}
                  render={({ field: { onChange, onBlur, value } }) => (
                    <TextInput
                      style={scaledStyles.textInput}
                      placeholder="Name of School"
                      onBlur={onBlur}
                      onChangeText={onChange}
                      value={value}
                    />
                  )}
                  name="school"
                />
              </View>
              <Text style={scaledStyles.textHelper}>
                If you've graduated, write your alma mater!
              </Text>
            </View>
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!school}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default Education;

const scaledStyles = scaleStyleSheet(onboardingStyles);
