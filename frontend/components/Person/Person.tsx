import { faBriefcase, faGraduationCap, faHouse, faRuler } from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import React from "react";
import { Image, StyleSheet, Text, View } from "react-native";
import HeaderScrollView from "react-native-header-scroll-view";
import InfoChips from "./InfoChips";
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
  relationshipType,
  religion,
  politicalAffiliation,
  alchoholFrequency,
  smokingFrequency,
  drugFrequency,
  marijuanaFrequency,
  instagram,
  mutualEvents,
  images, 
  isMatched, 
}: PersonProps) {
  const firstImage = images[0]?.image || "";
  const heightText = height ? `${height.feet}'${height.inches}"` : "";

  return (
    <HeaderScrollView
      headerContainerStyle={{}}
      containerStyle={{
        flex: 1
      }}
      title={
        <>
          <Text style={{ fontFamily: "DMSansBold", marginRight: 10 }}>{firstName} </Text>
          <Text id="hi" style={{ fontFamily: "DMSansRegular" }}>
            {age}
          </Text>
        </>
      }
    >
      <Text
        style={{
          fontFamily: "DMSansRegular",
          color: "gray",
          fontSize: 14,
          marginLeft: 20,
          marginBottom: 10
        }}
      >
        {pronouns}
      </Text>

      <View>
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
          <View>
            <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>

            <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>

            <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>

            <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>

            <Text style={styles.textStyle}>For our first date, let&apos;s go to...</Text>
            {/* REUSE THE ITEM */}
          </View>
        </View>
      </View>
    </HeaderScrollView>
  );
}

const styles = StyleSheet.create({
  headerStyle: {
    backgroundColor: "white"
  },
  imageStyle: {
    borderRadius: 10,
    width: "90%",
    height: 350,
    marginLeft: "auto",
    marginRight: "auto"
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
