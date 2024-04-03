import { useNavigation } from "@react-navigation/native";
import React, { useState } from "react";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  StyleSheet,
  TextInput,
  View
} from "react-native";

import { router } from "expo-router";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const JOB = require("../../assets/job.png");

export default function Career() {
  const [career, setCareer] = useState("");
  const navigation = useNavigation();
  const nextRoute = "Onboarding/Lifestyle/LifestyleReligion";

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <SafeAreaView style={{ height: "100%" }}>
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        style={scaledStyles.container}
      >
        <View>
          <TopBar
            onBackPress={() => goBack()}
            text="Education and Career"
            selectedCount={2}
            skipToRoute={nextRoute}
          />
        </View>
        <View>
          <Image source={JOB} style={{ height: 200, width: 200 }} resizeMode="contain" />
          <OnboardingTitle text="My job is..." />
          <TextInput
            editable
            value={career}
            onChangeText={setCareer}
            style={scaledStyles.textInput}
            placeholder="Job title"
          />
        </View>
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            onPress={() => router.push(nextRoute)}
            title="Continue"
            isDisabled={career.length === 0}
          />
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
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
