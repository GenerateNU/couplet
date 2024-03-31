import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import {
  Image,
  KeyboardAvoidingView,
  Platform,
  SafeAreaView,
  ScrollView,
  StyleSheet,
  TextInput,
  View
} from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const CAPTIONS = require("../../assets/profilecaptions.png");

export default function ProfileCaptions() {
  // Sample images for now, these are my local uris, replace if you want to test it out
  const images: string[] = [
    "file:///var/mobile/Containers/Data/Application/4A88A83D-ED82-43E2-8200-765E6547AA82/Library/Caches/ExponentExperienceData/@anonymous/myApp-da088f0a-c6dd-4e77-aa0f-f562e3466c64/ImagePicker/2BCF4800-CA67-438A-8239-B942084C6EF3.jpg",
    "file:///var/mobile/Containers/Data/Application/4A88A83D-ED82-43E2-8200-765E6547AA82/Library/Caches/ExponentExperienceData/@anonymous/myApp-da088f0a-c6dd-4e77-aa0f-f562e3466c64/ImagePicker/F81D349A-DCE2-43E3-91A5-8DCE8FAB8178.jpg",
    "file:///var/mobile/Containers/Data/Application/4A88A83D-ED82-43E2-8200-765E6547AA82/Library/Caches/ExponentExperienceData/@anonymous/myApp-da088f0a-c6dd-4e77-aa0f-f562e3466c64/ImagePicker/F1A713BF-0B8E-4CD4-8C09-9E2FED574AAE.jpg",
    "file:///var/mobile/Containers/Data/Application/4A88A83D-ED82-43E2-8200-765E6547AA82/Library/Caches/ExponentExperienceData/@anonymous/myApp-da088f0a-c6dd-4e77-aa0f-f562e3466c64/ImagePicker/A912B59B-02CD-4F93-815F-6286379F61DA.jpg"
  ];
  const [captions, setCaptions] = useState(["", "", "", ""]);
  const [continueDisabled, setContinueDisabled] = useState<boolean>(true);

  useEffect(() => {
    const captionCount = captions.filter((cap) => cap !== "").length;
    setContinueDisabled(captionCount !== 4);
  }, [captions]);

  const onContinue = () => {
    // TODO: Save captions into form state
    router.push("Onboarding/ProfileInsta");
  };

  const onSubmitCaption = (text: string, index: number) => {
    const newCaptions = [...captions];
    newCaptions[index] = text;
    setCaptions(newCaptions);
  };

  return (
    <SafeAreaView style={{ height: "100%" }}>
      {/* Does not work very well, weird delay */}
      <KeyboardAvoidingView
        behavior={Platform.OS === "ios" ? "padding" : "height"}
        style={{ height: "100%" }}
      >
        <ScrollView contentContainerStyle={scaledStyles.scrollContainer}>
          <View style={{ marginTop: 25 }}>
            <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={5} />
          </View>
          <ScrollView>
            {/* Spacer */}
            <View style={{ height: 100 }} />
            <Image source={CAPTIONS} />
            <OnboardingTitle text="Add a caption" />
            {images.map((img, i) => (
              <View>
                <Image source={{ uri: img }} style={scaledStyles.imageStyle} />
                <TextInput
                  style={scaledStyles.captionStyle}
                  onSubmitEditing={(e) => onSubmitCaption(e.nativeEvent.text, i)}
                  editable
                  placeholder="Caption"
                />
              </View>
            ))}
          </ScrollView>
          <View style={{ marginBottom: 25 }}>
            <ContinueButton title="Continue" isDisabled={continueDisabled} onPress={onContinue} />
          </View>
        </ScrollView>
      </KeyboardAvoidingView>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  scrollContainer: {
    alignItems: "center"
  },
  imageStyle: {
    height: 200,
    width: "100%",
    alignSelf: "center",
    borderRadius: 5
  },
  captionStyle: {
    width: "100%",
    height: 30,
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "#9EA3A2",
    color: "#000000",
    borderRadius: 5,
    padding: 10,
    marginTop: 10,
    marginBottom: 20,
    fontFamily: "DMSansRegular"
  }
});

const scaledStyles = scaleStyleSheet(styles);
