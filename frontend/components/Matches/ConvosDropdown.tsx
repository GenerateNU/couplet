import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React, { useState } from "react";
import { Image, Text, View } from "react-native";

export type ConvoUser = {
  userID: number;
  name: string;
  chatLine: string;
};

type ConvosDropdownProps = {
  convos: ConvoUser[];
  convoType: string;
};

const PROFILE = require("../../assets/profile.png");

function ConvosDropdown({ convos, convoType }: ConvosDropdownProps) {
  const [open, setOpen] = useState(false);

  const toggleItems = () => {
    setOpen(!open);
  };

  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <View style={{ marginTop: 10 }}>
      <View
        onTouchStart={toggleItems}
        style={{
          flexDirection: "row",
          justifyContent: "space-between",
          alignItems: "center",
          paddingHorizontal: "5%",
          marginBottom: 10
        }}
      >
        <Text style={{ fontSize: 14, fontFamily: "DMSansRegular" }}>
          {convoType === "activeConvos"
            ? `Active Convos (${convos.length})`
            : `Archives (${convos.length})`}
        </Text>
        <Text style={{ fontSize: 14 }}>{open ? "\u25B2" : "\u25BC"}</Text>
      </View>
      {open && (
        <View style={{ flexDirection: "column", marginBottom: 10 }}>
          {convos.map((convo, index) => (
            <View
              key={convo.userID}
              style={{
                flexDirection: "row",
                borderBottomWidth: 1,
                borderTopWidth: index === 0 ? 1 : 0,
                borderColor: "lightgray",
                padding: 15,
                alignItems: "center"
              }}
            >
              <Image source={PROFILE} style={{ width: 55, height: 55, marginRight: 10 }} />
              <View style={{ flexDirection: "column", justifyContent: "center" }}>
                <Text style={{ fontSize: 22, marginBottom: 5, fontFamily: "DMSansRegular" }}>
                  {convo.name}
                </Text>
                <Text style={{ fontSize: 14, marginBottom: 5, fontFamily: "DMSansRegular" }}>
                  {convo.chatLine}
                </Text>
              </View>
            </View>
          ))}
        </View>
      )}
    </View>
  );
}

export default ConvosDropdown;
