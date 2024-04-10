import { router } from "expo-router";
import React, { useEffect, useMemo, useState } from "react";
import { Image, ScrollView, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import COLORS from "../../../colors";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setPhotos } from "../../../state/formSlice";
import { useAppDispatch, useAppSelector } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const CAPTIONS = require("../../../assets/profilecaptions.png");

function ProfileCaptions() {
  const dispatch = useAppDispatch();
  const photos = useAppSelector((state) => state.form.photos);
  const images: string[] = useMemo(() => photos.map((photo) => photo.filePath), [photos]);
  const [captions, setCaptions] = useState(["", "", ""]);
  const [continueDisabled, setContinueDisabled] = useState<boolean>(true);

  useEffect(() => {
    const captionCount = captions.filter((cap) => cap !== "").length;
    setContinueDisabled(captionCount !== 3);
  }, [captions]);

  const onContinue = () => {
    // Create a copy of the photos state
    const photosState = [...photos];
    // eslint-disable-next-line no-plusplus
    for (let i = 0; i < photosState.length; i++) {
      if (i !== 0) {
        photosState[i] = { ...photosState[i], caption: captions[i] };
      }
    }
    dispatch(setPhotos(photosState));
    router.push("Onboarding/Profile/ProfileInsta");
  };

  const onSubmitCaption = (text: string, index: number) => {
    const newCaptions = [...captions];
    newCaptions[index] = text;
    setCaptions(newCaptions);
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
      <ScrollView contentContainerStyle={scaledStyles.mainContainer}>
        <View>
          <Image source={CAPTIONS} />
          <View style={scaledStyles.headingContainer}>
            <OnboardingTitle text="Captions speak louder than words." />
          </View>
          <View style={scaledStyles.headingContainer}>
            <Text style={scaledStyles.textHelper}> Add one to your photos! </Text>
          </View>
          <View style={scaledStyles.inputWrapper}>
            {images.map((img, i) => (
              // eslint-disable-next-line react/no-array-index-key
              <View style={scaledStyles.imageContainer} key={i}>
                <Image source={{ uri: img }} style={scaledStyles.imageStyle} />
                {i === 0 ? (
                  <View style={scaledStyles.pillWrapper}>
                    <View style={scaledStyles.pill}>
                      <Text style={scaledStyles.pillText}>Your top photo</Text>
                    </View>
                  </View>
                ) : (
                  <View style={scaledStyles.textInputWrapper}>
                    <TextInput
                      style={scaledStyles.textInput}
                      onChange={(e) => onSubmitCaption(e.nativeEvent.text, i - 1)}
                      editable
                      placeholder="Caption"
                    />
                  </View>
                )}
              </View>
            ))}
          </View>
        </View>
        <View>
          <ContinueButton title="Continue" isDisabled={continueDisabled} onPress={onContinue} />
        </View>
      </ScrollView>
    </SafeAreaView>
  );
}

export default ProfileCaptions;

const overrideStyles = {
  TopUiContainer: {
    alignItems: "center"
  },
  mainContainer: {
    paddingTop: "20%",
    marginLeft: 20,
    marginRight: 20
  },
  headingContainer: {
    marginTop: 8
  },
  imageStyle: {
    marginBottom: 8,
    height: 400,
    width: "100%",
    borderRadius: 5
  },
  pillWrapper: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center",
    marginBottom: "10%"
  },
  pill: {
    paddingHorizontal: 30,
    paddingVertical: 10,
    borderRadius: 25,
    backgroundColor: COLORS.primary,
    position: "absolute"
  },
  pillText: {
    fontFamily: "DMSansMedium",
    color: COLORS.white,
    fontWeight: 400
  }
};

const scaledStyles = scaleStyleSheet({ ...onboardingStyles, ...overrideStyles });
