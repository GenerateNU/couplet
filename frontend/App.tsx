import React, { useEffect, useState } from 'react';
import * as Google from 'expo-auth-session/providers/google';
import * as AppleAuthentication from 'expo-apple-authentication';
import { Button, View, Text } from 'react-native';
import { ResponseType, AuthSessionResult } from 'expo-auth-session';
import * as AuthSession from 'expo-auth-session';
import Home from './app/Home';

const redirectUri = process.env.NODE_ENV === 'development'
  ? 'https://auth.expo.io/@yourExpoUsername/yourApp'
  : AuthSession.makeRedirectUri({
      native: 'your.app://redirect',
    });

export default function App() {
  const [request, response, promptAsync] = Google.useIdTokenAuthRequest({
    clientId: 'your-google-client-id.apps.googleusercontent.com',
    redirectUri,
    responseType: ResponseType.IdToken,
    scopes: []
  });

  const [authToken, setAuthToken] = useState<string | null>(null);
  const [isAppleLoggedIn, setIsAppleLoggedIn] = useState(false); // Updated for clarity
  const isSignedIn = authToken !== null || isAppleLoggedIn;

  useEffect(() => {
    if (response?.type === 'success') {
      const { authentication } = response;
      if (authentication) {
        setAuthToken(authentication.accessToken);
      }
    }
  }, [response]);

  async function handleAppleSignIn() {
    try {
      const credential = await AppleAuthentication.signInAsync({
        requestedScopes: [
          AppleAuthentication.AppleAuthenticationScope.FULL_NAME,
          AppleAuthentication.AppleAuthenticationScope.EMAIL,
        ],
      });
      setIsAppleLoggedIn(true); // Updated to set a boolean flag indicating login status
    } catch (e) {
      if (e.code === 'ERR_CANCELED') {
        // Handle that the user canceled the sign-in flow
      } else {
        // Handle other possible errors
      }
    }
  }

  return (
    <View style={{ flex: 1 }}>
      {isSignedIn ? (
        <Home/>
      ) : (
<View style={{
  paddingTop: "100%",
  width: "100%",
  borderRadius: 12,
  alignSelf: "center",
  marginBottom: 10,
  minHeight: "100%",
  paddingBottom: "30%",
  alignItems: 'center', // Align items in the center horizontally
}}>
  <Button
    disabled={!request}
    title="Login with Google"
    onPress={() => {
      promptAsync();
    }}
  />
  {isAppleLoggedIn && <Text>Logged in with Apple!</Text>}
  {!isAppleLoggedIn && (
    <AppleAuthentication.AppleAuthenticationButton
      buttonType={AppleAuthentication.AppleAuthenticationButtonType.SIGN_IN}
      buttonStyle={AppleAuthentication.AppleAuthenticationButtonStyle.BLACK}
      cornerRadius={5}
      style={{ width: 200, height: 44, marginTop: 20 }} // Ensure consistent width and adjust marginTop if necessary
      onPress={handleAppleSignIn}
    />
  )}
  {authToken && <Text>Logged in with Google!</Text>}
</View>

    )}
    </View>
  );
}
