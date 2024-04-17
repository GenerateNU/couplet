import { useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import CardWrapper from "../../components/Profile/CardWrapper";
import EditAccountDetailCard from "../../components/Profile/EditAccountDetailCard";

export default function AccountDetails() {
  const router = useRouter();
  return (
    <SafeAreaView>
      <View>
        {/* <Button onPress={() => router.back()}> */}
        {/* <Text onPress={() => router.back()} style={styles.title}>{`< ${name}`}</Text> */}
        <Text onPress={() => router.back()} style={styles.title}>{`< Account Details`}</Text>

        {/* </Button> */}
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard description="Name" fieldInfo="NameDummy" />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="Pronouns"
            fieldInfo="PronounsDummy"
            editable
            onPress={() => router.push("ProfileScreens/EditPronouns")}
          />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard description="Birthday" fieldInfo="BirthdayDummy" />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            fieldInfo="EmailDummy"
            editable
            description="Email"
            onPress={() => router.push("ProfileScreens/EditEmail")}
          />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            fieldInfo="NeighborhoodDummy"
            editable
            last
            description="Neighborhood"
            onPress={() => router.push("ProfileScreens/EditNeighborhood")}
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
