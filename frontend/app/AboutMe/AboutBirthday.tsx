import { router } from "expo-router";
import React from "react";
import { Controller, useForm } from "react-hook-form";
import { Image, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import DropDownCalendar from "../../components/Onboarding/DropDownCalendar";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import { setBirthday } from "../../state/formSlice";
import { useAppDispatch } from "../../state/hooks";

const aboutBirthdayPicture = require("../../assets/calendarBirthday.png");

function AboutBirthday() {
  const dispatch = useAppDispatch();
  const { handleSubmit, setValue } = useForm({
    defaultValues: {
      birthday: new Date()
    }
  });
  const handleDateChange = (day: number, month: number, year: number) => {
    setValue("birthday", new Date(year, month - 1, day));
  };
  const onSubmit = (data: { birthday: Date }) => {
    //Store it as a string to satisfy Redux's required serialization values
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
          <DropDownCalendar onDateChange={handleDateChange} />
          <View style={scaledStyles.helperContainer}>
            <Text style={scaledStyles.textHelper}>You won&apos;t be able to change this</Text>
          </View>
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

const styles = StyleSheet.create({
  TopUiContainer: {
    alignItems: "center",
    flex: 0.3
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1,
    marginTop: 34,
    marginBottom: 36
  },
  helperContainer: {
    marginTop: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
