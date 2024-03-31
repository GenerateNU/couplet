import { Link } from "expo-router";
import React from "react";
import { Image, ImageSourcePropType, View } from "react-native";

export default function NavButton({ route, icon }: { route: string; icon: ImageSourcePropType }) {
  return (
    <View style={{ alignItems: "center", marginVertical: 15 }}>
      <Link
        href={`/${route}`}
      >
        <Image source={icon} style={{ height: 30, width: 30 }} resizeMode="contain" />
      </Link>
    </View>
  );
}
