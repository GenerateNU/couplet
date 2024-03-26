import { router } from "expo-router";
import React, { useRef, useState } from "react";
import { Image, Keyboard, StyleSheet, TextInput, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";

const BIO_IMAGE = require("../../assets/profilebio.png");

export default function ProfileBio() {
  const [open, setOpen] = useState<boolean>(false);
  const [prompt, setPrompt] = useState(null);
  const [options, setOptions] = useState([
    { label: "An ideal date", value: 0 },
    { label: "My perfect day consists of", value: 1 },
    { label: "On weekends you can find me", value: 2 },
    { label: "What I'm looking for on this app", value: 3 },
    { label: "One thing you should know about me", value: 4 }
  ]);
  const [response, setResponse] = useState<string>("");
  const [inputStyle, setInputStyle] = useState(styles.unfocusedResponse);
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
        <DropDownPicker
          style={prompt !== null ? styles.chosenPrompt : styles.unchosenPrompt}
          placeholder="Select a prompt"
          open={open}
          value={prompt}
          items={options}
          setOpen={setOpen}
          setValue={setPrompt}
          setItems={setOptions}
        />
        <TextInput
          ref={inputRef}
          onFocus={() => setInputStyle(styles.focusedResponse)}
          onBlur={() => setInputStyle(styles.unfocusedResponse)}
          onSubmitEditing={() => Keyboard.dismiss()}
          style={inputStyle}
          onChangeText={setResponse}
          value={response}
          placeholder="Your response here"
        />
      </View>
      <View />
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!(prompt !== null && response !== "")}
          onPress={onContinue}
        />
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
  chosenPrompt: {
    marginBottom: 10
  },
  unchosenPrompt: {
    marginBottom: 10
  },
  focusedResponse: {
    padding: 10,
    borderWidth: 1,
    borderColor: "black",
    borderRadius: 10
  },
  unfocusedResponse: {
    padding: 10,
    borderWidth: 1,
    borderColor: "#CDCDCD",
    borderRadius: 10
  }
});
