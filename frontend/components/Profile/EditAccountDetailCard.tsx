/* eslint-disable */

import { DMSans_400Regular as DMSansRegular } from "@expo-google-fonts/dm-sans";
import { useFonts } from "expo-font";
import React from "react";
import { Image, StyleSheet, Text, TouchableOpacity, View } from "react-native";

export default function EditAccountDetailCard({
  description,
  fieldInfo,
  onPress,
  editable,
  last = false,
  ...props
}: {
  description: string;
  fieldInfo: string;
  last?: boolean;
  editable?: boolean;
  onPress?: () => void;
}) {
  const [fontsLoaded] = useFonts({
    DMSansRegular
  });

  if (!fontsLoaded) {
    return null;
  }

  const ContainerComponent: React.ElementType = editable ? TouchableOpacity : View;

  return (
    <ContainerComponent onPress={editable ? onPress : undefined}>
      <View style={{ ...styles.container, borderBottomWidth: last ? 0 : 0.5 }}>
        <View style={styles.container2}>
          <Text style={styles.lightText}>{description}</Text>
          <Text style={styles.mainText}>{fieldInfo}</Text>
        </View>
        {/* // es-lint-disable-next-line */}
        {editable && <Image source={require("../../assets/editPencil.png")} style={styles.arrow} />}
      </View>
    </ContainerComponent>
  );
}

const styles = StyleSheet.create({
  container: {
    width: "100%",
    backgroundColor: "#ffffff",
    borderBottomColor: "#CDCDCD",
    flexDirection: "row"
  },
  container2: {
    paddingVertical: 15,
    flexDirection: "column"
  },
  lightText: {
    color: "#8A8A8A",
    fontFamily: "DMSansRegular",
    fontWeight: "400",
    fontSize: 17
  },
  mainText: {
    verticalAlign: "middle",
    fontFamily: "DMSansRegular",
    fontWeight: "400",
    fontSize: 17,
    paddingTop: 5
  },
  arrow: {
    width: 18,
    height: 18,
    alignSelf: "center",
    right: 20,
    position: "absolute"
  }
});
