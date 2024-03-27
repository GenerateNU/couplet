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

const RELIGION_IMAGE = require('../../assets/OnboardingReligion.png');

export default function LifestyleReligion() {
  const [religion, setReligion] = useState<string | null>(null);
  const listOfReligions = [
    'Christianity', 
    'Islam', 
    'Hindusim', 
    'Buddhism', 
    'Catholicism', 
    'Judaism', 
    'Agnosticisim', 
    'Atheism', 
    'Other', 
    'Prefer not to say'
  ]

  const handlePillPress = (selection: string) => {
    if (religion == selection) {
        setReligion(null);
    } else {
        setReligion(selection);
    }
  }

  const isContinueButtonEnabled = () => {
    return religion !== null;
  }

  const onContinue = () => {
    router.push('Onboarding/LifestylePolitics');
  };

  return (
    <ScrollView contentContainerStyle={scaledStyles.scrollContainer}>
        <View style={scaledStyles.container}>
            <View>
                <TopBar onBackPress={() => router.back()} text='Lifestyle' selectedCount={3} />
            </View>
            <View style={scaledStyles.contentContainer}>
                <Image source={RELIGION_IMAGE} />
                <OnboardingTitle text='I believe in...' />
                <View style={scaledStyles.buttonContainer}>
                {listOfReligions.map((option, index) => 
                    <OnboardingPillButton
                    key={index}
                    label={option}
                    onPress={() => handlePillPress(option)}
                    isSelected={religion === option}
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
    paddingTop: 100,
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

