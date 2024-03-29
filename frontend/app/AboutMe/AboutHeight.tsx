import { router } from "expo-router";
import React, { useState } from "react";
import { Image, SafeAreaView, StyleSheet, Text, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import ContinueButton from "../../components/Onboarding/ContinueButton";
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
          <View>
            <Text style={scaledStyles.headerContainer}>My height is...</Text>
          </View>
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
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/AboutMe/AboutLocation");
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
    flex: 0.3,
    alignItems: "center"
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
  ContinueButtonContainer: {
    marginBottom: 10
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
