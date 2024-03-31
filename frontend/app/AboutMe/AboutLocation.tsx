import { router } from "expo-router";
import React, { useState } from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import DropDownPicker from "react-native-dropdown-picker";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";

const aboutLocationPicture = require("../../assets/aboutlocation.png");
function AboutLocation() {
  const { control, handleSubmit } = useForm({
    defaultValues: {
      name: ""
    }
  });
  const name = useWatch({
    control,
    name: "name",
    defaultValue: ""
  });
  const onSubmit = (data: Object) => {
    console.log(name);
    router.push("/Onboarding/Education");
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
          <Image source={aboutLocationPicture} />
          <OnboardingTitle text="I live in..."/>
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

export default AboutLocation;

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
  },
  button: {
    marginBottom: 14
  },
  dropDownContainer: {
    flexDirection: "row"
  },
  dropdown: {
    flex: 1,
    marginRight: 5
  }
});

const scaledStyles = scaleStyleSheet(styles);

