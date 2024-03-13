import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Text } from "react-native";
import { Button } from "react-native-paper";
import { ReactionButtonProps } from "./ReactionButtonProps";

/* eslint-disable react/no-children-prop */

function ReactionButton({ like, icon, label, handleReact }: ReactionButtonProps) {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <Button
      style={{
        paddingVertical: 10,
        marginHorizontal: 10,
        borderRadius: 30,
        backgroundColor: like ? "black" : "white",
        flex: 1,
        alignSelf: "center" // This will center the button horizontally
      }}
      labelStyle={{ fontSize: 20, alignSelf: "flex-start", color: like ? "white" : "black" }}
      icon={icon}
      mode="elevated"
      onPress={() => handleReact(like)}
    >
      <Text style={{ fontFamily: "DMSansRegular" }}>{label}</Text>
    </Button>
  );
}

export default ReactionButton;
