import React, { useState } from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { COLORS } from '../../colors';

interface ButtonProps {
  title: string;
  onButtonClick: () => void; 
}

const OnboardingButton: React.FC<ButtonProps> = ({ title, onButtonClick }) => {
  const [isPressed, setIsPressed] = useState(false);

  const handlePress = () => {
    onButtonClick(); 
    setIsPressed(!isPressed);
  };

  return (
    <View style={styles.buttonContainer}>
      <TouchableOpacity
        onPress={handlePress}
        style={[styles.button, isPressed ? styles.buttonPressed : null]}
      >
        <Text style={styles.text}>{title}</Text>
      </TouchableOpacity>
    </View>
  );
};

const styles = StyleSheet.create({
  buttonContainer: {
    position: 'absolute',
    top: 402,
    left: 236,
  },
  text: {
    fontFamily: 'DMSansMedium', 
    fontSize: 12,
    fontWeight: 'bold',
    lineHeight: 15.62,
    textAlign: 'center',
  },
  button: {
    width: 73,
    height: 24,
    alignItems: 'center',
    justifyContent: 'center',
    paddingVertical: 4,
    paddingHorizontal: 12,
    borderRadius: 100,
    borderWidth: 1,
    borderColor: COLORS.secondary,
    backgroundColor: COLORS.white,
  },
  buttonPressed: {
    backgroundColor: COLORS.secondary,
  },
});

export default OnboardingButton;
