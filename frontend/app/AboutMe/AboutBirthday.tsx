import { router } from "expo-router";
import React, { useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import DropDownCalendar from "../../components/Onboarding/DropDownCalendar";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import { setBirthday } from "../../state/formSlice";
import { useAppDispatch } from "../../state/hooks";
import onboardingStyles from "../../styles/Onboarding/styles";

const aboutBirthdayPicture = require("../../assets/calendarBirthday.png");

function AboutBirthday() {
  const dispatch = useAppDispatch();
  const [isDropDownOpen, setIsDropDownOpen] = useState(false);
  const { control, handleSubmit } = useForm({
    defaultValues: {
      birthday: new Date()
    }
  });
  const handleDropDownOpen = (openDay: boolean, openMonth: boolean, openYear: boolean) => {
    const isOpen = openDay || openMonth || openYear;
    setIsDropDownOpen(isOpen);
  };
  const onSubmit = (data: { birthday: Date }) => {
    // console.log(data);
    // Store it as a string to satisfy Redux's required serialization values
    dispatch(setBirthday(data.birthday.toISOString()));
    router.push("/AboutMe/AboutGender");
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
          <Image source={aboutBirthdayPicture} />
          <OnboardingTitle text="My birthday is..." />
          <View style={scaledStyles.inputWrapper} />
          <Controller
            control={control}
            name="birthday"
            render={({ field: { onChange, value } }) => (
              <DropDownCalendar
                onDateChange={(day, month, year) => onChange(new Date(year, month - 1, day))}
                onDropDownOpen={handleDropDownOpen}
                selectedDate={value}
              />
            )}
          />
          {!isDropDownOpen && (
            <View style={scaledStyles.helperContainer}>
              <Text style={scaledStyles.textHelper}>You won&apos;t be able to change this</Text>
            </View>
          )}
        </View>
        <View>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutBirthday;

const scaledStyles = scaleStyleSheet(onboardingStyles);
