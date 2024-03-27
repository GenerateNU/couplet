import { useNavigation, router } from 'expo-router';
import React, { useState } from 'react';
import {
  Image,
  StyleSheet,
  View
} from 'react-native';
import ContinueButton from '../../components/Onboarding/ContinueButton';
import OnboardingTitle from '../../components/Onboarding/OnboardingTitle';
import TopBar from '../../components/Onboarding/TopBar';
import OnboardingPillButton from '../../components/Onboarding/OnboardingPillButton';
import scaleStyleSheet from '../../scaleStyles';

const POLITICS_IMAGE = require("../../assets/OnboardingPolitics.png");

export default function LifestylePolitics() {
    const navigation = useNavigation();
    const [politics, setPolitics] = useState<string | null>(null);
  const listOfPolitics = [
    'Liberal', 
    'Moderate', 
    'Conservative', 
    'Other', 
    'Prefer not to say', 
  ]

  const handlePillPress = (selection: string) => {
    if (politics == selection) {
        setPolitics(null);
    } else {
        setPolitics(selection);
    }
  }

  const isContinueButtonEnabled = () => {
    return politics !== null;
  }

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <View style={scaledStyles.container}>
        <View>
            <TopBar onBackPress={() => goBack()} text='Lifestyle' selectedCount={3} />
        </View>
        <View>
            <Image source={POLITICS_IMAGE} />
            <OnboardingTitle text='Politically, I am...' />
            <View style={scaledStyles.buttonContainer}>
            {listOfPolitics.map((option, index) => 
                <OnboardingPillButton
                key={index}
                label={option}
                onPress={() => handlePillPress(option)}
                isSelected={politics === option}
                />
            )}
            </View>
        </View>
        <View style={scaledStyles.continueContainer}>
            <ContinueButton
            title='Continue'
            isDisabled={!isContinueButtonEnabled()}
            onPress={() => router.push("Onboarding/LifestyleHabits")}
            />
        </View>
    </View>
  );
}

const styles = StyleSheet.create({
    buttonContainer: {
        flexDirection: 'row',
        flexWrap: 'wrap',
        alignItems: 'flex-start',
        marginTop: 20,
      },
    container: {
      flex: 1,
      justifyContent: "space-between",
      alignItems: "center",
      margin: 30
    },
    ContinueButtonContainer: {
      marginBottom: 10
    }
});

const scaledStyles = scaleStyleSheet(styles);
   