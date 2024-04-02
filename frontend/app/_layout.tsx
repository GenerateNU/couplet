import React from "react";
import { Provider } from "react-redux";
import { Slot } from "expo-router";
import { useFonts } from "expo-font";
import {
  DMSans_700Bold as DMSansBold,
  DMSans_500Medium as DMSansMedium,
  DMSans_400Regular as DMSansRegular
} from "@expo-google-fonts/dm-sans";
import store from "../state/store";

export default function Layout() {
  const [fontsLoaded] = useFonts({
    DMSansRegular,
    DMSansMedium,
    DMSansBold
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <Provider store={store}>
      <Slot />
    </Provider>
  );
}
