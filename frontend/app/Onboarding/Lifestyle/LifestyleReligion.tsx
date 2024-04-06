import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setReligion } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import chunkArray from "../../../utils/chunkArray";
import { onButtonClick } from "../../../utils/onButtonClick";

const RELIGION_IMAGE = require("../../../assets/OnboardingReligion.png");

function LifestyleReligion() {
  const dispatch = useAppDispatch();
  const [selectedButton, setSelectedButton] = useState("");

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
  const listReligionsChunked = chunkArray(listOfReligions, 3);

  const { control, handleSubmit } = useForm({
    defaultValues: {
      religion: ""
    }
  });
  const onSubmit = (data: { religion: string }) => {
    dispatch(setReligion(data.religion));
    router.push("Onboarding/Lifestyle/LifestylePolitics");
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
          <Image source={RELIGION_IMAGE} />
          <OnboardingTitle text="I believe in..." />
          <View style={scaledStyles.inputWrapper}>
            {listReligionsChunked.map((religionRow, rowIndex) => (
              // eslint-disable-next-line react/no-array-index-key
              <View key={rowIndex} style={scaledStyles.buttonContainer}>
                {religionRow.map((religion: string, index: React.Key | null | undefined) => (
                  // eslint-disable-next-line react/no-array-index-key
                  <View key={index} style={scaledStyles.button}>
                    <Controller
                      control={control}
                      name="religion"
                      render={({ field: { onChange, value } }) => (
                        <OnboardingButton
                          title={religion}
                          onButtonClick={() =>
                            onButtonClick(value, religion, setSelectedButton, onChange)
                          }
                          isDisabled={Boolean(value && value !== religion)}
                        />
                      )}
                    />
                  </View>
                ))}
              </View>
            ))}
          </View>
        </View>

        <View>
          <ContinueButton
            title="Continue"
            isDisabled={!selectedButton}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default LifestyleReligion;

const overrideStyles = {
  button: {
    marginRight: 8,
    marginBottom: 8
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start"
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
