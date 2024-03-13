import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import { useRouter } from "expo-router";
import React from "react";
import { Image, Pressable, Text, View } from "react-native";

const PIN = require("../../assets/pin.png");
const COIN = require("../../assets/coin.png");

export default function HomeEventCard() {
  const router = useRouter();
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <Pressable onPress={() => router.push("DummyEventDetails")}>
      <View
        style={{
          borderStyle: "solid",
          borderWidth: 1,
          borderColor: "black",
          borderRadius: 10,
          marginRight: 10
        }}
      >
        {/* <Image source={{uri: 'https://reactnative.dev/img/tiny_logo.png'}} style={{width: 150, height: 150}} /> */}
        <View
          style={{
            width: "100%",
            height: 150,
            backgroundColor: "rgb(200,200,200)",
            borderTopLeftRadius: 10,
            borderTopRightRadius: 10
          }}
        />
        <View>
          <Text
            style={{ textAlign: "center", padding: 10, fontSize: 14, fontFamily: "DMSansRegular" }}
          >
            Winter Ice Skating
          </Text>
          <View style={{ flexDirection: "row", padding: 10, borderRadius: 20, paddingTop: 0 }}>
            <Image source={PIN} style={{ width: 20, height: 20 }} />
            <Text
              style={{
                textAlign: "center",
                justifyContent: "center",
                verticalAlign: "middle",
                marginTop: 2,
                fontFamily: "DMSansRegular"
              }}
            >
              Frog Pond
            </Text>
            <Image source={COIN} style={{ width: 20, height: 20, marginLeft: 10 }} />
            <Text
              style={{
                textAlign: "center",
                justifyContent: "center",
                verticalAlign: "middle",
                marginTop: 2,
                marginHorizontal: 2,
                fontFamily: "DMSansRegular"
              }}
            >
              Cost
            </Text>
            <View />
          </View>
        </View>
      </View>
    </Pressable>
  );
}
