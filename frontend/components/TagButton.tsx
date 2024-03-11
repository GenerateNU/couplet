import React from "react";
import { Pressable, Text } from "react-native";
import {
  DMSans_400Regular as DMSansRegular
} from '@expo-google-fonts/dm-sans';
import { useFonts } from "expo-font";

export default function TagButton({
  text,
  selected,
  onPress
}: {
  text: string;
  selected: boolean;
  onPress: () => void;
}) {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null; 
  }
  
  return (
    <Pressable
      style={{
        borderStyle: "solid",
        borderColor: "black",
        backgroundColor: selected ? "black" : "white",
        borderWidth: 1,
        marginHorizontal: 10,
        padding: 15,
        borderRadius: 50
      }}
      onPress={onPress}
    >
      <Text style={{ color: selected ? "white" : "black", fontFamily: "DMSansRegular"}}>{text}</Text>
    </Pressable>
  );
}
