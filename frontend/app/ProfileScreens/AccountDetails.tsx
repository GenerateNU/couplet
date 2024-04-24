import { useRouter } from "expo-router";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import CardWrapper from "../../components/Profile/CardWrapper";
import EditAccountDetailCard from "../../components/Profile/EditAccountDetailCard";
<<<<<<< HEAD

export default function AccountDetails() {
  const router = useRouter();
  return (
    <SafeAreaView>
      <View>
        {/* <Button onPress={() => router.back()}> */}
        {/* <Text onPress={() => router.back()} style={styles.title}>{`< ${name}`}</Text> */}
        <Text onPress={() => router.back()} style={styles.title}>{`< Account Details`}</Text>

        {/* </Button> */}
=======
import { useAppSelector } from "../../state/hooks";

export default function AccountDetails() {
  const router = useRouter();
  const user = useAppSelector((state) => state.form);

  return (
    <SafeAreaView>
      <View>
        <Text onPress={() => router.back()} style={styles.title}>{`< Account Details`}</Text>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      </View>
      <View style={{ width: "100%" }}>
        <CardWrapper>
          {/* //eslint-disable-next-line global-require */}
<<<<<<< HEAD
          <EditAccountDetailCard description="Name" fieldInfo="NameDummy" />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="Pronouns"
            fieldInfo="PronounsDummy"
=======
          <EditAccountDetailCard description="Name" fieldInfo={user.name} />
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            description="Pronouns"
            fieldInfo={user.pronouns}
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
            editable
            onPress={() => router.push("ProfileScreens/EditPronouns")}
          />
          {/* //eslint-disable-next-line global-require */}
<<<<<<< HEAD
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
=======
          <EditAccountDetailCard
            description="Birthday"
            fieldInfo={new Date(user.birthday).toLocaleDateString()}
          />
          {/* //eslint-disable-next-line global-require */}
          {/* <EditAccountDetailCard
            fieldInfo={user.email}
            editable
            description="Email"
            onPress={() => router.push("ProfileScreens/EditEmail")}
          /> */}
          {/* //eslint-disable-next-line global-require */}
          <EditAccountDetailCard
            fieldInfo={user.location}
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
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
