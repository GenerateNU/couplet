import { router } from "expo-router";
import React from "react";
import { Image, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import COLORS from "../../../colors";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import onboardingStyles from "../../../styles/Onboarding/styles";

const NOTIFICATION_TOGGLE = require("../../../assets/notificationToggle.png");

function ProfileNotifications() {
  function goToNextPage() {
    router.push("Onboarding/Profile/ProfileConfirm");
  }

  const onAllowNotificationsPressed = () => {
    goToNextPage();
  };

  const onDisableNotificationsPressed = () => {
    goToNextPage();
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
        <View style={scaledStyles.notificationContainer}>
          <OnboardingTitle text="Turn on notifications" />
          <Text style={scaledStyles.textHelper}>Know when you get a match!</Text>
          <Image source={NOTIFICATION_TOGGLE} />
          <View style={scaledStyles.inputWrapper}>
            <ContinueButton
              title="Notify me"
              isDisabled={false}
              onPress={onAllowNotificationsPressed}
            />
          </View>
          <View style={scaledStyles.textInputWrapper}>
            <Text
              style={scaledStyles.disableNotificationText}
              onPress={onDisableNotificationsPressed}
            >
              Disable Notifications
            </Text>
          </View>
        </View>
      </View>
    </SafeAreaView>
  );
}

export default ProfileNotifications;

const overrideStyles = {
  notificationContainer: {
    alignItems: "center"
  },
  TopUiContainer: {
    alignItems: "center",
    flex: 0.45
  },
  textInputWrapper: {
    marginTop: 8
  },
  disableNotificationText: {
    fontSize: 15,
    fontWeight: "500",
    lineHeight: 20,
    letterSpacing: -0.15,
    fontFamily: "DMSansMedium",
    color: COLORS.darkGray
  },
  textHelper: {
    fontSize: 17,
    fontWeight: "400",
    lineHeight: 20,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium",
    color: COLORS.darkGray
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
