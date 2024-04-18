import { useRouter } from "expo-router";
import React from "react";
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
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="I'm interested in"
            fieldInfo={user.genderPreference}
            editable
            onPress={() => router.push("ProfileScreens/EditPreferredGender")}
          />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            editable
            last
            description="I'm looking for"
            fieldInfo={user.looking}
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
