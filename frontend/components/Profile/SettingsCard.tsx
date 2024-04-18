/* eslint-disable */
import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";

export default function SettingsCard({
  text,
  img,
  onPress,
  last = false,
  ...props
}: {
  text: string;
  img: any;
  last?: boolean;
  onPress: () => void;
}) {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  return (
    <TouchableOpacity onPress={onPress}>
      <View style={{ ...styles.container, borderBottomWidth: last ? 0 : 0.5 }}>
        <Image source={img} style={styles.imageStyle} />
        <Text style={styles.mainText}>{text}</Text>
        {/* // es-lint-disable-next-line */}
        <Image source={require("../../assets/Vector.png")} style={styles.arrow} />
      </View>
    </TouchableOpacity>
  );
}

const styles = StyleSheet.create({
  container: {
    padding: 15,
    width: "100%",
    alignSelf: "center",
    backgroundColor: "#ffffff",
    borderBottomColor: "#CDCDCD",
    flexDirection: "row"
  },
  imageStyle: {
    width: 38,
    height: 39,
    alignSelf: "center"
  },
  mainText: {
    verticalAlign: "middle",
    margin: "auto",
    alignSelf: "center",
    padding: 10,
    fontFamily: "DMSansRegular",
    fontWeight: "400",
    fontSize: 17
  },
  arrow: {
    width: 8,
    height: 15,
    alignSelf: "center",
    right: 20,
    position: "absolute"
  }
});
