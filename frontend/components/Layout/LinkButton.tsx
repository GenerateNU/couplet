import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import { Link } from "expo-router";
import React from "react";
import { Pressable, Text } from "react-native";

export default function LinkButton({ text }: { text: string }) {
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
        backgroundColor: "black",
        borderWidth: 1,
        padding: "5%",
        borderRadius: 100
      }}
    >
      <Link href="/People">
        <Text style={{ color: "white", fontFamily: "DMSansRegular" }}>{text}</Text>
      </Link>
    </Pressable>
  );
}
