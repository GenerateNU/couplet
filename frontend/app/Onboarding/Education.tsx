import { useNavigation } from "@react-navigation/native";
import React, { useState } from "react";
import { Image, KeyboardAvoidingView, Platform, StyleSheet, TextInput, View } from "react-native";

import { router } from "expo-router";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const SCHOOL = require("../../assets/school.png");

export default function Education() {
  const [school, setSchool] = useState("");
  const navigation = useNavigation();

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={scaledStyles.container}
    >
      <View>
        <TopBar onBackPress={() => goBack()} text="Education and Career" selectedCount={2} />
      </View>
      <View>
        <Image source={SCHOOL} style={{ height: 200, width: 200 }} resizeMode="contain" />
        <OnboardingTitle text="My school is..." />
        <TextInput
          editable
          value={school}
          onChangeText={setSchool}
          style={scaledStyles.textInput}
          placeholder="Name of School"
        />
      </View>
      <View style={scaledStyles.ContinueButtonContainer}>
        <ContinueButton
          onPress={() => router.push("Onboarding/Career")}
          title="Continue"
          isDisabled={false}
        />
      </View>
    </KeyboardAvoidingView>
  );
}

const styles = StyleSheet.create({
  textInput: {
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "#9EA3A2",
    color: "#000000",
    borderRadius: 10,
    padding: 10,
    fontFamily: "DMSansRegular"
  },
  container: {
    flex: 1,
    justifyContent: "space-between",
    alignItems: "center",
    margin: 30
  },
  ContinueButtonContainer: {
    marginBottom: 10
  }
});

const scaledStyles = scaleStyleSheet(styles);
