import { router } from "expo-router";
import React, { useState } from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const heightPicture = require("../../assets/height.png");

function AboutHeight() {
  const [openFeet, setOpenFeet] = useState(false);
  const [openInches, setOpenInches] = useState(false);
  const [foot, setFoot] = useState(null);
  const [inch, setInch] = useState(null);
  const feet = [1, 2, 3, 4, 5, 6, 7, 8].map((feetParam, index) => ({
    label: `${feetParam}`,
    value: index + 1
  }));
  const inches = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].map((inchParam, index) => ({
    label: `${inchParam}`,
    value: index + 1
  }));
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
    router.push("/AboutMe/AboutLocation");
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
          <View style={scaledStyles.dropDownContainer}>
            <DropDownPicker
              open={openFeet}
              value={foot}
              items={feet}
              setOpen={setOpenFeet}
              setValue={setFoot}
              placeholder="Feet"
              containerStyle={scaledStyles.dropdown}
            />
            <DropDownPicker
              open={openInches}
              value={inch}
              items={inches}
              setOpen={setOpenInches}
              setValue={setInch}
              placeholder="Inches"
              containerStyle={scaledStyles.dropdown}
            />
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

export default AboutHeight;

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
