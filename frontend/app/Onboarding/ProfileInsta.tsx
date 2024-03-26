import { router } from "expo-router";
import React, { useState } from "react";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  StyleSheet,
  Text,
  TextInput,
  View
} from "react-native";
import BackButton from "../../components/Onboarding/BackButton";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingBar from "../../components/Onboarding/OnboardingBar";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";

const AT_ICON = require("../../assets/3diconslmao.png");

export default function Profile() {
  const [instagram, setInstagram] = useState("");

  function goToNextPage() {
    console.log("Go to next page");
    router.push("Profile");
  }

  function onAllowNotifactionsPressed() {
    console.log("Notify me pressed");
    goToNextPage();
  }

  function onDisableNotificationsPressed() {
    console.log("Disable notificationspressed");
    goToNextPage();
  }

  return (
    <KeyboardAvoidingView
      behavior={Platform.OS === "ios" ? "padding" : "height"}
      style={{ flex: 1 }}
    >
      <View style={{ marginLeft: 15, marginRight: 15, marginTop: 15 }}>
        <BackButton onPress={() => {}} />
        <Text style={{ fontFamily: "DMSansRegular", textAlign: "center" }}>Profile</Text>
        <OnboardingBar selectedCount={5} />
      </View>
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        style={{ marginLeft: 15, marginRight: 15, marginTop: "10%", flex: 1 }}
      >
        <Image source={AT_ICON} style={{ height: 250, width: 250 }} resizeMode="contain" />
        <View style={{ flexDirection: "col", width: "90%", marginLeft: 15 }}>
          <View style={{ marginLeft: 15 }}>
            <OnboardingTitle text="What’s your Instagram????" />
          </View>
          <Text style={{ marginLeft: 0, top: 0 }}>This is how you will message your matches.</Text>
          <View>
            <TextInput
              style={{
                borderStyle: "solid",
                borderWidth: 1,
                borderColor: "#9EA3A2",
                color: "#000000",
                borderRadius: 10,
                padding: 15,
                fontFamily: "DMSansRegular"
              }}
              editable
              onChangeText={(text) => setInstagram(text)}
              placeholder="@couplet"
            />
          </View>
        </View>
      </KeyboardAvoidingView>
      <View style={{ bottom: 0, position: "absolute", height: "100vh", width: "100%" }}>
        <ContinueButton
          onPress={() => onAllowNotifactionsPressed()}
          title="Continue"
          isDisabled={instagram.length == 0}
        />
      </View>
    </KeyboardAvoidingView>
  );
}

const styles = StyleSheet.create({
  notificationMessage: {
    fontFamily: "DMSansRegular",
    textAlign: "center",
    marginTop: 20,
    marginBottom: 30
  },
  disableNotificationsButton: {
    fontFamily: "DMSansMedium",
    textAlign: "center",
    marginTop: 15,
    fontSize: 16
  }
});
