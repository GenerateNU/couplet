import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, TextInput, View } from "react-native";
import { Icon } from "react-native-paper";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";

function AboutBirthday() {
  return (
    <View style={styles.container}>
      <View style={styles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="About Me"
          selectedCount={1}
        />
      </View>
      <View style={styles.mainContainer}>
        <View>
          <Image source={require("../../assets/calendarBirthday.png")} />
          <View>
            <Text style={styles.headerContainer}>My birthday is...</Text>
          </View>
          <View style={styles.inputWrapper}>
            <TextInput style={styles.inputContainer} placeholder="DD/MM/YYYY" />
            <View style={styles.icon}>
              <Icon source={require("../../assets/calendar.png")} size={15} />
            </View>
          </View>
          <Text style={styles.textHelper}>You won't be able to change this</Text>
        </View>
        <View style={styles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/AboutMe/AboutGender");
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
    flex: 0.3,
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
  inputContainer: {
    padding: 8
  },
  inputWrapper: {
    borderRadius: 6,
    borderWidth: 1,
    borderColor: "grey",
    marginBottom: 8
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
  icon: {
    position: "absolute",
    right: 10,
    bottom: 10
  }
});
