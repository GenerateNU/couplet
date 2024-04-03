import { router, useNavigation } from "expo-router";
import React, { useState } from "react";
import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import COLORS from "../../../colors";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";

const PASSIONS_IMAGE = require("../../../assets/OnboardingPassions.png");

export default function LifestylePassions() {
  const navigation = useNavigation();
  const [interests, setInterests] = useState<string[]>([]);
  const listOfPassions = [
    "Acting",
    "Baking",
    "Board Games",
    "Cars",
    "Calligraphy",
    "Cooking",
    "Concerts",
    "Cycling",
    "Dancing",
    "DIY",
    "Fishing",
    "Hiking",
    "Interior Design",
    "Gaming",
    "Gardening",
    "Karaoke",
    "K-Pop",
    "Knitting",
    "Music",
    "Painting",
    "Parkour",
    "Photography",
    "Pilates",
    "Poetry",
    "Puzzles",
    "Running",
    "Rock Climbing",
    "Reading",
    "Swimming",
    "Surfing",
    "Sewing",
    "Singing",
    "Sports",
    "Traveling",
    "Trivia",
    "Video Games",
    "Volunteering",
    "Writing",
    "Weight Lifting",
    "Yoga"
  ];

  const handleInterestsButton = (selection: string) => {
    if (interests.includes(selection)) {
      const updatedInterests = interests.filter((interest) => interest !== selection);
      setInterests(updatedInterests);
    } else if (interests.length < 5) {
      setInterests([...interests, selection]);
    }
  };

  const isContinueButtonEnabled = () => interests.length === 5;

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
          <Image source={PASSIONS_IMAGE} />
          <OnboardingTitle text="What are you passionate about?" />
          <Text style={scaledStyles.subtitle}>Select your top five interests!</Text>
          <View style={scaledStyles.buttonContainer}>
            {listOfPassions.map((option) => (
              <View style={scaledStyles.onboardingButtonPadding}>
                <OnboardingButton
                  key={option}
                  title={option}
                  onButtonClick={() => handleInterestsButton(option)}
                  isDisabled={false}
                />
              </View>
            ))}
          </View>
        </View>
      </ScrollView>
      <View style={scaledStyles.ContinueButtonContainer}>
        <ContinueButton
          title={
            interests.length === 0 || interests.length === 5
              ? "Continue"
              : `Continue ${interests.length}/5`
          }
          isDisabled={!isContinueButtonEnabled()}
          onPress={() => router.push("Onboarding/Lifestyle/ProfileBio")}
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
    paddingRight: 8,
    paddingBottom: 8
  },
  subtitle: {
    fontFamily: "DMSansRegular",
    color: COLORS.darkGray
  },
  buttonContainer: {
    flexDirection: "row",
    flexWrap: "wrap",
    alignItems: "flex-start",
    marginTop: 20
  },
  ContinueButtonContainer: {
    paddingTop: 30,
    marginBottom: -20,
    backgroundColor: COLORS.white
  }
});

const scaledStyles = scaleStyleSheet(styles);
