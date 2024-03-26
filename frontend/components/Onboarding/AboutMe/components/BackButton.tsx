import { useNavigation } from "@react-navigation/native";
import React from "react";
import { Image, TouchableOpacity } from "react-native";

function BackButton() {
  const navigation = useNavigation();
  return (
    <TouchableOpacity
      onPress={() => {
        if (navigation.canGoBack()) {
          navigation.goBack();
        } else {
        }
      }}
    >
      <Image source={require("../../../../assets/Arrows.png")} />
    </TouchableOpacity>
  );
}

export default BackButton;
