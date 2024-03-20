import React from 'react';
import { StyleSheet, Text, TouchableOpacity, TouchableOpacityProps, ViewStyle } from 'react-native';
import { COLORS } from './colors';

interface CustomButtonProps extends TouchableOpacityProps {
  title: string; 
  isDisabled: boolean; 
}

const ContinueButton: React.FC<CustomButtonProps> = ({ title, isDisabled, onPress }) => {
  return (
    <TouchableOpacity
      onPress={onPress}
      disabled={isDisabled}
      style={[
        styles.button,
        isDisabled ? styles.buttonDisabled : styles.buttonEnabled, // Apply different styles based on the disabled state
      ]}
    >
      <Text style={styles.text}>{title}</Text>
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    width: 330,
    height: 41,
    paddingVertical: 10,
    paddingHorizontal: 130,
    borderRadius: 65,
    borderWidth: 1,
    borderColor: COLORS.disabled, 
    backgroundColor: COLORS.disabled, 
    shadowColor: '#000',
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 2,
    justifyContent: 'center', 
    alignItems: 'center',
  },
  buttonEnabled: {
    backgroundColor: COLORS.primary, 
    borderColor: COLORS.primary,
  },
  buttonDisabled: {
  },
  text: {
    fontFamily: 'DMSansRegular', 
    fontSize: 16,
    fontWeight: '500',
    lineHeight: 21, 
    textAlign: 'left', 
    color: COLORS.white, 
    width: 70,
    height: 21,
},

});

export default ContinueButton;
