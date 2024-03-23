import React, { useEffect, useState } from 'react';
import { View, Image } from 'react-native';
import { router } from 'expo-router'
import PhotoPicker from "../../components/PhotoPicker";
import ContinueButton from '../../components/Onboarding/ContinueButton';
import OnboardingTitle from '../../components/Onboarding/OnboardingTitle';
import TopBar from '../../components/Onboarding/TopBar';

const CAMERA_IMAGE = require('../../assets/profilecamera.png')

export default function ProfilePhotos() {
  const [imagesSelected, setImagesSelected] = useState(false);
  const [images, setImages] = useState<string[]>([]);

  useEffect(() => {
    if (images.length === 4) {
      setImagesSelected(true)
    } else {
      setImagesSelected(false)
    }
  }, [images])

  return (
    <View style={{ flex: 1, justifyContent: "space-between", margin: 30 }}>
      <View style={{alignSelf: "center"}}>
        <TopBar onBackPress={() => {}} text="Profile" selectedCount={5}/>
      </View>
      <View>
        <Image source={CAMERA_IMAGE} height={50}/>
        <OnboardingTitle text="Show your best angles"/>
        <PhotoPicker onPick={setImages}/>
      </View>
      <View>
        <ContinueButton title="Continue" isDisabled={!imagesSelected} onPress={() => router.push("Onboarding/ProfileInsta")}/>
      </View>
    </View>
  )
}