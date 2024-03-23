import { useRouter } from "expo-router";
import React from "react";
import { Image, Pressable, Text, View } from "react-native";

const PIN = require("../../assets/pin.png");
const COIN = require("../../assets/coin.png");

type HomeEventCardProps = {
  id: string;
  name: string;
  // TODO: we need Location and Cost, but these aren't in the endpoint response yet
};

export default function HomeEventCard({ id, name }: HomeEventCardProps) {
  const router = useRouter();

  return (
    <Pressable
      onPress={() => {
        router.push("Event");
        router.setParams({ collectionId: "", eventId: id });
      }}
    >
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
            {name}
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
              {/* TODO - replace with actual location */}
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
              $$
              {/* TODO - replace with actual cost */}
            </Text>
            <View />
          </View>
        </View>
      </View>
    </Pressable>
  );
}
