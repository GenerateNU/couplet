import { router } from "expo-router";
import React, { useState } from "react";
import { Image, SafeAreaView, StyleSheet, Text, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

function AboutHeight() {
    const [openLocation, setOpenLocation] = useState(false);
    const [location, setLocation] = useState(null);
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
          <Image source={require("../../assets/aboutlocation.png")} />
          <View>
            <Text style={scaledStyles.headerContainer}>I live in...</Text>
          </View>
          <View>
          <DropDownPicker
              open={openLocation}
              value={location}
              items={[]}
              setOpen={setOpenLocation}
              setValue={setLocation}
              placeholder="Select Your Location"
              containerStyle={scaledStyles.dropdown}
            />
          </View>
        </View>
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/Home");
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
  button: {
    backgroundColor: "#FFF",
    borderWidth: 1,
    borderColor: "#D1AAF6",
    borderRadius: 25,
    height: 40,
    justifyContent: "center",
    alignItems: "center",
    marginBottom: 16
  },
  buttonText: {
    color: "black",
    fontSize: 17,
    fontWeight: "500",
    letterSpacing: -0.17,
    fontFamily: "DMSansMedium"
  },
  dropDownContainer: {
    flexDirection : 'row',
  },
  dropdown: {
    flex: 1,
    marginRight : 5,
  }
});

const scaledStyles = scaleStyleSheet(styles);
