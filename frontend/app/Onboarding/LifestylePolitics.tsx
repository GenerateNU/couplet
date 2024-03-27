import { router } from 'expo-router';
import React, { useState } from 'react';
import {
  Image,
  StyleSheet,
  ScrollView,
  View
} from 'react-native';
import ContinueButton from '../../components/Onboarding/ContinueButton';
import OnboardingTitle from '../../components/Onboarding/OnboardingTitle';
import TopBar from '../../components/Onboarding/TopBar';
import OnboardingPillButton from '../../components/Onboarding/OnboardingPillButton';
import scaleStyleSheet from '../../scaleStyles';

const POLITICS_IMAGE = require("../../assets/OnboardingPolitics.png");

export default function LifestylePolitics() {
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

  const onContinue = () => {
    router.push('Onboarding/LifestyleHabits');
  };

  return (
    <ScrollView contentContainerStyle={scaledStyles.scrollContainer}>
        <View style={scaledStyles.container}>
            <View>
                <TopBar onBackPress={() => router.back()} text='Lifestyle' selectedCount={3} />
            </View>
            <View style={scaledStyles.contentContainer}>
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
                onPress={onContinue}
                />
            </View>
        </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  scrollContainer: {
    flexGrow: 1,
  },
  container: {
    flex: 1,
    justifyContent: 'space-between',
  },
  contentContainer: {
    paddingTop: 69,
    paddingLeft: 10,
    paddingRight: 10,
  },
  buttonContainer: {
    flexDirection: 'row',
    flexWrap: 'wrap',
    alignItems: 'flex-start',
    marginTop: 20,
  },
  continueContainer: {
    display: 'flex',
    width: '100%',
    height: 41,
    paddingTop: 10,
    paddingRight: 130,
    marginBottom: 30,
    paddingLeft: 130,
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
    flexShrink: 0,
  },
});

const scaledStyles = scaleStyleSheet(styles);