import React, { useState } from 'react';
import { StyleSheet, Text, TouchableOpacity, View } from 'react-native';
import { COLORS } from './colors';

// Defining props for the component
interface MyButtonProps {
  title: string; // Prop for the button title
  onButtonClick: () => void; // Prop for the button click action
}

const MyButton: React.FC<MyButtonProps> = ({ title, onButtonClick }) => {
  const [isPressed, setIsPressed] = useState(false);

  // Function to handle press action
  const handlePress = () => {
    onButtonClick(); // Call the passed onButtonClick function
    setIsPressed(!isPressed); // Toggle the isPressed state
  };

  return (
    <View style={styles.buttonContainer}>
      <TouchableOpacity
        onPress={handlePress}
        style={[styles.button, isPressed ? styles.buttonPressed : null]}
      >
        <Text>{title}</Text>
      </TouchableOpacity>
    </View>
  );
};

// Stylesheet for the button
const styles = StyleSheet.create({
  buttonContainer: {
    position: 'absolute',
    top: 402,
    left: 236,
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
    backgroundColor: COLORS.secondary, // Button background color when pressed
  },
});

export default MyButton;
