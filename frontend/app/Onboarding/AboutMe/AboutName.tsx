import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Image, KeyboardAvoidingView, Platform, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setName } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import { screenHeight } from "../../../utils/dimensions";

const aboutNamePicture = require("../../../assets/aboutName.png");

function AboutName() {
  const dispatch = useAppDispatch();
  // Use Form from React-Hook-Form
  const { control, handleSubmit } = useForm({
    defaultValues: {
      name: ""
    }
  });
  // Watch any changes made to the input form
  const name = useWatch({
    control,
    name: "name",
    defaultValue: ""
  });
  // On submit of the name form
  const onSubmit = (data: { name: string }) => {
    dispatch(setName(data.name));
    router.push("Onboarding/AboutMe/AboutBirthday");
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
      <KeyboardAvoidingView
        style={scaledStyles.avoidContainer}
        keyboardVerticalOffset={screenHeight * 0.1}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={aboutNamePicture} />
            <OnboardingTitle text="My first name is..." />
            <View style={scaledStyles.inputWrapper}>
              <View style={scaledStyles.textInputWrapper}>
                <Controller
                  control={control}
                  render={({ field: { onChange, onBlur, value } }) => (
                    <TextInput
                      style={scaledStyles.textInput}
                      placeholder="First Name"
                      onBlur={onBlur}
                      onChangeText={onChange}
                      value={value}
                    />
                  )}
                  name="name"
                />
              </View>
              <Text style={scaledStyles.textHelper}>
                This is how it will permanently appear on your profile
              </Text>
            </View>
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!name}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default AboutName;

const scaledStyles = scaleStyleSheet(onboardingStyles);
