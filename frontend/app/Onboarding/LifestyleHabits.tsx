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
import COLORS from '../../colors';
import scaleStyleSheet from '../../scaleStyles';
import OnboardingSmallTitle from '../../components/Onboarding/OnboardingSmallTitle';

const HABITS_IMAGE = require("../../assets/OnboardingHabits.png");

export default function LifestyleHabits() {
  const [drink, setDrink] = useState<string | null>(null);
  const [smoke, setSmoke] = useState<string | null>(null);
  const [weed, setWeed] = useState<string | null>(null);
  const [drugs, setDrugs] = useState<string | null>(null);
  const listOfChoices = [
    'Yes', 
    'No', 
    'Sometimes', 
    'Socially', 
    'Prefer not to say', 
  ]
  const drugList = [
    'Yes',
    'No',
    'Prefer not to say', 
  ]

  const handleDrinkChoice = (selection: string) => {
    if (drink == selection) {
      setDrink(null);
    } else {
      setDrink(selection);
    }
  }

  const handleSmokeChoice = (selection: string) => {
    if (smoke == selection) {
      setSmoke(null);
    } else {
      setSmoke(selection);
    }
  }

  const handleWeedChoice = (selection: string) => {
    if (weed == selection) {
      setWeed(null);
    } else {
      setWeed(selection);
    }
  }

  const handleDrugsChoice = (selection: string) => {
    if (drugs == selection) {
      setDrugs(null);
    } else {
      setDrugs(selection);
    }
  }

  const isContinueButtonEnabled = () => {
    return drink !== null && smoke !== null && weed !== null && drugs !== null;
  }

  const onContinue = () => {
    router.push('Onboarding/LifestylePassions');
  };

  return (
    <ScrollView contentContainerStyle={scaledStyles.scrollContainer}>
        <View style={scaledStyles.container}>
            <View>
                <TopBar onBackPress={() => router.back()} text='Lifestyle' selectedCount={3} />
            </View>
            <View style={scaledStyles.contentContainer}>
                <Image source={HABITS_IMAGE} />
                <OnboardingTitle text='What are your habits?' />
                <OnboardingSmallTitle text='Do you drink?' />
                <View style={scaledStyles.buttonContainer}>
                {listOfChoices.map((option, index) => 
                    <OnboardingPillButton
                    key={index}
                    label={option}
                    onPress={() => handleDrinkChoice(option)}
                    isSelected={drink === option}
                    />
                )}
                </View>
                <View style={scaledStyles.separator} />
                <OnboardingSmallTitle text='Do you smoke?' />
                <View style={scaledStyles.buttonContainer}>
                {listOfChoices.map((option, index) => 
                    <OnboardingPillButton
                    key={index}
                    label={option}
                    onPress={() => handleSmokeChoice(option)}
                    isSelected={smoke === option}
                    />
                )}
                </View>
                <View style={scaledStyles.separator} />
                <OnboardingSmallTitle text='Do you smoke weed?' />
                <View style={scaledStyles.buttonContainer}>
                {listOfChoices.map((option, index) => 
                    <OnboardingPillButton
                    key={index}
                    label={option}
                    onPress={() => handleWeedChoice(option)}
                    isSelected={weed === option}
                    />
                )}
                </View>
                <View style={scaledStyles.separator} />
                <OnboardingSmallTitle text='Do you do drugs?' />
                <View style={scaledStyles.buttonContainer}>
                {drugList.map((option, index) => 
                    <OnboardingPillButton
                    key={index}
                    label={option}
                    onPress={() => handleDrugsChoice(option)}
                    isSelected={drugs === option}
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
    paddingTop: 30,
    paddingRight: 130,
    marginBottom: 30,
    paddingLeft: 130,
    justifyContent: 'center',
    alignItems: 'center',
    gap: 10,
  },
  separator: {
    borderBottomColor: COLORS.lightGray, 
    borderBottomWidth: 1, 
    marginVertical: 30,
  },
});

const scaledStyles = scaleStyleSheet(styles);