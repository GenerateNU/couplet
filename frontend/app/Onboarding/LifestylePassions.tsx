import { router, useNavigation } from 'expo-router';
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
    const navigation = useNavigation();
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

    function goBack() {
        if (navigation.canGoBack()) {
          navigation.goBack();
        }
      }
  
    return (
      <ScrollView contentContainerStyle={scaledStyles.container}>
          <View>
            <TopBar onBackPress={() => goBack()} text='Lifestyle' selectedCount={3} />
        </View>
              <View>
                  <Image source={PASSIONS_IMAGE}/>
                  <OnboardingTitle text='What are you passionate about?' />
                  <OnboardingSmallTitle text='Select your top five interests!'/>
                  <View style={scaledStyles.buttonContainer}>
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
              <View style={scaledStyles.ContinueButtonContainer}>
                  <ContinueButton
                  title={(interests.length === 0 || interests.length == 5) ? 'Continue' : 'Continue ' + interests.length + '/5'}
                  isDisabled={!isContinueButtonEnabled()}
                  onPress={() => router.push("Onboarding/ProfileBio")}
                  />
              </View>
      </ScrollView>
    );
  }
  
  const styles = StyleSheet.create({
    container: {
        flexgrow: 1,
        justifyContent: "space-between",
        alignItems: "center",
        margin: 30
    },
    buttonContainer: {
      flexDirection: 'row',
      flexWrap: 'wrap',
      alignItems: 'flex-start',
      marginTop: 20,
    },
    ContinueButtonContainer: {
        marginTop: 10,
        marginBottom: 50,
      },
  });
  
  const scaledStyles = scaleStyleSheet(styles);