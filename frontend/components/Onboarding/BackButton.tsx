import React from 'react';
import { TouchableOpacity, View, StyleSheet } from 'react-native';
import Icon from 'react-native-vector-icons/AntDesign';
import { COLORS } from '../../colors';

type BackButtonProps = {
  onPress: () => void; 
};

export const BackButton: React.FC<BackButtonProps> = ({ onPress }) => {
  return (
    <TouchableOpacity onPress={onPress} style={styles.button}>
      <Icon
        name="leftcircle" 
        size={24} 
        color={COLORS.primary} 
      />
    </TouchableOpacity>
  );
};

const styles = StyleSheet.create({
  button: {
    width: 24,
    height: 24,
    justifyContent: 'center',
    alignItems: 'center',
    borderRadius: 100,
    backgroundColor: COLORS.white, 
    shadowOffset: {
      width: 0,
      height: 2,
    },
    shadowOpacity: 0.1,
    shadowRadius: 4,
    elevation: 5,
  },
});

export default BackButton;
