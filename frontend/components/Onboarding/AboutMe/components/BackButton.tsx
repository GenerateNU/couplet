import React from 'react';
import { Image, TouchableOpacity, View, StyleSheet } from 'react-native';
import { useNavigation } from '@react-navigation/native';

const BackButton = () => {
  const navigation = useNavigation();
  return (
    <TouchableOpacity onPress={() => {
        if (navigation.canGoBack()) {
            navigation.goBack()
        } else {
            //navigation.navigate('index');
            console.log("Back")
        }
    }}>
      <Image source={require("../../../../assets/Arrows.png")} />
    </TouchableOpacity>
  );
};

export default BackButton;
