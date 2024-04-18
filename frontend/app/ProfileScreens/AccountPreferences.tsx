import { useRouter } from "expo-router";
import React from "react";
<<<<<<< HEAD
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import CardWrapper from "../../components/Profile/CardWrapper";
import EditAccountDetailCard from "../../components/Profile/EditAccountDetailCard";

export default function AccountPreferences() {
  const router = useRouter();
  return (
    <SafeAreaView>
      <View>
        {/* <Button onPress={() => router.back()}> */}
        {/* <Text onPress={() => router.back()} style={styles.title}>{`< ${name}`}</Text> */}
        <Text onPress={() => router.back()} style={styles.title}>{`< Account Preferences`}</Text>

        {/* </Button> */}
=======
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import CardWrapper from "../../components/Profile/CardWrapper";
import EditAccountDetailCard from "../../components/Profile/EditAccountDetailCard";
import { useAppSelector } from "../../state/hooks";

export default function AccountPreferences() {
  const router = useRouter();
  const user = useAppSelector((state) => state.form);

  return (
    <SafeAreaView>
      <View>
        <TouchableOpacity onPress={() => router.back()}>
          <Text style={styles.title}>{`< Account Preferences`}</Text>
        </TouchableOpacity>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="I'm interested in"
<<<<<<< HEAD
            fieldInfo="men"
=======
            fieldInfo={user.genderPreference}
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
            editable
            onPress={() => router.push("ProfileScreens/EditPreferredGender")}
          />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            editable
            last
            description="I'm looking for"
<<<<<<< HEAD
            fieldInfo="Long term relationship"
=======
            fieldInfo={user.looking}
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
            onPress={() => router.push("ProfileScreens/EditPreferredRelationship")}
          />
        </CardWrapper>
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
  }
});
