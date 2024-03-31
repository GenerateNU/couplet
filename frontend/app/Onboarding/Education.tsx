// import { useNavigation } from "@react-navigation/native";
// import React, { useState } from "react";
// import {
//   Image,
//   KeyboardAvoidingView,
//   Platform,
//   SafeAreaView,
//   StyleSheet,
//   TextInput,
//   View
// } from "react-native";

// import { router } from "expo-router";
// import ContinueButton from "../../components/Onboarding/ContinueButton";
// import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
// import TopBar from "../../components/Onboarding/TopBar";
// import scaleStyleSheet from "../../scaleStyles";

// const SCHOOL = require("../../assets/lightbulb.png");

// export default function Education() {
//   const [school, setSchool] = useState("");
//   const navigation = useNavigation();
//   const nextRoute = "Onboarding/Career";

//   function goBack() {
//     if (navigation.canGoBack()) {
//       navigation.goBack();
//     }
//   }

//   return (
// <SafeAreaView style={scaledStyles.container}>
//   <KeyboardAvoidingView
//     behavior={Platform.OS === "ios" ? "padding" : "height"}
//     style={scaledStyles.container}
//   >
//     <View style={scaledStyles.TopUiContainer}>
//       <TopBar
// onBackPress={() => goBack()}
// text="Education and Career"
// selectedCount={2}
// skipToRoute={nextRoute}
//       />
//     </View>
//     <View style={scaledStyles.mainContainer}>
//       <View>
//   <Image source={SCHOOL} />
//   <OnboardingTitle text="My school is..." />
//   <TextInput
//     editable
//     value={school}
//     onChangeText={setSchool}
//     style={scaledStyles.textInput}
//     placeholder="Name of School"
//   />
// </View>

// <View>
// <ContinueButton
//   onPress={() => router.push(nextRoute)}
//   title="Continue"
//   isDisabled={school.length === 0}
// />
// </View>
//     </View>
//   </KeyboardAvoidingView>
// </SafeAreaView>
//   );
// }

// const styles = StyleSheet.create({
// textInput: {
//   borderStyle: "solid",
//   borderWidth: 1,
//   borderColor: "#9EA3A2",
//   color: "#000000",
//   borderRadius: 10,
//   padding: 10,
//   fontFamily: "DMSansRegular"
// },
//   container: {
//     flex: 1,
//     marginTop: 34,
//     marginBottom: 36
//   },
//   mainContainer: {
//     flex: 1,
//     marginLeft: 20,
//     marginRight: 20,
//     justifyContent: "space-between"
//   },
//   TopUiContainer: {
//     alignItems: "center",
//     flex: 0.3
//   },
// });

// const scaledStyles = scaleStyleSheet(styles);

import { router } from "expo-router";
import React, { useState } from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, KeyboardAvoidingView, Platform, StyleSheet, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

const SCHOOL = require("../../assets/lightbulb.png");

function Education() {
  const [school, setSchool] = useState("");
  const nextRoute = "Onboarding/Career";
  const { control, handleSubmit } = useForm({
    defaultValues: {
      name: ""
    }
  });
  const name = useWatch({
    control,
    name: "name",
    defaultValue: ""
  });
  const onSubmit = (data: Object) => {
    console.log(name);
    router.push("/Onboarding/Education");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.TopUiContainer}>
        <TopBar
          onBackPress={() => router.back()}
          text="Education and Career"
          selectedCount={2}
          skipToRoute={nextRoute}
        />
      </View>
      <KeyboardAvoidingView
        style={scaledStyles.avoidContainer}
        behavior={Platform.OS === "ios" ? "padding" : "height"}
      >
        <View style={scaledStyles.mainContainer}>
          <View>
            <Image source={SCHOOL} />
            <OnboardingTitle text="My school is..." />
            <TextInput
              editable
              value={school}
              onChangeText={setSchool}
              style={scaledStyles.textInput}
              placeholder="Name of School"
            />
          </View>
        </View>
      </KeyboardAvoidingView>
      <View>
        <ContinueButton
          onPress={() => router.push(nextRoute)}
          title="Continue"
          isDisabled={school.length === 0}
        />
      </View>
    </SafeAreaView>
  );
}

export default Education;

const styles = StyleSheet.create({
  TopUiContainer: {
    alignItems: "center",
    flex: 0.3
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1,
    marginTop: 34,
    marginBottom: 36
  },
  textInput: {
    borderStyle: "solid",
    borderWidth: 1,
    borderColor: "#9EA3A2",
    color: "#000000",
    borderRadius: 10,
    padding: 10,
    fontFamily: "DMSansRegular"
  },
  avoidContainer: {
    flex: 1
  }
});

const scaledStyles = scaleStyleSheet(styles);
