import { router } from "expo-router";
import React, { useState } from "react";
import { Image, StyleSheet, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import PhotoPicker from "../../components/PhotoPicker";
import scaleStyleSheet from "../../scaleStyles";

const CAMERA_IMAGE = require("../../assets/profilecamera.png");

export default function ProfilePhotos() {
  const [images, setImages] = useState<string[]>([]);

  const onContinue = () => {
    // TODO: Save images into form state
    router.push("Onboarding/ProfileInsta");
  };

  return (
    <View style={scaledStyles.container}>
      <View>
        <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={5} />
      </View>
      <View>
        <Image source={CAMERA_IMAGE} />
        <OnboardingTitle text="Show your best angles" />
        <PhotoPicker onPick={setImages} />
      </View>
      <View>
        <ContinueButton title="Continue" isDisabled={images.length !== 4} onPress={onContinue} />
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: "space-between",
    alignItems: "center",
    margin: 30
  }
});

const scaledStyles = scaleStyleSheet(styles);
