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
import { setPronouns } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import chunkArray from "../../../utils/chunkArray";
import { onButtonClick } from "../../../utils/onButtonClick";
import {
  PRONOUNS_HE_HIM,
  PRONOUNS_HE_THEY,
  PRONOUNS_PREFER_NOT_TO_SAY,
  PRONOUNS_SHE_HER,
  PRONOUNS_SHE_THEY,
  PRONOUNS_THEY_THEM,
  PRONOUNS_XE_XEM,
  PRONOUNS_ZE_ZIR
} from "../../../utils/pronouns";

const pronounPicture = require("../../../assets/pronouns.png");

function AboutPronouns() {
  const dispatch = useAppDispatch();
  const [selectedButton, setSelectedButton] = useState("");
  const { control, handleSubmit } = useForm({
    defaultValues: {
      pronouns: ""
    }
  });
  const onSubmit = (data: { pronouns: string }) => {
    dispatch(setPronouns(data.pronouns));
    router.push("Onboarding/AboutMe/AboutHeight");
  };
  const PRONOUNS = [
    { title: PRONOUNS_HE_HIM },
    { title: PRONOUNS_SHE_HER },
    { title: PRONOUNS_THEY_THEM },
    { title: PRONOUNS_HE_THEY },
    { title: PRONOUNS_SHE_THEY },
    { title: PRONOUNS_XE_XEM },
    { title: PRONOUNS_ZE_ZIR },
    { title: PRONOUNS_PREFER_NOT_TO_SAY }
  ];
  // Chunk the PRONOUNS array into sub-arrays of size 3
  const PRONOUNS_CHUNKED = chunkArray(PRONOUNS, 3);

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
      <View style={scaledStyles.mainContainer}>
        <View>
          <Image source={pronounPicture} />
          <OnboardingTitle text="My pronouns are..." />
          <View style={scaledStyles.inputWrapper}>
            {PRONOUNS_CHUNKED.map((pronounsRow, rowIndex) => (
              // eslint-disable-next-line react/no-array-index-key
              <View key={rowIndex} style={scaledStyles.buttonContainer}>
                {pronounsRow.map(
                  (pronoun: { title: string }, index: React.Key | null | undefined) => (
                    // eslint-disable-next-line react/no-array-index-key
                    <View key={index} style={scaledStyles.button}>
                      <Controller
                        control={control}
                        name="pronouns"
                        render={({ field: { onChange, value } }) => (
                          <OnboardingButton
                            title={pronoun.title}
                            onButtonClick={() =>
                              onButtonClick(value, pronoun.title, setSelectedButton, onChange)
                            }
                            isDisabled={Boolean(value && value !== pronoun.title)}
                          />
                        )}
                      />
                    </View>
                  )
                )}
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

export default AboutPronouns;

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
