import React, { useState } from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import COLORS from '../../colors';
import scaleStyleSheet from '../../scaleStyles';

interface ButtonProps {
  title: string;
  onButtonClick: () => void; 
}

function OnboardingButton({title, onButtonClick}: ButtonProps) {
  const [isPressed, setIsPressed] = useState(false);

  const handlePress = () => {
    onButtonClick(); 
    setIsPressed(!isPressed);
  };

  return (
    <View style={scaledStyles.buttonContainer}>
      <TouchableOpacity
        onPress={handlePress}
        style={[scaledStyles.button, isPressed ? scaledStyles.buttonPressed : null]}
      >
        <Text style={scaledStyles.text}>{title}</Text>
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
    paddingHorizontal: 8
  },
  button: {
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

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingButton;