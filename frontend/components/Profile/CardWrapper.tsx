import React from "react";
import { View } from "react-native";
export default function CardWrapper({ children }: { children: React.ReactNode }) {
  return (
    <View
      style={{
        padding: 20,
        borderRadius: 20,
        width: "90%",
        alignSelf: "center",
        top: 0,
        shadowColor: "#000000",
        shadowRadius: 4,
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        backgroundColor: "#ffffff"
      }}
    >
      {children}
    </View>
  );
}
