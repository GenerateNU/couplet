import { router } from "expo-router";
import React, { useState } from "react";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  StyleSheet,
  Text,
  TextInput,
  View
} from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const AT_ICON = require("../../assets/3diconslmao.png");

export default function Profile() {
  const [instagram, setInstagram] = useState("");

  return (
    <SafeAreaView style={{ height: "100%" }}>
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        style={scaledStyles.container}
      >
        <View>
          <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={0} />
        </View>
        <View>
          <Image source={AT_ICON} style={{ height: 250, width: 250 }} resizeMode="contain" />
          <View>
            <OnboardingTitle text="What’s your Instagram?" />
            <Text style={{ marginLeft: 0, top: 0 }}>
              This is how you will message your matches.
            </Text>
            <TextInput
              style={scaledStyles.textInput}
              editable
              onChangeText={setInstagram}
              placeholder="@couplet"
            />
          </View>
        </View>
        <View>
          <ContinueButton
            onPress={() => router.push("Onboarding/ProfileNotifications")}
            title="Continue"
            isDisabled={instagram.length === 0}
          />
        </View>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "space-between",
    alignItems: "center",
    margin: 30
  },
  textInput: {
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "#9EA3A2",
    color: "#000000",
    borderRadius: 10,
    padding: 10,
    fontFamily: "DMSansRegular"
  },
  disableNotificationsButton: {
    fontFamily: "DMSansMedium",
    textAlign: "center",
    marginTop: 15,
    fontSize: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
