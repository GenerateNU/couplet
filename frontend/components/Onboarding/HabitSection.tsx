import React from "react";
import { Control, Controller } from "react-hook-form";
import { StyleSheet, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";
import onboardingStyles from "../../styles/Onboarding/styles";
import chunkArray from "../../utils/chunkArray";
import { onButtonClick } from "../../utils/onButtonClick";
import OnboardingButton from "./OnboardingButton";
import OnboardingSmallTitle from "./OnboardingSmallTitle";

interface HabitSectionProps {
  title: string;
  options: string[];
  disableBar: boolean;
  parentControl: Control<
    {
      drinkHabit: string;
      smokeHabit: string;
      weedHabit: string;
      drugHabit: string;
    },
    any
  >;
  habit: "drinkHabit" | "smokeHabit" | "weedHabit" | "drugHabit";
  setHabit: (habit: string) => void;
}

function HabitSection({
  title,
  options,
  disableBar,
  parentControl,
  habit,
  setHabit
}: HabitSectionProps) {
  const chunkOptions = chunkArray(options, 3);
  return (
    <View>
      <View style={scaledStyles.titleContainer}>
        <OnboardingSmallTitle text={title} />
      </View>
      {chunkOptions.map((row, rowIndex) => (
        // eslint-disable-next-line react/no-array-index-key
        <View key={rowIndex} style={scaledStyles.buttonContainer}>
          {row.map((buttonTitle: string, index: React.Key | null | undefined) => (
            // eslint-disable-next-line react/no-array-index-key
            <View key={index} style={scaledStyles.button}>
              <Controller
                control={parentControl}
                name={habit}
                render={({ field: { onChange, value } }) => (
                  <OnboardingButton
                    title={buttonTitle}
                    onButtonClick={() => {
                      onButtonClick(value, buttonTitle, () => {}, onChange);
                      setHabit(buttonTitle);
                    }}
                    isDisabled={Boolean(value && value !== buttonTitle)}
                  />
                )}
              />
            </View>
          ))}
        </View>
      ))}
      {!disableBar && <View style={scaledStyles.line} />}
    </View>
  );
}

const overrideStyles = StyleSheet.create({
  button: {
    marginRight: 8,
    marginBottom: 8
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start"
  },
  line: {
    marginTop: 16,
    borderColor: "#CDCDCD",
    borderWidth: 0.75
  },
  titleContainer: {
    marginBottom: 16,
    marginTop: 16
  }
});

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });

export default HabitSection;
