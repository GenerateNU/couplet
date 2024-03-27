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
import OnboardingSmallTitle from '../../components/Onboarding/OnboardingSmallTitle';

const PASSIONS_IMAGE = require("../../assets/OnboardingPassions.png");

export default function LifestylePassions() {
    const [interests, setInterests] = useState<string[]>([]);
    const listOfPassions = [
      'Acting', 
      'Baking', 
      'Board Games', 
      'Cars', 
      'Calligraphy', 
      'Cooking',
      'Concerts',
      'Cycling',
      'Dancing', 
      'DIY',
      'Fishing',
      'Hiking',
      'Interior Design',
      'Gaming',
      'Gardening',
      'Karaoke',
      'K-Pop',
      'Knitting',
      'Music',
      'Painting',
      'Parkour',
      'Photography',
      'Pilates',
      'Poetry',
      'Puzzles',
      'Running',
      'Rock Climbing',
      'Reading',
      'Swimming',
      'Surfing',
      'Sewing',
      'Singing',
      'Sports',
      'Traveling',
      'Trivia',
      'Video Games',
      'Volunteering',
      'Writing',
      'Weight Lifting',
      'Yoga',
    ]
  
    const handleInterestsButton = (selection: string) => {
      if (interests.includes(selection)) {
        const updatedInterests = interests.filter(interest => interest !== selection);
        setInterests(updatedInterests);
      } else {
        if (interests.length < 5) {
            setInterests([...interests, selection]);
        }
      }
    }
  
    const isContinueButtonEnabled = () => {
      return interests.length == 5;
    }
  
    const onContinue = () => {
      router.push('Onboarding/ProfileBio');
    };
  
    return (
      <ScrollView contentContainerStyle={styles.scrollContainer}>
          <View style={styles.container}>
              <View>
                  <TopBar onBackPress={() => router.back()} text='Lifestyle' selectedCount={3} />
              </View>
              <View style={styles.contentContainer}>
                  <Image source={PASSIONS_IMAGE} style={styles.imageContainer}/>
                  <OnboardingTitle text='What are you passionate about?' />
                  <OnboardingSmallTitle text='Select your top five interests!'/>
                  <View style={styles.buttonContainer}>
                  {listOfPassions.map((option, index) => 
                      <OnboardingPillButton
                      key={index}
                      label={option}
                      onPress={() => handleInterestsButton(option)}
                      isSelected={interests.includes(option)}
                      />
                  )}
                  </View>
              </View>
              <View style={styles.continueContainer}>
                  <ContinueButton
                  title={(interests.length === 0 || interests.length == 5) ? 'Continue' : 'Continue ' + interests.length + '/5'}
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
    imageContainer: {
      width: 200,
      height: 200,
      flexShrink: 0,
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