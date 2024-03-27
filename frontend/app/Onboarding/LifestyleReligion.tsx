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

const RELIGION_IMAGE = require('../../assets/OnboardingReligion.png');

export default function LifestyleReligion() {
    const navigation = useNavigation();
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

  function goBack() {
    if (navigation.canGoBack()) {
      navigation.goBack();
    }
  }

  return (
    <View style={scaledStyles.container}>
        <View>
            <TopBar onBackPress={() => goBack()} text="Lifestyle" selectedCount={3} />
        </View>
        <View>
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
        <View style={scaledStyles.ContinueButtonContainer}>
        <ContinueButton
          onPress={() => router.push("Onboarding/LifestylePolitics")}
          title={"Continue"}
          isDisabled={!isContinueButtonEnabled()}
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
        marginTop: 10,
    }
});

const scaledStyles = scaleStyleSheet(styles);

