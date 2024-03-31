import { GoogleSignin } from "@react-native-google-signin/google-signin";
import * as AppleAuthentication from "expo-apple-authentication";
import React, { useState } from "react";
import {
  Image,
  ImageBackground,
  SafeAreaView,
  StyleSheet,
  Text,
  TouchableOpacity,
  View
} from "react-native";
import Config from "react-native-config";
import COLORS from "../../colors";
import scaleStyleSheet from "../../scaleStyles";

GoogleSignin.configure({
  scopes: ["https://www.googleapis.com/auth/drive.readonly"],
  webClientId: Config.GOOGLE_WEB_CLIENT_ID,
  iosClientId: Config.IOS_CLIENT_ID
});

export default function Login() {
  const [isGoogleLoggedIn, setIsGoogleLoggedIn] = useState(false);
  const [isAppleLoggedIn, setIsAppleLoggedIn] = useState(false);
  const isSignedIn = isGoogleLoggedIn || isAppleLoggedIn;

  async function handleGoogleSignIn() {
    try {
      await GoogleSignin.hasPlayServices();
      const userInfo = await GoogleSignin.signIn();
      setIsGoogleLoggedIn(true);
    } catch (error) {
      console.error(error);
      setIsGoogleLoggedIn(false);
    }
  }

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

  return (
    <ImageBackground source={require("./Gradient.png")} style={{ flex: 1 }} resizeMode="cover">
      <SafeAreaView style={scaledStyles.outerView}>
        <View style={scaledStyles.innerView}>
          <View style={scaledStyles.titleImageView}>
            <Image style={scaledStyles.coupletLogo} source={require("./Logo.png")} />
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
              <Image source={require("./AppleLogo.png")} style={scaledStyles.appleLogo} />
              <Text style={scaledStyles.buttonText}>Sign up with Apple</Text>
            </TouchableOpacity>
            <TouchableOpacity style={scaledStyles.button} onPress={handleGoogleSignIn}>
              <Image source={require("./GoogleLogo.png")} style={scaledStyles.googleLogo} />
              <Text style={scaledStyles.buttonText}>Sign up with Google</Text>
            </TouchableOpacity>
          </View>
        </View>
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
    background: COLORS.black,
    marginRight: 8
  },
  googleLogo: {
    width: 16,
    height: 16,
    background: COLORS.black,
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
