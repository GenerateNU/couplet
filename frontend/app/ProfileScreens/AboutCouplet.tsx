import { useRouter } from "expo-router";
import React from "react";
<<<<<<< HEAD
import { StyleSheet, Text, View } from "react-native";
=======
import { StyleSheet, Text, TouchableOpacity, View } from "react-native";
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
import { SafeAreaView } from "react-native-safe-area-context";

export default function AboutCouplet() {
  const router = useRouter();
  return (
    <SafeAreaView>
<<<<<<< HEAD
      <Text onPress={() => router.back()} style={styles.title}>{`< About Couplet`}</Text>
=======
      <TouchableOpacity onPress={() => router.back()}>
        <Text style={styles.title}>{`< About Couplet`}</Text>
      </TouchableOpacity>
>>>>>>> 526f5b11fbcdbafa24ca570f8948715853d839bb
      <View style={styles.container}>
        <Text style={styles.text1}>Who we are</Text>
        <Text style={styles.text2}>{mainText1}</Text>
        <Text style={styles.text2}>{mainText2}</Text>
        <Text style={styles.text2}>{mainText3}</Text>
        <Text style={styles.text2}>{mainText4}</Text>
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
    marginBottom: 10
  },
  text2: {
    fontFamily: "DMSansRegular",
    fontSize: 15,
    lineHeight: 24,
    fontWeight: "400",
    marginBottom: 10
  }
});

const mainText1 =
  "Couplet is a mobile app for young adults to meaningfully connect with others. In a world where superficial swipes often lead to short-lived interactions, Couplet serves as a movement towards deeper, experience-based relationships.";
const mainText2 =
  "Our journey was sparked by the vision of a dynamic duo, Victoria and Jan, who saw beyond the traditional dating scene. They imagined a platform that harnesses the energy of local events to ignite genuine connections among people. ";
const mainText3 =
  "This vision has become a reality, from ensuring a seamless user experience to creating an intuitive user interface that highlights events and shared interests. In doing so, Couplet serves as a platform where connections are not made by swiping on faces but by swiping on experiences. Imagine finding a fellow runner or sci-fi enthusiast based not on their profile picture but on their shared passion for the event. ";
const mainText4 =
  "Couplet is redefining social interactions while supporting the local economy and nonprofit sector. It's not just about finding a date for Friday night; it's about building a community, one shared experience at a time.";
