import { router } from "expo-router";
import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import BackButton from "../../components/Onboarding/BackButton";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingBar from "../../components/Onboarding/OnboardingBar";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";

const NOTIFICATION_TOGGLE = require("../../assets/notification-toggle.png");
const NOTIFICATION_ICON = require("../../assets/notification-icon.png");

export default function ProfileNotifications() {
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
    <View style={{ flex: 1 }}>
      <View style={{ marginLeft: 15, marginRight: 15, marginTop: 15 }}>
        <BackButton onPress={() => {}} />
        <Text style={{ fontFamily: "DMSansRegular", textAlign: "center" }}>Profile</Text>
        <OnboardingBar selectedCount={5} />
      </View>
      <View style={{ marginLeft: 15, marginRight: 15 }}>
        <Image
          source={NOTIFICATION_TOGGLE}
          style={{ height: 250, width: 250 }}
          resizeMode="contain"
        />
        <View style={{ flexDirection: "row", width: "93%" }}>
          <Image
            source={NOTIFICATION_ICON}
            style={{ height: 25, width: 25, marginTop: 5 }}
            resizeMode="contain"
          />
          <OnboardingTitle text="Turn on notifications" />
        </View>

        <Text style={styles.notificationMessage}> Know when you get a match! </Text>
        <ContinueButton
          onPress={() => onAllowNotifactionsPressed()}
          title="Notify Me"
          isDisabled={false}
        />
        <TouchableOpacity onPress={() => onDisableNotificationsPressed()} disabled={false}>
          <Text style={styles.disableNotificationsButton}>Disable Notifications</Text>
        </TouchableOpacity>
      </View>
    </View>
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
