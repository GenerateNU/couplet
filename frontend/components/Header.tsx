import React from "react";
import { Text, View } from "react-native";

export default function Header() {
  return (
    <View
      style={{
        backgroundColor: "black",
        height: "100%",
        width: "100%",
        flexDirection: "row",
        position: "absolute",
        top: 0
      }}
    >
        <Text style={{color:"white", textAlign:"center", marginTop: "10%", width:"100%", fontSize:20 }}>Couplet</Text>
    </View>
  );
}
