// import { router } from "expo-router";
// import React from "react";
// import { Image, SafeAreaView, StyleSheet, Text, View } from "react-native";
// import ContinueButton from "../../components/Onboarding/ContinueButton";
// import OnboardingButton from "../../components/Onboarding/OnboardingButton";
// import TopBar from "../../components/Onboarding/TopBar";
// import scaleStyleSheet from "../../scaleStyles";

// const aboutInterestedInPicture = require("../../assets/interestedin.png");

// function AboutInterestedIn() {
//   return (
//     <SafeAreaView style={scaledStyles.container}>
//       <View style={scaledStyles.TopUiContainer}>
//         <TopBar
//           onBackPress={() => {
//             router.back();
//           }}
//           text="About Me"
//           selectedCount={1}
//         />
//       </View>
//       <View style={scaledStyles.mainContainer}>
//         <View>
//           <Image source={aboutInterestedInPicture} />
//           <View>
//             <Text style={scaledStyles.headerContainer}>I&apos;m interested in...</Text>
//           </View>
//           <View>
//             <View style={scaledStyles.buttonContainer}>
//               <View style={scaledStyles.button}>
//                 <OnboardingButton title="Man" onButtonClick={() => {}} />
//               </View>
//               <View style={scaledStyles.button}>
//                 <OnboardingButton title="Woman" onButtonClick={() => {}} />
//               </View>
//               <View style={scaledStyles.button}>
//                 <OnboardingButton title="All" onButtonClick={() => {}} />
//               </View>
//             </View>
//           </View>
//         </View>
//         <View style={scaledStyles.ContinueButtonContainer}>
//           <ContinueButton
//             title="Continue"
//             isDisabled={false}
//             onPress={() => {
//               router.push("/AboutMe/AboutLooking");
//             }}
//           />
//         </View>
//       </View>
//     </SafeAreaView>
//   );
// }

// export default AboutInterestedIn;

// const styles = StyleSheet.create({
//   TopUiContainer: {
//     flex: 0.3,
//     alignItems: "center"
//   },
//   mainContainer: {
//     flex: 1,
//     marginLeft: 20,
//     marginRight: 20,
//     justifyContent: "space-between"
//   },
//   headerContainer: {
//     fontSize: 32,
//     fontWeight: "700",
//     lineHeight: 32,
//     letterSpacing: -0.32,
//     marginTop: 16,
//     marginBottom: 16,
//     fontFamily: "DMSansMedium"
//   },
//   textHelper: {
//     fontSize: 12,
//     fontWeight: "400",
//     lineHeight: 12,
//     letterSpacing: -0.12,
//     fontFamily: "DMSansMedium"
//   },
//   container: {
//     flex: 1,
//     marginTop : 30
//   },
//   ContinueButtonContainer: {
//     marginBottom: 10
//   },
//   button: {
//     marginBottom: 16
//   },
//   buttonText: {
//     color: "black",
//     fontSize: 17,
//     fontWeight: "500",
//     letterSpacing: -0.17,
//     fontFamily: "DMSansMedium"
//   }
// });

// const scaledStyles = scaleStyleSheet(styles);

import { router } from "expo-router";
import React from "react";
import { useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";
import OnboardingButton from "../../components/Onboarding/OnboardingButton";
import OnboardingTitle from "../../components/Onboarding/OnboardingTitle";

const aboutInterestedInPicture = require("../../assets/interestedin.png");

function AboutInterestedIn() {
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
    router.push("/AboutMe/AboutLooking");
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
      <View style={scaledStyles.mainContainer}>
        <View>
          <Image source={aboutInterestedInPicture} />
          <OnboardingTitle text="I&apos;m interested in..."/>
          <View style={scaledStyles.inputWrapper} />
          <View style={scaledStyles.buttonContainer}>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Man" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="Woman" onButtonClick={() => {}} />
              </View>
              <View style={scaledStyles.button}>
                <OnboardingButton title="All" onButtonClick={() => {}} />
              </View>
            </View>

        </View>

        <View>
          <ContinueButton
            title="Continue"
            isDisabled={false}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutInterestedIn;

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
  helperContainer: {
    marginTop: 16
  },
  button: {
    marginBottom: 16
  }
});

const scaledStyles = scaleStyleSheet(styles);
