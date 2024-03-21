import React from 'react';
import { View, Text, StyleSheet } from 'react-native';
import BackButton from './BackButton'; 
import OnboardingBar from './OnboardingBar'; 
import COLORS from '../../colors'; 

type TopBarProps = {
  onBackPress: () => void;
  text: string;
  selectedCount: number;
};

function TopBar(props: TopBarProps) {
  return (
    <View style={styles.container}>
      <BackButton onPress={props.onBackPress} />
      <View style={styles.textBarContainer}>
        <Text style={styles.informationText}>{props.text}</Text>
        <OnboardingBar selectedCount={props.selectedCount} />
      </View>
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    width: 346,
    height: 69,
    top: 68,
    left: 24,
    position: 'absolute',
    flexDirection: 'column', 
    justifyContent: 'space-between',
  },
  textBarContainer: {
    paddingTop: 8,
    width: 346, 
    height: 21, 
    justifyContent: 'flex-end',
  },
  informationText: {
    width: 46,
    height: 10,
    fontFamily: "DMSansRegular",
    fontSize: 10,
    fontWeight: '400',
    lineHeight: 10,
    textAlign: 'left',
    color: COLORS.darkGray,
    marginBottom: 2,
  },
});

export default TopBar;
