import { useRouter } from "expo-router";
import React from "react";
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
