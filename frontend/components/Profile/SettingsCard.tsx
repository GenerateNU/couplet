import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Image, Text, TouchableOpacity, View } from "react-native";

export default function SettingsCard({
  text,
  img,
  last
}: {
  text: string;
  img: any;
  last?: boolean;
}) {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }
  return (
    <TouchableOpacity>
      <View
        style={{
          padding: 20,
          width: "100%",
          alignSelf: "center",
          backgroundColor: "#ffffff",
          borderBottomWidth: last ? 0 : 1,
          borderBottomColor: "#CDCDCD",
          flexDirection: "row"
        }}
      >
        <Image
          source={img}
          style={{
            width: 50,
            height: 50,
            alignSelf: "center"
          }}
        />
        <Text
          style={{
            verticalAlign: "middle",
            margin: "auto",
            alignSelf: "center",
            padding: 10,
            fontFamily: "DMSansRegular",
            fontWeight: 800,
            fontSize: 16
          }}
        >
          {text}
        </Text>
        <Image
          source={require("../../assets/Vector.png")}
          style={{
            width: 8,
            height: 15,
            alignSelf: "center",
            right: 20,
            position: "absolute"
          }}
        />
      </View>
    </TouchableOpacity>
  );
}
