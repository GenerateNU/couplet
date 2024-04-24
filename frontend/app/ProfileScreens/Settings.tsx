import { useRouter } from "expo-router";
import React, { useState } from "react";
<<<<<<< HEAD
import { StyleSheet, Switch, Text, View } from "react-native";
=======
import { StyleSheet, Switch, Text, TouchableOpacity, View } from "react-native";
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
import { SafeAreaView } from "react-native-safe-area-context";

function ToggleSwitch({ text }: { text: string }) {
  const [isEnabled, setIsEnabled] = useState(false);
  const toggleSwitch = () => setIsEnabled((previousState) => !previousState);

  return (
    <View style={styles.textToggle}>
      <Text style={styles.text2}>{text}</Text>
      <Switch
        trackColor={{ false: "#F84949", true: "#F84949" }}
        onValueChange={toggleSwitch}
        value={isEnabled}
      />
    </View>
  );
}

export default function Settings() {
  const router = useRouter();
  return (
    <SafeAreaView>
<<<<<<< HEAD
      <Text onPress={() => router.back()} style={styles.title}>{`< Settings`}</Text>
=======
      <TouchableOpacity onPress={() => router.back()}>
        <Text style={styles.title}>{`< Settings`}</Text>
      </TouchableOpacity>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      <View style={styles.container}>
        <Text style={styles.text1}>Pause Profile</Text>
        <ToggleSwitch text="Pausing your profile will temporarily stop showing your account to others. You can still see your current matches." />
        <View style={styles.line} />
        <Text style={styles.text1}>NotificationSettings</Text>
        <ToggleSwitch text="New likes" />
        <ToggleSwitch text="New matches" />
        <ToggleSwitch text="Promotions and announcements" />
      </View>
    </SafeAreaView>
  );
}

const styles = StyleSheet.create({
  containerQ: {
    flex: 1,
    alignItems: "center",
    justifyContent: "center"
  },
  title: {
    fontFamily: "DMSansMedium",
    fontSize: 32,
    fontWeight: "700",
    lineHeight: 32,
    marginLeft: 16,
    marginTop: 16
  },
  container: {
    padding: 5,
    borderRadius: 20,
    width: "90%",
    alignSelf: "center",
    marginTop: 40
  },
  text1: {
    fontFamily: "DMSansRegular",
    fontSize: 15,
    lineHeight: 24,
    fontWeight: "400",
    color: "#8A8A8A",
    marginVertical: 5
  },
  text2: {
    fontFamily: "DMSansRegular",
    fontSize: 15,
    lineHeight: 24,
    fontWeight: "400",
    marginBottom: 10,
    maxWidth: "85%"
  },
  line: {
    borderBottomColor: "#CDCDCD",
    borderBottomWidth: 0.5,
    marginVertical: 5
  },
  textToggle: {
    flexDirection: "row",
    justifyContent: "space-between",
    alignItems: "center",
    marginVertical: 5
  }
});
