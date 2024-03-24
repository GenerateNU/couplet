import { router } from "expo-router";
import React, { useRef, useState } from "react";
import { Image, Keyboard, StyleSheet, TextInput, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";

const BIO_IMAGE = require("../../assets/profilebio.png");

export default function ProfileBio() {
  const [bio, setBio] = useState<string>("");
  const [inputStyle, setInputStyle] = useState(styles.unfocusedInput);
  const inputRef = useRef(null);

  const onContinue = () => {
    // TODO: Save bio into form state
    router.push("Onboarding/ProfilePhotos");
  };

  return (
    <View style={styles.container}>
      <View>
        <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={5} />
      </View>
      <View>
        <Image source={BIO_IMAGE} />
        <OnboardingTitle text="What's your bio?" />
        <TextInput
          ref={inputRef}
          onFocus={() => setInputStyle(styles.focusedInput)}
          onBlur={() => setInputStyle(styles.unfocusedInput)}
          onSubmitEditing={() => Keyboard.dismiss()}
          style={inputStyle}
          onChangeText={setBio}
          value={bio}
          placeholder="Your bio here (optional)"
        />
      </View>
      <View />
      <View>
        <ContinueButton title="Continue" isDisabled={false} onPress={onContinue} />
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
  },
  focusedInput: {
    padding: 5,
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 10
  },
  unfocusedInput: {
    padding: 5,
    borderWidth: 1,
    borderColor: "#CDCDCD",
    borderRadius: 10
  }
});
