import { router, useNavigation } from "expo-router";
import React, { useState } from "react";
import { Image, StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import scaleStyleSheet from "../../../scaleStyles";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";

const POLITICS_IMAGE = require("../../../assets/OnboardingPolitics.png");

export default function LifestylePolitics() {
  const navigation = useNavigation();
  const [politics, setPolitics] = useState<string | null>(null);
  const listOfPolitics = ["Liberal", "Moderate", "Conservative", "Other", "Prefer not to say"];

  const handlePillPress = (selection: string) => {
    if (politics === selection) {
      setPolitics(null);
    } else {
      setPolitics(selection);
    }
  };

  const isContinueButtonEnabled = () => politics !== null;

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
      <View>
        <Image source={POLITICS_IMAGE} />
        <OnboardingTitle text="Politically, I am..." />
        <View style={scaledStyles.buttonContainer}>
          {listOfPolitics.map((option) => (
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
      <View style={scaledStyles.continueContainer}>
        <ContinueButton
          title="Continue"
          isDisabled={!isContinueButtonEnabled()}
          onPress={() => router.push("Onboarding/Lifestyle/LifestyleHabits")}
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
  }
});

const scaledStyles = scaleStyleSheet(styles);
