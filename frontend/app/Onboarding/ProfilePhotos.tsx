import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import { Image, SafeAreaView, StyleSheet, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import PhotoPicker from "../../components/PhotoPicker";
import scaleStyleSheet from "../../scaleStyles";

const CAMERA_IMAGE = require("../../assets/profilecamera.png");

export default function ProfilePhotos() {
  const [images, setImages] = useState<string[]>(["", "", "", ""]);
  const [continueText, setContinueText] = useState<string>("Continue");
  const [continueEnabled, setContinueEnabled] = useState<boolean>(true);

  useEffect(() => {
    setContinueEnabled(
      images[0] === "" || images[1] === "" || images[2] === "" || images[3] === ""
    );

    const imgCount = images.filter((img) => img !== "").length;
    if (imgCount > 0) {
      setContinueText(`${imgCount}/4 Continue`);
    } else {
      setContinueText("Continue");
    }
  }, [images]);

  const onContinue = () => {
    // TODO: Save images into form state
    router.push("Onboarding/ProfileCaptions");
  };

  return (
    <SafeAreaView style={{ height: "100%" }}>
      <View style={scaledStyles.container}>
        <View>
          <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={4} />
        </View>
        <View>
          <Image source={CAMERA_IMAGE} style={{ alignSelf: "center" }} />
          <OnboardingTitle text="Show your best angles" />
          <PhotoPicker onPick={setImages} />
        </View>
        <View>
          <ContinueButton title={continueText} isDisabled={continueEnabled} onPress={onContinue} />
        </View>
      </View>
    </SafeAreaView>
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
