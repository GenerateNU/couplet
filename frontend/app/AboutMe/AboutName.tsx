import { router } from "expo-router";
import React from "react";
import { Controller, useForm, useWatch} from "react-hook-form";
import { Image, StyleSheet, Text, TextInput, View } from "react-native";
import ContinueButton from "../../components/Onboarding/ContinueButton";
import TopBar from "../../components/Onboarding/TopBar";
import scaleStyleSheet from "../../scaleStyles";

function AboutName() {
  const {
    control,
    handleSubmit,
  } = useForm({
    defaultValues: {
      name: ""
    }
  });
  const name = useWatch({
    control,
    name: "name",
    defaultValue: ""
  })
  const onSubmit = (data: Object) => { 
    router.push("/AboutMe/AboutBirthday")
  }
  return (
    <View style={scaledStyles.container}>
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
          <Image source={require("../../assets/aboutName.png")} />
          <View>
            <Text style={scaledStyles.headerContainer}>My first name is...</Text>
          </View>

          <View style={scaledStyles.inputWrapper}>
            <Controller
              control={control}
              render={({ field: { onChange, onBlur, value } }) => (
                <TextInput
                  style={scaledStyles.inputContainer}
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
        <View style={scaledStyles.ContinueButtonContainer}>
          <ContinueButton
            title="Continue"
            isDisabled={!name}
            onPress={() => {
              handleSubmit(onSubmit)();
            }}
          />
        </View>
      </View>
    </View>
  );
}

export default AboutName;

const styles = StyleSheet.create({
  TopUiContainer: {
    flex: 0.3
  },
  mainContainer: {
    flex: 1,
    marginLeft: 20,
    marginRight: 20,
    justifyContent: "space-between"
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
    padding: 8
  },
  inputWrapper: {
    borderRadius: 6,
    borderWidth: 1,
    borderColor: "grey",
    marginBottom: 8
  },
  textHelper: {
    fontSize: 12,
    fontWeight: "400",
    lineHeight: 12,
    letterSpacing: -0.12,
    fontFamily: "DMSansMedium"
  },
  container: {
    flex: 1
  },
  ContinueButtonContainer: {
    marginBottom: 10
  }
});

const scaledStyles = scaleStyleSheet(styles);
