// import { router } from "expo-router";
// import React, { useState } from "react";
// import {
//   Image,
//   KeyboardAvoidingView,
//   Platform,
//   SafeAreaView,
//   StyleSheet,
//   Text,
//   TextInput,
//   View
// } from "react-native";
// import ContinueButton from "../../../components/Onboarding/ContinueButton";
// import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
// import TopBar from "../../../components/Onboarding/TopBar";
// import scaleStyleSheet from "../../../scaleStyles";

// const AT_ICON = require("../../../assets/3diconslmao.png");

// export default function Profile() {
//   const [instagram, setInstagram] = useState("");

//   return (
//     <SafeAreaView style={{ height: "100%" }}>
//       <KeyboardAvoidingView
//         behavior={Platform.OS === "ios" ? "padding" : "height"}
//         style={scaledStyles.container}
//       >
//         <View>
//           <TopBar onBackPress={() => router.back()} text="Profile" selectedCount={0} />
//         </View>
//         <View>
//           <Image source={AT_ICON} style={{ height: 250, width: 250 }} resizeMode="contain" />
//           <View>
//             <OnboardingTitle text="What’s your Instagram?" />
//             <Text style={{ marginLeft: 0, top: 0 }}>
//               This is how you will message your matches.
//             </Text>
//             <TextInput
//               style={scaledStyles.textInput}
//               editable
//               onChangeText={setInstagram}
//               placeholder="@couplet"
//             />
//           </View>
//         </View>
//         <View>
//           <ContinueButton
//             onPress={() => router.push("Onboarding/ProfileNotifications")}
//             title="Continue"
//             isDisabled={instagram.length === 0}
//           />
//         </View>
//       </KeyboardAvoidingView>
//     </SafeAreaView>
//   );
// }

// const styles = StyleSheet.create({
//   container: {
//     flex: 1,
//     justifyContent: "space-between",
//     alignItems: "center",
//     margin: 30
//   },
//   textInput: {
//     borderStyle: "solid",
//     borderWidth: 1,
//     borderColor: "#9EA3A2",
//     color: "#000000",
//     borderRadius: 10,
//     padding: 10,
//     fontFamily: "DMSansRegular"
//   },
//   disableNotificationsButton: {
//     fontFamily: "DMSansMedium",
//     textAlign: "center",
//     marginTop: 15,
//     fontSize: 16
//   }
// });

// const scaledStyles = scaleStyleSheet(styles);

import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Image, KeyboardAvoidingView, Platform, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../../components/Onboarding/OnboardingTitle";
import TopBar from "../../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../../scaleStyles";
import { setInstagram } from "../../../state/formSlice";
import { useAppDispatch } from "../../../state/hooks";
import onboardingStyles from "../../../styles/Onboarding/styles";

const AT_ICON = require("../../../assets/instagramAt.png");

function ProfileInsta() {
  const dispatch = useAppDispatch();
  // Use Form from React-Hook-Form
  const { control, handleSubmit } = useForm({
    defaultValues: {
      username: ""
    }
  });
  // Watch any changes made to the input form
  const username = useWatch({
    control,
    name: "username",
    defaultValue: ""
  });
  // On submit of the name form
  const onSubmit = (data: { username: string }) => {
    dispatch(setInstagram(data.username));
    router.push("Onboarding/Profile/ProfileNotifications");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => {
            router.back();
          }}
          text="About Me"
          selectedCount={1}
        />
      </View>
      <KeyboardAvoidingView
        style={scaledStyles.avoidContainer}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={AT_ICON} />
            <OnboardingTitle text="What’s your Instagram handle?" />
            <View style={scaledStyles.inputWrapper}>
              <View style={scaledStyles.textInputWrapper}>
                <Controller
                  control={control}
                  render={({ field: { onChange, onBlur, value } }) => (
                    <TextInput
                      style={scaledStyles.textInput}
                      placeholder="@your_username"
                      onBlur={onBlur}
                      onChangeText={onChange}
                      value={value}
                    />
                  )}
                  name="username"
                />
              </View>
              <Text style={scaledStyles.textHelper}>This is how you will message your matches</Text>
            </View>
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          title="Continue"
          isDisabled={!username}
          onPress={() => {
            handleSubmit(onSubmit)();
          }}
        />
      </View>
    </SafeAreaView>
  );
}

export default ProfileInsta;

const scaledStyles = scaleStyleSheet(onboardingStyles);
