import { useRouter } from "expo-router";
import React from "react";
<<<<<<< HEAD
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import DropDownLocation from "../../components/Onboarding/DropDownLocation";

export default function EditNeighborhood() {
  const router = useRouter();
  return (
    <SafeAreaView>
      <Text onPress={() => router.back()} style={styles.title}>{`< Edit Neighborhood`}</Text>
      <View style={styles.container}>
        <DropDownLocation onLocationChange={() => "placeHolder"} selectedLocation="Allston" />
=======
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import DropDownLocation from "../../components/Onboarding/DropDownLocation";
import { useAppSelector } from "../../state/hooks";

export default function EditNeighborhood() {
  const router = useRouter();
  const user = useAppSelector((state) => state.form);

  return (
    <SafeAreaView>
      <TouchableOpacity onPress={() => router.back()}>
        <Text style={styles.title}>{`< Edit Neighborhood`}</Text>
      </TouchableOpacity>

      <View style={styles.container}>
        <DropDownLocation onLocationChange={() => "placeHolder"} selectedLocation={user.location} />
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
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
