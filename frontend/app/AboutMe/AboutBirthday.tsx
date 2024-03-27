import { router } from "expo-router";
import React from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, Text, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import DropDownCalendar from "../../components/Onboarding/DropDownCalendar";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

function AboutBirthday() {
  
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
    router.push("/AboutMe/AboutBirthday");
  };
  return (
    <View style={scaledStyles.container}>
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
          <Image source={require("../../assets/calendarBirthday.png")} />
          <Text style={scaledStyles.headerContainer}>My birthday is...</Text>
          <View style={scaledStyles.inputWrapper}></View>
          <DropDownCalendar />
          <View style={scaledStyles.helperContainer}>
          <Text style={scaledStyles.textHelper}>You won't be able to change this</Text>
          </View>
        </View>

        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={true}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </View>
  );
}

export default AboutBirthday;

const styles = StyleSheet.create({
  TopUiContainer: {
    flex: 0.3
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  headerContainer: {
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    letterSpacing: -0.32,
    marginTop: 16,
    marginBottom: 16,
    fontFamily: "DMSansMedium"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1
  },
  helperContainer: {
    marginTop : 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
