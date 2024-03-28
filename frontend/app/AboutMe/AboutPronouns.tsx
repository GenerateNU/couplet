import { router } from "expo-router";
import React from "react";
import { Image, SafeAreaView, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const pronounPicture = require("../../assets/pronouns.png");

function AboutPronouns() {
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
          <Image source={pronounPicture} />
          <View>
            <Text style={scaledStyles.headerContainer}>My pronouns are...</Text>
          </View>
          <View>
            <View style={scaledStyles.buttonContainer}>
              <TouchableOpacity style={styles.button}>
                <Text style={scaledStyles.buttonText}> He/Him </Text>
              </TouchableOpacity>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> She/Her </Text>
              </TouchableOpacity>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> They/Them </Text>
              </TouchableOpacity>
            </View>
            <View style={scaledStyles.buttonContainer}>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> He/They </Text>
              </TouchableOpacity>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> She/They </Text>
              </TouchableOpacity>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> Xe/Xem </Text>
              </TouchableOpacity>
            </View>
            <View style={scaledStyles.buttonContainer}>
              <TouchableOpacity style={scaledStyles.button}>
                <Text style={scaledStyles.buttonText}> Ze/Zir </Text>
              </TouchableOpacity>
            </View>
          </View>
        </View>
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              router.push("/AboutMe/AboutHeight");
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutPronouns;

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
    borderColor: "#E7D4FA",
    borderRadius: 25,
    justifyContent: "center",
    alignItems: "center",
    padding: 12,
    marginRight: 8,
    marginBottom: 8
  },
  buttonText: {
    color: "black",
    fontSize: 15,
    fontWeight: "400",
    letterSpacing: -0.15,
    fontFamily: "DMSansMedium"
  },
  buttonContainer: {
    flexDirection: "row",
    alignItems: "center",
    justifyContent: "flex-start"
  }
});

const scaledStyles = scaleStyleSheet(styles);
