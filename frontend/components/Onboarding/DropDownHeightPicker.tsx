import React, { useEffect, useState } from "react";
import { StyleSheet, View } from "react-native";
import DropDownPicker from "react-native-dropdown-picker";
import scaleStyleSheet from "../../scaleStyles";

interface DropDownCalendarProps {
  onHeightChange: (foot: number, inch: number) => void;
  selectedHeight: { foot: number; inch: number };
}

function DropDownHeightPicker({ onHeightChange, selectedHeight }: DropDownCalendarProps) {
  const [openFeet, setOpenFeet] = useState(false);
  const [openInches, setOpenInches] = useState(false);
  const [foot, setFoot] = useState(selectedHeight.foot);
  const [inch, setInch] = useState(selectedHeight.inch);
  const feet = [1, 2, 3, 4, 5, 6, 7, 8].map((feetParam, index) => ({
    label: `${feetParam}`,
    value: index + 1
  }));
  const inches = [0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11].map((inchParam, index) => ({
    label: `${inchParam}`,
    value: index + 1
  }));
  useEffect(() => {
    onHeightChange(foot, inch - 1);
  }, [foot, inch, onHeightChange]);
  return (
    <View style={scaledStyles.dropDownContainer}>
      <DropDownPicker
        open={openFeet}
        value={foot}
        items={feet}
        setOpen={setOpenFeet}
        setValue={setFoot}
        placeholder="Feet"
        containerStyle={scaledStyles.dropdown}
      />
      <DropDownPicker
        open={openInches}
        value={inch}
        items={inches}
        setOpen={setOpenInches}
        setValue={setInch}
        placeholder="Inches"
        containerStyle={scaledStyles.dropdown}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  dropDownContainer: {
    flexDirection: "row"
  },
  dropdown: {
    flex: 1,
    marginRight: 5
  }
});

const scaledStyles = scaleStyleSheet(styles);

export default DropDownHeightPicker;
