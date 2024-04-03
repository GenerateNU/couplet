import { router } from "expo-router";
import React from "react";
import { Image, SafeAreaView, StyleSheet, Text, TouchableOpacity, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const NOTIFICATION_TOGGLE = require("../../assets/notification-toggle.png");
const NOTIFICATION_ICON = require("../../assets/notification-icon.png");

export default function ProfileNotifications() {
  function goToNextPage() {
    // console.log("Go to next page");
    router.push("Onboarding/ProfileConfirm");
  }

  function onAllowNotifactionsPressed() {
    // console.log("Notify me pressed");
    goToNextPage();
  }

  function onDisableNotificationsPressed() {
    // console.log("Disable notificationspressed");
    goToNextPage();
  }

  return (
    <SafeAreaView style={{ height: "100%" }}>
      <View style={scaledStyles.container}>
        <View>
          <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={0} />
        </View>
        <View>
          <Image
            source={NOTIFICATION_TOGGLE}
            style={{ height: 250, width: 250 }}
            resizeMode="contain"
          />
          <View style={scaledStyles.titleMessage}>
            <Image source={NOTIFICATION_ICON} style={scaledStyles.bellIcon} resizeMode="contain" />
            <OnboardingTitle text="Turn on notifications" />
          </View>
          <Text style={scaledStyles.notificationMessage}> Know when you get a match! </Text>
          <View>
            <ContinueButton
              onPress={() => onAllowNotifactionsPressed()}
              title="Notify Me"
              isDisabled={false}
            />
          </View>
          <TouchableOpacity onPress={() => onDisableNotificationsPressed()} disabled={false}>
            <Text style={scaledStyles.disableNotificationsButton}>Disable Notifications</Text>
          </TouchableOpacity>
        </View>
        <View />
        <View />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    margin: 30,
    alignItems: "center",
    justifyContent: "space-between"
  },
  titleMessage: {
    flexDirection: "row"
  },
  bellIcon: {
    marginTop: 15
  },
  notificationMessage: {
    fontFamily: "DMSansRegular",
    textAlign: "center",
    marginBottom: 15
  },
  disableNotificationsButton: {
    fontFamily: "DMSansMedium",
    textAlign: "center",
    marginTop: 15,
    fontSize: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
