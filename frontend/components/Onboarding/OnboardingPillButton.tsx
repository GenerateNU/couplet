import React from 'react';
import { TouchableOpacity, Text, StyleSheet } from 'react-native';
import COLORS from '../../colors';
import scaleStyleSheet from '../../scaleStyles';

interface OnboardPillButtonProps {
  label: string;
  onPress: (label: string | null) => void;
  isSelected: boolean;
}

const OnboardingPillButton: React.FC<OnboardPillButtonProps> = ({ label, onPress, isSelected }) => {

  const handlePress = () => {
    if (isSelected) {
      onPress(null);
    } else {
      onPress(label);
    }
  };

  return (
    <TouchableOpacity
      style={[scaledStyles.pillButton, isSelected && scaledStyles.selected]}
      onPress={handlePress}
    >
      <Text style={scaledStyles.buttonText}>{label}</Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  pillButton: {
    borderRadius: 100, 
    borderWidth: 1.5,
    borderColor: COLORS.secondary, 
    backgroundColor: COLORS.white,
    paddingHorizontal: 20,
    paddingVertical: 10,
    margin: 5,
  },
  selected: {
    backgroundColor: COLORS.secondary,
  },
  buttonText: {
    fontSize: 15,
  },
});

export default OnboardingPillButton;

const scaledStyles = scaleStyleSheet(styles);