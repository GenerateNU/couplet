/* eslint-disable */
import * as AppleAuthentication from "expo-apple-authentication";
import { router } from "expo-router";
import * as SecureStore from "expo-secure-store";
import React, { useEffect, useState } from "react";
import {
  Image,
  ImageBackground,
  SafeAreaView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View
} from "react-native";
import appleLogo from "../../assets/appleLogo.png";
import googleLogo from "../../assets/googleLogo.png";
import gradient from "../../assets/gradient.png";
import logo from "../../assets/logo.png";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

// GoogleSignin.configure({
//   scopes: ["https://www.googleapis.com/auth/drive.readonly"],
//   webClientId: Config.GOOGLE_WEB_CLIENT_ID,
//   iosClientId: Config.IOS_CLIENT_ID
// });

export default function Login() {
  // eslint-disable-next-line no-unused-vars
  const [isGoogleLoggedIn, setIsGoogleLoggedIn] = useState(false);
  const [isAppleLoggedIn, setIsAppleLoggedIn] = useState(false);
  // eslint-disable-next-line no-unused-vars
  const isSignedIn = isGoogleLoggedIn || isAppleLoggedIn;
  const [isLoading, setIsLoading] = useState(false);

  const handleGoogleSignIn = async () => {
    router.push("Home");
    // try {
    //   await GoogleSignin.hasPlayServices();
    //   const userInfo = await GoogleSignin.signIn();
    //   setIsGoogleLoggedIn(true);
    // } catch (error) {
    //   console.error(error);
    //   setIsGoogleLoggedIn(false);
    // }
  };

  const handleAppleSignIn = async () => {
    try {
      const creds = await AppleAuthentication.signInAsync({
        requestedScopes: [
          AppleAuthentication.AppleAuthenticationScope.FULL_NAME,
          AppleAuthentication.AppleAuthenticationScope.EMAIL
        ]
      });

      await SecureStore.setItemAsync("appleAuth", creds.user);
      router.push("Home");
      setIsAppleLoggedIn(true);
      router.push("Onboarding/AboutMe/AboutName");
    } catch (e) {
      setIsAppleLoggedIn(false);
    }
  };

  useEffect(() => {
    const checkAppleAuth = async () => {
      setIsLoading(true);
      const appleAuth = await SecureStore.getItemAsync("appleAuth");

      // TODO: check apple credentials against backend user table, route to either
      // onboarding or home page based on whether we find a user

      setIsLoading(false);
      if (appleAuth !== null) {
        router.push("Home");
      }
    };
    checkAppleAuth();
  }, []);

  return (
    <ImageBackground source={gradient} style={{ flex: 1 }} resizeMode="cover">
      <SafeAreaView style={scaledStyles.outerView}>
        {isLoading ? (
          <View style={{ height: "100%" }}>
            <Text>LOADING</Text>
          </View>
        ) : (
          <View style={scaledStyles.innerView}>
            <View style={scaledStyles.titleImageView}>
              <Image style={scaledStyles.coupletLogo} source={logo} />
              <Text style={scaledStyles.coupletText}>Couplet</Text>
            </View>

            {/* Texts below the title/image */}
            <View style={scaledStyles.textsView}>
              <Text style={scaledStyles.headerText}>Create an account</Text>
              <Text style={scaledStyles.bodyText}>
                Linking your account makes it easier to sign in
              </Text>
            </View>

            {/* Buttons */}
            <View style={scaledStyles.buttonsView}>
              <TouchableOpacity style={scaledStyles.button} onPress={handleAppleSignIn}>
                <Image source={appleLogo} style={scaledStyles.appleLogo} />
                <Text style={scaledStyles.buttonText}>Sign up with Apple</Text>
              </TouchableOpacity>
              <TouchableOpacity style={scaledStyles.button} onPress={handleGoogleSignIn}>
                <Image source={googleLogo} style={scaledStyles.googleLogo} />
                <Text style={scaledStyles.buttonText}>Sign up with Google</Text>
              </TouchableOpacity>
            </View>
          </View>
        )}
      </SafeAreaView>
    </ImageBackground>
  );
}

const styles = StyleSheet.create({
  outerView: {
    width: 393,
    height: 852,
    justifyContent: "flex-start"
  },
  innerView: {
    width: 356,
    height: 438.34,
    position: "absolute",
    top: 205,
    left: 24,
    gap: 24
  },
  coupletLogo: {
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 0 },
    shadowRadius: 10,
    shadowOpacity: 0.5
  },
  appleLogo: {
    width: 16,
    height: 20,
    marginRight: 8
  },
  googleLogo: {
    width: 16,
    height: 16,
    marginRight: 8
  },
  titleImageView: {
    width: 356,
    height: 226.34,
    left: -12,
    justifyContent: "flex-end",
    alignItems: "center",
    marginBottom: 4
  },
  coupletText: {
    width: "100%",
    height: 60,
    fontFamily: "DMSansBold",
    fontSize: 60,
    lineHeight: 65,
    letterSpacing: -0.05,
    textAlign: "center",
    color: COLORS.white,
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 0 },
    shadowRadius: 25,
    shadowOpacity: 1
  },
  textsView: {
    width: 346,
    alignItems: "center",
    marginBottom: 8
  },
  headerText: {
    width: 346,
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    fontFamily: "DMSansBold",
    color: COLORS.white,
    alignItems: "center",
    textAlign: "center"
  },
  bodyText: {
    width: 330,
    fontSize: 15,
    lineHeight: 19.53,
    fontFamily: "DMSansRegular",
    color: COLORS.white,
    textAlign: "center"
  },
  buttonsView: {
    width: 346,
    flexDirection: "column",
    justifyContent: "space-between"
  },
  button: {
    width: 346,
    height: 50,
    paddingVertical: 14,
    paddingHorizontal: 40,
    borderRadius: 100,
    borderColor: "#000",
    justifyContent: "center",
    alignItems: "center",
    backgroundColor: COLORS.white,
    marginBottom: 10,
    flexDirection: "row",
    shadowColor: COLORS.white,
    shadowOffset: { width: 0, height: 2 },
    shadowRadius: 4,
    shadowOpacity: 0.5
  },
  buttonText: {
    fontFamily: "DMSansMedium",
    fontSize: 17,
    lineHeight: 22.13,
    textAlign: "left",
    color: COLORS.black
  }
});

const scaledStyles = scaleStyleSheet(styles);
