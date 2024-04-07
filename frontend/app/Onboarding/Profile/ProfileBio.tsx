import { router } from "expo-router";
import React, { useRef, useState } from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import {
  Image,
  Keyboard,
  KeyboardAvoidingView,
  Platform,
  StyleSheet,
  Text,
  TextInput,
  View
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import COLORS from "../../../colors";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import PromptDropDown from "../../../components/Onboarding/PromptDropDown";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setPromptBio, setResponseBio } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const BIO_IMAGE = require("../../../assets/profilebio.png");

function ProfileBio() {
  const dispatch = useAppDispatch();
  const [isPromptSelected, setIsPromptSelected] = useState(false);
  const handlePromptChange = (prompt: string) => {
    setIsPromptSelected(!!prompt);
  };
  const inputRef = useRef(null);
  // Use Form from React-Hook-Form
  const { control, handleSubmit, setValue } = useForm({
    defaultValues: {
      prompt: "",
      response: ""
    }
  });
  const responseStatus = useWatch({
    control,
    name: "response",
    defaultValue: ""
  });
  // On submit of the name form
  const onSubmit = (data: { prompt: string; response: string }) => {
    dispatch(setPromptBio(data.prompt));
    dispatch(setResponseBio(data.response));
    router.push("Onboarding/Profile/ProfilePhotos");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={4} />
      </View>

      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        style={scaledStyles.avoidContainer}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={BIO_IMAGE} />
            <OnboardingTitle text="What does your bio say?" />
            <View style={scaledStyles.inputWrapper}>
              <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                  <PromptDropDown
                    onPromptChange={(prompt: string) => {
                      onChange(prompt);
                      handlePromptChange(prompt);
                    }}
                    selectedPrompt={value}
                  />
                )}
                name="prompt"
              />
              <Controller
                control={control}
                render={({ field: { onChange, value } }) => (
                  <TextInput
                    ref={inputRef}
                    onSubmitEditing={() => Keyboard.dismiss()}
                    blurOnSubmit
                    multiline
                    maxLength={250}
                    style={scaledStyles.responseBox}
                    onChangeText={(val) => {
                      setValue("response", val);
                      onChange(val);
                    }}
                    value={value}
                    placeholder="Your response here"
                  />
                )}
                name="response"
              />
              <Text style={scaledStyles.charCount}>{responseStatus.length}/250</Text>
            </View>
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!isPromptSelected || !responseStatus}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default ProfileBio;

const styles = StyleSheet.create({
  charCount: {
    textAlign: "right",
    marginTop: 5
  },
  responseBox: {
    padding: 10,
    borderWidth: 1,
    borderColor: COLORS.darkGray,
    borderRadius: 10,
    height: "40%"
  }
});

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...styles });
