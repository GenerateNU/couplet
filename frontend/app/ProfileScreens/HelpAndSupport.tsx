import { useRouter } from "expo-router";
import React from "react";
import { Linking, StyleSheet, Text, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";

export default function HelpAndSupport() {
  const router = useRouter();

  const handleEmailPress = () => {
    Linking.openURL("mailto:support@couplet.com");
  };

  return (
    <SafeAreaView>
      <Text onPress={() => router.back()} style={styles.title}>{`< About Couplet`}</Text>
      <View style={styles.container}>
        <Text style={styles.text1}>Contact Support</Text>
        <Text style={styles.text2}>
          {mainText1}{" "}
          <Text style={styles.email} onPress={handleEmailPress}>
            support@couplet.com
          </Text>
        </Text>
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
    marginBottom: 5
  },
  text2: {
    fontFamily: "DMSansRegular",
    fontSize: 15,
    lineHeight: 24,
    fontWeight: "400",
    marginBottom: 10
  },
  email: {
    color: "red",
    textDecorationLine: "underline"
  }
});

const mainText1 = "Have questions or feedback for us? Get in touch by emailing us at";
