import { useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import DropDownGender from "../../components/Profile/DropDownGender";
import { useAppSelector } from "../../state/hooks";

export default function EditPreferredGender() {
  const router = useRouter();
  const user = useAppSelector((state) => state.form);

  return (
    <SafeAreaView>
      <TouchableOpacity onPress={() => router.back()}>
        <Text style={styles.title}>{`< Edit Preferred Gender`}</Text>
      </TouchableOpacity>
      <View style={styles.container}>
        <DropDownGender
          onGenderChange={() => "placeHolder"}
          selectedPronoun={user.genderPreference}
        />
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
