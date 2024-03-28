import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch } from "react-hook-form";
import { Image, StyleSheet, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

function AboutName() {
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
    router.push("/AboutMe/AboutBirthday");
  };
  return (
    <SafeAreaView style={scaledStyles.container}>
      <View style={scaledStyles.mainContainer}>
        <View style={scaledStyles.ProgressBarContainer}>
          <TopBar
            onBackPress={() => {
              router.back();
            }}
            text="About Me"
            selectedCount={1}
          />
        </View>

        <View style={scaledStyles.inputContainer}>
          <Image source={require("../../assets/aboutName.png")} />
          <Text style={scaledStyles.headerContainer}>My first name is...</Text>
          <View style={scaledStyles.inputWrapper}>
            <Controller
              control={control}
              render={({ field: { onChange, onBlur, value } }) => (
                <TextInput
                  style={scaledStyles.textContainer}
                  placeholder="First Name"
                  onBlur={onBlur}
                  onChangeText={onChange}
                  value={value}
                />
              )}
              name="name"
            />
          </View>
          <Text style={scaledStyles.textHelper}>
            This is how it will permanently appear on your profile
          </Text>
        </View>

        <View>
          <ContinueButton
            title="Continue"
            isDisabled={!name}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </SafeAreaView>
  );
}

export default AboutName;

const styles = StyleSheet.create({
  container: {
    flex: 1
  },
  ProgressBarContainer: {
    justifyContent: "center"
  },
  mainContainer: {
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between",
    flex: 1
  },
  headerContainer: {
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    letterSpacing: -0.32,
    marginTop: 16,
    marginBottom: 16,
    fontFamily: "DMSansMedium"
  },
  inputContainer: {
    flex: 0.55
  },
  textContainer: {
    padding: 8
  },
  inputWrapper: {
    borderRadius: 6,
    borderWidth: 1,
    borderColor: "grey",
    marginBottom: 8
  }
});

const scaledStyles = scaleStyleSheet(styles);
