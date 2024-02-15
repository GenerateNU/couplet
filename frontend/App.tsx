import React from "react";
import { Platform, StatusBar, View } from "react-native";
import { PaperProvider } from "react-native-paper";
import HomeScreen from "./components/HomeScreen";

export default function App() {
  return (
    <View>
      {Platform.OS === "ios" && <StatusBar barStyle="default" />}
      <PaperProvider>
        <HomeScreen />
      </PaperProvider>
    </View>
  );
}
