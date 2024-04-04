import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import { setPronouns } from "../../state/formSlice";
import { useAppDispatch } from "../../state/hooks";
import onboardingStyles from "../../styles/Onboarding/styles";
import onButtonClick from "../../utils/onButtonClick";
import {
  PRONOUNS_HE_HIM,
  PRONOUNS_HE_THEY,
  PRONOUNS_PREFER_NOT_TO_SAY,
  PRONOUNS_SHE_HER,
  PRONOUNS_SHE_THEY,
  PRONOUNS_THEY_THEM,
  PRONOUNS_XE_XEM,
  PRONOUNS_ZE_ZIR
} from "../../utils/pronouns";

const pronounPicture = require("../../assets/pronouns.png");

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
    router.push("/AboutMe/AboutHeight");
  };
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
          <View style={scaledStyles.buttonContainer}>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_HE_HIM}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_HE_HIM, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_HE_HIM)}
                  />
                )}
              />
            </View>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_SHE_HER}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_SHE_HER, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_SHE_HER)}
                  />
                )}
              />
            </View>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_THEY_THEM}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_THEY_THEM, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_THEY_THEM)}
                  />
                )}
              />
            </View>
          </View>
          <View style={scaledStyles.buttonContainer}>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_HE_THEY}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_HE_THEY, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_HE_THEY)}
                  />
                )}
              />
            </View>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_SHE_THEY}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_SHE_THEY, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_SHE_THEY)}
                  />
                )}
              />
            </View>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_XE_XEM}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_XE_XEM, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_XE_XEM)}
                  />
                )}
              />
            </View>
          </View>
          <View style={scaledStyles.buttonContainer}>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_ZE_ZIR}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_ZE_ZIR, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_ZE_ZIR)}
                  />
                )}
              />
            </View>
            <View style={scaledStyles.button}>
              <Controller
                control={control}
                name="pronouns"
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={PRONOUNS_PREFER_NOT_TO_SAY}
                    onButtonClick={() =>
                      onButtonClick(value, PRONOUNS_PREFER_NOT_TO_SAY, setSelectedButton, onChange)
                    }
                    isDisabled={Boolean(value && value !== PRONOUNS_PREFER_NOT_TO_SAY)}
                  />
                )}
              />
            </View>
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
