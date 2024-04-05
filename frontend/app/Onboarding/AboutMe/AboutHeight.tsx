import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import DropDownHeightPicker from "../../../components/Onboarding/DropDownHeightPicker";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setHeight } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const heightPicture = require("../../../assets/height.png");

function AboutHeight() {
  const dispatch = useAppDispatch();
  const [isHeightSelected, setIsHeightSelected] = useState(false);
  const handleHeightChange = (foot: number, inch: number) => {
    setIsHeightSelected(foot > 0 && inch >= 0);
  };
  const { control, handleSubmit } = useForm({
    defaultValues: {
      height: { foot: 0, inch: 0 }
    }
  });
  const onSubmit = (data: { height: { foot: number; inch: number } }) => {
    dispatch(setHeight(data.height));
    router.push("Onboarding/AboutMe/AboutLocation");
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
          <Image source={heightPicture} />
          <OnboardingTitle text="My height is..." />
          <View style={scaledStyles.inputWrapper}>
            <Controller
              control={control}
              name="height"
              render={({ field: { onChange, value } }) => (
                <DropDownHeightPicker
                  onHeightChange={(foot: number, inch: number) => {
                    onChange({ foot, inch });
                    handleHeightChange(foot, inch);
                  }}
                  selectedHeight={value}
                />
              )}
            />
          </View>
        </View>
        <View>
          <ContinueButton
            title="Continue"
            isDisabled={!isHeightSelected}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutHeight;

const scaledStyles = scaleStyleSheet(onboardingStyles);
