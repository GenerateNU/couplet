import { Link } from "expo-router";
import React from "react";
import { Image, View } from "react-native";
import { Button } from "react-native-paper";

export default function NavButton({ route, icon }: { route: string; icon: any }) {
  return (
    <View style={{ width: "100%", flex: 1, justifyContent: "center" }}>
      <Link
        href={`/${route}`}
        style={{
          width: "auto",
          justifyContent: "center"
        }}
      >
        <Button
          style={{
            width: "auto",
            height: "100%",
            justifyContent: "center",
            marginLeft: 15
          }}
        >
          <Image source={icon} style={{ height: 30, width: 30 }} resizeMode="contain" />
        </Button>
      </Link>
    </View>
  );
}
