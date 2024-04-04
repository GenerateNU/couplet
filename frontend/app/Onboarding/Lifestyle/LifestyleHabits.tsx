import { router } from "expo-router";
import React, { useState } from "react";
import { useForm } from "react-hook-form";
import { Image, ScrollView, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import HabitSection from "../../../components/Onboarding/HabitSection";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setDrinkHabit, setDrugHabit, setSmokeHabit, setWeedHabit } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const HABITS_IMAGE = require("../../../assets/OnboardingHabits.png");

function LifestyleHabits() {
  const genericHabits = ["Yes", "No", "Sometimes", "Socially", "Prefer not to say"];
  const drugHabits = ["Yes", "No", "Prefer not to say"];
  const dispatch = useAppDispatch();
  const [selectDrinkHabit, setSelectDrinkHabit] = useState("");
  const [selectSmokeHabit, setSelectSmokeHabit] = useState("");
  const [selectWeedHabit, setSelectWeedHabit] = useState("");
  const [selectDrugHabit, setSelectDrugHabit] = useState("");
  const { control, handleSubmit } = useForm({
    defaultValues: {
      drinkHabit: "",
      smokeHabit: "",
      weedHabit: "",
      drugHabit: ""
    }
  });
  const onSubmit = (data: {
    drinkHabit: string;
    smokeHabit: string;
    weedHabit: string;
    drugHabit: string;
  }) => {
    dispatch(setDrinkHabit(data.drinkHabit));
    dispatch(setSmokeHabit(data.smokeHabit));
    dispatch(setWeedHabit(data.weedHabit));
    dispatch(setDrugHabit(data.drugHabit));
    router.push("Onboarding/Lifestyle/LifestylePassions");
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
          <Image source={HABITS_IMAGE} />
          <OnboardingTitle text="What are your habits?" />
          <View style={scaledStyles.inputWrapper}>
            <ScrollView style={scaledStyles.habitWindow}>
              <HabitSection
                title="Do you drink?"
                options={genericHabits}
                disableBar={false}
                parentControl={control}
                habit="drinkHabit"
                setHabit={setSelectDrinkHabit}
              />
              <HabitSection
                title="Do you smoke?"
                options={genericHabits}
                disableBar={false}
                parentControl={control}
                habit="smokeHabit"
                setHabit={setSelectSmokeHabit}
              />
              <HabitSection
                title="Do you smoke weed?"
                options={genericHabits}
                disableBar={false}
                parentControl={control}
                habit="weedHabit"
                setHabit={setSelectWeedHabit}
              />
              <HabitSection
                title="Do you do drugs?"
                options={drugHabits}
                disableBar
                parentControl={control}
                habit="drugHabit"
                setHabit={setSelectDrugHabit}
              />
            </ScrollView>
          </View>
        </View>

        <View>
          <ContinueButton
            title="Continue"
            isDisabled={
              !selectDrinkHabit || !selectDrugHabit || !selectSmokeHabit || !selectWeedHabit
            }
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default LifestyleHabits;

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
    height: "55%"
  },
  inputWrapper: {
    marginTop: 0
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
