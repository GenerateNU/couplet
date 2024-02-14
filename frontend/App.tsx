import * as React from "react";
import HomeScreen from "./components/HomeScreen";
import { PaperProvider } from "react-native-paper";
import { View, Platform, StatusBar } from "react-native";

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
