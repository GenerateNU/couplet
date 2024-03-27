import React from "react";
import { StyleSheet, View } from "react-native";
import COLORS from "../../colors";
import { scaleWidth } from "../../ratio";
import scaleStyleSheet from "../../scaleStyles";

interface OnboardingBarProps {
  selectedCount: number;
}

function OnboardingBar({ selectedCount }: OnboardingBarProps) {
  return (
    <View style={[scaledStyles.container, { justifyContent: "center" }]}>
      <View style={{ flexDirection: "row", width: scaleWidth(346) }}>
        {Array.from({ length: 4 }).map((_, index) => (
          <View
            style={[
              scaledStyles.segment,
              index === selectedCount - 1
                ? scaledStyles.segmentSelected
                : scaledStyles.segmentUnselected,
              { marginRight: index < 3 ? 8 : 0 }
            ]}
          />
        ))}
      </View>
    </View>
  );
}

const styles = StyleSheet.create({
  centeringContainer: {
    alignItems: "center",
    width: "100%"
  },
  container: {
    width: 346,
    height: 61,
    justifyContent: "center",
    alignItems: "flex-start",
    opacity: 1
  },
  text: {
    fontFamily: "DMSansBold",
    fontSize: 32,
    fontWeight: "bold",
    lineHeight: 32,
    textAlign: "left",
    color: COLORS.black
  },
  segment: {
    width: 80.5,
    height: 3,
    borderRadius: 100
  },
  segmentSelected: {
    backgroundColor: COLORS.primary
  },
  segmentUnselected: {
    backgroundColor: COLORS.disabled
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default OnboardingBar;
