import { useRouter } from "expo-router";
import React from "react";
<<<<<<< HEAD
import { StyleSheet, Text, TextInput, View } from "react-native";
=======
import { StyleSheet, Text, TextInput, TouchableOpacity, View } from "react-native";
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
import { SafeAreaView } from "react-native-safe-area-context";

export default function EditEmail() {
  const router = useRouter();
  return (
    <SafeAreaView>
<<<<<<< HEAD
      <Text onPress={() => router.back()} style={styles.title}>{`< Edit Email`}</Text>
=======
      <TouchableOpacity onPress={() => router.back()}>
        <Text style={styles.title}>{`< Edit Email`}</Text>
      </TouchableOpacity>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      <View style={styles.container}>
        <TextInput
          style={{ height: 40, borderColor: "gray", borderWidth: 1 }}
          onChangeText={(text) => "PlaceHolder"}
          value="EmailDummy"
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
  },
  input: {
    height: 40,
    borderColor: "gray",
    borderWidth: 1
  }
});
