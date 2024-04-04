import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, ScrollView, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setPolitics } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import chunkArray from "../../../utils/chunkArray";
import { onButtonClickArray } from "../../../utils/onButtonClick";

const PASSIONS_IMAGE = require("../../../assets/OnboardingPassions.png");

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

function LifestylePassions() {
  const dispatch = useAppDispatch();
  const [selectedButtons, setSelectedButtons] = useState([]);

  const listOfPassionsChunked = chunkArray(listOfPassions, 3);

  const { control, handleSubmit } = useForm<{ passions: string[] }>({
    defaultValues: {
      passions: []
    }
  });
  const onSubmit = (data: { passions: string[] }) => {
    dispatch(setPolitics(data.passions));
    router.push("Onboarding/Profile/ProfileBio");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="Lifestyle"
          selectedCount={3}
        />
      </View>
      <View style={scaledStyles.mainContainer}>
        <View>
          <Image source={PASSIONS_IMAGE} />
          <OnboardingTitle text="What are you passionate about?" />
          <View style={scaledStyles.helperTextContainer}>
            <Text style={scaledStyles.textHelper}>Select your top five interests!</Text>
          </View>
          <View style={scaledStyles.inputWrapper}>
            <ScrollView style={scaledStyles.habitWindow}>
              {listOfPassionsChunked.map((Row, rowIndex) => (
                // eslint-disable-next-line react/no-array-index-key
                <View key={rowIndex} style={scaledStyles.buttonContainer}>
                  {Row.map((title: string, index: React.Key | null | undefined) => (
                    // eslint-disable-next-line react/no-array-index-key
                    <View key={index} style={scaledStyles.button}>
                      <Controller
                        control={control}
                        name="passions"
                        render={({ field: { onChange, value } }) => (
                          <OnboardingButton
                            title={title}
                            onButtonClick={() =>
                              onButtonClickArray(value, title, setSelectedButtons, onChange)
                            }
                            isDisabled={value.length >= 5 && !value.includes(title)}
                          />
                        )}
                      />
                    </View>
                  ))}
                </View>
              ))}
            </ScrollView>
          </View>
        </View>
        <View>
          <ContinueButton
            title="Continue"
            isDisabled={!(selectedButtons.length === 5)}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default LifestylePassions;

const overrideStyles = {
  button: {
    marginRight: 8,
    marginBottom: 8
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start"
  },
  TopUiContainer: {
    alignItems: "center"
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    marginTop: 24,
    justifyContent: "space-between"
  },
  habitWindow: {
    height: "50%"
  },
  helperTextContainer: {
    marginTop: 8
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
