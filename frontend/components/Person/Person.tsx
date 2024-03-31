import { faBriefcase, faGraduationCap, faHouse, faRuler } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import React from "react";
import { Image, ScrollView, StyleSheet, Text, View } from "react-native";
import InfoChips from "./InfoChips";
import PersonHeader from "./PersonHeader";
import { PersonProps } from "./PersonProps";

export default function Person({
  id,
  firstName,
  lastName,
  age,
  pronouns,
  location,
  school,
  work,
  height,
  bio,
  interests,
  habits,
  instagram,
  images
}: PersonProps) {
  const firstImage = images[0]?.image || "";
  const heightText = height ? `${height.feet}'${height.inches}"` : "";

  return (
    <ScrollView
      style={{
        width: "100%",
        alignSelf: "center",
        height: "90%"
      }}
    >
      <View>
        {/* can we assume there are 4 images.. does this structure make sense */}
        <PersonHeader firstName={firstName} age={age} pronouns={pronouns} />
        <Image style={styles.imageStyle} source={{ uri: firstImage }} />
        <View style={styles.infoContainer}>
          <View style={styles.basicInfoContainer}>
            {location && (
              <View style={styles.basicInfoRow}>
                <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faHouse} />
                <Text style={styles.basicInfoRowText}>{location}</Text>
              </View>
            )}
            {school && (
              <View style={styles.basicInfoRow}>
                <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faGraduationCap} />
                <Text style={styles.basicInfoRowText}>{school}</Text>
              </View>
            )}
            {work && (
              <View style={styles.basicInfoRow}>
                <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faBriefcase} />
                <Text style={styles.basicInfoRowText}>{work}</Text>
              </View>
            )}
            {height && (
              <View style={styles.basicInfoRow}>
                <FontAwesomeIcon style={styles.basicInfoRowIcon} icon={faRuler} />
                <Text style={styles.basicInfoRowText}>{heightText}</Text>
              </View>
            )}
          </View>
          <View style={styles.bioContainer}>
            <Text style={styles.textStyle}>{bio}</Text>
          </View>
          <InfoChips items={interests} textColor="black" backgroundColor="lavender" />
          {/* PROBALY NEED TO CHANGE THIS>>>> */}
          {habits && <InfoChips items={habits} textColor="white" backgroundColor="purple" />}
          <View>
            <Text style={styles.textStyle}>For our first date, let's go to...</Text>
            {/* REUSE THE ITEM */}
          </View>
        </View>
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  imageStyle: {
    borderRadius: 10,
    width: "90%",
    height: 350,
    marginLeft: "auto",
    marginRight: "auto",
    boxShadow: "0px 0px 10px 0px rgba(0,0,0,0.75)"
  },
  textStyle: {
    fontFamily: "DMSansRegular"
  },
  infoContainer: {
    margin: 20
  },
  basicInfoContainer: {
    marginTop: 20
  },
  basicInfoRow: {
    flexDirection: "row",
    display: "flex",
    alignItems: "center"
  },
  basicInfoRowIcon: {
    marginRight: 5,
    float: "left"
  },
  basicInfoRowText: {
    fontFamily: "DMSansRegular",
    float: "right"
  },
  bioContainer: {
    fontFamily: "DMSansRegular",
    marginTop: 20,
    marginBottom: 20
  }
});
