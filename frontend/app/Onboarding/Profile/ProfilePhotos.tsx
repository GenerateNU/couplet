import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import { Image, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import PhotoPicker from "../../../components/PhotoPicker";
import scaleStyleSheet from "../../../scaleStyles";
import { setPhotos } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const CAMERA_IMAGE = require("../../../assets/profilecamera.png");

function ProfilePhotos() {
  const dispatch = useAppDispatch();
  const [images, setImages] = useState<string[]>(["", "", "", ""]);
  const [continueText, setContinueText] = useState<string>("Continue");
  const [continueEnabled, setContinueEnabled] = useState<boolean>(true);
  useEffect(() => {
    setContinueEnabled(
      images[0] === "" || images[1] === "" || images[2] === "" || images[3] === ""
    );
    const imgCount = images.filter((img) => img !== "").length;
    if (imgCount === 4) {
      setContinueText("Continue");
    } else {
      setContinueText(`${imgCount}/4 Added`);
    }
  }, [images]);

  const onContinue = () => {
    dispatch(setPhotos(images.map((i) => ({ filePath: i, caption: "" }))));
    router.push("Onboarding/Profile/ProfileCaptions");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="Profile"
          selectedCount={4}
        />
      </View>
      <View style={scaledStyles.mainContainer}>
        <View>
          <Image source={CAMERA_IMAGE} />
          <OnboardingTitle text="Dont be shy, show your best angles!" />
          <View style={scaledStyles.inputWrapper}>
            <PhotoPicker onPick={setImages} />
          </View>
        </View>
      </View>
      <View>
        <ContinueButton title={continueText} isDisabled={continueEnabled} onPress={onContinue} />
      </View>
    </SafeAreaView>
  );
}

export default ProfilePhotos;

const overrideStyles = {
  TopUiContainer: {
    alignItems: "center",
    flex: 0.15
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
