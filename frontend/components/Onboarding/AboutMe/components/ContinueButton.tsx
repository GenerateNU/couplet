import React, { useState } from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { Link } from "expo-router";

const ContinueButton = ({ route, Component }) => {
  const [showComponent, setShowComponent] = useState(false);

  const handlePress = () => {
    if (Component) {
      setShowComponent(true);
    }
  };

  return (
    <View>
      <TouchableOpacity style={styles.button} onPress={handlePress}>
        <Link href={`/${route}`}>
          <Text style={styles.text}>Continue</Text>
        </Link>
      </TouchableOpacity>
      {showComponent && <Component />}
    </View>
  );
};

const styles = StyleSheet.create({
  button: {
    backgroundColor: "#FBA4A4",
    borderRadius: 25, 
    height: 50, 
    justifyContent: "center", 
    alignItems: "center", 
  },
  text: {
    color: "white", // choose your color
    fontSize: 16 // choose your size
  }
});

export default ContinueButton;
