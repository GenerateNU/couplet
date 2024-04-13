import { useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import DropDownPronoun from "../../components/Profile/DropDownPronouns";

export default function EditPronouns() {
  const router = useRouter();
  return (
    <SafeAreaView>
      <Text onPress={() => router.back()} style={styles.title}>{`< Edit Pronouns`}</Text>
      <View style={styles.container}>
        <DropDownPronoun onPronounChange={() => "placeHolder"} selectedPronoun="She/Her" />
      </View>
    </SafeAreaView>
  );
}
const styles = StyleSheet.create({
  title: {
    fontFamily: "DMSansMedium",
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    marginLeft: 16
  },
  container: {
    padding: 5,
    borderRadius: 20,
    width: "90%",
    alignSelf: "center",
    marginTop: 40
  }
});
