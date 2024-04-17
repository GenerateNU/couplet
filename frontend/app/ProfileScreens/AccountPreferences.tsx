import { useRouter } from "expo-router";
import React from "react";
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
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="I'm interested in"
            fieldInfo="men"
            editable
            onPress={() => router.push("ProfileScreens/EditPreferredGender")}
          />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            editable
            last
            description="I'm looking for"
            fieldInfo="Long term relationship"
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
