import {
  faBookOpen,
  faCannabis,
  faCapsules,
  faScaleBalanced,
  faSmoking,
  faUserGroup,
  faWineGlass
} from "@fortawesome/free-solid-svg-icons";
import { FontAwesomeIcon } from "@fortawesome/react-native-fontawesome";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import scaleStyleSheet from "../../scaleStyles";
import { LifestyleProps } from "./PersonProps";

export default function Lifestyle({
  relationshipType,
  religion,
  politicalAffiliation,
  alchoholFrequency,
  smokingFrequency,
  drugFrequency,
  cannabisFrequency
}: LifestyleProps) {
  return (
    <View style={scaledStyles.container}>
      {relationshipType && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faUserGroup} />
          <Text>{relationshipType}</Text>
        </View>
      )}
      {religion && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faBookOpen} />
          <Text>{religion}</Text>
        </View>
      )}
      {politicalAffiliation && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faScaleBalanced} />
          <Text>{politicalAffiliation}</Text>
        </View>
      )}
      {alchoholFrequency && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faWineGlass} />
          <Text>{alchoholFrequency}</Text>
        </View>
      )}
      {smokingFrequency && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faSmoking} />
          <Text>{smokingFrequency}</Text>
        </View>
      )}
      {drugFrequency && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faCapsules} />
          <Text>{drugFrequency}</Text>
        </View>
      )}
      {cannabisFrequency && (
        <View style={scaledStyles.lifestyleItemRow}>
          <FontAwesomeIcon style={scaledStyles.iconStyle} icon={faCannabis} />
          <Text>{cannabisFrequency}</Text>
        </View>
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    display: "flex",
    flexDirection: "row",
    flexWrap: "wrap",
    justifyContent: "flex-start"
  },
  lifestyleItemRow: {
    flexDirection: "row",
    display: "flex",
    alignItems: "center",
    margin: 7,
    marginRight: 10
  },
  iconStyle: {
    marginRight: 5,
    color: "#6B5CBF",
    fontSize: 15
  },
  iconTextStyle: {
    fontFamily: "DMSansRegular",
    fontSize: 15
  }
});

const scaledStyles = scaleStyleSheet(styles);
