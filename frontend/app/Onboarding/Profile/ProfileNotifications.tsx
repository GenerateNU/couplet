import { router } from "expo-router";
import React from "react";
import { Image, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { createUser } from "../../../api/users";
import COLORS from "../../../colors";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setNotifications } from "../../../state/formSlice";
import { useAppDispatch, useAppSelector } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";
import calculateAge from "../../../utils/calculateAge";

const NOTIFICATION_TOGGLE = require("../../../assets/notificationToggle.png");

interface userDataProps {
  firstName: string;
  lastName: string;
  age: number;
  bio: string;
  images: string[];
}

function ProfileNotifications() {
  const dispatch = useAppDispatch();
  const userState = useAppSelector((state) => state.form);
  async function goToNextPage() {
    const userData: userDataProps = {
      firstName: userState.name,
      lastName: "GET LAST NAME FROM AUTH", // Weird bug, if this is empty will throw an error
      age: calculateAge(new Date(userState.birthday)), // TODO get age from auth
      bio: userState.promptBio,
      images: userState.photos.map((photo) => photo.filePath)
    };
    try {
      const res = await createUser(userData);
      console.log(res);
    } catch (e) {
      if (e instanceof Error) {
        throw new Error(e.message);
      }
    }
    router.push("Onboarding/Profile/ProfileConfirm");
  }
  const onAllowNotificationsPressed = async () => {
    dispatch(setNotifications(true));
    await goToNextPage();
  };
  const onDisableNotificationsPressed = async () => {
    dispatch(setNotifications(false));
    await goToNextPage();
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
