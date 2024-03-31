import { router, useNavigation } from "expo-router";
import React, { useState } from "react";
import { Image, ScrollView, StyleSheet, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";
import ContinueButton from "../Onboarding/ContinueButton";
import OnboardingButton from "../Onboarding/OnboardingButton";
import OnboardingSmallTitle from "../Onboarding/OnboardingSmallTitle";
import OnboardingTitle from "../Onboarding/OnboardingTitle";
import TopBar from "../Onboarding/TopBar";

const HABITS_IMAGE = require("../../assets/OnboardingHabits.png");

export default function LifestyleHabits() {
  const navigation = useNavigation();
  const [drink, setDrink] = useState<string | null>(null);
  const [smoke, setSmoke] = useState<string | null>(null);
  const [weed, setWeed] = useState<string | null>(null);
  const [drugs, setDrugs] = useState<string | null>(null);
  const listOfChoices = ["Yes", "No", "Sometimes", "Socially", "Prefer not to say"];
  const drugList = ["Yes", "No", "Prefer not to say"];

  const handleDrinkChoice = (selection: string) => {
    if (drink === selection) {
      setDrink(null);
    } else {
      setDrink(selection);
    }
  };

  const handleSmokeChoice = (selection: string) => {
    if (smoke === selection) {
      setSmoke(null);
    } else {
      setSmoke(selection);
    }
  };

  const handleWeedChoice = (selection: string) => {
    if (weed === selection) {
      setWeed(null);
    } else {
      setWeed(selection);
    }
  };

  const handleDrugsChoice = (selection: string) => {
    if (drugs === selection) {
      setDrugs(null);
    } else {
      setDrugs(selection);
    }
  };

  const isContinueButtonEnabled = () =>
    drink !== null && smoke !== null && weed !== null && drugs !== null;

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <SafeAreaView style={scaledStyles.container}>
      <ScrollView>
        <View>
          <TopBar onBackPress={() => goBack()} text="Lifestyle" selectedCount={3} />
        </View>
        <View>
          <Image source={HABITS_IMAGE} />
          <OnboardingTitle text="What are your habits?" />
          <OnboardingSmallTitle text="Do you drink?" />
          <View style={scaledStyles.buttonContainer}>
            {listOfChoices.map((option) => (
              <View style={scaledStyles.onboardingButtonPadding}>
                <OnboardingButton
                  key={option}
                  title={option}
                  onButtonClick={() => handleDrinkChoice(option)}
                />
              </View>
            ))}
          </View>
          <View style={scaledStyles.separator} />
          <OnboardingSmallTitle text="Do you smoke?" />
          <View style={scaledStyles.buttonContainer}>
            {listOfChoices.map((option) => (
              <View style={scaledStyles.onboardingButtonPadding}>
                <OnboardingButton
                  key={option}
                  title={option}
                  onButtonClick={() => handleSmokeChoice(option)}
                />
              </View>
            ))}
          </View>
          <View style={scaledStyles.separator} />
          <OnboardingSmallTitle text="Do you smoke weed?" />
          <View style={scaledStyles.buttonContainer}>
            {listOfChoices.map((option) => (
              <View style={scaledStyles.onboardingButtonPadding}>
                <OnboardingButton
                  key={option}
                  title={option}
                  onButtonClick={() => handleWeedChoice(option)}
                />
              </View>
            ))}
          </View>
          <View style={scaledStyles.separator} />
          <OnboardingSmallTitle text="Do you do drugs?" />
          <View style={scaledStyles.buttonContainer}>
            {drugList.map((option) => (
              <View style={scaledStyles.onboardingButtonPadding}>
                <OnboardingButton
                  key={option}
                  title={option}
                  onButtonClick={() => handleDrugsChoice(option)}
                />
              </View>
            ))}
          </View>
        </View>
      </ScrollView>
      <View style={scaledStyles.ContinueButtonContainer}>
        <ContinueButton
          onPress={() => router.push("Onboarding/LifestylePassions")}
          title="Continue"
          isDisabled={!isContinueButtonEnabled()}
        />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flexgrow: 1,
    justifyContent: "space-between",
    alignItems: "center",
    margin: 30,
    flex: 1
  },
  onboardingButtonPadding: {
    paddingRight: 14,
    paddingBottom: 14
  },
  buttonContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    alignItems: "flex-start",
    marginTop: 20
  },
  ContinueButtonContainer: {
    paddingTop: 30,
    backgroundColor: COLORS.white,
    marginBottom: -20
  },
  separator: {
    borderBottomColor: COLORS.lightGray,
    borderBottomWidth: 1,
    marginBottom: 30,
    marginTop: 7
  }
});

const scaledStyles = scaleStyleSheet(styles);
