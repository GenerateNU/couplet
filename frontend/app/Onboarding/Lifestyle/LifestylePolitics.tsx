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
import { setPolitics } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import chunkArray from "../../../utils/chunkArray";
import { onButtonClick } from "../../../utils/onButtonClick";

const POLITICS_IMAGE = require("../../../assets/OnboardingPolitics.png");

function LifestylePolitics() {
  const dispatch = useAppDispatch();
  const [selectedButton, setSelectedButton] = useState("");

  const listOfPolitics = ["Liberal", "Moderate", "Conservative", "Other", "Prefer not to say"];
  const listReligionsChunked = chunkArray(listOfPolitics, 3);

  const { control, handleSubmit } = useForm({
    defaultValues: {
      politics: ""
    }
  });
  const onSubmit = (data: { politics: string }) => {
    dispatch(setPolitics(data.politics));
    router.push("Onboarding/Lifestyle/LifestyleHabits");
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
          <Image source={POLITICS_IMAGE} />
          <OnboardingTitle text="I believe in..." />
          <View style={scaledStyles.inputWrapper}>
            {listReligionsChunked.map((politicsRow, rowIndex) => (
              // eslint-disable-next-line react/no-array-index-key
              <View key={rowIndex} style={scaledStyles.buttonContainer}>
                {politicsRow.map((title: string, index: React.Key | null | undefined) => (
                  // eslint-disable-next-line react/no-array-index-key
                  <View key={index} style={scaledStyles.button}>
                    <Controller
                      control={control}
                      name="politics"
                      render={({ field: { onChange, value } }) => (
                        <OnboardingButton
                          title={title}
                          onButtonClick={() =>
                            onButtonClick(value, title, setSelectedButton, onChange)
                          }
                          isDisabled={Boolean(value && value !== title)}
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

export default LifestylePolitics;

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
