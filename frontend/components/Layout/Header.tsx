import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Text, View } from "react-native";

export default function Header() {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <View
      style={{
        backgroundColor: "white",
        flex: 1,
        flexDirection: "row",
        justifyContent: "center",
        padding: 20
      }}
    >
      <Text
        style={{
          color: "black",
          justifyContent: "center",
          textAlign: "left",
          width: "100%",
          fontSize: 20,
          fontFamily: "DMSansRegular"
        }}
      >
        Couplet
      </Text>
    </View>
  );
}
