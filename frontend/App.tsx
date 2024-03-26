import {
  DMSans_700Bold as DMSansBold,
  DMSans_500Medium as DMSansMedium,
  DMSans_400Regular as DMSansRegular
} from "@expo-google-fonts/dm-sans";
import * as AppleAuthentication from "expo-apple-authentication";
import { useFonts } from "expo-font";
import React, { useState } from "react";
import { Text, TouchableOpacity, View } from "react-native";
import Home from "./app/Home";
import AboutName from "./components/Onboarding/AboutMe/AboutName";
import AboutBirthday from "./components/Onboarding/AboutMe/AboutBirthday";
import AboutGender from "./components/Onboarding/AboutMe/AboutGender";

export default function App() {
  const [isGoogleLoggedIn, setIsGoogleLoggedIn] = useState(false);
  const [isAppleLoggedIn, setIsAppleLoggedIn] = useState(false);
  const isSignedIn = isGoogleLoggedIn || isAppleLoggedIn;
  const [fontsLoaded] = useFonts({
    DMSansRegular,
    DMSansMedium,
    DMSansBold
  });

  async function handleAppleSignIn() {
    try {
      await AppleAuthentication.signInAsync({
        requestedScopes: [
          AppleAuthentication.AppleAuthenticationScope.FULL_NAME,
          AppleAuthentication.AppleAuthenticationScope.EMAIL
        ]
      });
      setIsAppleLoggedIn(true);
    } catch (e) {
      setIsAppleLoggedIn(false);
    }
  }

  if (!fontsLoaded) {
    return null;
  }

  return (
    <View style={{ flex: 1 }}>
      {isSignedIn ? (
        <AboutName/>
      ) : (
        <View
          style={{
            paddingTop: "100%",
            width: "100%",
            borderRadius: 12,
            alignSelf: "center",
            marginBottom: 10,
            minHeight: "100%",
            paddingBottom: "30%",
            alignItems: "flex-start"
          }}
        >
          <Text
            style={{
              marginTop: -140,
              fontSize: 44,
              fontWeight: "bold",
              marginBottom: 15,
              paddingLeft: 30,
              fontFamily: "DMSansBold"
            }}
          >
            Welcome to Couplet!
          </Text>
          <Text
            style={{
              fontSize: 18,
              paddingHorizontal: 30,
              marginBottom: 20,
              fontFamily: "DMSansRegular"
            }}
          >
            Sign in with one of the providers below to get started.
          </Text>
          <View style={{ alignItems: "center", width: "100%", paddingHorizontal: 30 }}>
            <TouchableOpacity
              onPress={() => setIsGoogleLoggedIn(true)}
              style={{
                borderWidth: 1,
                borderColor: "black",
                backgroundColor: "white",
                padding: 10,
                borderRadius: 30,
                marginTop: 20,
                width: 300,
                alignItems: "center",
                justifyContent: "center",
                height: 44,
                shadowColor: "#000",
                shadowOffset: {
                  width: 0,
                  height: 2
                },
                shadowOpacity: 0.25,
                shadowRadius: 3.84,
                elevation: 5
              }}
            >
              <Text
                style={{
                  color: "black",
                  fontWeight: "500",
                  fontSize: 20,
                  fontFamily: "DMSansMedium"
                }}
              >
                Sign in with Google
              </Text>
            </TouchableOpacity>
            {isAppleLoggedIn && <Text style={{ marginTop: 20 }}>Logged in with Apple!</Text>}
            {!isAppleLoggedIn && (
              <AppleAuthentication.AppleAuthenticationButton
                buttonType={AppleAuthentication.AppleAuthenticationButtonType.SIGN_IN}
                buttonStyle={AppleAuthentication.AppleAuthenticationButtonStyle.WHITE_OUTLINE}
                cornerRadius={30}
                style={{
                  width: 200,
                  height: 50,
                  marginTop: 20,
                  shadowColor: "#000",
                  shadowOffset: {
                    width: 0,
                    height: 2
                  },
                  minWidth: 300,
                  shadowOpacity: 0.25,
                  shadowRadius: 3.84,
                  elevation: 5
                }}
                // eslint-disable-next-line react/jsx-no-bind
                onPress={handleAppleSignIn} // TODO: Remove
              />
            )}
          </View>
        </View>
      )}
    </View>
  );
}
