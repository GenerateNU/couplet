import { router } from "expo-router";
import React, { useEffect, useState } from "react";
import { Image, SafeAreaView, StyleSheet, TextInput, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

export default function ProfileCaptions() {
  const images: string[] = ["", "", "", ""];
  const [captions, setCaptions] = useState(["", "", "", ""]);
  const [continueEnabled, setContinueEnabled] = useState<boolean>(true);

  useEffect(() => {
    const captionCount = captions.filter((cap) => cap !== "").length;
    setContinueEnabled(captionCount === 4);
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
      <View style={scaledStyles.container}>
        <View>
          <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={5} />
        </View>
        <View>
          {/* <Image source={CAPTIONS}/> */}
          <OnboardingTitle text="Add a caption" />
          {images.map((img, i) => (
            <View>
              <Image source={{ uri: img }} />
              <TextInput onSubmitEditing={(e) => onSubmitCaption(e.nativeEvent.text, i)} />
            </View>
          ))}
        </View>
        <View>
          <ContinueButton title="Continue" isDisabled={continueEnabled} onPress={onContinue} />
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
