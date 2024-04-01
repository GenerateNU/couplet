import { router, useNavigation } from "expo-router";
import React, { useState } from "react";
import { Image, SafeAreaView, StyleSheet, View } from "react-native";
import scaleStyleSheet from "../../../scaleStyles";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";

const RELIGION_IMAGE = require("../../../assets/OnboardingReligion.png");

export default function LifestyleReligion() {
  const navigation = useNavigation();
  const [religion, setReligion] = useState<string | null>(null);
  const listOfReligions = [
    "Christianity",
    "Islam",
    "Hindusim",
    "Buddhism",
    "Catholicism",
    "Judaism",
    "Agnosticisim",
    "Atheism",
    "Other",
    "Prefer not to say"
  ];

  const handlePillPress = (selection: string) => {
    if (religion === selection) {
      setReligion(null);
    } else {
      setReligion(selection);
    }
  };

  const isContinueButtonEnabled = () => religion !== null;

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <SafeAreaView style={scaledStyles.container}>
      <View>
        <TopBar onBackPress={() => goBack()} text="Lifestyle" selectedCount={3} />
      </View>
      <View style={scaledStyles.innerContainer}>
        <Image source={RELIGION_IMAGE} />
        <OnboardingTitle text="I believe in..." />
        <View style={scaledStyles.buttonContainer}>
          {listOfReligions.map((option) => (
            <View style={scaledStyles.onboardingButtonPadding}>
              <OnboardingButton
                key={option}
                title={option}
                onButtonClick={() => handlePillPress(option)}
              />
            </View>
          ))}
        </View>
      </View>
      <View style={scaledStyles.ContinueButtonContainer}>
        <ContinueButton
          onPress={() => router.push("Onboarding/Lifestyle/LifestylePolitics")}
          title="Continue"
          isDisabled={!isContinueButtonEnabled()}
        />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  buttonContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    alignItems: "flex-start",
    marginTop: 20
  },
  onboardingButtonPadding: {
    paddingRight: 8,
    paddingBottom: 8
  },
  container: {
    flex: 1,
    justifyContent: "space-between",
    alignItems: "center",
    margin: 30
  },
  innerContainer: {
    top: 50
  },
  ContinueButtonContainer: {
    marginTop: 10
  }
});

const scaledStyles = scaleStyleSheet(styles);
