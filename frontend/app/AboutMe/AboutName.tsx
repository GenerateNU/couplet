import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  StyleSheet,
  Text,
  TextInput,
  View
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const aboutNamePicture = require("../../assets/aboutName.png");

function AboutName() {
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
  const onSubmit = (data: { name: string }) => {
    router.push({ pathname: "/AboutMe/AboutBirthday", params: data });
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
        behavior={Platform.OS === "ios" ? "padding" : "height"}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={aboutNamePicture} />
            <OnboardingTitle text="My first name is..." />
            <View style={scaledStyles.inputWrapper}>
              <Controller
                control={control}
                render={({ field: { onChange, onBlur, value } }) => (
                  <TextInput
                    style={scaledStyles.textContainer}
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
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={false}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default AboutName;

const styles = StyleSheet.create({
  TopUiContainer: {
    alignItems: "center",
    flex: 0.35
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
  textContainer: {
    padding: 8
  },
  inputWrapper: {
    borderRadius: 6,
    borderWidth: 1,
    borderColor: "grey",
    marginBottom: 8
  },
  avoidContainer: {
    flex: 1
  }
});

const scaledStyles = scaleStyleSheet(styles);
